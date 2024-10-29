// internal/domain/repositories/coupon_repository.go

package repositories

import (
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/models"
	"gorm.io/gorm"
)

type CouponRepository interface {
	GetCouponByCode(code string) (*models.Coupon, error)
	IsCouponUsedByUser(couponID, userID int32) (bool, error)
	MarkCouponAsUsed(usage *models.CouponUsage) error
}

type couponRepository struct {
	db *gorm.DB
}

func NewCouponRepository(db *gorm.DB) CouponRepository {
	return &couponRepository{db: db}
}

func (r *couponRepository) GetCouponByCode(code string) (*models.Coupon, error) {
	var coupon models.Coupon
	err := r.db.Where("code = ?", code).First(&coupon).Error
	return &coupon, err
}

func (r *couponRepository) IsCouponUsedByUser(couponID, userID int32) (bool, error) {
	var count int64
	err := r.db.Model(&models.CouponUsage{}).
		Where("coupon_id = ? AND wallet_id IN (SELECT id FROM wallets WHERE user_id = ?)", couponID, userID).
		Count(&count).Error
	return count > 0, err
}

func (r *couponRepository) MarkCouponAsUsed(usage *models.CouponUsage) error {
	return r.db.Create(usage).Error
}
