package app

import (
	"urlshorter/config"
	"urlshorter/database"
	"urlshorter/routes"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	config.InitConfig()

	database.StartDB()
	defer database.CloseDB()

	routes.InitRoutes(r, database.GetDB())

	r.Run(":8081")
}
