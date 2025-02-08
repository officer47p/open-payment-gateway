package config

// type EnvVariables struct {
// DBUrl       string
// DBPort      int64
// DBName      string
// DBUser      string
// DBPassword  string
// 	ProviderUrl string
// 	NatsUrl     string
// 	Environment string // This can be prod or dev
// }

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var ErrEnvVariableNotFound = errors.New("environment variable was not found")

func NewErrEnvVariableNotFound(key string) error {
	return errors.Join(ErrEnvVariableNotFound, fmt.Errorf("variable %s was not found", key))
}

type Env struct {
	ENV  string
	PORT int

	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DBNAME   string

	NATS_HOST string

	PROVIDER_HOST string

	// REDIS_HOST     string
	// REDIS_PORT     string
	// REDIS_PASSWORD string

	// ACCESS_TOKEN_EXPIRY_HOUR  int
	// REFRESH_TOKEN_EXPIRY_HOUR int
	// ACCESS_TOKEN_SECRET       string
	// REFRESH_TOKEN_SECRET      string
}

func LoadAndValidateEnvironmentVariables() (*Env, error) {
	ENV := os.Getenv("ENV")
	PORT := os.Getenv("PORT")

	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DBNAME := os.Getenv("POSTGRES_DBNAME")

	NATS_HOST := os.Getenv("NATS_HOST")

	PROVIDER_HOST := os.Getenv("PROVIDER_HOST")

	// REDIS_HOST := os.Getenv("REDIS_HOST")
	// REDIS_PORT := os.Getenv("REDIS_PORT")
	// REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")

	// ACCESS_TOKEN_EXPIRY_HOUR := os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR")
	// REFRESH_TOKEN_EXPIRY_HOUR := os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR")
	// ACCESS_TOKEN_SECRET := os.Getenv("ACCESS_TOKEN_SECRET")
	// REFRESH_TOKEN_SECRET := os.Getenv("REFRESH_TOKEN_SECRET")

	if ENV == "" {
		return nil, NewErrEnvVariableNotFound("ENV")
	}
	if PORT == "" {
		return nil, NewErrEnvVariableNotFound("PORT")
	}

	if POSTGRES_HOST == "" {
		return nil, NewErrEnvVariableNotFound("POSTGRES_HOST")
	}
	if POSTGRES_PORT == "" {
		return nil, NewErrEnvVariableNotFound("POSTGRES_PORT")
	}
	if POSTGRES_USER == "" {
		return nil, NewErrEnvVariableNotFound("POSTGRES_USER")
	}
	if POSTGRES_PASSWORD == "" {
		return nil, NewErrEnvVariableNotFound("POSTGRES_PASSWORD")
	}
	if POSTGRES_DBNAME == "" {
		return nil, NewErrEnvVariableNotFound("POSTGRES_DBNAME")
	}

	if NATS_HOST == "" {
		return nil, NewErrEnvVariableNotFound("NATS_HOST")
	}

	if PROVIDER_HOST == "" {
		return nil, NewErrEnvVariableNotFound("PROVIDER_HOST")
	}

	// if REDIS_HOST == "" {
	// 	return nil, NewErrEnvVariableNotFound("REDIS_HOST")
	// }
	// if REDIS_PORT == "" {
	// 	return nil, NewErrEnvVariableNotFound("REDIS_PORT")
	// }
	// if REDIS_PASSWORD == "" {
	// 	return nil, NewErrEnvVariableNotFound("REDIS_PASSWORD")
	// }
	// if ACCESS_TOKEN_EXPIRY_HOUR == "" {
	// 	return nil, NewErrEnvVariableNotFound("ACCESS_TOKEN_EXPIRY_HOUR")
	// }
	// if REFRESH_TOKEN_EXPIRY_HOUR == "" {
	// 	return nil, NewErrEnvVariableNotFound("REFRESH_TOKEN_EXPIRY_HOUR")
	// }
	// if ACCESS_TOKEN_SECRET == "" {
	// 	return nil, NewErrEnvVariableNotFound("ACCESS_TOKEN_SECRET")
	// }
	// if REFRESH_TOKEN_SECRET == "" {
	// 	return nil, NewErrEnvVariableNotFound("REFRESH_TOKEN_SECRET")
	// }

	portNumber, err := strconv.Atoi(PORT)
	if err != nil {
		return nil, NewErrEnvVariableNotFound("PORT")
	}

	// accessTokenExpiryHour, err := strconv.Atoi(ACCESS_TOKEN_EXPIRY_HOUR)
	// if err != nil {
	// 	return nil, NewErrEnvVariableNotFound("ACCESS_TOKEN_EXPIRY_HOUR")
	// }

	// refreshTokenExpiryHour, err := strconv.Atoi(REFRESH_TOKEN_EXPIRY_HOUR)
	// if err != nil {
	// 	return nil, NewErrEnvVariableNotFound("REFRESH_TOKEN_EXPIRY_HOUR")
	// }

	env := Env{
		ENV:  ENV,
		PORT: portNumber,

		POSTGRES_HOST:     POSTGRES_HOST,
		POSTGRES_PORT:     POSTGRES_PORT,
		POSTGRES_USER:     POSTGRES_USER,
		POSTGRES_PASSWORD: POSTGRES_PASSWORD,
		POSTGRES_DBNAME:   POSTGRES_DBNAME,

		NATS_HOST: NATS_HOST,

		PROVIDER_HOST: PROVIDER_HOST,

		// REDIS_HOST:     REDIS_HOST,
		// REDIS_PORT:     REDIS_PORT,
		// REDIS_PASSWORD: REDIS_PASSWORD,

		// ACCESS_TOKEN_EXPIRY_HOUR:  accessTokenExpiryHour,
		// REFRESH_TOKEN_EXPIRY_HOUR: refreshTokenExpiryHour,
		// ACCESS_TOKEN_SECRET:       ACCESS_TOKEN_SECRET,
		// REFRESH_TOKEN_SECRET:      REFRESH_TOKEN_SECRET,
	}

	return &env, nil

}

func Validate() (*Env, error) {
	if os.Getenv("ENV") != "production" && os.Getenv("ENV") != "staging" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	return LoadAndValidateEnvironmentVariables()
}
