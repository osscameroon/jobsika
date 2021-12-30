-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS salaries (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  city VARCHAR(255) NOT NULL,
  company VARCHAR(255) NOT NULL,
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

CREATE TABLE IF NOT EXISTS company_ratings (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  company_id INTEGER NOT NULL,
  rating INTEGER NOT NULL,
  comment VARCHAR(5000),
  createdat DATE,
  updatedat DATE
);


-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS salaries;
DROP TABLE IF EXISTS companies;
DROP TABLE IF EXISTS company_ratings;
-- +goose StatementEnd
