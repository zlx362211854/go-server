package form

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"serverGo/src/db"
	"serverGo/src/declare"
	"serverGo/src/utils"
	"time"
)


func CreateHandle(c *gin.Context) {
	var form declare.Form
	form.CreateTime = time.Now().UnixNano() / 1e6
	form.Id = utils.UniqueId()
	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Write("form", form.Id, form); err != nil {
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