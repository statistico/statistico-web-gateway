package rest_test

import (
	"github.com/statistico/statistico-web-gateway/internal/app/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCorsMiddleware(t *testing.T) {
	t.Run("applies response headers when receiving OPTIONS request", func(t *testing.T) {
		t.Helper()

		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := w.Header().Get("Access-Control-Allow-Origin")
			headers := w.Header().Get("Access-Control-Allow-Headers")
			methods := w.Header().Get("Access-Control-Allow-Methods")

			assert.Equal(t, "*", origin)
			assert.Equal(t, "Accept, Content-Type, Content-Length", headers)
			assert.Equal(t, "*", methods)
		})

		corsMiddleware := rest.CorsMiddleware(nextHandler)

		request := httptest.NewRequest("OPTIONS", "https://test.com", nil)

		corsMiddleware.ServeHTTP(httptest.NewRecorder(), request)
	})

	t.Run("does not apply header if request received is not an OPTIONS request", func(t *testing.T) {
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := w.Header().Get("Access-Control-Allow-Origin")
			headers := w.Header().Get("Access-Control-Allow-Headers")
			methods := w.Header().Get("Access-Control-Allow-Methods")

			assert.Equal(t, "", origin)
			assert.Equal(t, "", headers)
			assert.Equal(t, "", methods)
		})

		corsMiddleware := rest.CorsMiddleware(nextHandler)

		request := httptest.NewRequest("GET", "https://test.com", nil)

		corsMiddleware.ServeHTTP(httptest.NewRecorder(), request)
	})
}
