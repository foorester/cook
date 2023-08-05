-- name: SelectUserByID :one
SELECT id, username, name, email, password
FROM users
WHERE id = $1;

-- name: SelectAllUsers :many
SELECT id, username, name, email, password
FROM users;

-- name: InsertUser :one
INSERT INTO users (id, username, name, email, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET username = $1, name = $2, email = $3, password = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: SelectBookByID :one
SELECT b.id, b.name, b.description, u.id, u.username, u.name, u.email, u.password
FROM books b
         JOIN users u ON b.owner_id = u.id
WHERE b.id = $1;

-- name: SelectAllBooks :many
SELECT b.id, b.name, b.description, b.owner_id, b.created_at, b.updated_at
FROM books b
         JOIN users u ON b.owner_id = u.id
WHERE b.owner_id = $1;

-- name: InsertBook :one
INSERT INTO books (id, name, description, owner_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateBook :one
UPDATE books
SET name = $2, description = $3, owner_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: SelectRecipeByID :one
SELECT r.id, r.name, r.description, b.id, b.name, b.description, u.id, u.username, u.name, u.email, u.password
FROM recipes r
         JOIN books b ON r.book_id = b.id
         JOIN users u ON b.owner_id = u.id
WHERE r.id = $1;

-- name: SelectAllRecipes :many
SELECT r.id, r.name, r.description, b.id, b.name, b.description, u.id, u.username, u.name, u.email, u.password
FROM recipes r
         JOIN books b ON r.book_id = b.id
         JOIN users u ON b.owner_id = u.id;

-- name: InsertRecipe :one
INSERT INTO recipes (id, name, description, book_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateRecipe :one
UPDATE recipes
SET name = $2, description = $3, book_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteRecipe :exec
DELETE FROM recipes
WHERE id = $1;

-- name: SelectIngredientsByRecipeID :many
SELECT id, name, description, recipe_id, qty, unit
FROM ingredients
WHERE recipe_id = $1;

-- name: InsertIngredient :one
INSERT INTO ingredients (id, name, description, recipe_id, qty, unit)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateIngredient :one
UPDATE ingredients
SET name = $2, description = $3, recipe_id = $4, qty = $5, unit = $5
WHERE id = $1
RETURNING *;

-- name: DeleteIngredient :exec
DELETE FROM ingredients
WHERE id = $1;

-- name: SelectStepsByRecipeID :many
SELECT id, name, description, recipe_id, duration
FROM steps
WHERE recipe_id = $1;

-- name: InsertStep :one
INSERT INTO steps (id, name, description, recipe_id, duration)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateStep :one
UPDATE steps
SET name = $2, description = $3, recipe_id = $4, duration = $5
WHERE id = $1
RETURNING *;

-- name: DeleteStep :exec
DELETE FROM steps
WHERE id = $1;
