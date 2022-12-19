package downloader

import (
	"fmt"
	"io"
	"net/http"
)

type HttpDownloader struct {
}

func NewHttpDownloader() *HttpDownloader {
	return &HttpDownloader{}
}

func (d *HttpDownloader) Download(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unsuccessful  status [%s]", response.Status)
	}

	return io.ReadAll(response.Body) // TODO: is required to download by chunks
}
