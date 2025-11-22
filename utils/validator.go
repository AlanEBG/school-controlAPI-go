package utils

import (
	"regexp"
)

// Validar formato de email
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// Validar que la calificación esté en rango válido
func ValidateGrade(grade float64) bool {
	return grade >= 0 && grade <= 100
}

// Validar que el grupo no ande vacío
func ValidateGroup(group string) bool {
	return len(group) > 0 && len(group) <= 10
}

// Validar que tenga nombre xd
func ValidateName(name string) bool {
	return len(name) > 0 && len(name) <= 100
}
