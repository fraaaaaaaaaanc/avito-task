package allHandlers

import (
	"avito-tech/internal/logger"
	hlModel "avito-tech/internal/models/hanlders_models"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handlers) PatchBannerId(w http.ResponseWriter, r *http.Request) {
	var patchBannerModel hlModel.PatchBannerModel
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&patchBannerModel); err != nil {
		logger.Error("invalid request format", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	stringId := strings.Split(r.URL.Path, "/")[2]
	id, err := strconv.Atoi(stringId)
	if err != nil {
		logger.Error("invalid id parameter value", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	patchBannerModel.Id = id
}
