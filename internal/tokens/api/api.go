package api

import (
	"best-practices-golang"
	"best-practices-golang/configs"
	auditsrepositories "best-practices-golang/internal/audits/repositories"
	"best-practices-golang/internal/tokens/handlers"
	"best-practices-golang/internal/tokens/repositories"
	"best-practices-golang/internal/tokens/usecases"
	"best-practices-golang/internal/tokens/worker"
	"best-practices-golang/pkg/logger"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func SetupApi() {
	log := logger.NewLogger()

	db, err := configs.ConnectDB()

	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database")
		return
	}

	redisClientOpt, err := configs.ConnectRedis()

	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to Redis")
		return
	}

	redis, _, port := best_practices_golang.Env()
	portListener := ":" + port

	client, _ := configs.CreateAsynqClient(redisClientOpt, log)

	defer func() {
		if err := client.Close(); err != nil {
			log.Error().Err(err).Msg("Failed to close Redis client")
		} else {
			log.Info().Msg("Redis client closed successfully")
		}

	}()

	tr := repositories.NewTokenRepository(db)
	ar := auditsrepositories.NewAuditRepository(db)
	tuc := usecases.NewTokenUseCase(tr, ar)
	tokensHandler := handlers.NewTokensHandler(tuc, client)

	go worker.StartWorker(redis, tuc)

	http.Handle("/tokens", tokensHandler)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	server := &http.Server{Addr: portListener}

	go func() {
		log.Info().Msg("API started on " + portListener)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Failed to start HTTP server")
		}
	}()

	<-stop
	log.Info().Msg("Shutting down server...")

	if err := server.Close(); err != nil {
		log.Error().Err(err).Msg("Error shutting down server")
	}
	log.Info().Msg("Server stopped gracefully")

}
