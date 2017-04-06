package model

import (
	"time"
)

type Corporation struct {
	CorporationId     int64    `gorm:"primary_key"`
	CorporationName   string   `gorm:"varchar(100)"`
	CorporationTicker string   `gorm:"varchar(5)"`
	Alliance          Alliance `gorm:"ForeignKey:alliance_id;AssociationForeignKey:alliance_id;save_associations:false;"`
	AllianceId        *int64
	InsertedDt        *time.Time
	UpdatedDt         *time.Time
}
