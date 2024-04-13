# REST/JSON API with fiber

Using a framework to make api development and test easier

## Requirements

- go 1.20
- [fiber][fiber] v3
- [gorm][gorm] 1.25
- go-sqlite3 1.14 

## Project setup

```bash
go mod init 0015-rest-api
go get github.com/gofiber/fiber/v3
go get github.com/gofiber/template/pug/v2
go get gorm.io/gorm
go get gorm.io/driver/sqlite
mkdir -p app/{configs,todos,templates/todos}
touch .gitignore README.md main.go app/app.go \
  app/{configs/database.go,templates/{index.pug,todos/{detail.pug,list.pug,form.pug}}} \
  app/todos/{TodoController.go,TodoItem.go,TodoService.go}
```

## How to run

```bash
go run main.go
```

or


```bash
go build
./0015-rest-api
```


## How to test

```bash

```

## Noteworthy

[fiber]: https://docs.gofiber.io/
[gorm]: https://gorm.io
