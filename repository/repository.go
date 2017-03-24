package repository

import "github.com/jinzhu/gorm"

var DB *gorm.DB

var CharacterRepo CharacterRepository
var AllianceRepo AllianceRepository
var Accesses *accesses

func Setup ( dialect string, connectionString string ) error {
	db, err := gorm.Open(dialect, connectionString)

	if err != nil {
		return err
	}

	DB = db

	CharacterRepo = &character{db: DB}
	AllianceRepo = &alliance{db: DB}
	Accesses = &accesses{db: DB.DB()}

	return nil
}