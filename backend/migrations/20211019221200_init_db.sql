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

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS salaries;
-- +goose StatementEnd
