package main

import (
	"log"
	"net/http"
	"time"

	"gee"
)

func specialMiddleware() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for SpecialMiddleware", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	// 全局中间件
	r.Use(gee.Logger())
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	// 定义路由分组 /api1 下的单独中间件
	// 此时 /api1 下有两个中间件, 一个Logger, 一个specialMiddleware
	api1 := r.Group("/api/special")
	api1.Use(specialMiddleware())
	{
		api1.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/zzguo
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9009")
}
