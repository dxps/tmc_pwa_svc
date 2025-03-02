CREATE TABLE users (
    id          BIGSERIAL   PRIMARY KEY,
    eid         VARCHAR(36) NOT NULL,
    username    TEXT        UNIQUE NOT NULL,
    name        TEXT        NOT NULL,
    email       TEXT        UNIQUE NOT NULL,
    password    TEXT        NOT NULL
);
