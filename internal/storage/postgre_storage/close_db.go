package postgreStorage

import (
	"avito-tech/internal/logger"
	"go.uber.org/zap"
)

func (ps *PostgreStorage) CloseDB() {
	if err := ps.db.Close(); err != nil {
		logger.Error("error closing database connection", zap.Error(err))
	}
}