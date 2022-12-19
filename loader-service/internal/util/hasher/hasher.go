package hasher

type Hasher interface {
	HashString(str string) string
	HashBytes(bytes []byte) string
}
