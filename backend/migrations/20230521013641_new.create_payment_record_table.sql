-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payment_record (
    id BIGSERIAL NOT NULL,
    email VARCHAR(255) NOT NULL PRIMARY KEY,
    tier_id INTEGER NOT NULL,
	legacy_id INTEGER NOT NULL,
	slug VARCHAR(255) NOT NULL,
	tier_url VARCHAR(255) NOT NULL,
    createdat DATE,
    updatedat DATE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payment_record;
-- +goose StatementEnd
