package form

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"serverGo/src/db"
)
type Delete struct {
	id string `json:"id" binding:"required"`
}
func DeleteHandle(c *gin.Context) {
	var delete Delete
	if err := c.ShouldBindWith(&delete, binding.Query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Query("id")
	if err := db.DB.Delete("form", id); err != nil {
		fmt.Println("Error", err)
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  0,
			"result": "create success!",
		},
	)
}