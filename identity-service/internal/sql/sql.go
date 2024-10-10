package sql

import (
	"database/sql"
	"fmt"
	"github.com/MokhtarSMokhtar/online-wallet/identity-service/config"
	_ "github.com/lib/pq"
)

type Db struct {
	Sql *sql.DB
}

func NewIdentity() *Db {
	return &Db{}
}

func (Db *Db) GetConnection() (*Db, error) {
	config := config.NewConfig()
	conString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)
	db, err := sql.Open("postgres", conString)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	Db.Sql = db
	return Db, nil
}
