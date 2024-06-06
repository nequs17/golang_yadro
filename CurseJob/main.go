package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	//_ "path/to/your/project/docs" // путь до папки с docs

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Example API
// @version 1.0
// @description This is a sample server for a resource manager.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /

func main() {
	// Initialize Echo
	e := echo.New()

	// Load configuration
	loadConfig()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// JWT Middleware
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(viper.GetString("jwtSecret")),
	}))

	// Custom Middleware for timeouts
	e.Use(timeoutMiddleware)

	// Routes
	e.GET("/", home)
	e.GET("/api/v1/resources", getResources)
	e.POST("/api/v1/resources", createResource)
	e.PUT("/api/v1/resources/:id", updateResource)
	e.DELETE("/api/v1/resources/:id", deleteResource)

	// Swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// JIRA Routes
	e.GET("/api/v1/jira/updateProject", updateProject)
	e.GET("/api/v1/jira/projects", listProjects)

	// Start server
	port := viper.GetString("port")
	e.Logger.Fatal(e.Start(":" + port))
}

// @Summary Home
// @Description Welcome to the API
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func home(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Welcome to the API",
		"links": echo.Map{
			"resources": "/api/v1/resources",
			"jira":      "/api/v1/jira",
			// Add more links to other services here
		},
	})
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	requiredSettings := []string{"port", "resourceTimeout", "analyticsTimeout", "dbUser", "dbPassword", "dbHost", "dbPort", "dbName", "jwtSecret", "jiraUrl", "threadCount", "issueInOneRequest", "maxTimeSleep", "minTimeSleep"}
	for _, setting := range requiredSettings {
		if !viper.IsSet(setting) {
			log.Fatalf("The %s is not configured", setting)
		}
	}
}
