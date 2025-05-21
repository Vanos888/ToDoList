package main

import (
	config "ToDoList/internal/config"
	"ToDoList/internal/pkg/interceptor"
	"context"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"

	authgen "ToDoList/generated/api/auth"
	taskgen "ToDoList/generated/api/task"
	"ToDoList/internal/client/telegram"
	authapi "ToDoList/internal/handler/auth-api"
	taskapi "ToDoList/internal/handler/task-api"
	"ToDoList/internal/repository/auth"
	"ToDoList/internal/repository/task"
	"ToDoList/internal/repository/user"
	authservice "ToDoList/internal/service/auth"
	taskservice "ToDoList/internal/service/task"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err.Error())
	}

	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       cfg.RedisDB,
	})
	authRepo := auth.NewAuthRepo(client)

	connString := cfg.PgDSN
	conn, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatal(err)
	}
	if err := conn.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	userRepo := user.NewUserRepo(conn)
	tgClient := telegram.NewTelegramClient()
	authServ := authservice.NewAuthService(authRepo, userRepo, tgClient, &cfg)
	taskRepo := task.NewTaskRepo(conn)
	taskService := taskservice.NewTaskService(taskRepo)

	taskAPI := taskapi.NewTaskApi(taskService)
	authAPI := authapi.NewAuthApi(authServ)

	s := grpc.NewServer(

		grpc.ChainUnaryInterceptor(interceptor.UnaryJWTServerInterceptor()),
	)

	authgen.RegisterAuthApiServer(s, authAPI)
	taskgen.RegisterTaskApiServer(s, taskAPI)

	list, err := net.Listen("tcp", cfg.GrpcServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = s.Serve(list)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
