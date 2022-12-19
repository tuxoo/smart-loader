package downloader

type Downloader interface {
	Download(url string) ([]byte, error)
}
