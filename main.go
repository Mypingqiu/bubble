package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建默认的路由引擎
	r := gin.Default()
	//告诉gin框架模版文件引用的静态文件
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模版文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//测试
	r.Run(":8080")

}
