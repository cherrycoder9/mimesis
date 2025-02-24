CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER CHECK (age >= 0),
    email TEXT UNIQUE
);