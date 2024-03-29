
-- +migrate Up
CREATE TABLE users (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    username varchar(64) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    display_name varchar(128) NULL,
    first_name varchar(128) NULL,
    last_name varchar(128) NULL,
    joined_date DATETIME NOT NULL
);
-- +migrate Down
DROP TABLE users;
