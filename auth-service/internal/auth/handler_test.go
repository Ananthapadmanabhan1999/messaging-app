package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock AuthService for testing
type MockAuthService struct{}

func (m *MockAuthService) Authenticate(username, password string) (string, error) {
    if username == "validUser" && password == "validPass" {
        return "mockToken123", nil
    }
    return "", errors.New("invalid credentials")
}

func TestLoginHandler(t *testing.T) {
    mockService := NewAuthService()
    mockService.users = map[string]string{
        "validUser": "validPass",
    }

    handler := LoginHandler(mockService)

    t.Run("Successful Login", func(t *testing.T) {
        reqBody := LoginRequest{
            Username: "validUser",
            Password: "validPass",
        }
        body, _ := json.Marshal(reqBody)
        req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
        req.Header.Set("Content-Type", "application/json")
        rec := httptest.NewRecorder()

        handler.ServeHTTP(rec, req)

        if rec.Code != http.StatusOK {
            t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
        }

        var resp LoginResponse
        if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
            t.Fatalf("failed to decode response: %v", err)
        }

        if resp.Token == "" {
            t.Errorf("expected a token, got an empty string")
        }
    })
}