package task

import (
	"testing"
	"time"
)

func TestEndpointsFactory_GetTask(t *testing.T) {
	var err error
	user := task2.PostgreConfig{
		User:     "postgres",
		Password: "daha2000",
		Port:     "5432",
		Host:     "127.0.0.1",
	}
	db := task2.NewPostgre(user)
	endpoints := task2.NewEndpointsFactory(db)

	task , err := endpoints.taskTodo.GetTask(1)
	if err != nil {
		t.Error(err)
	}
	//По умолчанию тайтл как SSSS (пример)
	if task.Title != "SSSS" {
		t.Error("Title is not equal")
	}
}

func TestEndpointsFactory_CreateTask(t *testing.T) {
	user := task2.PostgreConfig{
		User:     "postgres",
		Password: "daha2000",
		Port:     "5432",
		Host:     "127.0.0.1",
	}
	db := task2.NewPostgre(user)
	endpoints := task2.NewEndpointsFactory(db)
	task := &task2.Task{
		ID:          3,
		Title:       "sdasd",
		Description: "DSde",
		Deadline:    time.Time{},
		IsDone:      true,
	}
	_, err := endpoints.taskTodo.CreateTask(task)
	if err!=nil{
		t.Error(err)
	}
	if task.IsDone==true{
		t.Error("Task can't be executed")
	}


}
func TestEndpointsFactory_DeleteTask(t *testing.T) {

}

func TestEndpointsFactory_UpdateTask(t *testing.T) {

}
func TestPostgreStore_GetListTask(t *testing.T) {

}
func TestEndpointsFactory_ExecuteTask(t *testing.T) {

}