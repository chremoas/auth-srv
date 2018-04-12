package main

import (
	"errors"
	"github.com/chremoas/auth-srv/handler"
	"github.com/chremoas/auth-srv/proto"
	"github.com/chremoas/auth-srv/repository"
	"github.com/chremoas/services-common/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro"
	"go.uber.org/zap"
)

var version = "1.0.0"
var name = "auth"
var logger *zap.Logger

func main() {
	service := config.NewService(version, "srv", name, initialize)

	abaeve_auth.RegisterUserAuthenticationHandler(service.Server(), &handler.AuthHandler{service.Client(), logger})
	abaeve_auth.RegisterEntityQueryHandler(service.Server(), &handler.EntityQueryHandler{service.Client(), logger})
	abaeve_auth.RegisterEntityAdminHandler(service.Server(), &handler.EntityAdminHandler{service.Client(), logger})

	service.Init(micro.BeforeStop(func() error {
		return repository.DB.Close()
	}))

	service.Run()
}

func initialize(configuration *config.Configuration) error {
	var err error

	// TODO pick stuff up from the config
	logger, err = zap.NewProduction()
	if err != nil {
		return err
	}
	defer logger.Sync()
	logger.Info("Initialized logger")

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
