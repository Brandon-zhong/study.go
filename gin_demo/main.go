package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	err := Routers().Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}

func Routers() *gin.Engine {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("get current path err")
		os.Exit(0)
	}
	dir += "/gin_demo"
	router := gin.Default()
	router.LoadHTMLGlob(dir + "/templates/*.html")
	router.LoadHTMLGlob(dir + "/templates/*/*.html")
	router.GET("/posts/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index/index", gin.H{
			"title": "posts" + strconv.Itoa(int(time.Now().Unix())),
		})
	})
	router.GET("/users/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "news/index.html", gin.H{
			"title": "users",
		})
	})

	return router
}
