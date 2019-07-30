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
var service micro.Service
var logger *zap.Logger
var name = "auth"

func main() {
	var err error
	service = config.NewService(version, "srv", name, initialize)

	// TODO pick stuff up from the config
	logger, err = zap.NewProduction()
	sugar := logger.Sugar()
	if err != nil {
		panic(err)
	}
	defer sugar.Sync()
	sugar.Info("Initialized logger")

	abaeve_auth.RegisterEntityQueryHandler(service.Server(), &handler.EntityQueryHandler{service.Client(), logger})
	abaeve_auth.RegisterEntityAdminHandler(service.Server(), &handler.EntityAdminHandler{service.Client(), logger})

	service.Init(micro.BeforeStop(func() error {
		return repository.DB.Close()
	}))

	err = service.Run()
	if err != nil {
		sugar.Errorf("Running service failed: %s", err.Error())
	}
}

func initialize(configuration *config.Configuration) error {
	abaeve_auth.RegisterUserAuthenticationHandler(service.Server(), handler.NewAuthHandler(configuration, service, logger))
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
