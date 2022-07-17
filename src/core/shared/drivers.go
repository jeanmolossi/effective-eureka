package shared

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// dbConnection is the database connection manager
type dbConnection struct {
	db *gorm.DB
}

// NewDBConnection creates a new database connection manager
func NewDbConnection() *dbConnection {
	return &dbConnection{}
}

// Connect connects to the database
func (d *dbConnection) Connect() error {
	db, err := gorm.Open(postgres.Open(dsn()), &gorm.Config{})
	if err != nil {
		return err
	}

	d.db = db

	return nil
}

// DB returns the database connection
func (d *dbConnection) DB() *gorm.DB {
	return d.db
}

// dsn returns the database connection string
func dsn() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password,
	)
}
