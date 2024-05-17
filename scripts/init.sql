DROP TABLE IF EXISTS Parkings;
DROP TABLE IF EXISTS Blocks;
DROP TABLE IF EXISTS Slots;
DROP TABLE IF EXISTS Vehicles;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS LoginToken;

DROP TYPE IF EXISTS UserRole;

CREATE TYPE UserRole AS ENUM ('root', 'manager', 'admin');

CREATE TABLE Parkings (
    id SERIAL PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL,
    capacity INTEGER NOT NULL,
    region VARCHAR
);

CREATE TABLE Blocks (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    parking_id INTEGER REFERENCES Parkings(id) NOT NULL,
    UNIQUE (name, parking_id)
);

CREATE TABLE Slots (
    id SERIAL PRIMARY KEY,
    number INTEGER NOT NULL,
    parking_id INTEGER REFERENCES Parkings(id) NOT NULL,
    block_id INTEGER REFERENCES Blocks(id) NOT NULL,
    is_empty BOOLEAN DEFAULT TRUE,
    UNIQUE (number, block_id),
    CHECK (number >= 1 AND number <= 50)
);

CREATE TABLE Vehicles (
    id SERIAL PRIMARY KEY,
    wincode VARCHAR UNIQUE NOT NULL,
    parking_id INTEGER REFERENCES Parkings(id) NOT NULL,
    block_id INTEGER REFERENCES Blocks(id) NOT NULL,
    slot_id INTEGER REFERENCES Slots(id) NOT NULL
);

CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    role UserRole NOT NULL,
    first_name VARCHAR,
    last_name VARCHAR,
    parking_id INTEGER REFERENCES Parkings(id)
);

CREATE TABLE LoginTokens (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id) NOT NULL,
    content VARCHAR NOT NULL,
    UNIQUE (user_id, content)
);
