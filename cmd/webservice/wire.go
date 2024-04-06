//go:build wireinject
// +build wireinject

package main

import (
	"todo/webservice/controllers"
	"todo/webservice/internal/entities"
	"todo/webservice/internal/usecases"
	"todo/webservice/services/config"
	"todo/webservice/services/db"
	"todo/webservice/services/log"
	"todo/webservice/services/server"

	"github.com/google/wire"
)

var (
	serviceProviders    = wire.NewSet(server.NewGoFiber, config.NewENVConfig, log.NewSTDLogger, db.NewDB)
	controllerProviders = wire.NewSet(controllers.NewTodoController)
	usecaseProviders    = wire.NewSet(usecases.NewTodoUsecase)
	entityProviders     = wire.NewSet(entities.NewTodoRepository)
)

func InitWebservice() *Webservice {
	wire.Build(NewWebService, controllerProviders, serviceProviders, usecaseProviders, entityProviders)
	return &Webservice{}
}
