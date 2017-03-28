package main

import (
	"fmt"
	"github.com/abaeve/auth-srv/repository"
	"github.com/micro/go-micro"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/abaeve/auth-srv/proto"
	"github.com/abaeve/auth-srv/handler"
)

type Configuration struct {
	Application struct {
		Namespace string
		Name      string
		Database  struct {
			Driver         string
			Protocol       string
			Host           string
			Port           uint
			Database       string
			Username       string
			Password       string
			Options        string
			MaxConnections int `yaml:"maxConnections"`
		}
	}
}

var version string = "1.0.0"

func main() {
	configuration := Configuration{}

	data, err := ioutil.ReadFile("application.yaml")

	//<editor-fold desc="Configuration Launch Sanity check">
	//TODO: Candidate for shared function for all my services.
	if err != nil {
		panic("Could not read application.yaml for configuration data.")
	}

	err = yaml.Unmarshal([]byte(data), &configuration)

	if err != nil {
		panic("Could not unmarshall application.yaml as yaml")
	}
	//</editor-fold>

	//TODO: Candidate for shared function for all my services
	service := micro.NewService(
		micro.Name(configuration.Application.Namespace+"."+configuration.Application.Name),
		micro.Version(version),
	)

	//<editor-fold desc="DB Initialization">
	//TODO: Candidate for shared function for all my services.
	connectionString := configuration.Application.Database.Username +
		":" +
		configuration.Application.Database.Password +
		"@" +
		configuration.Application.Database.Protocol +
		"(" +
		configuration.Application.Database.Host +
		":" +
		fmt.Sprintf("%d", configuration.Application.Database.Port) +
		")/" +
		configuration.Application.Database.Database

	err = repository.Setup(configuration.Application.Database.Driver, connectionString)

	if err != nil {
		panic("Could not open database connection using: " + connectionString)
	}
	repository.DB.DB().Ping()
	repository.DB.DB().SetMaxOpenConns(configuration.Application.Database.MaxConnections)
	defer repository.DB.Close()
	//</editor-fold>

	service.Init()
	abaeve_auth.RegisterUserAuthenticationAdminHandler(service.Server(), &handler.AdminHandler{service.Client()})
	abaeve_auth.RegisterUserAuthenticationHandler(service.Server(), &handler.AuthHandler{service.Client()})
	service.Run()
}
