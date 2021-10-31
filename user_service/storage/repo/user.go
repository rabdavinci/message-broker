package repo

import pb "github.com/rabdavinci/message-broker/user_service/genproto/user_service"

type UserRepoI interface {
	Create(user *pb.User) (string, error)
	Update(user *pb.User) (string, error)
	Get(id string) (*pb.User, error)
	GetAll(req *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error)
	Delete(id string) error
}
