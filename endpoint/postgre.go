package endpoint

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

)
//Config struct
type PostgreConfig struct {
	User     string
	Password string
	Port     string
	Host     string
	Database string
}

//Connection to postgre (for task)
func NewPostgre(config PostgreConfig) (TaskTodo,error) {
	db := pg.Connect(&pg.Options{
		Addr:     config.Host + ":" + config.Port,
		User:     config.User,
		Password: config.Password,
		Database:config.Database,
	})
	err := db.CreateTable(&Task{},&orm.CreateTableOptions{
		IfNotExists:   true,
	})
	if err!= nil {
		return nil,err
	}
	err = db.CreateTable(&User{},&orm.CreateTableOptions{
		IfNotExists:   true,
	})
	if err!= nil {
		return nil,err
	}
	return &postgreStore{db: db},nil
}



//db func.
type postgreStore struct {
	db *pg.DB
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






