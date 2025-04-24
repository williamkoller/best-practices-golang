package audits_domain

import (
	"errors"
	"strings"
)

type Audit struct {
	Token  string
	Reason string
}

func NewAudit(token, reason string) (*Audit, error) {
	if strings.TrimSpace(token) == "" {
		return nil, errors.New("token cannot be empty")
	}
	if strings.TrimSpace(reason) == "" {
		return nil, errors.New("reason cannot be empty")
	}
	return &Audit{Token: token, Reason: reason}, nil
}
