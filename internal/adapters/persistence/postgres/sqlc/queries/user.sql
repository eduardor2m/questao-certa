-- name: SignUp :exec

INSERT INTO "user" ("id", "name", "email", "password", "admin", "created_at", "updated_at") VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: SignIn :one

SELECT * FROM "user" WHERE "email" = $1 LIMIT 1;

-- name: VerifyUserIsLoggedOrAdmin :one

SELECT * FROM "user" WHERE "id" = $1 LIMIT 1;
