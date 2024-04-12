package allHandlers

import (
	"avito-tech/internal/logger"
	hlModel "avito-tech/internal/models/hanlders_models"
	storageModels "avito-tech/internal/models/storage_model"
	"encoding/json"
	"errors"
	"github.com/gorilla/schema"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handlers) GetUserBannerHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logger.Error("error parse query", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h.errorInternalServer)
		return
	}

	userBannerModel := new(hlModel.GetUserBannerModel)
	if err := schema.NewDecoder().Decode(userBannerModel, r.Form); err != nil {
		logger.Error("invalid request data format", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(h.errorIncorrectData)
		return
	}

	bannerContentModel, err := h.strg.GetUserBanner(*userBannerModel)
	if err != nil && !errors.Is(err, storageModels.ErrBannerNotFound) {
		logger.Error("error working with the database", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h.errorInternalServer)
		return
	} else if errors.Is(err, storageModels.ErrBannerNotFound) {
		logger.Error("no data was found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	enc := json.NewEncoder(w)
	if err = enc.Encode(bannerContentModel); err != nil {
		logger.Error("error forming the response body", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h.errorInternalServer)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	logger.Info("handler /user_banner worked correctly, the status 200 was received")
}
