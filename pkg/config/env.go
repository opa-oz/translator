package config

import (
	"log"

	"github.com/Netflix/go-env"
)

type Environment struct {
	Port       int64  `env:"PORT,default=8080"`
	Production bool   `env:"PRODUCTION,default=false"`
	ApiKey     string `env:"API_KEY"`
	FolderID   string `env:"FOLDER_ID"`

	Extras env.EnvSet
}

func GetConfig() (*Environment, error) {
	var environment Environment

	es, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	environment.Extras = es

	return &environment, nil
}
