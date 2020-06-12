package endpoint

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
)

//endpoint for tasks
func NewEndpointsFactory(tasktodo TaskTodo) *endpointsFactory {
	return &endpointsFactory{taskTodo: tasktodo}
}

//endpointsFactory ...
type endpointsFactory struct {
	taskTodo TaskTodo
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

