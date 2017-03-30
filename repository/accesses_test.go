package repository

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"strings"
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
				"role_from",
			},
		).AddRow("ABA", "ABA", "alliance").AddRow("VSKY", "VSKY", "corp"))

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
				"role_from",
			},
		).AddRow(
			"ROLE1", "ROLE1", "alliance",
		).AddRow(
			"ROLE2", "ROLE2", "corp",
		).AddRow(
			"ROLE3", "ROLE3", "corp",
		).AddRow(
			"ROLE4", "ROLE4", "corp",
		).AddRow(
			"ROLE5", "ROLE5", "corp",
		).AddRow(
			"ROLE6", "ROLE6", "corp",
		).AddRow(
			"ROLE7", "ROLE7", "corp",
		).AddRow(
			"ROLE8", "ROLE8", "corp",
		).AddRow(
			"ROLE9", "ROLE9", "corp",
		).AddRow(
			"ROLE10", "ROLE10", "corp",
		).AddRow(
			"ROLE11", "ROLE11", "corp",
		).AddRow(
			"ROLE12", "ROLE12", "corp",
		).AddRow(
			"ROLE13", "ROLE13", "corp",
		).AddRow(
			"ROLE14", "ROLE14", "corp",
		).AddRow(
			"ROLE15", "ROLE15", "corp",
		).AddRow(
			"ROLE16", "ROLE16", "corp",
		).AddRow(
			"ROLE17", "ROLE17", "corp",
		).AddRow(
			"ROLE18", "ROLE18", "corp",
		).AddRow(
			"ROLE19", "ROLE19", "corp",
		).AddRow(
			"ROLE20", "ROLE20", "corp",
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