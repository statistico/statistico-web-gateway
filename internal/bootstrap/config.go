package bootstrap

import "os"

type Config struct {
	StatisticoDataService
}

type StatisticoDataService struct {
	Host string
	Port string
}

func BuildConfig() *Config {
	config := Config{}

	config.StatisticoDataService = StatisticoDataService{
		Host: os.Getenv("STATISTICO_DATA_SERVICE_HOST"),
		Port: os.Getenv("STATISTICO_DATA_SERVICE_PORT"),
	}

	return &config
}
