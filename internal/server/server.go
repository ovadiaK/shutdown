package server

import (
	"context"
	"fmt"

	"github.com/ovadiaK/shutdown/internal/input"
)

type Server struct {
	in input.Input
}

func New(in input.Input) *Server {
	return &Server{in: in}
}
func (s *Server) Run(ctx context.Context) {

	_ = s.in.GetInts(ctx)
	fmt.Println("over")

}
