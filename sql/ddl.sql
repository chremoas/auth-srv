SET DATABASE = "chremoas-dev";

CREATE TABLE authentication_scopes
(
  authentication_scope_id   BIGSERIAL PRIMARY KEY NOT NULL,
  authentication_scope_name VARCHAR(255)           NOT NULL
);


CREATE TABLE users
(
  user_id BIGSERIAL     PRIMARY KEY NOT NULL,
  chat_id VARCHAR(255)           NOT NULL
);


CREATE TABLE alliances
(
  alliance_id     BIGINT PRIMARY KEY             NOT NULL,
  alliance_name   VARCHAR(100)                       NOT NULL,
  alliance_ticker VARCHAR(5),
  inserted_dt     TIMESTAMP                          NOT NULL,
  updated_dt      TIMESTAMP                          NOT NULL
);


CREATE TABLE corporations
(
  corporation_id     BIGINT PRIMARY KEY             NOT NULL,
  corporation_name   VARCHAR(100)                       NOT NULL,
  alliance_id        BIGINT,
  inserted_dt        TIMESTAMP                          NOT NULL,
  updated_dt         TIMESTAMP                          NOT NULL,
  corporation_ticker VARCHAR(5),
  CONSTRAINT corporation_alliance_alliance_id_fk FOREIGN KEY (alliance_id) REFERENCES alliances (alliance_id)
);


CREATE TABLE characters
(
  character_id   BIGINT PRIMARY KEY             NOT NULL,
  character_name VARCHAR(100)                       NOT NULL,
  inserted_dt    TIMESTAMP                          NOT NULL,
  updated_dt     TIMESTAMP                          NOT NULL,
  corporation_id BIGINT                         NOT NULL,
  token          VARCHAR(255)                       NOT NULL,
  CONSTRAINT character_corporation_corporation_id_fk FOREIGN KEY (corporation_id) REFERENCES corporations (corporation_id)
);


CREATE TABLE authentication_codes
(
  character_id        BIGINT PRIMARY KEY NOT NULL,
  authentication_code VARCHAR(20)            NOT NULL,
  is_used             BOOLEAN                NOT NULL DEFAULT FALSE,
  CONSTRAINT authentication_code_character_character_id_fk FOREIGN KEY (character_id) REFERENCES characters (character_id)
);


CREATE TABLE user_character_map
(
  user_id      BIGINT NOT NULL,
  character_id BIGINT NOT NULL,
  PRIMARY KEY (user_id, character_id),
  CONSTRAINT user_character_map__user_fk FOREIGN KEY (user_id) REFERENCES users (user_id),
  CONSTRAINT user_character_map__character_fk FOREIGN KEY (character_id) REFERENCES characters (character_id)
);


CREATE TABLE authentication_scope_character_map
(
  character_id            BIGINT NOT NULL,
  authentication_scope_id BIGINT NOT NULL,
  PRIMARY KEY (character_id, authentication_scope_id),
  CONSTRAINT scope_character_map__character_fk FOREIGN KEY (character_id) REFERENCES characters (character_id),
  CONSTRAINT scope_character_map__scope_fk FOREIGN KEY (authentication_scope_id) REFERENCES authentication_scopes (authentication_scope_id)
);
