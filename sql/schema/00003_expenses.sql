-- +goose Up
CREATE TABLE expenses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    amount NUMERIC(10, 2) NOT NULL CHECK (amount > 0),
    description VARCHAR(255),
    category_name VARCHAR(255) NOT NULL REFERENCES categories(name) ON UPDATE CASCADE, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    spend_date DATE NOT NULL
);

CREATE INDEX idx_expenses_user_id ON expenses(user_id, spend_date);

-- +goose Down
DROP TABLE IF EXISTS expenses;
