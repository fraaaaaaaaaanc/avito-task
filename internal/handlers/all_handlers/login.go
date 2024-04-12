package allHandlers

import (
	"avito-tech/internal/logger"
	tokenModels "avito-tech/internal/models/token_models"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

// TODO: объяснить почему тут повторяющийся код и назвать причину по которой я его не убрал
func (h *Handlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	accountName := r.URL.Query().Get("accountName")

	token, err := h.tokenAccount.NewAccountToken(accountName)
	if err != nil {
		logger.Error("error when creating a token during authorization for user", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resultToken := tokenModels.ResultToken{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	if err := enc.Encode(resultToken); err != nil {
		logger.Error("error forming the response body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info("handler /login worked correctly, the status 200 was received")
}
