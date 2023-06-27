package handlers

import (
	"github.com/anthonyvii27/api-rest-go/internal/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(todosController controllers.TodosController) http.Handler {
	router := gin.New()
	router.Use(gin.Logger())

	v1 := router.Group("/api/v1")
	{
		todos := v1.Group("/todos")
		{
			todos.POST("/", todosController.Create)
			todos.GET("/", todosController.FindAll)
			todos.GET("/:id", todosController.FindOne)
			todos.PUT("/:id", todosController.UpdateOne)
			todos.DELETE("/:id", todosController.DeleteOne)
		}
	}

	return router
}
