package repository

import (
	"fmt"

	"github.com/Zindiks/lookinlabs-test-task/configs"
	"github.com/Zindiks/lookinlabs-test-task/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB(config configs.Config) (*gorm.DB, error) {

	// DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB.Host, config.DB.User, config.DB.Pass, config.DB.Name, config.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})

	return db, nil
}
