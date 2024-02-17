package routes

import (
	"golang.org/x/net/context"

	orm "github.com/sudora1n/google-auth-bot/internal/microservice-api/database"
	"github.com/sudora1n/google-auth-bot/internal/microservice-api/proto"
)

type UserServer struct {
	proto.UnimplementedUsersServer
	ORMFunctions *orm.ORMFunctions
}

func (s *UserServer) CreateOrReturnUser(ctx context.Context, in *proto.CreateOrReturnUserRequest) (*proto.CreateOrReturnUserResponse, error) {
	user, err := s.ORMFunctions.CreateOrReturnUserByUserId(in.UserId)
	if err != nil {
		return nil, err
	}

	return &proto.CreateOrReturnUserResponse{
		Response: &proto.UserObject{
			UserId: user.Id,
			Lang:   user.LanguageISO,
		},
	}, nil
}

func (s *UserServer) ChangeLang(ctx context.Context, in *proto.ChangeLangRequest) (*proto.ChangeLangResponse, error) {
	err := s.ORMFunctions.ChangeLanguageByUserId(
		in.UserId,
		in.Lang,
	)
	if err != nil {
		return nil, err
	}

	return &proto.ChangeLangResponse{Status: true}, nil
}
