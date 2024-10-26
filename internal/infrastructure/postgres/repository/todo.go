package repository

import (
	"simple-todo/internal/entity"
	"simple-todo/internal/usecase/interface/repository"

	"github.com/jmoiron/sqlx"
)

type todoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) repository.TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) GetAll() ([]entity.Todo, error) {
	rows, err := r.db.Query("SELECT id, title, description, completed FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []entity.Todo
	for rows.Next() {
		var todo entity.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *todoRepository) GetByID(id int) (*entity.Todo, error) {
	row := r.db.QueryRow("SELECT id, title, description, completed FROM todo WHERE id = $1", id)

	var todo entity.Todo
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *todoRepository) Create(todo *entity.Todo) (*entity.Todo, error) {
	row := r.db.QueryRow("INSERT INTO todo (title, description, completed) VALUES ($1, $2, $3) RETURNING id", todo.Title, todo.Description, todo.Completed)
	if err := row.Scan(&todo.ID); err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *todoRepository) Update(id int, todo *entity.Todo) (*entity.Todo, error) {
	row := r.db.QueryRow("UPDATE todo SET title = $1, description = $2, completed = $3 WHERE id = $4 RETURNING id", todo.Title, todo.Description, todo.Completed, id)

	if err := row.Scan(&todo.ID); err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *todoRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM todo WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
