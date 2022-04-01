package service

import (
	"server/grampus-server/internal/conf"
	"server/grampus-server/internal/data"

	"git.openi.org.cn/OpenI/ClioneLimacina/server/adapter/api/v1"
)

type Service struct {
	v1.UnimplementedAdapterServer
	Data *data.Data
}

func NewService(conf *conf.Bootstrap, data *data.Data) v1.AdapterServer {
	service := &Service{
		Data: data,
	}
	return service
}
