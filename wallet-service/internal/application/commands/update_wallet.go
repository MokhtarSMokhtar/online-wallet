package commands

import (
	"errors"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/enums"
	"time"

	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/models"
	"gorm.io/gorm"
)

type UpdateWalletCommand struct {
	UserID int32
	Amount float32
	Reason string
}

func (h *CommandHandlers) UpdateWallet(cmd UpdateWalletCommand) error {
	return h.DB.Transaction(func(tx *gorm.DB) error {
		// Get the user's latest wallet balance
		walletTransaction, err := h.TransactionRepo.GetWalletByUserID(cmd.UserID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		var newBalance float32
		if walletTransaction != nil {
			newBalance = walletTransaction.Balance + cmd.Amount
		} else {
			newBalance = cmd.Amount
		}

		// Create a new wallet transaction
		newTransaction := models.WalletTransaction{
			Balance:           newBalance,
			UserId:            cmd.UserID,
			TransactionReason: cmd.Reason,
			TransactionType:   determineTransactionType(cmd.Amount),
			Credit:            cmd.Amount,
			CreatedAt:         time.Now(),
		}

		err = h.TransactionRepo.CreateWalletTransaction(&newTransaction)
		if err != nil {
			return err
		}

		return nil
	})
}

func determineTransactionType(amount float32) enums.TransactionType {
	if amount >= 0 {
		return enums.BalanceAddition
	}
	return enums.SendToUser
}
