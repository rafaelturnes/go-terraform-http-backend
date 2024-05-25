package http

import (
	"go-terraform-http-backend/internal/core/port"

	"github.com/gin-gonic/gin"
)

type Server struct {
	http     *gin.Engine
	stateSVC port.StateService
}

func NewServer(stateSVC port.StateService) *Server {
	r := gin.Default()

	return &Server{
		http:     r,
		stateSVC: stateSVC,
	}
}

func (s *Server) Routes() {

	// state
	// fetched via GET, updated via POST, and purged with DELETE
	s.http.GET("/state/:id", s.stateFetch)
	s.http.POST("/state/:id", s.stateUpdate)
	s.http.DELETE("/state/:id", s.statePurge)

	// state locking
	s.http.Handle("LOCK", "/state/:id", s.stateLock)
	s.http.Handle("UNLOCK", "/state/:id", s.stateUnlock)

}

func (s *Server) Run() {
	s.http.Run(":3000")
}
