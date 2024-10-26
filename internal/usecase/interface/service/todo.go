package service

import "simple-todo/internal/entity"

type TodoService interface {
	FindAll() ([]entity.Todo, error)
	FindOne(id int) (*entity.Todo, error)
	Create(todo *entity.Todo) (*entity.Todo, error)
	Update(id int, todo *entity.Todo) (*entity.Todo, error)
	Delete(id int) error
}
