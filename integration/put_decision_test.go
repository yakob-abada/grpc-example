package integration

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

func TestPutDecision(t *testing.T) {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	fmt.Println("Connected!")

	c := pb.NewExploreServiceClient(conn)
	t.Run("success match", func(t *testing.T) {
		response, err := c.PutDecision(context.TODO(), &pb.PutDecisionRequest{
			RecipientUserId: "1",
			ActorUserId:     "2",
			LikedRecipient:  true,
		})

		assert.Nil(t, err)
		assert.Equal(t, true, response.MutualLikes)
	})

	t.Run("success pass", func(t *testing.T) {
		response, err := c.PutDecision(context.TODO(), &pb.PutDecisionRequest{
			RecipientUserId: "1",
			ActorUserId:     "2",
			LikedRecipient:  false,
		})

		assert.Nil(t, err)
		assert.Equal(t, false, response.MutualLikes)
	})
}
