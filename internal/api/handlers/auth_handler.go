package handlers

import (
	"encoding/json"
	"fmt"
	"go-auth/internal/db/adapters"
	"go-auth/internal/db/repository"
	"go-auth/internal/models"
	"log"
	"net/http"
)

type AuthHandler struct {
	Email    string
	Password string
}

func (p *AuthHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var data AuthHandler

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRepository := repository.NewUserRepository(adapters.NewPostgreSQLAdapter())
	auth := models.Authenticator{UserRepository: userRepository}

	status := ""
	attr := "status"

	hasUser, user := auth.HasUser(data.Email)
	if hasUser && auth.CheckPasswordHash(data.Password, user[0].EncryptedPassword) {
		status, _ = auth.CreateNewJWT(data.Email)
		attr = "token"
	} else {
		status = "invalid credentials"
		log.Printf(fmt.Sprintf("invalid credentials :( " + data.Email))
	}

	dataResponse := map[string]interface{}{
		attr: status,
	}

	jsonData, _ := json.Marshal(dataResponse)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
