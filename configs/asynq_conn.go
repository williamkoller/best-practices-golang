package configs

import (
	"best-practices-golang"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
)

var redis, _, _ = best_practices_golang.Env()

func ConnectRedis() (*asynq.RedisClientOpt, error) {
	if redis == "" {
		return nil, nil
	}

	return &asynq.RedisClientOpt{
		Addr: redis,
	}, nil
}

func CreateAsynqClient(redis *asynq.RedisClientOpt, log zerolog.Logger) (*asynq.Client, error) {
	if redis == nil {
		return nil, nil
	}
	client := asynq.NewClient(redis)
	log.Info().Msg("Redis client created successfully")
	return client, nil
}
