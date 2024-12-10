package main

import (
	"context"
	"fmt"
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	fmt.Println("Connected!")

	c := pb.NewExploreServiceClient(conn)

	response, err := c.CountLikedYou(context.Background(), &pb.CountLikedYouRequest{})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Println("count is", response.Count)
}
