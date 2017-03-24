package repository

import (
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"github.com/abaeve/auth-srv/model"
	"github.com/jinzhu/gorm"
	"fmt"
	"encoding/json"
	"github.com/abaeve/auth-srv/util"
)

func TestCreateAndRetrieveThroughGORM(t *testing.T) {
	//<editor-fold desc="Setup code">
	err := Setup("sqlite3", "file:../test/authSrv.db?loc=auto")

	if err != nil {
		t.Fatalf("Could not open database connection: %s", err)
	}

	util.HackJoinTableHandlerData(DB)

	alliance := model.Alliance{AllianceId: 1, AllianceTicker: "TST", AllianceName: "Test AllianceRepo 1"}
	corporation := model.Corporation{CorporationId: 1, AllianceId: 1, CorporationName: "Test Corporation 1", CorporationTicker: "TST", Alliance: alliance}
	character := model.Character{CharacterId: 1, CharacterName: "Test Character 1", CorporationId: 1, Corporation: corporation, Token: ""}
	user := model.User{UserId: 1, ChatId: "1234567890", Characters: []model.Character{character}}

	DB.Save(&alliance)
	DB.Save(&corporation)
	DB.Save(&character)
	//DB.Save(&user)
	//</editor-fold>

	t.Run("Character-Retrieve", func(t *testing.T) {
		var characterAsRetrieved model.Character
		var corporationAsRetrieved model.Corporation

		DB.First(&characterAsRetrieved).Related(&corporationAsRetrieved)

		if character.Corporation.CorporationId != corporationAsRetrieved.CorporationId {
			t.Fatalf("Retrieved character's corporation (%d) doesn't equal original characters corporation: (%d)",
				corporationAsRetrieved.CorporationId, character.Corporation.CorporationId)
		}

		if character.CorporationId != characterAsRetrieved.CorporationId {
			t.Fatalf("Retrieved character's corporation id: (%d) doesn't equal origin characters corporation: (%d)",
				characterAsRetrieved.CorporationId, character.CorporationId)
		}
	})

	t.Run("Corporation-Count", func(t *testing.T) {
		rows, err := DB.DB().Query("select count(*) from corporation")

		if err != nil {
			t.Fatal("Raw corporation count query had an issue")
		}

		if rows.Next() {
			var count int64
			rows.Scan(&count)

			if count != 1 {
				t.Fatalf("Too few or too many corporations were found: %d", count)
			}
		} else {
			t.Fatal("Raw corporation count query returned rows")
		}
	})

	t.Run("Character-NewAndAttachToUser", func(t *testing.T) {
		//DB.Where
	})

	t.Run("User-PrintRelationships", func(t *testing.T) {
		var userCharRel *gorm.Relationship
		var charUserRel *gorm.Relationship

		for _, field := range DB.NewScope(model.User{}).GetStructFields() {
			if field.Name == "Characters" {
				userCharRel = field.Relationship // struct contains foreign keys informations
			}
		}

		for _, field := range DB.NewScope(model.Character{}).GetStructFields() {
			if field.Name == "Users" {
				charUserRel = field.Relationship
			}
		}

		if userCharRel.ForeignDBNames[0] != "character_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected character_id but got %s", userCharRel.ForeignDBNames)
		}

		if userCharRel.ForeignFieldNames[0] != "user_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected user_id but got %s", userCharRel.ForeignDBNames)
		}

		if userCharRel.AssociationForeignDBNames[0] != "user_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected user_id but got %s", userCharRel.ForeignDBNames)
		}

		if userCharRel.AssociationForeignFieldNames[0] != "character_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected character_id but got %s", userCharRel.ForeignDBNames)
		}

		if charUserRel.ForeignDBNames[0] != "user_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected user_id but got %s", userCharRel.ForeignDBNames)
		}

		if charUserRel.ForeignFieldNames[0] != "character_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected character_id but got %s", userCharRel.ForeignDBNames)
		}

		if charUserRel.AssociationForeignDBNames[0] != "character_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected character_id but got %s", userCharRel.ForeignDBNames)
		}

		if charUserRel.AssociationForeignFieldNames[0] != "user_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected user_id but got %s", userCharRel.ForeignDBNames)
		}
	})

	t.Run("User-PrintJoinTableHandlers", func(t *testing.T) {
		var userJoinTableHandler *gorm.JoinTableHandler
		var charJoinTableHandler *gorm.JoinTableHandler

		for _, field := range DB.NewScope(model.User{}).GetStructFields() {
			if field.Name == "Characters" {
				userJoinTableHandler = field.Relationship.JoinTableHandler.(*gorm.JoinTableHandler)
			}
		}

		for _, field := range DB.NewScope(model.Character{}).GetStructFields() {
			if field.Name == "Users" {
				charJoinTableHandler = field.Relationship.JoinTableHandler.(*gorm.JoinTableHandler)
			}
		}

		userJoinTableHandlerData, _ := json.Marshal(userJoinTableHandler)
		charJoinTableHandlerData, _ := json.Marshal(charJoinTableHandler)

		//fmt.Println(userJoinTableHandler)
		fmt.Println(string(userJoinTableHandlerData))
		fmt.Println(string(charJoinTableHandlerData))
	})

	//<editor-fold desc="Teardown Code">
	tx := DB.Begin()

	tx.Delete(&user)
	tx.Delete(&character)
	tx.Delete(&corporation)
	tx.Delete(&alliance)

	tx.Commit()
	//</editor-fold>
}

func TestCreateAndRetrieveThroughREPO(t *testing.T) {

}

func TestPrintUserForeignKeys(t *testing.T) {

}
