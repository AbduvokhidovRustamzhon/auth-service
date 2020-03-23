package main


const createTable = `CREATE TABLE IF NOT EXISTS   users (
   id               BIGSERIAL PRIMARY KEY,
   name             TEXT NOT NULL,
   login            TEXT NOT NULL,
   password         TEXT NOT NULL,
   role             TEXT NOT NULL,
   removed          BOOLEAN DEFAULT FALSE
);
`