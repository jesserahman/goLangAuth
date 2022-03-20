package main

import (
	"github.com/jesserahman/goLangAuth/app"
	"github.com/jesserahman/goLangAuth/logger"
)

func main() {
	logger.Info("starting application...")
	app.Run()
}
