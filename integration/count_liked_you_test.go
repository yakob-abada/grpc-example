package integration

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func TestCountLikedYou(t *testing.T) {
	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	fmt.Println("Connected!")

	c := pb.NewExploreServiceClient(conn)
	t.Run("success", func(t *testing.T) {
		response, err := c.CountLikedYou(context.TODO(), &pb.CountLikedYouRequest{
			RecipientUserId: "1",
		})

		assert.Nil(t, err)
		assert.Equal(t, uint64(1), response.Count)
	})

	t.Run("success no match for given user", func(t *testing.T) {
		response, err := c.CountLikedYou(context.TODO(), &pb.CountLikedYouRequest{
			RecipientUserId: "2",
		})

		assert.Nil(t, err)
		assert.Equal(t, uint64(0), response.Count)
	})
}
