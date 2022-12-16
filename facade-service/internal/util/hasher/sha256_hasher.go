package hasher

import (
	"crypto/sha256"
	"fmt"
)

type SHA256Hasher struct {
	salt string
}

func NewSHA256Hasher(salt string) *SHA256Hasher {
	return &SHA256Hasher{
		salt: salt,
	}
}

func (h *SHA256Hasher) Hash(content string) string {
	hasher := sha256.New()
	hasher.Write([]byte(content))
	return fmt.Sprintf("%x", hasher.Sum([]byte(h.salt)))
}
