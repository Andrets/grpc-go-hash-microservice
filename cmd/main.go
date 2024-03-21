package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Andrets/grpc-go-hash-microservice/internal/hash"
	pb "github.com/Andrets/grpc-go-hash-microservice/pkg/hash_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedHashServiceServer
}

func (s *server) HashPassword(ctx context.Context, req *pb.HashPasswordRequest) (*pb.HashPasswordResponse, error) {
	log.Printf("Password: %d", req.GetPassword())
	hashedpassword, err := hash.HashPassword(req.GetPassword())
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}
	return &pb.HashPasswordResponse{
	Hashedpassword: hashedpassword,
	}, nil
}

func (s *server) ValidatePassword(ctx context.Context, req *pb.ValidatePasswordRequest) (*pb.ValidatePasswordResponse, error) {
	log.Printf("Password: %d", req.GetPassword())
	log.Printf("Hashed password: %d", req.GetHashedpassword())

	isValid := hash.ValidatePassword(req.GetPassword(), req.GetHashedpassword())
	fmt.Printf("Is valid: %t\n", isValid)
	return &pb.ValidatePasswordResponse{
		IsValid: isValid,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterHashServiceServer(s, &server{})
	log.Println("Hash microservice is running on port 50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}