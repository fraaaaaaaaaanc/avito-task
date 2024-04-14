package allHandlers

import (
	"avito-tech/internal/logger"
	storageModels "avito-tech/internal/models/storage_model"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handlers) GetVersionBannerHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		logger.Error("error parse query", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h.errorInternalServer)
		return
	}

	stringBannerId := strings.Split(r.URL.Path, "/")[2]
	bannerId, err := strconv.Atoi(stringBannerId)
	if err != nil {
		logger.Error("invalid request data format", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(h.errorIncorrectData)
		return
	}

	bannerVersions, err := h.strg.GetVersionBanner(r.Context(), bannerId)
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
	if err = enc.Encode(bannerVersions); err != nil {
		logger.Error("error forming the response body", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h.errorInternalServer)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	logger.Info("handler /banner/versions/{id} worked correctly, the status 200 was received")
}
