package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
)

var DB *gorm.DB

var AllianceRepo AllianceRepository
var CorporationRepo CorporationRepository
var CharacterRepo CharacterRepository
var UserRepo UserRepository
var RoleRepo RoleRepository
var AccessRepo AccessesRepository
var AuthenticationCodeRepo AuthenticationCodeRepository

func Setup(dialect string, connectionString string) error {
	db, err := gorm.Open(dialect, connectionString)

	if err != nil {
		return err
	}

	DB = db

	hackJoinTableHandlerData(DB)

	AllianceRepo = &allianceRepository{db: DB}
	CorporationRepo = &corporationRepository{db: DB}
	CharacterRepo = &characterRepository{db: DB}
	UserRepo = &userRepository{db: DB}
	RoleRepo = &roleRepository{db: DB}
	AccessRepo = &accessesRepo{db: sqlx.NewDb(DB.DB(), dialect)}
	AuthenticationCodeRepo = &authCodeRepository{db: DB}

	return nil
}
