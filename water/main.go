package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-water/water"
	"github.com/go-water/water/multitemplate"
)

func main() {
	router := water.New()
	router.HTMLRender = createMyRender()

	router.Use(Logger)
	router.GET("/", Index)
	v2 := router.Group("/v2")
	{
		v2.GET("/hello", GetHello)
	}

	router.Serve(":80")
}

func Index(ctx *water.Context) {
	ctx.HTML(http.StatusOK, "index", water.H{"title": "我是标题", "body": "你好，朋友。"})
}

func GetHello(ctx *water.Context) {
	ctx.JSON(http.StatusOK, water.H{"msg": "Hello World!"})
}

func Logger(handlerFunc water.HandlerFunc) water.HandlerFunc {
	return func(ctx *water.Context) {
		start := time.Now()
		defer func() {
			msg := fmt.Sprintf("[WATER] %v | %15s | %13v | %-7s %s",
				time.Now().Format("2006/01/02 - 15:04:05"),
				ctx.ClientIP(),
				time.Since(start),
				ctx.Request.Method,
				ctx.Request.URL.Path,
			)

			fmt.Println(msg)
		}()

		handlerFunc(ctx)
	}
}

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "views/layout.html", "views/index.html", "views/_header.html", "views/_footer.html")
	return r
}
