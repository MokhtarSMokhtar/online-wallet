package queries

import "github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/models"

type GetTransactionsQuery struct {
	UserID int32
	Limit  int
	Offset int
}

func (h *QueryHandlers) GetTransactions(query GetTransactionsQuery) ([]models.WalletTransaction, error) {
	transactions, err := h.TransactionRepo.GetUserWalletTransactionHistory(query.UserID, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
