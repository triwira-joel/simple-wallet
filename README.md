## TODO
- Setup database
- Setup DB schema and seeds
- List down + Implement all endpoints
- UI/FE Design

## Setup MySQL
Assuming docker and mysql already installed

``` docker run --name simple-wallet -e MYSQL_ROOT_PASSWORD=12345678 -p 3333:3306 -d mysql:latest ```

to run mysql

```mysql --user=root --password``` 

Create database `simple_wallet`

``` CREATE DATABASE simple_wallet ```

## Migration
```go install -tags 'simple_wallet' github.com/golang-migrate/migrate/v4/cmd/migrate@latest ```

``` migrate -database "mysql://root:12345678@tcp(127.0.0.1:3333)/simple_wallet?parseTime=true" -path sqldb/migrations up ```