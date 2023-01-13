package downloader

import "io"

type Downloader interface {
	Download(url string) (io.ReadCloser, error)
}
