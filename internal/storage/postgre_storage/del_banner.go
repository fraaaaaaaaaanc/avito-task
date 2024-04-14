package postgreStorage

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (ps *PostgreStorage) DelBanner(ctx context.Context, idBanner int) error {
	var deletedID uint32

	if err := ps.db.QueryRow(deleteQuery, idBanner).
		Scan(
			&deletedID,
		); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.Wrapf(ErrorBannerNotFound, "with id %d", idBanner)
		}

		return errors.Wrapf(err, "can't delete banner with id %d", idBanner)
	}

	return nil
}