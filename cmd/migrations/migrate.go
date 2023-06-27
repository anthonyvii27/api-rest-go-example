package main

import (
	"github.com/anthonyvii27/api-rest-go/internal/configurations"
	"github.com/anthonyvii27/api-rest-go/internal/errors"
	"github.com/anthonyvii27/api-rest-go/internal/infrastructure"
	"github.com/anthonyvii27/api-rest-go/internal/initializers"
	"github.com/anthonyvii27/api-rest-go/internal/models"
	"github.com/rs/zerolog/log"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	databaseConfig := configurations.NewDatabaseConfiguration()
	databaseClient := infrastructure.NewDatabaseClient(databaseConfig)

	if err := databaseClient.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatal().Err(err).Msg(errors.ErrGeneratingMigrate.Error())
	}
}
