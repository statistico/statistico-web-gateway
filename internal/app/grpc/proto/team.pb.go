// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.6.1
// source: internal/app/grpc/proto/team.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ShortCode      *wrappers.StringValue `protobuf:"bytes,3,opt,name=short_code,json=shortCode,proto3" json:"short_code,omitempty"`
	CountryId      uint64                `protobuf:"varint,4,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	VenueId        uint64                `protobuf:"varint,5,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	IsNationalTeam *wrappers.BoolValue   `protobuf:"bytes,6,opt,name=is_national_team,json=isNationalTeam,proto3" json:"is_national_team,omitempty"`
	Founded        *wrappers.UInt64Value `protobuf:"bytes,7,opt,name=founded,proto3" json:"founded,omitempty"`
	Logo           *wrappers.StringValue `protobuf:"bytes,8,opt,name=logo,proto3" json:"logo,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_team_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_team_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_team_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Team) GetShortCode() *wrappers.StringValue {
	if x != nil {
		return x.ShortCode
	}
	return nil
}

func (x *Team) GetCountryId() uint64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *Team) GetVenueId() uint64 {
	if x != nil {
		return x.VenueId
	}
	return 0
}

func (x *Team) GetIsNationalTeam() *wrappers.BoolValue {
	if x != nil {
		return x.IsNationalTeam
	}
	return nil
}

func (x *Team) GetFounded() *wrappers.UInt64Value {
	if x != nil {
		return x.Founded
	}
	return nil
}

func (x *Team) GetLogo() *wrappers.StringValue {
	if x != nil {
		return x.Logo
	}
	return nil
}

var File_internal_app_grpc_proto_team_proto protoreflect.FileDescriptor

var file_internal_app_grpc_proto_team_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x02, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x76, 0x65, 0x6e, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x76, 0x65, 0x6e, 0x75, 0x65, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x69,
	0x73, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x36, 0x0a,
	0x07, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x32, 0x81, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x61, 0x6d, 0x42, 0x79, 0x49, 0x44, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x61, 0x6d, 0x73, 0x42, 0x79, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x65,
	0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x22, 0x00, 0x30, 0x01, 0x42, 0x19, 0x5a, 0x17, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_app_grpc_proto_team_proto_rawDescOnce sync.Once
	file_internal_app_grpc_proto_team_proto_rawDescData = file_internal_app_grpc_proto_team_proto_rawDesc
)

func file_internal_app_grpc_proto_team_proto_rawDescGZIP() []byte {
	file_internal_app_grpc_proto_team_proto_rawDescOnce.Do(func() {
		file_internal_app_grpc_proto_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_grpc_proto_team_proto_rawDescData)
	})
	return file_internal_app_grpc_proto_team_proto_rawDescData
}

var file_internal_app_grpc_proto_team_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_app_grpc_proto_team_proto_goTypes = []interface{}{
	(*Team)(nil),                 // 0: proto.Team
	(*wrappers.StringValue)(nil), // 1: google.protobuf.StringValue
	(*wrappers.BoolValue)(nil),   // 2: google.protobuf.BoolValue
	(*wrappers.UInt64Value)(nil), // 3: google.protobuf.UInt64Value
	(*TeamRequest)(nil),          // 4: proto.TeamRequest
	(*SeasonTeamsRequest)(nil),   // 5: proto.SeasonTeamsRequest
}
var file_internal_app_grpc_proto_team_proto_depIdxs = []int32{
	1, // 0: proto.Team.short_code:type_name -> google.protobuf.StringValue
	2, // 1: proto.Team.is_national_team:type_name -> google.protobuf.BoolValue
	3, // 2: proto.Team.founded:type_name -> google.protobuf.UInt64Value
	1, // 3: proto.Team.logo:type_name -> google.protobuf.StringValue
	4, // 4: proto.TeamService.GetTeamByID:input_type -> proto.TeamRequest
	5, // 5: proto.TeamService.GetTeamsBySeasonId:input_type -> proto.SeasonTeamsRequest
	0, // 6: proto.TeamService.GetTeamByID:output_type -> proto.Team
	0, // 7: proto.TeamService.GetTeamsBySeasonId:output_type -> proto.Team
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_app_grpc_proto_team_proto_init() }
func file_internal_app_grpc_proto_team_proto_init() {
	if File_internal_app_grpc_proto_team_proto != nil {
		return
	}
	file_internal_app_grpc_proto_requests_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_app_grpc_proto_team_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Team); i {
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
			RawDescriptor: file_internal_app_grpc_proto_team_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_app_grpc_proto_team_proto_goTypes,
		DependencyIndexes: file_internal_app_grpc_proto_team_proto_depIdxs,
		MessageInfos:      file_internal_app_grpc_proto_team_proto_msgTypes,
	}.Build()
	File_internal_app_grpc_proto_team_proto = out.File
	file_internal_app_grpc_proto_team_proto_rawDesc = nil
	file_internal_app_grpc_proto_team_proto_goTypes = nil
	file_internal_app_grpc_proto_team_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TeamServiceClient is the client API for TeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TeamServiceClient interface {
	GetTeamByID(ctx context.Context, in *TeamRequest, opts ...grpc.CallOption) (*Team, error)
	GetTeamsBySeasonId(ctx context.Context, in *SeasonTeamsRequest, opts ...grpc.CallOption) (TeamService_GetTeamsBySeasonIdClient, error)
}

type teamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamServiceClient(cc grpc.ClientConnInterface) TeamServiceClient {
	return &teamServiceClient{cc}
}

func (c *teamServiceClient) GetTeamByID(ctx context.Context, in *TeamRequest, opts ...grpc.CallOption) (*Team, error) {
	out := new(Team)
	err := c.cc.Invoke(ctx, "/proto.TeamService/GetTeamByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetTeamsBySeasonId(ctx context.Context, in *SeasonTeamsRequest, opts ...grpc.CallOption) (TeamService_GetTeamsBySeasonIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TeamService_serviceDesc.Streams[0], "/proto.TeamService/GetTeamsBySeasonId", opts...)
	if err != nil {
		return nil, err
	}
	x := &teamServiceGetTeamsBySeasonIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TeamService_GetTeamsBySeasonIdClient interface {
	Recv() (*Team, error)
	grpc.ClientStream
}

type teamServiceGetTeamsBySeasonIdClient struct {
	grpc.ClientStream
}

func (x *teamServiceGetTeamsBySeasonIdClient) Recv() (*Team, error) {
	m := new(Team)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TeamServiceServer is the server API for TeamService service.
type TeamServiceServer interface {
	GetTeamByID(context.Context, *TeamRequest) (*Team, error)
	GetTeamsBySeasonId(*SeasonTeamsRequest, TeamService_GetTeamsBySeasonIdServer) error
}

// UnimplementedTeamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTeamServiceServer struct {
}

func (*UnimplementedTeamServiceServer) GetTeamByID(context.Context, *TeamRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamByID not implemented")
}
func (*UnimplementedTeamServiceServer) GetTeamsBySeasonId(*SeasonTeamsRequest, TeamService_GetTeamsBySeasonIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTeamsBySeasonId not implemented")
}

func RegisterTeamServiceServer(s *grpc.Server, srv TeamServiceServer) {
	s.RegisterService(&_TeamService_serviceDesc, srv)
}

func _TeamService_GetTeamByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeamByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TeamService/GetTeamByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeamByID(ctx, req.(*TeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetTeamsBySeasonId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SeasonTeamsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TeamServiceServer).GetTeamsBySeasonId(m, &teamServiceGetTeamsBySeasonIdServer{stream})
}

type TeamService_GetTeamsBySeasonIdServer interface {
	Send(*Team) error
	grpc.ServerStream
}

type teamServiceGetTeamsBySeasonIdServer struct {
	grpc.ServerStream
}

func (x *teamServiceGetTeamsBySeasonIdServer) Send(m *Team) error {
	return x.ServerStream.SendMsg(m)
}

var _TeamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TeamService",
	HandlerType: (*TeamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeamByID",
			Handler:    _TeamService_GetTeamByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTeamsBySeasonId",
			Handler:       _TeamService_GetTeamsBySeasonId_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/app/grpc/proto/team.proto",
}
