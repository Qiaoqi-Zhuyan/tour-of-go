package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// gin.Default()生成了一个实例，这个实例即 WSGI 应用程序
	r := gin.Default()
	// 声明了一个路由，告诉 Gin 什么样的URL 能触发传入的函数
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world")
	})
	//
	r.Run(":9090") // listen and serve on 0.0.0.0:8080
}
