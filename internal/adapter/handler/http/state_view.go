package http

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) stateIDsView(c *gin.Context) {

	states, err := s.stateSVC.GetAllStateInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.HTML(
		http.StatusOK,
		"state_list.go.tpl.html",
		gin.H{
			"title":  "States",
			"states": states,
		},
	)
}

func (s *Server) stateView(c *gin.Context) {
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

	c.HTML(
		http.StatusOK,
		"state_view.go.tpl.html",
		gin.H{
			"title": "State",
			"id":    id,
			"state": string(state),
		},
	)
}
