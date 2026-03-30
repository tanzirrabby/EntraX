package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ErrorHandler(log zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		statusCode := c.Writer.Status()
		if statusCode < 400 {
			statusCode = http.StatusInternalServerError
		}

		log.Error().Err(err).Int("status", statusCode).Str("path", c.Request.URL.Path).Msg("request failed")
		c.JSON(statusCode, gin.H{"error": err.Error()})
	}
}
