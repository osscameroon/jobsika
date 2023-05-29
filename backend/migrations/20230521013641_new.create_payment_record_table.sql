-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payment_records (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    job_offer_id BIGINT NOT NULL,
    tier_id VARCHAR(255) NOT NULL,
	legacy_id INTEGER NOT NULL,
	slug VARCHAR(255) NOT NULL,
	tier_url VARCHAR(255) NOT NULL,
    createdat DATE,
    updatedat DATE,
    CONSTRAINT fk_job_offers
      FOREIGN KEY(job_offer_id)
	  REFERENCES job_offers(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payment_records;
-- +goose StatementEnd
