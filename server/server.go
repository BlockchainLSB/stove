package server

import (
	"strconv"

	"github.com/labstack/echo"
)

// ProtocolType describes a type of protocol.
type ProtocolType int

const (
	HTTP ProtocolType = iota
	gRPC
)

// ServerHandle acts behaviors.
type ServerHandle struct {
}

// ServerBuilder provides information to build a server.
type ServerBuilder interface {
	Name() string
	Protocol() ProtocolType
	Handlers()
}

// Server serves server applications.
type Server struct {
	*echo.Echo
	Port int
}

// New create a new Server.
func New() *Server {
	return &Server{
		Echo: echo.New(),
	}
}

// Start start a new server.
func (s *Server) Start() error {
	return s.Echo.Start(":" + strconv.Itoa(s.Port))
}
