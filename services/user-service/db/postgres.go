package db

import (
	"fmt"

	"github.com/prihodkoanton/go-services/services/user-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Name, cfg.Database.SSLMode, cfg.Database.Password,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
