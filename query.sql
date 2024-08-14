-- name: GetUserByID :one
SELECT * FROM users WHERE user_id = $1;

-- name: GetBooksByUserID :many
SELECT * FROM books WHERE user_id = $1;

-- name: GetAllBooks :many
SELECT * FROM books;

-- name: CreateUser :one
INSERT INTO users (username, email, password, profile_picture_url, bio, location)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: CreateBook :one
INSERT INTO books (user_id, title, author, genre, condition, description, image_url)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: CreateSwapRequest :one
INSERT INTO swap_requests (requester_id, book_id, status)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateReview :one
INSERT INTO reviews (reviewed_user_id, reviewer_id, rating, review_text)
VALUES ($1, $2, $3, $4)
RETURNING *;

