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
func NewPostgre(config PostgreConfig) (UserInfo,error) {
	db := pg.Connect(&pg.Options{
		Addr:     config.Host + ":" + config.Port,
		User:     config.User,
		Password: config.Password,
		Database:config.Database,
	})

	err := db.CreateTable(&User{},&orm.CreateTableOptions{
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

//GetUser...
func(p postgreStore) GetUser(id int) (*User,error){
	user := &User{ID: id}
	err := p.db.Select(user)
	if err != nil{
		return nil,err
	}
	return user,nil
}
// CreateUser...

func(p postgreStore) CreateUser(user *User) (*User,error) {
	res := p.db.Insert(user)
	return user,res
}
// UpdateUser...
func(p postgreStore) UpdateUser(id int,user *User) (*User,error){
	user.ID = id
	err := p.db.Update(user)
	return user,err
}

func(p postgreStore) DeleteUser(id int) error {
	user := User{ID: id}
	err := p.db.Delete(user)
	if err != nil {
		return err
	}
	return nil
}
