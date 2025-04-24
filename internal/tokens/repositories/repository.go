package repositories

import (
	"best-practices-golang/internal/tokens/domain"
	"best-practices-golang/pkg/logger"
	"context"
	"database/sql"
	"time"
)

type TokenRepositoryProvider interface {
	UpdateToken(ctx context.Context, tk *domain.Token) error
}

type TokenRepository struct {
	DB *sql.DB
}

func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{
		DB: db,
	}
}

func (tr *TokenRepository) UpdateToken(ctx context.Context, tk *domain.Token) error {
	log := logger.NewLogger()

	token := tk.Value
	now := time.Now()
	log.Info().Msg(token)

	var (
		query = `
  UPDATE dispositivo
  SET deleted_at = ?, count_updated_token = count_updated_token + 1
  WHERE fcm_token = ?
 `
	)

	log.Info().Msg(query)
	result, err := tr.DB.ExecContext(ctx,
		query,
		now,
		token)

	if err != nil {
		log.Error().Err(err).Msg("Failed to update token")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		log.Warn().Msg("No rows updated")
		return nil
	}

	log.Info().Msg("Token updated successfully")
	return err

}
