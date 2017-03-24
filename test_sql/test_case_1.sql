/*
Test scenario
'Rando Corp 1 Member' has the following roles
user{ id: '1234567890'}
CORP_ALLIANCE = RANDO_ROLE_1, RANDO_ROLE_2
ALLIANCE      = RANDO_ALLIANCE_ROLE
CORP          = RANDO_ROLE_4
CHAR          = SUPER_AWESOME_ADMIN
ALLIANCE LEAD = RANDO_ALLIANCE_1_LEADERSHIP_ROLE
CORP LEAD     = RANDO_CORP_1_LEADERSHIP_ROLE

'Rando Corp 2 Member' has the following roles
user{id: '1234567890'}
CORP_ALLIANCE = RANDO_ROLE_2
ALLIANCE      = RANDO_ALLIANCE_ROLE
CORP          = RANDO_ROLE_5, RANDO_ROLE_6
CHAR          =
ALLIANCE LEAD = RANDO_ALLIANCE_1_LEADERSHIP_ROLE
CORP LEAD     = RANDO_CORP_2_LEADERSHIP_ROLE

'Rando Corp 3 Member' has the following roles
user{id: '1234567891'}
CORP_ALLIANCE = RANDO_ROLE_3
ALLIANCE      = RANDO_ALLIANCE_ROLE
CORP          =
CHAR          =
*/
/*<editor-fold desc="Users">*/
INSERT INTO users (chat_id) VALUES ('1234567890');
/* Represents 'Rando Corp 1 Member' and 'Rando Corp 2 Member'*/
INSERT INTO users (chat_id) VALUES ('1234567891');
/* Represents the 'Rando Corp 3 Member' */
/*</editor-fold>*/

/*<editor-fold desc="Roles">*/
INSERT INTO roles (role_name, chatservice_group) VALUES ('RANDO_ROLE_1', NULL);
INSERT INTO roles (role_name, chatservice_group) VALUES ('RANDO_ROLE_2', 'RANDO_ROLE_2');
INSERT INTO roles (role_name, chatservice_group) VALUES ('RANDO_ROLE_3', 'RANDO_ROLE_3');
INSERT INTO roles (role_name, chatservice_group) VALUES ('RANDO_ROLE_4', 'RANDO_ROLE_4');
INSERT INTO roles (role_name, chatservice_group) VALUES ('RANDO_ROLE_5', 'RANDO_ROLE_5');
INSERT INTO roles (role_name, chatservice_group) VALUES ('RANDO_ROLE_6', 'RANDO_ROLE_6');
INSERT INTO roles (role_name, chatservice_group) VALUES ('SUPER_AWESOME_ADMIN', 'SUPER_AWESOME_ADMIN');
INSERT INTO roles (role_name, chatservice_group) VALUES ('RANDO_ALLIANCE_ROLE', 'RANDO_ALLIANCE_ROLE');
INSERT INTO roles (role_name, chatservice_group) VALUES ('CORP_2_LEADERSHIP', 'CORP_2_LEADERSHIP');
INSERT INTO roles (role_name, chatservice_group) VALUES ('ALLIANCE_LEADERSHIP', 'ALLIANCE_LEADERSHIP');
INSERT INTO roles (role_name, chatservice_group) VALUES ('RANDO_CORP_1_LEADERSHIP_ROLE', 'RANDO_CORP_1_LEADERSHIP_ROLE');
INSERT INTO roles (role_name, chatservice_group) VALUES ('RANDO_CORP_2_LEADERSHIP_ROLE', 'RANDO_CORP_1_LEADERSHIP_ROLE');
INSERT INTO roles (role_name, chatservice_group) VALUES ('RANDO_CORP_3_LEADERSHIP_ROLE', 'RANDO_CORP_1_LEADERSHIP_ROLE');
INSERT INTO roles (role_name, chatservice_group)
VALUES ('RANDO_ALLIANCE_1_LEADERSHIP_ROLE', 'RANDO_ALLIANCE_1_LEADERSHIP_ROLE');
/*</editor-fold>*/

/*<editor-fold desc="Alliances, corporations and characters">*/
INSERT INTO alliances (alliance_id, alliance_name, alliance_ticker) VALUES (1, 'Rando Alliance 1', 'RA1');
INSERT INTO corporations (corporation_id, corporation_name, corporation_ticker, alliance_id)
VALUES (1, 'Rando Corp 1', 'R1', 1);
INSERT INTO corporations (corporation_id, corporation_name, corporation_ticker, alliance_id)
VALUES (2, 'Rando Corp 2', 'R2', 1);
INSERT INTO corporations (corporation_id, corporation_name, corporation_ticker, alliance_id)
VALUES (3, 'Rando Corp 3', 'R3', 1);
INSERT INTO characters (character_id, character_name, corporation_id, token) VALUES (1, 'Rando Corp 1 Member', 1, '');
INSERT INTO characters (character_id, character_name, corporation_id, token) VALUES (2, 'Rando Corp 2 Member', 2, '');
INSERT INTO characters (character_id, character_name, corporation_id, token) VALUES (3, 'Rando Corp 3 Member', 3, '');
/*</editor-fold>*/

/*<editor-fold desc="User Mappings">*/
INSERT INTO user_character_map (user_id, character_id) VALUES
  (
    (SELECT user_id
     FROM users
     WHERE chat_id = '1234567890'),
    (SELECT character_id
     FROM characters
     WHERE character_name = 'Rando Corp 1 Member')
  );
INSERT INTO user_character_map (user_id, character_id) VALUES
  (
    (SELECT user_id
     FROM users
     WHERE chat_id = '1234567890'),
    (SELECT character_id
     FROM characters
     WHERE character_name = 'Rando Corp 2 Member')
  );
INSERT INTO user_character_map (user_id, character_id) VALUES
  (
    (SELECT user_id
     FROM users
     WHERE chat_id = '1234567891'),
    (SELECT character_id
     FROM characters
     WHERE character_name = 'Rando Corp 3 Member')
  );
/*</editor-fold>*/

/*<editor-fold desc="Role Mappings">*/
INSERT INTO alliance_corporation_role_map (alliance_id, corporation_id, role_id) VALUES
  (
    (SELECT alliance_id
     FROM alliances
     WHERE alliance_name = 'Rando Alliance 1'),
    (SELECT corporation_id
     FROM corporations
     WHERE corporation_name = 'Rando Corp 1'),
    (SELECT role_id
     FROM roles
     WHERE role_name = 'RANDO_ROLE_1')
  );
INSERT INTO alliance_corporation_role_map (alliance_id, corporation_id, role_id) VALUES
  (
    (SELECT alliance_id
     FROM alliances
     WHERE alliance_name = 'Rando Alliance 1'),
    (SELECT corporation_id
     FROM corporations
     WHERE corporation_name = 'Rando Corp 1'),
    (SELECT role_id
     FROM roles
     WHERE role_name = 'RANDO_ROLE_2')
  );
INSERT INTO alliance_corporation_role_map (alliance_id, corporation_id, role_id) VALUES
  (
    (SELECT alliance_id
     FROM alliances
     WHERE alliance_name = 'Rando Alliance 1'),
    (SELECT corporation_id
     FROM corporations
     WHERE corporation_name = 'Rando Corp 3'),
    (SELECT role_id
     FROM roles
     WHERE role_name = 'RANDO_ROLE_3')
  );
INSERT INTO alliance_corporation_role_map (alliance_id, corporation_id, role_id) VALUES
  (
    (SELECT alliance_id
     FROM alliances
     WHERE alliance_name = 'Rando Alliance 1'),
    (SELECT corporation_id
     FROM corporations
     WHERE corporation_name = 'Rando Corp 2'),
    (SELECT role_id
     FROM roles
     WHERE role_name = 'RANDO_ROLE_6')
  );
INSERT INTO alliance_role_map (role_id, alliance_id) VALUES
  (
    (SELECT role_id
     FROM roles
     WHERE role_name = 'RANDO_ALLIANCE_ROLE'),
    (SELECT alliance_id
     FROM alliances
     WHERE alliance_name = 'Rando Alliance 1')
  );
INSERT INTO corporation_role_map (role_id, corporation_id) VALUES
  (
    (SELECT role_id
     FROM roles
     WHERE role_name = 'RANDO_ROLE_4'),
    (SELECT corporation_id
     FROM corporations
     WHERE corporation_name = 'Rando Corp 1')
  );
INSERT INTO corporation_role_map (role_id, corporation_id) VALUES
  (
    (SELECT role_id
     FROM roles
     WHERE role_name = 'RANDO_ROLE_5'),
    (SELECT corporation_id
     FROM corporations
     WHERE corporation_name = 'Rando Corp 2')
  );
INSERT INTO character_role_map (character_id, role_id) VALUES
  (
    (SELECT character_id
     FROM characters
     WHERE character_name = 'Rando Corp 1 Member'),
    (SELECT role_id
     FROM roles
     WHERE role_name = 'SUPER_AWESOME_ADMIN')
  );
INSERT INTO alliance_character_leadership_role_map (alliance_id, character_id, role_id) VALUES
  (
    (SELECT alliance_id
     FROM alliances
     WHERE alliance_name = 'Rando Alliance 1'),
    (SELECT character_id
     FROM characters
     WHERE character_name = 'Rando Corp 1 Member'),
    (SELECT role_id
     FROM roles
     WHERE role_name = 'RANDO_ALLIANCE_1_LEADERSHIP_ROLE')
  );
INSERT INTO alliance_character_leadership_role_map (alliance_id, character_id, role_id) VALUES
  (
    (SELECT alliance_id
     FROM alliances
     WHERE alliance_name = 'Rando Alliance 1'),
    (SELECT character_id
     FROM characters
     WHERE character_name = 'Rando Corp 2 Member'),
    (SELECT role_id
     FROM roles
     WHERE role_name = 'RANDO_ALLIANCE_1_LEADERSHIP_ROLE')
  );
INSERT INTO corp_character_leadership_role_map (corporation_id, character_id, role_id) VALUES
  (
    (SELECT corporation_id
     FROM corporations
     WHERE corporation_name = 'Rando Corp 2'),
    (SELECT character_id
     FROM characters
     WHERE character_name = 'Rando Corp 2 Member'),
    (SELECT role_id
     FROM roles
     WHERE role_name = 'RANDO_CORP_1_LEADERSHIP_ROLE')
  );
/*</editor-fold>*/

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
WHERE user.chat_id = '1234567890'
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
WHERE user.chat_id = '1234567890'
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
WHERE user.chat_id = '1234567890'
UNION
SELECT
  role.*,
  'character' AS role_from
FROM users user
  JOIN user_character_map ucm ON (user.user_id = ucm.user_id)
  JOIN characters chars ON (ucm.character_id = chars.character_id)
  JOIN character_role_map crm ON (chars.corporation_id = crm.character_id)
  JOIN roles role ON (crm.role_id = role.role_id)
WHERE user.chat_id = '1234567890'
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
WHERE user.chat_id = '1234567890'
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
WHERE user.chat_id = '1234567890';

/*
Starting over
 */
/*
 */
DELETE FROM alliance_corporation_role_map;
DELETE FROM alliance_role_map;
DELETE FROM corporation_role_map;
DELETE FROM character_role_map;
DELETE FROM user_character_map;
DELETE FROM alliance_character_leadership_role_map;
DELETE FROM corp_character_leadership_role_map;
DELETE FROM characters;
DELETE FROM corporations;
DELETE FROM alliances;
DELETE FROM roles;
DELETE FROM users;
