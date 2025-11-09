-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_sale_created_at ON sale (created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_sale_created_at;
-- +goose StatementEnd
