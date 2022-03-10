// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/api/service/framesystem/v1/frame_system.proto

/*
Package v1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_FrameSystemService_Config_0(ctx context.Context, marshaler runtime.Marshaler, client FrameSystemServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ConfigRequest
	var metadata runtime.ServerMetadata

	msg, err := client.Config(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_FrameSystemService_Config_0(ctx context.Context, marshaler runtime.Marshaler, server FrameSystemServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ConfigRequest
	var metadata runtime.ServerMetadata

	msg, err := server.Config(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterFrameSystemServiceHandlerServer registers the http handlers for service FrameSystemService to "mux".
// UnaryRPC     :call FrameSystemServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterFrameSystemServiceHandlerFromEndpoint instead.
func RegisterFrameSystemServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server FrameSystemServiceServer) error {

	mux.Handle("GET", pattern_FrameSystemService_Config_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/proto.api.service.framesystem.v1.FrameSystemService/Config", runtime.WithHTTPPathPattern("/viam/api/v1/service/frame_system/config"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_FrameSystemService_Config_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FrameSystemService_Config_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterFrameSystemServiceHandlerFromEndpoint is same as RegisterFrameSystemServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterFrameSystemServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterFrameSystemServiceHandler(ctx, mux, conn)
}

// RegisterFrameSystemServiceHandler registers the http handlers for service FrameSystemService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterFrameSystemServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterFrameSystemServiceHandlerClient(ctx, mux, NewFrameSystemServiceClient(conn))
}

// RegisterFrameSystemServiceHandlerClient registers the http handlers for service FrameSystemService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "FrameSystemServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "FrameSystemServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "FrameSystemServiceClient" to call the correct interceptors.
func RegisterFrameSystemServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client FrameSystemServiceClient) error {

	mux.Handle("GET", pattern_FrameSystemService_Config_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/proto.api.service.framesystem.v1.FrameSystemService/Config", runtime.WithHTTPPathPattern("/viam/api/v1/service/frame_system/config"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FrameSystemService_Config_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FrameSystemService_Config_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_FrameSystemService_Config_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 2, 5}, []string{"viam", "api", "v1", "service", "frame_system", "config"}, ""))
)

var (
	forward_FrameSystemService_Config_0 = runtime.ForwardResponseMessage
)
