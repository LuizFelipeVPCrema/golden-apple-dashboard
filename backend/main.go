package main

import (
	"log"

	"github.com/LuizFelipeVPCrema/minecraft-dashboard-go/config"
	"github.com/LuizFelipeVPCrema/minecraft-dashboard-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	router := gin.Default()
	routes.RegisterRouter(router)

	log.Fatal(router.Run(":8080"))
}
