CREATE TABLE users (
    id bigserial NOT NULL PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    encrypted_password VARCHAR NOT NULL
)