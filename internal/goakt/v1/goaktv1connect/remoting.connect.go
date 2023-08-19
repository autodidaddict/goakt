// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: goakt/v1/remoting.proto

package goaktv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/tochemey/goakt/internal/goakt/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// RemoteMessagingServiceName is the fully-qualified name of the RemoteMessagingService service.
	RemoteMessagingServiceName = "goakt.v1.RemoteMessagingService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RemoteMessagingServiceRemoteAskProcedure is the fully-qualified name of the
	// RemoteMessagingService's RemoteAsk RPC.
	RemoteMessagingServiceRemoteAskProcedure = "/goakt.v1.RemoteMessagingService/RemoteAsk"
	// RemoteMessagingServiceRemoteTellProcedure is the fully-qualified name of the
	// RemoteMessagingService's RemoteTell RPC.
	RemoteMessagingServiceRemoteTellProcedure = "/goakt.v1.RemoteMessagingService/RemoteTell"
	// RemoteMessagingServiceRemoteLookupProcedure is the fully-qualified name of the
	// RemoteMessagingService's RemoteLookup RPC.
	RemoteMessagingServiceRemoteLookupProcedure = "/goakt.v1.RemoteMessagingService/RemoteLookup"
)

// RemoteMessagingServiceClient is a client for the goakt.v1.RemoteMessagingService service.
type RemoteMessagingServiceClient interface {
	// RemoteAsk is used to send a message to an actor remotely and expect a response
	// immediately. With this type of message the receiver cannot communicate back to Sender
	// except reply the message with a response. This one-way communication
	RemoteAsk(context.Context, *connect.Request[v1.RemoteAskRequest]) (*connect.Response[v1.RemoteAskResponse], error)
	// RemoteTell is used to send a message to an actor remotely by another actor
	// This is the only way remote actors can interact with each other. The actor on the
	// other line can reply to the sender by using the Sender in the message
	RemoteTell(context.Context, *connect.Request[v1.RemoteTellRequest]) (*connect.Response[v1.RemoteTellResponse], error)
	// Lookup for an actor on a remote host.
	RemoteLookup(context.Context, *connect.Request[v1.RemoteLookupRequest]) (*connect.Response[v1.RemoteLookupResponse], error)
}

// NewRemoteMessagingServiceClient constructs a client for the goakt.v1.RemoteMessagingService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRemoteMessagingServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) RemoteMessagingServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &remoteMessagingServiceClient{
		remoteAsk: connect.NewClient[v1.RemoteAskRequest, v1.RemoteAskResponse](
			httpClient,
			baseURL+RemoteMessagingServiceRemoteAskProcedure,
			opts...,
		),
		remoteTell: connect.NewClient[v1.RemoteTellRequest, v1.RemoteTellResponse](
			httpClient,
			baseURL+RemoteMessagingServiceRemoteTellProcedure,
			opts...,
		),
		remoteLookup: connect.NewClient[v1.RemoteLookupRequest, v1.RemoteLookupResponse](
			httpClient,
			baseURL+RemoteMessagingServiceRemoteLookupProcedure,
			opts...,
		),
	}
}

// remoteMessagingServiceClient implements RemoteMessagingServiceClient.
type remoteMessagingServiceClient struct {
	remoteAsk    *connect.Client[v1.RemoteAskRequest, v1.RemoteAskResponse]
	remoteTell   *connect.Client[v1.RemoteTellRequest, v1.RemoteTellResponse]
	remoteLookup *connect.Client[v1.RemoteLookupRequest, v1.RemoteLookupResponse]
}

// RemoteAsk calls goakt.v1.RemoteMessagingService.RemoteAsk.
func (c *remoteMessagingServiceClient) RemoteAsk(ctx context.Context, req *connect.Request[v1.RemoteAskRequest]) (*connect.Response[v1.RemoteAskResponse], error) {
	return c.remoteAsk.CallUnary(ctx, req)
}

// RemoteTell calls goakt.v1.RemoteMessagingService.RemoteTell.
func (c *remoteMessagingServiceClient) RemoteTell(ctx context.Context, req *connect.Request[v1.RemoteTellRequest]) (*connect.Response[v1.RemoteTellResponse], error) {
	return c.remoteTell.CallUnary(ctx, req)
}

// RemoteLookup calls goakt.v1.RemoteMessagingService.RemoteLookup.
func (c *remoteMessagingServiceClient) RemoteLookup(ctx context.Context, req *connect.Request[v1.RemoteLookupRequest]) (*connect.Response[v1.RemoteLookupResponse], error) {
	return c.remoteLookup.CallUnary(ctx, req)
}

// RemoteMessagingServiceHandler is an implementation of the goakt.v1.RemoteMessagingService
// service.
type RemoteMessagingServiceHandler interface {
	// RemoteAsk is used to send a message to an actor remotely and expect a response
	// immediately. With this type of message the receiver cannot communicate back to Sender
	// except reply the message with a response. This one-way communication
	RemoteAsk(context.Context, *connect.Request[v1.RemoteAskRequest]) (*connect.Response[v1.RemoteAskResponse], error)
	// RemoteTell is used to send a message to an actor remotely by another actor
	// This is the only way remote actors can interact with each other. The actor on the
	// other line can reply to the sender by using the Sender in the message
	RemoteTell(context.Context, *connect.Request[v1.RemoteTellRequest]) (*connect.Response[v1.RemoteTellResponse], error)
	// Lookup for an actor on a remote host.
	RemoteLookup(context.Context, *connect.Request[v1.RemoteLookupRequest]) (*connect.Response[v1.RemoteLookupResponse], error)
}

// NewRemoteMessagingServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRemoteMessagingServiceHandler(svc RemoteMessagingServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	remoteMessagingServiceRemoteAskHandler := connect.NewUnaryHandler(
		RemoteMessagingServiceRemoteAskProcedure,
		svc.RemoteAsk,
		opts...,
	)
	remoteMessagingServiceRemoteTellHandler := connect.NewUnaryHandler(
		RemoteMessagingServiceRemoteTellProcedure,
		svc.RemoteTell,
		opts...,
	)
	remoteMessagingServiceRemoteLookupHandler := connect.NewUnaryHandler(
		RemoteMessagingServiceRemoteLookupProcedure,
		svc.RemoteLookup,
		opts...,
	)
	return "/goakt.v1.RemoteMessagingService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RemoteMessagingServiceRemoteAskProcedure:
			remoteMessagingServiceRemoteAskHandler.ServeHTTP(w, r)
		case RemoteMessagingServiceRemoteTellProcedure:
			remoteMessagingServiceRemoteTellHandler.ServeHTTP(w, r)
		case RemoteMessagingServiceRemoteLookupProcedure:
			remoteMessagingServiceRemoteLookupHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRemoteMessagingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRemoteMessagingServiceHandler struct{}

func (UnimplementedRemoteMessagingServiceHandler) RemoteAsk(context.Context, *connect.Request[v1.RemoteAskRequest]) (*connect.Response[v1.RemoteAskResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("goakt.v1.RemoteMessagingService.RemoteAsk is not implemented"))
}

func (UnimplementedRemoteMessagingServiceHandler) RemoteTell(context.Context, *connect.Request[v1.RemoteTellRequest]) (*connect.Response[v1.RemoteTellResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("goakt.v1.RemoteMessagingService.RemoteTell is not implemented"))
}

func (UnimplementedRemoteMessagingServiceHandler) RemoteLookup(context.Context, *connect.Request[v1.RemoteLookupRequest]) (*connect.Response[v1.RemoteLookupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("goakt.v1.RemoteMessagingService.RemoteLookup is not implemented"))
}
