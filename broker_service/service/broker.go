package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	pb "github.com/rabdavinci/message-broker/broker_service/genproto/broker_service"
	"github.com/rabdavinci/message-broker/broker_service/genproto/producer_service"
	"github.com/rabdavinci/message-broker/broker_service/genproto/user_service"
	"github.com/rabdavinci/message-broker/broker_service/pkg/helper"
	"github.com/rabdavinci/message-broker/broker_service/pkg/logger"
	grpc_client "github.com/rabdavinci/message-broker/broker_service/service/grpc_clients"
	"github.com/rabdavinci/message-broker/broker_service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type brokerService struct {
	logger  logger.Logger
	storage storage.StorageI
	client  *grpc_client.GrpcClients
}

func NewUserService(db *sqlx.DB, log logger.Logger, transaction *grpc_client.GrpcClients) *brokerService {
	return &brokerService{
		logger:  log,
		storage: storage.NewStoragePg(db),
		client:  transaction,
	}
}

func (s *brokerService) AddUserTopic(ctx context.Context, req *pb.UserTopic) (*emptypb.Empty, error) {
	_, err := s.storage.BrokerService().AddUserTopic(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating new user topic", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}

func (s *brokerService) SendMessageToUser(ctx context.Context, req *pb.TopicMessage) (*emptypb.Empty, error) {
	userTopics, err := s.storage.BrokerService().GetAll(&pb.GetUserTopicsRequest{
		Offset:  0,
		Limit:   1000,
		TopicId: req.TopicId,
	})
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting user", req, codes.Internal)
	}

	for _, user_topic := range userTopics.UserTopics {
		_, err := s.client.UserService().Receiver(
			context.Background(),
			&user_service.UserMessage{
				UserId:  user_topic.UserId,
				Message: req.Message,
			},
		)

		if err != nil {
			return nil, helper.HandleError(s.logger, err, "error while sending message to user", req, codes.Internal)
		}
	}

	return &emptypb.Empty{}, nil
}

func (s *brokerService) GetUserTopic(ctx context.Context, req *pb.GetUserTopicsRequest) (*pb.GetUserTopicsResponse, error) {
	userTopics, err := s.storage.BrokerService().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all user_topics", req, codes.Internal)
	}

	for index, user_topic := range userTopics.UserTopics {
		resp, err := s.client.ProducerService().Get(
			context.Background(),
			&producer_service.TopicId{
				Id: user_topic.TopicId,
			},
		)

		if err != nil {
			return nil, helper.HandleError(s.logger, err, "error while getting topic", req, codes.Internal)
		}

		userTopics.UserTopics[index].TopicName = resp.Name
	}

	return userTopics, nil
}
