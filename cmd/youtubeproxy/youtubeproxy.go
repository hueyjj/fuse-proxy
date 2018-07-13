package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "github.com/hueyjj/fuse-proxy/proto/youtubeproxy"
)

var (
	youtubeEndpoint = flag.String("youtube_endpoint", "localhost:9090", "endpoint of youtube service")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := gw.RegisterYoutubeServiceHandlerFromEndpoint(ctx, mux, *youtubeEndpoint, opts)
	if err != nil {
		return err
	}

	log.Printf("Starting listening at %s", *youtubeEndpoint)

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
