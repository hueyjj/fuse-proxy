// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/youtubeproxy/youtube.proto

/*
Package youtube is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package youtube

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_YoutubeService_FindYoutubeMusic_0(ctx context.Context, marshaler runtime.Marshaler, client YoutubeServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq YtMusicRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.FindYoutubeMusic(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_YoutubeService_DownloadYoutubeMusic_0(ctx context.Context, marshaler runtime.Marshaler, client YoutubeServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq YtMusicRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.DownloadYoutubeMusic(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterYoutubeServiceHandlerFromEndpoint is same as RegisterYoutubeServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterYoutubeServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterYoutubeServiceHandler(ctx, mux, conn)
}

// RegisterYoutubeServiceHandler registers the http handlers for service YoutubeService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterYoutubeServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterYoutubeServiceHandlerClient(ctx, mux, NewYoutubeServiceClient(conn))
}

// RegisterYoutubeServiceHandlerClient registers the http handlers for service YoutubeService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "YoutubeServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "YoutubeServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "YoutubeServiceClient" to call the correct interceptors.
func RegisterYoutubeServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client YoutubeServiceClient) error {

	mux.Handle("POST", pattern_YoutubeService_FindYoutubeMusic_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_YoutubeService_FindYoutubeMusic_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_YoutubeService_FindYoutubeMusic_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_YoutubeService_DownloadYoutubeMusic_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_YoutubeService_DownloadYoutubeMusic_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_YoutubeService_DownloadYoutubeMusic_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_YoutubeService_FindYoutubeMusic_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"youtube", "music", "search"}, ""))

	pattern_YoutubeService_DownloadYoutubeMusic_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"youtube", "music", "download"}, ""))
)

var (
	forward_YoutubeService_FindYoutubeMusic_0 = runtime.ForwardResponseMessage

	forward_YoutubeService_DownloadYoutubeMusic_0 = runtime.ForwardResponseMessage
)
