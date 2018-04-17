// FIX: I don't even know right now. Will circle back. -brian
package repository

import (
	"github.com/jmoiron/sqlx"
)

type AccessesRepository interface {
	FindByChatId(chatId string) ([]string, error)
	GetMembership() (members []membership, err error)
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

var getMembership = `
SELECT a.alliance_name
    , a.alliance_ticker
    , c.corporation_name
    , c.corporation_ticker
    , u.chat_id
FROM corporations c
INNER JOIN alliances a
    ON c.alliance_id = a.alliance_id
INNER JOIN characters ch
    ON c.corporation_id = ch.corporation_id
INNER JOIN user_character_map ucm
    ON ch.character_id = ucm.character_id
INNER JOIN users u
    ON u.user_id = ucm.user_id;
`

type membership struct {
	AllianceName   string `db:"alliance_name"`
	AllianceTicker string `db:"alliance_ticker"`
	CorpName       string `db:"corporation_name"`
	CorpTicker     string `db:"corporation_ticker"`
	ChatId         string `db:"chat_id"`
}

func (acc *accessesRepo) GetMembership() (members []membership, err error) {
	rows, err := acc.db.Queryx(getMembership)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for idx := 0; rows.Next(); idx++ {
		var m membership

		if err = rows.StructScan(&m); err != nil {
			return nil, err
		}
		members = append(members, m)
	}

	return members, nil
}
