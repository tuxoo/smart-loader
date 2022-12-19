package hasher

import (
	"crypto/sha1"
	"fmt"
)

type SHA1Hasher struct {
	salt string
}

func NewSHA1Hasher(salt string) *SHA1Hasher {
	return &SHA1Hasher{
		salt: salt,
	}
}

func (h *SHA1Hasher) HashString(content string) string {
	hasher := sha1.New()
	hasher.Write([]byte(content))
	return fmt.Sprintf("%x", hasher.Sum([]byte(h.salt)))
}
