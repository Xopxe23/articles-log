package main

import (
	"context"
	"log"
	"time"

	"github.com/xopxe23/articles-log/config"
	"github.com/xopxe23/articles-log/internal/repository"
	"github.com/xopxe23/articles-log/internal/server"
	"github.com/xopxe23/articles-log/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("error occured init config: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})
	opts.ApplyURI(cfg.DB.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("error occured mongo client connect: %s", err.Error())
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	db := dbClient.Database(cfg.DB.Database)
	logsRepos := repository.NewLogsRepository(db)
	logsService := service.NewLogsService(logsRepos)
	srv, err := server.NewServer(logsService)
	if err != nil {
		log.Fatal(err)
	}
	err = srv.Consume("logs")
	if err != nil {
		log.Fatal(err)
	}
}
