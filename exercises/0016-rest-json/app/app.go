package app

import (
	"0016-rest-json/app/configs"
	"0016-rest-json/app/controllers"
	"0016-rest-json/app/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

type App struct {
	echo           *echo.Echo
	dbconfig       *configs.DBConfig
	todoService    *services.TodoService
	todoController *controllers.TodoController
}

// NewApp - instantiates and configures the app
func NewApp() (app *App) {
	app = &App{echo: echo.New()}
	app.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Printf("%v, %v\n", v.Status, v.URI)
			return nil
		},
	}))

	// Database setup
	var err error
	app.dbconfig, err = configs.NewDBConfig()
	if err != nil {
		app.echo.Logger.Fatal(err)
	}

	// business setup
	app.todoService = services.NewTodoService(app.dbconfig)
	app.todoController = controllers.NewTodoController(app.echo, app.todoService)

	return
}

// Run - makes app listens to requests
func (app *App) Run() {
	app.echo.Logger.Fatal(app.echo.Start(":1323"))
}
