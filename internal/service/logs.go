package service

import (
	"context"
	"encoding/json"

	"github.com/xopxe23/articles-log/internal/domain"
)

type LogsRepository interface {
	Insert(ctx context.Context, input domain.LogItem) error
}

type LogsService struct {
	repo LogsRepository
}

func NewLogsService(repo LogsRepository) *LogsService {
	return &LogsService{repo: repo}
}

func (s *LogsService) Insert(ctx context.Context, logString []byte) error {
	var logItem domain.LogItem
	err := json.Unmarshal(logString, &logItem)
	if err != nil {
		return err
	}
	return s.repo.Insert(ctx, logItem)
}
