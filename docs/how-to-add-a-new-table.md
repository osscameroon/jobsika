1. Open [`migrations/20211019221200_init_db.sql`](https://github.com/osscameroon/camerdevs/blob/main/backend/migrations/20211019221200_init_db.sql)
2. Add your new table schema in the file, between the `-- +goose Up \n -- +goose StatementBegin` and `-- +goose StatementEnd` lines
Example:
```sql
CREATE TABLE IF NOT EXISTS <your_new_table> (
  id BIGSERIAL NOT NULL,
/*
	... Some fields
*/
  createdat DATE,
  updatedat DATE
);
```

3. Add a drop table statement `-- +goose Down \n -- +goose StatementBegin` and `-- +goose StatementEnd` lines
Example:
```sql
DROP TABLE IF EXISTS <your_new_table>;
```

4. Create a fixture file `<your_new_table>.sql` for that will be used to populate your table in this [location](https://github.com/osscameroon/camerdevs/tree/main/backend/fixtures/postgres/sql).
 Note: You can use [mockaroo](https://www.mockaroo.com/) to generate the fixture data.

5. Add the `../../fixtures/postgres/sql/<your_new_table>.sql` in the appropriate position on that [`sql_file_ordered.txt`](https://github.com/osscameroon/camerdevs/blob/main/backend/scripts/setup-postgres/sql_file_ordered.txt) file. This  is used to specify the order in which your sql fixtures files will be applied to the Database.
