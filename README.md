# go-api-starter-kit

## Curl actions

``` bash
# list
curl -X GET http://localhost:8080/api/v2/examples

# create
curl -X POST http://localhost:8080/api/v2/examples -H 'Content-Type: application/json' -d '{"name":"test"}'

# show
curl -X GET http://localhost:8080/api/v2/examples/4

# update
curl -X PUT http://localhost:8080/api/v2/examples/4 -H 'Content-Type: application/json' -d '{"name":"test3"}'

# delete
curl -X DELETE http://localhost:8080/api/v2/examples/6
```

## Rerurns and statuses

In case everything is alright

``` txt
message = OK, HTTP status = 200
```

In case of an issue:

``` txt
message = dbError, HTTP status = 400 (database query syntax error)
message = apiIssue, HTTP status = 400 (misc error like parsing issue, transaction start up issue, basically issues we didnt predicted)
message = notFound, HTTP status = 404 (query is searching for something that cannot be found in the DB)
```
