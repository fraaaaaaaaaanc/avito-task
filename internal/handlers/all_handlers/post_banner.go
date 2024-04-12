package allHandlers

import (
	"avito-tech/internal/logger"
	hlModel "avito-tech/internal/models/hanlders_models"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handlers) PostBannerHandler(w http.ResponseWriter, r *http.Request) {
	var addBannerModel hlModel.PostBannerModel
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&addBannerModel); err != nil {
		logger.Error("invalid request data format", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(h.errorIncorrectData)
		return
	}

	respBannerModel, err := h.strg.SetBanner(addBannerModel)
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
	w.WriteHeader(http.StatusCreated)
	logger.Info("handler /banner worked correctly, the status 201 was received")
}
