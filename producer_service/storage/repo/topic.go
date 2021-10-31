package repo

import pb "github.com/rabdavinci/message-broker/producer_service/genproto/producer_service"

type TopicRepoI interface {
	Create(user *pb.Topic) (string, error)
	Update(user *pb.Topic) (string, error)
	Get(id string) (*pb.Topic, error)
	GetAll(req *pb.GetAllTopicRequest) (*pb.GetAllTopicResponse, error)
	Delete(id string) error
}
