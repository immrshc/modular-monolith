package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/immrshch/modular-monolith/proto/workflow"
	"github.com/immrshch/modular-monolith/transfer/adapter/boundary"
	tu "github.com/immrshch/modular-monolith/transfer/usecase"
	"github.com/immrshch/modular-monolith/workflow/adapter/rpc/application"
	"github.com/immrshch/modular-monolith/workflow/adapter/transfer"
	wu "github.com/immrshch/modular-monolith/workflow/usecase"
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
	server := &application.Server{
		Usecase: wu.NewTransferApplication(
			transfer.NewClient(
				boundary.NewTransfer(
					tu.NewAccountTransfer(),
				),
			),
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
