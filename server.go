package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vovovoC/lofo-backend/config"
	"github.com/vovovoC/lofo-backend/controller"
	"github.com/vovovoC/lofo-backend/middleware"
	"github.com/vovovoC/lofo-backend/repository"
	"github.com/vovovoC/lofo-backend/service"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	// bookRepository repository.BookRepository = repository.NewBookRepository(db)
	userService service.UserService = service.NewUserService(userRepository)
	// bookService    service.BookService       = service.NewBookService(bookRepository)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()

	authRoutes := r.Group("api/auth", middleware.AuthorizeJWT(jwtService))
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run()
}
