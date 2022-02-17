
-- +migrate Up
CREATE TABLE csrfs (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    key varchar(255) NOT NULL UNIQUE,
    value varchar(255) NOT NULL,
    expire_date DATETIME NOT NULL
);
-- +migrate Down
DROP TABLE csrfs;
