package grpc_client

import (
	"fmt"

	"github.com/rabdavinci/message-broker/producer_service/config"
	"github.com/rabdavinci/message-broker/producer_service/genproto/broker_service"
	"google.golang.org/grpc"
)

type ServiceManager interface {
	BrokerService() broker_service.BrokerServiceClient
}

type GrpcClients struct {
	brokerService broker_service.BrokerServiceClient
}

func NewGrpcClients(conf *config.Config) (*GrpcClients, error) {
	connBrokerService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.BrokerServiceHost, conf.BrokerServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		brokerService: broker_service.NewBrokerServiceClient(connBrokerService),
	}, nil
}

func (g *GrpcClients) BrokerService() broker_service.BrokerServiceClient {
	return g.brokerService
}
