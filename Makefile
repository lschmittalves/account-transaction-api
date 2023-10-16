get-docs:
	go get -u github.com/swaggo/swag/cmd/swag

docs: get-docs
	swag init -g cmd/api/main.go --output docs

build: docs
	go mod tidy
	go build -o bin/account-transaction-api cmd/api/main.go

test: mocks
	go test ./tests/... -v

build-docker: build
	docker build . -t account-transaction-api

run-docker: build-docker
	docker-compose up

stop-docker:
	docker-compose down

mocks:
	go install go.uber.org/mock/mockgen@latest
	mockgen -source=internal/repositories/account_repository.go -destination=tests/mocks/repositories/account.go
	mockgen -source=internal/repositories/operationType_repository.go -destination=tests/mocks/repositories/operationType.go
	mockgen -source=internal/repositories/transaction_repository.go -destination=tests/mocks/repositories/transaction.go
	mockgen -source=internal/cache/cache.go -destination=tests/mocks/clients/cache.go
	go mod tidy
