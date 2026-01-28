CREATE TABLE users
(
    id serial primary key,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE certificates
(
    id serial primary key,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    url varchar(255) not null,
    valid_from timestamp not null,
    valid_to timestamp not null
);
