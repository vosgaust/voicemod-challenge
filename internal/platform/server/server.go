package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	userSvc "github.com/vosgaust/voicemod-challenge.git/internal/application/user"
	"github.com/vosgaust/voicemod-challenge.git/internal/platform/server/handler/health"
	"github.com/vosgaust/voicemod-challenge.git/internal/platform/server/handler/user"
)

type Server struct {
	address string
	engine  *gin.Engine

	userService userSvc.UserService
}

func New(host string, port uint, userService userSvc.UserService) Server {
	srv := Server{
		address:     fmt.Sprintf("%s:%d", host, port),
		engine:      gin.New(),
		userService: userService,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running at: ", s.address)
	return s.engine.Run(s.address)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())

	s.engine.POST("/user", user.CreateHandler(s.userService))
	s.engine.PATCH("/user/:user_id", user.UpdateHandler())
	s.engine.DELETE("user/:user_id", user.DeleteHandler())

	// TODO: login handler
}
