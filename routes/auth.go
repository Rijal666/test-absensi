package routes

import (
	"test-absensi/handlers"
	"test-absensi/pkg/middleware"
	"test-absensi/pkg/mysql"
	"test-absensi/repositories"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	r.POST("/register", middleware.UploadFile(h.Register))
	r.POST("/login", middleware.UploadFile(h.Login))
	r.GET("/logout", middleware.Auth(h.GetUser))
}