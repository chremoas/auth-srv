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
	Alliance          Alliance `gorm:"ForeignKey:alliance_id;AssociationForeignKey:alliance_id"`
	AllianceId        int64
	InsertedDt        time.Time
	UpdatedDt         time.Time
}

type Character struct {
	CharacterId   int64 `gorm:"primary_key"`
	CharacterName string `gorm:"varchar(100)"`
	Corporation   Corporation `gorm:"ForeignKey:corporation_id;AssociationForeignKey:corporation_id"`
	CorporationId int64
}

type Role struct {
	RoleId           int64 `gorm:"primary_key"`
	RoleName         string `gorm:"varchar(70)"`
	ChatServiceGroup string `gorm:"varchar(70)"`
	InsertedDt       time.Time
	UpdatedDt        time.Time
}
