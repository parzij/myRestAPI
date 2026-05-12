package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "newRestAPI/docs"
	"newRestAPI/internal/router"
)

// @title           NewRestAPI
// @version         1.0
// @description     Простой REST API для работы с сообщениями
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.email  support@example.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /
// @schemes   http
func main() {
	if err := godotenv.Load(); err != nil {
	}

	port := os.Getenv("APP_PORT")
	
	e := echo.New()

	// Swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// API routes
	router.SetupRouter(e)

	e.Logger.Fatal(e.Start(port))
}