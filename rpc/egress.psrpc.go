// Code generated by protoc-gen-psrpc v0.3.0, DO NOT EDIT.
// source: rpc/egress.proto

package rpc

import (
	"context"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/client"
	"github.com/livekit/psrpc/pkg/info"
	"github.com/livekit/psrpc/pkg/server"
	"github.com/livekit/psrpc/version"
)
import livekit1 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_3_0

// ===============================
// EgressInternal Client Interface
// ===============================

type EgressInternalClient interface {
	SubscribeStartEgress(ctx context.Context) (psrpc.Subscription[*livekit1.EgressInfo], error)

	ListActiveEgress(ctx context.Context, req *ListActiveEgressRequest, opts ...psrpc.RequestOption) (*ListActiveEgressResponse, error)
}

// ===================================
// EgressInternal ServerImpl Interface
// ===================================

type EgressInternalServerImpl interface {
	ListActiveEgress(context.Context, *ListActiveEgressRequest) (*ListActiveEgressResponse, error)
	ListActiveEgressAffinity(*ListActiveEgressRequest) float32
}

// ===============================
// EgressInternal Server Interface
// ===============================

type EgressInternalServer interface {
	PublishStartEgress(ctx context.Context, msg *livekit1.EgressInfo) error

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =====================
// EgressInternal Client
// =====================

type egressInternalClient struct {
	client *client.RPCClient
}

// NewEgressInternalClient creates a psrpc client that implements the EgressInternalClient interface.
func NewEgressInternalClient(clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (EgressInternalClient, error) {
	sd := &info.ServiceDefinition{
		Name: "EgressInternal",
		ID:   clientID,
	}

	sd.RegisterMethod("StartEgress", false, false, true)
	sd.RegisterMethod("ListActiveEgress", true, false, true)

	rpcClient, err := client.NewRPCClientWithStreams(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &egressInternalClient{
		client: rpcClient,
	}, nil
}

func (c *egressInternalClient) SubscribeStartEgress(ctx context.Context) (psrpc.Subscription[*livekit1.EgressInfo], error) {
	return client.JoinQueue[*livekit1.EgressInfo](ctx, c.client, "StartEgress", nil)
}

func (c *egressInternalClient) ListActiveEgress(ctx context.Context, req *ListActiveEgressRequest, opts ...psrpc.RequestOption) (*ListActiveEgressResponse, error) {
	return client.RequestSingle[*ListActiveEgressResponse](ctx, c.client, "ListActiveEgress", nil, req, opts...)
}

// =====================
// EgressInternal Server
// =====================

type egressInternalServer struct {
	svc EgressInternalServerImpl
	rpc *server.RPCServer
}

// NewEgressInternalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewEgressInternalServer(serverID string, svc EgressInternalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (EgressInternalServer, error) {
	sd := &info.ServiceDefinition{
		Name: "EgressInternal",
		ID:   serverID,
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("StartEgress", false, false, true)
	sd.RegisterMethod("ListActiveEgress", true, false, true)
	var err error
	err = server.RegisterHandler(s, "ListActiveEgress", nil, svc.ListActiveEgress, svc.ListActiveEgressAffinity)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	return &egressInternalServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *egressInternalServer) PublishStartEgress(ctx context.Context, msg *livekit1.EgressInfo) error {
	return s.rpc.Publish(ctx, "StartEgress", nil, msg)
}

func (s *egressInternalServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *egressInternalServer) Kill() {
	s.rpc.Close(true)
}

// ==============================
// EgressHandler Client Interface
// ==============================

type EgressHandlerClient interface {
	SubscribeUpdateStream(ctx context.Context) (psrpc.Subscription[*livekit1.EgressInfo], error)

	SubscribeStopEgress(ctx context.Context) (psrpc.Subscription[*livekit1.EgressInfo], error)
}

// ==================================
// EgressHandler ServerImpl Interface
// ==================================

type EgressHandlerServerImpl interface {
}

// ==============================
// EgressHandler Server Interface
// ==============================

type EgressHandlerServer interface {
	PublishUpdateStream(ctx context.Context, msg *livekit1.EgressInfo) error

	PublishStopEgress(ctx context.Context, msg *livekit1.EgressInfo) error

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ====================
// EgressHandler Client
// ====================

type egressHandlerClient struct {
	client *client.RPCClient
}

// NewEgressHandlerClient creates a psrpc client that implements the EgressHandlerClient interface.
func NewEgressHandlerClient(clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (EgressHandlerClient, error) {
	sd := &info.ServiceDefinition{
		Name: "EgressHandler",
		ID:   clientID,
	}

	sd.RegisterMethod("UpdateStream", false, false, true)
	sd.RegisterMethod("StopEgress", false, false, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &egressHandlerClient{
		client: rpcClient,
	}, nil
}

func (c *egressHandlerClient) SubscribeUpdateStream(ctx context.Context) (psrpc.Subscription[*livekit1.EgressInfo], error) {
	return client.JoinQueue[*livekit1.EgressInfo](ctx, c.client, "UpdateStream", nil)
}

func (c *egressHandlerClient) SubscribeStopEgress(ctx context.Context) (psrpc.Subscription[*livekit1.EgressInfo], error) {
	return client.JoinQueue[*livekit1.EgressInfo](ctx, c.client, "StopEgress", nil)
}

// ====================
// EgressHandler Server
// ====================

type egressHandlerServer struct {
	svc EgressHandlerServerImpl
	rpc *server.RPCServer
}

// NewEgressHandlerServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewEgressHandlerServer(serverID string, svc EgressHandlerServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (EgressHandlerServer, error) {
	sd := &info.ServiceDefinition{
		Name: "EgressHandler",
		ID:   serverID,
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("UpdateStream", false, false, true)
	sd.RegisterMethod("StopEgress", false, false, true)
	return &egressHandlerServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *egressHandlerServer) PublishUpdateStream(ctx context.Context, msg *livekit1.EgressInfo) error {
	return s.rpc.Publish(ctx, "UpdateStream", nil, msg)
}

func (s *egressHandlerServer) PublishStopEgress(ctx context.Context, msg *livekit1.EgressInfo) error {
	return s.rpc.Publish(ctx, "StopEgress", nil, msg)
}

func (s *egressHandlerServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *egressHandlerServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor0 = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x6d, 0x6f, 0xd2, 0x40,
	0x1c, 0xf7, 0xd6, 0xc1, 0xe8, 0x9f, 0x81, 0xe4, 0xc4, 0x50, 0xba, 0x2d, 0x21, 0xa8, 0x09, 0x59,
	0xb4, 0x35, 0xec, 0x8d, 0xfa, 0x8e, 0x19, 0x92, 0x91, 0xcc, 0x18, 0x8b, 0xcb, 0x12, 0xdf, 0x90,
	0xd2, 0xde, 0xb0, 0xa1, 0xed, 0x9d, 0x77, 0x07, 0x84, 0x8f, 0xb0, 0xcf, 0xe0, 0x2b, 0xbf, 0x81,
	0xd9, 0x27, 0x34, 0xdc, 0x95, 0x0a, 0x2c, 0x18, 0x5f, 0x41, 0x7f, 0x4f, 0xf7, 0x7f, 0x82, 0x1a,
	0x67, 0x81, 0x4b, 0x26, 0x9c, 0x08, 0xe1, 0x30, 0x4e, 0x25, 0xc5, 0x06, 0x67, 0x81, 0xdd, 0x9c,
	0x50, 0x3a, 0x89, 0x89, 0xab, 0xa0, 0xf1, 0xec, 0xce, 0xf5, 0xd3, 0xa5, 0xe6, 0xed, 0x0a, 0x65,
	0x32, 0xa2, 0x69, 0x26, 0xb7, 0xeb, 0x71, 0x34, 0x27, 0xd3, 0x48, 0x8e, 0x36, 0x43, 0xda, 0x3f,
	0x0f, 0x01, 0x0f, 0xa5, 0xcf, 0x65, 0x5f, 0xa1, 0x1e, 0xf9, 0x31, 0x23, 0x42, 0xe2, 0x13, 0x30,
	0xb5, 0x6c, 0x14, 0x85, 0x16, 0x6a, 0xa1, 0x8e, 0xe9, 0x95, 0x34, 0x30, 0x08, 0xf1, 0x35, 0x54,
	0x39, 0xa5, 0xc9, 0x28, 0xa0, 0x09, 0xa3, 0x22, 0x92, 0xc4, 0x2a, 0xb4, 0x50, 0xa7, 0xdc, 0x7d,
	0xe1, 0x64, 0x4f, 0x38, 0x1e, 0xa5, 0xc9, 0xc7, 0x35, 0xbb, 0x95, 0x7c, 0xf5, 0xc4, 0xab, 0xf0,
	0x4d, 0x16, 0x7f, 0x86, 0xa7, 0x92, 0xfb, 0xc1, 0x74, 0x23, 0xae, 0xa8, 0xe2, 0x5e, 0xe6, 0x71,
	0x5f, 0x57, 0xfc, 0xde, 0xbc, 0xaa, 0xdc, 0xa2, 0xf1, 0x05, 0x14, 0x14, 0x62, 0x1d, 0xa9, 0x98,
	0x93, 0xed, 0x98, 0x5d, 0xb7, 0xd6, 0xe2, 0x37, 0x60, 0x2c, 0xc8, 0xd8, 0x2a, 0x2b, 0x4b, 0x33,
	0xb7, 0xdc, 0x92, 0xf1, 0xae, 0x61, 0xa5, 0xc3, 0x0d, 0x38, 0x52, 0x23, 0x88, 0x42, 0xcb, 0x50,
	0xd3, 0x29, 0xae, 0x3e, 0x07, 0x21, 0xae, 0x43, 0x41, 0xd2, 0x29, 0x49, 0xad, 0x92, 0x82, 0xf5,
	0x07, 0x7e, 0x0e, 0xc5, 0x85, 0x18, 0xcd, 0x78, 0x6c, 0x99, 0x1a, 0x5e, 0x88, 0x1b, 0x1e, 0xe3,
	0x1e, 0x94, 0x12, 0x22, 0xfd, 0xd0, 0x97, 0xbe, 0x75, 0xdc, 0x32, 0x3a, 0xe5, 0xee, 0x2b, 0x87,
	0xb3, 0xc0, 0x79, 0xbc, 0x10, 0xe7, 0x53, 0xa6, 0xeb, 0xa7, 0x92, 0x2f, 0xbd, 0xdc, 0x66, 0x7f,
	0x81, 0xca, 0x16, 0x85, 0x6b, 0x60, 0x4c, 0xc9, 0x32, 0xdb, 0xd9, 0xea, 0x2f, 0x3e, 0x87, 0xc2,
	0xdc, 0x8f, 0x67, 0xc4, 0x3a, 0x50, 0xcd, 0xd5, 0x1d, 0x7d, 0x32, 0xce, 0xfa, 0x64, 0x9c, 0x5e,
	0xba, 0xf4, 0xb4, 0xe4, 0xc3, 0xc1, 0x3b, 0x74, 0x69, 0xc2, 0x11, 0xd7, 0xaf, 0xb6, 0x9b, 0xd0,
	0xb8, 0x8e, 0x84, 0xec, 0x05, 0x32, 0x9a, 0x6f, 0xcf, 0xbd, 0xfd, 0x1e, 0xac, 0xc7, 0x94, 0x60,
	0x34, 0x15, 0x04, 0x9f, 0x01, 0xe4, 0xd7, 0x23, 0x2c, 0xd4, 0x32, 0x3a, 0xa6, 0x67, 0xae, 0xcf,
	0x47, 0x74, 0x7f, 0x23, 0xa8, 0x6a, 0xc7, 0x20, 0x95, 0x84, 0xa7, 0x7e, 0x8c, 0xfb, 0x50, 0xde,
	0x68, 0x1a, 0x37, 0xf6, 0x8c, 0xc1, 0x7e, 0x96, 0x6f, 0x66, 0x1d, 0x70, 0x47, 0xdb, 0xa5, 0x87,
	0x7b, 0x74, 0x58, 0x43, 0x6f, 0x11, 0xbe, 0x85, 0xda, 0x6e, 0x51, 0xf8, 0x54, 0x65, 0xed, 0x69,
	0xc3, 0x3e, 0xdb, 0xc3, 0xea, 0x4e, 0xda, 0xc5, 0x87, 0x7b, 0x74, 0xd0, 0x41, 0xdd, 0x5f, 0x08,
	0x2a, 0x9a, 0xba, 0xf2, 0xd3, 0x30, 0x26, 0x1c, 0x0f, 0xe0, 0xf8, 0x86, 0x85, 0xbe, 0x24, 0x43,
	0xc9, 0x89, 0x9f, 0xe0, 0xd3, 0xbc, 0xb2, 0x4d, 0xf8, 0x9f, 0x75, 0xab, 0xf0, 0x1a, 0xc2, 0x7d,
	0x80, 0xa1, 0xa4, 0x2c, 0xab, 0xd7, 0xce, 0xa5, 0x7f, 0xc1, 0xff, 0x89, 0xb9, 0x7c, 0xfd, 0xed,
	0x7c, 0x12, 0xc9, 0xef, 0xb3, 0xb1, 0x13, 0xd0, 0xc4, 0xcd, 0x84, 0xf9, 0x2f, 0x9b, 0x4e, 0x5c,
	0x41, 0xf8, 0x3c, 0x0a, 0x88, 0xcb, 0x59, 0x30, 0x2e, 0xaa, 0xf5, 0x5f, 0xfc, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x01, 0xc3, 0x6c, 0x5c, 0x58, 0x04, 0x00, 0x00,
}
