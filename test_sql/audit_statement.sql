SELECT
  role.role_name,
  role.chatservice_group,
  alli.alliance_name,
  corp.corporation_name,
  chars.character_name,
  'alliance_corp' AS role_from
FROM alliances alli
  JOIN corporations corp ON (alli.alliance_id = corp.alliance_id)
  JOIN characters chars ON (corp.corporation_id = chars.corporation_id)
  JOIN alliance_corporation_role_map acrm
    ON (alli.alliance_id = acrm.alliance_id AND corp.corporation_id = acrm.corporation_id)
  JOIN roles role ON (acrm.role_id = role.role_id)
UNION
SELECT
  role.role_name,
  role.chatservice_group,
  alli.alliance_name,
  corp.corporation_name,
  chars.character_name,
  'alliance' AS role_from
FROM alliances alli
  JOIN corporations corp ON (alli.alliance_id = corp.alliance_id)
  JOIN characters chars ON (corp.corporation_id = chars.corporation_id)
  JOIN alliance_role_map arm ON (alli.alliance_id = arm.alliance_id)
  JOIN roles role ON (arm.role_id = role.role_id)
UNION
SELECT
  role.role_name,
  role.chatservice_group,
  alli.alliance_name,
  corp.corporation_name,
  chars.character_name,
  'corporation' AS role_from
FROM alliances alli
  JOIN corporations corp ON (alli.alliance_id = corp.alliance_id)
  JOIN characters chars ON (corp.corporation_id = chars.corporation_id)
  JOIN corporation_role_map crm ON (corp.corporation_id = crm.corporation_id)
  JOIN roles role ON (crm.role_id = role.role_id)
UNION
SELECT
  role.role_name,
  role.chatservice_group,
  alli.alliance_name,
  corp.corporation_name,
  chars.character_name,
  'character' AS role_from
FROM alliances alli
  JOIN corporations corp ON (alli.alliance_id = corp.alliance_id)
  JOIN characters chars ON (corp.corporation_id = chars.corporation_id)
  JOIN character_role_map crm ON (chars.character_id = crm.character_id)
  JOIN roles role ON (crm.role_id = role.role_id)
UNION
SELECT
  role.role_name,
  role.chatservice_group,
  alli.alliance_name,
  corp.corporation_name,
  chars.character_name,
  'alliance_character_leadership' AS role_from
FROM alliances alli
  JOIN corporations corp ON (alli.alliance_id = corp.alliance_id)
  JOIN characters chars ON (corp.corporation_id = chars.corporation_id)
  JOIN alliance_character_leadership_role_map aclrm
    ON (chars.character_id = aclrm.character_id AND alli.alliance_id = aclrm.alliance_id)
  JOIN roles role ON (aclrm.role_id = role.role_id)
UNION
SELECT
  role.role_name,
  role.chatservice_group,
  alli.alliance_name,
  corp.corporation_name,
  chars.character_name,
  'corporation_character_leadership' AS role_from
FROM alliances alli
  JOIN corporations corp ON (alli.alliance_id = corp.alliance_id)
  JOIN characters chars ON (corp.corporation_id = chars.corporation_id)
  JOIN corp_character_leadership_role_map cclrm
    ON (chars.character_id = cclrm.character_id AND corp.corporation_id = cclrm.corporation_id)
  JOIN roles role ON (cclrm.role_id = role.role_id);