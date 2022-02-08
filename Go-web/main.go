package main

import (
	"log"
	"net/http"
	"time"
	"web"
)

func specialMiddleware() web.HandlerFunc {
	return func(c *web.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for SpecialMiddleware", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := web.Default()
	// 定义正常分组
	r.GET("/hello", func(c *web.Context) {
		c.String(http.StatusOK, "Hello zzguo!\n")
	})

	// 定义路由分组 /api/special 下的单独中间件 SpecialMiddleware
	// 此时 /api/special 下有三个中间件,两个全局Logger、recovery,一个specialMiddleware
	api1 := r.Group("/api/special")
	api1.Use(specialMiddleware())
	{
		api1.GET("/hello/:name", func(c *web.Context) {
			// expect /hello/zzguo
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	// 数组越界测试 错误恢复
	r.GET("/panic", func(c *web.Context) {
		names := []string{"zzguo so hansome!"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
