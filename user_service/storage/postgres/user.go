package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	pb "github.com/rabdavinci/message-broker/user_service/genproto/user_service"
	"github.com/rabdavinci/message-broker/user_service/storage/repo"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo(db *sqlx.DB) repo.UserRepoI {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *pb.User) (string, error) {
	id, _ := uuid.NewRandom()
	query := `INSERT INTO users (
                            id,
                            name
						)
                        VALUES ($1, $2) `

	_, err := r.db.Exec(
		query,
		id,
		user.Name,
	)

	return id.String(), err
}

func (r *userRepo) Update(user *pb.User) (string, error) {
	query := `UPDATE users 
                    SET
                        name = $1,
                        updated_at = current_timestamp
                WHERE id = $2`

	_, err := r.db.Exec(
		query,
		user.Name,
		user.Id,
	)
	if err != nil {
		return "", err
	}

	return user.Id, nil
}

func (r *userRepo) Get(id string) (*pb.User, error) {
	var user pb.User
	query := `SELECT 
                    id,
                    name
                FROM users
                WHERE deleted_at = 0 AND id = $1 `

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&user.Id,
		&user.Name,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetAll(req *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error) {
	var (
		filter string
		count  int32
		users  []*pb.User
		args   = make(map[string]interface{})
	)

	if req.Name != "" {
		filter += " AND name ilike '%' || :name || '%' "
		args["name"] = req.Name
	}

	countQuery := `SELECT count(1) FROM users WHERE deleted_at = 0 ` + filter
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
                FROM users 
                WHERE deleted_at = 0 %s`
	rows, err = r.db.NamedQuery(fmt.Sprintf(query, filter), args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user pb.User
		err = rows.Scan(
			&user.Id,
			&user.Name,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return &pb.GetAllUserResponse{
		Users: users,
		Count: count,
	}, nil
}

func (r *userRepo) Delete(id string) error {
	query := `UPDATE users 
                SET 
                    deleted_at=date_part('epoch', CURRENT_TIMESTAMP)::int 
                WHERE id = $1 AND deleted_at=0 `

	_, err := r.db.Exec(query, id)
	return err
}
