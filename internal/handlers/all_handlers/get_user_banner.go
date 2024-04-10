package allHandlers

import (
	"avito-tech/internal/logger"
	hlModel "avito-tech/internal/models/hanlders_models"
	"github.com/gorilla/schema"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handlers) GetUserBanner(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logger.Error("error parse query", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userBannerModel := new(hlModel.GetUserBannerModel)
	if err := schema.NewDecoder().Decode(userBannerModel, r.Form); err != nil {
		logger.Error("invalid request format", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
