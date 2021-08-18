package internal

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Server gin.Engine
}

func (s *Server) Run() error {
	err := s.Server.Run(":8000")
	return err
}
