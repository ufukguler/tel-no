package main

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"io"
	goLog "log"
	"telno/config"
)

func main() {
	goLog.SetOutput(io.Discard)
	e := echo.New()
	initServer(e)

	port := config.GetEnv("SERVER_PORT")

	if err := e.Start(port); err != nil {
		log.Panic(err)
	}
}
