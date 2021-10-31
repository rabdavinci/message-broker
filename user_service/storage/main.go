package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/rabdavinci/message-broker/user_service/storage/postgres"
	"github.com/rabdavinci/message-broker/user_service/storage/repo"
)

type StorageI interface {
	UserService() repo.UserRepoI
}

type storagePg struct {
	db          *sqlx.DB
	companyRepo repo.UserRepoI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:          db,
		companyRepo: postgres.NewUserRepo(db),
	}
}

func (s *storagePg) UserService() repo.UserRepoI {
	return s.companyRepo
}
