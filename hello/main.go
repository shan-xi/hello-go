package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc HelloService
	svc = helloService{}

	sayHelloHandler := httptransport.NewServer(
		makeSayHelloEndpoint(svc),
		decodeSayHelloRequest,
		encodeResponse,
	)

	http.Handle("/sayHello", sayHelloHandler)
	logger.Log("msg", "HTTP", "addr", ":3000")
	logger.Log("err", http.ListenAndServe(":3000", nil))
}
