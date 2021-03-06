// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/youtubeproxy/youtube.proto

package youtube

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Status tells whether or not the music file exists
type YtMusicReply_Status int32

const (
	YtMusicReply_UNKNOWN          YtMusicReply_Status = 0
	YtMusicReply_MUSIC_EXIST      YtMusicReply_Status = 1
	YtMusicReply_MUSIC_NOT_EXIST  YtMusicReply_Status = 2
	YtMusicReply_DOWNLOAD_FAIL    YtMusicReply_Status = 3
	YtMusicReply_DOWNLOAD_SUCCESS YtMusicReply_Status = 4
)

var YtMusicReply_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "MUSIC_EXIST",
	2: "MUSIC_NOT_EXIST",
	3: "DOWNLOAD_FAIL",
	4: "DOWNLOAD_SUCCESS",
}
var YtMusicReply_Status_value = map[string]int32{
	"UNKNOWN":          0,
	"MUSIC_EXIST":      1,
	"MUSIC_NOT_EXIST":  2,
	"DOWNLOAD_FAIL":    3,
	"DOWNLOAD_SUCCESS": 4,
}

func (x YtMusicReply_Status) String() string {
	return proto.EnumName(YtMusicReply_Status_name, int32(x))
}
func (YtMusicReply_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_youtube_ad8766d59382941c, []int{1, 0}
}

type YtMusicRequest struct {
	// Url is the url of the youtube music video
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *YtMusicRequest) Reset()         { *m = YtMusicRequest{} }
func (m *YtMusicRequest) String() string { return proto.CompactTextString(m) }
func (*YtMusicRequest) ProtoMessage()    {}
func (*YtMusicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_youtube_ad8766d59382941c, []int{0}
}
func (m *YtMusicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YtMusicRequest.Unmarshal(m, b)
}
func (m *YtMusicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YtMusicRequest.Marshal(b, m, deterministic)
}
func (dst *YtMusicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YtMusicRequest.Merge(dst, src)
}
func (m *YtMusicRequest) XXX_Size() int {
	return xxx_messageInfo_YtMusicRequest.Size(m)
}
func (m *YtMusicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_YtMusicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_YtMusicRequest proto.InternalMessageInfo

func (m *YtMusicRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type YtMusicReply struct {
	// Name is the name of the music file
	Url                  string              `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Name                 string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               YtMusicReply_Status `protobuf:"varint,3,opt,name=status,proto3,enum=youtube.YtMusicReply_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *YtMusicReply) Reset()         { *m = YtMusicReply{} }
func (m *YtMusicReply) String() string { return proto.CompactTextString(m) }
func (*YtMusicReply) ProtoMessage()    {}
func (*YtMusicReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_youtube_ad8766d59382941c, []int{1}
}
func (m *YtMusicReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YtMusicReply.Unmarshal(m, b)
}
func (m *YtMusicReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YtMusicReply.Marshal(b, m, deterministic)
}
func (dst *YtMusicReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YtMusicReply.Merge(dst, src)
}
func (m *YtMusicReply) XXX_Size() int {
	return xxx_messageInfo_YtMusicReply.Size(m)
}
func (m *YtMusicReply) XXX_DiscardUnknown() {
	xxx_messageInfo_YtMusicReply.DiscardUnknown(m)
}

var xxx_messageInfo_YtMusicReply proto.InternalMessageInfo

func (m *YtMusicReply) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *YtMusicReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *YtMusicReply) GetStatus() YtMusicReply_Status {
	if m != nil {
		return m.Status
	}
	return YtMusicReply_UNKNOWN
}

func init() {
	proto.RegisterType((*YtMusicRequest)(nil), "youtube.YtMusicRequest")
	proto.RegisterType((*YtMusicReply)(nil), "youtube.YtMusicReply")
	proto.RegisterEnum("youtube.YtMusicReply_Status", YtMusicReply_Status_name, YtMusicReply_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// YoutubeServiceClient is the client API for YoutubeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type YoutubeServiceClient interface {
	SearchYoutubeMusic(ctx context.Context, in *YtMusicRequest, opts ...grpc.CallOption) (*YtMusicReply, error)
	DownloadYoutubeMusic(ctx context.Context, in *YtMusicRequest, opts ...grpc.CallOption) (*YtMusicReply, error)
}

type youtubeServiceClient struct {
	cc *grpc.ClientConn
}

func NewYoutubeServiceClient(cc *grpc.ClientConn) YoutubeServiceClient {
	return &youtubeServiceClient{cc}
}

func (c *youtubeServiceClient) SearchYoutubeMusic(ctx context.Context, in *YtMusicRequest, opts ...grpc.CallOption) (*YtMusicReply, error) {
	out := new(YtMusicReply)
	err := c.cc.Invoke(ctx, "/youtube.YoutubeService/SearchYoutubeMusic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *youtubeServiceClient) DownloadYoutubeMusic(ctx context.Context, in *YtMusicRequest, opts ...grpc.CallOption) (*YtMusicReply, error) {
	out := new(YtMusicReply)
	err := c.cc.Invoke(ctx, "/youtube.YoutubeService/DownloadYoutubeMusic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YoutubeServiceServer is the server API for YoutubeService service.
type YoutubeServiceServer interface {
	SearchYoutubeMusic(context.Context, *YtMusicRequest) (*YtMusicReply, error)
	DownloadYoutubeMusic(context.Context, *YtMusicRequest) (*YtMusicReply, error)
}

func RegisterYoutubeServiceServer(s *grpc.Server, srv YoutubeServiceServer) {
	s.RegisterService(&_YoutubeService_serviceDesc, srv)
}

func _YoutubeService_SearchYoutubeMusic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(YtMusicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeServiceServer).SearchYoutubeMusic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/youtube.YoutubeService/SearchYoutubeMusic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeServiceServer).SearchYoutubeMusic(ctx, req.(*YtMusicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YoutubeService_DownloadYoutubeMusic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(YtMusicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeServiceServer).DownloadYoutubeMusic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/youtube.YoutubeService/DownloadYoutubeMusic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeServiceServer).DownloadYoutubeMusic(ctx, req.(*YtMusicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _YoutubeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "youtube.YoutubeService",
	HandlerType: (*YoutubeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchYoutubeMusic",
			Handler:    _YoutubeService_SearchYoutubeMusic_Handler,
		},
		{
			MethodName: "DownloadYoutubeMusic",
			Handler:    _YoutubeService_DownloadYoutubeMusic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/youtubeproxy/youtube.proto",
}

func init() {
	proto.RegisterFile("proto/youtubeproxy/youtube.proto", fileDescriptor_youtube_ad8766d59382941c)
}

var fileDescriptor_youtube_ad8766d59382941c = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0xff, 0x0b, 0x04, 0xf2, 0x1f, 0x14, 0xea, 0x0a, 0x81, 0x10, 0x0e, 0xcd, 0x9e, 0x08,
	0x07, 0x9a, 0xa0, 0x27, 0x6e, 0x04, 0x30, 0x21, 0x42, 0x9b, 0xb0, 0x10, 0xe4, 0x44, 0x16, 0x58,
	0xb1, 0xa6, 0x74, 0x6b, 0xbb, 0x55, 0xb9, 0xfa, 0x0a, 0xbe, 0x98, 0x89, 0xaf, 0xc0, 0x83, 0x18,
	0xb6, 0x85, 0x18, 0xf1, 0xe6, 0x6d, 0xe6, 0x37, 0x5f, 0xe6, 0xfb, 0x76, 0x07, 0x74, 0xcf, 0x17,
	0x52, 0x18, 0x5b, 0x11, 0xca, 0x70, 0xc1, 0x3d, 0x5f, 0xbc, 0x6e, 0x0f, 0x4d, 0x43, 0x8d, 0x70,
	0x26, 0x6e, 0x2b, 0xd5, 0xb5, 0x10, 0x6b, 0x87, 0x1b, 0xcc, 0xb3, 0x0d, 0xe6, 0xba, 0x42, 0x32,
	0x69, 0x0b, 0x37, 0x88, 0x64, 0x84, 0x40, 0x6e, 0x26, 0x87, 0x61, 0x60, 0x2f, 0x47, 0xfc, 0x29,
	0xe4, 0x81, 0xc4, 0x1a, 0x24, 0x43, 0xdf, 0x29, 0x23, 0x1d, 0xd5, 0xfe, 0x8f, 0xf6, 0x25, 0xf9,
	0x40, 0x70, 0x76, 0x14, 0x79, 0xce, 0xf6, 0x54, 0x82, 0x31, 0xa4, 0x5c, 0xb6, 0xe1, 0xe5, 0x84,
	0x42, 0xaa, 0xc6, 0xd7, 0x90, 0x0e, 0x24, 0x93, 0x61, 0x50, 0x4e, 0xea, 0xa8, 0x96, 0x6b, 0x56,
	0x1b, 0x87, 0x84, 0xdf, 0x97, 0x35, 0xa8, 0xd2, 0x8c, 0x62, 0x2d, 0x59, 0x41, 0x3a, 0x22, 0x38,
	0x0b, 0x99, 0x89, 0x79, 0x6b, 0x5a, 0x53, 0x53, 0xfb, 0x87, 0xf3, 0x90, 0x1d, 0x4e, 0x68, 0xbf,
	0x33, 0xef, 0xdd, 0xf5, 0xe9, 0x58, 0x43, 0xf8, 0x12, 0xf2, 0x11, 0x30, 0xad, 0x71, 0x0c, 0x13,
	0xf8, 0x02, 0xce, 0xbb, 0xd6, 0xd4, 0x1c, 0x58, 0xed, 0xee, 0xfc, 0xa6, 0xdd, 0x1f, 0x68, 0x49,
	0x5c, 0x00, 0xed, 0x88, 0xe8, 0xa4, 0xd3, 0xe9, 0x51, 0xaa, 0xa5, 0x9a, 0x3b, 0x04, 0xb9, 0x59,
	0x94, 0x86, 0x72, 0xff, 0xd9, 0x5e, 0x72, 0x7c, 0x0f, 0x98, 0x72, 0xe6, 0x2f, 0x1f, 0x62, 0xae,
	0x22, 0xe2, 0xd2, 0x69, 0x68, 0xf5, 0x4d, 0x95, 0xe2, 0xaf, 0xaf, 0x21, 0xfa, 0xdb, 0xe7, 0xee,
	0x3d, 0x51, 0x21, 0xc5, 0xc3, 0x39, 0x8c, 0xcd, 0x7e, 0x68, 0x04, 0x6a, 0x75, 0x0b, 0xd5, 0xf1,
	0x23, 0x14, 0xba, 0xe2, 0xc5, 0x75, 0x04, 0x5b, 0xfd, 0xc9, 0x89, 0x28, 0xa7, 0x2a, 0x29, 0xfd,
	0x70, 0x5a, 0xc5, 0xcb, 0x5b, 0xa8, 0xbe, 0x48, 0xab, 0x23, 0x5f, 0x7d, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x4e, 0x55, 0xce, 0x6b, 0x2f, 0x02, 0x00, 0x00,
}
