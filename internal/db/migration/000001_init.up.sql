CREATE TABLE IF NOT EXISTS users (
    id serial not null unique,
    name VARCHAR(255) not null,
    surname VARCHAR(255) not null,
    patronymic VARCHAR(255),
    age int,
    gender VARCHAR(255),
    nationality VARCHAR(255)
)