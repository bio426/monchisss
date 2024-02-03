package datasource

import "os"

type Configuration struct {
	PROD        bool
	PG_HOST     string
	PG_PORT     string
	PG_USER     string
	PG_PASSWORD string
	PG_DATABASE string
}

var Config = Configuration{}

func InitConfig() Configuration {
	_, isProd := os.LookupEnv("PROD")
	Config.PROD = isProd

	Config.PG_HOST = os.Getenv("PG_HOST")
	Config.PG_PORT = os.Getenv("PG_PORT")
	Config.PG_USER = os.Getenv("PG_USER")
	Config.PG_PASSWORD = os.Getenv("PG_PASSWORD")
	Config.PG_DATABASE = os.Getenv("PG_DATABASE")

	return Config
}
