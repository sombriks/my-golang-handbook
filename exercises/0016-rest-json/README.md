# REST/JSON API with [echo][echo]

Using a framework to make api development and test easier

## Requirements

- go 1.22
- echo v4 (api)
- goqu v9 (db/query builder)
- docker 25 (to provision a postgres)

## How to run

Be sure to have docker compose running the development postgres database:

```bash
# cd exercises/0016-rest-json
docker compose -f infrastructure/docker-compose.yml up -d
```

Then build and run like any normal go application:

```bash
go build 
./0016-rest-json
```

## How to test

Perform some requests to the running API:

```bash
# list todos
curl -X GET http://localhost:1323/todos
# find by id
curl -X GET http://localhost:1323/todos/2
# insert todo
curl -X POST --json '{"description":"new task"}' http://localhost:1323/todos
# update todo
curl -X PUT --json '{"description":"task updated"}' http://localhost:1323/todos/2
# delete todo
curl -X DELETE http://localhost:1323/todos/3
```



## Noteworthy

- I have a fond for query builders since [knex][knex] and [goqu][goqu] is a nice
  surprise.
- The echo's [bind api][echo-bind] delivers the best data handling experience in
  go ecosystem so far.

[echo]: https://echo.labstack.com/docs/quick-start
[knex]: https://knexjs.org/
[echo-bind]: https://echo.labstack.com/docs/binding
