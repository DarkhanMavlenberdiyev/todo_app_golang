package user

// model for user
type User struct {
	ID 				int 		`json:"id"`
	FirstName 		string 		`json:"first_name"`
	LastName		string 		`json:"last_name"`
	Email			string 		`json:"email"`
	Password		string 		`json:"password"`
}

// interfaces for user
type UserInfo interface {
	CreateUser(user *User) (*User,error)
	GetUser(id int) (*User,error)
	UpdateUser(id int, user *User) (*User,error)
	DeleteUser(id int) error
}
