package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
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

func main() {
	// Initialize Gin router
	router := gin.Default()
	router.Use(cors.Default())

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
	employeeRepository := repository.NewEmployeeRepository(db)
	userRepository := repository.NewUserRepository(db)

	// Initialize services
	employeeService := service.NewEmployeeService(employeeRepository)
	userService := service.NewUserService(userRepository)

	// Initialize controllers
	employeeController := controller.NewEmployeeController(employeeService)
	userController := controller.NewUserController(userService)

	// Define routes
	router.POST(constants.InsertLeaveEndPoint, middleware.BasicAuth(userService), employeeController.Insert)
	router.GET(constants.LeaveDetailsEndPoint, middleware.BasicAuth(userService), employeeController.LeaveDetails)
	router.POST(constants.DeleteLeaveEndPoint, middleware.BasicAuth(userService), employeeController.Delete)
	router.POST(constants.SignUpEndPoint, userController.SignUp)
	router.POST(constants.LoginEndPoint, userController.Login)

	// Run the server
	err = router.Run("localhost:8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
