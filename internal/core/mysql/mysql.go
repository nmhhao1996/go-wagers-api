package mysql

import (
	"github.com/nmhhao1996/go-wagers-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect connects to the MySQL database
func Connect(cfg config.MySQLConfig) (*gorm.DB, error) {
	logMode := logger.Silent
	if cfg.Debug {
		logMode = logger.Error
	}

	db, err := gorm.Open(mysql.Open(cfg.URI), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
