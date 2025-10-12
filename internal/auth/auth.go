package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type User struct {
	ID    string
	Email string
	Name  string
}

var (
	ErrMissingUserID = errors.New("missing user id")
)

func UserFromRequest(r *http.Request) (User, error) {
	keyset, err := jwk.Fetch(r.Context(), os.Getenv("AUTH_SERVER"))
	if err != nil {
		return User{}, fmt.Errorf("fetch jwk: %w", err)
	}

	token, err := jwt.ParseRequest(r, jwt.WithKeySet(keyset))
	if err != nil {
		return User{}, fmt.Errorf("parse request: %w", err)
	}

	userID, exists := token.Subject()
	if !exists {
		return User{}, ErrMissingUserID
	}

	var email string
	var name string

	token.Get("email", &email)
	token.Get("name", &name)

	return User{
		ID:    userID,
		Email: email,
		Name:  name,
	}, nil
}
