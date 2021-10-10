package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gcp-kit/stalog"
	"github.com/labstack/echo/v4"
	"github.com/office-aska/lwgc/pkg/environ"
)

func SetErrorHandler(e *echo.Echo) {
	config := stalog.NewConfig("")
	config.RequestLogOut = os.Stderr       // request log to stderr
	config.ContextLogOut = os.Stdout       // context log to stdout
	config.Severity = stalog.SeverityDebug // only over INFO logs are logged
	e.Use(stalog.RequestLoggingWithEcho(config))
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		var code int
		var message interface{}
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			message = he.Message
		} else {
			code = http.StatusInternalServerError
			message = "Internal Server Error"
		}
		if environ.IsLocal() {
			_ = c.String(code, fmt.Sprintf("%v", err))
			fmt.Printf("\x1b[31m%s\x1b[0m\n", err)
			return
		}

		_ = c.JSON(code, message)
		r := c.Request()
		logger := stalog.RequestContextLogger(r)
		logger.AdditionalData["request"] = map[string]interface{}{
			"method": r.Method,
			"path":   r.URL.Path,
			"error":  err,
			"agent":  r.UserAgent(),
		}
		if code == http.StatusNotFound {
			logger.Warningf("NotFound")
		} else {
			logger.Errorf("HTTPErrorHandler %v", err)
		}
	}
}
