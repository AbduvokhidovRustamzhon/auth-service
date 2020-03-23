CREATE TABLE IF NOT EXISTS   users (
   id               BIGSERIAL PRIMARY KEY,
   name             TEXT NOT NULL,
   login            TEXT NOT NULL,
   password         TEXT NOT NULL,
   role             BOOLEAN DEFAULT FALSE,
   removed          BOOLEAN DEFAULT FALSE
);





--

SELECT id FROM users WHERE login = 'akram8008';

---
SELECT id FROM users WHERE login = "sasas";
