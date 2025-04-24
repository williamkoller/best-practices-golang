package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

const TypeProcessToken = "process:token"

func NewProcessTokenTask(token string) (*asynq.Task, error) {
	payload, err := json.Marshal(map[string]string{"token": token})
	if err != nil {
		return nil, err
	}

	task := asynq.NewTask(TypeProcessToken, payload)

	return task, nil
}
