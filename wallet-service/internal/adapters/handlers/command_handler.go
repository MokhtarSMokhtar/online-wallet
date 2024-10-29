// internal/adapters/handlers/command_handler.go

package handlers

import (
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/application/commands"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CommandHandler handles command-related HTTP requests
type CommandHandler struct {
	CommandHandlers *commands.CommandHandlers
}

// RedeemCouponRequest represents the request body for redeeming a coupon
type RedeemCouponRequest struct {
	Code string `json:"code" binding:"required"`
}

// TransferFundsRequest represents the request body for transferring funds
type TransferFundsRequest struct {
	ToUserID int32   `json:"to_user_id" binding:"required"`
	Amount   float32 `json:"amount" binding:"required,gt=0"`
}

// MessageResponse represents a generic success message response
type MessageResponse struct {
	Message string `json:"message"`
}

// ErrorResponse represents a generic error message response
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewCommandHandler creates a new CommandHandler
func NewCommandHandler(ch *commands.CommandHandlers) *CommandHandler {
	return &CommandHandler{CommandHandlers: ch}
}

// RedeemCoupon godoc
// @Summary      Redeem a coupon
// @Description  Allows a user to redeem a coupon code
// @Tags         Wallet
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        request body handlers.RedeemCouponRequest true "Coupon Code"
// @Success      200  {object}  handlers.MessageResponse  "Coupon redeemed successfully"
// @Failure      400  {object}  handlers.ErrorResponse    "Bad Request"
// @Failure      401  {object}  handlers.ErrorResponse    "Unauthorized"
// @Router       /wallet/redeem-coupon [post]
func (h *CommandHandler) RedeemCoupon(c *gin.Context) {
	// Get user ID from context
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userID, ok := userIDValue.(int32)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	// Bind JSON input to named struct
	var req RedeemCouponRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create and execute command
	cmd := commands.RedeemCouponCommand{
		UserID: userID,
		Code:   req.Code,
	}
	if err := h.CommandHandlers.RedeemCoupon(cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Coupon redeemed successfully"})
}

// TransferFunds godoc
// @Summary      Transfer funds to another user
// @Description  Allows a user to transfer funds to another user
// @Tags         Wallet
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        request body handlers.TransferFundsRequest true "Transfer Details"
// @Success      200  {object}  handlers.MessageResponse  "Transfer successful"
// @Failure      400  {object}  handlers.ErrorResponse    "Bad Request"
// @Failure      401  {object}  handlers.ErrorResponse    "Unauthorized"
// @Router       /wallet/transfer [post]
func (h *CommandHandler) TransferFunds(c *gin.Context) {
	// Get user ID from context
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	fromUserID, ok := userIDValue.(int32)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	// Bind JSON input to named struct
	var req TransferFundsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create and execute command
	cmd := commands.TransferFundsCommand{
		FromUserID: fromUserID,
		ToUserID:   req.ToUserID,
		Amount:     req.Amount,
	}
	if err := h.CommandHandlers.TransferFunds(cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
}
