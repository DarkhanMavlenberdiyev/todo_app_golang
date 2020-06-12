package endpoint

import "time"


// model for task, that todo
type Task struct {
	ID 			int 		`json:"id"`
	Title		string		`json:"title"`
	Description	string 		`json:"description"`
	Deadline	time.Time	`json:"deadline"`
	IsDone		bool		`json:"is_done"`
}


//Interfaces for task
type TaskTodo interface {
	CreateTask(task *Task) (*Task,error)
	GetTask(id int) (*Task,error)
	DeleteTask(id int) error
	UpdateTask(id int,task *Task) (*Task,error)
	GetListTask() ([]*Task,error)
}

