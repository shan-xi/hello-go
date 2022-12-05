package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeSayHelloEndpoint(svc HelloService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(sayeHelloRequest)
		v := svc.SayHello(req.Name)
		return sayHelloResponse{Message: v}, nil
	}
}

type sayeHelloRequest struct {
	Name string `json:"name"`
}

type sayHelloResponse struct {
	Message string `json:"message"`
}

func decodeSayHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println()
	if r.Method == "POST" {
		var request sayeHelloRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}
		return request, nil
	} else {
		return sayeHelloRequest{}, nil
	}
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
