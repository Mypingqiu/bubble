package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//创建默认的路由引擎
	r := gin.Default()
	r.GET("/bubble", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "测试记事本",
		})
	})
	//测试
	r.Run(":8080")

}
