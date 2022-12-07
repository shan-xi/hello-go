package transport

import (
	"context"
	"encoding/json"
	"errors"
	"hello/service"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeTodosEndpoint(svc service.TodoService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		ts, err := svc.GetTodos()
		if err != nil {
			return nil, err
		}
		todosRes := todosResponse{}
		for _, todo := range ts {
			todoRes := todoResponse{Item: todo.Item}
			todosRes = append(todosRes, todoRes)
		}
		return todosRes, nil
	}
}

func MakeCreateTodoEndpoint(svc service.TodoService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createTodoRequest)
		err := svc.CreateTodo(req.Item)
		if err != nil {
			return "fail", err
		}
		return "success", nil
	}
}

func MakeDeleteTodoEndpoint(svc service.TodoService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createTodoRequest)
		err := svc.DeleteTodo(req.Item)
		if err != nil {
			return "fail", err
		}
		return "success", nil
	}
}

type todoResponse struct {
	Item string `json:"item"`
}
type todosResponse []todoResponse

type createTodoRequest struct {
	Item string `json:"item"`
}

func DecodeTodosRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
func DecodeCreateTodoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if r.Method == "POST" {
		var request createTodoRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}
		return request, nil
	} else {
		return nil, errors.New("not supported method")
	}
}

func DecodeDeleteTodoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if r.Method == "DELETE" {
		var request createTodoRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}
		return request, nil
	} else {
		return nil, errors.New("not supported method")
	}
}
