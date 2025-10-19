package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type user struct {
	ID    string
	Email string
	Name  string
}

type User interface {
	IsAuthenticated() func(http.Handler) http.Handler
}

var (
	ErrMissingUserID = errors.New("missing user id")
)

func NewUser() *user {
	return &user{}
}

func UserFromRequest(r *http.Request) (*user, error) {
	url := os.Getenv("AUTH_SERVER")
	keyset, err := jwk.Fetch(r.Context(), fmt.Sprintf("%s/api/auth/jwks", url))
	if err != nil {
		slog.Error("fetch jwk", "error", err)
		return &user{}, fmt.Errorf("fetch jwk: %w", err)
	}

	token, err := jwt.ParseRequest(r, jwt.WithKeySet(keyset))
	if err != nil {
		slog.Error("parse request", "error", err)
		return &user{}, fmt.Errorf("parse request: %w", err)
	}

	userID, exists := token.Subject()
	if !exists {
		return &user{}, ErrMissingUserID
	}

	var email string
	var name string

	token.Get("email", &email)
	token.Get("name", &name)

	return &user{
		ID:    userID,
		Email: email,
		Name:  name,
	}, nil
}

func (u *user) IsAuthenticated() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var err error
			u, err = UserFromRequest(r)
			if err != nil {
				slog.Error("error getting user from request", "error", err)
				http.Error(w, err.Error(), http.StatusUnauthorized)
				_ = json.NewEncoder(w).Encode(map[string]string{"message": "unauthorized", "error": err.Error()})
				return
			}
			slog.Info("user authenticated", "user", u)
			ctx := r.Context()
			ctx = context.WithValue(ctx, "user", u)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
