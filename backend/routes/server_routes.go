package routes

import (
	"github.com/LuizFelipeVPCrema/minecraft-dashboard-go/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		servers := api.Group("/servers")
		{
			servers.POST("/", controllers.CreateServer)
			servers.GET("/", controllers.ListServers)
			servers.DELETE("/:name", controllers.DeleteServer)

			servers.PUT("/:name/start", controllers.StartServer)
			servers.PUT("/:name/stop", controllers.StopServer)
		}
	}
}
