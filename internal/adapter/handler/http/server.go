package http

import "github.com/gin-gonic/gin"

type Server struct {
	http *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()

	return &Server{
		http: r,
	}
}

func (s *Server) Routes() {

	// state
	// fetched via GET, updated via POST, and purged with DELETE
	s.http.GET("/state", fetch)
	s.http.POST("/state", update)
	s.http.DELETE("/state", purge)

	// state locking
	s.http.Handle("LOCK", "/state", lock)
	s.http.Handle("UNLOCK", "/state", unlock)

}

func (s *Server) Run() {
	s.http.Run(":3000")
}
