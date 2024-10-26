package service

import (
	"simple-todo/internal/entity"
	"simple-todo/internal/usecase/interface/repository"
	"simple-todo/internal/usecase/interface/service"
)

type todoServiceImpl struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) service.TodoService {
	return &todoServiceImpl{repo: repo}
}

func (s *todoServiceImpl) FindAll() ([]entity.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoServiceImpl) FindOne(id int) (*entity.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *todoServiceImpl) Create(todo *entity.Todo) (*entity.Todo, error) {
	return s.repo.Create(todo)
}

func (s *todoServiceImpl) Update(id int, todo *entity.Todo) (*entity.Todo, error) {
	return s.repo.Update(id, todo)
}

func (s *todoServiceImpl) Delete(id int) error {
	return s.repo.Delete(id)
}
