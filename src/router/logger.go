package router

import (
	"github.com/gin-gonic/gin"
	"time"
	"qiniupkg.com/x/log.v7"
)

func ScreenShotLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		log.Printf(
			"%s\t%d",
			latency,
			status,
		)
	}
}
