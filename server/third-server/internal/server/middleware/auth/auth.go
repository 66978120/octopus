package auth

import (
	"context"
	"net/http"
	commctx "server/common/context"
	"server/common/errors"
	"strings"

	oserver "github.com/go-oauth2/oauth2/v4/server"

	"github.com/go-kratos/kratos/v2/middleware"
	kratosHttp "github.com/go-kratos/kratos/v2/transport/http"
)

// Option is HTTP logging option.
type Option func(*Options)

type Options struct {
	NoAuthUris []string
	Server     *oserver.Server
}

const (
	AUTHORIZATION      = "Authorization"
	AUTHORIZATION_TYPE = "Bearer"
)

// Server is an server logging middleware.
func Server(opts ...Option) middleware.Middleware {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var request *http.Request
			if info, ok := kratosHttp.FromServerContext(ctx); ok {
				request = info.Request
			} else {
				return handler(ctx, req)
			}

			needAuth := true
			for _, i := range options.NoAuthUris {
				if strings.Contains(request.RequestURI, i) {
					needAuth = false
					break
				}
			}

			if needAuth {
				authorization := request.Header.Get(AUTHORIZATION)
				if authorization == "" {
					return nil, errors.Errorf(nil, errors.ErrorNotAuthorized)
				}

				if strings.Index(authorization, AUTHORIZATION_TYPE) != 0 || authorization[0:len(AUTHORIZATION_TYPE)] != AUTHORIZATION_TYPE {
					return nil, errors.Errorf(nil, errors.ErrorTokenInvalid)
				}

				tokenInfo, err := options.Server.Manager.LoadAccessToken(ctx, authorization[len(AUTHORIZATION_TYPE)+1:])
				if err != nil {
					return nil, err
				}
				ctx = commctx.PlatformIdToContext(ctx, tokenInfo.GetClientID())
			}

			return handler(ctx, req)
		}
	}
}
