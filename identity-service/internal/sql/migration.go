package sql

import (
	"database/sql"
	"fmt"
)

func MigrateDatabase() error {
	db := NewIdentity()
	conn, erro := db.GetConnection()
	if erro != nil {
		return fmt.Errorf("failed to connect to database: %w", erro)

	}
	defer func(sql *sql.DB) {
		err := sql.Close()
		if err != nil {

		}
	}(conn.Sql)
	query := `            CREATE TABLE IF NOT EXISTS users (
                id SERIAL PRIMARY KEY,
                full_name VARCHAR(100) NOT NULL,
                email VARCHAR(100) UNIQUE NOT NULL,
                country VARCHAR(50),
                country_code VARCHAR(10),
                phone VARCHAR(20),
                password_hash BYTEA NOT NULL,
                user_type VARCHAR(20) NOT NULL,
                gender VARCHAR(10),
                created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
            );`
	_, err := conn.Sql.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}
	return nil
}
