get-docs:
	go get -u github.com/swaggo/swag/cmd/swag

docs: get-docs
	swag init --dir cmd/api --parseDependency --output docs

build:
	go mod tidy
	go build -o bin/account-transaction-api cmd/api/main.go

test:
	go test -v ./test/...

build-docker: build
	docker build . -t account-transaction-api

run-docker:
	docker compose up