package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func (s *Server) fetch(c *gin.Context) {
	reqDump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("REQUEST:\n%s", string(reqDump))
	id := c.Param("id")

	reader, err := s.stateSVC.Fetch(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	state, err := io.ReadAll(reader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.String(http.StatusOK, string(state))
}

func (s *Server) update(c *gin.Context) {
	reqDump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("REQUEST:\n%s", string(reqDump))

	id := c.Param("id")

	err = s.stateSVC.Update(id, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) purge(c *gin.Context) {
	reqDump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("REQUEST:\n%s", string(reqDump))

	c.String(http.StatusOK, "delete state", nil)
}
