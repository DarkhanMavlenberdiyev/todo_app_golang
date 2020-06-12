package user

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
