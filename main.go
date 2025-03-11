package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	if len(messages) == 0 {
		return c.JSON(http.StatusNotFound, Response{
			Status:  "Error",
			Message: "No messages found",
		})
	}
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
	ech := echo.New()

	ech.GET("/messages", Gethandler)
	ech.POST("/messages", PostHandler

	ech.Logger.Fatal(ech.Start(":8080"))
}
