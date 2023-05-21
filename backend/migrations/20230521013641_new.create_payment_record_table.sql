-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payment_records (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    tier_id VARCHAR(255) NOT NULL,
	legacy_id INTEGER NOT NULL,
	slug VARCHAR(255) NOT NULL,
	tier_url VARCHAR(255) NOT NULL,
    createdat DATE,
    updatedat DATE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payment_records;
-- +goose StatementEnd
