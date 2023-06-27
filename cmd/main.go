package main

import (
	"fmt"
	"github.com/anthonyvii27/api-rest-go/internal/configurations"
	"github.com/anthonyvii27/api-rest-go/internal/controllers"
	"github.com/anthonyvii27/api-rest-go/internal/errors"
	"github.com/anthonyvii27/api-rest-go/internal/handlers"
	"github.com/anthonyvii27/api-rest-go/internal/infrastructure"
	"github.com/anthonyvii27/api-rest-go/internal/initializers"
	"github.com/anthonyvii27/api-rest-go/internal/repositories"
	"github.com/anthonyvii27/api-rest-go/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	applicationConfig := configurations.NewAppConfiguration()
	databaseConfig := configurations.NewDatabaseConfiguration()

	databaseClient := infrastructure.NewDatabaseClient(databaseConfig)

	todosRepository := repositories.NewTodosRepository(databaseClient)
	todosService := services.NewTodoService(todosRepository)
	todosController := controllers.NewTodosController(todosService)

	if applicationConfig.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := handlers.NewRouter(todosController)

	log.Info().Int("port", applicationConfig.Port).Msg("starting the application server")

	if err := http.ListenAndServe(fmt.Sprintf(":%d", applicationConfig.Port), router); err != nil {
		log.Panic().Err(err).Msg(errors.ErrInitializingServer.Error())
	}
}
