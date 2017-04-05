package model

import "time"

type Character struct {
	CharacterId   int64       `gorm:"primary_key"`
	CharacterName string      `gorm:"varchar(100)"`
	Corporation   Corporation `gorm:"ForeignKey:corporation_id;AssociationForeignKey:corporation_id;save_associations:false;"`
	CorporationId int64
	Token         string
	// this really isn't meant as a many to many but the flow of auth linking forces this or another method
	// which I didn't want to use
	Users      []User               `gorm:"many2many:user_character_map;"` //AssociationForeignKey:user_id;ForeignKey:character_id;"`
	AuthCodes  []AuthenticationCode `gorm:"ForeignKey:character_id"`
	InsertedDt time.Time
	UpdatedDt  time.Time
}
