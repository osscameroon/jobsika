-- +goose Up
-- +goose StatementBegin
ALTER TABLE job_offers
ADD COLUMN company_image_location VARCHAR(255);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd