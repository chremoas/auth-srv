package repository

import (
	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/go-sql-driver/mysql"
	"testing"
	"github.com/abaeve/auth-srv/model"
	"github.com/jinzhu/gorm"
	"github.com/abaeve/auth-srv/util"
)

func TestCreateAndRetrieveThroughGORM(t *testing.T) {
	//<editor-fold desc="Setup code">
	err := Setup("sqlite3", "file:../test/authSrv.db?loc=auto")
	/*err := Setup("mysql", "root@tcp(localhost:3306)/microservices")*/

	if err != nil {
		t.Fatalf("Could not open database connection: %s", err)
	}

	DB.LogMode(true)
	util.HackJoinTableHandlerData(DB)

	alliance := model.Alliance{AllianceId: 1, AllianceTicker: "TST", AllianceName: "Test AllianceRepo 1"}
	corporation := model.Corporation{CorporationId: 1, AllianceId: 1, CorporationName: "Test Corporation 1", CorporationTicker: "TST", Alliance: alliance}
	character := model.Character{CharacterId: 1, CharacterName: "Test Character 1", CorporationId: 1, Corporation: corporation, Token: ""}
	user := model.User{UserId: 1, ChatId: "1234567890", Characters: []model.Character{character}}

	DB.Create(&alliance)
	DB.Create(&corporation)
	DB.Create(&character)
	DB.Create(&user)
	//</editor-fold>

	t.Run("Character-Retrieve", func(t *testing.T) {
		tx := DB.Begin()

		var characterAsRetrieved model.Character
		var corporationAsRetrieved model.Corporation

		tx.First(&characterAsRetrieved).Related(&corporationAsRetrieved)

		if character.Corporation.CorporationId != corporationAsRetrieved.CorporationId {
			t.Fatalf("Retrieved character's corporation (%d) doesn't equal original characters corporation: (%d)",
				corporationAsRetrieved.CorporationId, character.Corporation.CorporationId)
		}

		if character.CorporationId != characterAsRetrieved.CorporationId {
			t.Fatalf("Retrieved character's corporation id: (%d) doesn't equal origin characters corporation: (%d)",
				characterAsRetrieved.CorporationId, character.CorporationId)
		}

		tx.Rollback()
	})

	t.Run("Corporation-Count", func(t *testing.T) {
		tx := DB.Begin()

		rows, err := DB.DB().Query("select count(*) from corporations")

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

		tx.Rollback()
	})

	t.Run("Character-NewAndAttachToUser", func(t *testing.T) {
		tx := DB.Begin()

		var userAsRetrieved model.User
		newCharacter := model.Character{CharacterId: 2, CharacterName: "Test Character 2", CorporationId: 1, Corporation: corporation, Token: ""}

		tx.Create(&newCharacter)

		tx.Where("user_id = ?", 1).First(&userAsRetrieved)

		tx.Model(&user).Association("Characters").Append(newCharacter)

		charCount := tx.Model(&user).Association("Characters").Count()

		if charCount != 2{
			t.Fatalf("Expected 2 characters, only got %d", charCount)
		}

		err = tx.Model(&user).Association("Characters").Delete(&newCharacter).Error

		if err != nil {
			t.Fatalf("Could not remove the new character -> user association: %s", err)
		}

		rowsAffected := tx.Delete(&newCharacter).RowsAffected

		if rowsAffected != 1 {
			t.Fatalf("Did not remove the character, removed %d records", rowsAffected)
		}

		tx.Rollback()
	})

	t.Run("User-Relationships", func(t *testing.T) {
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

		if userCharRel.ForeignDBNames[0] != "user_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected user_id but got %s", userCharRel.ForeignDBNames)
		}

		if userCharRel.ForeignFieldNames[0] != "user_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected user_id but got %s", userCharRel.ForeignFieldNames)
		}

		if userCharRel.AssociationForeignDBNames[0] != "character_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected character_id but got %s", userCharRel.AssociationForeignDBNames)
		}

		if userCharRel.AssociationForeignFieldNames[0] != "character_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected character_id but got %s", userCharRel.AssociationForeignFieldNames)
		}

		if charUserRel.ForeignDBNames[0] != "character_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected character_id but got %s", charUserRel.ForeignDBNames)
		}

		if charUserRel.ForeignFieldNames[0] != "character_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected character_id but got %s", charUserRel.ForeignFieldNames)
		}

		if charUserRel.AssociationForeignDBNames[0] != "user_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected user_id but got %s", charUserRel.AssociationForeignDBNames)
		}

		if charUserRel.AssociationForeignFieldNames[0] != "user_id" {
			t.Fatalf("User Character many2many had wrong mapping, expected user_id but got %s", charUserRel.AssociationForeignFieldNames)
		}
	})

	t.Run("User-JoinTableHandlers", func(t *testing.T) {
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

		if userJoinTableHandler.Source.ForeignKeys[0].DBName != "user_id" {
			t.Fatalf("user join table handler has wrong value for user_id association column: %s",
				userJoinTableHandler.Source.ForeignKeys[0].DBName)
		}

		if userJoinTableHandler.Destination.ForeignKeys[0].DBName != "character_id" {
			t.Fatalf("user join table handler has wrong value for character_id association column: %s",
				userJoinTableHandler.Destination.ForeignKeys[0].DBName)
		}

		if charJoinTableHandler.Source.ForeignKeys[0].DBName != "character_id" {
			t.Fatalf("character join table handler has wrong value for character_id association column: %s",
				charJoinTableHandler.Source.ForeignKeys[0].DBName)
		}

		if charJoinTableHandler.Destination.ForeignKeys[0].DBName != "user_id" {
			t.Fatalf("character join table handler has wrong value for user_id association column: %s",
				charJoinTableHandler.Destination.ForeignKeys[0].DBName)
		}
	})

	//<editor-fold desc="Teardown Code">
	tx := DB.Begin()

	err = tx.Model(&user).Association("Characters").Delete(&character).Error

	if err != nil {
		t.Errorf("Could not remove user -> character association: (%s)", err)
	}

	rowsAffected := tx.Delete(&user).RowsAffected

	if rowsAffected != 1 {
		t.Errorf("Did not remove the user, removed %d records", rowsAffected)
	}

	rowsAffected = tx.Delete(&character).RowsAffected

	if rowsAffected != 1 {
		t.Errorf("Did not remove the character, removed %d records", rowsAffected)
	}

	rowsAffected = tx.Delete(&corporation).RowsAffected

	if rowsAffected != 1 {
		t.Errorf("Did not remove the corporation, removed %d records", rowsAffected)
	}

	rowsAffected = tx.Delete(&alliance).RowsAffected

	if rowsAffected != 1 {
		t.Errorf("Did not remove the alliance, removed %d records", rowsAffected)
	}

	tx.Commit()
	//</editor-fold>
}

func TestCreateAndRetrieveThroughREPO(t *testing.T) {

}
