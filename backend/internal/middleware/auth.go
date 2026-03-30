package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/example/entrax/backend/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AzureADJWTMiddleware struct {
	cfg config.Config
}

func NewAzureADJWTMiddleware(cfg config.Config) AzureADJWTMiddleware {
	return AzureADJWTMiddleware{cfg: cfg}
}

func (m AzureADJWTMiddleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(errors.New("missing authorization header"))
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, _, err := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
		if err != nil {
			c.Error(errors.New("invalid bearer token format"))
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["aud"] != m.cfg.ExpectedAudience {
			c.Error(errors.New("token audience mismatch"))
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
