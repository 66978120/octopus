package grampus

import (
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
)

type grampusService struct {
	api.UnimplementedGrampusServiceServer
	conf *conf.Bootstrap
	data *data.Data
}

func NewGrampusService(conf *conf.Bootstrap, data *data.Data) api.GrampusServiceServer {
	s := &grampusService{
		conf: conf,
		data: data,
	}

	return s
}
