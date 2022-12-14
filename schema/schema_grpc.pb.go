// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: schema.proto

package schema

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

// TodoManagerClient is the client API for TodoManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoManagerClient interface {
	AddTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ReadTodos(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Todos, error)
}

type todoManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoManagerClient(cc grpc.ClientConnInterface) TodoManagerClient {
	return &todoManagerClient{cc}
}

func (c *todoManagerClient) AddTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/schema.TodoManager/AddTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoManagerClient) ReadTodos(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Todos, error) {
	out := new(Todos)
	err := c.cc.Invoke(ctx, "/schema.TodoManager/ReadTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoManagerServer is the server API for TodoManager service.
// All implementations must embed UnimplementedTodoManagerServer
// for forward compatibility
type TodoManagerServer interface {
	AddTodo(context.Context, *Todo) (*emptypb.Empty, error)
	ReadTodos(context.Context, *emptypb.Empty) (*Todos, error)
	mustEmbedUnimplementedTodoManagerServer()
}

// UnimplementedTodoManagerServer must be embedded to have forward compatible implementations.
type UnimplementedTodoManagerServer struct {
}

func (UnimplementedTodoManagerServer) AddTodo(context.Context, *Todo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTodo not implemented")
}
func (UnimplementedTodoManagerServer) ReadTodos(context.Context, *emptypb.Empty) (*Todos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadTodos not implemented")
}
func (UnimplementedTodoManagerServer) mustEmbedUnimplementedTodoManagerServer() {}

// UnsafeTodoManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoManagerServer will
// result in compilation errors.
type UnsafeTodoManagerServer interface {
	mustEmbedUnimplementedTodoManagerServer()
}

func RegisterTodoManagerServer(s grpc.ServiceRegistrar, srv TodoManagerServer) {
	s.RegisterService(&TodoManager_ServiceDesc, srv)
}

func _TodoManager_AddTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).AddTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.TodoManager/AddTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).AddTodo(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoManager_ReadTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).ReadTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.TodoManager/ReadTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).ReadTodos(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoManager_ServiceDesc is the grpc.ServiceDesc for TodoManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schema.TodoManager",
	HandlerType: (*TodoManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTodo",
			Handler:    _TodoManager_AddTodo_Handler,
		},
		{
			MethodName: "ReadTodos",
			Handler:    _TodoManager_ReadTodos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema.proto",
}
