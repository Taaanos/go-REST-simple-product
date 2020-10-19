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

You need to have `golang-migrate` installed or use a tool of your choice.

```bash
brew install golang-migrate
```

```bash
# cd to the project dir
docker-compose up --build

# in a new window, in the same dir run
make migrateup
```

### How to run tests

You need to have `go` and `migrate` installed.

```bash
make postgres createdb migrateup
make build test
```

### Config and env files

`.config` is read by the go app, change vars according to your needs.
`.env` is read by docker, changes here should be synced with `.config`.
