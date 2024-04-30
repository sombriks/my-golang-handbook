# REST/JSON API with [echo][echo]

Using a framework to make api development and test easier

## Requirements

- go 1.22
- echo v4
- goqu v9
- docker 25 (for postgres)

## How to run

Be sure to have docker compose running the test postgres database:

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

```bash
```

## Noteworthy


[echo]: https://echo.labstack.com/docs/quick-start
