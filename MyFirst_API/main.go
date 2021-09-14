package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	//定義新路由並回傳數據到HTML
	router.GET("/data", func(c *gin.Context) {
		c.HTML(200, "data.html", gin.H{"data": "Hello Go's Gin World."})
	})

	//定義新路由並回傳數據到HTML
	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "ok",
			"data":      "Hello Go's Gin World.",
			"developer": "DarkCorNer",
		})
	})

	//定義新路由並讓前端回傳值到後端
	router.GET("/form", func(c *gin.Context) {
		c.HTML(200, "form.html", gin.H{})
	})

	//接收前端post數據，返回Json數據
	router.POST("/service", func(c *gin.Context) {
		uname := c.PostForm("uname")
		c.JSON(200, gin.H{
			"result": "ok",
			"Hello":  uname,
		})
	})

	//定義新路由前端post值，後端Get後回傳當前頁面
	router.GET("/user", HelloWorldGet)
	router.POST("/user", HelloWorldPost)

	router.Run(":8080")
}

func HelloWorldGet(c *gin.Context) {
	c.HTML(http.StatusOK, "user.html", gin.H{
		"username": "",
	})
}

func HelloWorldPost(c *gin.Context) {
	username := c.PostForm("UserName")
	c.HTML(http.StatusOK, "user.html", gin.H{
		"username": username,
	})
}
