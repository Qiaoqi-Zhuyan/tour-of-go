package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r2 := gin.Default()
	r2.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "hello %s", name)
	})

	r2.Run(":9090")
}
