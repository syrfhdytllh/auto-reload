package routers

import (
	"auto-reload/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.PUT("/save/:id", controllers.UpdateStatus)

	return router
}
