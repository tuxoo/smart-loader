package hasher

type Hasher interface {
	HashString(password string) string
}
