postgres: 
	docker run --name postgres13 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root products

dropdb:
	docker exec -it postgres13 dropdb products

migrateup: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/products?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/products?sslmode=disable" -verbose down

reset:
	make dropdb createdb migrateup

build:
	go build -v -o go-api

run :
	./go-api

test:
	go test -v

.PHONY: postgres createdb dropdb migrateup migratedown reset

