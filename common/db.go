package common

import (
	"fmt"
	"os"
	"time"

	"github.com/mustafawidiarto/go-boilerplate/model/entity"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDatabase returns a new Gorm DB instance that is initialized with PostgreSQL
// database connection using environment variables for configuration.
func NewDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal().Err(err).Msg("unable to open db connection")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get sql db")
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(15)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = db.AutoMigrate(
		&entity.Room{},
	)

	if err != nil {
		log.Fatal().Err(err).Msg("database migration failed")
	}

	return db
}
