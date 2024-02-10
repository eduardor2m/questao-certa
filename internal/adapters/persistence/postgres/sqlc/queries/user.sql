-- name: SignUp :exec

INSERT INTO "user" ("id", "name", "email", "password", "admin", "created_at", "updated_at") VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: SignIn :one

SELECT * FROM "user" WHERE "email" = $1 LIMIT 1;

-- name: DeleteUserTest :exec

DELETE FROM "user" WHERE "email" = $1 AND "name" = $2;

-- name: ListUsers :many

SELECT * FROM "user" ORDER BY "created_at" DESC;

-- name: GetUserByEmail :one

SELECT * FROM "user" WHERE "email" = $1 LIMIT 1;

-- name: VerifyUserIsLoggedOrAdmin :one

SELECT * FROM "user" WHERE "id" = $1 LIMIT 1;
