package main

import (
	"coding-test/backend/controllers"
	"coding-test/backend/middleware"
	"coding-test/backend/models"
	"encoding/gob"

	"github.com/gin-gonic/gin"
)

func init() {
	controllers.Store.Options.HttpOnly = true
	controllers.Store.Options.Secure = true
	controllers.Store.Options.MaxAge = 86400 * 7
	gob.Register(&models.Login{})
}

func main() {
	router := setupServer()

	router.Run("0.0.0.0:3000")
}

func setupServer() *gin.Engine {
	router := gin.Default()

	api := router.Group("/")
	api.GET("/hello", controllers.Hello)
	api.POST("/sortnum", controllers.Sortnum)
	api.POST("/login", controllers.Login)
	api.GET("/is_auth", middleware.Require, controllers.Is_auth)

	return router
}
