package hash

import (
	"context"

	hash "github.com/Andrets/grpc-go-hash-microservice/pkg/hash_v1"
)

type HashService struct{}

func NewService() *HashService {
	return &HashService{}
}

func (s *HashService) HashPassword(ctx context.Context, req *hash.HashPasswordRequest) (*hash.HashPasswordResponse, error) {
	return &hash.HashPasswordResponse{
		Hashedpassword: "hashed_password_here",
	}, nil
}

func (s *HashService) ValidatePassword(ctx context.Context, req *hash.ValidatePasswordRequest) (*hash.ValidatePasswordResponse, error) {
	return &hash.ValidatePasswordResponse{
		IsValid: true,
	}, nil
}
