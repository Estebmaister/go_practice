migrate-up:
	go run cli.go migrate up

migrate-down:
	go run cli.go migrate down

migrate-test-up:
	GO_ENVIRONMENT=test go run cli.go migrate up

migrate-test-down:
	GO_ENVIRONMENT=test go run cli.go migrate down

install:
	go get ./...

prepare: install migrate-up

start:
	go run app/main.go

test:
	go test ./app/tests/...

lint:
	golangci-lint run ./...
