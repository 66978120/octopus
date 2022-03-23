package server

import (
	nethttp "net/http"
	"server/common/http/codec"
	"server/common/middleware/logging"

	"server/common/middleware/validate"

	"server/common/errors"
	v1 "server/grampus-server/api/v1"
	"server/grampus-server/internal/conf"

	comhttp "server/common/http"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ResponseErr struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, service v1.GrampusServiceServer) *http.Server {
	var opts = []http.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	options := []http.HandleOption{
		http.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				logging.Server(),
				validate.Server(),
			),
		),
		http.RequestDecoder(comhttp.DecodeRequest),
		http.ResponseEncoder(func(res nethttp.ResponseWriter, req *nethttp.Request, v interface{}) error {
			return codec.EncodeResponse(res, req, v, func(v interface{}) interface{} {
				return v
			})
		}),
		http.ErrorEncoder(func(res nethttp.ResponseWriter, req *nethttp.Request, err error) {
			codec.EncodeError(res, req, err, func(e *errors.OctopusError) interface{} {
				return convertErrCode(e)
			})
		}),
	}

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", v1.NewGrampusServiceHandler(service, options...))
	return srv
}

func convertErrCode(e *errors.OctopusError) ResponseErr {
	var errorCode ErrorCode
	switch e.Code {
	case errors.ErrorInvalidRequestParameter:
		errorCode = ErrorCodeInternalError
	case errors.ErrorAuthenticationFailed:
		errorCode = ErrorAuthenticationFailed
	default:
		errorCode = ErrorCodeInternalError
	}

	return ResponseErr{
		ErrorCode: int(errorCode),
		ErrorMsg:  DefaultErrorMsg[errorCode],
	}
}
