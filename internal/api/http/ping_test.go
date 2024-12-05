package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
)

func TestPingRoute(t *testing.T) {
	r := gin.New()
	PingRoute(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	t.Run("sent correct status code", func(t *testing.T) {
		if http.StatusOK != w.Code {
			t.Errorf("Expected %v, received %v", http.StatusOK, w.Code)
		}
	})

	t.Run("sent correct response in body", func(t *testing.T) {
		if "pong" != w.Body.String() {
			t.Errorf("Expected %v, received %v", "pong", w.Body.String())
		}
	})
}
