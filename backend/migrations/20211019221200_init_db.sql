-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS salaries (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  title_id INTEGER NOT NULL,
  city_id INTEGER NOT NULL,
  country VARCHAR(255),
  company_id INTEGER NOT NULL,
  company_rating_id INTEGER,
  seniority VARCHAR(255) NOT NULL,
  salary INTEGER NOT NULL,
  createdat DATE,
  updatedat DATE
);

CREATE TABLE IF NOT EXISTS companies (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  rating INTEGER NOT NULL,
  createdat DATE,
  updatedat DATE
);

CREATE TABLE IF NOT EXISTS jobtitles (
  id BIGSERIAL NOT NULL,
  title VARCHAR(255) NOT NULL PRIMARY KEY,
  createdat DATE,
  updatedat DATE
);

CREATE TABLE IF NOT EXISTS company_ratings (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  company_id INTEGER NOT NULL,
  rating INTEGER NOT NULL,
  comment VARCHAR(5000),
  createdat DATE,
  updatedat DATE
);

CREATE TABLE IF NOT EXISTS cities (
    id BIGSERIAL NOT NULL,
    name VARCHAR(255) NOT NULL PRIMARY KEY,
    createdat DATE,
    updatedat DATE
);

CREATE TABLE IF NOT EXISTS job_offers (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    email VARCHAR(255),
    company_name VARCHAR(255),
    title_id INTEGER NOT NULL,
    is_remote BOOLEAN,
    description VARCHAR(5000),
    how_to_apply VARCHAR(1000),
    apply_url VARCHAR(255),
    apply_email_address VARCHAR(255),
    createdat DATE,
    updatedat DATE
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS salaries;
DROP TABLE IF EXISTS companies;
DROP TABLE IF EXISTS jobtitles;
DROP TABLE IF EXISTS cities;
DROP TABLE IF EXISTS company_ratings;
DROP TABLE IF EXISTS job_offers;
-- +goose StatementEnd
