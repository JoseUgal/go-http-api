package server

import (
	"fmt"
	"log"

	"github.com/JoseUgal/go-http-api/internal/creating"
	"github.com/JoseUgal/go-http-api/internal/platform/server/handler/courses"
	"github.com/JoseUgal/go-http-api/internal/platform/server/handler/health"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine		*gin.Engine
	httpAddr	string

	// deps
	creatingCourseService creating.CourseService
}

func New(host string, port uint, creatingCourseService creating.CourseService) Server {
	srv := Server{
		engine: gin.New(),
		httpAddr: fmt.Sprintf("%v:%v", host, port),

		creatingCourseService: creatingCourseService,
	}

	srv.registerRoutes()
	return srv
}

// Method to start server
func (s *Server) Run() error {
	log.Println("Server running", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

// Method to register all API routes
func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.creatingCourseService))
}