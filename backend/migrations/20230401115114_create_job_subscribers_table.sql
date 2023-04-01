-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS job_subscribers (
    id BIGSERIAL NOT NULL,
    email VARCHAR(255) NOT NULL PRIMARY KEY,
    createdat DATE,
    updatedat DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS job_subscribers;
-- +goose StatementEnd
