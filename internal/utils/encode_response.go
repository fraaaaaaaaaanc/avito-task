package utils

import (
	"avito-tech/internal/logger"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

func EncodeResponse[T Responses](w http.ResponseWriter, responseModel T) (http.ResponseWriter, error) {
	enc := json.NewEncoder(w)
	if err := enc.Encode(responseModel); err != nil {
		logger.Error("the response body could not be formed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return w, err
	}
	return w, nil
}

//if err != nil && !errors.Is(err, storageModels.ErrBannerNotFound) {
//	logger.Error("error working with the database", zap.Error(err))
//	wResponse, err := utils.EncodeResponse(w)
//	w = wResponse
//	if err != nil {
//		logger.Error("internal server error", zap.Error(err))
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write(h.errorMessage)
//		return
//	} else if errors.Is(err, storageModels.ErrBannerNotFound) {
//		logger.Error("no data was found", zap.Error(err))
//		wResponse, err := utils.EncodeResponse(w, bannerContenModel)
//		w = wResponse
//		if err != nil {
//			logger.Error("internal server error", zap.Error(err))
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write(h.errorMessage)
//			return
//
