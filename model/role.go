package model

import "time"

type Role struct {
	RoleId           int64 `gorm:"primary_key;AUTO_INCREMENT"`
	RoleName         string `gorm:"varchar(70)"`
	ChatServiceGroup string `gorm:"varchar(70)"`
	InsertedDt       time.Time
	UpdatedDt        time.Time
}
