-- name: GetPerson :one
SELECT *
FROM people
WHERE id = $1
LIMIT 1;

-- name: CreatePerson :one
INSERT INTO people (name, nickname, birthdate, stacks)
VALUES ($1, $2, $3, $4)
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
   OR LOWER(nickname) LIKE '%' || $2 || '%'
   OR LOWER(stacks) LIKE '%' || $3 || '%';

-- name: CountAllPeople :one
SELECT COUNT(*) as qtt
FROM people;
