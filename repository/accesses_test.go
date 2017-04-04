package repository

import (
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

func TestAccessesWithAllianceRoleAndCorpRole(t *testing.T) {
	SharedSetup(t)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	accesses := AccessRepo.(*accessesRepo)

	accesses.db = db

	sql := strings.Replace(roleQuery, "(", ".", -1)
	sql = strings.Replace(sql, ")", ".", -1)
	sql = strings.Replace(sql, "?", ".", -1)

	mock.ExpectPrepare(sql).ExpectQuery().WithArgs(
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
}

func TestAccessesWith20Roles(t *testing.T) {
	SharedSetup(t)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	accesses := AccessRepo.(*accessesRepo)

	accesses.db = db

	sql := strings.Replace(roleQuery, "(", ".", -1)
	sql = strings.Replace(sql, ")", ".", -1)
	sql = strings.Replace(sql, "?", ".", -1)

	mock.ExpectPrepare(sql).ExpectQuery().WithArgs(
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
}

func TestErrorOnPrepare(t *testing.T) {
	SharedSetup(t)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	accesses := AccessRepo.(*accessesRepo)

	accesses.db = db

	sql := strings.Replace(roleQuery, "(", ".", -1)
	sql = strings.Replace(sql, ")", ".", -1)
	sql = strings.Replace(sql, "?", ".", -1)

	mock.ExpectPrepare(sql).WillReturnError(&accessError{message: "Database connection lost"})

	_, err = AccessRepo.FindByChatId("")

	if err == nil && err.Error() != "Database connection lost" {
		t.Fatal("Expected 'Database connection lost' as the error but received nothing or the wrong thing")
	}
}

func TestErrorOnQuery(t *testing.T) {
	SharedSetup(t)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	accesses := AccessRepo.(*accessesRepo)

	accesses.db = db

	sql := strings.Replace(roleQuery, "(", ".", -1)
	sql = strings.Replace(sql, ")", ".", -1)
	sql = strings.Replace(sql, "?", ".", -1)

	mock.ExpectPrepare(sql).ExpectQuery().WillReturnError(&accessError{message: "Database connection lost"})

	_, err = AccessRepo.FindByChatId("")

	if err == nil && err.Error() != "Database connection lost" {
		t.Fatal("Expected 'Database connection lost' as the error but received nothing or the wrong thing")
	}
}
