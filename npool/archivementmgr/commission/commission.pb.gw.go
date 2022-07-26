// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: npool/archivementmgr/commission/commission.proto

/*
Package commission is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package commission

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

func request_Commission_CreateAppUserGoodCommissions_0(ctx context.Context, marshaler runtime.Marshaler, client CommissionClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateAppUserGoodCommissionsRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateAppUserGoodCommissions(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Commission_CreateAppUserGoodCommissions_0(ctx context.Context, marshaler runtime.Marshaler, server CommissionServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateAppUserGoodCommissionsRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CreateAppUserGoodCommissions(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterCommissionHandlerServer registers the http handlers for service Commission to "mux".
// UnaryRPC     :call CommissionServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterCommissionHandlerFromEndpoint instead.
func RegisterCommissionHandlerServer(ctx context.Context, mux *runtime.ServeMux, server CommissionServer) error {

	mux.Handle("POST", pattern_Commission_CreateAppUserGoodCommissions_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/archivement.manager.commission.v1.Commission/CreateAppUserGoodCommissions", runtime.WithHTTPPathPattern("/v1/create/app/user/good/commissions"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Commission_CreateAppUserGoodCommissions_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Commission_CreateAppUserGoodCommissions_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterCommissionHandlerFromEndpoint is same as RegisterCommissionHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCommissionHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterCommissionHandler(ctx, mux, conn)
}

// RegisterCommissionHandler registers the http handlers for service Commission to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCommissionHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCommissionHandlerClient(ctx, mux, NewCommissionClient(conn))
}

// RegisterCommissionHandlerClient registers the http handlers for service Commission
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CommissionClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CommissionClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CommissionClient" to call the correct interceptors.
func RegisterCommissionHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CommissionClient) error {

	mux.Handle("POST", pattern_Commission_CreateAppUserGoodCommissions_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/archivement.manager.commission.v1.Commission/CreateAppUserGoodCommissions", runtime.WithHTTPPathPattern("/v1/create/app/user/good/commissions"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Commission_CreateAppUserGoodCommissions_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Commission_CreateAppUserGoodCommissions_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Commission_CreateAppUserGoodCommissions_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 2, 5}, []string{"v1", "create", "app", "user", "good", "commissions"}, ""))
)

var (
	forward_Commission_CreateAppUserGoodCommissions_0 = runtime.ForwardResponseMessage
)
