// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.1
// source: pki.proto

package portal_apis

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

const (
	SSLService_ListCertificates_FullMethodName   = "/pkiService.SSLService/ListCertificates"
	SSLService_CertificateDetails_FullMethodName = "/pkiService.SSLService/CertificateDetails"
	SSLService_IssueCertificate_FullMethodName   = "/pkiService.SSLService/IssueCertificate"
	SSLService_RevokeCertificate_FullMethodName  = "/pkiService.SSLService/RevokeCertificate"
)

// SSLServiceClient is the client API for SSLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SSLServiceClient interface {
	ListCertificates(ctx context.Context, in *ListSslRequest, opts ...grpc.CallOption) (*ListSslResponse, error)
	CertificateDetails(ctx context.Context, in *CertificateDetailsRequest, opts ...grpc.CallOption) (*SslCertificateDetails, error)
	IssueCertificate(ctx context.Context, in *IssueSslRequest, opts ...grpc.CallOption) (*IssueSslResponse, error)
	RevokeCertificate(ctx context.Context, in *RevokeSslRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type sSLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSSLServiceClient(cc grpc.ClientConnInterface) SSLServiceClient {
	return &sSLServiceClient{cc}
}

func (c *sSLServiceClient) ListCertificates(ctx context.Context, in *ListSslRequest, opts ...grpc.CallOption) (*ListSslResponse, error) {
	out := new(ListSslResponse)
	err := c.cc.Invoke(ctx, SSLService_ListCertificates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLServiceClient) CertificateDetails(ctx context.Context, in *CertificateDetailsRequest, opts ...grpc.CallOption) (*SslCertificateDetails, error) {
	out := new(SslCertificateDetails)
	err := c.cc.Invoke(ctx, SSLService_CertificateDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLServiceClient) IssueCertificate(ctx context.Context, in *IssueSslRequest, opts ...grpc.CallOption) (*IssueSslResponse, error) {
	out := new(IssueSslResponse)
	err := c.cc.Invoke(ctx, SSLService_IssueCertificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLServiceClient) RevokeCertificate(ctx context.Context, in *RevokeSslRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SSLService_RevokeCertificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SSLServiceServer is the server API for SSLService service.
// All implementations must embed UnimplementedSSLServiceServer
// for forward compatibility
type SSLServiceServer interface {
	ListCertificates(context.Context, *ListSslRequest) (*ListSslResponse, error)
	CertificateDetails(context.Context, *CertificateDetailsRequest) (*SslCertificateDetails, error)
	IssueCertificate(context.Context, *IssueSslRequest) (*IssueSslResponse, error)
	RevokeCertificate(context.Context, *RevokeSslRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSSLServiceServer()
}

// UnimplementedSSLServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSSLServiceServer struct {
}

func (UnimplementedSSLServiceServer) ListCertificates(context.Context, *ListSslRequest) (*ListSslResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCertificates not implemented")
}
func (UnimplementedSSLServiceServer) CertificateDetails(context.Context, *CertificateDetailsRequest) (*SslCertificateDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CertificateDetails not implemented")
}
func (UnimplementedSSLServiceServer) IssueCertificate(context.Context, *IssueSslRequest) (*IssueSslResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueCertificate not implemented")
}
func (UnimplementedSSLServiceServer) RevokeCertificate(context.Context, *RevokeSslRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeCertificate not implemented")
}
func (UnimplementedSSLServiceServer) mustEmbedUnimplementedSSLServiceServer() {}

// UnsafeSSLServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SSLServiceServer will
// result in compilation errors.
type UnsafeSSLServiceServer interface {
	mustEmbedUnimplementedSSLServiceServer()
}

func RegisterSSLServiceServer(s grpc.ServiceRegistrar, srv SSLServiceServer) {
	s.RegisterService(&SSLService_ServiceDesc, srv)
}

func _SSLService_ListCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSslRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLServiceServer).ListCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLService_ListCertificates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLServiceServer).ListCertificates(ctx, req.(*ListSslRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLService_CertificateDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLServiceServer).CertificateDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLService_CertificateDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLServiceServer).CertificateDetails(ctx, req.(*CertificateDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLService_IssueCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueSslRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLServiceServer).IssueCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLService_IssueCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLServiceServer).IssueCertificate(ctx, req.(*IssueSslRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLService_RevokeCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeSslRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLServiceServer).RevokeCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLService_RevokeCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLServiceServer).RevokeCertificate(ctx, req.(*RevokeSslRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SSLService_ServiceDesc is the grpc.ServiceDesc for SSLService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SSLService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pkiService.SSLService",
	HandlerType: (*SSLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCertificates",
			Handler:    _SSLService_ListCertificates_Handler,
		},
		{
			MethodName: "CertificateDetails",
			Handler:    _SSLService_CertificateDetails_Handler,
		},
		{
			MethodName: "IssueCertificate",
			Handler:    _SSLService_IssueCertificate_Handler,
		},
		{
			MethodName: "RevokeCertificate",
			Handler:    _SSLService_RevokeCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pki.proto",
}

const (
	SmimeService_ListCertificates_FullMethodName  = "/pkiService.SmimeService/ListCertificates"
	SmimeService_IssueCertificate_FullMethodName  = "/pkiService.SmimeService/IssueCertificate"
	SmimeService_RevokeCertificate_FullMethodName = "/pkiService.SmimeService/RevokeCertificate"
)

// SmimeServiceClient is the client API for SmimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmimeServiceClient interface {
	ListCertificates(ctx context.Context, in *ListSmimeRequest, opts ...grpc.CallOption) (*ListSmimeResponse, error)
	IssueCertificate(ctx context.Context, in *IssueSmimeRequest, opts ...grpc.CallOption) (*IssueSmimeResponse, error)
	RevokeCertificate(ctx context.Context, in *RevokeSmimeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type smimeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmimeServiceClient(cc grpc.ClientConnInterface) SmimeServiceClient {
	return &smimeServiceClient{cc}
}

func (c *smimeServiceClient) ListCertificates(ctx context.Context, in *ListSmimeRequest, opts ...grpc.CallOption) (*ListSmimeResponse, error) {
	out := new(ListSmimeResponse)
	err := c.cc.Invoke(ctx, SmimeService_ListCertificates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smimeServiceClient) IssueCertificate(ctx context.Context, in *IssueSmimeRequest, opts ...grpc.CallOption) (*IssueSmimeResponse, error) {
	out := new(IssueSmimeResponse)
	err := c.cc.Invoke(ctx, SmimeService_IssueCertificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smimeServiceClient) RevokeCertificate(ctx context.Context, in *RevokeSmimeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SmimeService_RevokeCertificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmimeServiceServer is the server API for SmimeService service.
// All implementations must embed UnimplementedSmimeServiceServer
// for forward compatibility
type SmimeServiceServer interface {
	ListCertificates(context.Context, *ListSmimeRequest) (*ListSmimeResponse, error)
	IssueCertificate(context.Context, *IssueSmimeRequest) (*IssueSmimeResponse, error)
	RevokeCertificate(context.Context, *RevokeSmimeRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSmimeServiceServer()
}

// UnimplementedSmimeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSmimeServiceServer struct {
}

func (UnimplementedSmimeServiceServer) ListCertificates(context.Context, *ListSmimeRequest) (*ListSmimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCertificates not implemented")
}
func (UnimplementedSmimeServiceServer) IssueCertificate(context.Context, *IssueSmimeRequest) (*IssueSmimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueCertificate not implemented")
}
func (UnimplementedSmimeServiceServer) RevokeCertificate(context.Context, *RevokeSmimeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeCertificate not implemented")
}
func (UnimplementedSmimeServiceServer) mustEmbedUnimplementedSmimeServiceServer() {}

// UnsafeSmimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmimeServiceServer will
// result in compilation errors.
type UnsafeSmimeServiceServer interface {
	mustEmbedUnimplementedSmimeServiceServer()
}

func RegisterSmimeServiceServer(s grpc.ServiceRegistrar, srv SmimeServiceServer) {
	s.RegisterService(&SmimeService_ServiceDesc, srv)
}

func _SmimeService_ListCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSmimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmimeServiceServer).ListCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmimeService_ListCertificates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmimeServiceServer).ListCertificates(ctx, req.(*ListSmimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmimeService_IssueCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueSmimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmimeServiceServer).IssueCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmimeService_IssueCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmimeServiceServer).IssueCertificate(ctx, req.(*IssueSmimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmimeService_RevokeCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeSmimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmimeServiceServer).RevokeCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmimeService_RevokeCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmimeServiceServer).RevokeCertificate(ctx, req.(*RevokeSmimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmimeService_ServiceDesc is the grpc.ServiceDesc for SmimeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmimeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pkiService.SmimeService",
	HandlerType: (*SmimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCertificates",
			Handler:    _SmimeService_ListCertificates_Handler,
		},
		{
			MethodName: "IssueCertificate",
			Handler:    _SmimeService_IssueCertificate_Handler,
		},
		{
			MethodName: "RevokeCertificate",
			Handler:    _SmimeService_RevokeCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pki.proto",
}
