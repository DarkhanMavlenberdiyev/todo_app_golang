package user

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
)

// init endpointFactory for user
func NewEndpointsFactory(userinfo UserInfo) *endpointsFactory {
	return &endpointsFactory{userInfo: userinfo}
}


type endpointsFactory struct {
	userInfo UserInfo
}


//GetUser
func (ef *endpointsFactory) GetUser() func (ctx *fasthttp.RequestCtx){
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
//CreateUser...
func (ef *endpointsFactory) CreateUser() func (ctx *fasthttp.RequestCtx){
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

func writeResponse(ctx *fasthttp.RequestCtx,status int,msg []byte) {
	ctx.SetStatusCode(status)
	ctx.Write(msg)
}
