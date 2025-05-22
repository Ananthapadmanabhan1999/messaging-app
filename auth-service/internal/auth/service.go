package auth

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type User struct {
    Username string
    Password string
}

type AuthService struct {
    users  map[string]string // username -> password
    tokens map[string]string // token -> username
    mu     sync.RWMutex
}

func NewAuthService() *AuthService {
    return &AuthService{
        users: map[string]string{
            "admin": "admin123", // demo
            "user":  "password",
        },
        tokens: make(map[string]string),
    }
}

func (a *AuthService) Authenticate(username, password string) (string, error) {
    a.mu.RLock()
    stored, ok := a.users[username]
    a.mu.RUnlock()

    if !ok || stored != password {
        return "", errors.New("invalid credentials")
    }

    token := uuid.NewString()
    a.mu.Lock()
    a.tokens[token] = username
    a.mu.Unlock()

    return token, nil
}

func (a *AuthService) ValidateToken(token string) (string, bool) {
    a.mu.RLock()
    username, ok := a.tokens[token]
    a.mu.RUnlock()
    return username, ok
}