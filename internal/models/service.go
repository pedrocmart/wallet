package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// NewDBService handles the database connection
func NewDBService() (*DBService, error) {
	db, err := gorm.Open("mysql", getConnectionString())
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &DBService{
		Wallet: &walletGorm{db},
		db:     db,
	}, nil
}

// DBService represents the database connection service
type DBService struct {
	Wallet *walletGorm
	db     *gorm.DB
}

// Close closes the database connection
func (s *DBService) Close() error {
	return s.db.Close()
}

// AutoMigrate will attempt to automatically migrate all tables
func (s *DBService) AutoMigrate() error {
	return s.db.AutoMigrate(&Wallet{}).Error
}

func getConnectionString() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "127.0.0.23"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3307"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}

	password := os.Getenv("DB_PASS")
	if password == "" {
		password = "root"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "wallet"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)
}
