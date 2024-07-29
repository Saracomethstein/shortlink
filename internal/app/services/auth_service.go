package services

import (
	"crypto/sha256"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/teris-io/shortid"
	"net/http"
	"regexp"
	"shortlink/internal/app/repositories"
)

type IAuthService interface {
	Authorization(c echo.Context) error
	Registration(c echo.Context) error
	EncryptData(user repositories.User) repositories.User
	CheckCorrectPassword(password string) bool
	CheckAuthorization(next echo.HandlerFunc) echo.HandlerFunc
	GenerateSessionID() string
}

type AuthService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Authorization(login, password string) (string, error) {
	var answer bool
	var err error
	var user repositories.User

	user.Login = login
	user.Password = password

	user = EncryptData(user)
	answer, err = s.userRepo.CheckUserExists(user)

	if err != nil {
		return "", err
	}

	if answer == true {
		sessionID := GenerateSessionID()
		return sessionID, nil
	}
	return "", err
}

func (s *AuthService) Registration(login, password string) error {
	var user repositories.User
	var err error
	var answer bool

	user.Login = login
	user.Password = password

	answer = CheckCorrectPassword(user.Password)

	if answer != true {
		return err
	}

	user = EncryptData(user)
	answer, err = s.userRepo.CheckUserExistsByLogin(user.Login)

	if answer == true {
		return nil
	}

	err = s.userRepo.CreateUser(&user)

	if err != nil {
		return err
	}

	return nil
}

func EncryptData(user repositories.User) repositories.User {
	loginHash := sha256.Sum256([]byte(user.Login))
	passwordHash := sha256.Sum256([]byte(user.Password))
	user.Login = fmt.Sprintf("%x", loginHash)
	user.Password = fmt.Sprintf("%x", passwordHash)
	return user
}

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
