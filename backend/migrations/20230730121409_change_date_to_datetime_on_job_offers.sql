-- +goose Up
-- +goose StatementBegin
ALTER TABLE job_offers
ALTER COLUMN createdat TYPE TIMESTAMP WITH TIME ZONE USING createdat::TIMESTAMP AT TIME ZONE 'UTC';
ALTER TABLE job_offers
ALTER COLUMN updatedat TYPE TIMESTAMP WITH TIME ZONE USING updatedat::TIMESTAMP AT TIME ZONE 'UTC';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd