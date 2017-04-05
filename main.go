package main

import (
	"github.com/abaeve/services-common/config"
	"github.com/abaeve/auth-srv/handler"
	"github.com/abaeve/auth-srv/proto"
	"github.com/abaeve/auth-srv/repository"
	_ "github.com/go-sql-driver/mysql"
)

var version string = "1.0.0"

func main() {
	configuration := config.Configuration{}
	// These needs to be a commandline argument eventually
	configuration.Load("application.yaml")
	service, serr := configuration.NewService(version)
	if serr != nil {
		panic(serr)
	}
	connectionString, cerr := configuration.NewConnectionString()
	if cerr != nil {
		panic(cerr)
	}

	err := repository.Setup(configuration.Database.Driver, connectionString)

	if err != nil {
		panic("Could not open database connection using: " + connectionString + " received error: " + err.Error())
	}
	repository.DB.DB().Ping()
	repository.DB.DB().SetMaxOpenConns(configuration.Database.MaxConnections)
	defer repository.DB.Close()

	service.Init()
	abaeve_auth.RegisterUserAuthenticationAdminHandler(service.Server(), &handler.AdminHandler{service.Client()})
	abaeve_auth.RegisterUserAuthenticationHandler(service.Server(), &handler.AuthHandler{service.Client()})
	service.Run()
}
