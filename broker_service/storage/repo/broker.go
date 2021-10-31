package repo

import pb "github.com/rabdavinci/message-broker/broker_service/genproto/broker_service"

type BrokerRepoI interface {
	AddUserTopic(broker *pb.UserTopic) (string, error)
	GetAll(req *pb.GetUserTopicsRequest) (*pb.GetUserTopicsResponse, error)
}
