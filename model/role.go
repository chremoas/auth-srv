package model

import "time"

type Role struct {
	RoleId           int64  `gorm:"primary_key"`
	RoleName         string `gorm:"varchar(70)"`
	ChatServiceGroup string `gorm:"varchar(70);column:chatservice_group"`
	InsertedDt       *time.Time
	UpdatedDt        *time.Time
}
