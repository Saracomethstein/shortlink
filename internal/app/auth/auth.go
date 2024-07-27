package auth

import (
	"crypto/sha256"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/teris-io/shortid"
	"net/http"
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

func CheckAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session_id")
		if err != nil || cookie.Value == "" {
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}
		return next(c)
	}
}

func GenerateSessionID() string {
	sessionID, err := shortid.Generate()
	if err != nil {
		return ""
	}
	return sessionID
}
