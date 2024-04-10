package allHandlers

import (
	"avito-tech/internal/logger"
	tokenModels "avito-tech/internal/models/token_models"
	"avito-tech/internal/token"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

// TODO: объяснить почему тут повторяющийся код и назвать причину по которой я его не убрал
func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	accountName := r.Header.Get("accountName")

	var resultToken tokenModels.ResultToken
	switch accountName {
	case "user":
		token, err := h.tokenAccount.NewAccountToken(accountName)
		if err != nil {
			logger.Error("error when creating a token during authorization for user", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		resultToken = tokenModels.ResultToken{
			Token: token,
		}
	case "admin":
		token, err := h.tokenAccount.NewAccountToken(accountName)
		if err != nil {
			logger.Error("error when creating a token during authorization for admin", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		resultToken = tokenModels.ResultToken{
			Token: token,
		}
	default:
		logger.Error("the non-existent value of the account parameter")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	if err := enc.Encode(resultToken); err != nil {
		logger.Error("error forming the response body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getToken(tokenAccount *token.TokenAccount, w http.ResponseWriter,
	accountName string) (*tokenModels.ResultToken, error) {
	token, err := tokenAccount.NewAccountToken(accountName)
	if err != nil {
		http.Error(w, "error creating the token", http.StatusInternalServerError)
		logger.Error("error when creating a token during authorization", zap.Error(err))
		return nil, err
	}
	resultToken := &tokenModels.ResultToken{
		Token: token,
	}

	return resultToken, nil
}
