CREATE TABLE users (
    user_id INT GENERATED ALWAYS AS IDENTITY,
    email VARCHAR NOT NULL UNIQUE,
    PRIMARY KEY(user_id)
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE auth_data (
     auth_data_id uuid DEFAULT uuid_generate_v4 (),
     features JSONB,
     user_id INT,
     PRIMARY KEY(auth_data_id),
     CONSTRAINT fk_user
         FOREIGN KEY(user_id)
             REFERENCES users(user_id)
);