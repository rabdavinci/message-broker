package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rabdavinci/message-broker/producer_service/genproto/broker_service"
	pb "github.com/rabdavinci/message-broker/producer_service/genproto/producer_service"
	"github.com/rabdavinci/message-broker/producer_service/pkg/helper"
	"github.com/rabdavinci/message-broker/producer_service/pkg/logger"
	grpc_client "github.com/rabdavinci/message-broker/producer_service/service/grpc_clients"
	"github.com/rabdavinci/message-broker/producer_service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type producerService struct {
	logger  logger.Logger
	storage storage.StorageI
	client  *grpc_client.GrpcClients
}

func NewProducerService(db *sqlx.DB, log logger.Logger, transaction *grpc_client.GrpcClients) *producerService {
	return &producerService{
		logger:  log,
		storage: storage.NewStoragePg(db),
		client:  transaction,
	}
}

func (s *producerService) CreateTopic(ctx context.Context, req *pb.Topic) (*pb.TopicId, error) {
	id, err := s.storage.TopicService().Create(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating new topic", req, codes.Internal)
	}

	return &pb.TopicId{
		Id: id,
	}, nil
}

func (s *producerService) Update(ctx context.Context, req *pb.Topic) (*pb.TopicId, error) {
	id, err := s.storage.TopicService().Update(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating new topic", req, codes.Internal)
	}

	return &pb.TopicId{
		Id: id,
	}, nil
}

func (s *producerService) Get(ctx context.Context, req *pb.TopicId) (*pb.Topic, error) {
	topic, err := s.storage.TopicService().Get(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting topic", req, codes.Internal)
	}

	return topic, nil
}

func (s *producerService) GetAll(ctx context.Context, req *pb.GetAllTopicRequest) (*pb.GetAllTopicResponse, error) {
	companies, err := s.storage.TopicService().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all topics", req, codes.Internal)
	}

	return companies, nil
}

func (s *producerService) Delete(ctx context.Context, req *pb.TopicId) (*emptypb.Empty, error) {
	err := s.storage.TopicService().Delete(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while deleting topic", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}

func (s *producerService) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*emptypb.Empty, error) {
	_, err := s.client.BrokerService().SendMessageToUser(
		context.Background(),
		&broker_service.TopicMessage{
			TopicId: req.TopicId,
			Message: req.Message,
		},
	)

	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while sending message to topic", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}
