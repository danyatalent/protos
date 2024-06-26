// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: question-service/question-service.proto

package questionv1

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

// QuestionsClient is the client API for Questions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionsClient interface {
	GetQuestion(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*QuestionResponse, error)
}

type questionsClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionsClient(cc grpc.ClientConnInterface) QuestionsClient {
	return &questionsClient{cc}
}

func (c *questionsClient) GetQuestion(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*QuestionResponse, error) {
	out := new(QuestionResponse)
	err := c.cc.Invoke(ctx, "/question.Questions/GetQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionsServer is the server API for Questions service.
// All implementations must embed UnimplementedQuestionsServer
// for forward compatibility
type QuestionsServer interface {
	GetQuestion(context.Context, *Empty) (*QuestionResponse, error)
	mustEmbedUnimplementedQuestionsServer()
}

// UnimplementedQuestionsServer must be embedded to have forward compatible implementations.
type UnimplementedQuestionsServer struct {
}

func (UnimplementedQuestionsServer) GetQuestion(context.Context, *Empty) (*QuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestion not implemented")
}
func (UnimplementedQuestionsServer) mustEmbedUnimplementedQuestionsServer() {}

// UnsafeQuestionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionsServer will
// result in compilation errors.
type UnsafeQuestionsServer interface {
	mustEmbedUnimplementedQuestionsServer()
}

func RegisterQuestionsServer(s grpc.ServiceRegistrar, srv QuestionsServer) {
	s.RegisterService(&Questions_ServiceDesc, srv)
}

func _Questions_GetQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionsServer).GetQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/question.Questions/GetQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionsServer).GetQuestion(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Questions_ServiceDesc is the grpc.ServiceDesc for Questions service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Questions_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "question.Questions",
	HandlerType: (*QuestionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuestion",
			Handler:    _Questions_GetQuestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "question-service/question-service.proto",
}
