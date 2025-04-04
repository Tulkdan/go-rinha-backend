CREATE TABLE IF NOT EXISTS people (
    id UUID PRIMARY KEY
    , name TEXT
    , nickname TEXT
    , birthdate TIMESTAMP
    , stacks TEXT[]
);
