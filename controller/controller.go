package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
url ---> controller   ---> logic  ---> model
请求来了  --》 控制器--》业务逻辑--》模型层的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func CetTodoList(c *gin.Context) {
	//todo := &Todo{}
	//todoList := DB.Find(&todo)
	//c.JSON(200, gin.H{
	//	"message": todoList,
	//})
}

func UpdateATodo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "put todo:id",
	})
}

func DeleteATodo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete todo:id",
	})
}
