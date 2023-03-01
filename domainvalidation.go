package domainvalidation

import (
	"regexp"
)

// IsValidDomain checks if a domain is valid.
func IsValidDomain(domain string) bool {
	// A valid domain name contains only letters, digits, hyphens, and dots, and does not start or end with a hyphen.
	regex := regexp.MustCompile(`^[a-zA-Z0-9\-]+(\.[a-zA-Z0-9\-]+)*\.[a-zA-Z]{2,}$`)
	return regex.MatchString(domain)
}

// GetValidationResult returns a string representation of the validation result.
func GetValidationResult(valid bool) string {
	if valid {
		return "valid"
	}
	return "invalid"
}
