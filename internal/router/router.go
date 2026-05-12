package router

import (
	"newRestAPI/internal/handler"

	"github.com/labstack/echo/v4"
)

func SetupRouter (e *echo.Echo) {
	e.GET("/messages", handler.Gethandler)
	e.POST("/messages", handler.PostHandler)
	e.PUT("/messages/:id", handler.PutHandler)
	e.DELETE("/messages/:id", handler.DeleteHandler)
}
