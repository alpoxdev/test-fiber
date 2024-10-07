package database

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"test-fiber/config"
	"test-fiber/models"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable Timezone=Asia/Seoul",
		config.AppConfig.DBHost,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
		config.AppConfig.DBPort,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	log.Info("Database connected successfully")

	Migrate(DB)
}

func Migrate(DB *gorm.DB) {
	log.Info("Database migration started")

	DB.AutoMigrate(&models.User{})
	Session := DB.Session(&gorm.Session{PrepareStmt: false})

	if Session != nil {
		log.Info("Database migration completed")
	}
}
