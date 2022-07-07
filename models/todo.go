package models

import "bubble/dao"

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
   Todo 增删改查
*/

//创建Todo
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	if err != nil {
		return err
	}
	return
}

//查询所有Todo
func GetTodoList() (err error, todolist []*Todo) {
	err = dao.DB.Find(&todolist).Error
	if err != nil {
		return err, nil
	}
	return
}

//查看某个Todo
func GetATodo(id string) (err error, todo *Todo) {
	todo = new(Todo)
	err = dao.DB.Where("id=?", id).First(todo).Error
	if err != nil {
		return err, nil
	}
	return
}

//修改某个Todo
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Debug().Save(&todo).Error
	if err != nil {
		return err
	}
	return
}

//删除某个Todo
func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	if err != nil {
		return err
	}
	return
}
