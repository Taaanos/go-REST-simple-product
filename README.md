# go-REST-simple-product

A simple REST API to manage products in a database.

## Functionality

* Create a new product
* Update an existing product
* Delete an existing product
* Get an existing product
* Get a list of products

### API Spec

[Postman Collection](https://documenter.getpostman.com/view/13097698/TVRoYmdb)

## Tools

* DB: PostgreSQL
* Routing: [gorilla/mux](https://github.com/gorilla/mux)
* Migrations: [migrate](https://github.com/golang-migrate/migrate)
* CI: circleci

## How to run

```bash
docker-compose up --build
make createdb migrateup
```

You need to have `migrate` installed or use a tool of your choice.

Tests are currently not containirized.

### How to run tests

```bash
make postgres createdb migrateup
make build test
```

### Config and env files

`.config` is read by the go app, change vars according to your needs
`.env` is read by docker, changes here should be synced with `.config`. (not automated yet)
