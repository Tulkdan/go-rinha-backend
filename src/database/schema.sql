CREATE TABLE IF NOT EXISTS people (
    id TEXT NOT NULL PRIMARY KEY,
    name TEXT,
    nickname TEXT,
    birthdate TIMESTAMP,
    stacks JSONB
);
