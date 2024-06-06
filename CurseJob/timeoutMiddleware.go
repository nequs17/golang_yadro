package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func timeoutMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var timeout int
		if c.Path() == "/api/v1/connector/*" || c.Path() == "/api/v1/graph/*" {
			timeout = viper.GetInt("analyticsTimeout")
		} else {
			timeout = viper.GetInt("resourceTimeout")
		}

		timer := time.AfterFunc(time.Duration(timeout)*time.Second, func() {
			c.Error(&echo.HTTPError{
				Code:    http.StatusRequestTimeout,
				Message: "Request Timeout",
			})
		})
		defer timer.Stop()

		return next(c)
	}
}
