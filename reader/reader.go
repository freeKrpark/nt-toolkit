package reader

type Reader interface {
	GetWords(dir string) ([]string, error)
}
