package main

import (
	"errors"
	"github.com/abaeve/auth-srv/handler"
	"github.com/abaeve/auth-srv/proto"
	"github.com/abaeve/auth-srv/repository"
	"github.com/abaeve/services-common/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro"
)

var version string = "1.0.0"

func main() {
	service := config.NewService(version, "auth-srv", initialize)

	abaeve_auth.RegisterUserAuthenticationAdminHandler(service.Server(), &handler.AdminHandler{service.Client()})
	abaeve_auth.RegisterUserAuthenticationHandler(service.Server(), &handler.AuthHandler{service.Client()})

	service.Init(micro.BeforeStop(func() error {
		return repository.DB.Close()
	}))

	service.Run()
}

func initialize(configuration *config.Configuration) error {
	connectionString, err := configuration.NewConnectionString()
	if err != nil {
		return err
	}

	err = repository.Setup(configuration.Database.Driver, connectionString)

	if err != nil {
		return errors.New("Could not open database connection using: " + connectionString + " received error: " + err.Error())
	}
	repository.DB.DB().Ping()
	repository.DB.DB().SetMaxOpenConns(configuration.Database.MaxConnections)

	return nil
}
