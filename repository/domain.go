package repository

import (
	"errors"
	"fmt"
	"github.com/chremoas/auth-srv/model"
	"github.com/chremoas/auth-srv/util"
	"github.com/jinzhu/gorm"
)

type AllianceRepository interface {
	Save(alliance *model.Alliance) error
	FindByAllianceId(allianceId int64) *model.Alliance
	FindAll() []*model.Alliance
	Delete(allianceId int64) error
}

type CorporationRepository interface {
	Save(corporation *model.Corporation) error
	FindByCorporationId(corporationId int64) *model.Corporation
	FindAll() []*model.Corporation
	Delete(corporationId int64) error
}

type CharacterRepository interface {
	Save(character *model.Character) error
	FindByCharacterId(characterId int64) *model.Character
	FindByAutenticationCode(authCode string) *model.Character
	FindAll() []*model.Character
	Delete(characterId int64) error
}

type UserRepository interface {
	Save(user *model.User) error
	FindByChatId(chatId string) *model.User
	LinkCharacterToUserByAuthCode(authCode string, user *model.User) error
}

type RoleRepository interface {
	Save(role *model.Role) error
	FindByRoleName(roleName string) *model.Role
	FindAll() []*model.Role
	Delete(roleName string) error
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

	if util.BeforeEpoc(alliance.InsertedDt) {
		alliance.InsertedDt = util.NewTimeNow()
	}

	alliance.UpdatedDt = util.NewTimeNow()

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

func (all *allianceRepository) FindAll() []*model.Alliance {
	alliances := []*model.Alliance{}

	all.db.Find(&alliances)

	return alliances
}

func (all *allianceRepository) Delete(allianceId int64) error {
	err := all.db.Exec("update corporations set alliance_id = null where alliance_id = ?", allianceId).Error
	if err != nil {
		return err
	}

	result := all.db.Where("alliance_id = ?", allianceId).Delete(&model.Alliance{})

	return result.Error
}

//END AllianceRepo accessor methods

//BGN Corporation accessor methods
func (corp *corporationRepository) Save(corporation *model.Corporation) error {
	if corporation.CorporationId == 0 {
		return errors.New("Primary key must not be 0")
	}

	if util.BeforeEpoc(corporation.InsertedDt) {
		corporation.InsertedDt = util.NewTimeNow()
	}

	corporation.UpdatedDt = util.NewTimeNow()

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

func (corp *corporationRepository) FindAll() []*model.Corporation {
	var corporations []*model.Corporation

	corp.db.Find(&corporations)

	return corporations
}

func (corp *corporationRepository) Delete(corporationId int64) error {
	// FIX: This doesn't seem to have anything left now that that the role maps are gone. -brian
	return corp.db.Where("corporation_id = ?", corporationId).Delete(&model.Corporation{}).Error
}

//END Corporation accessor methods

//BGN Character accessor methods
func (chr *characterRepository) Save(character *model.Character) error {
	if character.CharacterId == 0 {
		return errors.New("Primary key must not be 0")
	}

	if util.BeforeEpoc(character.InsertedDt) {
		character.InsertedDt = util.NewTimeNow()
	}

	character.UpdatedDt = util.NewTimeNow()

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

func (chr *characterRepository) FindAll() []*model.Character {
	characters := []*model.Character{}

	chr.db.Find(&characters)

	return characters
}

func (chr *characterRepository) Delete(characterId int64) error {
	tx := chr.db.Begin()
	err := chr.db.Where("character_id = ?", characterId).Delete(&model.AuthenticationCode{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = chr.db.Exec("delete from user_character_map where character_id = ?", characterId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	defer tx.Commit()

	return chr.db.Where("character_id = ?", characterId).Delete(&model.Character{}).Error
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
	//First we check that we haven't already consumed this auth code
	var authCodeModel model.AuthenticationCode
	//TODO: Refactor this into the auth code repo?usr.db.Where("authentication_code = ? and is_used = ?", authCode, false).Find(&authCodeModel)
	usr.db.Where("authentication_code = ?", authCode).Find(&authCodeModel)

	if authCodeModel.IsUsed {
		return errors.New("Authentication Code is invalid or used.")
	}

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
	authCodeModel.IsUsed = true

	usr.db.Save(&authCodeModel)

	return nil
}

//END User accessor methods

//BGN Role accessor methods
func (rle *roleRepository) Save(role *model.Role) error {
	if util.BeforeEpoc(role.InsertedDt) {
		role.InsertedDt = util.NewTimeNow()
	}

	role.UpdatedDt = util.NewTimeNow()

	err := rle.db.Save(&role).Error
	return err
}

func (rle *roleRepository) FindByRoleName(roleName string) *model.Role {
	var role model.Role

	rle.db.Where("role_name = ?", roleName).Find(&role)

	return &role
}

func (rle *roleRepository) FindAll() []*model.Role {
	roles := []*model.Role{}

	rle.db.Find(&roles)

	return roles
}

func (rle *roleRepository) Delete(roleName string) error {
	result := rle.db.Where("role_name = ?", roleName).Delete(&model.Role{})

	if result.RowsAffected == 0 {
		return fmt.Errorf("No such role: %s", roleName)
	}

	return nil
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
