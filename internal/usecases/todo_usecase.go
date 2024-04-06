package usecases

import (
	"todo/webservice/internal/entities"
)

type TodoUsecase struct {
	todoRepository *entities.TodoRepository
}

func NewTodoUsecase(todoRepository *entities.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		todoRepository: todoRepository,
	}
}

func (uc *TodoUsecase) GetByID(id string) (entities.Todo, error) {
	return uc.todoRepository.GetByID(id)
}

func (uc *TodoUsecase) Create(draft *entities.DraftTodo) (entities.Todo, error) {
	return uc.todoRepository.CreateTodo(*draft)
}
