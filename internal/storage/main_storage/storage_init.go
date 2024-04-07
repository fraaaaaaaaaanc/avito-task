package mainStorage

import (
	"avito-tech/internal/logger"
	"context"
	"database/sql"
	"go.uber.org/zap"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(storageDBAddress string) (*Storage, error) {
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

	_, err = db.ExecContext(ctx, ``)

	if err != nil {
		return nil, err
	}
	return &Storage{
		db: db,
	}, nil
}
