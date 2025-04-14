-- name: GetPerson :one
SELECT *
FROM people
WHERE id = $1
LIMIT 1;

-- name: CreatePerson :one
INSERT INTO people (id, name, nickname, birthdate, stacks)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: SearchPerson :many
SELECT
    id
    , name
    , nickname
    , birthdate
    , stacks
FROM people
WHERE LOWER(name) LIKE '%' || $1 || '%'
   OR LOWER(nickname) LIKE '%' || $1 || '%'
   OR LOWER(ARRAY_TO_STRING(stacks, ',')) LIKE '%' || $1 || '%';

-- name: CountAllPeople :one
SELECT COUNT(*) as qtt
FROM people;
