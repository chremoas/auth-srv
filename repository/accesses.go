package repository

import (
	"database/sql"
	"github.com/abaeve/auth-srv/model"
)

type AccessesRepository interface {
	FindByChatId(chatId string) []string
}

type accesses struct {
	db *sql.DB
}

var roleQuery string = `
SELECT
  role.*,
  'alliance_corp' AS role_from
FROM user user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN character char ON (ucm.character_id = char.character_id)
  JOIN corporation corp ON (char.corporation_id = corp.corporation_id)
  JOIN alliance alliance ON (corp.alliance_id = alliance.alliance_id)
  JOIN alliance_corporation_role_map acrm
    ON (alliance.alliance_id = acrm.alliance_id AND corp.corporation_id = acrm.corporation_id)
  JOIN role role ON (acrm.role_id = role.role_id)
WHERE user.id = '%s'
UNION
SELECT
  role.*,
  'alliance' AS role_from
FROM user user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN character char ON (ucm.character_id = char.character_id)
  JOIN corporation corp ON (char.corporation_id = corp.corporation_id)
  JOIN alliance alliance ON (corp.alliance_id = alliance.alliance_id)
  JOIN alliance_role_map arm ON (alliance.alliance_id = arm.alliance_id)
  JOIN role role ON (arm.role_id = role.role_id)
WHERE user.id = '%s'
UNION
SELECT
  role.*,
  'corp' AS role_from
FROM user user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN character char ON (ucm.character_id = char.character_id)
  JOIN corporation corp ON (char.corporation_id = corp.corporation_id)
  JOIN corporation_role_map crm ON (corp.corporation_id = crm.corporation_id)
  JOIN role role ON (crm.role_id = role.role_id)
WHERE user.id = '%s'
UNION
SELECT
  role.*,
  'character' AS role_from
FROM user user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN character char ON (ucm.character_id = char.character_id)
  JOIN character_role_map crm ON (char.corporation_id = crm.character_id)
  JOIN role role ON (crm.role_id = role.role_id)
WHERE user.id = '%s'
UNION
SELECT
  role.*,
  'alliance_corporation_leadership' AS role_from
FROM user user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN character char ON (ucm.character_id = char.character_id)
  JOIN corporation corp ON (char.corporation_id = corp.corporation_id)
  JOIN alliance alliance ON (corp.alliance_id = alliance.alliance_id)
  JOIN alliance_character_leadership_role_map aclrm
    ON (char.character_id = aclrm.character_id AND alliance.alliance_id = aclrm.alliance_id)
  JOIN role role ON (aclrm.role_id = role.role_id)
WHERE user.id = '%s'
UNION
SELECT
  role.*,
  'corporation_character_leadership' AS role_from
FROM user user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN character char ON (ucm.character_id = char.character_id)
  JOIN corp_character_leadership_role_map cclrm
    ON (char.character_id = cclrm.character_id AND char.corporation_id = cclrm.corporation_id)
  JOIN role role ON (cclrm.role_id = role.role_id)
WHERE user.id = '%s'
`

func (acc *accesses) findByUserId ( userId int64 ) []string {
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
