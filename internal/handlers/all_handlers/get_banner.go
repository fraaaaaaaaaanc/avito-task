package allHandlers

import (
	"avito-tech/internal/logger"
	hlModel "avito-tech/internal/models/hanlders_models"
	"encoding/json"
	"github.com/gorilla/schema"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handlers) GetBannerHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logger.Error("error parse query", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h.errorInternalServer)
		return
	}

	bannerModel := new(hlModel.GetBannerModel)
	if err := schema.NewDecoder().Decode(bannerModel, r.Form); err != nil {
		logger.Error("invalid request data format", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(h.errorIncorrectData)
		return
	}

	respBannerModel, err := h.strg.GetBanner(*bannerModel)
	if err != nil {
		logger.Error("error working with the database", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h.errorInternalServer)
		return
	}

	enc := json.NewEncoder(w)
	if err = enc.Encode(respBannerModel); err != nil {
		logger.Error("error forming the response body", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h.errorInternalServer)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	logger.Info("handler /banner worked correctly, the status 200 was received")
}
