package endpoint

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//endpoint for tasks
func NewEndpointsFactory(tasktodo TaskTodo) *endpointsFactory {
	return &endpointsFactory{taskTodo: tasktodo}
}

//endpoint for users
func NewEndpointsFactoryUser(userInt UserInt) *endpointsFactory{
	return &endpointsFactory{UserInt: userInt}
}

type endpointsFactory struct {
	taskTodo TaskTodo
	UserInt	 UserInt
}

func (ef *endpointsFactory) CreateUser() func (w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		user := &User{}
		if err := json.Unmarshal(data, user); err != nil {
			writeResponse(w,http.StatusBadRequest,[]byte("Error: "+err.Error()))
			return
		}

		er := ValidateEmail(user)
		if er!=nil {
			writeResponse(w,http.StatusUnauthorized,[]byte("Error: Incorrect email or password"))
			return
		}

		result, err := ef.UserInt.CreateUser(user)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		response, err := json.Marshal(result)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		writeResponse(w,http.StatusCreated,response)
	}
}

// Check email and passwords valid
func ValidateEmail(user *User) error{
	return validation.ValidateStruct(user,validation.Field(&user.Email,validation.Required,is.Email),
		validation.Field(&user.Password,validation.Required,validation.Length(8,100)))
}


func (ef *endpointsFactory) SessionUser() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		reqUser := &User{}
		if err := json.Unmarshal(data, reqUser); err != nil {
			writeResponse(w,http.StatusBadRequest,[]byte("Error: "+err.Error()))
			return
		}
		user,err := ef.UserInt.GetUser(reqUser.Email)
		if err!=nil || CompareTwoPasswords(reqUser.Password,user.Password){
			writeResponse(w,http.StatusUnauthorized,[]byte("Error: incorrect email or password"))
			return
		}

		writeResponse(w,http.StatusOK,[]byte("OK"))


	}
}

//compare two passwords
func CompareTwoPasswords(p1 string,p2 string) bool{
	return bcrypt.CompareHashAndPassword([]byte(p1),[]byte(p2))==nil
}


func (ef *endpointsFactory) GetTask(idParam string) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars[idParam]
		if !ok {
			writeResponse(w,http.StatusBadRequest,[]byte("Task not found"))
			return
		}
		idd, _ := strconv.Atoi(id)
		res, err := ef.taskTodo.GetTask(idd)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		data, err := json.Marshal(res)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		writeResponse(w,http.StatusOK,data)
	}
}

func (ef *endpointsFactory) CreateTask() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		task := &Task{}
		if err := json.Unmarshal(data, task); err != nil {
			writeResponse(w,http.StatusBadRequest,[]byte("Error: "+err.Error()))
			return
		}
		result, err := ef.taskTodo.CreateTask(task)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		response, err := json.Marshal(result)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		writeResponse(w,http.StatusCreated,response)
	}
}
func (ef *endpointsFactory) UpdateTask(idParam string) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars[idParam]
		if !ok {
			writeResponse(w,http.StatusBadRequest,[]byte("Task not found"))
			return
		}
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		task := &Task{}
		if err := json.Unmarshal(data, task); err != nil {
			writeResponse(w,http.StatusBadRequest,[]byte("Error: "+err.Error()))
			return
		}
		idd, _ := strconv.Atoi(id)
		res, err := ef.taskTodo.UpdateTask(idd, task)
		if err != nil {
			writeResponse(w,http.StatusBadRequest,[]byte("Error: "+err.Error()))
			return
		}
		response, err := json.Marshal(res)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		writeResponse(w,http.StatusCreated,response)

	}
}
func (ef *endpointsFactory) DeleteTask(idParam string) func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars[idParam]
		if !ok {
			writeResponse(w,http.StatusBadRequest,[]byte("Error: Not Found"))
			return
		}
		idd, _ := strconv.Atoi(id)
		err := ef.taskTodo.DeleteTask(idd)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		writeResponse(w,http.StatusOK,[]byte("Task is deleted successfully"))


	}
}
func (ef *endpointsFactory) GetListTask() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		list,err := ef.taskTodo.GetListTask()
		if err!=nil{
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		data, err := json.Marshal(list)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		writeResponse(w,http.StatusOK,data)



	}
}
func (ef *endpointsFactory) ExecuteTask(idParam string) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars[idParam]
		if !ok {
			writeResponse(w,http.StatusBadRequest,[]byte("Task not found"))
			return
		}
		idd, _ := strconv.Atoi(id)
		res, err := ef.taskTodo.GetTask(idd)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		executedTask := &Task{
			ID:          res.ID,
			Title:       res.Title,
			Description: res.Description,
			Deadline:    res.Deadline,
			IsDone:      true,
		}


		res, er := ef.taskTodo.UpdateTask(idd, executedTask)
		if er != nil {
			writeResponse(w,http.StatusBadRequest,[]byte("Error: "+err.Error()))
			return
		}
		response, err := json.Marshal(res)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}


		//declare Rabbit for producer
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"email", // name
			true,    // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)
		failOnError(err, "Failed to declare a queue")

		msg := amqp.Publishing{
			Body: []byte(res.Title),
		}

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			msg)
		failOnError(err, "Failed to publish a message")


		writeResponse(w,http.StatusCreated,response)
	}
}

func (ef *endpointsFactory) GetListUsers() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		listUsers,err := ef.taskTodo.GetListTask()
		if err!=nil{
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		data, err := json.Marshal(listUsers)
		if err != nil {
			writeResponse(w,http.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		writeResponse(w,http.StatusOK,data)

	}
}


func writeResponse(w http.ResponseWriter,status int,msg []byte) {
	w.WriteHeader(status)
	w.Write(msg)
}


func failOnError(err error, msg string) {
	if err != nil {
		fmt.Errorf("%s: %s", msg, err)
	}
}