package main

import "github.com/tuxoo/smart-loader/facade-service/internal/app"

const (
	configPath = "config/config"
)

func main() {
	app.Run(configPath)
}
