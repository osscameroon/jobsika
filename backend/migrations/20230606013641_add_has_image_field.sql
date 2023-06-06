-- +goose Up
-- +goose StatementBegin
ALTER TABLE job_offers
ADD COLUMN has_image BOOLEAN;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd