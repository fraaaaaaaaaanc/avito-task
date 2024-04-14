package postgreStorage

import (
	"avito-tech/internal/logger"
	"avito-tech/internal/storage"
	"context"
	"database/sql"

	"go.uber.org/zap"
)

type PostgreStorage struct {
	db *sql.DB
}

func NewStorage(storageDBAddress string) (storage.StorageBanner, error) {
	db, err := sql.Open("pgx", storageDBAddress)
	if err != nil {
		logger.Error("error connect to the database", zap.Error(err))
		return nil, err
	}

	ctx, cansel := context.WithCancel(context.Background())
	defer cansel()

	if err = db.PingContext(ctx); err != nil {
		logger.Error("error ping to the database", zap.Error(err))
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &PostgreStorage{
		db: db,
	}, nil
}