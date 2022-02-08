package main

import (
	"gee"
	"net/http"
)

/**
/panic：
Traceback:
        /usr/local/go/src/runtime/panic.go:1038
        /usr/local/go/src/runtime/panic.go:90
        /Users/kokorozashiguohuang/GolandProjects/Go-web/Practice/6-panic-recover/main.go:16
        /Users/kokorozashiguohuang/GolandProjects/Go-web/Practice/6-panic-recover/gee/Context.go:44
        /Users/kokorozashiguohuang/GolandProjects/Go-web/Practice/6-panic-recover/gee/recovery.go:37
        /Users/kokorozashiguohuang/GolandProjects/Go-web/Practice/6-panic-recover/gee/Context.go:44
        /Users/kokorozashiguohuang/GolandProjects/Go-web/Practice/6-panic-recover/gee/logger.go:15
        /Users/kokorozashiguohuang/GolandProjects/Go-web/Practice/6-panic-recover/gee/Context.go:44
        /Users/kokorozashiguohuang/GolandProjects/Go-web/Practice/6-panic-recover/gee/router.go:102
        /Users/kokorozashiguohuang/GolandProjects/Go-web/Practice/6-panic-recover/gee/gee.go:107
        /usr/local/go/src/net/http/server.go:2879
        /usr/local/go/src/net/http/server.go:1930
        /usr/local/go/src/runtime/asm_amd64.s:1582

2022/02/08 16:13:07 [500] /panic in 926.021µs

*/
func main() {
	r := gee.Default()
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello zzguo!\n")
	})
	// 数组越界测试 错误恢复功能
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"zzguo so hansome"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9119")
}
