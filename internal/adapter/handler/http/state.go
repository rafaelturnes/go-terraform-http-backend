package http

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func fetch(c *gin.Context) {
	reqDump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("REQUEST:\n%s", string(reqDump))

	c.String(http.StatusOK, "get state", nil)
}

func update(c *gin.Context) {
	reqDump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("REQUEST:\n%s", string(reqDump))

	c.String(http.StatusOK, "update state", nil)
}

func purge(c *gin.Context) {
	reqDump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("REQUEST:\n%s", string(reqDump))

	c.String(http.StatusOK, "delete state", nil)
}
