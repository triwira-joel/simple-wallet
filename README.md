## TODO
- Setup database
- Setup DB schema and seeds
- List down + Implement all endpoints
- UI/FE Design

## Setup MySQL
Assuming docker and mysql already installed

``` docker run --name simple-wallet -e MYSQL_ROOT_PASSWORD=12345678 -p 3333:3306 -d mysql:latest ```

Run mysql

```mysql --user=root --password``` 

Create database `simple_wallet`

``` CREATE DATABASE simple_wallet ```

## Migration
Install golang-migrate

```go install -tags 'simple_wallet' github.com/golang-migrate/migrate/v4/cmd/migrate@latest ```

Create a migration file

``` migrate create -ext sql -dir sqldb/migrations <file_name> ```

Run the migration up

``` migrate -database "mysql://root:12345678@tcp(127.0.0.1:3333)/simple_wallet?parseTime=true" -path sqldb/migrations up ```

## Add Query
To add query, first install sqlc

``` go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest ```

Update schema.sql and query.sql when you have updated database. Then run this to generate the code 

``` sqlc generate ```
