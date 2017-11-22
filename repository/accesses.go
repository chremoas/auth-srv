package repository

import (
	"errors"
	"github.com/chremoas/auth-srv/model"
	"github.com/jmoiron/sqlx"
)

type AccessesRepository interface {
	SaveAllianceAndCorpRole(allianceId, corporationId int64, role *model.Role) error
	SaveAllianceRole(allianceId int64, role *model.Role) error
	SaveCorporationRole(corporationId int64, role *model.Role) error
	SaveCharacterRole(characterId int64, role *model.Role) error
	SaveAllianceCharacterLeadershipRole(allianceId, characterId int64, role *model.Role) error
	SaveCorporationCharacterLeadershipRole(corporationId, characterId int64, role *model.Role) error

	DeleteAllianceAndCorpRole(allianceId, corporationId int64, role *model.Role) (int64, error)
	DeleteAllianceRole(allianceId int64, role *model.Role) (int64, error)
	DeleteCorporationRole(corporationId int64, role *model.Role) (int64, error)
	DeleteCharacterRole(characterId int64, role *model.Role) (int64, error)
	DeleteAllianceCharacterLeadershipRole(allianceId, characterId int64, role *model.Role) (int64, error)
	DeleteCorporationCharacterLeadershipRole(corporationId, characterId int64, role *model.Role) (int64, error)

	FindByChatId(chatId string) ([]string, error)
}

type accessesRepo struct {
	db *sqlx.DB
}

var roleQuery string = `
SELECT
  role.role_name,
  role.chatservice_group
FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN corporations corp ON (chars.corporation_id = corp.corporation_id)
  JOIN alliances alliance ON (corp.alliance_id = alliance.alliance_id)
  JOIN alliance_corporation_role_map acrm
    ON (alliance.alliance_id = acrm.alliance_id AND corp.corporation_id = acrm.corporation_id)
  JOIN roles role ON (acrm.role_id = role.role_id)
WHERE user.chat_id = ?
UNION
SELECT
  role.role_name,
  role.chatservice_group
  FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN corporations corp ON (chars.corporation_id = corp.corporation_id)
  JOIN alliances alliance ON (corp.alliance_id = alliance.alliance_id)
  JOIN alliance_role_map arm ON (alliance.alliance_id = arm.alliance_id)
  JOIN roles role ON (arm.role_id = role.role_id)
WHERE user.chat_id = ?
UNION
SELECT
  role.role_name,
  role.chatservice_group
  FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN corporations corp ON (chars.corporation_id = corp.corporation_id)
  JOIN corporation_role_map crm ON (corp.corporation_id = crm.corporation_id)
  JOIN roles role ON (crm.role_id = role.role_id)
WHERE user.chat_id = ?
UNION
SELECT
  role.role_name,
  role.chatservice_group
  FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN character_role_map crm ON (chars.character_id = crm.character_id)
  JOIN roles role ON (crm.role_id = role.role_id)
WHERE user.chat_id = ?
UNION
SELECT
  role.role_name,
  role.chatservice_group
  FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN corporations corp ON (chars.corporation_id = corp.corporation_id)
  JOIN alliances alliance ON (corp.alliance_id = alliance.alliance_id)
  JOIN alliance_character_leadership_role_map aclrm
    ON (chars.character_id = aclrm.character_id AND alliance.alliance_id = aclrm.alliance_id)
  JOIN roles role ON (aclrm.role_id = role.role_id)
WHERE user.chat_id = ?
UNION
SELECT
  role.role_name,
  role.chatservice_group
  FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN corp_character_leadership_role_map cclrm
    ON (chars.character_id = cclrm.character_id AND chars.corporation_id = cclrm.corporation_id)
  JOIN roles role ON (cclrm.role_id = role.role_id)
WHERE user.chat_id = ?
`

const (
	allianceCorpInsert = "insert into alliance_corporation_role_map (alliance_id, corporation_id, role_id) values (?, ?, ?)"

	allianceCorpDelete = "delete from alliance_corporation_role_map where alliance_id = ? and corporation_id = ? and role_id = ?"

	allianceInsert = "insert into alliance_role_map (alliance_id, role_id) values (?, ?)"

	allianceDelete = "delete from alliance_role_map where alliance_id = ? and role_id = ?"

	corporationInsert = "insert into corporation_role_map (corporation_id, role_id) values (?, ?)"

	corporationDelete = "delete from corporation_role_map where corporation_id = ? and role_id = ?"

	characterInsert = "insert into character_role_map (character_id, role_id) values (?, ?)"

	characterDelete = "delete from character_role_map where character_id = ? and role_id = ?"

	characterAllianceInsert = "insert into alliance_character_leadership_role_map (alliance_id, character_id, role_id) values(?, ?, ?)"

	characterAllianceDelete = "delete from alliance_character_leadership_role_map where alliance_id = ? and character_id = ? and role_id = ?"

	characterCorpInsert = "insert into corp_character_leadership_role_map (corporation_id, character_id, role_id) values (?, ?, ?)"

	characterCorpDelete = "delete from corp_character_leadership_role_map where corporation_id = ? and character_id = ? and role_id = ?"
)

// Saves a role that is linked to an alliance AND a corporation
// alliance_coporation_role_map table
func (acc *accessesRepo) SaveAllianceAndCorpRole(allianceId, corporationId int64, role *model.Role) error {
	_, err := acc.doubleEntityRoleQuery(allianceId, corporationId, role.RoleId, allianceCorpInsert)
	return err
}

// Saves a role that is linked to an alliance
// alliance_role_map table
func (acc *accessesRepo) SaveAllianceRole(allianceId int64, role *model.Role) error {
	_, err := acc.singleEntityRoleQuery(allianceId, role.RoleId, allianceInsert)
	return err
}

// Saves a role that is linked to a corporation
// corporation_role_map table
func (acc *accessesRepo) SaveCorporationRole(corporationId int64, role *model.Role) error {
	_, err := acc.singleEntityRoleQuery(corporationId, role.RoleId, corporationInsert)
	return err
}

// Saves a role that is linked to a character
// character_role_map table
func (acc *accessesRepo) SaveCharacterRole(characterId int64, role *model.Role) error {
	_, err := acc.singleEntityRoleQuery(characterId, role.RoleId, characterInsert)
	return err
}

// Saves a role that is linked to an alliance and a character in an alliance leadership position
// alliance_character_leadership_role_map table
func (acc *accessesRepo) SaveAllianceCharacterLeadershipRole(allianceId, characterId int64, role *model.Role) error {
	_, err := acc.doubleEntityRoleQuery(allianceId, characterId, role.RoleId, characterAllianceInsert)
	return err
}

// Saves a role that is linked to a corporation and a character in a corporation leadership position
// corp_character_leadership_role_map table
func (acc *accessesRepo) SaveCorporationCharacterLeadershipRole(corporationId, characterId int64, role *model.Role) error {
	_, err := acc.doubleEntityRoleQuery(corporationId, characterId, role.RoleId, characterCorpInsert)
	return err
}

func (acc *accessesRepo) DeleteAllianceAndCorpRole(allianceId, corporationId int64, role *model.Role) (int64, error) {
	return acc.doubleEntityRoleQuery(allianceId, corporationId, role.RoleId, allianceCorpDelete)
}

func (acc *accessesRepo) DeleteAllianceRole(allianceId int64, role *model.Role) (int64, error) {
	return acc.singleEntityRoleQuery(allianceId, role.RoleId, allianceDelete)
}

func (acc *accessesRepo) DeleteCorporationRole(corporationId int64, role *model.Role) (int64, error) {
	return acc.singleEntityRoleQuery(corporationId, role.RoleId, corporationDelete)
}

func (acc *accessesRepo) DeleteCharacterRole(characterId int64, role *model.Role) (int64, error) {
	return acc.singleEntityRoleQuery(characterId, role.RoleId, characterDelete)
}

func (acc *accessesRepo) DeleteAllianceCharacterLeadershipRole(allianceId, characterId int64, role *model.Role) (int64, error) {
	return acc.doubleEntityRoleQuery(allianceId, characterId, role.RoleId, characterAllianceDelete)
}

func (acc *accessesRepo) DeleteCorporationCharacterLeadershipRole(corporationId, characterId int64, role *model.Role) (int64, error) {
	return acc.doubleEntityRoleQuery(corporationId, characterId, role.RoleId, characterCorpDelete)
}

// Will take the two entityId's, roleId and query (as long as it follows the insert into table (entityOneId, entityTwoId, roleId) values (?, ?, ?))
// flow and execute it in a transaction returning any error that happened.
func (acc *accessesRepo) doubleEntityRoleQuery(entityOneId, entityTwoId, roleId int64, query string) (int64, error) {
	tx, err := acc.db.Begin()
	if err != nil {
		return 0, errors.New("Error opening a transaction: " + err.Error())
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return 0, errors.New("Error during prepare: " + err.Error())
	}

	result, err := stmt.Exec(entityOneId, entityTwoId, roleId)
	if err != nil {
		tx.Rollback()
		return 0, errors.New("Error during exec: " + err.Error())
	}

	rows, _ := result.RowsAffected()

	if rows != int64(1) {
		tx.Rollback()
		return rows, errors.New("nserted more than one record?")
	}

	stmt.Close()
	tx.Commit()

	return rows, nil
}

// Will take the entityId, roleId and query (as long as it follows the insert into table (entityId, roleId) values (?, ?))
// flow and execute it in a transaction returning any error that happened.
func (acc *accessesRepo) singleEntityRoleQuery(entityId, roleId int64, query string) (int64, error) {
	tx, err := acc.db.Begin()
	if err != nil {
		return 0, errors.New("Error opening a transaction: " + err.Error())
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return 0, errors.New("Error during prepare: " + err.Error())
	}

	result, err := stmt.Exec(entityId, roleId)
	if err != nil {
		tx.Rollback()
		return 0, errors.New("Error during exec: " + err.Error())
	}

	rows, _ := result.RowsAffected()

	if rows != int64(1) {
		tx.Rollback()
		return rows, errors.New("inserted more than one record")
	}

	stmt.Close()
	tx.Commit()

	return rows, nil
}

// Will be the main usage in anything automated.  This method will lookup all the available roles for the given
// Chat user id and return them as a map[string]string where the key is the role and the value is the chat
// group to apply.
func (acc *accessesRepo) FindByChatId(chatId string) ([]string, error) {
	var roles []string

	statement, err := acc.db.Prepare(roleQuery)
	if err != nil {
		return []string{}, err
	}

	rows, err := statement.Query(chatId, chatId, chatId, chatId, chatId, chatId)
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	roleName, chatRoleName := "", ""
	for idx := 0; rows.Next(); idx++ {
		rows.Scan(&roleName, &chatRoleName)

		//TODO: Is this innefficient?  Should I be going about growing my array differently?
		roles = append(roles, chatRoleName)
	}

	statement.Close()

	return roles, nil
}
