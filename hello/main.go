package main

import (
	"hello/service"
	"hello/transport"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

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
