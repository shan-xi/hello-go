package service

import "fmt"

type HelloService interface {
	SayHello(string) string
}

func HelloServiceInstance() HelloService {
	return helloService{}
}

type helloService struct{}

func (helloService) SayHello(s string) string {
	if s == "" {
		return "Hello World!"
	}
	return fmt.Sprintf("Hello %s", s)
}
