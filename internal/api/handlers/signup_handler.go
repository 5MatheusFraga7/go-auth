package handlers

import (
	"encoding/json"
	"go-auth/internal/db/adapters"
	"go-auth/internal/db/repository"
	"go-auth/internal/models"
	"net/http"
)

type SignupHandler struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	ConfirmationPassword string `json:"confirmation_password"`
}

func (sh *SignupHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var data models.UserParams
	userRepository := repository.NewUserRepository(adapters.NewPostgreSQLAdapter())

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		errorResponse(err.Error(), w)
		return
	}

	validator := models.UserSignupValidator{data, userRepository}
	valid, errorMessage := validator.ValidateUserParams()

	if !valid {
		errorResponse(errorMessage, w)
		return
	}

	userRepository.Create(data)

	dataResponse := map[string]interface{}{
		"Status": "Success",
		"User":   data,
	}

	jsonData, _ := json.Marshal(dataResponse)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
func errorResponse(errorMessage string, w http.ResponseWriter) {
	dataResponse := map[string]interface{}{
		"Error": errorMessage,
	}

	jsonData, _ := json.Marshal(dataResponse)

	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonData)
}
