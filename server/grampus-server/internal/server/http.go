package server

import (
	nethttp "net/http"
	"server/common/http/codec"
	"server/common/middleware/logging"

	"server/common/middleware/validate"

	"server/common/errors"
	"server/grampus-server/internal/conf"

	v1 "git.openi.org.cn/OpenI/Grampus/server/adapter/api/v1"

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
func NewHTTPServer(c *conf.Server, service v1.AdapterServer) *http.Server {
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
	srv.HandlePrefix("/", v1.NewAdapterHandler(service, options...))
	return srv
}

var DefaultErrorMsg = map[v1.ErrorCode]string{
	v1.ErrorCode_InternalError:           "internal error",
	v1.ErrorCode_InvalidRequestParameter: "invalid request parameter",
	v1.ErrorCode_AuthenticationFailed:    "authentication failed",
	v1.ErrorCode_ApiNotSupport:           "api not support",
}

func convertErrCode(e *errors.OctopusError) v1.ErrorResponse {
	var errorCode v1.ErrorCode
	switch e.Code {
	case errors.ErrorInvalidRequestParameter:
		errorCode = v1.ErrorCode_InternalError
	case errors.ErrorAuthenticationFailed:
		errorCode = v1.ErrorCode_AuthenticationFailed
	default:
		errorCode = v1.ErrorCode_InternalError
	}

	return v1.ErrorResponse{
		ErrorCode: errorCode,
		ErrorMsg:  DefaultErrorMsg[errorCode],
	}
}
