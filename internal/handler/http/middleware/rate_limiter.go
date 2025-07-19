package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"kumparan-tech-test/internal/utils"
	"time"
)

var limiter = rate.NewLimiter(rate.Every(1*time.Second), 500)

func RateLimiter(c *gin.Context) {
	if !limiter.Allow() {
		utils.ResponseError(c, utils.ErrTooManyRequest("Rate limit exceeded", "middleware.RateLimiter"))
		return
	}
	c.Next()
}
