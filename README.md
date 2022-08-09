# Basket Service

Basket service is a REST API to handle CRUD operations in an easy and fast manner 
with relational database PostgreSQL operationally.

## Prerequisites

[Docker](https://www.docker.com/), [Golang](https://go.dev/), [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate), [Sqlc](https://github.com/kyleconroy/sqlc#installation)

## Makefile commands

- Create postgres14 container
    ```console
    make postgres
    ```
- Create PostgreSQL database
    ```console
    make createdb
    ```
- Drop PostgreSQL database
    ```console
    make dropdb
    ```
- Run database migration up
    ```console
    make migrateup
    ```
- Run database migration down
    ```console
    make migratedown
    ```
- Generate CRUDs operations with sqlc 
    ```console
    make sqlc
    ```
- Run server 
    ```console
    make server
    ```
- Run test 
    ```console
    make test
    ```

## Usage

- List all products
    ```console
    http://localhost:8080/products
    ```
- List all products in basket
    ```console
    http://localhost:8080/basket
    ```
- Add specified product to basket
    ```console
    http://localhost:8080/basket/product_id
    ```
- Delete specified product from basket
    ```console
    http://localhost:8080/basket/del/product_id
    ```
- Complete order
    ```console
    http://localhost:8080/complete
    ```