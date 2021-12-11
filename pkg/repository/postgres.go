package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const driverName = "postgres"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// NewConfig returns pointer on new instance of Config
func NewConfig(host string, port string, username string, password string, dbName string, sslMode string) *Config {
	result := &Config{
		Host:     host,
		Port:     port,
		Username: username,
		DBName:   dbName,
		SSLMode:  sslMode,
		Password: password,
	}

	return result
}

// NewPostgresDB returns pointer on new instance of postgres db and error
func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}
