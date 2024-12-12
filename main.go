package main

import (
	"fmt"
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"github.com/yakob-abada/backend-match/pkg/bootstrap"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	db, err := connectDB()
	if err != nil {
		log.Fatalf("error connecting to database : %v", err)
	}

	grpcServer := grpc.NewServer()

	fmt.Println("starting grpc server")

	pb.RegisterExploreServiceServer(grpcServer, bootstrap.NewExploreServer(db))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func connectDB() (*gorm.DB, error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/dating_app?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}