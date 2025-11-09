-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sale (
    id SERIAL PRIMARY KEY,
    item VARCHAR(50) NOT NULL,
    income NUMERIC CHECK (income >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sale;
-- +goose StatementEnd
