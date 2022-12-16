package token_manager

type TokenManager interface {
	GenerateToken(userId string) (string, error)
	ParseToken(accessToken string) (string, error)
}
