package services

import (
	"github.com/anthonyvii27/api-rest-go/internal/models"
	"github.com/anthonyvii27/api-rest-go/internal/repositories"
)

type TodosService struct {
	repository repositories.TodosRepository
}

func NewTodoService(repository repositories.TodosRepository) TodosService {
	return TodosService{
		repository: repository,
	}
}

func (s TodosService) Create(data models.Todo) (models.Todo, error) {
	todo, err := s.repository.Create(data)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (s TodosService) FindAll() ([]models.Todo, error) {
	todos, err := s.repository.FindAll()

	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (s TodosService) FindOne(id string) (models.Todo, error) {
	todo, err := s.repository.FindOne(id)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (s TodosService) UpdateOne(id string, data models.Todo) (models.Todo, error) {
	var updated models.Todo

	todo, err := s.repository.FindOne(id)

	if err != nil {
		return updated, err
	}

	updated, err = s.repository.UpdateOne(todo, data)

	if err != nil {
		return updated, err
	}

	return updated, nil
}

func (s TodosService) DeleteOne(id string) error {
	todo, err := s.repository.FindOne(id)

	if err != nil {
		return err
	}

	if err = s.repository.DeleteOne(todo); err != nil {
		return err
	}

	return nil
}
