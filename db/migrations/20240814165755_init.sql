-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    user_id             SERIAL PRIMARY KEY,
    username            VARCHAR(255) UNIQUE NOT NULL,
    email               VARCHAR(255) UNIQUE NOT NULL,
    password            VARCHAR(255) NOT NULL,
    profile_picture_url VARCHAR(255),
    bio                 TEXT,
    location            VARCHAR(255),
    created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);

-- Book Table
CREATE TABLE books (
    book_id      SERIAL PRIMARY KEY,
    user_id      INT NOT NULL,
    title        VARCHAR(255) NOT NULL,
    author       VARCHAR(255) NOT NULL,
    genre        VARCHAR(255) NOT NULL,
    condition    VARCHAR(50) NOT NULL,
    description  TEXT,
    image_url    VARCHAR(255),
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- SwapRequest Table
CREATE TABLE swap_requests (
    swap_request_id SERIAL PRIMARY KEY,
    requester_id    INT NOT NULL,
    book_id         INT NOT NULL,
    status          VARCHAR(50) NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    FOREIGN KEY (requester_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES books(book_id) ON DELETE CASCADE
);

-- Review Table
CREATE TABLE reviews (
    review_id        SERIAL PRIMARY KEY,
    reviewed_user_id INT NOT NULL,
    reviewer_id      INT NOT NULL,
    rating           INT NOT NULL,
    review_text      TEXT,
    created_at       TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    FOREIGN KEY (reviewed_user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (reviewer_id) REFERENCES users(user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users, books, swap_requests, reviews;
-- +goose StatementEnd
