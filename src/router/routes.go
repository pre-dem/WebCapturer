package router

import "github.com/gin-gonic/gin"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetScreenShot",
		"GET",
		"/v1/get_screenshot",
		GetScreenShot_v1,
	},
}
