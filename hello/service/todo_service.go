package service

import (
	"hello/db"
	"hello/model"

	"github.com/go-kit/log"
)

type TodoService interface {
	GetTodos() ([]model.Todos, error)
	CreateTodo(string) error
	DeleteTodo(string) error
}

func TodoServiceInstance(l log.Logger) TodoService {
	return &todoService{logger: l}
}

type todoService struct {
	logger log.Logger
}

func (todoService) GetTodos() ([]model.Todos, error) {
	conn, err := db.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	var todos []model.Todos
	conn.Find(&todos)
	return todos, nil
}

func (todoService) CreateTodo(item string) error {
	conn, err := db.GetDatabaseConnection()
	if err != nil {
		return err
	}
	todo := model.Todos{Item: item}
	result := conn.Create(&todo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (todoService) DeleteTodo(item string) error {
	conn, err := db.GetDatabaseConnection()
	if err != nil {
		return err
	}
	todo := model.Todos{Item: item}
	result := conn.Where("item=?", item).Delete(&todo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
