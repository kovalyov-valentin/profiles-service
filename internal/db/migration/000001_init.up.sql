CREATE TABLE IF NOT EXISTS profiles (
    profiles_uid VARCHAR(255) PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    surname VARCHAR(255),
    patronymic VARCHAR(255),
    age VARCHAR(255),
    gender int,
    country VARCHAR(255)
)