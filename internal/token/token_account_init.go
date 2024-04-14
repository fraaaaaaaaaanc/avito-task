package token

const (
	roleAdmin = "admin"
	roleUser  = "user"
)

type TokenAccount struct {
	secretKeyJWTToken string
	roles             map[string]struct{}
}

func NewTokenAccount(secretKeyJWTToken string) *TokenAccount {
	roles := map[string]struct{}{
		roleUser:  {},
		roleAdmin: {},
	}
	return &TokenAccount{
		secretKeyJWTToken: secretKeyJWTToken,
		roles:             roles,
	}
}
