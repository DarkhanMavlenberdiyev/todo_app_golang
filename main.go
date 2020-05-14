package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"net/smtp"
	"os"
	_ "github.com/valyala/fasthttp"
	"./endpoint"
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
	dbUser := endpoint.NewPostgreUser(user)


	endpoints := endpoint.NewEndpointsFactory(db)
	endpointUser := endpoint.NewEndpointsFactoryUser(dbUser)



	// router
	router := mux.NewRouter()



	go func() {

		// declare rabbit for consumer
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

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		failOnError(err, "Failed to register a consumer")

		forever := make(chan bool)

		emailSender := "" // email of sender
		pass := "" // password

		go func() {
			// for loop for checking queue body
			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)
				emails,err := endpointUser.UserInt.GetListUsers()
				if err!=nil {
					fmt.Print(err.Error())
				}

				//with for loop send email to all users
				for _,email := range emails {
					SendEmail(emailSender,pass,email.Email,string(d.Body))
				}

			}
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	}()

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

	// Start the server
	http.ListenAndServe("0.0.0.0:8000", router)


	return nil
}

//if don't connect to rabbitMQ
func failOnError(err error, msg string) {
	if err != nil {
		fmt.Errorf("%s: %s", msg, err)
	}
}

// SendEmail ...
func SendEmail (emailSender string,pass string,emailReceiver string,taskTitle string){
	hostURL := ""// host
	hostPort := "" //port


	emailAuth := smtp.PlainAuth(
		"",
		emailSender,
		pass,
		hostURL,
	)
	msg := []byte("To: "+emailReceiver+"\r\n"+
		"Subject: "+"Hello"+"\r\n"+taskTitle+" task has been executed")
	err := smtp.SendMail(
		hostURL+":"+hostPort,
		emailAuth,
		emailSender,
		[]string{emailReceiver},
		msg,)
	if err!=nil {
		fmt.Print(err)
	}
}