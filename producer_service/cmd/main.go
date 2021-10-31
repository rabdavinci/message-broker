package main

import (
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rabdavinci/message-broker/producer_service/config"
	"github.com/rabdavinci/message-broker/producer_service/genproto/producer_service"
	"github.com/rabdavinci/message-broker/producer_service/pkg/logger"
	"github.com/rabdavinci/message-broker/producer_service/service"
	grpc_client "github.com/rabdavinci/message-broker/producer_service/service/grpc_clients"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Environment, "ur_go_user_service")
	defer logger.Cleanup(log)

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
		"disable",
	)
	// conStr = `host=localhost port=5432 user=postgres password=20072003 dbname=producer_service sslmode=disable`
	db, err := sqlx.Connect("postgres", conStr)
	if err != nil {
		log.Error("error while connecting database", logger.Error(err))
		return
	}
	client, err := grpc_client.NewGrpcClients(&cfg)
	if err != nil {
		log.Error("error while connecting to clients", logger.Error(err))
		return
	}

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		return
	}

	producerService := service.NewProducerService(db, log, client)

	s := grpc.NewServer()
	reflection.Register(s)

	producer_service.RegisterProducerServiceServer(s, producerService)

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Error("error while listening: %v", logger.Error(err))
	}
}
