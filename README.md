# Overview
This is an account transaction REST api implemented with echo framework.
Accounts and transactions are store in a postgres database, and the clients are able to query and regiter new accounts.

There is also an endpoint where the client can perform transactions to an account, this endpoint has a
cache configured that locks each transaction with a small TTL in order to avoid duplications.


## What's inside:

- Create new accounts and get by ID
- Create new transactions and get by ID
- Seed default Operation Types
- Postgres integration
- Redis integration
- Request validation
- Swagger docs
- Environment configuration
- Docker development environment

## Usage
1. Run your application using the command in the terminal:
    `docker-compose up` or `make run-docker` 
2. Browse to localhost:7788/swagger/index.html. You will see Swagger 2.0 API documents.
3. Using the API documentation, make requests to register an account and registering new transactions.
4. These are the default operation types that can be used to register the new transactions
```
{
   "ef4dc378-e57e-4951-ad43-77b8d4af403d": {
      "description": "COMPRA A VISTA",
      "is_debit": "true",
   },
   "443a4215-80db-4614-888c-dc9be9b29656": {
      "description": "COMPRA PARCELADA",
      "is_debit": "true",
   },
   "6f025e29-937f-4ca1-af4e-4fa03838f27e": {
      "description": "SAQUE",
      "is_debit": "true",
   },
   "fce2fa7e-a698-40c8-a765-268d13190329": {
      "description": "PAGAMENTO",
      "is_debit": "false",
   },
}
```
5. To execute the tests run the command  `make test` in the terminal.

## Endpoints

### Create Account
`curl -X 'POST' \ 'http://localhost:7788/accounts' \ -H 'accept: application/json' \ -H 'Content-Type: application/json' \ -d '{ "document_number": "23dddf", "name": "John Doe" }'`

### Get Account By Id
`curl -X 'GET' \ 'http://localhost:7788/accounts/34ccb16d-5854-4d3f-9712-6eb4820d68ea' \ -H 'accept: application/json'`

### Create Transaction

### Get Transaction By Id

## License
The project is developed by [lschmittalves@gmail.com]() under [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)