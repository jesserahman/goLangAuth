package main

import (
	"goLangAuth/app"
	"goLangAuth/logger"
)

func main() {
	logger.Info("starting application...")
	app.Run()
}
