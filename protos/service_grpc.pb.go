// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: protos/service.proto

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

// InterviewClient is the client API for Interview service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InterviewClient interface {
	OneOnOne(ctx context.Context, in *User, opts ...grpc.CallOption) (*Result, error)
	OnlineScreening(ctx context.Context, in *Batch, opts ...grpc.CallOption) (Interview_OnlineScreeningClient, error)
	FinalRound(ctx context.Context, opts ...grpc.CallOption) (Interview_FinalRoundClient, error)
	CampusDrive(ctx context.Context, opts ...grpc.CallOption) (Interview_CampusDriveClient, error)
}

type interviewClient struct {
	cc grpc.ClientConnInterface
}

func NewInterviewClient(cc grpc.ClientConnInterface) InterviewClient {
	return &interviewClient{cc}
}

func (c *interviewClient) OneOnOne(ctx context.Context, in *User, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/Interview/OneOnOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interviewClient) OnlineScreening(ctx context.Context, in *Batch, opts ...grpc.CallOption) (Interview_OnlineScreeningClient, error) {
	stream, err := c.cc.NewStream(ctx, &Interview_ServiceDesc.Streams[0], "/Interview/OnlineScreening", opts...)
	if err != nil {
		return nil, err
	}
	x := &interviewOnlineScreeningClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Interview_OnlineScreeningClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type interviewOnlineScreeningClient struct {
	grpc.ClientStream
}

func (x *interviewOnlineScreeningClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *interviewClient) FinalRound(ctx context.Context, opts ...grpc.CallOption) (Interview_FinalRoundClient, error) {
	stream, err := c.cc.NewStream(ctx, &Interview_ServiceDesc.Streams[1], "/Interview/FinalRound", opts...)
	if err != nil {
		return nil, err
	}
	x := &interviewFinalRoundClient{stream}
	return x, nil
}

type Interview_FinalRoundClient interface {
	Send(*User) error
	CloseAndRecv() (*ResultReport, error)
	grpc.ClientStream
}

type interviewFinalRoundClient struct {
	grpc.ClientStream
}

func (x *interviewFinalRoundClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *interviewFinalRoundClient) CloseAndRecv() (*ResultReport, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ResultReport)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *interviewClient) CampusDrive(ctx context.Context, opts ...grpc.CallOption) (Interview_CampusDriveClient, error) {
	stream, err := c.cc.NewStream(ctx, &Interview_ServiceDesc.Streams[2], "/Interview/CampusDrive", opts...)
	if err != nil {
		return nil, err
	}
	x := &interviewCampusDriveClient{stream}
	return x, nil
}

type Interview_CampusDriveClient interface {
	Send(*User) error
	Recv() (*Result, error)
	grpc.ClientStream
}

type interviewCampusDriveClient struct {
	grpc.ClientStream
}

func (x *interviewCampusDriveClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *interviewCampusDriveClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InterviewServer is the server API for Interview service.
// All implementations must embed UnimplementedInterviewServer
// for forward compatibility
type InterviewServer interface {
	OneOnOne(context.Context, *User) (*Result, error)
	OnlineScreening(*Batch, Interview_OnlineScreeningServer) error
	FinalRound(Interview_FinalRoundServer) error
	CampusDrive(Interview_CampusDriveServer) error
	mustEmbedUnimplementedInterviewServer()
}

// UnimplementedInterviewServer must be embedded to have forward compatible implementations.
type UnimplementedInterviewServer struct {
}

func (UnimplementedInterviewServer) OneOnOne(context.Context, *User) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OneOnOne not implemented")
}
func (UnimplementedInterviewServer) OnlineScreening(*Batch, Interview_OnlineScreeningServer) error {
	return status.Errorf(codes.Unimplemented, "method OnlineScreening not implemented")
}
func (UnimplementedInterviewServer) FinalRound(Interview_FinalRoundServer) error {
	return status.Errorf(codes.Unimplemented, "method FinalRound not implemented")
}
func (UnimplementedInterviewServer) CampusDrive(Interview_CampusDriveServer) error {
	return status.Errorf(codes.Unimplemented, "method CampusDrive not implemented")
}
func (UnimplementedInterviewServer) mustEmbedUnimplementedInterviewServer() {}

// UnsafeInterviewServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InterviewServer will
// result in compilation errors.
type UnsafeInterviewServer interface {
	mustEmbedUnimplementedInterviewServer()
}

func RegisterInterviewServer(s grpc.ServiceRegistrar, srv InterviewServer) {
	s.RegisterService(&Interview_ServiceDesc, srv)
}

func _Interview_OneOnOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterviewServer).OneOnOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Interview/OneOnOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterviewServer).OneOnOne(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interview_OnlineScreening_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Batch)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InterviewServer).OnlineScreening(m, &interviewOnlineScreeningServer{stream})
}

type Interview_OnlineScreeningServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type interviewOnlineScreeningServer struct {
	grpc.ServerStream
}

func (x *interviewOnlineScreeningServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func _Interview_FinalRound_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InterviewServer).FinalRound(&interviewFinalRoundServer{stream})
}

type Interview_FinalRoundServer interface {
	SendAndClose(*ResultReport) error
	Recv() (*User, error)
	grpc.ServerStream
}

type interviewFinalRoundServer struct {
	grpc.ServerStream
}

func (x *interviewFinalRoundServer) SendAndClose(m *ResultReport) error {
	return x.ServerStream.SendMsg(m)
}

func (x *interviewFinalRoundServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Interview_CampusDrive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InterviewServer).CampusDrive(&interviewCampusDriveServer{stream})
}

type Interview_CampusDriveServer interface {
	Send(*Result) error
	Recv() (*User, error)
	grpc.ServerStream
}

type interviewCampusDriveServer struct {
	grpc.ServerStream
}

func (x *interviewCampusDriveServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func (x *interviewCampusDriveServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Interview_ServiceDesc is the grpc.ServiceDesc for Interview service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Interview_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Interview",
	HandlerType: (*InterviewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OneOnOne",
			Handler:    _Interview_OneOnOne_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OnlineScreening",
			Handler:       _Interview_OnlineScreening_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FinalRound",
			Handler:       _Interview_FinalRound_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CampusDrive",
			Handler:       _Interview_CampusDrive_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/service.proto",
}