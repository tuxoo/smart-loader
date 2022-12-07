package util

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
)

type Hasher struct {
	cfg *config.AppConfig
}

func NewHasher(cfg *config.AppConfig) *Hasher {
	return &Hasher{
		cfg: cfg,
	}
}

func (h *Hasher) SHA1Hash(content string) string {
	hasher := sha1.New()
	hasher.Write([]byte(content))
	return fmt.Sprintf("%x", hasher.Sum([]byte(h.cfg.HashSalt)))
}

func (h *Hasher) SHA256Hash(content string) string {
	hasher := sha256.New()
	hasher.Write([]byte(content))
	return fmt.Sprintf("%x", hasher.Sum([]byte(h.cfg.HashSalt)))
}
