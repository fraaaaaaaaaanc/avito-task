package tokenModels

import "errors"

var ErrTokenIsNotValid = errors.New("token is not valid")

var ErrTokenEmpty = errors.New("the value for the token field is missing")

var ErrAccessDenied = errors.New("access denied")

var ErrRoleIsNotValid = errors.New("role is not valid")

type ResultToken struct {
	Token string `json:"token"`
}
