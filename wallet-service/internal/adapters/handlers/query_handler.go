package handlers

import (
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/application/queries"
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type QueryHandler struct {
	QueryHandlers *queries.QueryHandlers
}

func NewQueryHandler(qh *queries.QueryHandlers) *QueryHandler {
	return &QueryHandler{QueryHandlers: qh}
}

type BalanceResponse struct {
	Balance float32 `json:"balance"`
}

// TransactionsResponse represents the response body for GetTransactions
type TransactionsResponse struct {
	Transactions []models.WalletTransaction `json:"transactions"`
}

// GetBalance godoc
// @Summary      Get wallet balance
// @Description  Retrieves the current balance of the user's wallet
// @Tags         Wallet
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {object}  handlers.BalanceResponse    "Balance amount"
// @Failure      401  {object}  handlers.ErrorResponse      "Unauthorized"
// @Router       /wallet/balance [get]
func (h *QueryHandler) GetBalance(c *gin.Context) {
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

	// Create and execute query
	query := queries.GetBalanceQuery{
		UserID: userID,
	}
	balance, err := h.QueryHandlers.GetBalance(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

// GetTransactions godoc
// @Summary      Get transaction history
// @Description  Retrieves the user's wallet transaction history
// @Tags         Wallet
// @Produce      json
// @Security     ApiKeyAuth
// @Param        limit   query     int     false  "Number of transactions to retrieve"
// @Param        offset  query     int     false  "Offset for pagination"
// @Success      200  {object}  handlers.TransactionsResponse  "List of transactions"
// @Failure      400  {object}  handlers.ErrorResponse         "Bad Request"
// @Failure      401  {object}  handlers.ErrorResponse         "Unauthorized"
// @Router       /wallet/transactions [get]
func (h *QueryHandler) GetTransactions(c *gin.Context) {
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

	// Get query parameters for pagination
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}

	// Create and execute query
	query := queries.GetTransactionsQuery{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	}
	transactions, err := h.QueryHandlers.GetTransactions(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}
