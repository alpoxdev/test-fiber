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
		config.AppConfig.DB_HOST,
		config.AppConfig.DB_USER,
		config.AppConfig.DB_PASSWORD,
		config.AppConfig.DB_NAME,
		config.AppConfig.DB_PORT,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	// Connection pool 설정
	sqlDB.SetMaxIdleConns(config.AppConfig.DB_MAX_CONNECTION) // 유휴 상태로 유지할 수 있는 최대 연결 수를 5로 설정합니다.
	sqlDB.SetMaxOpenConns(config.AppConfig.DB_MAX_CONNECTION) // 동시에 열 수 있는 최대 연결 수를 5로 설정합니다.
	sqlDB.SetConnMaxLifetime(0) // 연결의 최대 수명을 설정합니다. 0은 무제한 수명을 의미합니다.

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
