CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE users (
    id char(20) PRIMARY KEY NOT NULL,
    username citext NOT NULL UNIQUE,
    role tinyint NOT NULL DEFAULT 0,
    password text NOT NULL,
    verified boolean NOT NULL DEFAULT false
);