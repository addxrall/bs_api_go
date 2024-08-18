package services

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/addxrall/bs_api_go/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = os.Getenv("JWT")

var q *db.Queries

func InitServices(dbtx db.DBTX) {
	q = db.New(dbtx)
}

func Register(c *fiber.Ctx) error {
	var data struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if data.Username == "" || data.Email == "" || data.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing fields"})
	}

	existingUserEmail, err := q.FindUserByEmail(context.Background(), data.Email)
	if err == nil && existingUserEmail.Email != "" {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "Email already in use"})
	}

	existingUserUsername, err := q.FindUserByUsername(context.Background(), data.Username)
	if err == nil && existingUserUsername.Username != "" {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "Username already in use"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	newUser, err := q.CreateUser(context.Background(), db.CreateUserParams{
		Username: data.Username,
		Email:    data.Email,
		Password: string(hashedPassword),
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": newUser.UserID,
		"email":   newUser.Email,
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create token"})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(2 * time.Hour),
	})

	return c.Status(http.StatusOK).JSON(fiber.Map{"token": tokenString})
}

func Login(c *fiber.Ctx) error {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if data.Email == "" || data.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing fields"})
	}

	user, err := q.FindUserByEmail(context.Background(), data.Email)
	if err != nil || user.Email == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"email":   user.Email,
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create token"})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(2 * time.Hour),
	})

	return c.Status(http.StatusOK).JSON(fiber.Map{"token": tokenString})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-(time.Hour * 2)),
	})
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Logout Successful"})
}

func Session(c *fiber.Ctx) error {
	tokenString := c.Cookies("token")
	if tokenString == "" {
		return c.Status(http.StatusOK).JSON(nil)
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(http.StatusNotFound).JSON(nil)
	}

	return c.Status(http.StatusOK).JSON(claims)
}
