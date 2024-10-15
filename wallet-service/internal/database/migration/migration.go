package migration

import (
	"fmt"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/database"
	"log"
)

func MigrateDb() {
	db := database.GetDB()
	for i, migrationFunc := range migrations {
		query := migrationFunc()
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Migration %d failed: %v", i+1, err)
		} else {
			fmt.Printf("Migration %d executed successfully\n", i+1)
		}
	}
}

var migrations = []func() string{
	// Migration to create the wallet table
	func() string {
		return `
            CREATE TABLE IF NOT EXISTS wallets (
                id SERIAL PRIMARY KEY,
                user_id VARCHAR(255) NOT NULL,
                debit DECIMAL(15, 2) DEFAULT 0.00,
                credit DECIMAL(15, 2) DEFAULT 0.00,
                balance DECIMAL(15, 2) DEFAULT 0.00,
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
            )
        `

	},

	func() string {
		return `
            CREATE TABLE IF NOT EXISTS wallet_trans (
                id SERIAL PRIMARY KEY,
                wallet_id INT NOT NULL,
                transaction_type VARCHAR(255) NOT NULL, 
                amount DECIMAL(15, 2) NOT NULL,
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                FOREIGN KEY (wallet_id) REFERENCES wallets(id) ON DELETE CASCADE
            )
        `
	},
}
