package allHandlers

import (
	"avito-tech/internal/logger"
	hlModel "avito-tech/internal/models/hanlders_models"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handlers) PostBanner(w http.ResponseWriter, r *http.Request) {
	var addBannerModel hlModel.PostBannerModel
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&addBannerModel); err != nil {
		logger.Error("invalid request format", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
