package configurations

import (
	"github.com/anthonyvii27/api-rest-go/internal/errors"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type DatabaseConfiguration struct {
	ConnectionString string `required:"true" split_words:"true" json:"connectionString"`
}

func databaseConfigurationFromLocalEnvironment() DatabaseConfiguration {
	var cfg DatabaseConfiguration

	err := envconfig.Process("DATABASE", &cfg)

	if err != nil {
		log.Fatal().Str("context", "database").Err(err).Msg(errors.ErrProcessingEnvironmentIntoConfiguration.Error())
	}

	return cfg
}

func NewDatabaseConfiguration() *DatabaseConfiguration {
	cfg := databaseConfigurationFromLocalEnvironment()
	return &cfg
}
