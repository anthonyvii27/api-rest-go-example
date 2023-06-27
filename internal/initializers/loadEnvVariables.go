package initializers

import (
	"github.com/anthonyvii27/api-rest-go/internal/errors"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal().Err(err).Msg(errors.ErrLoadingDotEnvFile.Error())
	}
}
