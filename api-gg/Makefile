build:
	@go build -o bin/gobank

run: build
	@./bin/gobank

db:
	docker run --name gobank -p 5434:5432 -e POSTGRES_PASSWORD=12345 -e POSTGRES_DB=gobank -d postgres

test:
	@go test -v ./...