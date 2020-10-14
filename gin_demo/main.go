package main

import (
	"github.com/gin-gonic/gin"
	"study.go/gin_demo/middleware"
)

func main() {

}


func Routers() *gin.Engine {

	engine := gin.Default()

	//跨域处理
	engine.Use(middleware.Cors)

	group := engine.Group("")

	group.Group()


	return nil
}




