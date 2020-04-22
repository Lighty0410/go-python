package http

import (
	"context"
	"errors"
	api "test/server/grpc/api"
	"test/server/response"

	"github.com/valyala/fasthttp"
)

// Service represents a service for the HTTP-server.
type Service interface {
	GetTopics(word string) *api.Topics
	FilterText(ctx context.Context, textCase string) (map[string]int64, error)
}

// Server represents HTTP server.
type Server struct {
	Ctrl Service
}

// NewHTTPServer creates a new instance of the HTTP-server.
func NewHTTPServer(ctrl Service) Server {
	return Server{
		Ctrl: ctrl,
	}
}

// Run runs HTTP-server.
func (s *Server) Run(ctx *fasthttp.RequestCtx) {
	switch {
	case string(ctx.Path()) == "/test-text" && string(ctx.Method()) == fasthttp.MethodPost:
		s.testText(ctx)
		return
	default:
		response.Incorrect(ctx, fasthttp.StatusBadRequest, errors.New("incorrect endpoint"))
	}
}
