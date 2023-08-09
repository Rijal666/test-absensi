package main

import (
	"fmt"
	"test-absensi/migration"
	"test-absensi/routes"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	"test-absensi/pkg/mysql"
)

func main() {

	errEnv := godotenv.Load()
    if errEnv != nil {
      panic("Failed to load env file")
    }

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	mysql.DatabaseInit()
	migration.RunAutoMigrate()
	
	routes.RouteInit(r.Group("/api/v1"))

	r.Static("/uploads", "./uploads")

	fmt.Println("server running localhost:5002")
	r.Run("localhost:5002")
}