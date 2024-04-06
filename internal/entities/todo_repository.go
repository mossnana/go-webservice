package entities

import (
	"time"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (t *TodoRepository) GetByID(id string) (todo Todo, e error) {
	todo = Todo{}
	e = t.db.Where("id = ?", id).First(&todo).Error
	return
}

func (t *TodoRepository) CreateTodo(draft DraftTodo) (todo Todo, e error) {
	todo = Todo{
		Title:       draft.Title,
		Description: draft.Description,
		CreateDate:  time.Now(),
	}
	e = t.db.Create(&todo).Error
	return
}
