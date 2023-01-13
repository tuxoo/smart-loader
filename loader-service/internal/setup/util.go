package setup

import (
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model/config"
	"github.com/tuxoo/smart-loader/loader-service/internal/util/downloader"
	"github.com/tuxoo/smart-loader/loader-service/internal/util/hasher"
)

func provideDownloader() downloader.Downloader {
	return downloader.NewHttpDownloader()
}

func provideHasher(cfg *config.AppConfig) hasher.Hasher {
	return hasher.NewSHA256Hasher(cfg.HashSalt)
}
