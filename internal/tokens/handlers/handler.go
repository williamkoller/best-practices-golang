package handlers

import (
	"best-practices-golang/internal/tokens/tasks"
	"best-practices-golang/internal/tokens/usecases"
	"encoding/json"
	"net/http"

	"github.com/hibiken/asynq"
)

type TokenRequest struct {
	Token string `json:"token"`
}

type TokensHandler struct {
	usecase     *usecases.TokenUseCase
	asynqClient *asynq.Client
}

func NewTokensHandler(uc *usecases.TokenUseCase, client *asynq.Client) *TokensHandler {
	return &TokensHandler{
		usecase:     uc,
		asynqClient: client,
	}
}

func (h *TokensHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req TokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		h.usecase.Log.Error().Err(err).Msg("Failed to decode request body")
		return
	}

	task, err := tasks.NewProcessTokenTask(req.Token)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		h.usecase.Log.Error().Err(err).Msg("Failed to create task")
		return
	}

	if _, err := h.asynqClient.Enqueue(task); err != nil {
		http.Error(w, "Failed to enqueue task", http.StatusInternalServerError)
		h.usecase.Log.Error().Err(err).Msg("Failed to enqueue task")
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Token enqueued successfully"))
}
