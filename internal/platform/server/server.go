package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/JoseUgal/go-http-api/internal/platform/server/handler/courses"
	"github.com/JoseUgal/go-http-api/internal/platform/server/handler/health"
	"github.com/JoseUgal/go-http-api/kit/command"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine			*gin.Engine
	httpAddr		string

	shutdownTimeout time.Duration

	// deps
	commandBus 		command.Bus
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, commandBus command.Bus ) (context.Context, Server) {
	srv := Server{
		engine: gin.Default(), // GIN with middleware [logging, recovery]
		httpAddr: fmt.Sprintf("%v:%v", host, port),

		shutdownTimeout: shutdownTimeout,

		commandBus: commandBus,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}


// Method to register all API routes
func (s *Server) registerRoutes() {
	//s.engine.Use(recovery.Middleware(), logging.Middleware())

	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.commandBus))
}

// Method to start server
func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running", s.httpAddr)

	// Implements this method to allow Graceful shutdown
	// using GIN Framework.
	srv := &http.Server{
		Addr: s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done() // SI el contexto se cancela se desbloqueará y pasrá lo siguiente

	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
