package model

import "time"

type Alliance struct {
	AllianceId     int64      `gorm:"primary_key" db:"alliance_id"`
	AllianceName   string     `gorm:"varchar(100)" db:"alliance_name"`
	AllianceTicker string     `gorm:"varchar(5)" db:"alliance_ticker"`
	InsertedDt     *time.Time `db:"inserted_dt"`
	UpdatedDt      *time.Time `db:"updated_dt"`
}
