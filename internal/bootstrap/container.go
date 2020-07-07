package bootstrap

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
)

type Container struct {
	Config                      *Config
	Logger                      *logrus.Logger
	StatisticoDataServiceConnection *grpc.ClientConn
}

func BuildContainer(config *Config) *Container {
	c := Container{Config: config}

	c.Logger = logger()
	c.StatisticoDataServiceConnection = statisticoDataServiceConnection(config)

	return &c
}

func logger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	return logger
}

func statisticoDataServiceConnection(config *Config) *grpc.ClientConn {
	address := config.StatisticoDataService.Host + ":" + config.StatisticoDataService.Port

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		logger().Warnf("Error initializing statistico data service grpc client %s", err.Error())
	}

	return conn
}
