package helper

import (
	"errors"
	"fmt"
	"go-ecommerce-app/internal/domain"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

func SetupAuth(secret string) Auth {
	return Auth{
		Secret: secret,
	}
}

func (a Auth) GenerateHashedPassword(password string) (string, error) {
	if len(password) < 6 {
		return "", errors.New("password length be at least 6 characters long")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("password hash failed")
	}

	return string(hashed), nil
}

func (a Auth) GenerateToken(id uint, email, role string) (string, error) {
	if id == 0 || email == "" || role == "" {
		return "", errors.New("required input are missing")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		return "", errors.New("signed token failed")
	}

	return tokenStr, nil
}

func (a Auth) VerifyPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("password does not match")
	}

	return nil
}

func (a Auth) VerifyToken(t string) (domain.User, error) {
	tokenArr := strings.Split(t, " ")
	if len(tokenArr) != 2 {
		return domain.User{}, nil
	}

	if tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New("invalid token")
	}

	token, err := jwt.Parse(tokenArr[1], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknow signing method: %v", t.Header)
		}
		return []byte(a.Secret), nil
	})
	if err != nil {
		return domain.User{}, errors.New("invalid signing method")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token is expired")
		}

		user := domain.User{}
		user.ID = uint(claims["user_id"].(float64))
		user.Email = claims["email"].(string)
		user.UserType = claims["role"].(string)

		return user, nil
	}

	return domain.User{}, errors.New("token verification failed")
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {
	authHeader := ctx.GetReqHeaders()["Authorization"]

	if authHeader == nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "auth header is missing",
		})
	}

	user, err := a.VerifyToken(authHeader[0])
	if err == nil && user.ID > 0 {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "authorization failed",
			"error":   err,
		})
	}
}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")
	return user.(domain.User)
}

func (a Auth) GenerateCode() (string, error) {
	return RandomNumbers(6)
}