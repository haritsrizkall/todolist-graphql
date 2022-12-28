package repository

import (
	"todolist-graphql/entity"

	"gorm.io/gorm"
)

type TodoRepository interface {
	FindAll() ([]*entity.Todo, error)
	FindById(id int32) (*entity.Todo, error)
	Create(todo *entity.Todo) (*entity.Todo, error)
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) FindAll() ([]*entity.Todo, error) {
	var todos []*entity.Todo
	err := r.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepository) FindById(id int32) (*entity.Todo, error) {
	var todo entity.Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) Create(todo *entity.Todo) (*entity.Todo, error) {
	err := r.db.Create(todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}
