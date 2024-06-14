package main

import (
	"fmt"

	// "github.com/gin-contrib/cors"
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/leaves/constants"
	"main.go/leaves/controller"
	"main.go/leaves/model"
	"main.go/leaves/repository"
	"main.go/leaves/service"
	middleware "main.go/middleware/security"
)

const (
	dsn = "host=localhost user=postgres password=pass123 dbname=employee_leave_management_system port=5432 sslmode=disable"
)

var db *gorm.DB

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main() {
	// Initialize Gin router
	router := gin.Default()
	// router.Use(cors.Default())
	router.Use(CORSMiddleware())
	

	// Initialize database connection
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	fmt.Println("Database connection established")

	// AutoMigrate the models
	db.AutoMigrate(&model.Emp{}, &model.Leave{}, &model.User{})

	// Initialize repositories
	leaveRepository := repository.NewLeaveRepository(db)
	userRepository := repository.NewUserRepository(db)

	// Initialize services
	leaveService := service.NewLeaveService(leaveRepository)
	userService := service.NewUserService(userRepository)

	// Initialize controllers
	leaveController := controller.NewLeaveController(leaveService)
	userController := controller.NewUserController(userService)

	// Define routes
	router.POST(constants.InsertLeaveEndPoint, middleware.JWTAuthMiddleware(), leaveController.Insert)
	router.GET(constants.LeaveDetailsEndPoint,middleware.JWTAuthMiddleware(), leaveController.LeaveDetailsNew)
	router.POST(constants.DeleteLeaveEndPoint, leaveController.Delete)
	router.POST(constants.SignUpEndPoint, userController.SignUp)
	router.POST(constants.LoginEndPoint, userController.Login)
	router.GET(constants.TeamLeaveDetailsEndPoint, leaveController.LeaveDetails)

	// Run the server
	err = router.Run("localhost:8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}


