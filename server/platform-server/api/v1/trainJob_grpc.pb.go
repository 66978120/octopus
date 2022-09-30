// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TrainJobServiceClient is the client API for TrainJobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainJobServiceClient interface {
	// 创建训练任务
	TrainJob(ctx context.Context, in *TrainJobRequest, opts ...grpc.CallOption) (*TrainJobReply, error)
	//停止训练任务
	StopJob(ctx context.Context, in *StopJobRequest, opts ...grpc.CallOption) (*StopJobReply, error)
	// 获取训练任务详情
	GetJobInfo(ctx context.Context, in *TrainJobInfoRequest, opts ...grpc.CallOption) (*TrainJobInfoReply, error)
	//获取训练任务列表
	TrainJobList(ctx context.Context, in *TrainJobListRequest, opts ...grpc.CallOption) (*TrainJobListReply, error)
	//获取训练任务统计信息
	TrainJobStastics(ctx context.Context, in *TrainJobStasticsRequest, opts ...grpc.CallOption) (*TrainJobStasticsReply, error)
	//获取集群资源信息
	PlatformResources(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PlatformResourcesReply, error)
}

type trainJobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainJobServiceClient(cc grpc.ClientConnInterface) TrainJobServiceClient {
	return &trainJobServiceClient{cc}
}

func (c *trainJobServiceClient) TrainJob(ctx context.Context, in *TrainJobRequest, opts ...grpc.CallOption) (*TrainJobReply, error) {
	out := new(TrainJobReply)
	err := c.cc.Invoke(ctx, "/platformserver.api.v1.TrainJobService/TrainJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainJobServiceClient) StopJob(ctx context.Context, in *StopJobRequest, opts ...grpc.CallOption) (*StopJobReply, error) {
	out := new(StopJobReply)
	err := c.cc.Invoke(ctx, "/platformserver.api.v1.TrainJobService/StopJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainJobServiceClient) GetJobInfo(ctx context.Context, in *TrainJobInfoRequest, opts ...grpc.CallOption) (*TrainJobInfoReply, error) {
	out := new(TrainJobInfoReply)
	err := c.cc.Invoke(ctx, "/platformserver.api.v1.TrainJobService/GetJobInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainJobServiceClient) TrainJobList(ctx context.Context, in *TrainJobListRequest, opts ...grpc.CallOption) (*TrainJobListReply, error) {
	out := new(TrainJobListReply)
	err := c.cc.Invoke(ctx, "/platformserver.api.v1.TrainJobService/TrainJobList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainJobServiceClient) TrainJobStastics(ctx context.Context, in *TrainJobStasticsRequest, opts ...grpc.CallOption) (*TrainJobStasticsReply, error) {
	out := new(TrainJobStasticsReply)
	err := c.cc.Invoke(ctx, "/platformserver.api.v1.TrainJobService/TrainJobStastics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainJobServiceClient) PlatformResources(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PlatformResourcesReply, error) {
	out := new(PlatformResourcesReply)
	err := c.cc.Invoke(ctx, "/platformserver.api.v1.TrainJobService/PlatformResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainJobServiceServer is the server API for TrainJobService service.
// All implementations must embed UnimplementedTrainJobServiceServer
// for forward compatibility
type TrainJobServiceServer interface {
	// 创建训练任务
	TrainJob(context.Context, *TrainJobRequest) (*TrainJobReply, error)
	//停止训练任务
	StopJob(context.Context, *StopJobRequest) (*StopJobReply, error)
	// 获取训练任务详情
	GetJobInfo(context.Context, *TrainJobInfoRequest) (*TrainJobInfoReply, error)
	//获取训练任务列表
	TrainJobList(context.Context, *TrainJobListRequest) (*TrainJobListReply, error)
	//获取训练任务统计信息
	TrainJobStastics(context.Context, *TrainJobStasticsRequest) (*TrainJobStasticsReply, error)
	//获取集群资源信息
	PlatformResources(context.Context, *emptypb.Empty) (*PlatformResourcesReply, error)
	mustEmbedUnimplementedTrainJobServiceServer()
}

// UnimplementedTrainJobServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrainJobServiceServer struct {
}

func (UnimplementedTrainJobServiceServer) TrainJob(context.Context, *TrainJobRequest) (*TrainJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainJob not implemented")
}
func (UnimplementedTrainJobServiceServer) StopJob(context.Context, *StopJobRequest) (*StopJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopJob not implemented")
}
func (UnimplementedTrainJobServiceServer) GetJobInfo(context.Context, *TrainJobInfoRequest) (*TrainJobInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobInfo not implemented")
}
func (UnimplementedTrainJobServiceServer) TrainJobList(context.Context, *TrainJobListRequest) (*TrainJobListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainJobList not implemented")
}
func (UnimplementedTrainJobServiceServer) TrainJobStastics(context.Context, *TrainJobStasticsRequest) (*TrainJobStasticsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainJobStastics not implemented")
}
func (UnimplementedTrainJobServiceServer) PlatformResources(context.Context, *emptypb.Empty) (*PlatformResourcesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlatformResources not implemented")
}
func (UnimplementedTrainJobServiceServer) mustEmbedUnimplementedTrainJobServiceServer() {}

// UnsafeTrainJobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainJobServiceServer will
// result in compilation errors.
type UnsafeTrainJobServiceServer interface {
	mustEmbedUnimplementedTrainJobServiceServer()
}

func RegisterTrainJobServiceServer(s grpc.ServiceRegistrar, srv TrainJobServiceServer) {
	s.RegisterService(&TrainJobService_ServiceDesc, srv)
}

func _TrainJobService_TrainJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainJobServiceServer).TrainJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platformserver.api.v1.TrainJobService/TrainJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainJobServiceServer).TrainJob(ctx, req.(*TrainJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainJobService_StopJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainJobServiceServer).StopJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platformserver.api.v1.TrainJobService/StopJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainJobServiceServer).StopJob(ctx, req.(*StopJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainJobService_GetJobInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainJobInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainJobServiceServer).GetJobInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platformserver.api.v1.TrainJobService/GetJobInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainJobServiceServer).GetJobInfo(ctx, req.(*TrainJobInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainJobService_TrainJobList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainJobListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainJobServiceServer).TrainJobList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platformserver.api.v1.TrainJobService/TrainJobList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainJobServiceServer).TrainJobList(ctx, req.(*TrainJobListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainJobService_TrainJobStastics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainJobStasticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainJobServiceServer).TrainJobStastics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platformserver.api.v1.TrainJobService/TrainJobStastics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainJobServiceServer).TrainJobStastics(ctx, req.(*TrainJobStasticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainJobService_PlatformResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainJobServiceServer).PlatformResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platformserver.api.v1.TrainJobService/PlatformResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainJobServiceServer).PlatformResources(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainJobService_ServiceDesc is the grpc.ServiceDesc for TrainJobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainJobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "platformserver.api.v1.TrainJobService",
	HandlerType: (*TrainJobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TrainJob",
			Handler:    _TrainJobService_TrainJob_Handler,
		},
		{
			MethodName: "StopJob",
			Handler:    _TrainJobService_StopJob_Handler,
		},
		{
			MethodName: "GetJobInfo",
			Handler:    _TrainJobService_GetJobInfo_Handler,
		},
		{
			MethodName: "TrainJobList",
			Handler:    _TrainJobService_TrainJobList_Handler,
		},
		{
			MethodName: "TrainJobStastics",
			Handler:    _TrainJobService_TrainJobStastics_Handler,
		},
		{
			MethodName: "PlatformResources",
			Handler:    _TrainJobService_PlatformResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/platform-server/api/v1/trainJob.proto",
}
