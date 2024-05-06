package handlers

import (
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
}

func (p *AuthHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	dataResponse := map[string]interface{}{
		"status": "SUCCESS",
	}
	jsonData, _ := json.Marshal(dataResponse)

	// Resposta de exemplo
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
