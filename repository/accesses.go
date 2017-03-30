package repository

import (
	"database/sql"
	"github.com/abaeve/auth-srv/model"
)

type AccessesRepository interface {
	SaveAllianceAndCorpRole(allianceId int64, corporationId int64, role model.Role) error
	SaveAllianceRole(allianceId int64, role model.Role) error
	SaveCorporationRole(corporationId int64, role model.Role) error
	SaveCharacterRole(characterId int64, role model.Role) error
	SaveAllianceCharacterLeadershipRole(allianceId int64, role model.Role) error
	SaveCorporationCharacterLeadershipRole(corporationId int64, role model.Role) error

	FindByChatId(chatId string) ([]string, error)
}

type accessesRepo struct {
	db *sql.DB
}

var roleQuery string = `
SELECT
  role.role_name,
  role.chatservice_group,
  'alliance_corp' AS role_from
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
  role.chatservice_group,
  'alliance' AS role_from
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
  role.chatservice_group,
  'corp' AS role_from
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
  role.chatservice_group,
  'character' AS role_from
  FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN character_role_map crm ON (chars.corporation_id = crm.character_id)
  JOIN roles role ON (crm.role_id = role.role_id)
WHERE user.chat_id = ?
UNION
SELECT
  role.role_name,
  role.chatservice_group,
  'alliance_corporation_leadership' AS role_from
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
  role.chatservice_group,
  'corporation_character_leadership' AS role_from
  FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN corp_character_leadership_role_map cclrm
    ON (chars.character_id = cclrm.character_id AND chars.corporation_id = cclrm.corporation_id)
  JOIN roles role ON (cclrm.role_id = role.role_id)
WHERE user.chat_id = ?
`

// Saves a role that is linked to an alliance AND a corporation
// alliance_coporation_role_map table
func (acc *accessesRepo) SaveAllianceAndCorpRole(allianceId int64, corporationId int64, role model.Role) error {
	return nil
}

// Saves a role that is linked to an alliance
// alliance_role_map table
func (acc *accessesRepo) SaveAllianceRole(allianceId int64, role model.Role) error {
	return nil
}

// Saves a role that is linked to a corporation
// corporation_role_map table
func (acc *accessesRepo) SaveCorporationRole(corporationId int64, role model.Role) error {
	return nil
}

// Saves a role that is linked to a character
// character_role_map table
func (acc *accessesRepo) SaveCharacterRole(characterId int64, role model.Role) error {
	return nil
}

// Saves a role that is linked to an alliance and a character in an alliance leadership position
// alliance_character_leadership_role_map table
func (acc *accessesRepo) SaveAllianceCharacterLeadershipRole(allianceId int64, role model.Role) error {
	return nil
}

// Saves a role that is linked to a corporation and a character in a corporation leadership position
// corp_character_leadership_role_map table
func (acc *accessesRepo) SaveCorporationCharacterLeadershipRole(corporationId int64, role model.Role) error {
	return nil
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

	roleName, chatRoleName, from := "", "", ""
	for idx := 0; rows.Next(); idx++ {
		rows.Scan(&roleName, &chatRoleName, &from)

		//TODO: Is this innefficient?  Should I be going about growing my array differently?
		roles = append(roles, chatRoleName)
	}

	return roles, nil
}

func (acc *accessesRepo) findByCharacter(character model.Character) map[string]string {
	return nil
}

func (acc *accessesRepo) findByUserId(userId int64) []string {
	return nil
}
