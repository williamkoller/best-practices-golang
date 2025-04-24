package domain

import (
	"errors"
	"strings"
)

type Token struct {
	Value string
}

func NewToken(value string) (*Token, error) {
	if strings.TrimSpace(value) == "" {
		return nil, errors.New("token cannot be empty")
	}
	if len(value) < 10 {
		return nil, errors.New("token too short")
	}
	return &Token{Value: value}, nil
}
