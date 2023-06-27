package infrastructure

import (
	"github.com/anthonyvii27/api-rest-go/internal/configurations"
	"github.com/anthonyvii27/api-rest-go/internal/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabaseClient(configuration *configurations.DatabaseConfiguration) *gorm.DB {
	client, err := gorm.Open(postgres.Open(configuration.ConnectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal().Err(err).Msg(errors.ErrConnectingToPostgresDatabase.Error())
	}

	return client
}
