package auth

import (
	"crypto/sha256"
	"fmt"
	"regexp"
)

func EcnryptData(login, password string) (string, string) {
	loginHash := sha256.Sum256([]byte(login))
	passwordHash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", loginHash), fmt.Sprintf("%x", passwordHash)
}

/*
password must meet the following conditions:
- minimum length 10 characters
- include special characters
- include capital letters
- include numbers
*/

func CheckCorrectPassword(password string) bool {
	if len(password) < 10 {
		return false
	}

	uppercasePattern := `[A-Z]`
	matched, _ := regexp.MatchString(uppercasePattern, password)
	if !matched {
		return false
	}

	specialCharPattern := `[!@#$%^&*(),.?":{}|<>]`
	matched, _ = regexp.MatchString(specialCharPattern, password)
	if !matched {
		return false
	}

	digitPattern := `[0-9]`
	matched, _ = regexp.MatchString(digitPattern, password)
	if !matched {
		return false
	}

	return true
}
