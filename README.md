# Test Invoice System 
Test the ESB Invoice System. This application was built using Golang 1.20, MySQL 5.7, and Docker.

## Link Documentation Postman
[Click Link Postman](https://www.postman.com/bookingtogo/workspace/go-invoice-system/collection/31320401-1a01bfbf-e244-487f-a245-d5881cb508ca?action=share&creator=31320401&active-environment=31320401-2f3b4a60-8752-40ae-92a7-403dda46a8dd)


## Installation
Use the package manager docker to install app MySQL 5.7, Redis and RabbitMQ.
```bash
docker-compose up -d --build
```

Create a new MySQL 5.7 database with name `invoice_system`.

## Mac OS / Linux users
Run the `make migrate` command to populate the database tables.
```bash
make migrate
```

Run the `make seed` command to fill data in several tables.
```bash
make seed
```

Run the `make all` command to run the application.
```bash
make all
```

## Windows users
Run the following command to populate the database table.
``` bash
go mod vendor -v
rm -f cmd/migrate/migrate
go build -o cmd/migrate/migrate cmd/migrate/migrate.go
./cmd/migrate/migrate
```

Run the following command to fill data in several tables.
``` bash
go mod vendor -v
rm -f cmd/seed/seed
go build -o cmd/seed/seed cmd/seed/seed.go
./cmd/seed/seed
```
Run the following command to run the application.
``` bash
go mod vendor -v
rm -f cmd/app/app
go build -o cmd/app/app cmd/app/app.go
./cmd/app/app
```