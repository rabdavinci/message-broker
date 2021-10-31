package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	pb "github.com/rabdavinci/message-broker/broker_service/genproto/broker_service"
	"github.com/rabdavinci/message-broker/broker_service/storage/repo"
)

type brokerRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewBrokerRepo(db *sqlx.DB) repo.BrokerRepoI {
	return &brokerRepo{db: db}
}

func (r *brokerRepo) AddUserTopic(broker *pb.UserTopic) (string, error) {
	query := `INSERT INTO broker (
                            user_id,
                            topic_id
						)
                        VALUES ($1, $2) `

	_, err := r.db.Exec(
		query,
		broker.UserId,
		broker.TopicId,
	)

	return broker.UserId, err
}

func (r *brokerRepo) GetAll(req *pb.GetUserTopicsRequest) (*pb.GetUserTopicsResponse, error) {
	fmt.Println(req)
	var (
		filter  string
		count   int32
		brokers []*pb.UserTopic
		args    = make(map[string]interface{})
	)

	if req.UserId != "" {
		filter += " AND user_id = :user_id"
		args["user_id"] = req.UserId
	}

	if req.TopicId != "" {
		filter += " AND topic_id = :topic_id"
		args["topic_id"] = req.TopicId
	}

	countQuery := `SELECT count(1) FROM broker WHERE deleted_at = 0 ` + filter
	rows, err := r.db.NamedQuery(countQuery, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return nil, err
		}
	}

	filter += " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
	args["limit"] = req.Limit
	args["offset"] = req.Offset

	query := `SELECT
					user_id,
					topic_id
                FROM broker 
                WHERE deleted_at = 0 %s`
	rows, err = r.db.NamedQuery(fmt.Sprintf(query, filter), args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var broker pb.UserTopic
		err = rows.Scan(
			&broker.UserId,
			&broker.TopicId,
		)
		if err != nil {
			return nil, err
		}

		brokers = append(brokers, &broker)
	}

	return &pb.GetUserTopicsResponse{
		UserTopics: brokers,
		Count:      count,
	}, nil
}
