
-- +migrate Up
CREATE TABLE users (
    id int AUTO_INCREMENT NOT NULL UNIQUE PRIMARY KEY,
    username varchar(64) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    display_name varchar(128) NOT NULL,
    first_name varchar(128) NULL,
    last_name varchar(128) NULL,
    age int NULL
);
-- +migrate Down
DROP TABLE users;
