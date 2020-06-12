package task

import "time"


// model for task, that todo
type Task struct {
	ID 			int 		`json:"id"`
	Title		string		`json:"title"`
	Description	string 		`json:"description"`
	Deadline	time.Time	`json:"deadline"`
	IsDone		bool		`json:"is_done"`
}

// model for user
type User struct {
	ID 				int 		`json:"id"`
	FirstName 		string 		`json:"first_name"`
	LastName		string 		`json:"last_name"`
	Email			string 		`json:"email"`
	Password		string 		`json:"password"`
}


//Interfaces for task
type TaskTodo interface {
	CreateTask(task *Task) (*Task,error)
	GetTask(id int) (*Task,error)
	DeleteTask(id int) error
	UpdateTask(id int,task *Task) (*Task,error)
	GetListTask() ([]*Task,error)
}

// interfaces for user
type UserInfo interface {
	CreateUser(user *User) (*User,error)
	GetUser(id int) (*User,error)
	UpdateUser(id int, user *User) (*User,error)
	DeleteUser(id int) error
}