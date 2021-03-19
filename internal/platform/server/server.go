package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	sessionSvc "github.com/vosgaust/voicemod-challenge.git/internal/application/session"
	userSvc "github.com/vosgaust/voicemod-challenge.git/internal/application/user"
	"github.com/vosgaust/voicemod-challenge.git/internal/platform/server/handler/health"
	"github.com/vosgaust/voicemod-challenge.git/internal/platform/server/handler/user"
)

type Server struct {
	address string
	engine  *gin.Engine

	userService    userSvc.UserService
	sessionService sessionSvc.SessionService
}

func New(host string, port uint, userService userSvc.UserService, sessionService sessionSvc.SessionService) Server {
	srv := Server{
		address:        fmt.Sprintf("%s:%d", host, port),
		engine:         gin.New(),
		userService:    userService,
		sessionService: sessionService,
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
	s.engine.PATCH("/user/:user_id", user.UpdateHandler(s.userService))
	s.engine.DELETE("user/:user_id", user.DeleteHandler(s.userService))

	s.engine.POST("/login")

	// TODO: login handler
}
