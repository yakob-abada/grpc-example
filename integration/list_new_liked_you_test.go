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

func TestListNewLikedYou(t *testing.T) {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	fmt.Println("Connected!")

	c := pb.NewExploreServiceClient(conn)
	t.Run("first page", func(t *testing.T) {
		response, err := c.ListNewLikedYou(context.TODO(), &pb.ListLikedYouRequest{
			RecipientUserId: "1",
		})

		assert.Nil(t, err)
		assert.Equal(t, 2, len(response.Likers))
		assert.NotNil(t, response.NextPaginationToken)
	})
}
