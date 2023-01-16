package setup

import (
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/util/hasher"
	token_manager "github.com/tuxoo/smart-loader/facade-service/internal/util/token-manager"
)

func provideHasher(cfg *config.AppConfig) hasher.Hasher {
	return hasher.NewSHA256Hasher(cfg.HashSalt)
}

func provideTokenManager(cfg *config.AppConfig) token_manager.TokenManager {
	return token_manager.NewJWTTokenManager(cfg.AccessTokenTTL, cfg.SigningKey)
}
