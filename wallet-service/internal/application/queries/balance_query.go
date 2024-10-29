package queries

import (
	"errors"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/repositories"
	"gorm.io/gorm"
)

type QueryHandlers struct {
	TransactionRepo repositories.WalletRepository
}
type GetBalanceQuery struct {
	UserID int32
}

func (h *QueryHandlers) GetBalance(query GetBalanceQuery) (float32, error) {
	latestTransaction, err := h.TransactionRepo.GetWalletByUserID(query.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}
	return latestTransaction.Balance, nil
}
