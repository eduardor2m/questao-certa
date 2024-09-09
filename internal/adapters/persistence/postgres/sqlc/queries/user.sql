-- name: Register :exec

INSERT INTO "user" ("id", "name", "email", "password", "admin", "is_active", "created_at", "updated_at") VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: Authenticate :one

SELECT * FROM "user" WHERE "email" = $1 LIMIT 1;

-- name: Delete :exec

UPDATE "user" SET "is_active" = false WHERE "email" = $1 and "is_active" = true and "admin" = false and "name" = $2;

-- name: List :many

SELECT * FROM "user" ORDER BY "created_at" DESC;

-- name: FindByEmail :one

SELECT * FROM "user" WHERE "email" = $1 LIMIT 1;

-- name: CheckType :one

SELECT * FROM "user" WHERE "id" = $1 LIMIT 1;
