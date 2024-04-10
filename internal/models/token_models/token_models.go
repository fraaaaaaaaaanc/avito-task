package tokenModels

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

var ErrTokenIsNotValid = errors.New("token is not valid")

var ErrTokenEmpty = errors.New("the value for the token field is missing")

var ErrAccessDenied = errors.New("access denied")

type ResultToken struct {
	Token string `json:"token"`
}

type Claims struct {
	jwt.RegisteredClaims
	AccountName string
}
