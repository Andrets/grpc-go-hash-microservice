package main

import (
	"log"
	"net"

	app "github.com/Andrets/grpc-go-hash-microservice/internal"
	hash "github.com/Andrets/grpc-go-hash-microservice/internal/hash"
	pb "github.com/Andrets/grpc-go-hash-microservice/pkg/hash_v1"
	"google.golang.org/grpc"
)

func main() {
	app.Init()

	server := grpc.NewServer()

	hashService := hash.NewService()

	pb.RegisterHashServiceServer(server, hashService)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Hash microservice is running on port 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}