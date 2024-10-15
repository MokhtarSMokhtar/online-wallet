package repository

import (
	"database/sql"
	"fmt"
	"log"
)

type WalletRepository struct {
	Conn *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{Conn: db}
}
func (r *WalletRepository) AddUserWallet(userId int) error {
	// Check if wallet already exists to ensure idempotency
	var existingWalletID int
	err := r.Conn.QueryRow(`
        SELECT id FROM wallets WHERE user_id = $1
    `, userId).Scan(&existingWalletID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to query existing wallet: %w", err)
	}
	if existingWalletID != 0 {
		log.Printf("Wallet already exists for UserID: %d", userId)
		return nil // Wallet already exists, no action needed
	}

	// Proceed to create wallet
	query := `
        INSERT INTO wallets (user_id, debit, credit, balance, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
    `
	_, err = r.Conn.Exec(query, userId, 0.00, 0.00, 0.00)
	if err != nil {
		return fmt.Errorf("failed to create wallet: %w", err)
	}
	return nil
}
