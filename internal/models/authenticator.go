package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	FindByEmail(email string) []User
}

type Authenticator struct {
	UserRepository UserRepository
}

func (auth *Authenticator) HasUser(email string) (bool, []User) {
	user := auth.UserRepository.FindByEmail(email)

	return len(user) != 0, user
}

func (auth *Authenticator) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (auth *Authenticator) CreateNewJWT(email string) (string, error) {
	secretKey := getSecretKey()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		log.Printf(fmt.Sprintf("Error new token: %s", err))
		return "", err
	}

	return tokenString, nil
}

func (auth *Authenticator) VerifyJWT(tokenString string) error {
	secretKey := getSecretKey()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func getSecretKey() []byte {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	envFile := filepath.Join(dir, ".env")

	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
		return nil
	}

	err = godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
		return nil
	}

	return []byte(os.Getenv("SECRET_KEY"))
}
