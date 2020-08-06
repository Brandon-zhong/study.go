package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/get", func(context *gin.Context) {
		fmt.Println("")
		context.JSON(http.StatusOK, "jfaklsjflskjdf")
	})
	_ = engine.Run()
}
