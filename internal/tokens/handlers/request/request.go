package request

import (
	"best-practices-golang/internal/tokens/handlers/request/errors"
	"fmt"
	validate "github.com/go-playground/validator/v10"
)

type TokenRequest struct {
	Token string `json:"token" validate:"required"`
}

func (tr *TokenRequest) Validate() error {
	validator := validate.New(validate.WithRequiredStructEnabled())

	err := validator.Struct(tr)

	if err != nil {
		var fields = make([]string, 0)
		for _, err := range err.(validate.ValidationErrors) {
			fields = append(fields, fmt.Sprintf("%s is %s", err.StructField(), err.ActualTag()))
		}

		return errors.NewBodyError(fields)
	}

	return nil
}
