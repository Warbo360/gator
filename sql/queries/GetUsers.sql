-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    $1 = name;
