// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: npool/inspire/gw/v1/coupon/app/scope/scope.proto

/*
Package scope is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package scope

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

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func request_Gateway_GetAppGoodScopes_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAppGoodScopesRequest
<<<<<<< HEAD
=======
func request_Gateway_CreateAppScope_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateAppScopeRequest
>>>>>>> 789287843 (update scope)
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

<<<<<<< HEAD
	msg, err := client.GetAppGoodScopes(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
=======
	msg, err := client.CreateAppScope(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
>>>>>>> 789287843 (update scope)
	return msg, metadata, err

}

<<<<<<< HEAD
func local_request_Gateway_GetAppGoodScopes_0(ctx context.Context, marshaler runtime.Marshaler, server GatewayServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAppGoodScopesRequest
=======
func local_request_Gateway_CreateAppScope_0(ctx context.Context, marshaler runtime.Marshaler, server GatewayServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateAppScopeRequest
>>>>>>> 789287843 (update scope)
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

<<<<<<< HEAD
	msg, err := server.GetAppGoodScopes(ctx, &protoReq)
=======
	msg, err := server.CreateAppScope(ctx, &protoReq)
	return msg, metadata, err

}

func request_Gateway_GetAppScopes_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
=======
func request_Gateway_GetAppGoodScopes_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
>>>>>>> ce3f1742c (delete createappscope)
	var protoReq GetAppScopesRequest
=======
>>>>>>> 808e18b40 (update req)
=======
=======
>>>>>>> e6bdbe852 (delete createappscope)
func request_Gateway_GetAppGoodScopes_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAppGoodScopesRequest
<<<<<<< HEAD
=======
func request_Gateway_CreateAppScope_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateAppScopeRequest
>>>>>>> 569d52611 (update scope)
>>>>>>> 2e441a23e (update scope)
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

<<<<<<< HEAD
	msg, err := client.GetAppGoodScopes(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
=======
	msg, err := client.CreateAppScope(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
>>>>>>> 569d52611 (update scope)
	return msg, metadata, err

}

<<<<<<< HEAD
func local_request_Gateway_GetAppGoodScopes_0(ctx context.Context, marshaler runtime.Marshaler, server GatewayServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAppGoodScopesRequest
=======
func local_request_Gateway_CreateAppScope_0(ctx context.Context, marshaler runtime.Marshaler, server GatewayServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateAppScopeRequest
>>>>>>> 569d52611 (update scope)
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

<<<<<<< HEAD
<<<<<<< HEAD
	msg, err := server.GetAppScopes(ctx, &protoReq)
>>>>>>> 789287843 (update scope)
=======
	msg, err := server.GetAppGoodScopes(ctx, &protoReq)
>>>>>>> ce3f1742c (delete createappscope)
=======
	msg, err := server.GetAppGoodScopes(ctx, &protoReq)
=======
	msg, err := server.CreateAppScope(ctx, &protoReq)
	return msg, metadata, err

}

func request_Gateway_GetAppScopes_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
=======
func request_Gateway_GetAppGoodScopes_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
>>>>>>> 80bc5ee31 (delete createappscope)
	var protoReq GetAppScopesRequest
=======
>>>>>>> f7e76ad3e (update req)
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetAppGoodScopes(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Gateway_GetAppGoodScopes_0(ctx context.Context, marshaler runtime.Marshaler, server GatewayServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAppGoodScopesRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

<<<<<<< HEAD
	msg, err := server.GetAppScopes(ctx, &protoReq)
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
	msg, err := server.GetAppGoodScopes(ctx, &protoReq)
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)
	return msg, metadata, err

}

func request_Gateway_CreateAppGoodScope_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateAppGoodScopeRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateAppGoodScope(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Gateway_CreateAppGoodScope_0(ctx context.Context, marshaler runtime.Marshaler, server GatewayServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateAppGoodScopeRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CreateAppGoodScope(ctx, &protoReq)
	return msg, metadata, err

}

func request_Gateway_DeleteAppGoodScope_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DeleteAppGoodScopeRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.DeleteAppGoodScope(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Gateway_DeleteAppGoodScope_0(ctx context.Context, marshaler runtime.Marshaler, server GatewayServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DeleteAppGoodScopeRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.DeleteAppGoodScope(ctx, &protoReq)
	return msg, metadata, err

}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> 6712e122f (update scope)
=======
func request_Gateway_GetAppGoodScopes_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAppGoodScopesRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetAppGoodScopes(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Gateway_GetAppGoodScopes_0(ctx context.Context, marshaler runtime.Marshaler, server GatewayServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAppGoodScopesRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetAppGoodScopes(ctx, &protoReq)
	return msg, metadata, err

}

<<<<<<< HEAD
>>>>>>> 789287843 (update scope)
=======
>>>>>>> 60afcdcd5 (update scope)
=======
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
>>>>>>> 909af8f64 (update scope)
>>>>>>> 6712e122f (update scope)
// RegisterGatewayHandlerServer registers the http handlers for service Gateway to "mux".
// UnaryRPC     :call GatewayServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterGatewayHandlerFromEndpoint instead.
func RegisterGatewayHandlerServer(ctx context.Context, mux *runtime.ServeMux, server GatewayServer) error {

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	mux.Handle("POST", pattern_Gateway_GetAppGoodScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
=======
	mux.Handle("POST", pattern_Gateway_CreateAppScope_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
>>>>>>> 789287843 (update scope)
=======
	mux.Handle("POST", pattern_Gateway_GetAppGoodScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
>>>>>>> ce3f1742c (delete createappscope)
=======
=======
>>>>>>> e6bdbe852 (delete createappscope)
	mux.Handle("POST", pattern_Gateway_GetAppGoodScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
=======
	mux.Handle("POST", pattern_Gateway_CreateAppScope_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
	mux.Handle("POST", pattern_Gateway_GetAppGoodScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes", runtime.WithHTTPPathPattern("/v1/get/appgoodscopes"))
=======
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/CreateAppScope", runtime.WithHTTPPathPattern("/v1/create/appscope"))
>>>>>>> 789287843 (update scope)
=======
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes", runtime.WithHTTPPathPattern("/v1/get/appgoodscopes"))
>>>>>>> ce3f1742c (delete createappscope)
=======
=======
>>>>>>> e6bdbe852 (delete createappscope)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes", runtime.WithHTTPPathPattern("/v1/get/appgoodscopes"))
=======
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/CreateAppScope", runtime.WithHTTPPathPattern("/v1/create/appscope"))
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes", runtime.WithHTTPPathPattern("/v1/get/appgoodscopes"))
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
		resp, md, err := local_request_Gateway_GetAppGoodScopes_0(rctx, inboundMarshaler, server, req, pathParams)
=======
		resp, md, err := local_request_Gateway_CreateAppScope_0(rctx, inboundMarshaler, server, req, pathParams)
>>>>>>> 789287843 (update scope)
=======
		resp, md, err := local_request_Gateway_GetAppGoodScopes_0(rctx, inboundMarshaler, server, req, pathParams)
>>>>>>> ce3f1742c (delete createappscope)
=======
=======
>>>>>>> e6bdbe852 (delete createappscope)
		resp, md, err := local_request_Gateway_GetAppGoodScopes_0(rctx, inboundMarshaler, server, req, pathParams)
=======
		resp, md, err := local_request_Gateway_CreateAppScope_0(rctx, inboundMarshaler, server, req, pathParams)
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
		resp, md, err := local_request_Gateway_GetAppGoodScopes_0(rctx, inboundMarshaler, server, req, pathParams)
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> e6bdbe852 (delete createappscope)
		forward_Gateway_GetAppGoodScopes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
=======
		forward_Gateway_CreateAppScope_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Gateway_GetAppScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppScopes", runtime.WithHTTPPathPattern("/v1/get/appscopes"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Gateway_GetAppScopes_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_GetAppScopes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
<<<<<<< HEAD
>>>>>>> 789287843 (update scope)
=======
		forward_Gateway_GetAppGoodScopes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
>>>>>>> ce3f1742c (delete createappscope)
=======
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
		forward_Gateway_GetAppGoodScopes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)

	})

	mux.Handle("POST", pattern_Gateway_CreateAppGoodScope_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/CreateAppGoodScope", runtime.WithHTTPPathPattern("/v1/create/appgoodscope"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Gateway_CreateAppGoodScope_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_CreateAppGoodScope_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Gateway_DeleteAppGoodScope_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/DeleteAppGoodScope", runtime.WithHTTPPathPattern("/v1/delete/appgoodscope"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Gateway_DeleteAppGoodScope_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_DeleteAppGoodScope_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> 6712e122f (update scope)
=======
	mux.Handle("POST", pattern_Gateway_GetAppGoodScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes", runtime.WithHTTPPathPattern("/v1/get/appgoodscopes"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Gateway_GetAppGoodScopes_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_GetAppGoodScopes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

<<<<<<< HEAD
>>>>>>> 789287843 (update scope)
=======
>>>>>>> 60afcdcd5 (update scope)
=======
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
>>>>>>> 909af8f64 (update scope)
>>>>>>> 6712e122f (update scope)
	return nil
}

// RegisterGatewayHandlerFromEndpoint is same as RegisterGatewayHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterGatewayHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterGatewayHandler(ctx, mux, conn)
}

// RegisterGatewayHandler registers the http handlers for service Gateway to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterGatewayHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterGatewayHandlerClient(ctx, mux, NewGatewayClient(conn))
}

// RegisterGatewayHandlerClient registers the http handlers for service Gateway
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "GatewayClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "GatewayClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "GatewayClient" to call the correct interceptors.
func RegisterGatewayHandlerClient(ctx context.Context, mux *runtime.ServeMux, client GatewayClient) error {

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> e6bdbe852 (delete createappscope)
	mux.Handle("POST", pattern_Gateway_GetAppGoodScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes", runtime.WithHTTPPathPattern("/v1/get/appgoodscopes"))
=======
	mux.Handle("POST", pattern_Gateway_CreateAppScope_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/CreateAppScope", runtime.WithHTTPPathPattern("/v1/create/appscope"))
<<<<<<< HEAD
>>>>>>> 789287843 (update scope)
=======
	mux.Handle("POST", pattern_Gateway_GetAppGoodScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes", runtime.WithHTTPPathPattern("/v1/get/appgoodscopes"))
>>>>>>> ce3f1742c (delete createappscope)
=======
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
	mux.Handle("POST", pattern_Gateway_GetAppGoodScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes", runtime.WithHTTPPathPattern("/v1/get/appgoodscopes"))
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
		resp, md, err := request_Gateway_GetAppGoodScopes_0(rctx, inboundMarshaler, client, req, pathParams)
=======
		resp, md, err := request_Gateway_CreateAppScope_0(rctx, inboundMarshaler, client, req, pathParams)
>>>>>>> 789287843 (update scope)
=======
		resp, md, err := request_Gateway_GetAppGoodScopes_0(rctx, inboundMarshaler, client, req, pathParams)
>>>>>>> ce3f1742c (delete createappscope)
=======
=======
>>>>>>> e6bdbe852 (delete createappscope)
		resp, md, err := request_Gateway_GetAppGoodScopes_0(rctx, inboundMarshaler, client, req, pathParams)
=======
		resp, md, err := request_Gateway_CreateAppScope_0(rctx, inboundMarshaler, client, req, pathParams)
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
		resp, md, err := request_Gateway_GetAppGoodScopes_0(rctx, inboundMarshaler, client, req, pathParams)
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> e6bdbe852 (delete createappscope)
		forward_Gateway_GetAppGoodScopes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
=======
		forward_Gateway_CreateAppScope_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Gateway_GetAppScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppScopes", runtime.WithHTTPPathPattern("/v1/get/appscopes"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Gateway_GetAppScopes_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_GetAppScopes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
<<<<<<< HEAD
>>>>>>> 789287843 (update scope)
=======
		forward_Gateway_GetAppGoodScopes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
>>>>>>> ce3f1742c (delete createappscope)
=======
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
		forward_Gateway_GetAppGoodScopes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)

	})

	mux.Handle("POST", pattern_Gateway_CreateAppGoodScope_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/CreateAppGoodScope", runtime.WithHTTPPathPattern("/v1/create/appgoodscope"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Gateway_CreateAppGoodScope_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_CreateAppGoodScope_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Gateway_DeleteAppGoodScope_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/DeleteAppGoodScope", runtime.WithHTTPPathPattern("/v1/delete/appgoodscope"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Gateway_DeleteAppGoodScope_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_DeleteAppGoodScope_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> 6712e122f (update scope)
=======
	mux.Handle("POST", pattern_Gateway_GetAppGoodScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes", runtime.WithHTTPPathPattern("/v1/get/appgoodscopes"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Gateway_GetAppGoodScopes_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_GetAppGoodScopes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

<<<<<<< HEAD
>>>>>>> 789287843 (update scope)
=======
>>>>>>> 60afcdcd5 (update scope)
=======
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
>>>>>>> 909af8f64 (update scope)
>>>>>>> 6712e122f (update scope)
	return nil
}

var (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> e6bdbe852 (delete createappscope)
	pattern_Gateway_GetAppGoodScopes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "get", "appgoodscopes"}, ""))
=======
	pattern_Gateway_CreateAppScope_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "create", "appscope"}, ""))

	pattern_Gateway_GetAppScopes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "get", "appscopes"}, ""))
<<<<<<< HEAD
>>>>>>> 789287843 (update scope)
=======
	pattern_Gateway_GetAppGoodScopes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "get", "appgoodscopes"}, ""))
>>>>>>> ce3f1742c (delete createappscope)
=======
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
	pattern_Gateway_GetAppGoodScopes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "get", "appgoodscopes"}, ""))
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)

	pattern_Gateway_CreateAppGoodScope_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "create", "appgoodscope"}, ""))

	pattern_Gateway_DeleteAppGoodScope_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "delete", "appgoodscope"}, ""))
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> 6712e122f (update scope)
)

var (
	forward_Gateway_GetAppGoodScopes_0 = runtime.ForwardResponseMessage
=======

	pattern_Gateway_GetAppGoodScopes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "get", "appgoodscopes"}, ""))
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 60afcdcd5 (update scope)
)

var (
<<<<<<< HEAD
	forward_Gateway_CreateAppScope_0 = runtime.ForwardResponseMessage

	forward_Gateway_GetAppScopes_0 = runtime.ForwardResponseMessage
>>>>>>> 789287843 (update scope)
=======
	forward_Gateway_GetAppGoodScopes_0 = runtime.ForwardResponseMessage
>>>>>>> ce3f1742c (delete createappscope)
=======
=======
=======
>>>>>>> 909af8f64 (update scope)
>>>>>>> 6712e122f (update scope)
)

var (
<<<<<<< HEAD
	forward_Gateway_CreateAppScope_0 = runtime.ForwardResponseMessage

	forward_Gateway_GetAppScopes_0 = runtime.ForwardResponseMessage
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
	forward_Gateway_GetAppGoodScopes_0 = runtime.ForwardResponseMessage
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)

	forward_Gateway_CreateAppGoodScope_0 = runtime.ForwardResponseMessage

	forward_Gateway_DeleteAppGoodScope_0 = runtime.ForwardResponseMessage
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======

	forward_Gateway_GetAppGoodScopes_0 = runtime.ForwardResponseMessage
>>>>>>> 789287843 (update scope)
=======
>>>>>>> 60afcdcd5 (update scope)
=======
=======
>>>>>>> 6712e122f (update scope)
=======

	forward_Gateway_GetAppGoodScopes_0 = runtime.ForwardResponseMessage
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
>>>>>>> 909af8f64 (update scope)
>>>>>>> 6712e122f (update scope)
)
