package repositories

import (
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/models"
	"gorm.io/gorm"
)

type WalletRepository interface {
	GetWalletByUserID(userId int32) (*models.WalletTransaction, error)
	CreateWalletTransaction(wallet *models.WalletTransaction) error
	GetUserWalletTransactionHistory(userId int32, limit, offset int) ([]models.WalletTransaction, error)
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{db: db}
}

func (r *walletRepository) CreateWalletTransaction(wallet *models.WalletTransaction) error {
	if err := r.db.Create(wallet).Error; err != nil {
		return err
	}
	return nil
}

func (r *walletRepository) GetUserWalletTransactionHistory(userId int32, limit, offset int) ([]models.WalletTransaction, error) {
	var transactions []models.WalletTransaction
	err := r.db.Where("user_id = ?", userId).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&transactions).Error
	return transactions, err
}

func (r *walletRepository) GetWalletByUserID(userId int32) (*models.WalletTransaction, error) {
	var wallet models.WalletTransaction
	if err := r.db.Where("user_id = ?", userId).First(&wallet).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}
