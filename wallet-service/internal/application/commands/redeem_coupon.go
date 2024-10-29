// internal/application/commands/redeem_coupon.go

package commands

import (
	"database/sql"
	"errors"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/enums"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/models"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/repositories"
	"gorm.io/gorm"
	"time"
)

type CommandHandlers struct {
	TransactionRepo repositories.WalletRepository
	CouponRepo      repositories.CouponRepository
	DB              *gorm.DB
}

type RedeemCouponCommand struct {
	UserID int32
	Code   string
}

func (h *CommandHandlers) RedeemCoupon(cmd RedeemCouponCommand) error {
	// Check if coupon exists and is valid
	txOptions := &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	}
	return h.DB.Transaction(func(tx *gorm.DB) error {
		coupon, err := h.CouponRepo.GetCouponByCode(cmd.Code)
		if err != nil {
			return err
		}
		if time.Now().After(coupon.ExpiredAt) {
			return errors.New("coupon has expired")
		}
		// Check if user has already used the coupon
		used, err := h.CouponRepo.IsCouponUsedByUser(coupon.ID, cmd.UserID)
		if err != nil {
			return err
		}
		if used {
			return errors.New("coupon already used")
		}
		// update the user wallet
		userWall, err := h.TransactionRepo.GetWalletByUserID(cmd.UserID)
		if err != nil {
			return err
		}
		walletTran := models.WalletTransaction{
			Balance:           userWall.Balance + coupon.Amount,
			UserId:            cmd.UserID,
			TransactionReason: "Redeemed coupon",
			TransactionType:   enums.BalanceAddition,
			Credit:            coupon.Amount,
		}
		err = h.TransactionRepo.CreateWalletTransaction(&walletTran)
		if err != nil {
			return err
		}
		usage := &models.CouponUsage{
			CouponID: coupon.ID,
			WalletID: cmd.UserID,
			UsedAt:   time.Now(),
		}
		if err := h.CouponRepo.MarkCouponAsUsed(usage); err != nil {
			return err
		}
		return nil
	}, txOptions)
}
