package repository

import (
	"github.com/abaeve/auth-srv/model"
	"github.com/jinzhu/gorm"
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
	FindByCharacterId(characterId int64) *model.AuthenticationCode
}

type alliance struct {
	db *gorm.DB
}

type corporation struct {
	db *gorm.DB
}

type character struct {
	db *gorm.DB
}

type user struct {
	db *gorm.DB
}

type role struct {
	db *gorm.DB
}

type authCodes struct {
	db *gorm.DB
}

type daoError struct {
	message string
}

//BGN AllianceRepo accessor methods
func (all *alliance) Save(alliance *model.Alliance) error {
	//Apparently GORM sends no error back when a primary key isn't populated?  GARBAGE!
	if alliance.AllianceId == 0 {
		return &daoError{message: "Primary key must not be 0"}
	}

	err := all.db.Save(alliance).Error

	return err
}

func (all *alliance) FindByAllianceId(allianceId int64) *model.Alliance {
	var alliance model.Alliance

	all.db.Where("alliance_id = ?", allianceId).Find(&alliance)

	return &alliance
}

func (all *alliance) findByAllianceName(allianceName string) *model.Alliance {
	return &model.Alliance{}
}
//END AllianceRepo accessor methods

//BGN Corporation accessor methods
func (corp *corporation) Save(corporation *model.Corporation) error {
	if corporation.CorporationId == 0 {
		return &daoError{message: "Primary key must not be 0"}
	}

	err := corp.db.Save(corporation).Error
	return err
}

func (corp *corporation) FindByCorporationId(corporationId int64) *model.Corporation {
	var corporation model.Corporation

	corp.db.Where("corporation_id = ?", corporationId).Find(&corporation)
	corp.db.Model(&corporation).Association("Alliance").Find(&corporation.Alliance)

	return &corporation
}
//END Corporation accessor methods

//BGN Character accessor methods
func (chr *character) Save(character *model.Character) error {
	err := chr.db.Save(character).Error
	return err
}

func (chr *character) FindByCharacterId(characterId int64) *model.Character {
	var character model.Character

	chr.db.Where("character_id = ?", characterId).Find(&character)
	chr.db.Model(&character).Association("Corporation").Find(&character.Corporation)
	chr.db.Model(&character).Association("Users").Find(&character.Users)
	chr.db.Model(&character).Association("AuthCodes").Find(&character.AuthCodes)

	return &character
}

func (chr *character) FindByAutenticationCode(authCode string) *model.Character {
	var authCodeModel model.AuthenticationCode
	var character model.Character

	chr.db.Where("authentication_code = ?", authCode).Find(&authCodeModel)

	if authCodeModel.CharacterId != 0 {
		chr.db.Model(&authCodeModel).Association("Character").Find(&character)
	}

	//character.AuthCodes = []model.AuthenticationCode{authCodeModel}

	return &character
}

func (chr *character) findByCharacterName(characterName string) *model.Character {
	return &model.Character{}
}
//END Character accessor methods

//BGN User accessor methods
func (usr *user) Save(user *model.User) error {
	return nil
}

func (usr *user) FindByChatId(chatId string) *model.User {
	return &model.User{}
}

func (usr *user) LinkCharacterToUserByAuthCode(authCode string, user *model.User) error {
	return nil
}
//END User accessor methods

//BGN Role accessor methods
func (rle *role) Save(role *model.Role) error {
	return nil
}
//END Role accessor methods

//BGN Authentication Code methods
func (ac *authCodes) Save(character *model.Character, authCode string) error {
	return nil
}
func (ac *authCodes) FindByCharacterId(characterId int64) *model.AuthenticationCode {
	return &model.AuthenticationCode{}
}
//END Authentication Code methods

//Make daoError implement error
func (err *daoError) Error() string {
	return err.message
}
