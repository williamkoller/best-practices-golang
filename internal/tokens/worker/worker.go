package worker

import (
	"best-practices-golang/internal/tokens/handlers/request"
	"best-practices-golang/internal/tokens/tasks"
	"best-practices-golang/internal/tokens/usecases"
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
)

type TokenProcessor struct {
	uc *usecases.TokenUseCase
}

func NewTokenProcessor(uc *usecases.TokenUseCase) *TokenProcessor {
	return &TokenProcessor{
		uc: uc,
	}
}

func (tp *TokenProcessor) HandleProcessTokenTask(ctx context.Context, t *asynq.Task) error {
	var payload request.TokenRequest
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		tp.uc.Log.Error().Err(err).Msg("Failed to unmarshal token payload")
		return err
	}

	if err := payload.Validate(); err != nil {
		tp.uc.Log.Error().Err(err).Msg("Validation error in token payload")
		return err
	}

	if err := tp.uc.Execute(ctx, payload); err != nil {
		tp.uc.Log.Error().Err(err).Str("token", payload.Token).Msg("Failed to store token")
		return err
	}

	tp.uc.Log.Info().Str("token", payload.Token).Msg("Token processed and stored successfully")
	return nil
}

func StartWorker(redis string, uc *usecases.TokenUseCase) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redis},
		asynq.Config{
			Concurrency: 10,
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, t *asynq.Task, err error) {
				uc.Log.Error().
					Str("task_type", t.Type()).
					Str("payload", string(t.Payload())).
					Err(err).
					Msg("Task failed")

				// Create and enqueue a Dead Letter Queue (DLQ) task
				dlqTask, dlqErr := tasks.NewDeadLetterTask(t.Type(), string(t.Payload()), err.Error())
				if dlqErr != nil {
					uc.Log.Error().Err(dlqErr).Msg("Failed to create DLQ task")
					return
				}

				client := asynq.NewClient(asynq.RedisClientOpt{Addr: redis})
				defer func(client *asynq.Client) {
					err := client.Close()
					if err != nil {

					}
				}(client)

				if _, enqueueErr := client.Enqueue(dlqTask); enqueueErr != nil {
					uc.Log.Error().Err(enqueueErr).Msg("Failed to enqueue DLQ task")
				}
			}),
		},
	)

	processor := NewTokenProcessor(uc)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeProcessToken, processor.HandleProcessTokenTask)
	mux.HandleFunc(tasks.TypeDeadLetterQueue, func(ctx context.Context, t *asynq.Task) error {
		return tasks.HandleDeadLetterTask(ctx, t, uc.Log)
	})

	if err := srv.Run(mux); err != nil {
		log.Fatalf("Failed to start worker: %v", err)
	}
}
