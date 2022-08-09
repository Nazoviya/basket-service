include app.env

postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root basketService

dropdb:
	docker exec -it postgres14 dropdb basketService

migrateup:
	migrate -path db/migration -database ${DB_SOURCE} -verbose up

migratedown:
	migrate -path db/migration -database ${DB_SOURCE} -verbose down

sqlc:
	docker run --rm -v "${CURDIR}:/src" -w /src kjconroy/sqlc generate

server:
	go run main.go

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server test