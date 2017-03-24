package repository

import (
	"database/sql"
	"github.com/abaeve/auth-srv/model"
)

type AccessesRepository interface {
	SaveAllianceAndCorpLevelRole(allianceId int64, corporationId int64, role model.Role) error
	SaveAllianceRole(allianceId int64, corporationId int64, role model.Role) error
	FindByChatId(chatId string) []string
}

type accesses struct {
	db *sql.DB
}

var roleQuery string = `
SELECT
  role.*,
  'alliance_corp' AS role_from
FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN corporations corp ON (chars.corporation_id = corp.corporation_id)
  JOIN alliances alliance ON (corp.alliance_id = alliance.alliance_id)
  JOIN alliance_corporation_role_map acrm
    ON (alliance.alliance_id = acrm.alliance_id AND corp.corporation_id = acrm.corporation_id)
  JOIN roles role ON (acrm.role_id = role.role_id)
WHERE user.chat_id = '%s'
UNION
SELECT
  role.*,
  'alliance' AS role_from
FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN corporations corp ON (chars.corporation_id = corp.corporation_id)
  JOIN alliances alliance ON (corp.alliance_id = alliance.alliance_id)
  JOIN alliance_role_map arm ON (alliance.alliance_id = arm.alliance_id)
  JOIN roles role ON (arm.role_id = role.role_id)
WHERE user.chat_id = '%s'
UNION
SELECT
  role.*,
  'corp' AS role_from
FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN corporations corp ON (chars.corporation_id = corp.corporation_id)
  JOIN corporation_role_map crm ON (corp.corporation_id = crm.corporation_id)
  JOIN roles role ON (crm.role_id = role.role_id)
WHERE user.chat_id = '%s'
UNION
SELECT
  role.*,
  'character' AS role_from
FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN character_role_map crm ON (chars.corporation_id = crm.character_id)
  JOIN roles role ON (crm.role_id = role.role_id)
WHERE user.chat_id = '%s'
UNION
SELECT
  role.*,
  'alliance_corporation_leadership' AS role_from
FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN corporations corp ON (chars.corporation_id = corp.corporation_id)
  JOIN alliances alliance ON (corp.alliance_id = alliance.alliance_id)
  JOIN alliance_character_leadership_role_map aclrm
    ON (chars.character_id = aclrm.character_id AND alliance.alliance_id = aclrm.alliance_id)
  JOIN roles role ON (aclrm.role_id = role.role_id)
WHERE user.chat_id = '%s'
UNION
SELECT
  role.*,
  'corporation_character_leadership' AS role_from
FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN corp_character_leadership_role_map cclrm
    ON (chars.character_id = cclrm.character_id AND chars.corporation_id = cclrm.corporation_id)
  JOIN roles role ON (cclrm.role_id = role.role_id)
WHERE user.chat_id = '%s'
`

func (acc *accesses) findByUserId ( userId int64 ) []string {
	return nil
}

// Saves a role that is linked to an alliance AND a corporation
func (acc *accesses) SaveAllianceAndCorpLevelRole(allianceId int64, corporationId int64, role model.Role) error {
	return nil
}

func (acc *accesses) SaveAllianceRole(allianceId int64, corporationId int64, role model.Role) error {
	return nil
}

// Will be the main usage in anything automated.  This method will lookup all the available roles for the given
// Chat user id and return them as a map[string]string where the key is the role and the value is the chat
// group to apply.
func (acc *accesses) FindByChatId(chatId string) []string {
	return nil
}

func (acc *accesses) findByCharacter ( character model.Character ) map[string]string {
	return nil
}
