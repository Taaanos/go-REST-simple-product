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
* Routing: gorilla/mux

## How to run it

The app currently assumes a working installation of postgres. Create a database in your postgres instance and provide your credentials and info in the .config file.

## API Spec

[Postman Collection](https://documenter.getpostman.com/view/13097698/TVRoYmdb)

* GET /products returns a list of products
* GET /products?start= returns a list of products from the start point in the db
* POST /product creates a new product
* PUT /product/:id updates the given product
* DELETE /product/:id deletes the given product

example JSON request

```json
{"id":1,"name":"speaker","price":101}
```
