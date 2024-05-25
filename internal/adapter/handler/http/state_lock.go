package http

import (
	"errors"
	"fmt"
	"go-terraform-http-backend/internal/adapter/storage"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func (s *Server) stateLock(c *gin.Context) {
	reqDump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("REQUEST:\n%s\n", string(reqDump))

	id := c.Param("id")

	err = s.stateSVC.Lock(id)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrNotExists):
			c.Status(http.StatusLocked)
		case errors.Is(err, storage.ErrAlreadyLocked):
			c.Status(http.StatusConflict)
		}
		return
	}

	// locked and state info saved
	c.Status(http.StatusOK)
}

func (s *Server) stateUnlock(c *gin.Context) {
	reqDump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("REQUEST:\n%s\n", string(reqDump))

	id := c.Param("id")

	err = s.stateSVC.Unlock(id)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrNotExists):
			c.Status(http.StatusNotFound)
		case errors.Is(err, storage.ErrAlreadyUnlocked):
			c.Status(http.StatusOK)
		}
		return
	}

	c.Status(http.StatusOK)
}
