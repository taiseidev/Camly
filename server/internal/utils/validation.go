package utils

import "regexp"

func IsValidEmail(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	// Add email validation logic using regex or a validation package
	return emailRegex.MatchString(email)
}
