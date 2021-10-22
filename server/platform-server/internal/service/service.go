package service

import (
	"server/platform-server/internal/conf"
	"server/platform-server/internal/data"
)

type Service struct {
	Data            *data.Data
	OauthService    OauthService
	TrainJobService TrainJobService
}

func NewService(conf *conf.Bootstrap, data *data.Data) *Service {
	service := &Service{
		OauthService:    NewOauthService(conf, data),
		TrainJobService: NewTrainJobService(conf, data),
	}
	return service
}
