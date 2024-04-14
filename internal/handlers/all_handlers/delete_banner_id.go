package allHandlers

import (
	"avito-tech/internal/logger"
	storageModels "avito-tech/internal/models/storage_model"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handlers) DeleteBannerIdHandler(w http.ResponseWriter, r *http.Request) {
	stringIdBanner := strings.Split(r.URL.Path, "/")[2]
	idBanner, err := strconv.Atoi(stringIdBanner)
	if err != nil {
		logger.Error("invalid request data format", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(h.errorIncorrectData)
		return
	}

	err = h.strg.DelBanner(r.Context(), idBanner)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	logger.Info("handler /banner/{id} worked correctly, the status 200 was received")
}
