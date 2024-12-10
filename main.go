package main

import (
	"fmt"
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"github.com/yakob-abada/backend-match/pkg/bootstrap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// add a listener address
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("error starting server : %v", err)
	}

	defer lis.Close()

	grpcServer := grpc.NewServer()

	fmt.Println("starting grpc server")

	pb.RegisterExploreServiceServer(grpcServer, bootstrap.NewExploreServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
