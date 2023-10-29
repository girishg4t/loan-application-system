package middleware

import (
	"net/http"
	"strings"

	"github.com/loan-application-system/pkg/model"
)

const (
	apiKeyHeaderName = "X-API-KEY"
)

type AuthProvider struct {
	Config model.Config
}

func NewAuthProvider(config model.Config) AuthProvider {
	return AuthProvider{
		Config: config,
	}
}

// Authenticate check if frontend application is valid
func (ap *AuthProvider) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var apiKeyValue = r.Header.Get(apiKeyHeaderName)
		if len(strings.TrimSpace(apiKeyValue)) == 0 {
			http.Error(w, "api key is not present", http.StatusUnauthorized)
			return
		}

		if apiKeyValue != ap.Config.ApiKey {
			http.Error(w, "api key is not valid", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}
