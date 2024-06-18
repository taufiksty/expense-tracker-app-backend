package main

import (
	"github.com/gin-gonic/gin"
	"github.com/taufiksty/expense-tracker-app-backend/config"
	"github.com/taufiksty/expense-tracker-app-backend/routes"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	routes.SetupRoutes(r)

	r.Run(":3000")
}
