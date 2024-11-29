package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {

	t.Run("Status Code is correct", func(t *testing.T) {
		r := SetupRouter()
		PingRoute(r)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ping", nil)
		r.ServeHTTP(w, req)

		if http.StatusOK != w.Code {
			t.Errorf("Expected %v, received %v", http.StatusOK, w.Code)
		}
	})

	t.Run("Received correct response in body", func(t *testing.T) {
		r := SetupRouter()
		PingRoute(r)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ping", nil)
		r.ServeHTTP(w, req)

		if "pong" != w.Body.String() {
			t.Errorf("Expected %v, received %v", http.StatusOK, w.Code)
		}
	})

	// assert.Equal(t, http.StatusOK, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}
