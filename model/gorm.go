package model

import ()
import "time"

type Alliance struct {
	AllianceId     int64 `gorm:"primary_key"`
	AllianceName   string `gorm:"varchar(100)"`
	AlliacneTicker string `gorm:"varchar(5)"`
	InsertedDt     time.Time
	UpdatedDt      time.Time
}

type Corporation struct {
	CorporationId     int64 `gorm:"primary_key"`
	CorporationName   string `gorm:"varchar(100)"`
	CorporationTicket string `gorm:"varchar(5)"`
	Alliance          Alliance `gorm:"ForeignKey:AllianceId;AssociationForeignKey:AllianceId"`
	AllianceId        int64
	InsertedDt        time.Time
	UpdatedDt         time.Time
}

type Character struct {
	CharacterId   int64 `gorm:"primary_key"`
	CharacterName string `gorm:"varchar(100)"`
	Corporation   Corporation `gorm:"ForeignKey:CorporationId;AssociationForeignKey:CorporationId"`
	CorporationId int64
	// this really isn't meant as a many to many but the flow of auth linking forces this or another method
	// which I didn't want to use
	User []User `gorm:"many2many:user_character_map"`
}

type Role struct {
	RoleId           int64 `gorm:"primary_key;AUTO_INCREMENT"`
	RoleName         string `gorm:"varchar(70)"`
	ChatServiceGroup string `gorm:"varchar(70)"`
	InsertedDt       time.Time
	UpdatedDt        time.Time
}

type User struct {
	UserId int64 `gorm:"primary_key"`
	ChatId string `gorm:"varchar(255)"`
	Characters []Character `gorm:"many2many:user_character_map;"`
}
