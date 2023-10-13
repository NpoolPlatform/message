// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: npool/inspire/gw/v1/coupon/scope/scope.proto

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

func request_Gateway_CreateScope_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateScopeRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateScope(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Gateway_CreateScope_0(ctx context.Context, marshaler runtime.Marshaler, server GatewayServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateScopeRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CreateScope(ctx, &protoReq)
	return msg, metadata, err

}

func request_Gateway_DeleteScope_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DeleteScopeRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.DeleteScope(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Gateway_DeleteScope_0(ctx context.Context, marshaler runtime.Marshaler, server GatewayServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DeleteScopeRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.DeleteScope(ctx, &protoReq)
	return msg, metadata, err

}

func request_Gateway_GetScopes_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetScopesRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetScopes(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Gateway_GetScopes_0(ctx context.Context, marshaler runtime.Marshaler, server GatewayServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetScopesRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetScopes(ctx, &protoReq)
	return msg, metadata, err

}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b738f188 (update scope)
=======
=======
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
=======
>>>>>>> 1fb4844cc (update scope)
>>>>>>> 509fa36c7 (update scope)
=======
=======
>>>>>>> 343ac084f (delete get app scopes)
=======
=======
>>>>>>> 1fb4844cc (update scope)
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
>>>>>>> c3156b2f2 (add coupon scope in gw)
func request_Gateway_GetAppScopes_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAppScopesRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetAppScopes(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Gateway_GetAppScopes_0(ctx context.Context, marshaler runtime.Marshaler, server GatewayServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAppScopesRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetAppScopes(ctx, &protoReq)
	return msg, metadata, err

}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> f213f6ed2 (add coupon scope in gw)
=======
>>>>>>> 447c10d15 (delete get app scopes)
=======
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b0ed6935 (delete get app scopes)
=======
>>>>>>> 4b738f188 (update scope)
=======
func request_Gateway_GetNAppScopes_0(ctx context.Context, marshaler runtime.Marshaler, client GatewayClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetNAppScopesRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetNAppScopes(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Gateway_GetNAppScopes_0(ctx context.Context, marshaler runtime.Marshaler, server GatewayServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetNAppScopesRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetNAppScopes(ctx, &protoReq)
	return msg, metadata, err

}

>>>>>>> 36e1391e1 (update scope pb)
=======
>>>>>>> 789287843 (update scope)
=======
=======
>>>>>>> 509fa36c7 (update scope)
=======
>>>>>>> c3156b2f2 (add coupon scope in gw)
>>>>>>> e243bad0c (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
>>>>>>> feba7689e (delete get app scopes)
<<<<<<< HEAD
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
>>>>>>> 1fb4844cc (update scope)
<<<<<<< HEAD
>>>>>>> 509fa36c7 (update scope)
=======
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 67de434f0 (delete get app scopes)
>>>>>>> 343ac084f (delete get app scopes)
// RegisterGatewayHandlerServer registers the http handlers for service Gateway to "mux".
// UnaryRPC     :call GatewayServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterGatewayHandlerFromEndpoint instead.
func RegisterGatewayHandlerServer(ctx context.Context, mux *runtime.ServeMux, server GatewayServer) error {

	mux.Handle("POST", pattern_Gateway_CreateScope_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope", runtime.WithHTTPPathPattern("/v1/create/scope"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Gateway_CreateScope_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_CreateScope_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Gateway_DeleteScope_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope", runtime.WithHTTPPathPattern("/v1/delete/scope"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Gateway_DeleteScope_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_DeleteScope_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Gateway_GetScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes", runtime.WithHTTPPathPattern("/v1/get/scopes"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Gateway_GetScopes_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_GetScopes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b738f188 (update scope)
=======
=======
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
=======
>>>>>>> 1fb4844cc (update scope)
>>>>>>> 509fa36c7 (update scope)
=======
=======
>>>>>>> 343ac084f (delete get app scopes)
=======
=======
>>>>>>> 1fb4844cc (update scope)
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
>>>>>>> c3156b2f2 (add coupon scope in gw)
	mux.Handle("POST", pattern_Gateway_GetAppScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.scope.v1.Gateway/GetAppScopes", runtime.WithHTTPPathPattern("/v1/get/app/scopes"))
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

	})

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> f213f6ed2 (add coupon scope in gw)
=======
>>>>>>> 447c10d15 (delete get app scopes)
=======
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b0ed6935 (delete get app scopes)
=======
>>>>>>> 4b738f188 (update scope)
=======
	mux.Handle("POST", pattern_Gateway_GetNAppScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/inspire.gateway.coupon.scope.v1.Gateway/GetNAppScopes", runtime.WithHTTPPathPattern("/v1/get/n/app/scopes"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Gateway_GetNAppScopes_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_GetNAppScopes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

>>>>>>> 36e1391e1 (update scope pb)
=======
>>>>>>> 789287843 (update scope)
=======
=======
>>>>>>> 509fa36c7 (update scope)
=======
>>>>>>> c3156b2f2 (add coupon scope in gw)
>>>>>>> e243bad0c (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
>>>>>>> feba7689e (delete get app scopes)
<<<<<<< HEAD
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
>>>>>>> 1fb4844cc (update scope)
<<<<<<< HEAD
>>>>>>> 509fa36c7 (update scope)
=======
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 67de434f0 (delete get app scopes)
>>>>>>> 343ac084f (delete get app scopes)
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

	mux.Handle("POST", pattern_Gateway_CreateScope_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope", runtime.WithHTTPPathPattern("/v1/create/scope"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Gateway_CreateScope_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_CreateScope_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Gateway_DeleteScope_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope", runtime.WithHTTPPathPattern("/v1/delete/scope"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Gateway_DeleteScope_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_DeleteScope_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Gateway_GetScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes", runtime.WithHTTPPathPattern("/v1/get/scopes"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Gateway_GetScopes_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_GetScopes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b738f188 (update scope)
=======
=======
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
=======
>>>>>>> 1fb4844cc (update scope)
>>>>>>> 509fa36c7 (update scope)
=======
=======
>>>>>>> 343ac084f (delete get app scopes)
=======
=======
>>>>>>> 1fb4844cc (update scope)
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
>>>>>>> c3156b2f2 (add coupon scope in gw)
	mux.Handle("POST", pattern_Gateway_GetAppScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/inspire.gateway.coupon.scope.v1.Gateway/GetAppScopes", runtime.WithHTTPPathPattern("/v1/get/app/scopes"))
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

	})

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> f213f6ed2 (add coupon scope in gw)
=======
>>>>>>> 447c10d15 (delete get app scopes)
=======
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b0ed6935 (delete get app scopes)
=======
>>>>>>> 4b738f188 (update scope)
=======
	mux.Handle("POST", pattern_Gateway_GetNAppScopes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/inspire.gateway.coupon.scope.v1.Gateway/GetNAppScopes", runtime.WithHTTPPathPattern("/v1/get/n/app/scopes"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Gateway_GetNAppScopes_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Gateway_GetNAppScopes_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

>>>>>>> 36e1391e1 (update scope pb)
=======
>>>>>>> 789287843 (update scope)
=======
=======
>>>>>>> 509fa36c7 (update scope)
=======
>>>>>>> c3156b2f2 (add coupon scope in gw)
>>>>>>> e243bad0c (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
>>>>>>> feba7689e (delete get app scopes)
<<<<<<< HEAD
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
>>>>>>> 1fb4844cc (update scope)
<<<<<<< HEAD
>>>>>>> 509fa36c7 (update scope)
=======
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 67de434f0 (delete get app scopes)
>>>>>>> 343ac084f (delete get app scopes)
	return nil
}

var (
	pattern_Gateway_CreateScope_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "create", "scope"}, ""))

	pattern_Gateway_DeleteScope_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "delete", "scope"}, ""))

	pattern_Gateway_GetScopes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "get", "scopes"}, ""))
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======

	pattern_Gateway_GetAppScopes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v1", "get", "app", "scopes"}, ""))
<<<<<<< HEAD
>>>>>>> f213f6ed2 (add coupon scope in gw)
=======
>>>>>>> 447c10d15 (delete get app scopes)
=======

	pattern_Gateway_GetAppScopes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v1", "get", "app", "scopes"}, ""))
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b0ed6935 (delete get app scopes)
=======

	pattern_Gateway_GetAppScopes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v1", "get", "app", "scopes"}, ""))
>>>>>>> 4b738f188 (update scope)
=======

	pattern_Gateway_GetNAppScopes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"v1", "get", "n", "app", "scopes"}, ""))
>>>>>>> 36e1391e1 (update scope pb)
=======
>>>>>>> 789287843 (update scope)
=======
=======
>>>>>>> e8948f8dc (delete get app scopes)
=======
>>>>>>> 509fa36c7 (update scope)
=======
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
>>>>>>> 343ac084f (delete get app scopes)
=======

	pattern_Gateway_GetAppScopes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v1", "get", "app", "scopes"}, ""))
>>>>>>> e243bad0c (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
>>>>>>> feba7689e (delete get app scopes)
<<<<<<< HEAD
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======

	pattern_Gateway_GetAppScopes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v1", "get", "app", "scopes"}, ""))
>>>>>>> 1fb4844cc (update scope)
<<<<<<< HEAD
>>>>>>> 509fa36c7 (update scope)
=======
=======

	pattern_Gateway_GetAppScopes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v1", "get", "app", "scopes"}, ""))
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 67de434f0 (delete get app scopes)
>>>>>>> 343ac084f (delete get app scopes)
)

var (
	forward_Gateway_CreateScope_0 = runtime.ForwardResponseMessage

	forward_Gateway_DeleteScope_0 = runtime.ForwardResponseMessage

	forward_Gateway_GetScopes_0 = runtime.ForwardResponseMessage
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======

	forward_Gateway_GetAppScopes_0 = runtime.ForwardResponseMessage
<<<<<<< HEAD
>>>>>>> f213f6ed2 (add coupon scope in gw)
=======
>>>>>>> 447c10d15 (delete get app scopes)
=======

	forward_Gateway_GetAppScopes_0 = runtime.ForwardResponseMessage
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b0ed6935 (delete get app scopes)
=======

	forward_Gateway_GetAppScopes_0 = runtime.ForwardResponseMessage
>>>>>>> 4b738f188 (update scope)
=======

	forward_Gateway_GetNAppScopes_0 = runtime.ForwardResponseMessage
>>>>>>> 36e1391e1 (update scope pb)
=======
>>>>>>> 789287843 (update scope)
=======
=======
>>>>>>> e8948f8dc (delete get app scopes)
=======
>>>>>>> 509fa36c7 (update scope)
=======
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
>>>>>>> 343ac084f (delete get app scopes)
=======

	forward_Gateway_GetAppScopes_0 = runtime.ForwardResponseMessage
>>>>>>> e243bad0c (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
>>>>>>> feba7689e (delete get app scopes)
<<<<<<< HEAD
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======

	forward_Gateway_GetAppScopes_0 = runtime.ForwardResponseMessage
>>>>>>> 1fb4844cc (update scope)
<<<<<<< HEAD
>>>>>>> 509fa36c7 (update scope)
=======
=======

	forward_Gateway_GetAppScopes_0 = runtime.ForwardResponseMessage
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 67de434f0 (delete get app scopes)
>>>>>>> 343ac084f (delete get app scopes)
)
