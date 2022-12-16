package hasher

type Hasher interface {
	Hash(password string) string
}
