package auth

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
    // Mock AuthService
    mockService := NewAuthService()
    mockService.tokens = map[string]string{
        "validToken": "validUser",
    }

    // Mock next handler
    nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("success"))
    })

    // Middleware with mock service and next handler
    middleware := AuthMiddleware(mockService, nextHandler)

    // Define test cases
    tests := []struct {
        name           string
        authHeader     string
        expectedStatus int
        expectedBody   string
    }{
        {
            name:           "Valid Token",
            authHeader:     "Bearer validToken",
            expectedStatus: http.StatusOK,
            expectedBody:   "success",
        },
        {
            name:           "Missing Token",
            authHeader:     "",
            expectedStatus: http.StatusUnauthorized,
            expectedBody:   "missing token",
        },
        {
            name:           "Invalid Token",
            authHeader:     "Bearer invalidToken",
            expectedStatus: http.StatusUnauthorized,
            expectedBody:   "invalid token",
        },
    }

    // Run test cases in a loop
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            req := httptest.NewRequest(http.MethodGet, "/protected", nil)
            if tc.authHeader != "" {
                req.Header.Set("Authorization", tc.authHeader)
            }
            rec := httptest.NewRecorder()

            middleware.ServeHTTP(rec, req)

            if rec.Code != tc.expectedStatus {
                t.Errorf("expected status %d, got %d", tc.expectedStatus, rec.Code)
            }

            actualBody := strings.TrimSpace(rec.Body.String())
            expectedBody := strings.TrimSpace(tc.expectedBody)
            if actualBody != expectedBody {
                t.Errorf("expected body %q, got %q", expectedBody, actualBody)
            }
        })
    }
}