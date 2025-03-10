package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var messages []Message

func Gethandler(c echo.Context) error {
	return c.JSON(http.StatusOK, &messages)
}

func PostHandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}
	messages = append(messages, message)
	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message added successfully",
	})
}

func main() {
	e := echo.New()

	e.GET("/messages", Gethandler)
	e.POST("/messages", PostHandler)
	e.Start("localhost:8000")
}

// 15:35 video
