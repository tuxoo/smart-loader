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

func (d *HttpDownloader) Download(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unsuccessful  status [%s]", response.Status)
	}

	return response.Body, nil
}
