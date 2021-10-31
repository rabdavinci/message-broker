package postgres_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/rabdavinci/message-broker/user_service/config"
	"github.com/rabdavinci/message-broker/user_service/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/manveru/faker"
	"github.com/stretchr/testify/assert"
)

var (
	postgresConn *sqlx.DB
	err          error
	cfg          config.Config
	strg         storage.StorageI
	fakeData     *faker.Faker
)

func createRandomId(t *testing.T) string {
	id, err := uuid.NewRandom()
	assert.NoError(t, err)
	return id.String()
}

func TestMain(m *testing.M) {
	cfg = config.Load()

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
		"disable",
	)
	conStr = `host=localhost port=5432 user=postgres password=20072003 dbname=user_service sslmode=disable`
	fakeData, _ = faker.New("en")
	postgresConn, err = sqlx.Open("postgres", conStr)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(postgresConn)

	strg = storage.NewStoragePg(postgresConn)

	os.Exit(m.Run())
}
