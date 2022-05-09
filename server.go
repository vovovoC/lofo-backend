package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vovovoC/lofo-backend/config"
	"github.com/vovovoC/lofo-backend/controller"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthControlller()
)

func main() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run()
}
