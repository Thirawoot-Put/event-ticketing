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

	// r := gin.Default()
	// r.GET("/healthz-check", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "App health OK",
	// 	})
	// })

	s.app.GET("/healthz-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "App health OK",
		})
	})

	s.HttpListenPort(url)
}
