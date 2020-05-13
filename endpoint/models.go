package endpoint

import "time"

// model for user
type User struct {
	Email	string `json:"email"`
	Password	string 	`json:"password"`
}
// model for task, that todo
type Task struct {
	ID 			int 		`json:"id"`
	Title		string		`json:"title"`
	Description	string 		`json:"description"`
	Deadline	time.Time	`json:"deadline"`
	IsDone		bool		`json:"is_done"`
}


type TaskTodo interface {
	CreateTask(task *Task) (*Task,error)
	GetTask(id int) (*Task,error)
	DeleteTask(id int) error
	UpdateTask(id int,task *Task) (*Task,error)
	GetListTask() ([]*Task,error)
}

type UserInt interface {
	CreateUser(user *User) (*User,error)
	GetUser(email string) (*User,error)
}
