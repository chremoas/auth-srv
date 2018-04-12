package main

import (
	"errors"
	"github.com/chremoas/auth-srv/handler"
	"github.com/chremoas/auth-srv/proto"
	"github.com/chremoas/auth-srv/repository"
	"github.com/chremoas/services-common/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro"
)

var version = "1.0.0"
var name = "auth"

func main() {
	service := config.NewService(version, "srv", name, initialize)

	abaeve_auth.RegisterUserAuthenticationHandler(service.Server(), &handler.AuthHandler{service.Client()})
	abaeve_auth.RegisterEntityQueryHandler(service.Server(), &handler.EntityQueryHandler{service.Client()})
	abaeve_auth.RegisterEntityAdminHandler(service.Server(), &handler.EntityAdminHandler{service.Client()})

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
