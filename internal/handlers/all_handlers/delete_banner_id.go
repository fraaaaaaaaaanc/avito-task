package allHandlers

import (
	"avito-tech/internal/logger"
	hlModel "avito-tech/internal/models/hanlders_models"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handlers) DeleteBannerId(w http.ResponseWriter, r *http.Request) {
	stringId := strings.Split(r.URL.Path, "/")[2]
	id, err := strconv.Atoi(stringId)
	if err != nil {
		logger.Error("invalid id parameter value", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	deleteBannerModel := hlModel.DeleteBannerModel{
		Id: id,
	}

	fmt.Println(deleteBannerModel)
}
