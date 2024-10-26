package repository

import "simple-todo/internal/entity"

type TodoRepository interface {
	GetAll() ([]entity.Todo, error)
	GetByID(id int) (*entity.Todo, error)
	Create(todo *entity.Todo) (*entity.Todo, error)
	Update(id int, todo *entity.Todo) (*entity.Todo, error)
	Delete(id int) error
}
