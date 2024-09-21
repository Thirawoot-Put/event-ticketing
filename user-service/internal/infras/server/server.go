package server

import (
	"github.com/Thirawoot-Put/event-ticketing/user-service/internal/infras/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	app *gin.Engine
}

func AppServer() *Server {
	ginApp := gin.Default()

	return &Server{
		app: ginApp,
	}
}

func (s *Server) HttpListenPort(port string) {
	err := s.app.Run(port)

	if err != nil {
		panic("Failed to start user service")
	}
}

func (s *Server) Start(port string) {

	db.Connect()
	s.HttpListenPort(":" + port)
}
