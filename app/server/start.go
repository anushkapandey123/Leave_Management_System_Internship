package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	

	// "github.com/gorilla/mux"
	"main.go/config"
	"main.go/leaves/constants"
	"main.go/leaves/controller"
	"main.go/leaves/database/connection"
	"main.go/leaves/repository"
	"main.go/leaves/service"
)

func Init(cfg config.AppConfig) error {
	handler := connection.NewDBHandler(cfg.Database)
	db := handler.Instance()

	// router := mux.NewRouter()
	router := setupApp(cfg)

	employeeRepository := repository.NewEmployeeRepository(db)

	employeeService := service.NewEmplyeeService(employeeRepository)

	employeeController := controller.NewEmployeeController(employeeService)

	// router.HandleFunc("/details", employeeController.Detail).Methods("GET")
	router.GET(constants.DetailEndPoint, employeeController.Detail)
	

	err := start(router, cfg.Server)
	if err != nil {
		return err
	}
	return nil


}

func start(r *gin.Engine, cfg config.ServerConfig) error {
	s := &http.Server{
		Addr:         port(cfg),
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
	}
	err := s.ListenAndServe()
	if err != nil {
		return fmt.Errorf("unable to start gin server. error: %w", err)
	}
	return nil
}

func port(c config.ServerConfig) string {
	return fmt.Sprintf(":%d", c.Port)
}


func setupApp(cfg config.AppConfig) *gin.Engine {
	gin.SetMode(cfg.Server.GineMode)
	engine := gin.New()
	// binding.Validator = new(validator.DtoValidator)
	return setupMiddleware(engine, cfg)
}

func setupMiddleware(engine *gin.Engine, cfg config.AppConfig) *gin.Engine {
	// engine.Use(cors.SetupCORS())
	// engine.Use(ginzap.Ginzap(logger.GetLogger(), time.RFC3339, true))
	// engine.Use(ginzap.RecoveryWithZap(logger.GetLogger(), true))
	return engine
}

