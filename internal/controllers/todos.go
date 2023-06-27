package controllers

import (
	"errors"
	customerrors "github.com/anthonyvii27/api-rest-go/internal/errors"
	"github.com/anthonyvii27/api-rest-go/internal/models"
	"github.com/anthonyvii27/api-rest-go/internal/requests"
	"github.com/anthonyvii27/api-rest-go/internal/responses"
	"github.com/anthonyvii27/api-rest-go/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type TodosController struct {
	service services.TodosService
}

func NewTodosController(service services.TodosService) TodosController {
	return TodosController{
		service: service,
	}
}

func (c TodosController) Create(ctx *gin.Context) {
	var request requests.CreateTodoRequest
	var response responses.CreateTodoResponse
	var data models.Todo

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg(customerrors.ErrBindingRequestBody.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": customerrors.ErrBindingRequestBody.Error()})
		return
	}

	if err := copier.Copy(&data, &request); err != nil {
		log.Error().Err(err).Msg(customerrors.ErrHandlingRequestBody.Error())
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": customerrors.ErrHandlingRequestBody.Error()})
		return
	}

	todo, err := c.service.Create(data)

	if err != nil {
		log.Error().Err(err).Msg(customerrors.ErrExecutingRequestedOperation.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": customerrors.ErrExecutingRequestedOperation.Error()})
		return
	}

	if err = copier.Copy(&response.Data, &todo); err != nil {
		log.Error().Err(err).Msg(customerrors.ErrBuildingResponse.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": customerrors.ErrBuildingResponse.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c TodosController) FindAll(ctx *gin.Context) {
	var response responses.FindAllTodosResponse

	todos, err := c.service.FindAll()

	if err != nil {
		log.Error().Err(err).Msg(customerrors.ErrExecutingRequestedOperation.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": customerrors.ErrExecutingRequestedOperation.Error()})
		return
	}

	response.Data = make([]struct {
		Id          string     `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Closed      bool       `json:"closed"`
		CreatedAt   time.Time  `json:"createdAt"`
		UpdatedAt   *time.Time `json:"updatedAt"`
	}, 0)

	if err = copier.Copy(&response.Data, &todos); err != nil {
		log.Error().Err(err).Msg(customerrors.ErrBuildingResponse.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": customerrors.ErrBuildingResponse.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c TodosController) FindOne(ctx *gin.Context) {
	var response responses.FindOneTodoResponse
	id := ctx.Param("id")

	todo, err := c.service.FindOne(id)

	if err != nil {
		if errors.Is(err, customerrors.ErrTodoNotFound) {
			log.Error().Err(err).Msg(err.Error())
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			log.Error().Err(err).Msg(customerrors.ErrExecutingRequestedOperation.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": customerrors.ErrExecutingRequestedOperation.Error()})
		}
		return
	}

	if err = copier.Copy(&response.Data, &todo); err != nil {
		log.Error().Err(err).Msg(customerrors.ErrBuildingResponse.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": customerrors.ErrBuildingResponse.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c TodosController) UpdateOne(ctx *gin.Context) {
	var request requests.UpdateTodoRequest
	var response responses.UpdateTodoResponse
	var data models.Todo

	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg(customerrors.ErrBindingRequestBody.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": customerrors.ErrBindingRequestBody.Error()})
		return
	}

	if err := copier.Copy(&data, &request); err != nil {
		log.Error().Err(err).Msg(customerrors.ErrHandlingRequestBody.Error())
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": customerrors.ErrHandlingRequestBody.Error()})
		return
	}

	todo, err := c.service.UpdateOne(id, data)

	if err != nil {
		if errors.Is(err, customerrors.ErrTodoNotFound) {
			log.Error().Msg(err.Error())
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			log.Error().Err(err).Msg(customerrors.ErrExecutingRequestedOperation.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": customerrors.ErrExecutingRequestedOperation.Error()})
		}
		return
	}

	if err := copier.Copy(&response.Data, &todo); err != nil {
		log.Error().Err(err).Msg(customerrors.ErrBuildingResponse.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": customerrors.ErrBuildingResponse.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c TodosController) DeleteOne(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.service.DeleteOne(id); err != nil {
		if errors.Is(err, customerrors.ErrTodoNotFound) {
			log.Error().Msg(err.Error())
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			log.Error().Err(err).Msg(customerrors.ErrExecutingRequestedOperation.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": customerrors.ErrExecutingRequestedOperation.Error()})
		}
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
