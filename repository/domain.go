package repository

import (
	"errors"
	"fmt"
	"github.com/abaeve/auth-srv/model"
	"github.com/jinzhu/gorm"
	"time"
)

type AllianceRepository interface {
	Save(alliance *model.Alliance) error
	FindByAllianceId(allianceId int64) *model.Alliance
}

type CorporationRepository interface {
	Save(corporation *model.Corporation) error
	FindByCorporationId(corporationId int64) *model.Corporation
}

type CharacterRepository interface {
	Save(character *model.Character) error
	FindByCharacterId(characterId int64) *model.Character
	FindByAutenticationCode(authCode string) *model.Character
}

type UserRepository interface {
	Save(user *model.User) error
	FindByChatId(chatId string) *model.User
	LinkCharacterToUserByAuthCode(authCode string, user *model.User) error
}

type RoleRepository interface {
	Save(role *model.Role) error
}

type AuthenticationCodeRepository interface {
	Save(character *model.Character, authCode string) error
	//FindByAuthenticationCode(authCode string) *model.AuthenticationCode
	/*Excluding this from the interface for now, maybe we will need it later?
	FindByCharacterId(characterId int64) *model.AuthenticationCode
	*/
}

type allianceRepository struct {
	db *gorm.DB
}

type corporationRepository struct {
	db *gorm.DB
}

type characterRepository struct {
	db *gorm.DB
}

type userRepository struct {
	db *gorm.DB
}

type roleRepository struct {
	db *gorm.DB
}

type authCodeRepository struct {
	db *gorm.DB
}

//BGN AllianceRepo accessor methods
func (all *allianceRepository) Save(alliance *model.Alliance) error {
	//Apparently GORM sends no error back when a primary key isn't populated?  GARBAGE!
	if alliance.AllianceId == 0 {
		return errors.New("Primary key must not be 0")
	}

	if beforeEpoc(alliance.InsertedDt) {
		alliance.InsertedDt = newTimeNow()
	}

	alliance.UpdatedDt = newTimeNow()

	fmt.Printf("Alliance: (%+v)\n", alliance)

	err := all.db.Save(alliance).Error

	return err
}

func (all *allianceRepository) FindByAllianceId(allianceId int64) *model.Alliance {
	var alliance model.Alliance

	all.db.Where("alliance_id = ?", allianceId).Find(&alliance)
	if alliance.AllianceId == 0 {
		return nil
	}

	return &alliance
}
//END AllianceRepo accessor methods

//BGN Corporation accessor methods
func (corp *corporationRepository) Save(corporation *model.Corporation) error {
	if corporation.CorporationId == 0 {
		return errors.New("Primary key must not be 0")
	}

	if beforeEpoc(corporation.InsertedDt) {
		corporation.InsertedDt = newTimeNow()
	}

	corporation.UpdatedDt = newTimeNow()

	fmt.Printf("Corporation: (%+v)\n", corporation)

	err := corp.db.Save(corporation).Error

	return err
}

func (corp *corporationRepository) FindByCorporationId(corporationId int64) *model.Corporation {
	var corporation model.Corporation

	corp.db.Where("corporation_id = ?", corporationId).Find(&corporation)
	if corporation.CorporationId != 0 {
		corp.db.Model(&corporation).Association("Alliance").Find(&corporation.Alliance)
	} else {
		return nil
	}

	return &corporation
}
//END Corporation accessor methods

//BGN Character accessor methods
func (chr *characterRepository) Save(character *model.Character) error {
	if character.CharacterId == 0 {
		return errors.New("Primary key must not be 0")
	}

	if beforeEpoc(character.InsertedDt) {
		character.InsertedDt = newTimeNow()
	}

	character.UpdatedDt = newTimeNow()

	fmt.Printf("Character: (%+v)\n", character)

	err := chr.db.Save(character).Error
	return err
}

func (chr *characterRepository) FindByCharacterId(characterId int64) *model.Character {
	var character model.Character

	chr.db.Where("character_id = ?", characterId).Find(&character)
	if character.CharacterId != 0 {
		chr.db.Model(&character).Association("Corporation").Find(&character.Corporation)
		chr.db.Model(&character).Association("Users").Find(&character.Users)
		chr.db.Model(&character).Association("AuthCodes").Find(&character.AuthCodes)
	} else {
		return nil
	}

	return &character
}

func (chr *characterRepository) FindByAutenticationCode(authCode string) *model.Character {
	var authCodeModel model.AuthenticationCode
	var character model.Character

	chr.db.Where("authentication_code = ?", authCode).Find(&authCodeModel)

	if authCodeModel.CharacterId != 0 {
		chr.db.Model(&authCodeModel).Association("Character").Find(&character)
	} else {
		return nil
	}

	//character.AuthCodes = []model.AuthenticationCode{authCodeModel}

	return &character
}
//END Character accessor methods

//BGN User accessor methods
func (usr *userRepository) Save(user *model.User) error {
	err := usr.db.Save(&user).Error

	return err
}

func (usr *userRepository) FindByChatId(chatId string) *model.User {
	var user model.User

	usr.db.Where("chat_id = ?", chatId).Find(&user)

	if user.UserId == 0 {
		return nil
	}

	usr.db.Model(&user).Association("Characters").Find(&user.Characters)

	return &user
}

func (usr *userRepository) LinkCharacterToUserByAuthCode(authCode string, user *model.User) error {
	foundCharacter := CharacterRepo.FindByAutenticationCode(authCode)

	if foundCharacter == nil {
		return errors.New("No user with that auth code found")
	}

	//Ensure we have all the current associations
	usr.db.Model(&user).Association("Characters").Find(&user.Characters)

	//BGN Do I really have to do this to grow the size of an array?
	characters := user.Characters
	user.Characters = make([]model.Character, len(characters)+1)

	for idx, character := range characters {
		user.Characters[idx] = character
	}

	user.Characters[len(user.Characters)-1] = *foundCharacter
	//END Do I really have to do this to grow the size of an array?

	usr.db.Save(&user)

	//Now use up the auth code
	var authCodeModel model.AuthenticationCode
	//TODO: Refactor this into the auth code repo?usr.db.Where("authentication_code = ? and is_used = ?", authCode, false).Find(&authCodeModel)
	usr.db.Where("authentication_code = ?", authCode).Find(&authCodeModel)

	if authCodeModel.IsUsed {
		return errors.New("Authentication Code is invalid or used.")
	}

	authCodeModel.IsUsed = true

	usr.db.Save(&authCodeModel)

	return nil
}
//END User accessor methods

//BGN Role accessor methods
func (rle *roleRepository) Save(role *model.Role) error {
	if beforeEpoc(role.InsertedDt) {
		role.InsertedDt = newTimeNow()
	}

	role.UpdatedDt = newTimeNow()

	err := rle.db.Save(&role).Error
	return err
}
//END Role accessor methods

//BGN Authentication Code methods
func (ac *authCodeRepository) Save(character *model.Character, authCode string) error {
	newAuthCode := model.AuthenticationCode{AuthenticationCode: authCode, CharacterId: character.CharacterId, Character: *character, IsUsed: false}

	err := ac.db.Save(&newAuthCode).Error

	return err
}

//func (ac *authCodeRepository) FindByAuthenticationCode(authCode string) *model.AuthenticationCode {
//	return nil
//}
//END Authentication Code methods

func beforeEpoc(timeToCheck *time.Time) bool {
	epoch, _ := time.Parse(time.RFC822, "01 Jan 70 00:01 UTC")

	if timeToCheck == nil {
		return true
	}

	return timeToCheck.Before(epoch)
}

func newTimeNow() *time.Time {
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	return &now
}
