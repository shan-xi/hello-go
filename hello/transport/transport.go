package transport

import (
	"context"
	"encoding/json"
	"errors"
	"hello/service"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeSayHelloEndpoint(svc service.HelloService) endpoint.Endpoint {
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

func DecodeSayHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func MakeVisitEndpoint(svc service.HelloService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(visitRequest)
		v, err := svc.GetVisit(req.Id)
		if err != nil {
			return visitResponse{}, err
		}
		return visitResponse{v.Name, v.Timestamp.String()}, err
	}
}

type visitRequest struct {
	Id string `json:"id"`
}

type visitResponse struct {
	Name      string `json:"name"`
	Timestamp string `json:"time"`
}

type visitsRequest struct{}
type visitsResponse []visitResponse

func DecodeVisitRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if r.Method == "POST" {
		var request visitRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}
		return request, nil
	} else if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		request := visitRequest{Id: id}
		return request, nil
	} else {

		return visitRequest{}, nil
	}
}

func MakeVisitsEndpoint(svc service.HelloService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		vs, err := svc.GetVisits()
		if err != nil {
			return visitsResponse{}, err
		}

		vvs := visitsResponse{}
		for _, v := range vs {
			t := visitResponse{v.Name, v.Timestamp.String()}
			vvs = append(vvs, t)
		}
		return vvs, err
	}
}

func DecodeVisitsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return visitsRequest{}, nil
}

func MakeDeleteVisitEndpoint(svc service.HelloService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(visitRequest)
		err := svc.DeleteVisit(req.Id)
		if err != nil {
			return "fail", err
		}
		return "success", nil
	}
}

func DecodeDeleteVisitRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if r.Method == "DELETE" {
		var request visitRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}
		return request, nil
	} else {
		return "not supported method", errors.New("not supported method")
	}
}
