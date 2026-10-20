-- name: GetFeed :one

SELECT
    feeds.id
FROM
    feeds
WHERE
    feeds.url = $1;
