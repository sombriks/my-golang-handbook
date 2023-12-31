# Databases

Sampling go capabilities of talk with databases

Module created locally with `go mod init 0013-databases`

## what kind of database we're sampling here

- Relational SQL databases
  - <https://github.com/mattn/go-sqlite3> 
- NoSQL, document databases
  - <https://github.com/ostafen/clover> 
- Key/Value databases
  - <https://github.com/syndtr/goleveldb> 

## Building

```bash
go build -o example0013 .
```

## Running tests and building coverage report

```bash
go test -v -coverprofile=coverage.out ./relational ./key-value-based ./document-based
go tool cover -html=coverage.out -o coverage.html
```
