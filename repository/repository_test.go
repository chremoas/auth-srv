package repository

import (
	"github.com/chremoas/auth-srv/model"
	"github.com/chremoas/auth-srv/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func SharedSetup(t *testing.T) (model.Alliance, model.Corporation, [2]model.Character, model.User, [2]model.AuthenticationCode) {
	//<editor-fold desc="Setup code">
	err := Setup("mysql", "root@tcp(localhost:3306)/authservices?parseTime=true")

	if err != nil {
		t.Fatalf("Could not open database connection: %s", err)
	}

	DB.LogMode(true)

	var allianceId int64
	allianceId = 1

	alliance := model.Alliance{
		AllianceId:     1,
		AllianceTicker: "TST",
		AllianceName:   "Test AllianceRepo 1",
		InsertedDt:     util.NewTimeNow(),
		UpdatedDt:      util.NewTimeNow(),
	}
	corporation := model.Corporation{
		CorporationId:     1,
		AllianceId:        &allianceId,
		CorporationName:   "Test Corporation 1",
		CorporationTicker: "TST",
		Alliance:          alliance,
		InsertedDt:        util.NewTimeNow(),
		UpdatedDt:         util.NewTimeNow(),
	}
	character := model.Character{
		CharacterId:   1,
		CharacterName: "Test Character 1",
		CorporationId: 1,
		Corporation:   corporation,
		Token:         "",
		InsertedDt:    util.NewTimeNow(),
		UpdatedDt:     util.NewTimeNow(),
	}
	character10k := model.Character{
		CharacterId:   10000,
		CharacterName: "I'm lonely, auth me",
		CorporationId: 1,
		Corporation:   corporation,
		Token:         "",
		InsertedDt:    util.NewTimeNow(),
		UpdatedDt:     util.NewTimeNow(),
	}
	user := model.User{UserId: 1, ChatId: "1234567890", Characters: []model.Character{character}}
	authCode := model.AuthenticationCode{CharacterId: 1, Character: character, AuthenticationCode: "123456789", IsUsed: true}
	authCode2 := model.AuthenticationCode{CharacterId: 10000, Character: character10k, AuthenticationCode: "abcdefghijk", IsUsed: false}

	DB.Create(&alliance)
	DB.Create(&corporation)
	DB.Create(&character)
	DB.Create(&character10k)
	DB.Create(&user)
	DB.Create(&authCode)
	DB.Create(&authCode2)
	//</editor-fold>

	return alliance, corporation, [2]model.Character{character, character10k}, user, [2]model.AuthenticationCode{authCode, authCode2}
}

func SharedTearDown() {
	//<editor-fold desc="Teardown Code">
	/*TODO: WTF?*/
	/*
		This works... but doesn't take into account all the data created by whatever test calls this teardown
		DB.Delete(&authCode[0])
		DB.Delete(&authCode[1])
		DB.Model(&user).Association("Characters").Delete(&character)
		DB.Delete(&user)
		DB.Delete(&character[0])
		DB.Delete(&character[1])
		DB.Delete(&corporation)
		DB.Delete(&alliance)
	*/
	/*This doesn't?  I need this to work to clean up the other testing cases... or do I?
	tx, _ := DB.DB().Begin()
	t.Log("Delete authentication_code entries")
	tx.Exec("delete from authentication_codes")
	t.Log("Deleting user_character_map entries")
	tx.Exec("delete from user_character_map")
	t.Log("Deleting user entries")
	tx.Exec("delete from users")
	t.Log("Deleting character entries")
	tx.Exec("delete from characters")
	t.Log("Deleting corporation entries")
	tx.Exec("delete from corporations")
	t.Log("Deleting alliance entries")
	tx.Exec("delete from alliacnes")
	tx.Commit()*/

	var authCodes []model.AuthenticationCode
	var users []model.User
	var characters []model.Character
	var corporations []model.Corporation
	var alliances []model.Alliance
	var roles []model.Role

	tx := DB.Begin()

	tx.Find(&authCodes)
	tx.Find(&users)
	tx.Find(&characters)
	tx.Find(&corporations)
	tx.Find(&alliances)
	tx.Find(&roles)

	for _, authCode := range authCodes {
		tx.Delete(&authCode)
	}

	for _, user := range users {
		tx.Model(&user).Association("Characters").Find(&user.Characters)

		for _, character := range user.Characters {
			tx.Model(&user).Association("Characters").Delete(&character)
		}

		tx.Delete(&user)
	}

	for _, character := range characters {
		tx.Delete(&character)
	}

	for _, corporation := range corporations {
		tx.Delete(&corporation)
	}

	for _, alliance := range alliances {
		tx.Delete(&alliance)
	}

	for _, role := range roles {
		tx.Delete(&role)
	}

	tx.Commit()
	//</editor-fold>
}

func Test_CreateAndRetrieve_ThroughGORM(t *testing.T) {
	var err error
	_, corporation, character, user, _ := SharedSetup(t)

	t.Run("Character-Retrieve", func(t *testing.T) {
		tx := DB.Begin()

		var characterAsRetrieved model.Character
		var corporationAsRetrieved model.Corporation

		tx.First(&characterAsRetrieved).Related(&corporationAsRetrieved)

		if character[0].Corporation.CorporationId != corporationAsRetrieved.CorporationId {
			t.Fatalf("Retrieved character's corporation (%d) doesn't equal original characters corporation: (%d)",
				corporationAsRetrieved.CorporationId, character[0].Corporation.CorporationId)
		}

		if character[0].CorporationId != characterAsRetrieved.CorporationId {
			t.Fatalf("Retrieved character's corporation id: (%d) doesn't equal origin characters corporation: (%d)",
				characterAsRetrieved.CorporationId, character[0].CorporationId)
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

	t.Run("Character-CreateAndAttachToUser", func(t *testing.T) {
		tx := DB.Begin()

		var userAsRetrieved model.User
		newCharacter := model.Character{
			CharacterId:   2,
			CharacterName: "Test Character 2",
			CorporationId: 1,
			Corporation:   corporation,
			Token:         "",
			InsertedDt:    util.NewTimeNow(),
			UpdatedDt:     util.NewTimeNow(),
		}

		tx.Create(&newCharacter)

		tx.Where("user_id = ?", 1).First(&userAsRetrieved)

		tx.Model(&user).Association("Characters").Append(newCharacter)

		charCount := tx.Model(&user).Association("Characters").Count()

		if charCount != 2 {
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

	SharedTearDown()
}

func Test_AlliancesCRUD_ThroughREPO(t *testing.T) {
	alliance, _, _, _, _ := SharedSetup(t)
	allianceRepo := AllianceRepo.(*allianceRepository)
	allianceRepo.db = allianceRepo.db.Begin()

	Convey("RetrieveByAllianceId", t, func() {
		var allianceAsRetrieved *model.Alliance

		allianceAsRetrieved = AllianceRepo.FindByAllianceId(alliance.AllianceId)

		So(allianceAsRetrieved.AllianceId, ShouldEqual, alliance.AllianceId)
		So(allianceAsRetrieved.AllianceName, ShouldEqual, alliance.AllianceName)
		So(allianceAsRetrieved.AllianceTicker, ShouldEqual, alliance.AllianceTicker)
	})

	Convey("RetrieveByAllianceId_WhereAllianceDoesn'tExist",t, func() {
		var allianceAsRetrieved *model.Alliance

		allianceAsRetrieved = AllianceRepo.FindByAllianceId(20000)

		So(allianceAsRetrieved, ShouldBeNil)
	})

	Convey("Create", t, func() {
		var allianceAsRetrieved model.Alliance
		allianceAsCreated := model.Alliance{AllianceId: 2, AllianceName: "Test Alliacnce 2", AllianceTicker: "TST2"}

		err := AllianceRepo.Save(&allianceAsCreated)

		So(err, ShouldBeNil)

		allianceRepo.db.Where("alliance_id = 2").Find(&allianceAsRetrieved)

		So(allianceAsRetrieved.AllianceId, ShouldEqual, allianceAsCreated.AllianceId)
		So(allianceAsRetrieved.AllianceName, ShouldEqual, allianceAsCreated.AllianceName)
		So(allianceAsRetrieved.AllianceTicker, ShouldEqual, allianceAsCreated.AllianceTicker)
	})

	Convey("CreateWithoutId", t, func() {
		allianceAsCreated := model.Alliance{AllianceName: "Test Alliance No ID", AllianceTicker: "TST3"}

		err := AllianceRepo.Save(&allianceAsCreated)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "Primary key must not be 0")
	})

	Convey("FindAll", t, func() {
		allianceAsCreated := model.Alliance{AllianceId: 2, AllianceName: "Test Alliance No ID", AllianceTicker: "TST3"}

		//First lets create one, so we can get 2 back
		err := AllianceRepo.Save(&allianceAsCreated)

		So(err, ShouldBeNil)

		alliances := AllianceRepo.FindAll()

		So(len(alliances), ShouldEqual, 2)
	})

	allianceRepo.db.Rollback()
	SharedTearDown()
}

func Test_CorporationsCRUD_ThroughREPO(t *testing.T) {
	alliance, corporation, _, _, _ := SharedSetup(t)
	corpRepo := CorporationRepo.(*corporationRepository)
	corpRepo.db = corpRepo.db.Begin()

	Convey("RetrieveByCorporationId", t, func() {
		var corporationAsRetrieved *model.Corporation

		corporationAsRetrieved = CorporationRepo.FindByCorporationId(corporation.CorporationId)

		So(corporationAsRetrieved.CorporationId, ShouldEqual, corporation.CorporationId)
		So(corporationAsRetrieved.CorporationName, ShouldEqual, corporation.CorporationName)
		So(corporationAsRetrieved.CorporationTicker, ShouldEqual, corporation.CorporationTicker)
		So(corporationAsRetrieved.AllianceId, ShouldEqual, corporation.AllianceId)
		So(corporationAsRetrieved.Alliance.AllianceId, ShouldEqual, alliance.AllianceId)
		So(corporationAsRetrieved.Alliance.AllianceName, ShouldEqual, alliance.AllianceName)
		So(corporationAsRetrieved.Alliance.AllianceTicker, ShouldEqual, alliance.AllianceTicker)
	})

	Convey("RetrieveByCorporationId_WhereCorporationDoesn'tExist", t, func() {
		var corporationAsRetrieved *model.Corporation

		corporationAsRetrieved = CorporationRepo.FindByCorporationId(20000)

		So(corporationAsRetrieved, ShouldBeNil)
	})

	Convey("Create", t, func() {
		var corporationAsRetrieved model.Corporation
		corporationAsCreated := model.Corporation{
			CorporationId:     2,
			CorporationName:   "Test Corporation 2",
			CorporationTicker: "TST2",
			AllianceId:        &alliance.AllianceId,
			Alliance:          alliance,
		}

		err := CorporationRepo.Save(&corporationAsCreated)

		So(err, ShouldBeNil)

		corpRepo.db.Where("corporation_id = 2").Find(&corporationAsRetrieved)

		So(corporationAsRetrieved.CorporationId, ShouldEqual, corporationAsCreated.CorporationId)
		So(corporationAsRetrieved.CorporationName, ShouldEqual, corporationAsCreated.CorporationName)
		So(corporationAsRetrieved.CorporationTicker, ShouldEqual, corporationAsCreated.CorporationTicker)
		So(corporationAsRetrieved.AllianceId, ShouldEqual, corporationAsCreated.AllianceId)
	})

	Convey("CreateWithoutId", t, func() {
		corporationAsCreated := model.Corporation{
			CorporationName:   "Test Corporation 2",
			CorporationTicker: "TST2",
			AllianceId:        &alliance.AllianceId,
			Alliance:          alliance,
		}

		err := CorporationRepo.Save(&corporationAsCreated)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "Primary key must not be 0")
	})

	Convey("FindAll", t, func() {
		corporationAsCreated := model.Corporation{
			CorporationId:     2,
			CorporationName:   "Test Corporation 2",
			CorporationTicker: "TST2",
			AllianceId:        &alliance.AllianceId,
			Alliance:          alliance,
		}

		err := CorporationRepo.Save(&corporationAsCreated)

		So(err, ShouldBeNil)

		corporations := CorporationRepo.FindAll()

		So(len(corporations), ShouldEqual, 2)
	})

	corpRepo.db.Rollback()
	SharedTearDown()
}

func Test_CharactersCRUD_ThroughREPO(t *testing.T) {
	_, _, character, _, _ := SharedSetup(t)
	charRepo := CharacterRepo.(*characterRepository)
	charRepo.db = charRepo.db.Begin()

	Convey("RetrieveByCharacterId", t, func() {
		characterAsRetrieved := CharacterRepo.FindByCharacterId(character[0].CharacterId)

		So(characterAsRetrieved, ShouldResemble, character)
	})

	Convey("RetrieveByCharacterId_WhereCharacterDoesn'tExist", t, func() {
		characterAsRetrieved := CharacterRepo.FindByCharacterId(20000)

		So(characterAsRetrieved, ShouldBeNil)
	})

	Convey("RetrieveByAuthenticationCode", t, func() {
		characterAsRetrieved := CharacterRepo.FindByAutenticationCode("123456789")

		So(characterAsRetrieved.CharacterId, ShouldEqual, character[0].CharacterId)
		So(characterAsRetrieved.CharacterName, ShouldEqual, character[0].CharacterName)
	})

	Convey("RetrieveByAuthenticationCode_WhereAuthCodeDoesn'tExist", t, func() {
		characterAsRetrieved := CharacterRepo.FindByAutenticationCode("fjksadljfahuoifsoda")

		So(characterAsRetrieved, ShouldBeNil)
	})

	Convey("Create", t, func() {
		characterAsCreated := model.Character{CharacterId: 2, CorporationId: 1, Token: "123456", CharacterName: "Test Character 2"}
		var characterAsRetrieved model.Character

		err := CharacterRepo.Save(&characterAsCreated)

		So(err, ShouldBeNil)

		charRepo.db.Where("character_id = 2").Find(&characterAsRetrieved)

		So(characterAsRetrieved.CharacterId, ShouldEqual, characterAsCreated.CharacterId)
		So(characterAsRetrieved.CharacterName, ShouldEqual, characterAsCreated.CharacterName)
		So(characterAsRetrieved.CorporationId, ShouldEqual, characterAsCreated.CorporationId)
	})

	Convey("Create_WithBadId", t, func() {
		characterAsCreated := model.Character{CharacterId: 0, CorporationId: 1, Token: "123456", CharacterName: "Test Character 2"}

		err := CharacterRepo.Save(&characterAsCreated)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "Primary key must not be 0")
	})

	Convey("FindAll", t, func() {
		characterAsCreated := model.Character{CharacterId: 2, CorporationId: 1, Token: "123456", CharacterName: "Test Character 2"}

		err := CharacterRepo.Save(&characterAsCreated)

		So(err, ShouldBeNil)

		characters := CharacterRepo.FindAll()

		So(len(characters), ShouldEqual, 3)
	})

	charRepo.db.Rollback()
	SharedTearDown()
}

func Test_CreateAndRetrieve_UsersThroughREPO(t *testing.T) {
	_, _, character, user, authCode := SharedSetup(t)
	usersRepo := UserRepo.(*userRepository)
	usersRepo.db = usersRepo.db.Begin()

	t.Run("RetrieveByChatIdNoUser", func(t *testing.T) {
		userAsRetrieved := UserRepo.FindByChatId("123456")

		if userAsRetrieved != nil {
			t.Fatalf("Expected a nil user but instead got: %+v", user)
		}
	})

	t.Run("RetrieveByChatId", func(t *testing.T) {
		userAsRetrieved := UserRepo.FindByChatId(user.ChatId)

		if userAsRetrieved.ChatId != user.ChatId {
			t.Fatalf("Retrieved user's chat id: (%s) does not equal original: (%s)",
				userAsRetrieved.ChatId, user.ChatId)
		}

		if userAsRetrieved.UserId != user.UserId {
			t.Fatalf("Retrieved user's user id: (%d) does not equal original: (%d)",
				userAsRetrieved.UserId, user.UserId)
		}

		if len(userAsRetrieved.Characters) != len(user.Characters) {
			t.Fatalf("Retrieved user's character list size: (%d) does not equal original: (%d)",
				len(userAsRetrieved.Characters), len(user.Characters))
		}
	})

	t.Run("LinkCharacterToUserByAuthCode", func(t *testing.T) {
		err := UserRepo.LinkCharacterToUserByAuthCode(authCode[1].AuthenticationCode, &user)

		if err != nil {
			t.Fatalf("Had an error while linking a character to a user: %s", err)
		}

		var linkedUser []model.User
		var authCodeAsRetrieved model.AuthenticationCode
		var userAsRetrieved model.User
		usersRepo.db.Where("user_id = ?", user.UserId).Find(&userAsRetrieved)
		usersRepo.db.Model(&userAsRetrieved).Association("Characters").Find(&userAsRetrieved.Characters)
		usersRepo.db.Model(&character[1]).Association("Users").Find(&linkedUser)
		usersRepo.db.Where("character_id = ?", character[1].CharacterId).Find(&authCodeAsRetrieved)

		if len(linkedUser) == 0 {
			t.Fatal("Expected at least one linked user")
		}

		if len(userAsRetrieved.Characters) != 2 {
			t.Fatalf("User should have 2 characters instead they have: %d", len(userAsRetrieved.Characters))
		}

		if linkedUser[0].UserId != user.UserId {
			t.Fatalf("Linked user's user id: (%d) does not equal original: (%d)",
				linkedUser[0].UserId, user.UserId)
		}

		if authCodeAsRetrieved.IsUsed == false {
			t.Fatal("Auth code was not used up")
		}
	})

	t.Run("Create", func(t *testing.T) {
		var userAsRetrieved model.User
		userAsCreated := model.User{ChatId: "01234567890"}

		err := UserRepo.Save(&userAsCreated)

		if err != nil {
			t.Fatalf("Had an error while saving the user: %s", err)
		}

		usersRepo.db.Where("user_id = ?", userAsCreated.UserId).Find(&userAsRetrieved)

		if userAsCreated.ChatId != userAsRetrieved.ChatId {
			t.Fatalf("Retrieved user's chat id: (%s) does not equal original: (%s)",
				userAsCreated.ChatId, userAsRetrieved.ChatId)
		}

		if userAsCreated.UserId != userAsRetrieved.UserId {
			t.Fatalf("Retrieved user's user id: (%d) does not equal original: (%d)",
				userAsCreated.UserId, userAsRetrieved.UserId)
		}
	})

	usersRepo.db.Rollback()
	SharedTearDown()
}

func Test_LinkCharacterToUser_WithInvalidAuthCode(t *testing.T) {
	_, _, character, user, authCode := SharedSetup(t)

	user2 := model.User{UserId: 2, ChatId: "12345678901"}

	usersRepo := UserRepo.(*userRepository)
	usersRepo.db = usersRepo.db.Begin()

	err := usersRepo.Save(&user2)

	if err != nil {
		t.Fatalf("Could not create a new user: %s", err.Error())
	}

	err = UserRepo.LinkCharacterToUserByAuthCode(authCode[0].AuthenticationCode, &user2)

	if err == nil {
		t.Fatal("Expected an error while linking a character to a user")
	}

	var linkedUser []model.User
	var authCodeAsRetrieved model.AuthenticationCode
	var userAsRetrieved model.User
	usersRepo.db.Where("user_id = ?", user2.UserId).Find(&userAsRetrieved)
	usersRepo.db.Model(&userAsRetrieved).Association("Characters").Find(&userAsRetrieved.Characters)
	usersRepo.db.Model(&character[0]).Association("Users").Find(&linkedUser)
	usersRepo.db.Where("character_id = ?", character[0].CharacterId).Find(&authCodeAsRetrieved)

	if len(linkedUser) == 0 {
		t.Error("Expected at least one linked user")
	}

	if len(linkedUser) != 1 {
		t.Errorf("Expected 1 user but found: %d", len(linkedUser))
	}

	if len(userAsRetrieved.Characters) != 0 {
		t.Errorf("User should have 0 characters instead they have: %d", len(userAsRetrieved.Characters))
	}

	if linkedUser[0].UserId != user.UserId {
		t.Errorf("Linked user's user id: (%d) does not equal original: (%d)",
			linkedUser[0].UserId, user.UserId)
	}

	if authCodeAsRetrieved.IsUsed == false {
		t.Error("Auth code was not used up")
	}

	usersRepo.db.Rollback()
	SharedTearDown()
}

func Test_LinkCharacterToUser_WithNonExistingCharacter(t *testing.T) {
	_, _, _, user, _ := SharedSetup(t)
	usersRepo := UserRepo.(*userRepository)

	usersRepo.db = usersRepo.db.Begin()
	err := UserRepo.LinkCharacterToUserByAuthCode("sfjakdslfjaksdlajl", &user)

	if err == nil {
		t.Fatal("Expected an error while linking a character to a user")
	}
}

func Test_RolesCRUD_ThroughREPO(t *testing.T) {
	SharedSetup(t)
	rolesRepo := RoleRepo.(*roleRepository)
	rolesRepo.db = rolesRepo.db.Begin()

	t.Run("CreateNoChatServiceGroup", func(t *testing.T) {
		var roleAsRetrieved model.Role
		newRole := model.Role{RoleName: "TEST_ROLE_FOR_TESTING"}
		err := RoleRepo.Save(&newRole)

		if err != nil {
			t.Fatalf("Had an error while saving the role: %s", err)
		}

		rolesRepo.db.Where("role_name = 'TEST_ROLE_FOR_TESTING'").Find(&roleAsRetrieved)

		if roleAsRetrieved.RoleName != newRole.RoleName {
			t.Fatalf("Retrieved role's name: (%s) does not match original: (%s)",
				roleAsRetrieved.RoleName, newRole.RoleName)
		}

		if roleAsRetrieved.ChatServiceGroup != newRole.ChatServiceGroup {
			t.Fatalf("Retrieved role's chatservice group: (%s) does not match original: (%s)",
				roleAsRetrieved.ChatServiceGroup, newRole.ChatServiceGroup)
		}
	})

	t.Run("CreateWithChatServiceGroup", func(t *testing.T) {
		var roleAsRetrieved model.Role
		newRole := model.Role{RoleName: "TEST_ROLE_FOR_TESTING2", ChatServiceGroup: "SUPER_COOL_CHAT"}
		err := RoleRepo.Save(&newRole)

		if err != nil {
			t.Fatalf("Had an error while saving the role: %s", err)
		}

		rolesRepo.db.Where("role_name = 'TEST_ROLE_FOR_TESTING2'").Find(&roleAsRetrieved)

		if roleAsRetrieved.RoleName != newRole.RoleName {
			t.Fatalf("Retrieved role's name: (%s) does not match original: (%s)",
				roleAsRetrieved.RoleName, newRole.RoleName)
		}

		if roleAsRetrieved.ChatServiceGroup != newRole.ChatServiceGroup {
			t.Fatalf("Retrieved role's chatservice group: (%s) does not match original: (%s)",
				roleAsRetrieved.ChatServiceGroup, newRole.ChatServiceGroup)
		}
	})

	t.Run("FindByRoleName", func(t *testing.T) {
		newRole := model.Role{RoleName: "TEST_ROLE_FOR_TESTING3", ChatServiceGroup: "SUPER_COOL_CHAT2"}
		err := RoleRepo.Save(&newRole)

		if err != nil {
			t.Fatalf("Had an error while saving the role: %s", err)
		}

		roleAsRetrieved := RoleRepo.FindByRoleName("TEST_ROLE_FOR_TESTING3")

		if roleAsRetrieved == nil {
			t.Fatal("Expected a role but retrieved nil")
		}

		if roleAsRetrieved.RoleId != newRole.RoleId {
			t.Errorf("Role id incorrect, expected: (%d) but retrieved: (%d)", newRole.RoleId, roleAsRetrieved.RoleId)
		}

		if roleAsRetrieved.RoleName != newRole.RoleName {
			t.Errorf("Role name, expected: (%s) but retrieved: (%s)", newRole.RoleName, roleAsRetrieved.RoleName)
		}

		if roleAsRetrieved.ChatServiceGroup != newRole.ChatServiceGroup {
			t.Errorf("Role chat service group, expected: (%s) but retrieved: (%s)", newRole.ChatServiceGroup, roleAsRetrieved.ChatServiceGroup)
		}
	})

	rolesRepo.db.Rollback()
	SharedTearDown()
}

func Test_CreateAndRetrieveAuthenticationCodes_ThroughREPO(t *testing.T) {
	_, _, characters, _, _ := SharedSetup(t)
	authCodeRepo := AuthenticationCodeRepo.(*authCodeRepository)
	authCodeRepo.db = authCodeRepo.db.Begin()

	t.Run("Create", func(t *testing.T) {
		var authCodesAsRetrieved []model.AuthenticationCode

		AuthenticationCodeRepo.Save(&characters[0], "testtest123")

		authCodeRepo.db.Where("character_id = ?", characters[0].CharacterId).Find(&authCodesAsRetrieved)

		aMatchFound := false
		aMatchFoundIdx := -1
		for idx, authCode := range authCodesAsRetrieved {
			if authCode.AuthenticationCode == "testtest123" {
				aMatchFound = true
				aMatchFoundIdx = idx
			}
		}

		if !aMatchFound {
			t.Fatal("Expected one auth code match, found none")
		}

		if authCodesAsRetrieved[aMatchFoundIdx].CharacterId != characters[0].CharacterId {
			t.Fatalf("Retrieved matching auth code's character id: (%d) did not match original: (%d)",
				authCodesAsRetrieved[aMatchFoundIdx].CharacterId, characters[0].CharacterId)
		}
	})

	authCodeRepo.db.Rollback()
	SharedTearDown()
}

func TestSetup_BadConnectionString(t *testing.T) {
	err := Setup("derp", "yep")
	if err == nil {
		t.Error("Should have received an error... what kind of connection string is yep on the derp driver?")
	}
}
