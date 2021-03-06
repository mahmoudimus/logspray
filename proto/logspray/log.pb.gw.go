// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/logspray/log.proto

/*
Package logspray is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package logspray

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

func request_LogService_LogStream_0(ctx context.Context, marshaler runtime.Marshaler, client LogServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var metadata runtime.ServerMetadata
	stream, err := client.LogStream(ctx)
	if err != nil {
		grpclog.Printf("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	for {
		var protoReq Message
		err = dec.Decode(&protoReq)
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Printf("Failed to decode request: %v", err)
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
		if err = stream.Send(&protoReq); err != nil {
			grpclog.Printf("Failed to send request: %v", err)
			return nil, metadata, err
		}
	}

	if err := stream.CloseSend(); err != nil {
		grpclog.Printf("Failed to terminate client stream: %v", err)
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		grpclog.Printf("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header

	msg, err := stream.CloseAndRecv()
	metadata.TrailerMD = stream.Trailer()
	return msg, metadata, err

}

func request_LogService_Log_0(ctx context.Context, marshaler runtime.Marshaler, client LogServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Message
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Log(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_LogService_Tail_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_LogService_Tail_0(ctx context.Context, marshaler runtime.Marshaler, client LogServiceClient, req *http.Request, pathParams map[string]string) (LogService_TailClient, runtime.ServerMetadata, error) {
	var protoReq TailRequest
	var metadata runtime.ServerMetadata

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_LogService_Tail_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.Tail(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

var (
	filter_LogService_Search_0 = &utilities.DoubleArray{Encoding: map[string]int{"from": 0, "seconds": 1, "to": 2, "query": 3}, Base: []int{1, 1, 1, 5, 4, 0, 3, 0, 0}, Check: []int{0, 1, 2, 1, 1, 3, 4, 7, 5}}
)

func request_LogService_Search_0(ctx context.Context, marshaler runtime.Marshaler, client LogServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SearchRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["from.seconds"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "from.seconds")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "from.seconds", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "from.seconds", err)
	}

	val, ok = pathParams["to.seconds"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "to.seconds")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "to.seconds", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "to.seconds", err)
	}

	val, ok = pathParams["query"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "query")
	}

	protoReq.Query, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "query", err)
	}

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_LogService_Search_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Search(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_LogService_SearchStream_0 = &utilities.DoubleArray{Encoding: map[string]int{"from": 0, "seconds": 1, "to": 2, "query": 3}, Base: []int{1, 1, 1, 5, 4, 0, 3, 0, 0}, Check: []int{0, 1, 2, 1, 1, 3, 4, 7, 5}}
)

func request_LogService_SearchStream_0(ctx context.Context, marshaler runtime.Marshaler, client LogServiceClient, req *http.Request, pathParams map[string]string) (LogService_SearchStreamClient, runtime.ServerMetadata, error) {
	var protoReq SearchRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["from.seconds"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "from.seconds")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "from.seconds", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "from.seconds", err)
	}

	val, ok = pathParams["to.seconds"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "to.seconds")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "to.seconds", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "to.seconds", err)
	}

	val, ok = pathParams["query"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "query")
	}

	protoReq.Query, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "query", err)
	}

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_LogService_SearchStream_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.SearchStream(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

var (
	filter_LogService_Labels_0 = &utilities.DoubleArray{Encoding: map[string]int{"from": 0, "seconds": 1, "to": 2}, Base: []int{1, 1, 1, 4, 0, 3, 0}, Check: []int{0, 1, 2, 1, 3, 4, 6}}
)

func request_LogService_Labels_0(ctx context.Context, marshaler runtime.Marshaler, client LogServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq LabelsRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["from.seconds"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "from.seconds")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "from.seconds", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "from.seconds", err)
	}

	val, ok = pathParams["to.seconds"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "to.seconds")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "to.seconds", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "to.seconds", err)
	}

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_LogService_Labels_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Labels(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_LogService_LabelValues_0 = &utilities.DoubleArray{Encoding: map[string]int{"from": 0, "seconds": 1, "to": 2, "name": 3}, Base: []int{1, 1, 1, 5, 4, 0, 3, 0, 0}, Check: []int{0, 1, 2, 1, 1, 3, 4, 7, 5}}
)

func request_LogService_LabelValues_0(ctx context.Context, marshaler runtime.Marshaler, client LogServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq LabelValuesRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["from.seconds"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "from.seconds")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "from.seconds", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "from.seconds", err)
	}

	val, ok = pathParams["to.seconds"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "to.seconds")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "to.seconds", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "to.seconds", err)
	}

	val, ok = pathParams["name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "name")
	}

	protoReq.Name, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "name", err)
	}

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_LogService_LabelValues_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.LabelValues(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterLogServiceHandlerFromEndpoint is same as RegisterLogServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterLogServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterLogServiceHandler(ctx, mux, conn)
}

// RegisterLogServiceHandler registers the http handlers for service LogService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterLogServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterLogServiceHandlerClient(ctx, mux, NewLogServiceClient(conn))
}

// RegisterLogServiceHandler registers the http handlers for service LogService to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "LogServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "LogServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "LogServiceClient" to call the correct interceptors.
func RegisterLogServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client LogServiceClient) error {

	mux.Handle("POST", pattern_LogService_LogStream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
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
		resp, md, err := request_LogService_LogStream_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogService_LogStream_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_LogService_Log_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
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
		resp, md, err := request_LogService_Log_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogService_Log_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_LogService_Tail_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
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
		resp, md, err := request_LogService_Tail_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogService_Tail_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_LogService_Search_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
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
		resp, md, err := request_LogService_Search_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogService_Search_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_LogService_SearchStream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
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
		resp, md, err := request_LogService_SearchStream_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogService_SearchStream_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_LogService_Labels_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
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
		resp, md, err := request_LogService_Labels_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogService_Labels_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_LogService_LabelValues_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
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
		resp, md, err := request_LogService_LabelValues_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogService_LabelValues_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_LogService_LogStream_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "log", "stream"}, ""))

	pattern_LogService_Log_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "log"}, ""))

	pattern_LogService_Tail_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "tail", "stream"}, ""))

	pattern_LogService_Search_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 1, 0, 4, 1, 5, 3, 1, 0, 4, 1, 5, 4}, []string{"v1", "search", "from.seconds", "to.seconds", "query"}, ""))

	pattern_LogService_SearchStream_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 1, 0, 4, 1, 5, 3, 1, 0, 4, 1, 5, 4, 2, 5}, []string{"v1", "search", "from.seconds", "to.seconds", "query", "stream"}, ""))

	pattern_LogService_Labels_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 1, 0, 4, 1, 5, 3}, []string{"v1", "labels", "from.seconds", "to.seconds"}, ""))

	pattern_LogService_LabelValues_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 1, 0, 4, 1, 5, 3, 1, 0, 4, 1, 5, 4}, []string{"v1", "labels", "from.seconds", "to.seconds", "name"}, ""))
)

var (
	forward_LogService_LogStream_0 = runtime.ForwardResponseMessage

	forward_LogService_Log_0 = runtime.ForwardResponseMessage

	forward_LogService_Tail_0 = runtime.ForwardResponseStream

	forward_LogService_Search_0 = runtime.ForwardResponseMessage

	forward_LogService_SearchStream_0 = runtime.ForwardResponseStream

	forward_LogService_Labels_0 = runtime.ForwardResponseMessage

	forward_LogService_LabelValues_0 = runtime.ForwardResponseMessage
)
