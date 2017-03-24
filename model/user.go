package model

type User struct {
	UserId     int64 `gorm:"primary_key"`
	ChatId     string `gorm:"varchar(255)"`
	Characters []Character `gorm:"many2many:user_character_map;"` //AssociationForeignKey:character_id;ForeignKey:user_id;"`
}
