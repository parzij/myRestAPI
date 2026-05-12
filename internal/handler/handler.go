package handler

import (
	"net/http"
	"newRestAPI/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetMessages godoc
// @Summary      Get all messages
// @Description  Get a list of all stored messages
// @Tags         messages
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.Message
// @Failure      404  {object}  models.Response
// @Router       /messages [get]
func Gethandler(c echo.Context) error {
	if len(models.Messages) == 0 {
		return c.JSON(http.StatusNotFound, models.Response{
			Status:  "Error",
			Message: "No messages found",
		})
	}
	return c.JSON(http.StatusOK, &models.Messages)
}

// CreateMessage godoc
// @Summary      Create a new message
// @Description  Add a new message to the list
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        message  body  models.Message  true  "Message object"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /messages [post]
func PostHandler(c echo.Context) error {
	var message models.Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}
	models.Messages = append(models.Messages, message)
	return c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Message added successfully",
	})
}

// UpdateMessage godoc
// @Summary      Update a message
// @Description  Update an existing message by ID
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "Message ID"
// @Param        message  body  models.Message  true  "Updated message object"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /messages/{id} [put]
func PutHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 || id >= len(models.Messages) {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid message ID",
		})
	}
	var updatedMessage models.Message
	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Could not update the message",
		})
	}
	models.Messages[id] = updatedMessage
	return c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Message updated successfully",
	})
}

// DeleteMessage godoc
// @Summary      Delete a message
// @Description  Delete a message by ID
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "Message ID"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /messages/{id} [delete]
func DeleteHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 || id >= len(models.Messages) {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid message ID",
		})
	}
	models.Messages = append(models.Messages[:id], models.Messages[id+1:]...)
	return c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Message deleted successfully",
	})
}