POSTGRES_CONT_NAME = go-rest-simple-product_database_1
postgres: 
	docker run --name $(POSTGRES_CONT_NAME) -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:13-alpine

createdb:
	docker exec -it $(POSTGRES_CONT_NAME) createdb --username=root --owner=root products

dropdb:
	docker exec -it $(POSTGRES_CONT_NAME) dropdb products

migrateup: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/products?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/products?sslmode=disable" -verbose down

reset:
	make dropdb createdb migrateup

test:
	make postgres 
	sleep 3 
	make createdb migrateup
	go test -v

stop: 
	docker-compose down
	docker stop $(POSTGRES_CONT_NAME)
	docker rm $(POSTGRES_CONT_NAME)
	
run:
	docker-compose up --build
	make migrateup

.PHONY: postgres createdb dropdb migrateup migratedown reset run test timeout
