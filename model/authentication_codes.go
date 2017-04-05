package model

type AuthenticationCode struct {
	CharacterId        int64 `gorm:"primary_key"`
	AuthenticationCode string
	Character          Character `gorm:"ForeignKey:CharacterId;save_associations:false;"`
	IsUsed             bool
}
