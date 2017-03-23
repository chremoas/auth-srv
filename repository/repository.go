package repository

import "github.com/jinzhu/gorm"

var DB *gorm.DB

var Character *character
var Alliance *alliance
var Accesses *accesses

func Setup ( dialect string, connectionString string ) error {
	DB,err := gorm.Open(dialect, connectionString)

	if err != nil {
		return nil
	}

	Character = &character{db: DB}
	Alliance = &alliance{db: DB}
	Accesses = &accesses{db: DB.DB()}

	return nil
}