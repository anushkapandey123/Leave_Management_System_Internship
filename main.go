package main

import (
	// "errors"
	// "errors"
	"fmt"
	// "flag"
	// "fmt"
	// "os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/leaves/constants"
	"main.go/leaves/controller"
	"main.go/leaves/model"
	"main.go/leaves/repository"
	"main.go/leaves/service"
	// "main.go/app/server"
	// "main.go/config"
	// "main.go/integration_test/db"
	// "main.go/leaves/constants"
	// "main.go/leaves/controller"
	// "main.go/leaves/model"
	// "main.go/leaves/repository"
	// "main.go/leaves/repository"
	// "main.go/leaves/service"
)

const (
	configFileKey     = "configFile"
	defaultConfigFile = ""
	configFileUsage   = "/path/to/configfile/wrto/pwd"
)

// @title Booking API
// @version 1.0
// @description This is a skyfox

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
//	@securityDefinitions.basic	BasicAuth

// func main() {
// 	var configFile string

// 	flag.StringVar(&configFile, configFileKey, defaultConfigFile, configFileUsage)
// 	flag.Parse()

// 	cfg, err := config.LoadConfig(configFile)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	err = server.Init(cfg)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }

var db *gorm.DB

func main() {
    // Initialize Gin router
    router := gin.Default()
	

	// Initialize database connection
	dsn := "host=localhost user=postgres password=pass123 dbname=employee_leave_management_system port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	fmt.Println("Database connection established")
	fmt.Println("in main :", db)

	// Auto migrate the User model
	db.AutoMigrate(&model.Emp{})
	db.AutoMigrate(&model.Leave{})

	fmt.Println(db.Migrator().CurrentDatabase())

	fmt.Println(db.Migrator().HasTable(&model.Emp{}))

	// db.Migrator().DropTable("emptable")

	fmt.Println(db.Migrator().HasTable(&model.Emp{}))

	
	

	employeeRepository := repository.NewEmployeeRepository(db)

	employeeService := service.NewEmployeeService(employeeRepository)

	employeeController := controller.NewEmployeeController(employeeService)

	// router.HandleFunc("/details", employeeController.Detail).Methods("GET")
	router.GET(constants.DetailEndPoint, employeeController.Detail)

	router.POST(constants.InsertLeaveEndPoint, employeeController.Insert)

	router.GET(constants.LeaveDetailsEndPoint, employeeController.LeaveDetails)
	
	// Run the server
	err = router.Run("localhost:8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}



