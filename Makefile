get-docs:
	go get -u github.com/swaggo/swag/cmd/swag

docs: get-docs
	swag init -g cmd/api/main.go --output docs

build: docs
	go mod tidy
	go build -o bin/account-transaction-api cmd/api/main.go

test:
	go test -v ./test/...

build-docker: build
	docker build . -t account-transaction-api

run-docker: build-docker
	docker-compose up

stop-docker:
	docker-compose down