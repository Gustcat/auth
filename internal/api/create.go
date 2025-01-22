package api

import (
	"context"
	"github.com/Gustcat/auth/internal/converter"
	desc "github.com/Gustcat/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req.GetPassword() != req.GetPasswordConfirm() {
		return nil, status.Error(codes.InvalidArgument, "password does not match")
	}

	id, err := i.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Created user with id %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
