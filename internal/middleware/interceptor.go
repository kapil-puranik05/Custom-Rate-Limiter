package middleware

import (
	"net/http"
	"rate_limiter/internal/database"
	"time"

	"github.com/gin-gonic/gin"
)

func InterceptRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Requestor-id")
		if header == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Requestor id is required"})
			c.Abort()
			return
		}
		count, err := database.Client.Incr(database.Ctx, header).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
			c.Abort()
			return
		}
		if count == 1 {
			database.Client.Expire(database.Ctx, header, 1*time.Minute)
		}
		if count > 1 {
			ttl, _ := database.Client.TTL(database.Ctx, header).Result()
			if ttl == -1 {
				database.Client.Expire(database.Ctx, header, 1*time.Minute)
			}
		}
		if count > 10 {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded. Try again in a minute.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
