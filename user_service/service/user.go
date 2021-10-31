package service

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	pb "github.com/rabdavinci/message-broker/user_service/genproto/user_service"
	"github.com/rabdavinci/message-broker/user_service/pkg/helper"
	"github.com/rabdavinci/message-broker/user_service/pkg/logger"
	grpc_client "github.com/rabdavinci/message-broker/user_service/service/grpc_clients"
	"github.com/rabdavinci/message-broker/user_service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userService struct {
	logger  logger.Logger
	storage storage.StorageI
	client  *grpc_client.GrpcClients
}

func NewUserService(db *sqlx.DB, log logger.Logger, transaction *grpc_client.GrpcClients) *userService {
	return &userService{
		logger:  log,
		storage: storage.NewStoragePg(db),
		client:  transaction,
	}
}

func (s *userService) Create(ctx context.Context, req *pb.User) (*pb.UserId, error) {
	id, err := s.storage.UserService().Create(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating new user", req, codes.Internal)
	}

	return &pb.UserId{
		Id: id,
	}, nil
}

func (s *userService) Update(ctx context.Context, req *pb.User) (*pb.UserId, error) {
	id, err := s.storage.UserService().Update(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating new user", req, codes.Internal)
	}

	return &pb.UserId{
		Id: id,
	}, nil
}

func (s *userService) Get(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	user, err := s.storage.UserService().Get(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting user", req, codes.Internal)
	}

	return user, nil
}

func (s *userService) GetAll(ctx context.Context, req *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error) {
	users, err := s.storage.UserService().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all users", req, codes.Internal)
	}

	return users, nil
}

func (s *userService) Delete(ctx context.Context, req *pb.UserId) (*emptypb.Empty, error) {
	err := s.storage.UserService().Delete(req.Id)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while deleting user", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}

func (s *userService) Receiver(ctx context.Context, req *pb.UserMessage) (*emptypb.Empty, error) {
	user, err := s.storage.UserService().Get(req.UserId)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting user", req, codes.Internal)
	}

	fmt.Printf("\n %s получил сообщение: %s \n", user.Name, req.Message)
	return &emptypb.Empty{}, nil
}
