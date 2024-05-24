package connection

import (
	"errors"
	// "fmt"
	// "sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "main.go/config"
	// "main.go/leaves/database/common"
	"main.go/leaves/model"
)

// var once sync.Once

// type DBHandler interface {
// 	Instance() *common.BaseDB
// }

// type dbHandler struct {
// 	config config.DbConfig
// }

// func (dh *dbHandler) Instance() *common.BaseDB {

// 	var db *gorm.DB
// 	var err error

// 	once.Do(func() {
// 		dsn := connectionString(dh.config)
// 		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		})
// 	})
// 	if err != nil {
// 		panic("could not establish database connection")
// 	}
// 	return common.NewBaseDB(db)
// }

// func NewDBHandler(config config.DbConfig) *dbHandler {
// 	return &dbHandler{
// 		config: config,
// 	}
// }

// func connectionString(cfg config.DbConfig) string {
// 	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
// }

func Init() *gorm.DB {
    dbURL := "postgres://pg:pass123@localhost:5432/employee_leave_management_system"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        errors.New("error in connection")
    }

    db.AutoMigrate(&model.Emp{})

    return db
}