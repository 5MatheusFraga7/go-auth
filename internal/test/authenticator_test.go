package test

import (
	"fmt"
	"go-auth/internal/db/adapters"
	"go-auth/internal/db/repository"
	"go-auth/internal/models"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHasUser(t *testing.T) {
	userRepository := repository.NewUserRepository(adapters.NewPostgreSQLAdapter())
	auth := models.Authenticator{UserRepository: userRepository}

	emailTest := "fragamatheus823@gmail.com"
	hasUser, user := auth.HasUser(emailTest)

	if !hasUser {
		t.Errorf("Expected find user with email %v got: false", emailTest)
	}
	if len(user) != 1 {
		t.Errorf("Expected to find user with email %s, but no user was returned", emailTest)
	}
	if user[0].Email != emailTest {
		t.Errorf("Expected find user with email %v got: %+v", emailTest, user)
	}
}

func TestCheckPassword(t *testing.T) {
	stringPassword := "123456678"
	hashPassword := hashPassword(stringPassword)
	auth := models.Authenticator{}

	if !auth.CheckPassword(stringPassword, hashPassword) {
		t.Errorf("Failed check Password: Expected password will be valid")
	}
}

func TestCreateNewToken(t *testing.T) {
	auth := models.Authenticator{}
	emailTest := "fragamatheus823@gmail.com"

	token, err := auth.CreateNewJWT(emailTest)
	invalidToken := auth.VerifyJWT(token)

	if err != nil {
		t.Errorf("Fail in CreateNewJWT: We have and error: %v", err)
	} else if invalidToken != nil {
		t.Errorf("Fail in CreateNewJWT: We have and invalid token: %v", token)
	}
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(hashedPassword)
}
