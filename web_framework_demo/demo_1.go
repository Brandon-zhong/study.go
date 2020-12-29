package main

import (
	"fmt"
	"net/http"
	"strings"
	"study.go/web_framework_demo/jiangyu"
)

func main() {
	engine := jiangyu.New()
	engine.Get("/", func(ctx *jiangyu.Context) {
		ctx.String(http.StatusOK, "url.path --> %s\n", ctx.Req.URL.Path)
	})
	engine.Post("/hello", func(ctx *jiangyu.Context) {
		buf := new(strings.Builder)
		for k, v := range ctx.Req.Header {
			buf.WriteString(fmt.Sprintf("header [%s] --> %s\n", k, v))
		}
		ctx.String(http.StatusOK, buf.String())
	})
	engine.Run(":9999")
}
