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
touch main.go app/sample.go app/sample_test.go \
 app/config/database.go app/model/contact.go app/model/address.go \
 app/service/agenda-service.go
```

## How to run

```bash
go build

```

## How to test

## Noteworthy

[gorm]: https://gorm.io/
