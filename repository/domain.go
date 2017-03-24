package repository

import (
	"github.com/abaeve/auth-srv/model"
	"github.com/jinzhu/gorm"
)

type CharacterRepository interface {
	FindByCharacterId(characterId int64) model.Character
}

type AllianceRepository interface {
	FindByAllianceId(allianceId int64) model.Alliance
}

type character struct {
	db *gorm.DB
}

type alliance struct {
	db *gorm.DB
}

//BGN Character accessor methods
func (chr *character) FindByCharacterId(characterId int64) model.Character {
	return model.Character{}
}

func (char *character) findByCharacterName ( characterName string ) model.Character {
	return model.Character{}
}
//END Character accessor methods

//BGN AllianceRepo accessor methods
func (all *alliance) FindByAllianceId(allianceId int64) model.Alliance {
	return model.Alliance{}
}

func (all *alliance) findByAllianceName ( allianceName string ) model.Alliance {
	return model.Alliance{}
}

//END AllianceRepo accessor methods
