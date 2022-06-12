package form

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
	data, err := db.DB.ReadAll("form")
	if err != nil {
		fmt.Println("Error", err)
	}
	lists := []declare.Form{}
	for _, f := range data {
		formFound := declare.Form{}
		if err := json.Unmarshal([]byte(f), &formFound); err != nil {
			fmt.Println("Error", err)
		}
		lists = append(lists, formFound)
	}
	sort.Sort(utils.FormList(lists))

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  0,
			"result": lists,
		},
	)
}