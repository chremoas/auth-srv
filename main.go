package main

import (
	"fmt"
	"github.com/abaeve/auth-common/config"
	"github.com/abaeve/auth-srv/repository"
	"github.com/micro/go-micro"
	"github.com/abaeve/auth-srv/proto"
	"github.com/abaeve/auth-srv/handler"
	_ "github.com/go-sql-driver/mysql"
)

var version string = "1.0.0"

func main() {
	configuration := config.Configuration{}
	// These needs to be a commandline argument eventually
	configuration.Load("application.yaml")

	//TODO: Candidate for shared function for all my services
	service := micro.NewService(
		micro.Name(configuration.Namespace+"."+configuration.Name),
		micro.Version(version),
	)

	//<editor-fold desc="DB Initialization">
	//TODO: Candidate for shared function for all my services.
	connectionString := configuration.Database.Username +
		":" +
		configuration.Database.Password +
		"@" +
		configuration.Database.Protocol +
		"(" +
		configuration.Database.Host +
		":" +
		fmt.Sprintf("%d", configuration.Database.Port) +
		")/" +
		configuration.Database.Database

	err := repository.Setup(configuration.Database.Driver, connectionString)

	if err != nil {
		panic("Could not open database connection using: " + connectionString + " received error: " + err.Error())
	}
	repository.DB.DB().Ping()
	repository.DB.DB().SetMaxOpenConns(configuration.Database.MaxConnections)
	defer repository.DB.Close()
	//</editor-fold>

	service.Init()
	abaeve_auth.RegisterUserAuthenticationAdminHandler(service.Server(), &handler.AdminHandler{service.Client()})
	abaeve_auth.RegisterUserAuthenticationHandler(service.Server(), &handler.AuthHandler{service.Client()})
	service.Run()
}
