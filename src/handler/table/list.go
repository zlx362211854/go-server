package table

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"serverGo/src/db"
	"serverGo/src/declare"
	"serverGo/src/utils"
	"sort"
)

func ListHandle(c *gin.Context) {
	data, err := db.DB.ReadAll("table")
	if err != nil {
		fmt.Println("Error", err)
	}
	lists := []declare.Table{}
	for _, f := range data {
		tableFound := declare.Table{}
		if err := json.Unmarshal([]byte(f), &tableFound); err != nil {
			fmt.Println("Error", err)
		}
		lists = append(lists, tableFound)
	}
	sort.Sort(utils.TableList(lists))

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  0,
			"result": lists,
		},
	)
}