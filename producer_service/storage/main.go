package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/rabdavinci/message-broker/producer_service/storage/postgres"
	"github.com/rabdavinci/message-broker/producer_service/storage/repo"
)

type StorageI interface {
	TopicService() repo.TopicRepoI
}

type storagePg struct {
	db        *sqlx.DB
	topicRepo repo.TopicRepoI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:        db,
		topicRepo: postgres.NewTopicRepo(db),
	}
}

func (s *storagePg) TopicService() repo.TopicRepoI {
	return s.topicRepo
}
