// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"todo/webservice/controllers"
	"todo/webservice/internal/entities"
	"todo/webservice/internal/usecases"
	"todo/webservice/services/config"
	"todo/webservice/services/db"
	"todo/webservice/services/log"
	"todo/webservice/services/server"
)

// Injectors from wire.go:

func InitWebservice() *Webservice {
	goFiber := server.NewGoFiber()
	envConfig := config.NewENVConfig()
	stdLogger := log.NewSTDLogger(envConfig)
	gormDB := db.NewDB(envConfig, stdLogger)
	todoRepository := entities.NewTodoRepository(gormDB)
	todoUsecase := usecases.NewTodoUsecase(todoRepository)
	todoController := controllers.NewTodoController(todoUsecase)
	webservice := NewWebService(goFiber, envConfig, stdLogger, gormDB, todoController)
	return webservice
}

// wire.go:

var (
	serviceProviders    = wire.NewSet(server.NewGoFiber, config.NewENVConfig, log.NewSTDLogger, db.NewDB)
	controllerProviders = wire.NewSet(controllers.NewTodoController)
	usecaseProviders    = wire.NewSet(usecases.NewTodoUsecase)
	entityProviders     = wire.NewSet(entities.NewTodoRepository)
)
