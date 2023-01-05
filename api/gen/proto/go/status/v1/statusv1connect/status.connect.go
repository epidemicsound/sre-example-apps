// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: status/v1/status.proto

package statusv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/grafana/phlare/api/gen/proto/go/status/v1"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// StatusServiceName is the fully-qualified name of the StatusService service.
	StatusServiceName = "status.v1.StatusService"
)

// StatusServiceClient is a client for the status.v1.StatusService service.
type StatusServiceClient interface {
	// Retrieve build information about the binary
	GetBuildInfo(context.Context, *connect_go.Request[v1.GetBuildInfoRequest]) (*connect_go.Response[v1.GetBuildInfoResponse], error)
	// Retrieve the running config
	GetConfig(context.Context, *connect_go.Request[v1.GetConfigRequest]) (*connect_go.Response[httpbody.HttpBody], error)
	// Retrieve the diff config to the defaults
	GetDiffConfig(context.Context, *connect_go.Request[v1.GetConfigRequest]) (*connect_go.Response[httpbody.HttpBody], error)
	GetDefaultConfig(context.Context, *connect_go.Request[v1.GetConfigRequest]) (*connect_go.Response[httpbody.HttpBody], error)
}

// NewStatusServiceClient constructs a client for the status.v1.StatusService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStatusServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) StatusServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &statusServiceClient{
		getBuildInfo: connect_go.NewClient[v1.GetBuildInfoRequest, v1.GetBuildInfoResponse](
			httpClient,
			baseURL+"/status.v1.StatusService/GetBuildInfo",
			opts...,
		),
		getConfig: connect_go.NewClient[v1.GetConfigRequest, httpbody.HttpBody](
			httpClient,
			baseURL+"/status.v1.StatusService/GetConfig",
			opts...,
		),
		getDiffConfig: connect_go.NewClient[v1.GetConfigRequest, httpbody.HttpBody](
			httpClient,
			baseURL+"/status.v1.StatusService/GetDiffConfig",
			opts...,
		),
		getDefaultConfig: connect_go.NewClient[v1.GetConfigRequest, httpbody.HttpBody](
			httpClient,
			baseURL+"/status.v1.StatusService/GetDefaultConfig",
			opts...,
		),
	}
}

// statusServiceClient implements StatusServiceClient.
type statusServiceClient struct {
	getBuildInfo     *connect_go.Client[v1.GetBuildInfoRequest, v1.GetBuildInfoResponse]
	getConfig        *connect_go.Client[v1.GetConfigRequest, httpbody.HttpBody]
	getDiffConfig    *connect_go.Client[v1.GetConfigRequest, httpbody.HttpBody]
	getDefaultConfig *connect_go.Client[v1.GetConfigRequest, httpbody.HttpBody]
}

// GetBuildInfo calls status.v1.StatusService.GetBuildInfo.
func (c *statusServiceClient) GetBuildInfo(ctx context.Context, req *connect_go.Request[v1.GetBuildInfoRequest]) (*connect_go.Response[v1.GetBuildInfoResponse], error) {
	return c.getBuildInfo.CallUnary(ctx, req)
}

// GetConfig calls status.v1.StatusService.GetConfig.
func (c *statusServiceClient) GetConfig(ctx context.Context, req *connect_go.Request[v1.GetConfigRequest]) (*connect_go.Response[httpbody.HttpBody], error) {
	return c.getConfig.CallUnary(ctx, req)
}

// GetDiffConfig calls status.v1.StatusService.GetDiffConfig.
func (c *statusServiceClient) GetDiffConfig(ctx context.Context, req *connect_go.Request[v1.GetConfigRequest]) (*connect_go.Response[httpbody.HttpBody], error) {
	return c.getDiffConfig.CallUnary(ctx, req)
}

// GetDefaultConfig calls status.v1.StatusService.GetDefaultConfig.
func (c *statusServiceClient) GetDefaultConfig(ctx context.Context, req *connect_go.Request[v1.GetConfigRequest]) (*connect_go.Response[httpbody.HttpBody], error) {
	return c.getDefaultConfig.CallUnary(ctx, req)
}

// StatusServiceHandler is an implementation of the status.v1.StatusService service.
type StatusServiceHandler interface {
	// Retrieve build information about the binary
	GetBuildInfo(context.Context, *connect_go.Request[v1.GetBuildInfoRequest]) (*connect_go.Response[v1.GetBuildInfoResponse], error)
	// Retrieve the running config
	GetConfig(context.Context, *connect_go.Request[v1.GetConfigRequest]) (*connect_go.Response[httpbody.HttpBody], error)
	// Retrieve the diff config to the defaults
	GetDiffConfig(context.Context, *connect_go.Request[v1.GetConfigRequest]) (*connect_go.Response[httpbody.HttpBody], error)
	GetDefaultConfig(context.Context, *connect_go.Request[v1.GetConfigRequest]) (*connect_go.Response[httpbody.HttpBody], error)
}

// NewStatusServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStatusServiceHandler(svc StatusServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/status.v1.StatusService/GetBuildInfo", connect_go.NewUnaryHandler(
		"/status.v1.StatusService/GetBuildInfo",
		svc.GetBuildInfo,
		opts...,
	))
	mux.Handle("/status.v1.StatusService/GetConfig", connect_go.NewUnaryHandler(
		"/status.v1.StatusService/GetConfig",
		svc.GetConfig,
		opts...,
	))
	mux.Handle("/status.v1.StatusService/GetDiffConfig", connect_go.NewUnaryHandler(
		"/status.v1.StatusService/GetDiffConfig",
		svc.GetDiffConfig,
		opts...,
	))
	mux.Handle("/status.v1.StatusService/GetDefaultConfig", connect_go.NewUnaryHandler(
		"/status.v1.StatusService/GetDefaultConfig",
		svc.GetDefaultConfig,
		opts...,
	))
	return "/status.v1.StatusService/", mux
}

// UnimplementedStatusServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedStatusServiceHandler struct{}

func (UnimplementedStatusServiceHandler) GetBuildInfo(context.Context, *connect_go.Request[v1.GetBuildInfoRequest]) (*connect_go.Response[v1.GetBuildInfoResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("status.v1.StatusService.GetBuildInfo is not implemented"))
}

func (UnimplementedStatusServiceHandler) GetConfig(context.Context, *connect_go.Request[v1.GetConfigRequest]) (*connect_go.Response[httpbody.HttpBody], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("status.v1.StatusService.GetConfig is not implemented"))
}

func (UnimplementedStatusServiceHandler) GetDiffConfig(context.Context, *connect_go.Request[v1.GetConfigRequest]) (*connect_go.Response[httpbody.HttpBody], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("status.v1.StatusService.GetDiffConfig is not implemented"))
}

func (UnimplementedStatusServiceHandler) GetDefaultConfig(context.Context, *connect_go.Request[v1.GetConfigRequest]) (*connect_go.Response[httpbody.HttpBody], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("status.v1.StatusService.GetDefaultConfig is not implemented"))
}
