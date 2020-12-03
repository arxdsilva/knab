.PHONY: postgres migrate run

postgres:
	docker run --rm -ti -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres:13

migrate:
	go run cmd/api/main.go migrate up

run:
	go run cmd/api/main.go