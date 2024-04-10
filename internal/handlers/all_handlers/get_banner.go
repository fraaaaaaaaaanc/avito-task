package allHandlers

import (
	"avito-tech/internal/logger"
	hlModel "avito-tech/internal/models/hanlders_models"
	"github.com/gorilla/schema"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handlers) GetBanner(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logger.Error("Error parse query", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bannerModel := new(hlModel.GetBannerModel)
	if err := schema.NewDecoder().Decode(bannerModel, r.Form); err != nil {
		logger.Error("invalid request data format", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
