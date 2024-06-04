package connection

import (
	"fmt"

	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/config"
	// "main.go/leaves/database/common"
)

var once sync.Once

type DBHandler interface {
	Instance() *gorm.DB
}

type dbHandler struct {
	config config.DbConfig
}

func (dh *dbHandler) Instance() *gorm.DB {

	var db *gorm.DB
	var err error

	once.Do(func() {
		dsn := connectionString(dh.config)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			//	todo - use the zap logger
		})
	})
	if err != nil {
		panic("could not establish database connection")
	}
	return db
	
}

func NewDBHandler(config config.DbConfig) *dbHandler {
	return &dbHandler{
		config: config,
	}
}

func connectionString(cfg config.DbConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
}