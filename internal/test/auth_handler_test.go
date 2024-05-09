package test

import (
	"bytes"
	"go-auth/internal/api/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticateUser(t *testing.T) {
	requestBody := bytes.NewBuffer([]byte(`{"email": "fragamatheus823@gmail.com", "password": "user123"}`))
	req, err := http.NewRequest("POST", "/signin", requestBody)

	if err != nil {
		t.Fatal(err)
	}

	p := handlers.AuthHandler{}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(p.AuthenticateUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Falha na autenticação do usuário: Código de status incorreto. Recebemos: %v Queríamos: %v", rr.Code, http.StatusOK)
	}
}
