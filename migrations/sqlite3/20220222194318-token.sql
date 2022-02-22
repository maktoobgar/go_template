
-- +migrate Up
CREATE TABLE rtokens (
    token VARCHAR(511) NOT NULL UNIQUE PRIMARY KEY
);
-- +migrate Down
DROP TABLE rtokens;