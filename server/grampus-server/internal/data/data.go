package data

import (
	"context"
	"server/common/errors"
	"server/common/middleware/ctxcopy"
	"server/grampus-server/internal/conf"
	"time"

	api "server/base-server/api/v1"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type Data struct {
	GrampusServiceClient api.GrampusServiceClient
}

func NewData(confData *conf.Data) (*Data, error) {
	baseServerRequestTimeout, err := time.ParseDuration(confData.BaseServerRequestTimeout)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(confData.BaseServerAddr),
		grpc.WithTimeout(baseServerRequestTimeout),
		grpc.WithMiddleware(
			middleware.Chain(
				status.Client(status.WithHandler(errors.ErrorDecode)),
				recovery.Recovery(),
				ctxcopy.Client(),
			),
		),
	)
	if err != nil {
		return nil, err
	}

	return &Data{
		GrampusServiceClient: api.NewGrampusServiceClient(conn),
	}, nil
}