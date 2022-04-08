# go-api-starter-kit

## Basic charasteristics

- JWT authorization
- MySQL or Postgresql database
- Cloudwatch or DynamoDB cloudtrail
- Gin router with zap logs and cors security
- Accessed allowed based on user roles
- Things the API can do is object based

## Database

The starter kit is configured with mysql DB, but migration to postgres should not be a problem. Just modify the utils/sql.go and env files

replace this

``` go
# utils/sql.go
db, err = vssdb.InitMysqlDB(dbAccountUser, dbAccountPass, dbAccountHost, dbAccountPort, dbAccountName)
# .env
VSS_DB_MYSQL_MAX_OPEN_CONNECTIONS=50
VSS_DB_MYSQL_MAX_IDLE_CONNECTIONS=50
```

with

``` go
# utils/sql.go
db, err = vssdb.InitPostgresDB(dbAccountUser, dbAccountPass, dbAccountHost, dbAccountPort, dbAccountName)
# .env
VSS_DB_PSQL_MAX_OPEN_CONNECTIONS=50
VSS_DB_PSQL_MAX_IDLE_CONNECTIONS=50
```

## Migration

Migration is done by using a golang-migrate library. For more details, please have look at <https://github.com/golang-migrate/migrate>

### CLI

Installation

``` bash
# for mysql
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Migration creation:

``` bash
migrate create -ext sql -dir migrations test
```

### Migrate

Migration is run on API start up and runs all migrations from ./migrations directory

## Audit trail

At this moment the starter kit works with Cloudwatch log groups, in the future I will add DynamoDB too.

## Claims

Configuration in middlewares. At this moment i am getting just User ID from the JWT, you might need to modify that to add extra fileds. See moddlewares/auth.go

``` go
func getClaims(c *gin.Context, token *jwt.Token, log *zap.Logger) error {
    // ...
    if ok {
        c.Set("foo1", claims["bar1"])
        c.Set("foo2", claims["bar2"])
        c.Set("foo3", claims["bar3"])
        c.Set("foo4", claims["bar4"])
    }
    // ...
}
```

## Errors

see helpers/errors for more details. At this moment, those errors are returned in message string.

- STATUS_OK_STRING = "OK"
- STATUS_DB_ERROR_STRING = "dbError"
- STATUS_SERVER_ERROR_STRING = "apiError"
- STATUS_NOTFOUND_ERROR_STRING = "notFound"
- STATUS_REQUEST_ERROR_STRING = "requestError"
- STATUS_VALIDATION_ERROR_STRING = "validationError"
- STATUS_AUTH_ERROR_STRING = "notAllowed"

## Objects

I provided an object called example with CRUD actions to create, update, delete, show and list the objects. This example can be copied and renamed to anything. Users, Accounts. See examples' object for more details.

Each object has own routes, validation and methods and then connected to the API.

## Testing with CURL

``` bash
# Start the application
touch .env
cp emv.example .env
# Configure your env file
go run main.go

# list
curl -X GET http://localhost:8080/api/v1/examples

# create
curl -X POST http://localhost:8080/api/v1/examples -H 'Content-Type: application/json' -d '{"name":"test"}'

# show
curl -X GET http://localhost:8080/api/v1/examples/4

# update
curl -X PUT http://localhost:8080/api/v1/examples/4 -H 'Content-Type: application/json' -d '{"name":"test3"}'

# delete
curl -X DELETE http://localhost:8080/api/v1/examples/6
```

## Testing

``` bash
go-api-starter-kit
```
