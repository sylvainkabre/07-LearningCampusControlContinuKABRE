package utils

import (
	"errors"
	"regexp"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("Le mot de passe doit contenir au moins 8 caractères")
	}

	var (
		upper   = regexp.MustCompile(`[A-Z]`)
		lower   = regexp.MustCompile(`[a-z]`)
		number  = regexp.MustCompile(`[0-9]`)
		special = regexp.MustCompile(`[!@#~$%^&*()_+|<>?:{}]`)
	)

	if !upper.MatchString(password) {
		return errors.New("Le mot de passe doit contenir au moins une majuscule")
	}

	if !lower.MatchString(password) {
		return errors.New("Le mot de passe doit contenir au moins une minuscule")
	}

	if !number.MatchString(password) {
		return errors.New("Le mot de passe doit contenir au moins un chiffre")
	}

	if !special.MatchString(password) {
		return errors.New("Le mot de passe doit contenir au moins un caractère spécial")
	}
	return nil
}
