package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"github.com/valyala/fasthttp"
	"github.com/gospodinzerkalo/todo_app_golang/endpoint"
)

func main() {
	// CLI command for starting APP
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


	//User for coonection to db
	user := endpoint.PostgreConfig{
		User:     "postgres",
		Password: "daha2000",
		Port:     "5432",
		Host:     "127.0.0.1",
	}
	//Connect db
	db := endpoint.NewPostgre(user)


	endpoints := endpoint.NewEndpointsFactory(db)



	// router
	router := fasthttprouter.New()




	//endpoints for tasks todo
	router.GET("/api/todo",endpoints.GetListTask())
	router.GET("/api/todo/:id",endpoints.GetTask())
	router.PUT("/api/todo/:id",endpoints.UpdateTask())
	router.POST("/api/todo/:id/execute",endpoints.ExecuteTask())
	router.POST("/api/todo",endpoints.CreateTask())
	router.DELETE("/api/todo/:id",endpoints.DeleteTask())


	// Start the server
	log.Fatal(fasthttp.ListenAndServe(":8000",router.Handler))

	return nil
}


