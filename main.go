package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/gospodinzerkalo/todo_app_golang/endpoint/task"
	"github.com/gospodinzerkalo/todo_app_golang/endpoint/user"
	"github.com/gospodinzerkalo/todo_app_golang/redis"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)


// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	redis.InitCache()
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
	dbUser,exist := os.LookupEnv("DB_USER")
	if !exist || dbUser ==""{
		log.Println("db_user not found in .env file")
		return nil
	}
	dbPassword,exist := os.LookupEnv("DB_PASSWORD")
	if !exist || dbPassword =="" {
		log.Println("db_password not found in .env file")
		return nil
	}
	dbPort,exist := os.LookupEnv("DB_PORT")
	if !exist || dbPort ==""{
		log.Println("db_port not found in .env file")
		return nil
	}
	dbHost,exist := os.LookupEnv("DB_HOST")
	if !exist || dbHost ==""{
		log.Println("db_host not found in .env file")
		return nil
	}
	dbDatabaseName,exist := os.LookupEnv("DATABASE_NAME")
	if !exist{
		log.Println("database_name not found in .env file")
	}


	//User for connection to db
	userConfig := task.PostgreConfig{
		User:     dbUser,
		Password: dbPassword,
		Port:     dbPort,
		Host:     dbHost,
		Database: dbDatabaseName,
	}
	userConfig2 := user.PostgreConfig{
		User:     dbUser,
		Password: dbPassword,
		Port:     dbPort,
		Host:     dbHost,
		Database: dbDatabaseName,
	}

	//Connect db
	db,err := task.NewPostgre(userConfig)

	if err != nil {
		log.Println(err)
		return nil
	}
	pgUser,err := user.NewPostgre(userConfig2)
	if err != nil {
		log.Println(err)
		return err
	}


	endpoints := task.NewEndpointsFactory(db)
	endpointsUser := user.NewEndpointsFactory(pgUser)



	// router
	router := fasthttprouter.New()


	//endpoints for tasks todo
	router.GET("/api/todo",endpoints.GetListTask())
	router.GET("/api/todo/:id",endpoints.GetTask())
	router.PUT("/api/todo/:id",endpoints.UpdateTask())
	router.POST("/api/todo/:id/execute",endpoints.ExecuteTask())
	router.POST("/api/todo",endpoints.CreateTask())
	router.DELETE("/api/todo/:id",endpoints.DeleteTask())

	// endpoints for user auth
	router.POST("/signup",endpointsUser.CreateUser())
	router.POST("/signin",endpointsUser.SignIn())
	router.GET("/welcome",endpointsUser.Welcome())
	router.GET("/list",endpointsUser.GetListUsers())


	// Start the server
	log.Fatal(fasthttp.ListenAndServe(":8000",router.Handler))

	return nil
}


