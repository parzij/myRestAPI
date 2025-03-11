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

func PutHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 || id >= len(messages) {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Invalid message ID",
		})
	}
	var updatedMessage Message
	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not update the message",
		})
	}
	messages[id] = updatedMessage
	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message updated successfully",
	})
}


func main() {
	ech := echo.New()

	ech.GET("/messages", Gethandler)
	ech.POST("/messages", PostHandler)
	ech.PUT("/messages/:id", PutHandler)

	ech.Logger.Fatal(ech.Start(":8080"))
}
