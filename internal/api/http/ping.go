package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func PingRoute(r *gin.Engine) {
	r.GET("/ping", ping)
}
