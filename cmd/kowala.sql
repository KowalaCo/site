CREATE TABLE users (
    id char(20) PRIMARY KEY NOT NULL,
    username varchar(16) NOT NULL UNIQUE,
    token text NOT NULL,
    role smallint NOT NULL DEFAULT 0,
    password text NOT NULL,
    verified boolean NOT NULL DEFAULT false
);