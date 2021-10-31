package postgres_test

import (
	"testing"

	pb "github.com/rabdavinci/message-broker/user_service/genproto/user_service"
	"github.com/stretchr/testify/assert"
)

func createUser(t *testing.T) *pb.User {

	user := &pb.User{
		Id:   createRandomId(t),
		Name: fakeData.UserName(),
	}

	res, err := strg.UserService().Create(user)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	return user
}

func deleteUser(t *testing.T, id string) {
	err := strg.UserService().Delete(id)

	assert.NoError(t, err)
}

func TestNewUserRepo(t *testing.T) {
	user := createUser(t)
	deleteUser(t, user.Id)
}

func TestUpdateUser(t *testing.T) {
	user := createUser(t)

	user.Name = fakeData.UserName()

	id, err := strg.UserService().Update(user)

	assert.NotEmpty(t, id)
	assert.NoError(t, err)
	deleteUser(t, user.Id)
}

func TestUserGet(t *testing.T) {
	user := createUser(t)

	_, err := strg.UserService().Get(user.Id)

	assert.NoError(t, err)
	deleteUser(t, user.Id)
}

func TestUserGetAll(t *testing.T) {
	user := createUser(t)
	resp, err := strg.UserService().GetAll(&pb.GetAllUserRequest{
		Offset: 0,
		Limit:  10,
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Users)
	deleteUser(t, user.Id)
}

func TestUserDelete(t *testing.T) {
	user := createUser(t)
	deleteUser(t, user.Id)
}
