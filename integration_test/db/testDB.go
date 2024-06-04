package db

import (
	"sync"

	"gorm.io/gorm"
	"main.go/config"
	"main.go/leaves/database/connection"
	"main.go/leaves/model"
)

type testDB struct {
	db *gorm.DB
}

var once sync.Once
var instance *testDB

func InitDB(cfg config.DbConfig) *testDB {
	once.Do(func() {
		handler := connection.NewDBHandler(cfg)
		db := handler.Instance()
		instance = &testDB{db: db}
	})
	return instance
}

func (s *testDB) Seed() {

	err := s.db.AutoMigrate(model.Emp{}, model.Leave{})
	if err != nil {
		// logger.Error("error occurred while migrating schema. error: %v", err)
		return
	}
}

func GetDB() *gorm.DB {
	return instance.db
}
