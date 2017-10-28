//TODO: Make some tests for role insertion and delete to bounce off a real DB.
package repository

import (
	"database/sql"
	"errors"
	"git.maurer-it.net/abaeve/auth-srv/model"
	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"strings"
	"testing"
)

type accessError struct {
	message string
}

func (ae *accessError) Error() string {
	return ae.message
}

func AccessesSharedSetup(t *testing.T, mainQuery string) (*sql.DB, sqlmock.Sqlmock, string) {
	SharedSetup(t)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}

	accesses := AccessRepo.(*accessesRepo)

	accesses.db = sqlx.NewDb(db, "mysql")

	query := makeQueryStringRegex(mainQuery)

	return db, mock, query
}

func TestAccessesWithAllianceRoleAndCorpRole(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, roleQuery)
	defer db.Close()

	mock.ExpectPrepare(query).ExpectQuery().WithArgs(
		"1234567890", "1234567890", "1234567890", "1234567890", "1234567890", "1234567890",
	).WillReturnRows(
		sqlmock.NewRows(
			[]string{
				"role_name",
				"chatservice_group",
			},
		).AddRow("ABA", "ABA").AddRow("VSKY", "VSKY"))

	roles, err := AccessRepo.FindByChatId("1234567890")

	if err != nil {
		t.Fatalf("Received an error when one wasn't expected: %s", err)
	}

	if len(roles) != 2 {
		t.Fatalf("Expected 2 roles but received %d", len(roles))
	}

	if roles[0] != "ABA" {
		t.Errorf("First role was not as expected, got '%s' expected 'ABA'", roles[0])
	}

	if roles[1] != "VSKY" {
		t.Errorf("Second role was not as expected, got '%s' expected 'VSKY'", roles[1])
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesWith20Roles(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, roleQuery)
	defer db.Close()

	mock.ExpectPrepare(query).ExpectQuery().WithArgs(
		"1234567890", "1234567890", "1234567890", "1234567890", "1234567890", "1234567890",
	).WillReturnRows(
		sqlmock.NewRows(
			[]string{
				"role_name",
				"chatservice_group",
			},
		).AddRow(
			"ROLE1", "ROLE1",
		).AddRow(
			"ROLE2", "ROLE2",
		).AddRow(
			"ROLE3", "ROLE3",
		).AddRow(
			"ROLE4", "ROLE4",
		).AddRow(
			"ROLE5", "ROLE5",
		).AddRow(
			"ROLE6", "ROLE6",
		).AddRow(
			"ROLE7", "ROLE7",
		).AddRow(
			"ROLE8", "ROLE8",
		).AddRow(
			"ROLE9", "ROLE9",
		).AddRow(
			"ROLE10", "ROLE10",
		).AddRow(
			"ROLE11", "ROLE11",
		).AddRow(
			"ROLE12", "ROLE12",
		).AddRow(
			"ROLE13", "ROLE13",
		).AddRow(
			"ROLE14", "ROLE14",
		).AddRow(
			"ROLE15", "ROLE15",
		).AddRow(
			"ROLE16", "ROLE16",
		).AddRow(
			"ROLE17", "ROLE17",
		).AddRow(
			"ROLE18", "ROLE18",
		).AddRow(
			"ROLE19", "ROLE19",
		).AddRow(
			"ROLE20", "ROLE20",
		),
	)

	roles, err := AccessRepo.FindByChatId("1234567890")

	if err != nil {
		t.Fatalf("Received an error when one wasn't expected: %s", err)
	}

	if len(roles) != 20 {
		t.Fatalf("Expected 2 roles but received %d", len(roles))
	}

	if roles[0] != "ROLE1" {
		t.Errorf("First role was not as expected, got '%s' expected 'ROLE1'", roles[0])
	}

	if roles[19] != "ROLE20" {
		t.Errorf("20th role was not as expected, got '%s' expected 'ROLE20'", roles[1])
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestErrorOnPrepare(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, roleQuery)
	defer db.Close()

	mock.ExpectPrepare(query).WillReturnError(&accessError{message: "Database connection lost"})

	_, err := AccessRepo.FindByChatId("")

	if err == nil && err.Error() != "Database connection lost" {
		t.Fatal("Expected 'Database connection lost' as the error but received nothing or the wrong thing")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestErrorOnQuery(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, roleQuery)
	defer db.Close()

	mock.ExpectPrepare(query).ExpectQuery().WillReturnError(&accessError{message: "Database connection lost"})

	_, err := AccessRepo.FindByChatId("")

	if err == nil && err.Error() != "Database connection lost" {
		t.Fatal("Expected 'Database connection lost' as the error but received nothing or the wrong thing")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_SaveAllianceAndCorpRole(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, allianceCorpInsert)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(
		query,
	).ExpectExec().WithArgs(
		int64(1),
		int64(2),
		int64(3),
	).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := AccessRepo.SaveAllianceAndCorpRole(int64(1), int64(2), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	if err != nil {
		t.Errorf("Received an error but expected none (%s).", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_SaveAllianceAndCorpRole_WithPrepareError(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, allianceCorpInsert)
	defer db.Close()

	expectedError := "I derped somewhere on the floor over there?"

	mock.ExpectBegin()
	mock.ExpectPrepare(query).WillReturnError(errors.New(expectedError))
	mock.ExpectRollback()

	err := AccessRepo.SaveAllianceAndCorpRole(int64(1), int64(2), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	expectedError = "Error during prepare: " + expectedError

	if err == nil {
		t.Fatal("Expected an error but received nil")
	}

	if err.Error() != expectedError {
		t.Errorf("Expected error text: (%s) but received: (%s)", expectedError, err.Error())
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_SaveAllianceAndCorpRole_WithExecError(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, allianceCorpInsert)
	defer db.Close()

	expectedError := "I derped somewhere on the floor over there?"

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(2), int64(3)).WillReturnError(errors.New(expectedError))
	mock.ExpectRollback()

	err := AccessRepo.SaveAllianceAndCorpRole(int64(1), int64(2), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	expectedError = "Error during exec: " + expectedError

	if err == nil {
		t.Fatal("Expected an error but received nil")
	}

	if err.Error() != expectedError {
		t.Errorf("Expected error text: (%s) but received: (%s)", expectedError, err.Error())
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_SaveAllianceRole(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, allianceInsert)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(3)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := AccessRepo.SaveAllianceRole(int64(1), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	if err != nil {
		t.Errorf("Received an error but expected none: (%s).", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_SaveAllianceRole_WithPrepareError(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, allianceInsert)
	defer db.Close()

	expectedError := "I dropped the bomb right there ->"

	mock.ExpectBegin()
	mock.ExpectPrepare(query).WillReturnError(errors.New(expectedError))
	mock.ExpectRollback()

	err := AccessRepo.SaveAllianceRole(int64(1), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	expectedError = "Error during prepare: " + expectedError

	if err == nil {
		t.Error("Expected an error but received nil")
	}

	if err.Error() != expectedError {
		t.Errorf("Expected error text: (%s) but received: (%s)", expectedError, err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_SaveAllianceRole_WithExecError(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, allianceInsert)
	defer db.Close()

	expectedError := "I dropped the bomb right there ->"

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(3)).WillReturnError(errors.New(expectedError))
	mock.ExpectRollback()

	err := AccessRepo.SaveAllianceRole(int64(1), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	expectedError = "Error during exec: " + expectedError

	if err == nil {
		t.Error("Expected an error but received nil")
	}

	if err.Error() != expectedError {
		t.Errorf("Expected error text: (%s) but received: (%s)", expectedError, err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_SaveCorporationRole(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, corporationInsert)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(3)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := AccessRepo.SaveCorporationRole(int64(1), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	if err != nil {
		t.Errorf("Received an error but expected none: (%s).", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_SaveCorporationRole_WithPrepareError(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, corporationInsert)
	defer db.Close()

	expectedError := "I dropped the bomb right there ->"

	mock.ExpectBegin()
	mock.ExpectPrepare(query).WillReturnError(errors.New(expectedError))
	mock.ExpectRollback()

	err := AccessRepo.SaveCorporationRole(int64(1), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	expectedError = "Error during prepare: " + expectedError

	if err == nil {
		t.Error("Expected an error but received nil")
	}

	if err.Error() != expectedError {
		t.Errorf("Expected error text: (%s) but received: (%s)", expectedError, err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_SaveCorporationRole_WithExecError(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, corporationInsert)
	defer db.Close()

	expectedError := "I dropped the bomb right there ->"

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(3)).WillReturnError(errors.New(expectedError))
	mock.ExpectRollback()

	err := AccessRepo.SaveCorporationRole(int64(1), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	expectedError = "Error during exec: " + expectedError

	if err == nil {
		t.Error("Expected an error but received nil")
	}

	if err.Error() != expectedError {
		t.Errorf("Expected error text: (%s) but received: (%s)", expectedError, err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

// Finally after the writing the same test code... and same implementation code I wrote some shared code... now I don't want to delete the code
// above this point that I wrote...  Most of the above *WithPrepareError and *WithExecError will just live on until someone decides we don't need
// to test that shared functionality again... and again... and again... :P
func TestAccessesRepo_SaveCharacterRole(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, characterInsert)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(3)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := AccessRepo.SaveCharacterRole(int64(1), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	if err != nil {
		t.Errorf("Received an error but expected none: (%s).", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_SaveAllianceCharacterLeadershipRole(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, characterAllianceInsert)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(2), int64(3)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := AccessRepo.SaveAllianceCharacterLeadershipRole(int64(1), int64(2), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	if err != nil {
		t.Errorf("Received an error but expected none: (%s).", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_SaveCorporationCharacterLeadershipRole(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, characterCorpInsert)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(2), int64(3)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := AccessRepo.SaveCorporationCharacterLeadershipRole(int64(1), int64(2), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	if err != nil {
		t.Errorf("Received an error but expected none: (%s).", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_DeleteAllianceAndCorpRole(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, allianceCorpDelete)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(2), int64(3)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	rows, err := AccessRepo.DeleteAllianceAndCorpRole(int64(1), int64(2), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	if err != nil {
		t.Errorf("Received an error but expected none: (%s).", err)
	}

	if rows != 1 {
		t.Errorf("Expected 1 modification but recieved: (%d)", rows)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_DeleteAllianceRole(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, allianceDelete)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(3)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	rows, err := AccessRepo.DeleteAllianceRole(int64(1), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	if err != nil {
		t.Errorf("Received an error but expected none: (%s).", err)
	}

	if rows != 1 {
		t.Errorf("Expected 1 modification but recieved: (%d)", rows)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_DeleteCorporationRole(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, corporationDelete)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(3)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	rows, err := AccessRepo.DeleteCorporationRole(int64(1), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	if err != nil {
		t.Errorf("Received an error but expected none: (%s).", err)
	}

	if rows != 1 {
		t.Errorf("Expected 1 modification but recieved: (%d)", rows)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_DeleteCharacterRole(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, characterDelete)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(3)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	rows, err := AccessRepo.DeleteCharacterRole(int64(1), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	if err != nil {
		t.Errorf("Received an error but expected none: (%s).", err)
	}

	if rows != 1 {
		t.Errorf("Expected 1 modification but recieved: (%d)", rows)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_DeleteAllianceCharacterLeadershipRole(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, characterAllianceDelete)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(2), int64(3)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	rows, err := AccessRepo.DeleteAllianceCharacterLeadershipRole(int64(1), int64(2), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	if err != nil {
		t.Errorf("Received an error but expected none: (%s).", err)
	}

	if rows != 1 {
		t.Errorf("Expected 1 modification but recieved: (%d)", rows)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestAccessesRepo_DeleteCorporationCharacterLeadershipRole(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, characterCorpDelete)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(2), int64(3)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	rows, err := AccessRepo.DeleteCorporationCharacterLeadershipRole(int64(1), int64(2), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	if err != nil {
		t.Errorf("Received an error but expected none: (%s).", err)
	}

	if rows != 1 {
		t.Errorf("Expected 1 modification but recieved: (%d)", rows)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_doubleEntityRoleQuery_WithTransactionBeginError(t *testing.T) {
	db, mock, _ := AccessesSharedSetup(t, allianceCorpInsert)
	defer db.Close()

	expectedError := "I'm sorry, Dave. I'm afraid I can't do that."

	mock.ExpectBegin().WillReturnError(errors.New(expectedError))

	err := AccessRepo.SaveAllianceAndCorpRole(int64(1), int64(2), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	expectedError = "Error opening a transaction: " + expectedError

	if err == nil {
		t.Error("Expected an error but received nil")
	}

	if err.Error() != expectedError {
		t.Errorf("Expected error text: (%s) but received: (%s)", expectedError, err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_doubleEntityRoleQuery_WithMultipleInsertions(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, allianceCorpInsert)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(2), int64(3)).WillReturnResult(sqlmock.NewResult(0, 2))
	mock.ExpectRollback()

	err := AccessRepo.SaveAllianceAndCorpRole(int64(1), int64(2), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	expectedError := "Inserted more than one record?"

	if err == nil {
		t.Error("Expected an error but received nil")
	}

	if err.Error() != expectedError {
		t.Errorf("Expected error text: (%s) but received: (%s)", expectedError, err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_singleEntityRoleQuery_WithTransactionBeginError(t *testing.T) {
	db, mock, _ := AccessesSharedSetup(t, corporationInsert)
	defer db.Close()

	expectedError := "I'm sorry, Dave. I'm afraid I can't do that."

	mock.ExpectBegin().WillReturnError(errors.New(expectedError))

	err := AccessRepo.SaveCorporationRole(int64(1), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	expectedError = "Error opening a transaction: " + expectedError

	if err == nil {
		t.Error("Expected an error but received nil")
	}

	if err.Error() != expectedError {
		t.Errorf("Expected error text: (%s) but received: (%s)", expectedError, err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_singleEntityRoleQuery_WithMultipleInsertions(t *testing.T) {
	db, mock, query := AccessesSharedSetup(t, corporationInsert)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(query).ExpectExec().WithArgs(int64(1), int64(3)).WillReturnResult(sqlmock.NewResult(0, 2))
	mock.ExpectRollback()

	err := AccessRepo.SaveCorporationRole(int64(1), &model.Role{
		RoleId:           int64(3),
		RoleName:         "TEST_ROLE1",
		ChatServiceGroup: "TEST_ROLE1",
	})

	expectedError := "Inserted more than one record?"

	if err == nil {
		t.Error("Expected an error but received nil")
	}

	if err.Error() != expectedError {
		t.Errorf("Expected error text: (%s) but received: (%s)", expectedError, err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func makeQueryStringRegex(queryString string) string {
	sqlRegex := strings.Replace(queryString, "(", ".", -1)
	sqlRegex = strings.Replace(sqlRegex, ")", ".", -1)
	sqlRegex = strings.Replace(sqlRegex, "?", ".", -1)

	return sqlRegex
}
