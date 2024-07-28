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

func (s *AuthService) Authorization(c echo.Context) error {
	var answer bool
	var err error
	var user repositories.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	user = EncryptData(user)
	answer, err = s.userRepo.CheckUserExists(user)

	if err != nil {
		return err
	}

	// TODO: change generate and set session_id
	if answer == true {
		sessionID := GenerateSessionID()
		cookie := new(http.Cookie)
		cookie.Name = "session_id"
		cookie.Value = sessionID
		cookie.Path = "/"
		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, map[string]string{"success": "true", "session_id": sessionID})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Authentication failed"})
}

func (s *AuthService) Registration(c echo.Context) error {
	var user repositories.User
	var err error
	var answer bool

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	answer = CheckCorrectPassword(user.Password)

	if answer != true {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Password not correct"})
	}

	user = EncryptData(user)
	answer, err = s.userRepo.CheckUserExistsByLogin(user.Login)

	if answer == true {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Username is taken"})
	}

	err = s.userRepo.CreateUser(&user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Add user successful"})
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
