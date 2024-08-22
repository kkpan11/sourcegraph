// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: repoupdater.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RepoUpdaterService_RepoUpdateSchedulerInfo_FullMethodName = "/repoupdater.v1.RepoUpdaterService/RepoUpdateSchedulerInfo"
	RepoUpdaterService_EnqueueRepoUpdate_FullMethodName       = "/repoupdater.v1.RepoUpdaterService/EnqueueRepoUpdate"
	RepoUpdaterService_RecloneRepository_FullMethodName       = "/repoupdater.v1.RepoUpdaterService/RecloneRepository"
	RepoUpdaterService_EnqueueChangesetSync_FullMethodName    = "/repoupdater.v1.RepoUpdaterService/EnqueueChangesetSync"
)

// RepoUpdaterServiceClient is the client API for RepoUpdaterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepoUpdaterServiceClient interface {
	// RepoUpdateSchedulerInfo returns information about the state of the repo in the update scheduler.
	RepoUpdateSchedulerInfo(ctx context.Context, in *RepoUpdateSchedulerInfoRequest, opts ...grpc.CallOption) (*RepoUpdateSchedulerInfoResponse, error)
	// EnqueueRepoUpdate requests that the named repository be updated in the near
	// future. It does not wait for the update.
	EnqueueRepoUpdate(ctx context.Context, in *EnqueueRepoUpdateRequest, opts ...grpc.CallOption) (*EnqueueRepoUpdateResponse, error)
	// RecloneRepository reclones the repository on gitserver. This is useful when
	// the repository is likely broken.
	// Note that after this call finished the repository is not immediately available
	// again, so this method should be used with care.
	RecloneRepository(ctx context.Context, in *RecloneRepositoryRequest, opts ...grpc.CallOption) (*RecloneRepositoryResponse, error)
	EnqueueChangesetSync(ctx context.Context, in *EnqueueChangesetSyncRequest, opts ...grpc.CallOption) (*EnqueueChangesetSyncResponse, error)
}

type repoUpdaterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRepoUpdaterServiceClient(cc grpc.ClientConnInterface) RepoUpdaterServiceClient {
	return &repoUpdaterServiceClient{cc}
}

func (c *repoUpdaterServiceClient) RepoUpdateSchedulerInfo(ctx context.Context, in *RepoUpdateSchedulerInfoRequest, opts ...grpc.CallOption) (*RepoUpdateSchedulerInfoResponse, error) {
	out := new(RepoUpdateSchedulerInfoResponse)
	err := c.cc.Invoke(ctx, RepoUpdaterService_RepoUpdateSchedulerInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoUpdaterServiceClient) EnqueueRepoUpdate(ctx context.Context, in *EnqueueRepoUpdateRequest, opts ...grpc.CallOption) (*EnqueueRepoUpdateResponse, error) {
	out := new(EnqueueRepoUpdateResponse)
	err := c.cc.Invoke(ctx, RepoUpdaterService_EnqueueRepoUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoUpdaterServiceClient) RecloneRepository(ctx context.Context, in *RecloneRepositoryRequest, opts ...grpc.CallOption) (*RecloneRepositoryResponse, error) {
	out := new(RecloneRepositoryResponse)
	err := c.cc.Invoke(ctx, RepoUpdaterService_RecloneRepository_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoUpdaterServiceClient) EnqueueChangesetSync(ctx context.Context, in *EnqueueChangesetSyncRequest, opts ...grpc.CallOption) (*EnqueueChangesetSyncResponse, error) {
	out := new(EnqueueChangesetSyncResponse)
	err := c.cc.Invoke(ctx, RepoUpdaterService_EnqueueChangesetSync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepoUpdaterServiceServer is the server API for RepoUpdaterService service.
// All implementations must embed UnimplementedRepoUpdaterServiceServer
// for forward compatibility
type RepoUpdaterServiceServer interface {
	// RepoUpdateSchedulerInfo returns information about the state of the repo in the update scheduler.
	RepoUpdateSchedulerInfo(context.Context, *RepoUpdateSchedulerInfoRequest) (*RepoUpdateSchedulerInfoResponse, error)
	// EnqueueRepoUpdate requests that the named repository be updated in the near
	// future. It does not wait for the update.
	EnqueueRepoUpdate(context.Context, *EnqueueRepoUpdateRequest) (*EnqueueRepoUpdateResponse, error)
	// RecloneRepository reclones the repository on gitserver. This is useful when
	// the repository is likely broken.
	// Note that after this call finished the repository is not immediately available
	// again, so this method should be used with care.
	RecloneRepository(context.Context, *RecloneRepositoryRequest) (*RecloneRepositoryResponse, error)
	EnqueueChangesetSync(context.Context, *EnqueueChangesetSyncRequest) (*EnqueueChangesetSyncResponse, error)
	mustEmbedUnimplementedRepoUpdaterServiceServer()
}

// UnimplementedRepoUpdaterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRepoUpdaterServiceServer struct {
}

func (UnimplementedRepoUpdaterServiceServer) RepoUpdateSchedulerInfo(context.Context, *RepoUpdateSchedulerInfoRequest) (*RepoUpdateSchedulerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepoUpdateSchedulerInfo not implemented")
}
func (UnimplementedRepoUpdaterServiceServer) EnqueueRepoUpdate(context.Context, *EnqueueRepoUpdateRequest) (*EnqueueRepoUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnqueueRepoUpdate not implemented")
}
func (UnimplementedRepoUpdaterServiceServer) RecloneRepository(context.Context, *RecloneRepositoryRequest) (*RecloneRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecloneRepository not implemented")
}
func (UnimplementedRepoUpdaterServiceServer) EnqueueChangesetSync(context.Context, *EnqueueChangesetSyncRequest) (*EnqueueChangesetSyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnqueueChangesetSync not implemented")
}
func (UnimplementedRepoUpdaterServiceServer) mustEmbedUnimplementedRepoUpdaterServiceServer() {}

// UnsafeRepoUpdaterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepoUpdaterServiceServer will
// result in compilation errors.
type UnsafeRepoUpdaterServiceServer interface {
	mustEmbedUnimplementedRepoUpdaterServiceServer()
}

func RegisterRepoUpdaterServiceServer(s grpc.ServiceRegistrar, srv RepoUpdaterServiceServer) {
	s.RegisterService(&RepoUpdaterService_ServiceDesc, srv)
}

func _RepoUpdaterService_RepoUpdateSchedulerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoUpdateSchedulerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoUpdaterServiceServer).RepoUpdateSchedulerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RepoUpdaterService_RepoUpdateSchedulerInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoUpdaterServiceServer).RepoUpdateSchedulerInfo(ctx, req.(*RepoUpdateSchedulerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoUpdaterService_EnqueueRepoUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueRepoUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoUpdaterServiceServer).EnqueueRepoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RepoUpdaterService_EnqueueRepoUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoUpdaterServiceServer).EnqueueRepoUpdate(ctx, req.(*EnqueueRepoUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoUpdaterService_RecloneRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecloneRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoUpdaterServiceServer).RecloneRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RepoUpdaterService_RecloneRepository_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoUpdaterServiceServer).RecloneRepository(ctx, req.(*RecloneRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoUpdaterService_EnqueueChangesetSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueChangesetSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoUpdaterServiceServer).EnqueueChangesetSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RepoUpdaterService_EnqueueChangesetSync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoUpdaterServiceServer).EnqueueChangesetSync(ctx, req.(*EnqueueChangesetSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RepoUpdaterService_ServiceDesc is the grpc.ServiceDesc for RepoUpdaterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RepoUpdaterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "repoupdater.v1.RepoUpdaterService",
	HandlerType: (*RepoUpdaterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RepoUpdateSchedulerInfo",
			Handler:    _RepoUpdaterService_RepoUpdateSchedulerInfo_Handler,
		},
		{
			MethodName: "EnqueueRepoUpdate",
			Handler:    _RepoUpdaterService_EnqueueRepoUpdate_Handler,
		},
		{
			MethodName: "RecloneRepository",
			Handler:    _RepoUpdaterService_RecloneRepository_Handler,
		},
		{
			MethodName: "EnqueueChangesetSync",
			Handler:    _RepoUpdaterService_EnqueueChangesetSync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repoupdater.proto",
}
