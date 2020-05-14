package endpoint

import (
	"github.com/go-pg/pg"
)
//Config struct
type PostgreConfig struct {
	User     string
	Password string
	Port     string
	Host     string
}

//Connection to postgre (for task)
func NewPostgre(config PostgreConfig) (TaskTodo) {
	db := pg.Connect(&pg.Options{
		Addr:     config.Host + ":" + config.Port,
		User:     config.User,
		Password: config.Password,
		Database:"todo",
	})

	return &postgreStore{db: db}
}

//Connection to postgre (for user)
func NewPostgreUser(config PostgreConfig) (UserInt) {
	db := pg.Connect(&pg.Options{
		Addr:     config.Host + ":" + config.Port,
		User:     config.User,
		Password: config.Password,
		Database:"todo",
	})

	return &postgreStore{db: db}
}

//db func.
type postgreStore struct {
	db *pg.DB
}

func (p postgreStore) GetListUsers() ([]*User, error) {
	var users []*User
	err := p.db.Model(&users).Select()

	if err != nil {
		return nil, err
	}
	return users, err
}

//CreateUser ...
func (p postgreStore) CreateUser(user *User) (*User, error) {
	res := p.db.Insert(user)
	return user, res
}

//GetUser ...
func (p postgreStore) GetUser(email string) (*User, error) {
	user := &User{Email: email}
	err := p.db.Select(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//DeleteTask ...
func (p postgreStore) DeleteTask(id int) error {
	task := &Task{ID: id}
	err := p.db.Delete(task)
	if err != nil{
		return err
	}

	return nil
}

//GetListTask ...
func (p postgreStore) GetListTask() ([]*Task, error) {
	var tasks []*Task
	err := p.db.Model(&tasks).Select()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//UpdateTask ...
func (p postgreStore) UpdateTask(id int, task *Task) (*Task, error) {
	task.ID = id
	err := p.db.Update(task)
	return task, err
}

//GetTask ...
func (p postgreStore) GetTask(id int) (*Task, error) {
	task := &Task{ID: id}
	err := p.db.Select(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

//CreateTask ...
func (p postgreStore) CreateTask(task *Task) (*Task, error) {
	res := p.db.Insert(task)
	return task, res
}






