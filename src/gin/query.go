package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r3 := gin.Default()

	r3.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"content": "hello world",
		})
	})

	// 获取Query参数
	r3.GET("/users", func(context *gin.Context) {
		name := context.Query("name")
		role := context.DefaultQuery("role", "teacher")
		context.String(http.StatusOK, "%s is a %s", name, role)
	})

	// 获取POST参数
	r3.POST("/form", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.DefaultPostForm("password", "000000")

		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	}) // Query和POST混合参数

	r3.POST("/posts", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		username := context.PostForm("username")
		password := context.DefaultPostForm("password", "张三")

		context.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	// Map参数(字典参数)
	r3.POST("/post", func(context *gin.Context) {
		ids := context.QueryMap("ids")
		names := context.PostFormMap("names")

		context.JSON(http.StatusOK, gin.H{
			"ids":  ids,
			"name": names,
		})
	})

	//重定向(Redirect)
	r3.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/")
	})

	r3.GET("/goindex", func(context *gin.Context) {
		context.Request.URL.Path = "/"
		r3.HandleContext(context)
	})

	// 分组路由(Grouping Routes)
	defaultHandler := func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"path": context.FullPath(),
		})
	}

	// group v1
	v1 := r3.Group("/v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}

	// group v2
	v2 := r3.Group("/v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}

	err := r3.Run(":9090")
	if err != nil {
		return
	}

}
