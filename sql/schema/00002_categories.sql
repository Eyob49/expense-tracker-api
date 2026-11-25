-- +goose Up
CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE
);

INSERT INTO categories (name) VALUES
    ('Groceries'),
    ('Leisure'),
    ('Electronics'),
    ('Utilities'),
    ('Clothing'),
    ('Health'),
    ('Transportation'),
    ('Personal Care'),
    ('Entertainment'),
    ('Travel'),
    ('Subscriptions'),
    ('Others');

-- +goose Down
DROP TABLE IF EXISTS categories;
