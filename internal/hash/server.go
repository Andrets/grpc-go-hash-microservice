package hash

import (
	"context"

	pb "github.com/Andrets/grpc-go-hash-microservice/pkg/hash_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	pb.UnimplementedHashServiceServer
	hash Hash
}

type Hash interface {
	HashPassword(
		ctx context.Context,
		password string,
	) (hashedpassword string, err error)
	VerifyPassword(
		ctx context.Context,
		password string,
		hashedpassword string,
	) (IsValid bool, err error)
}

func Register(gRPCServer *grpc.Server, hash Hash) {
	pb.RegisterHashServiceServer(gRPCServer, &serverAPI{hash: hash})
}

func (s *serverAPI) HashPassword(
	ctx context.Context,
	in *pb.HashPasswordRequest,
) (*pb.HashPasswordResponse, error) {
	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password cannot be empty")
	}
	hashedpassword, err := s.hash.HashPassword(ctx, in.GetPassword())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid password")
	}
	return &pb.HashPasswordResponse{Hashedpassword: hashedpassword}, nil
}

// func (s *serverAPI) VerifyPassword(
// 	ctx context.Context,
// 	in *pb.ValidatePasswordRequest,
// ) (*pb.ValidatePasswordResponse, error) {
//   // TODO: implement
// }
