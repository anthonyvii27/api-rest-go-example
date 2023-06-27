package configurations

import (
	"github.com/anthonyvii27/api-rest-go/internal/errors"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type AppConfiguration struct {
	Environment string `required:"true"`
	Port        int    `required:"true"`
}

func appConfigurationFromEnvironment() AppConfiguration {
	var cfg AppConfiguration

	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatal().Str("context", "application").Err(err).Msg(errors.ErrProcessingEnvironmentIntoConfiguration.Error())
	}

	return cfg
}

func NewAppConfiguration() *AppConfiguration {
	cfg := appConfigurationFromEnvironment()

	return &cfg
}
