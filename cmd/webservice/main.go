package main

import (
	"todo/webservice/controllers"
	"todo/webservice/internal/entities"
	"todo/webservice/services/config"
	"todo/webservice/services/log"
	"todo/webservice/services/server"

	"gorm.io/gorm"
)

type Webservice struct {
	App            *server.GoFiber
	DB             *gorm.DB
	ENVConfig      *config.ENVConfig
	Log            *log.STDLogger
	TodoController *controllers.TodoController
}

func NewWebService(app *server.GoFiber, envConfig *config.ENVConfig, log *log.STDLogger, db *gorm.DB, todoController *controllers.TodoController) *Webservice {
	return &Webservice{
		App:            app,
		DB:             db,
		Log:            log,
		ENVConfig:      envConfig,
		TodoController: todoController,
	}
}

func main() {
	webservice := InitWebservice()
	webservice.Log.Info("initiate web service finish")

	webservice.DB.AutoMigrate(
		&entities.Todo{},
	)
	webservice.Log.Info("migrate database finish")

	webservice.App.Get("/todo/:id", webservice.TodoController.GetByID)
	webservice.App.Post("/todo", webservice.TodoController.Create)
	webservice.Log.Info("register all paths finish")

	webservice.App.Listen(webservice.ENVConfig.Port)
}
