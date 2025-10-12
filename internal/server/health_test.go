package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alcb1310/bca-auth/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestCheckHealth(t *testing.T) {
	s := server.NewServer()

	testData := []struct {
		name   string
		status int
		body   string
	}{
		{
			name:   "check health",
			status: http.StatusOK,
			body:   "{\"message\":\"It's healthy\"}\n",
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/api/v1/health", nil)
			assert.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()

			s.Router.ServeHTTP(res, req)
			assert.Equal(t, tt.status, res.Code)

			body := res.Body.String()
			assert.Equal(t, tt.body, body)
		})
	}
}
