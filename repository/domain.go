package repository

import (
	"github.com/abaeve/auth-srv/model"
	"github.com/jinzhu/gorm"
)

type character struct {
	db *gorm.DB
}

type alliance struct {
	db *gorm.DB
}

//BGN Character accessor methods
func (chr *character) findByCharacterId ( characterId int64 ) model.Character {
	return nil
}

func (char *character) findByCharacterName ( characterName string ) model.Character {
	return nil
}
//END Character accessor methods

//BGN Alliance accessor methods
func (all *alliance) findByAllianceId ( allianceId int64 ) model.Alliance {
	return nil
}

func (all *alliance) findByAllianceName ( allianceName string ) model.Alliance {
	return nil
}
//END Alliance accessor methods
