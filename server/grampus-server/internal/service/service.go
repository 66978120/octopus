package service

import (
	api "server/grampus-server/api/v1"
	"server/grampus-server/internal/conf"
	"server/grampus-server/internal/data"
)

type Service struct {
	api.UnimplementedGrampusServiceServer
	Data *data.Data
}

func NewService(conf *conf.Bootstrap, data *data.Data) api.GrampusServiceServer {
	service := &Service{
		Data: data,
	}
	return service
}
