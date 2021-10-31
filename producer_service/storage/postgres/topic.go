package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	pb "github.com/rabdavinci/message-broker/producer_service/genproto/producer_service"
	"github.com/rabdavinci/message-broker/producer_service/storage/repo"
	"github.com/xtgo/uuid"
)

type topicRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewTopicRepo(db *sqlx.DB) repo.TopicRepoI {
	return &topicRepo{db: db}
}

func (r *topicRepo) Create(topic *pb.Topic) (string, error) {
	id := uuid.NewRandom()
	query := `INSERT INTO topic (
                            id,
                            name
						)
                        VALUES ($1, $2) `

	_, err := r.db.Exec(
		query,
		id.String(),
		topic.Name,
	)

	return id.String(), err
}

func (r *topicRepo) Update(topic *pb.Topic) (string, error) {
	query := `UPDATE topic 
                    SET
                        name = $1,
                        updated_at = current_timestamp
                WHERE id = $2`

	_, err := r.db.Exec(
		query,
		topic.Name,
		topic.Id,
	)
	if err != nil {
		return "", err
	}

	return topic.Id, nil
}

func (r *topicRepo) Get(id string) (*pb.Topic, error) {
	var topic pb.Topic
	query := `SELECT 
                    id,
                    name
                FROM topic
                WHERE id = $1 `

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&topic.Id,
		&topic.Name,
	)
	if err != nil {
		return nil, err
	}

	return &topic, nil
}

func (r *topicRepo) GetAll(req *pb.GetAllTopicRequest) (*pb.GetAllTopicResponse, error) {
	var (
		filter string
		count  int32
		users  []*pb.Topic
		args   = make(map[string]interface{})
	)

	if req.Name != "" {
		filter += " AND name ilike '%' || :name || '%' "
		args["name"] = req.Name
	}

	countQuery := `SELECT count(1) FROM topic WHERE deleted_at = 0 ` + filter
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
                    id,
                    name
                FROM topic 
                WHERE deleted_at = 0 %s`
	rows, err = r.db.NamedQuery(fmt.Sprintf(query, filter), args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var topic pb.Topic
		err = rows.Scan(
			&topic.Id,
			&topic.Name,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, &topic)
	}

	return &pb.GetAllTopicResponse{
		Topics: users,
		Count:  count,
	}, nil
}

func (r *topicRepo) Delete(id string) error {
	query := `UPDATE topic 
                SET 
                    deleted_at=date_part('epoch', CURRENT_TIMESTAMP)::int 
                WHERE id = $1 AND deleted_at=0 `

	_, err := r.db.Exec(query, id)
	return err
}
