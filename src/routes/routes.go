package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Run will start the server
func Run() {
	//middleware.RegisterMiddleWare()
	getRoutes()
	router.Run(":3060")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	v1 := router.Group("/v1")
	addFormRoutes(v1)
	addTableRoutes(v1)
}