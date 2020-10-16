# go-REST-simple-product

A simple REST API to manage products in a database.

## Functionality

* Create a new product
* Update an existing product
* Delete an existing product
* Fetch an existing product
* Fetch a list of products

## Tools

* DB: PostgreSQL
* Routing: [gorilla/mux](https://github.com/gorilla/mux)
* Migrations: [migrate](https://github.com/golang-migrate/migrate)

## How to run it

```
make postgres createdb migrateup
```

## API Spec

[Postman Collection](https://documenter.getpostman.com/view/13097698/TVRoYmdb)
