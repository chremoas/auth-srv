package repository

import (
	"git.maurer-it.net/abaeve/auth-srv/model"
	"github.com/jinzhu/gorm"
)

//TODO: Make this generic enough that it won't only hack one specific many2many relationship
func hackJoinTableHandlerData(db *gorm.DB) {
	//<editor-fold desc="Stupid hack because the developer won't fix it">
	//Due to a bug being rejected this will never be fixed...
	//https://github.com/jinzhu/gorm/issues/707
	var userCharRel *gorm.Relationship
	var charUserRel *gorm.Relationship
	var userJoinTableHandler *gorm.JoinTableHandler
	var charJoinTableHandler *gorm.JoinTableHandler

	for _, field := range db.NewScope(model.User{}).GetStructFields() {
		if field.Name == "Characters" {
			userCharRel = field.Relationship // struct contains foreign keys informations
			userJoinTableHandler = field.Relationship.JoinTableHandler.(*gorm.JoinTableHandler)
		}
	}

	for _, field := range db.NewScope(model.Character{}).GetStructFields() {
		if field.Name == "Users" {
			charUserRel = field.Relationship
			charJoinTableHandler = field.Relationship.JoinTableHandler.(*gorm.JoinTableHandler)
		}
	}

	userCharRel.ForeignDBNames = []string{"user_id"}
	userCharRel.AssociationForeignDBNames = []string{"character_id"}
	charUserRel.ForeignDBNames = []string{"character_id"}
	charUserRel.AssociationForeignDBNames = []string{"user_id"}

	userJoinTableHandler.Source.ForeignKeys[0].DBName = "user_id"
	userJoinTableHandler.Destination.ForeignKeys[0].DBName = "character_id"
	charJoinTableHandler.Source.ForeignKeys[0].DBName = "character_id"
	charJoinTableHandler.Destination.ForeignKeys[0].DBName = "user_id"
	//</editor-fold>
}
