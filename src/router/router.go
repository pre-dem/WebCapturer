package router

import (
	"github.com/gin-gonic/gin"
	"base"
)

var config *base.Config

// NewRouter create a router engine
func RunNewRouter(c *base.Config) {
	config = c

	// Init router
	router := gin.New()
	router.Use(ScreenShotLogger())

	// Register routes to router
	for _, route := range routes {
		var handler gin.HandlerFunc

		handler = route.HandlerFunc

		router.Handle(route.Method, route.Pattern, handler)
	}

	router.Run(config.ListenAddress)
}
