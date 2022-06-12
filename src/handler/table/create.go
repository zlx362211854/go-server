package table

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
	var table declare.Table
	table.CreateTime = time.Now().UnixNano() / 1e6
	table.Id = utils.UniqueId()
	if err := c.ShouldBindWith(&table, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Write("table", table.Id, table); err != nil {
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