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

	if auth.HasUser(data.Email) {
		log.Printf(fmt.Sprintf("ACHAMOS O USER!!! " + data.Email))
	} else {
		log.Printf(fmt.Sprintf("USER NOT FOUND :( " + data.Email))
	}

	dataResponse := map[string]interface{}{
		"status": "SUCCESS",
	}

	jsonData, _ := json.Marshal(dataResponse)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
