package bootstrap

import (
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-web-gateway/internal/app/grpc/proto"
	"google.golang.org/grpc"
	"os"
)

type Container struct {
	Config                      *Config
	Logger                      *logrus.Logger
	StatisticoDataServiceClient proto.TeamServiceClient
}

func BuildContainer(config *Config) *Container {
	c := Container{Config: config}

	c.Logger = logger()
	c.StatisticoDataServiceClient = statisticoDataServiceClient(config)

	return &c
}

func logger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	return logger
}

func statisticoDataServiceClient(config *Config) proto.TeamServiceClient {
	address := config.StatisticoDataService.Host + ":" + config.StatisticoDataService.Port

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		logger().Warnf("Error initializing statistico data service grpc client %s", err.Error())
	}

	return proto.NewTeamServiceClient(conn)
}
