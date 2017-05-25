package server

import (
    "net/http"
    "github.com/labstack/echo"
)

func router(e *echo.Echo) {
    /*
     * Hello World
     */
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
}
