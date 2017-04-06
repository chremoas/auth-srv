package model

import "time"

type Alliance struct {
	AllianceId     int64  `gorm:"primary_key"`
	AllianceName   string `gorm:"varchar(100)"`
	AllianceTicker string `gorm:"varchar(5)"`
	InsertedDt     *time.Time
	UpdatedDt      *time.Time
}
