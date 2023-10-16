# Overview
This is an account transaction REST api implemented with echo framework.
Accounts and transactions are store in a postgres database, and the clients are able to query and regiter new accounts.

There is also an endpoint where the client can perform transactions to an account, this endpoint has a
cache configured that locks each transaction with a small TTL in order to avoid duplications.


## What's inside:

- Create new accounts and get by ID
- Create new transactions and get by ID
- Postgres integration
- Redis integration
- Request validation
- Swagger docs
- Environment configuration
- Docker development environment

## Usage
1. Running without docker, execute the command to create the open-api files `make docs` and set the
3. Run your application using the command in the terminal:

    `docker-compose up`
3. Browse to {HOST}:{PORT}/swagger/index.html. You will see Swagger 2.0 API documents.
4. Using the API documentation, make requests to register an account and registering new transactions.

## License
The project is developed by [lschmittalves@gmail.com]() under [MIT LICENSE](https://github.com/nixsolutions/golang-echo-boilerplate/blob/master/LICENSE)