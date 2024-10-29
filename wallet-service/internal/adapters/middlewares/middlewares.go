package middlewares

import (
	"github.com/MokhtarSMokhtar/online-wallet/comman/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		// Assuming claims.UserId is a string, convert it to int32
		userID, err := strconv.ParseInt(claims.UserId, 10, 32)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
			return
		}
		c.Set("userID", int32(userID))
		c.Next()
	}
}
