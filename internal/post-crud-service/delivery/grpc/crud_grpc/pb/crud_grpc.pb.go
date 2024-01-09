// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// CrudServiceClient is the client API for CrudService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrudServiceClient interface {
	GetList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PostList, error)
	GetPost(ctx context.Context, in *PostRequestId, opts ...grpc.CallOption) (*Post, error)
	Update(ctx context.Context, in *PostUpdate, opts ...grpc.CallOption) (*Result, error)
	Delete(ctx context.Context, in *PostRequestId, opts ...grpc.CallOption) (*Empty, error)
}

type crudServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrudServiceClient(cc grpc.ClientConnInterface) CrudServiceClient {
	return &crudServiceClient{cc}
}

func (c *crudServiceClient) GetList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PostList, error) {
	out := new(PostList)
	err := c.cc.Invoke(ctx, "/crud_grpc.CrudService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudServiceClient) GetPost(ctx context.Context, in *PostRequestId, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/crud_grpc.CrudService/GetPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudServiceClient) Update(ctx context.Context, in *PostUpdate, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/crud_grpc.CrudService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudServiceClient) Delete(ctx context.Context, in *PostRequestId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/crud_grpc.CrudService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrudServiceServer is the server API for CrudService service.
// All implementations must embed UnimplementedCrudServiceServer
// for forward compatibility
type CrudServiceServer interface {
	GetList(context.Context, *Empty) (*PostList, error)
	GetPost(context.Context, *PostRequestId) (*Post, error)
	Update(context.Context, *PostUpdate) (*Result, error)
	Delete(context.Context, *PostRequestId) (*Empty, error)
	mustEmbedUnimplementedCrudServiceServer()
}

// UnimplementedCrudServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCrudServiceServer struct {
}

func (UnimplementedCrudServiceServer) GetList(context.Context, *Empty) (*PostList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedCrudServiceServer) GetPost(context.Context, *PostRequestId) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedCrudServiceServer) Update(context.Context, *PostUpdate) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCrudServiceServer) Delete(context.Context, *PostRequestId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCrudServiceServer) mustEmbedUnimplementedCrudServiceServer() {}

// UnsafeCrudServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrudServiceServer will
// result in compilation errors.
type UnsafeCrudServiceServer interface {
	mustEmbedUnimplementedCrudServiceServer()
}

func RegisterCrudServiceServer(s grpc.ServiceRegistrar, srv CrudServiceServer) {
	s.RegisterService(&CrudService_ServiceDesc, srv)
}

func _CrudService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud_grpc.CrudService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServiceServer).GetList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrudService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud_grpc.CrudService/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServiceServer).GetPost(ctx, req.(*PostRequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrudService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud_grpc.CrudService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServiceServer).Update(ctx, req.(*PostUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrudService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud_grpc.CrudService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServiceServer).Delete(ctx, req.(*PostRequestId))
	}
	return interceptor(ctx, in, info, handler)
}

// CrudService_ServiceDesc is the grpc.ServiceDesc for CrudService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrudService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crud_grpc.CrudService",
	HandlerType: (*CrudServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _CrudService_GetList_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _CrudService_GetPost_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CrudService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CrudService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/post-crud-service/delivery/grpc/crud_grpc/crud.proto",
}
