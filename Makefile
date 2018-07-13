ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

all: phoney

phoney:
	echo "nothing"
	echo $(GOPATH)

youtube-gateway:
	protoc -I. -I$(GOPATH)/src -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. proto/youtubeproxy/youtube.proto
	protoc -I. -I$(GOPATH)/src -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. proto/youtubeproxy/youtube.proto

build-gateway:
	go build -o bin/youtubeproxy.exe cmd/youtubeproxy/youtubeproxy.go

clean:
	# TODO
