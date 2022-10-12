package application

import (
	"context"
	"log"

	"github.com/immrshch/modular-monolith/proto/workflow"
	"github.com/immrshch/modular-monolith/workflow/usecase/port"
)

type Server struct {
	workflow.UnimplementedGreeterServer
	Usecase port.ApplicationUsecase
}

func (s *Server) SayHello(ctx context.Context, in *workflow.HelloRequest) (*workflow.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	if err := s.Usecase.ApproveApplication(ctx, 10000); err != nil {
		return nil, err
	}
	return &workflow.HelloReply{Message: "Hello " + in.GetName()}, nil
}
