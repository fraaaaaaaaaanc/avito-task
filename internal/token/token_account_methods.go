package token

import (
	"avito-tech/internal/logger"
	cookieModels "avito-tech/internal/models/token_models"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const timeLiveToken = time.Hour * 6

func (t TokenAccount) buildJwtString(accountName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &cookieModels.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(timeLiveToken)),
		},
		AccountName: accountName,
	})

	tokenString, err := token.SignedString([]byte(t.secretKeyJWTToken))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t TokenAccount) NewAccountToken(accountName string) (string, error) {
	tokenString, err := t.buildJwtString(accountName)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t TokenAccount) getAccountName(tokenString string) (string, error) {
	claims := &cookieModels.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secretKeyJWTToken), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", cookieModels.ErrTokenIsNotValid
	}
	return claims.AccountName, nil
}

func (t TokenAccount) MiddlewareCheckToken() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("token")
			if tokenString == "" {
				logger.Error("the r.Header.Get(\"token\") parameter is missing",
					zap.Error(cookieModels.ErrTokenEmpty))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			accountName, err := t.getAccountName(tokenString)
			if err != nil {
				logger.Error("an error occurred while working with the authorization token", zap.Error(err))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if accountName == "user" && r.URL.Path != "/user_banner" {
				logger.Error(fmt.Sprint("the user is denied access to the address: %s", r.URL.Path),
					zap.Error(cookieModels.ErrAccessDenied))
				w.WriteHeader(http.StatusForbidden)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}
