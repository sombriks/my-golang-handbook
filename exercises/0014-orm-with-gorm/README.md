# Object/Relational mapping with [GORM][gorm] 

Using a framework to make database access easier 

## Project setup

```bash
mkdir 0014-orm-with-gorm
cd 0014-orm-with-gorm
go mod init 0014-orm-with-gorm
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
mkdir -p app/{config,model,service}
touch README.md main.go app/sample.go app/sample_test.go \
 app/config/database.go app/model/contact.go app/model/address.go \
 app/service/agenda-service.go
```

## How to run

```bash
go build
./0014-orm-with-gorm
```

## How to test

```bash
go test -v -coverprofile=coverage.out ./app
```

If you want a better coverage report (line above just prints total coverage %)
add the following command:

```bash
go tool cover -html=coverage.out -o coverage.html
```

## Noteworthy

- gorm has a very opinionated way to deal with databases. tables must be
  [pluralized][gorm-plural], timestamps must exists and primary key is `id`.
- associations, when desired, must be explicitly [preloaded][gorm-preload] on 
  queries, but will be saved in a transparent way. 
- golang has a [rudimentary setup/teardown system for testcases][go-test]. But
  it does the job and we can gracefully prepare and close the database.

[gorm]: https://gorm.io/
[gorm-plural]: https://gorm.io/docs/models.html#Conventions
[gorm-preload]: https://gorm.io/docs/preload.html
[go-test]: https://pkg.go.dev/testing#hdr-Main
