package service

import (
	"context"
	innterapi "server/base-server/api/v1"
	"server/common/errors"
	"server/common/jwt"
	"server/common/log"
	"server/common/utils"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"
	"time"
)

type AuthService struct {
	api.UnimplementedAuthServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewAuthService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.AuthServer {
	return &AuthService{
		conf: conf,
		log:  log.NewHelper("AuthService", logger),
		data: data,
	}
}

func (s *AuthService) GetToken(ctx context.Context, req *api.GetTokenRequest) (*api.GetTokenReply, error) {
	if req.Username == "" {
		return nil, errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	reply, err := s.data.UserClient.FindUser(ctx, &innterapi.FindUserRequest{
		Email: req.Username,
	})
	if err != nil {
		return nil, err
	}
	if reply.User == nil {
		return nil, errors.Errorf(nil, errors.ErrorAuthenticationFailed)
	}
	if reply.User.Status != innterapi.UserStatus_ACTIVITY {
		return nil, errors.Errorf(nil, errors.ErrorAuthenticationForbidden)
	}
	if !utils.ValidatePassword(reply.User.Password, req.Password) {
		return nil, errors.Errorf(err, errors.ErrorAuthenticationFailed)
	}

	token, err := jwt.CreateToken(reply.User.Id, s.conf.Server.Http.JwtSecrect, time.Second*time.Duration(s.conf.Service.TokenExpirationSec))
	if err != nil {
		return nil, err
	}

	return &api.GetTokenReply{
		Token:      token,
		Expiration: s.conf.Service.TokenExpirationSec,
	}, nil
}

func (s *AuthService) DeleteToken(ctx context.Context, req *api.DeleteTokenRequest) (*api.DeleteTokenReply, error) {
	return &api.DeleteTokenReply{}, nil
}
