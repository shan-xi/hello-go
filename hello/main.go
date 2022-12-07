package main

import (
	"fmt"
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

	con := db.GetMongoDB(fmt.Sprintf("%v", viper.Get("MONGODB_URI")))
	fmt.Println(con)

	var svc service.HelloService
	svc = service.HelloServiceInstance()

	sayHelloHandler := httptransport.NewServer(
		transport.MakeSayHelloEndpoint(svc),
		transport.DecodeSayHelloRequest,
		transport.EncodeResponse,
	)

	http.Handle("/sayHello", sayHelloHandler)
	logger.Log("msg", "HTTP", "addr", ":3000")
	logger.Log("err", http.ListenAndServe(":3000", nil))
}
