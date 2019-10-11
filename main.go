package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"imjuni.logrustest.com/tools"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server Healthy")
	})

	e.GET("/hello", func(c echo.Context) error {
		who := c.QueryParam("who")
		return c.String(http.StatusOK, tools.Hello(who))
	})

	logrus.Infof("Start Server - :%d", 8082)
	e.Start(fmt.Sprintf(":%d", 8082))
}
