package main

import (
	"fmt"
	"net"
	"log"
	"google.golang.org/grpc"
	pb "github.com/Xepobopa/helloworldgrcp/services/genproto"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if (err != nil) {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()


}