package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/rabdavinci/message-broker/broker_service/storage/postgres"
	"github.com/rabdavinci/message-broker/broker_service/storage/repo"
)

type StorageI interface {
	BrokerService() repo.BrokerRepoI
}

type storagePg struct {
	db         *sqlx.DB
	brokerRepo repo.BrokerRepoI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:         db,
		brokerRepo: postgres.NewBrokerRepo(db),
	}
}

func (s *storagePg) BrokerService() repo.BrokerRepoI {
	return s.brokerRepo
}
