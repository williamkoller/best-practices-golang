package audits_repositories

import (
	auditsdomain "best-practices-golang/internal/audits/domain"
	"context"
	"database/sql"
)

type AuditRepositoryProvider interface {
	LogFailure(ctx context.Context, audit *auditsdomain.Audit) error
}

type AuditRepository struct {
	DB *sql.DB
}

func NewAuditRepository(db *sql.DB) *AuditRepository {
	return &AuditRepository{
		DB: db,
	}
}

func (ar *AuditRepository) LogFailure(ctx context.Context, audit *auditsdomain.Audit) error {
	auditDomain := audit
	query := `
		INSERT INTO audit (token, reason)
		VALUES (?, ?)
	`

	_, err := ar.DB.ExecContext(ctx, query, auditDomain, auditDomain)
	if err != nil {
		return err
	}

	return nil
}
