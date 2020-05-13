package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"net/http"
	"os"
	_ "github.com/valyala/fasthttp"
	"./endpoint"

)

func main() {

	app := cli.NewApp()
	app.Commands = cli.Commands{
		&cli.Command{
			Name:   "start",
			Usage:  "start the local server",
			Action: StartServer,
		},
	}
	app.Run(os.Args)

}

func StartServer(c *cli.Context) error {
	user := endpoint.PostgreConfig{
		User:     "postgres",
		Password: "daha2000",
		Port:     "5432",
		Host:     "127.0.0.1",
	}
	db := endpoint.NewPostgre(user)
	dbUser := endpoint.NewPostgreUser(user)

	endpoints := endpoint.NewEndpointsFactory(db)
	endpointUser := endpoint.NewEndpointsFactoryUser(dbUser)
	router := mux.NewRouter()


	//endpoints for tasks todo
	router.Methods("GET").Path("/api/todo").HandlerFunc(endpoints.GetListTask())
	router.Methods("GET").Path("/api/todo/{id}").HandlerFunc(endpoints.GetTask("id"))
	router.Methods("PUT").Path("/api/todo/{id}").HandlerFunc(endpoints.UpdateTask("id"))
	router.Methods("POST").Path("/api/todo/{id}/execute").HandlerFunc(endpoints.ExecuteTask("id"))
	router.Methods("POST").Path("/api/todo").HandlerFunc(endpoints.CreateTask())
	router.Methods("DELETE").Path("/api/todo/{id}").HandlerFunc(endpoints.DeleteTask("id"))

	//endpoints for authorization
	router.Methods("POST").Path("/create").HandlerFunc(endpointUser.CreateUser())
	router.Methods("POST").Path("/sessions").HandlerFunc(endpointUser.SessionUser())

	http.ListenAndServe("0.0.0.0:8000", router)


	return nil
}
