CREATE TABLE user_character_map
(
  user_id      BIGINT(20) NOT NULL,
  character_id BIGINT(20) NOT NULL,
  CONSTRAINT `PRIMARY` PRIMARY KEY (user_id, character_id),
  CONSTRAINT user_character_map__user_fk FOREIGN KEY (user_id) REFERENCES user (user_id),
  CONSTRAINT user_character_map__character_fk FOREIGN KEY (character_id) REFERENCES `character` (character_id)
);
CREATE INDEX user_character_map__character_fk
  ON user_character_map (character_id);


CREATE TABLE authentication_scope_character_map
(
  character_id            BIGINT(20) NOT NULL,
  authentication_scope_id BIGINT(20) NOT NULL,
  CONSTRAINT `PRIMARY` PRIMARY KEY (character_id, authentication_scope_id),
  CONSTRAINT scope_character_map__character_fk FOREIGN KEY (character_id) REFERENCES `character` (character_id),
  CONSTRAINT scope_character_map__scope_fk FOREIGN KEY (authentication_scope_id) REFERENCES authentication_scope (authentication_scope_id)
);
CREATE INDEX scope_character_map__scope_fk
  ON authentication_scope_character_map (authentication_scope_id);


CREATE TABLE alliance_role_map
(
  role_id     BIGINT(20) NOT NULL,
  alliance_id BIGINT(20) NOT NULL,
  CONSTRAINT `PRIMARY` PRIMARY KEY (role_id, alliance_id),
  CONSTRAINT alliance_role_map__role_fk FOREIGN KEY (role_id) REFERENCES role (role_id),
  CONSTRAINT alliance_role_map__alliance_fk FOREIGN KEY (alliance_id) REFERENCES alliance (alliance_id)
);
CREATE INDEX alliance_role_map__alliance_fk
  ON alliance_role_map (alliance_id);


CREATE TABLE alliance_corporation_role_map
(
  alliance_id    BIGINT(20) NOT NULL,
  corporation_id BIGINT(20) NOT NULL,
  role_id        BIGINT(20) NOT NULL,
  CONSTRAINT `PRIMARY` PRIMARY KEY (alliance_id, corporation_id, role_id),
  CONSTRAINT alliance_corporation_role_map__alliance_fk FOREIGN KEY (alliance_id) REFERENCES alliance (alliance_id),
  CONSTRAINT alliance_corporation_role_map__corporation_fk FOREIGN KEY (corporation_id) REFERENCES corporation (corporation_id),
  CONSTRAINT alliance_corporation_role_map__role_fk FOREIGN KEY (role_id) REFERENCES role (role_id)
);
CREATE INDEX alliance_corporation_role_map__corporation_fk
  ON alliance_corporation_role_map (corporation_id);
CREATE INDEX alliance_corporation_role_map__role_fk
  ON alliance_corporation_role_map (role_id);


CREATE TABLE authentication_scope
(
  authentication_scope_id   BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  authentication_scope_name VARCHAR(255)           NOT NULL
);


CREATE TABLE corporation
(
  corporation_id     BIGINT(20) PRIMARY KEY             NOT NULL,
  corporation_name   VARCHAR(100)                       NOT NULL,
  alliance_id        BIGINT(20),
  inserted_dt        DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_dt         DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  corporation_ticker VARCHAR(5)                         NOT NULL,
  CONSTRAINT corporation_alliance_alliance_id_fk FOREIGN KEY (alliance_id) REFERENCES alliance (alliance_id)
);
CREATE INDEX corporation_alliance_alliance_id_fk
  ON corporation (alliance_id);


CREATE TABLE corp_character_leadership_role_map
(
  corporation_id BIGINT(20) NOT NULL,
  character_id   BIGINT(20) NOT NULL,
  role_id        BIGINT(20) NOT NULL,
  CONSTRAINT `PRIMARY` PRIMARY KEY (corporation_id, character_id, role_id),
  CONSTRAINT leadership_role__corporation_fk FOREIGN KEY (corporation_id) REFERENCES corporation (corporation_id),
  CONSTRAINT leadership_role__character_fk FOREIGN KEY (character_id) REFERENCES `character` (character_id),
  CONSTRAINT leadership_role__role_fk FOREIGN KEY (role_id) REFERENCES role (role_id)
);
CREATE INDEX leadership_role__character_fk
  ON corp_character_leadership_role_map (character_id);
CREATE INDEX leadership_role__role_fk
  ON corp_character_leadership_role_map (role_id);


CREATE TABLE user
(
  user_id BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  id      VARCHAR(255)           NOT NULL
  COMMENT 'Should be large enough to hold id''s from other systems such as twitter, github, discord, slack, etc'
);


CREATE TABLE character_role_map
(
  character_id BIGINT(20) NOT NULL,
  role_id      BIGINT(20) NOT NULL,
  CONSTRAINT `PRIMARY` PRIMARY KEY (character_id, role_id),
  CONSTRAINT character_role_map__character_fk FOREIGN KEY (character_id) REFERENCES `character` (character_id),
  CONSTRAINT character_role_map__role_fk FOREIGN KEY (role_id) REFERENCES role (role_id)
);
CREATE INDEX character_role_map__role_fk
  ON character_role_map (role_id);


CREATE TABLE alliance_character_leadership_role_map
(
  alliance_id  BIGINT(20) NOT NULL,
  character_id BIGINT(20) NOT NULL,
  role_id      BIGINT(20) NOT NULL,
  CONSTRAINT `PRIMARY` PRIMARY KEY (alliance_id, character_id, role_id),
  CONSTRAINT alliance_leadership__alliance_fk FOREIGN KEY (alliance_id) REFERENCES alliance (alliance_id),
  CONSTRAINT alliance_leadership__character_fk FOREIGN KEY (character_id) REFERENCES `character` (character_id),
  CONSTRAINT alliance_leadership__role_fk FOREIGN KEY (role_id) REFERENCES role (role_id)
);
CREATE INDEX alliance_leadership__character_fk
  ON alliance_character_leadership_role_map (character_id);
CREATE INDEX alliance_leadership__role_fk
  ON alliance_character_leadership_role_map (role_id);


CREATE TABLE role
(
  role_name         VARCHAR(70)                        NOT NULL,
  inserted_dt       DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_dt        DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  role_id           BIGINT(20) PRIMARY KEY             NOT NULL AUTO_INCREMENT,
  chatservice_group VARCHAR(70)
);
CREATE UNIQUE INDEX role_role_name_uindex
  ON role (role_name);


CREATE TABLE alliance
(
  alliance_id     BIGINT(20) PRIMARY KEY             NOT NULL,
  alliance_name   VARCHAR(100)                       NOT NULL,
  alliance_ticker VARCHAR(5)                         NOT NULL,
  inserted_dt     DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_dt      DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE TABLE corporation_role_map
(
  role_id        BIGINT(20) NOT NULL,
  corporation_id BIGINT(20) NOT NULL,
  CONSTRAINT `PRIMARY` PRIMARY KEY (role_id, corporation_id),
  CONSTRAINT corporation_role_map__role_fk FOREIGN KEY (role_id) REFERENCES role (role_id),
  CONSTRAINT corporation_role_map__corporation_fk FOREIGN KEY (corporation_id) REFERENCES corporation (corporation_id)
);
CREATE INDEX corporation_role_map__corporation_fk
  ON corporation_role_map (corporation_id);


CREATE TABLE authentication_code
(
  character_id        BIGINT(20) PRIMARY KEY NOT NULL,
  authentication_code VARCHAR(20)            NOT NULL,
  CONSTRAINT authentication_code_character_character_id_fk FOREIGN KEY (character_id) REFERENCES `character` (character_id)
);


CREATE TABLE `character`
(
  character_id   BIGINT(20) PRIMARY KEY             NOT NULL,
  character_name VARCHAR(100)                       NOT NULL,
  inserted_dt    DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  udpated_dt     DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  corporation_id BIGINT(20)                         NOT NULL,
  token          VARCHAR(255)                       NOT NULL,
  CONSTRAINT character_corporation_corporation_id_fk FOREIGN KEY (corporation_id) REFERENCES corporation (corporation_id)
);
CREATE INDEX character_corporation_corporation_id_fk
  ON `character` (corporation_id);