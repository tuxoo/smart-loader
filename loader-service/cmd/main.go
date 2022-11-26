package main

import "github.com/tuxoo/smart-loader/loader-service/internal/app"

const (
	configPath = "config/config"
)

func main() {
	app.Run(configPath)
}
