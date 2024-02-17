package routes

import (
	"golang.org/x/net/context"

	orm "github.com/sudora1n/google-auth-bot/internal/microservice-api/database"
	"github.com/sudora1n/google-auth-bot/internal/microservice-api/proto"
)

type ToTPServer struct {
	proto.UnimplementedToTPsServer
	ORMFunctions *orm.ORMFunctions
}

func (s *ToTPServer) AddToTP(ctx context.Context, in *proto.AddToTPRequest) (*proto.AddToTPResponse, error) {
	err := s.ORMFunctions.AddToTPByUserId(
		in.UserId,
		in.Totp,
		in.Name,
	)
	if err != nil {
		return nil, err
	}

	return &proto.AddToTPResponse{Status: true}, nil
}

func (s *ToTPServer) FindAllToTP(ctx context.Context, in *proto.FindAllToTPRequest) (*proto.FindAllToTPResponse, error) {
	totps, err := s.ORMFunctions.FindAllToTPByUserId(in.UserId)
	if err != nil {
		return nil, err
	}

	var resp []*proto.ToTPObject

	for _, totp := range totps {
		resp = append(
			resp,
			&proto.ToTPObject{
				Id:   totp.Id,
				Name: totp.Name,
			})
	}

	return &proto.FindAllToTPResponse{Response: resp}, nil
}

func (s *ToTPServer) RemoveToTP(ctx context.Context, in *proto.RemoveToTPRequest) (*proto.RemoveToTPResponse, error) {
	err := s.ORMFunctions.RemoveToTPByUserId(
		in.UserId,
		in.Totp,
	)
	if err != nil {
		return nil, err
	}

	return &proto.RemoveToTPResponse{Status: true}, nil
}
