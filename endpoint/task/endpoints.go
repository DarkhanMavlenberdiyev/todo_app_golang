package task

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
)

//endpoint for tasks
func NewEndpointsFactory(tasktodo TaskTodo) *endpointsFactory {
	return &endpointsFactory{taskTodo: tasktodo}
}
//endpointsFactory task...
type endpointsFactory struct {
	taskTodo TaskTodo
}

//endpoint for user
func NewEndpointsFactoryUser(userinfo UserInfo) *endpointsFactoryUser {
	return &endpointsFactoryUser{userInfo: userinfo}
}

//endpointsFactory user
type endpointsFactoryUser struct {
	userInfo UserInfo
}

//GetUser
func (ef *endpointsFactoryUser) GetUser() func (ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		vars := ctx.FormValue("id")
		id, err := strconv.Atoi(string(vars))
		if err != nil {
			writeResponse(ctx,fasthttp.StatusBadRequest,[]byte("Error: "+err.Error()))
			return
		}
		res,err := ef.userInfo.GetUser(id)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusNotFound,[]byte("Not found"))
			return
		}
		data,err := json.Marshal(res)
		if err != nil{
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error! Try again"))
			return
		}
		writeResponse(ctx,fasthttp.StatusOK,data)
	}
}

//GetTask ...
func (ef *endpointsFactory) GetTask() func(ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		vars := ctx.FormValue("id")

		idd, _ := strconv.Atoi(string(vars))
		res, err := ef.taskTodo.GetTask(idd)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		data, err := json.Marshal(res)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		writeResponse(ctx,fasthttp.StatusOK,data)
	}
}

//CreateUser...
func (ef *endpointsFactoryUser) CreateUser() func (ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		data := ctx.PostBody()
		user := &User{}
		if err := json.Unmarshal(data,user);err != nil {
			writeResponse(ctx,fasthttp.StatusBadRequest,[]byte("Error: invalid input"))
			return
		}
		res, err := ef.userInfo.CreateUser(user)
		if err !=nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error: try again"))
			return
		}
		response,err := json.Marshal(res)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error!"))
			return
		}
		writeResponse(ctx,fasthttp.StatusCreated,response)
	}
}

//CreateTask ...
func (ef *endpointsFactory) CreateTask() func(ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		data := ctx.PostBody()

		task := &Task{}
		if err := json.Unmarshal(data, task); err != nil {
			writeResponse(ctx,fasthttp.StatusBadRequest,[]byte("Error: "+err.Error()))
			return
		}
		result, err := ef.taskTodo.CreateTask(task)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		response, err := json.Marshal(result)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		writeResponse(ctx,fasthttp.StatusCreated,response)
	}
}

//UpdateTask ...
func (ef *endpointsFactory) UpdateTask() func(ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		vars := ctx.FormValue("id")

		data := ctx.PostBody()

		task := &Task{}
		if err := json.Unmarshal(data, task); err != nil {
			writeResponse(ctx,fasthttp.StatusBadRequest,[]byte("Error: "+err.Error()))
			return
		}
		idd, _ := strconv.Atoi(string(vars))
		res, err := ef.taskTodo.UpdateTask(idd, task)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusBadRequest,[]byte("Error: "+err.Error()))
			return
		}
		response, err := json.Marshal(res)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		writeResponse(ctx,fasthttp.StatusCreated,response)

	}
}

//DeleteTask ...
func (ef *endpointsFactory) DeleteTask() func(ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		vars := ctx.FormValue("id")
		idd, _ := strconv.Atoi(string(vars))
		err := ef.taskTodo.DeleteTask(idd)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		writeResponse(ctx,fasthttp.StatusOK,[]byte("Task is deleted successfully"))


	}
}
//GetListTask ...
func (ef *endpointsFactory) GetListTask() func(ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		list,err := ef.taskTodo.GetListTask()
		if err!=nil{
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		data, err := json.Marshal(list)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		writeResponse(ctx,fasthttp.StatusOK,data)

	}
}

// ExecuteTask ...
func (ef *endpointsFactory) ExecuteTask() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		vars := ctx.FormValue("id")
		idd, _ := strconv.Atoi(string(vars))
		res, err := ef.taskTodo.GetTask(idd)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error: "+err.Error()))
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
			writeResponse(ctx,fasthttp.StatusBadRequest,[]byte("Error: "+err.Error()))
			return
		}
		response, err := json.Marshal(res)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}




		writeResponse(ctx,fasthttp.StatusCreated,response)
	}
}


func writeResponse(ctx *fasthttp.RequestCtx,status int,msg []byte) {
	ctx.SetStatusCode(status)
	ctx.Write(msg)
}

