package allHandlers

import (
	"avito-tech/internal/logger"
	hlModel "avito-tech/internal/models/hanlders_models"
	storageModels "avito-tech/internal/models/storage_model"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handlers) PatchBannerIdHandler(w http.ResponseWriter, r *http.Request) {
	var patchBannerModel hlModel.PatchBannerModel
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&patchBannerModel); err != nil {
		logger.Error("invalid request data format", zap.Error(err))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(h.errorIncorrectData)
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
	patchBannerModel.Id = bannerId

	err = h.strg.PatchBanner(r.Context(), &patchBannerModel)
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
	w.WriteHeader(http.StatusOK)
	logger.Info("handler /banner/{id} worked correctly, the status 200 was received")
}
