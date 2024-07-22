package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/taufiksty/expense-tracker-app-backend/auth"
	"github.com/taufiksty/expense-tracker-app-backend/config"
	"github.com/taufiksty/expense-tracker-app-backend/models"
	"github.com/taufiksty/expense-tracker-app-backend/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.RespondError(c, http.StatusUnauthorized, "Authorization header is required")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Check if the token is blacklisted
		var blacklistToken models.BlacklistToken
		if err := config.DB.Where("token = ?", tokenString).First(&blacklistToken).Error; err == nil {
			utils.RespondError(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				utils.RespondError(c, http.StatusUnauthorized, "Invalid token signature")
			} else {
				utils.RespondError(c, http.StatusUnauthorized, "Invalid token")
			}
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
