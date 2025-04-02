package validations

import (
	"errors"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Valida se uma string contém apenas letras e espaços
func ValidateName(name string) error {
	matched, err := regexp.MatchString(`^[a-zA-Z\s]+$`, name)
	if err != nil || !matched {
		return errors.New("name must contain only letters and spaces")
	}
	return nil
}

// Valida se um e-mail é válido
func ValidateEmail(email string) error {
	err := validate.Var(email, "required,email")
	if err != nil {
		return errors.New("invalid email format")
	}
	return nil
}

// Valida se uma string não está vazia
func ValidateNotEmpty(field, fieldName string) error {
	if strings.TrimSpace(field) == "" {
		return errors.New(fieldName + " cannot be empty")
	}
	return nil
}

// Valida se um valor é positivo
func ValidatePositive(value float64, fieldName string) error {
	if value <= 0 {
		return errors.New(fieldName + " must be a positive value")
	}
	return nil
}
