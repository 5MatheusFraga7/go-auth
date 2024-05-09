package test

import (
	"go-auth/internal/db/adapters"
	"go-auth/internal/db/repository"
	"go-auth/internal/models"
	"testing"
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

func TestCreateNewToken(t *testing.T) {

}

func TestCheckPassword(t *testing.T) {

}
