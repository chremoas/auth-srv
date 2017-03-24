package repository

import (
	"github.com/abaeve/auth-srv/model"
	"github.com/jinzhu/gorm"
)

type AllianceRepository interface {
	Save(alliance model.Alliance) error
	FindByAllianceId(allianceId int64) model.Alliance
}

type CorporationRepository interface {
	Save(corporation model.Corporation) error
	FindByCorporationId(corporationId int64) model.Corporation
}

type CharacterRepository interface {
	Save(character model.Character) error
	FindByCharacterId(characterId int64) model.Character
	FindByAutenticationCode(authCode string) model.Character
}

type UserRepository interface {
	Save(user model.User) error
	FindByChatId(chatId string) model.User
	LinkCharacterToUser(character model.Character, user model.User) error
}
type RoleRepository interface {
	Save(role model.Role) error
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

//BGN AllianceRepo accessor methods
func (all *alliance) Save(alliance model.Alliance) error {
	return nil
}

func (all *alliance) FindByAllianceId(allianceId int64) model.Alliance {
	return model.Alliance{}
}

func (all *alliance) findByAllianceName ( allianceName string ) model.Alliance {
	return model.Alliance{}
}
//END AllianceRepo accessor methods

//BGN Corporation accessor methods
func (corp *corporation) Save(corporation model.Corporation) error {
	return nil
}

func (corp *corporation) FindByCorporationId(corporationId int64) model.Corporation {
	return model.Corporation{}
}
//END Corporation accessor methods

//BGN Character accessor methods
func (chr *character) Save(character model.Character) error {
	return nil
}

func (chr *character) FindByCharacterId(characterId int64) model.Character {
	return model.Character{}
}

func (chr *character) FindByAutenticationCode(authCode string) model.Character {
	return model.Character{}
}

func (chr *character) findByCharacterName ( characterName string ) model.Character {
	return model.Character{}
}
//END Character accessor methods

//BGN User accessor methods
func (usr *user) Save(user model.User) error {
	return nil
}

func (usr *user) FindByChatId(chatId string) model.User {
	return model.User{}
}

func (usr *user) LinkCharacterToUser(character model.Character, user model.User) error {
	return nil
}
//END User accessor methods

//BGN Role accessor methods
func (rle *role) Save(role model.Role) error {
	return nil
}
//END Role accessor methods
