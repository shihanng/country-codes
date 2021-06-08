# country-codes

Tools to obtain latest versions of country codes from [ISO's Online Browsing Platform (OBP)](https://www.iso.org/obp/).

## Development note

### Migration

We use [`migration`](https://github.com/golang-migrate/migrate) to manage schema migration of our SQLite3 DB.
The tool needs to be installed with `sqlite3` tag so that it can support SQLite3.

```
go install -tags "sqlite3" github.com/golang-migrate/migrate/v4/cmd/migrate
```

### How to generate migration files

Migration files can be found in `db/migrations/`.

```
cd db/migrations/
migrate create -ext sql <FILENAME BASE>
```

Migrations are applied by the app internally, so we do not need to run `migrate up`.

### How to rollback a migration

Example of rolling back one step:

```
migrate -source file://./db/migrations -database "sqlite3://./country_code.db" down 1
```
