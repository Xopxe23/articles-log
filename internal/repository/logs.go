package repository

import (
	"context"

	"github.com/xopxe23/articles-log/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogsRepository struct {
	DB *mongo.Database
}

func NewLogsRepository(db *mongo.Database) *LogsRepository {
	return &LogsRepository{DB: db}
}

func (r *LogsRepository) Insert(ctx context.Context, input domain.LogItem) error {
	_, err := r.DB.Collection("logs").InsertOne(ctx, input)

	return err
}
