-- name: GetUrlById :one
SELECT * FROM urls
WHERE id = ? LIMIT 1;

-- name: GetUrlByCode :one
SELECT * FROM urls
WHERE code = ? LIMIT 1;

-- name: CreateUrl :one
INSERT INTO urls
  ( url, code, md5 )
VALUES
  ( ?, ?, ? )
ON CONFLICT (md5)
DO UPDATE SET md5 = EXCLUDED.md5
RETURNING *;

-- name: IncrementUrlHitsById :exec
UPDATE urls
SET hits = hits + 1,
    last_used = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteUrlByLastUsed :exec
DELETE FROM urls
WHERE last_used < ?;
