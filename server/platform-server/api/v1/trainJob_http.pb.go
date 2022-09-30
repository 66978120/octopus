// Code generated by protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(http.Request)
var _ = new(context.Context)
var _ = binding.MapProto
var _ = mux.NewRouter

const _ = http1.SupportPackageIsVersion1

type TrainJobServiceHandler interface {
	GetJobInfo(context.Context, *TrainJobInfoRequest) (*TrainJobInfoReply, error)

	PlatformResources(context.Context, *emptypb.Empty) (*PlatformResourcesReply, error)

	StopJob(context.Context, *StopJobRequest) (*StopJobReply, error)

	TrainJob(context.Context, *TrainJobRequest) (*TrainJobReply, error)

	TrainJobList(context.Context, *TrainJobListRequest) (*TrainJobListReply, error)

	TrainJobStastics(context.Context, *TrainJobStasticsRequest) (*TrainJobStasticsReply, error)
}

func NewTrainJobServiceHandler(srv TrainJobServiceHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/v1/platform/trainjob", func(w http.ResponseWriter, r *http.Request) {
		var in TrainJobRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TrainJob(ctx, req.(*TrainJobRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*TrainJobReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/v1/platform/trainjob/{id}/stop", func(w http.ResponseWriter, r *http.Request) {
		var in StopJobRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.StopJob(ctx, req.(*StopJobRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*StopJobReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/v1/platform/trainjob/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in TrainJobInfoRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetJobInfo(ctx, req.(*TrainJobInfoRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*TrainJobInfoReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/v1/platform/trainjob", func(w http.ResponseWriter, r *http.Request) {
		var in TrainJobListRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TrainJobList(ctx, req.(*TrainJobListRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*TrainJobListReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/v1/overall/trainjobStastics", func(w http.ResponseWriter, r *http.Request) {
		var in TrainJobStasticsRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TrainJobStastics(ctx, req.(*TrainJobStasticsRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*TrainJobStasticsReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/v1/platform/resource", func(w http.ResponseWriter, r *http.Request) {
		var in emptypb.Empty
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PlatformResources(ctx, req.(*emptypb.Empty))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*PlatformResourcesReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	return r
}
