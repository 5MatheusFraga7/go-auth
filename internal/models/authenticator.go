package models

import (
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	FindByEmail(email string) []User
}

type Authenticator struct {
	UserRepository UserRepository
}

func (auth *Authenticator) HasUser(email string) bool {
	user := auth.UserRepository.FindByEmail(email)

	return len(user) != 0
}

func (auth *Authenticator) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (auth *Authenticator) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
