package utils

import "regexp"

func IsValidEmail(email string) bool {
	// Add email validation logic using regex or a validation package
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}
