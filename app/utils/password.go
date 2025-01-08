package utils

import (
	"errors"

	"github.com/go-passwd/validator"
)

func ValidatePassword(password string) error {
	passwordValidator := validator.New(
		validator.MinLength(8, errors.New("password should be min. 8 characters")),
		validator.MaxLength(20, errors.New("password should be max. 20 characters")),
		validator.ContainsAtLeast("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1, errors.New("password should include at least one uppercase letter")),
		validator.ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 1, errors.New("password should include at least one lowercase letter")),
		validator.ContainsAtLeast("0123456789", 1, errors.New("password should include at least one number")))
	err := passwordValidator.Validate(password)

	return err
}
