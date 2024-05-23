package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func lock(c *gin.Context) {
	c.String(http.StatusOK, "lock state", nil)
}

func unlock(c *gin.Context) {
	c.String(http.StatusOK, "lock state", nil)
}
