package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/immrshc/modular-monolith/proto/workflow"
	"github.com/immrshc/modular-monolith/workflow/adapter/registration"
	"github.com/immrshc/modular-monolith/workflow/adapter/repository"
	"github.com/immrshc/modular-monolith/workflow/adapter/rpc/application"
	"github.com/immrshc/modular-monolith/workflow/adapter/transfer"
	"github.com/immrshc/modular-monolith/workflow/usecase"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	tb := transfer.NewClientBuilder(1, 2)
	rb := registration.NewClientBuilder(1, 2)
	server := &application.Server{
		Usecase: usecase.NewTransferApplication(
			repository.NewDB(1),
			tb.NewClient(1),
			rb.NewClient(1),
			repository.NewTxExecutor(1, tb, rb),
		),
	}
	s := grpc.NewServer()
	workflow.RegisterGreeterServer(s, server)
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
