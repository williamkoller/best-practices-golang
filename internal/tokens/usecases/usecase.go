package usecases

import (
	auditsdomain "best-practices-golang/internal/audits/domain"
	auditsrepositories "best-practices-golang/internal/audits/repositories"
	"best-practices-golang/internal/tokens/domain"
	"best-practices-golang/internal/tokens/handlers/request"
	"best-practices-golang/internal/tokens/repositories"
	"context"
	"github.com/rs/zerolog"
)

type TokenUseCaseProvider interface {
	Execute(ctx context.Context, rt request.TokenRequest) error
}

type TokenUseCase struct {
	tr  repositories.TokenRepositoryProvider
	ar  auditsrepositories.AuditRepositoryProvider
	Log zerolog.Logger
}

func NewTokenUseCase(tr repositories.TokenRepositoryProvider, ar auditsrepositories.AuditRepositoryProvider) *TokenUseCase {
	return &TokenUseCase{
		tr: tr,
		ar: ar,
	}
}

func (tuc *TokenUseCase) Execute(ctx context.Context, rt request.TokenRequest) error {
	token, err := domain.NewToken(rt.Token)

	if err != nil {
		tuc.Log.Error().Err(err).Msg("[UseCase:Execute] Failed to create token domain object")
		return err
	}

	audit, err := auditsdomain.NewAudit(token.Value, "Token update failed")

	if err != nil {
		tuc.Log.Error().Err(err).Msg("[UseCase:Execute] Failed to create audit domain object")
		return err
	}

	if err := tuc.tr.UpdateToken(ctx, token); err != nil {
		err := tuc.ar.LogFailure(ctx, audit)
		if err != nil {
			return err
		}
		tuc.Log.Error().Err(err).Msg("[UseCase:Execute] Failed to update token in repository")
		return err
	}

	tuc.Log.Info().Str("token", rt.Token).Msg("[UseCase:Execute] Token successfully updated")

	return nil
}
