package server

import (
	"fmt"
	"net/http"

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

	url := fmt.Sprintf(":%s", port)

	s.app.GET("/healthz-check", healthCheckHandler)

	s.HttpListenPort(url)
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "App healthy")
}
