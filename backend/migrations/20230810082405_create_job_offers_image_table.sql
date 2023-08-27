-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS job_offers_image (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    job_offer_id BIGINT NOT NULL,
    image_location VARCHAR(255) NOT NULL,
    createdat TIMESTAMP WITH TIME ZONE,
    updatedat TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_job_offers 
     FOREIGN KEY (job_offer_id) 
     REFERENCES job_offers(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS job_offers_image;
-- +goose StatementEnd