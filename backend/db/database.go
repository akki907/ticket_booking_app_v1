package db

import (
	"github.com/akki907/ticket_booking_app_v1/config"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig, DBMigrator func(*gorm.DB) error) *gorm.DB {
	// uri := fmt.Sprintf(`
	// host=%s user=%s dbname=%s password=%s sslmode=%s port=5432`,
	// 	config.DBHost, config.DBUser, config.DBName, config.DBPassword, config.DBSSLMode,
	// )

	uri := `postgresql://lingo_owner:JrcO4FDzf8YE@ep-snowy-bar-a5vkxsr2.us-east-2.aws.neon.tech/ticket_booking?sslmode=require`

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Info("Connected to the database")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Unable to migrate: %v", err)
	}

	return db

}
