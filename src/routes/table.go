package routes

import (
	"github.com/gin-gonic/gin"
	"serverGo/src/handler/table"
)

func addTableRoutes(versionGroup *gin.RouterGroup) {
	users := versionGroup.Group("/table")

	users.POST("/create", table.CreateHandle)
	users.GET("/list", table.ListHandle)
	users.GET("/delete", table.DeleteHandle)
}
