package routes

import (
	"serverGo/src/handler/form"

	"github.com/gin-gonic/gin"
)

func addFormRoutes(versionGroup *gin.RouterGroup) {
	users := versionGroup.Group("/form")

	users.POST("/create", form.CreateHandle)
	users.GET("/list", form.ListHandle)
	users.GET("/delete", form.DeleteHandle)
}
