package crud

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/sudoblockio/icon-transformer/config"
)

var postgresSession *gorm.DB
var postgresSessionOnce sync.Once

func formatPostgresDSN(host string, port string, user string, password string, dbname string, sslmode string, timezone string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)
}

func getPostgresConn() *gorm.DB {
	postgresSessionOnce.Do(func() {
		dsn := formatPostgresDSN(
			config.Config.DbHost,
			config.Config.DbPort,
			config.Config.DbUser,
			config.Config.DbPassword,
			config.Config.DbName,
			config.Config.DbSslmode,
			config.Config.DbTimezone,
		)

		var err error
		postgresSession, err = retryGetPostgresSession(dsn)
		if err != nil {
			zap.S().Fatal("Cannot create a connection to postgres", err)
		}

		zap.S().Info("Successful connection to postgres")
	})

	return postgresSession
}

func retryGetPostgresSession(dsn string) (*gorm.DB, error) {
	var session *gorm.DB
	operation := func() error {
		sess, err := createSession(dsn)
		if err != nil {
			zap.S().Info("POSTGRES SESSION Error : ", err.Error())
		} else {
			session = sess
		}
		return err
	}
	neb := backoff.NewExponentialBackOff()
	err := backoff.Retry(operation, neb)
	return session, err
}

func getLogLevel() logger.LogLevel {
	switch config.Config.LogLevel {
	case "warn":
		return logger.Warn
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "info":
		return logger.Info
	default:
		return logger.Warn
	}
}

func createSession(dsn string) (*gorm.DB, error) {

	slowThreshold := (time.Duration(config.Config.GormLoggingThresholdMilli) * time.Millisecond)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             slowThreshold, // Slow SQL threshold
			LogLevel:                  getLogLevel(), // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		zap.S().Info("err:", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.Config.DbMaxIdleConnections)
	sqlDB.SetMaxOpenConns(config.Config.DbMaxOpenConnections)

	return db, err
}
