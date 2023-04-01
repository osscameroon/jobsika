-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS job_offers (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    company_name VARCHAR(255),
    company_email VARCHAR(255),
    title_id INTEGER NOT NULL,
    is_remote BOOLEAN,
    city VARCHAR(255),
    country VARCHAR(255),
    department VARCHAR(255),
    salary_range_min BIGSERIAL,
    salary_range_max BIGSERIAL,
    description VARCHAR(5000),
    benefits VARCHAR(1000),
    how_to_apply VARCHAR(1000),
    application_url VARCHAR(255),
    application_email_address VARCHAR(255),
    application_phone_number VARCHAR(255),
    tags VARCHAR(255),

    createdat DATE,
    updatedat DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS job_offers;
-- +goose StatementEnd
