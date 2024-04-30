package app

import (
	"0016-rest-json/app/configs"
	"0016-rest-json/app/controllers"
	"0016-rest-json/app/services"
	"github.com/labstack/echo/v4"
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

	// Database setup
	var err error
	app.dbconfig, err = configs.NewDBConfig()
	if err != nil {
		app.echo.Logger.Fatal(err)
	}

	// business setup
	app.todoService = services.NewTodoService(app.dbconfig)
	app.todoController = controllers.NewTodoController(app.todoService)

	// wiring routes -- could we down that to controllers setup? yes?
	app.echo.GET("/todos", app.todoController.List)
	app.echo.POST("/todos", app.todoController.Insert)
	app.echo.GET("/todos/:id", app.todoController.Find)
	app.echo.PUT("/todos/:id", app.todoController.Update)
	app.echo.DELETE("/todos/:id", app.todoController.Delete)

	return
}

// Run - makes app listens to requests
func (app *App) Run() {
	app.echo.Logger.Fatal(app.echo.Start(":1323"))
}
