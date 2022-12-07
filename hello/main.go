package main

import (
	"hello/db"
	"hello/service"
	"hello/transport"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/spf13/viper"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	db.SetupMongoDBConnection()
	db.GetPostgresqlDBConnection()

	var svc service.HelloService
	svc = service.HelloServiceInstance(logger)
	var todoSvc service.TodoService
	todoSvc = service.TodoServiceInstance(logger)

	sayHelloHandler := httptransport.NewServer(
		transport.MakeSayHelloEndpoint(svc),
		transport.DecodeSayHelloRequest,
		transport.EncodeResponse,
	)
	visitHandler := httptransport.NewServer(
		transport.MakeVisitEndpoint(svc),
		transport.DecodeVisitRequest,
		transport.EncodeResponse,
	)
	visitsHandler := httptransport.NewServer(
		transport.MakeVisitsEndpoint(svc),
		transport.DecodeVisitsRequest,
		transport.EncodeResponse,
	)
	deleteVisitHandler := httptransport.NewServer(
		transport.MakeDeleteVisitEndpoint(svc),
		transport.DecodeDeleteVisitRequest,
		transport.EncodeResponse,
	)

	todosHandler := httptransport.NewServer(
		transport.MakeTodosEndpoint(todoSvc),
		transport.DecodeTodosRequest,
		transport.EncodeResponse,
	)
	createTodoHandler := httptransport.NewServer(
		transport.MakeCreateTodoEndpoint(todoSvc),
		transport.DecodeCreateTodoRequest,
		transport.EncodeResponse,
	)
	deleteTodoHandler := httptransport.NewServer(
		transport.MakeDeleteTodoEndpoint(todoSvc),
		transport.DecodeDeleteTodoRequest,
		transport.EncodeResponse,
	)
	http.Handle("/sayHello", sayHelloHandler)
	http.Handle("/visit", visitHandler)
	http.Handle("/visits", visitsHandler)
	http.Handle("/deleteVisit", deleteVisitHandler)
	http.Handle("/todos", todosHandler)
	http.Handle("/createTodo", createTodoHandler)
	http.Handle("/deleteTodo", deleteTodoHandler)
	logger.Log("msg", "HTTP", "addr", ":3000")
	logger.Log("err", http.ListenAndServe(":3000", nil))
}
