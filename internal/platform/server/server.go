package server

import (
	"fmt"
	"log"

	mooc "github.com/JoseUgal/go-http-api/internal"
	"github.com/JoseUgal/go-http-api/internal/platform/server/handler/courses"
	"github.com/JoseUgal/go-http-api/internal/platform/server/handler/health"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine		*gin.Engine
	httpAddr	string

	// deps
	courseRepository mooc.CourseRepository
}

func New(host string, port uint, courseRepository mooc.CourseRepository) Server {
	srv := Server{
		engine: gin.New(),
		httpAddr: fmt.Sprintf("%v:%v", host, port),

		courseRepository: courseRepository,
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
	s.engine.POST("/courses", courses.CreateHandler(s.courseRepository))
}