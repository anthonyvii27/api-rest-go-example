package configurations

import (
	"github.com/anthonyvii27/api-rest-go/internal/errors"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type AuthConfig struct {
	APIKey string `json:"api_key" default: "admin" split_words: "true"`
}

func apiKeyFromEnvironment() AuthConfig {
	var cfg AuthConfig

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal().Str("context", "auth").Err(err).Msg(errors.ErrProcessingEnvironmentIntoConfiguration.Error())
	}

	return cfg
}

func NewAuthConfig() *AuthConfig {
	cfg := apiKeyFromEnvironment()
	return &cfg
}
