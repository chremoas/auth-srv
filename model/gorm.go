package model

func (Alliance) TableName() string {
	return "alliance"
}

func (Corporation) TableName() string {
	return "corporation"
}

func (Character) TableName() string {
	return "character"
}

func (Role) TableName() string {
	return "role"
}

func (User) TableName() string {
	return "user"
}
