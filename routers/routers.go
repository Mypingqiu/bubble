package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//创建默认的路由引擎
	r := gin.Default()
	//告诉gin框架模版文件引用的静态文件
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模版文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	// v1 设置路由组
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加待办事项
		v1Group.POST("/todo", controller.CreateATodo)
		// 查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)

		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
