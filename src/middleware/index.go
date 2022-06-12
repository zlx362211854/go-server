package middleware

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)


func RegisterMiddleWare() {

	gin.DisableConsoleColor()
	f, _ := os.Create("../log/server.log")
	gin.DefaultWriter = io.MultiWriter(f)

}
