// Let's try something completely different

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: protos/todo.proto

package protos

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
	TodoService_List_FullMethodName   = "/protos.TodoService/List"
	TodoService_Insert_FullMethodName = "/protos.TodoService/Insert"
	TodoService_Find_FullMethodName   = "/protos.TodoService/Find"
	TodoService_Update_FullMethodName = "/protos.TodoService/Update"
	TodoService_Delete_FullMethodName = "/protos.TodoService/Delete"
)

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	List(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	Insert(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	Find(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	Update(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	Delete(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) List(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, TodoService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Insert(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, TodoService_Insert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Find(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, TodoService_Find_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Update(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, TodoService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Delete(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, TodoService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	List(context.Context, *TodoRequest) (*TodoResponse, error)
	Insert(context.Context, *TodoRequest) (*TodoResponse, error)
	Find(context.Context, *TodoRequest) (*TodoResponse, error)
	Update(context.Context, *TodoRequest) (*TodoResponse, error)
	Delete(context.Context, *TodoRequest) (*TodoResponse, error)
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) List(context.Context, *TodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTodoServiceServer) Insert(context.Context, *TodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedTodoServiceServer) Find(context.Context, *TodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedTodoServiceServer) Update(context.Context, *TodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTodoServiceServer) Delete(context.Context, *TodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).List(ctx, req.(*TodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_Insert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Insert(ctx, req.(*TodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_Find_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Find(ctx, req.(*TodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Update(ctx, req.(*TodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Delete(ctx, req.(*TodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TodoService_List_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _TodoService_Insert_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _TodoService_Find_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TodoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TodoService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/todo.proto",
}
