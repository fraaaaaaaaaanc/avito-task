package postgreStorage

import (
	storageModels "avito-tech/internal/models/storage_model"
	"context"
)

func (ps *PostgreStorage) DelBannerFeatureOrTag(ctx context.Context, tagChan []storageModels.DelFeatureOrTagChan) error {
	return nil
}
