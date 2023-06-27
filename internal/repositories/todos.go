package repositories

import (
	"errors"
	customerrors "github.com/anthonyvii27/api-rest-go/internal/errors"
	"github.com/anthonyvii27/api-rest-go/internal/models"
	"gorm.io/gorm"
)

type TodosRepository struct {
	postgres *gorm.DB
}

func NewTodosRepository(postgres *gorm.DB) TodosRepository {
	return TodosRepository{
		postgres: postgres,
	}
}

func (r TodosRepository) Create(todo models.Todo) (models.Todo, error) {
	if err := r.postgres.Create(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (r TodosRepository) FindAll() ([]models.Todo, error) {
	var todos []models.Todo

	if err := r.postgres.Find(&todos).Error; err != nil {
		return todos, err
	}

	return todos, nil
}

func (r TodosRepository) FindOne(id string) (models.Todo, error) {
	var todo models.Todo

	if err := r.postgres.First(&todo, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return todo, customerrors.ErrTodoNotFound
		}

		return todo, err
	}

	return todo, nil
}

func (r TodosRepository) UpdateOne(todo, data models.Todo) (models.Todo, error) {
	if err := r.postgres.Model(&todo).Updates(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return todo, customerrors.ErrTodoNotFound
		}

		return todo, err
	}

	return todo, nil
}

func (r TodosRepository) DeleteOne(todo models.Todo) error {
	if err := r.postgres.Delete(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customerrors.ErrTodoNotFound
		}

		return err
	}

	return nil
}
