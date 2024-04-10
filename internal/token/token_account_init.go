package token

type TokenAccount struct {
	secretKeyJWTToken string
}

func NewTokenAccount(secretKeyJWTToken string) *TokenAccount {
	return &TokenAccount{
		secretKeyJWTToken: secretKeyJWTToken,
	}
}
