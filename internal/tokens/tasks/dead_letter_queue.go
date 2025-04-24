package tasks

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
)

const TypeDeadLetterQueue = "deadletter:queue"

type DeadLetterPayload struct {
	TaskType string `json:"task_type"`
	Payload  string `json:"payload"`
	Error    string `json:"error"`
}

func NewDeadLetterTask(taskType, payload, errMsg string) (*asynq.Task, error) {
	dlqPayload := DeadLetterPayload{
		TaskType: taskType,
		Payload:  payload,
		Error:    errMsg,
	}
	data, err := json.Marshal(dlqPayload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeDeadLetterQueue, data), nil
}

func HandleDeadLetterTask(ctx context.Context, t *asynq.Task, log zerolog.Logger) error {
	var payload DeadLetterPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal dead letter payload")
		return err
	}

	log.Error().
		Str("task_type", payload.TaskType).
		Str("payload", payload.Payload).
		Str("error", payload.Error).
		Msg("Processing dead letter task")

	return nil
}
