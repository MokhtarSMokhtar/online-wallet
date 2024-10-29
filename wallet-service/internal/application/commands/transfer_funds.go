package commands

import (
	"errors"
	"fmt"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/enums"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/models"
	"gorm.io/gorm"
	"time"
)

type TransferFundsCommand struct {
	FromUserID int32
	ToUserID   int32
	Amount     float32
}

func (h *CommandHandlers) TransferFunds(cmd TransferFundsCommand) error {
	if cmd.FromUserID == cmd.ToUserID {
		return errors.New("cannot transfer to the same user")
	}
	if cmd.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	return h.DB.Transaction(func(tx *gorm.DB) error {
		//Get sender's latest balance
		senderTransaction, err := h.TransactionRepo.GetWalletByUserID(cmd.FromUserID)
		var senderBalance float32
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				senderBalance = 0
			} else {
				return err
			}
		} else {
			senderBalance = senderTransaction.Balance
		}

		// Step 2: Check if sender has enough balance
		if senderBalance < cmd.Amount {
			return errors.New("insufficient funds")
		}

		//Create debit transaction for sender
		senderNewBalance := senderBalance - cmd.Amount
		debitTransaction := models.WalletTransaction{
			Balance:           senderNewBalance,
			UserId:            cmd.FromUserID,
			TransactionReason: "Transfer to user " + fmt.Sprint(cmd.ToUserID),
			TransactionType:   enums.SendToUser,
			Debit:             cmd.Amount,
			CreatedAt:         time.Now(),
		}
		err = h.TransactionRepo.CreateWalletTransaction(&debitTransaction)
		if err != nil {
			return err
		}

		// Get receiver's latest balance
		receiverTransaction, err := h.TransactionRepo.GetWalletByUserID(cmd.ToUserID)
		var receiverBalance float32
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				receiverBalance = 0
			} else {
				return err
			}
		} else {
			receiverBalance = receiverTransaction.Balance
		}

		//  Create credit transaction for receiver
		receiverNewBalance := receiverBalance + cmd.Amount
		creditTransaction := models.WalletTransaction{
			Balance:           receiverNewBalance,
			UserId:            cmd.ToUserID,
			TransactionReason: "Transfer from user " + fmt.Sprint(cmd.FromUserID),
			TransactionType:   enums.BalanceAddition,
			Credit:            cmd.Amount,
			CreatedAt:         time.Now(),
		}
		err = h.TransactionRepo.CreateWalletTransaction(&creditTransaction)
		if err != nil {
			return err
		}

		return nil
	})
}
