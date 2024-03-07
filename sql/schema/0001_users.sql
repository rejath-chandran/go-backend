-- +goose Up
CREATE TABLE users(
    id  INT PRIMARY KEY,
    email TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    UNIQUE(email)
);

-- +goose Down
DROP TABLE users;