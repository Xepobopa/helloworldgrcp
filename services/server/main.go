package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Xepobopa/helloworldgrcp/services/genproto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}


func (s *server) SayHello(_ context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.Reply{Message: "Hello, " + in.GetName() + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
