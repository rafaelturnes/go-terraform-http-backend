package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func fetch(c *gin.Context) {
	c.String(http.StatusOK, "get state", nil)
}

func update(c *gin.Context) {
	c.String(http.StatusOK, "update state", nil)
}

func purge(c *gin.Context) {
	c.String(http.StatusOK, "delete state", nil)
}
