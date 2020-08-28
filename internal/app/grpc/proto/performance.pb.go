// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.6.1
// source: internal/app/grpc/proto/performance.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TeamStatPerformanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action  string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Games   uint32   `protobuf:"varint,2,opt,name=games,proto3" json:"games,omitempty"`
	Measure string   `protobuf:"bytes,3,opt,name=measure,proto3" json:"measure,omitempty"`
	Metric  string   `protobuf:"bytes,4,opt,name=metric,proto3" json:"metric,omitempty"`
	Seasons []uint64 `protobuf:"varint,5,rep,packed,name=seasons,proto3" json:"seasons,omitempty"`
	Stat    string   `protobuf:"bytes,6,opt,name=stat,proto3" json:"stat,omitempty"`
	Value   float32  `protobuf:"fixed32,7,opt,name=value,proto3" json:"value,omitempty"`
	Venue   string   `protobuf:"bytes,8,opt,name=venue,proto3" json:"venue,omitempty"`
}

func (x *TeamStatPerformanceRequest) Reset() {
	*x = TeamStatPerformanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_performance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamStatPerformanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamStatPerformanceRequest) ProtoMessage() {}

func (x *TeamStatPerformanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_performance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamStatPerformanceRequest.ProtoReflect.Descriptor instead.
func (*TeamStatPerformanceRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_performance_proto_rawDescGZIP(), []int{0}
}

func (x *TeamStatPerformanceRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *TeamStatPerformanceRequest) GetGames() uint32 {
	if x != nil {
		return x.Games
	}
	return 0
}

func (x *TeamStatPerformanceRequest) GetMeasure() string {
	if x != nil {
		return x.Measure
	}
	return ""
}

func (x *TeamStatPerformanceRequest) GetMetric() string {
	if x != nil {
		return x.Metric
	}
	return ""
}

func (x *TeamStatPerformanceRequest) GetSeasons() []uint64 {
	if x != nil {
		return x.Seasons
	}
	return nil
}

func (x *TeamStatPerformanceRequest) GetStat() string {
	if x != nil {
		return x.Stat
	}
	return ""
}

func (x *TeamStatPerformanceRequest) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TeamStatPerformanceRequest) GetVenue() string {
	if x != nil {
		return x.Venue
	}
	return ""
}

type TeamStatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teams []*Team `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
}

func (x *TeamStatResponse) Reset() {
	*x = TeamStatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_performance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamStatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamStatResponse) ProtoMessage() {}

func (x *TeamStatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_performance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamStatResponse.ProtoReflect.Descriptor instead.
func (*TeamStatResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_performance_proto_rawDescGZIP(), []int{1}
}

func (x *TeamStatResponse) GetTeams() []*Team {
	if x != nil {
		return x.Teams
	}
	return nil
}

var File_internal_app_grpc_proto_performance_proto protoreflect.FileDescriptor

var file_internal_app_grpc_proto_performance_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a, 0x1a, 0x54, 0x65, 0x61, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x74, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x65, 0x6e,
	0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x22,
	0x35, 0x0a, 0x10, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52,
	0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x32, 0x6a, 0x0a, 0x12, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_app_grpc_proto_performance_proto_rawDescOnce sync.Once
	file_internal_app_grpc_proto_performance_proto_rawDescData = file_internal_app_grpc_proto_performance_proto_rawDesc
)

func file_internal_app_grpc_proto_performance_proto_rawDescGZIP() []byte {
	file_internal_app_grpc_proto_performance_proto_rawDescOnce.Do(func() {
		file_internal_app_grpc_proto_performance_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_grpc_proto_performance_proto_rawDescData)
	})
	return file_internal_app_grpc_proto_performance_proto_rawDescData
}

var file_internal_app_grpc_proto_performance_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_app_grpc_proto_performance_proto_goTypes = []interface{}{
	(*TeamStatPerformanceRequest)(nil), // 0: proto.TeamStatPerformanceRequest
	(*TeamStatResponse)(nil),           // 1: proto.TeamStatResponse
	(*Team)(nil),                       // 2: proto.Team
}
var file_internal_app_grpc_proto_performance_proto_depIdxs = []int32{
	2, // 0: proto.TeamStatResponse.teams:type_name -> proto.Team
	0, // 1: proto.PerformanceService.GetTeamsMatchingStat:input_type -> proto.TeamStatPerformanceRequest
	1, // 2: proto.PerformanceService.GetTeamsMatchingStat:output_type -> proto.TeamStatResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_app_grpc_proto_performance_proto_init() }
func file_internal_app_grpc_proto_performance_proto_init() {
	if File_internal_app_grpc_proto_performance_proto != nil {
		return
	}
	file_internal_app_grpc_proto_team_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_app_grpc_proto_performance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamStatPerformanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_app_grpc_proto_performance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamStatResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_app_grpc_proto_performance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_app_grpc_proto_performance_proto_goTypes,
		DependencyIndexes: file_internal_app_grpc_proto_performance_proto_depIdxs,
		MessageInfos:      file_internal_app_grpc_proto_performance_proto_msgTypes,
	}.Build()
	File_internal_app_grpc_proto_performance_proto = out.File
	file_internal_app_grpc_proto_performance_proto_rawDesc = nil
	file_internal_app_grpc_proto_performance_proto_goTypes = nil
	file_internal_app_grpc_proto_performance_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PerformanceServiceClient is the client API for PerformanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PerformanceServiceClient interface {
	GetTeamsMatchingStat(ctx context.Context, in *TeamStatPerformanceRequest, opts ...grpc.CallOption) (*TeamStatResponse, error)
}

type performanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPerformanceServiceClient(cc grpc.ClientConnInterface) PerformanceServiceClient {
	return &performanceServiceClient{cc}
}

func (c *performanceServiceClient) GetTeamsMatchingStat(ctx context.Context, in *TeamStatPerformanceRequest, opts ...grpc.CallOption) (*TeamStatResponse, error) {
	out := new(TeamStatResponse)
	err := c.cc.Invoke(ctx, "/proto.PerformanceService/GetTeamsMatchingStat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PerformanceServiceServer is the server API for PerformanceService service.
type PerformanceServiceServer interface {
	GetTeamsMatchingStat(context.Context, *TeamStatPerformanceRequest) (*TeamStatResponse, error)
}

// UnimplementedPerformanceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPerformanceServiceServer struct {
}

func (*UnimplementedPerformanceServiceServer) GetTeamsMatchingStat(context.Context, *TeamStatPerformanceRequest) (*TeamStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamsMatchingStat not implemented")
}

func RegisterPerformanceServiceServer(s *grpc.Server, srv PerformanceServiceServer) {
	s.RegisterService(&_PerformanceService_serviceDesc, srv)
}

func _PerformanceService_GetTeamsMatchingStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamStatPerformanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceServiceServer).GetTeamsMatchingStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PerformanceService/GetTeamsMatchingStat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceServiceServer).GetTeamsMatchingStat(ctx, req.(*TeamStatPerformanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PerformanceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PerformanceService",
	HandlerType: (*PerformanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeamsMatchingStat",
			Handler:    _PerformanceService_GetTeamsMatchingStat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/grpc/proto/performance.proto",
}
