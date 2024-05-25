package http

import (
	"fmt"
	"go-terraform-http-backend/internal/core/port"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"

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

	// html
	s.http.HTMLRender = s.loadTemplates("./internal/adapter/template", []string{"debug"})
	//s.http.LoadHTMLGlob("./internal/adapter/template/**/*") // internal/adapter/handler/http internal/adapter/template

	s.http.GET("/debug/state/", s.stateIDsView)
	s.http.GET("/debug/state/:id", s.stateView)

}

func (s *Server) Run() {
	s.http.Run(":3000")
}

func (s *Server) loadTemplates(templatesPath string, pages []string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesPath + "/layout/*.html")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("layouts:", layouts)

	for _, page := range pages {
		fmt.Println("page", fmt.Sprintf("%s/%s/*.hmtl", templatesPath, page))
		includes, err := filepath.Glob(fmt.Sprintf("%s/%s/*.html", templatesPath, page))
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("includes:", includes)

		// Generate our templates map from our layouts/ and includes/ directories
		for _, include := range includes {
			layoutCopy := make([]string, len(layouts))
			copy(layoutCopy, layouts)
			files := append(layoutCopy, include)
			r.AddFromFiles(filepath.Base(include), files...)
		}
	}

	fmt.Println("templates", r)
	return r
}
