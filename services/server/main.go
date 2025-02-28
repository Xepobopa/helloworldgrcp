package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/protobuf/types/known/emptypb"
	pb "github.com/Xepobopa/helloworldgrcp/services/genproto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSendBroadcastServer
}


func (s *server) Broadcast(_ context.Context, in *pb.ReqMessage) (*emptypb.Empty, error) {
	log.Printf("Received: %v", in.Username)
	log.Printf("Text: %v", in.Text)
	log.Println("")

	return &emptypb.Empty{}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSendBroadcastServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
