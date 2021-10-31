package grpc_client

import (
	"fmt"

	pbp "github.com/rabdavinci/message-broker/broker_service/genproto/producer_service"
	pbu "github.com/rabdavinci/message-broker/broker_service/genproto/user_service"
	"google.golang.org/grpc"

	"github.com/rabdavinci/message-broker/broker_service/config"
)

type ServiceManager interface {
	UserService() pbu.UserServiceClient
	ProducerService() pbp.ProducerServiceClient
}

type GrpcClients struct {
	userService     pbu.UserServiceClient
	producerService pbp.ProducerServiceClient
}

func NewGrpcClients(conf *config.Config) (*GrpcClients, error) {
	connUserService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connProducerService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.ProducerServiceHost, conf.ProducerServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		userService:     pbu.NewUserServiceClient(connUserService),
		producerService: pbp.NewProducerServiceClient(connProducerService),
	}, nil
}

func (g *GrpcClients) UserService() pbu.UserServiceClient {
	return g.userService
}

func (g *GrpcClients) ProducerService() pbp.ProducerServiceClient {
	return g.producerService
}
