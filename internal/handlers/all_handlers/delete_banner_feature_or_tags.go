package allHandlers

import (
	hlModel "avito-tech/internal/models/hanlders_models"
	"avito-tech/internal/logger"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handlers) DeleteBannerFeatureOrTag(w http.ResponseWriter, r *http.Request) {
	var deleteBannerFeatureOrTag hlModel.DeleteBannerFeatureOrTagModel
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&deleteBannerFeatureOrTag); err != nil {
		logger.Error("invalid request data format", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(h.errorIncorrectData)
		return
	}
	
	// h.delQueryChan <- deleteBannerFeatureOrTag
}