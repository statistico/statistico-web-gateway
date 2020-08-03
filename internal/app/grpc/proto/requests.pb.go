// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.6.1
// source: internal/app/grpc/proto/requests.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type CompetitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A filter to limit the competitions returned associated to a specific countries
	CountryIds []uint64 `protobuf:"varint,1,rep,packed,name=country_ids,json=countryIds,proto3" json:"country_ids,omitempty"`
	// Order the ID column to return competitions in specific order
	Sort *wrappers.StringValue `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty"`
	// A filter to limit the competitions returned depending on if they are a cup competition or not
	IsCup *wrappers.BoolValue `protobuf:"bytes,3,opt,name=is_cup,json=isCup,proto3" json:"is_cup,omitempty"`
}

func (x *CompetitionRequest) Reset() {
	*x = CompetitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompetitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompetitionRequest) ProtoMessage() {}

func (x *CompetitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompetitionRequest.ProtoReflect.Descriptor instead.
func (*CompetitionRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_requests_proto_rawDescGZIP(), []int{0}
}

func (x *CompetitionRequest) GetCountryIds() []uint64 {
	if x != nil {
		return x.CountryIds
	}
	return nil
}

func (x *CompetitionRequest) GetSort() *wrappers.StringValue {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *CompetitionRequest) GetIsCup() *wrappers.BoolValue {
	if x != nil {
		return x.IsCup
	}
	return nil
}

type FixtureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FixtureId uint64 `protobuf:"varint,1,opt,name=fixture_id,json=fixtureId,proto3" json:"fixture_id,omitempty"`
}

func (x *FixtureRequest) Reset() {
	*x = FixtureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FixtureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FixtureRequest) ProtoMessage() {}

func (x *FixtureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FixtureRequest.ProtoReflect.Descriptor instead.
func (*FixtureRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_requests_proto_rawDescGZIP(), []int{1}
}

func (x *FixtureRequest) GetFixtureId() uint64 {
	if x != nil {
		return x.FixtureId
	}
	return 0
}

type HistoricalResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Home Team ID that the Result set relates to
	HomeTeamId uint64 `protobuf:"varint,1,opt,name=home_team_id,json=homeTeamId,proto3" json:"home_team_id,omitempty"`
	// The Away Team ID that the Result set relates to
	AwayTeamId uint64 `protobuf:"varint,2,opt,name=away_team_id,json=awayTeamId,proto3" json:"away_team_id,omitempty"`
	// The number of results to return
	Limit uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// A filter to return Results before a specific date
	// RFC3339 formatted string i.e "2006-01-02T15:04:05Z07:00"
	DateBefore string `protobuf:"bytes,4,opt,name=date_before,json=dateBefore,proto3" json:"date_before,omitempty"`
}

func (x *HistoricalResultRequest) Reset() {
	*x = HistoricalResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoricalResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoricalResultRequest) ProtoMessage() {}

func (x *HistoricalResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoricalResultRequest.ProtoReflect.Descriptor instead.
func (*HistoricalResultRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_requests_proto_rawDescGZIP(), []int{2}
}

func (x *HistoricalResultRequest) GetHomeTeamId() uint64 {
	if x != nil {
		return x.HomeTeamId
	}
	return 0
}

func (x *HistoricalResultRequest) GetAwayTeamId() uint64 {
	if x != nil {
		return x.AwayTeamId
	}
	return 0
}

func (x *HistoricalResultRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *HistoricalResultRequest) GetDateBefore() string {
	if x != nil {
		return x.DateBefore
	}
	return ""
}

type SeasonFixtureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonId uint64 `protobuf:"varint,1,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
	// RFC3339 formatted string i.e. "2006-01-02T15:04:05Z07:00"
	DateFrom string `protobuf:"bytes,2,opt,name=date_from,json=dateFrom,proto3" json:"date_from,omitempty"`
	// RFC3339 formatted string i.e "2006-01-02T15:04:05Z07:00"
	DateTo string `protobuf:"bytes,3,opt,name=date_to,json=dateTo,proto3" json:"date_to,omitempty"`
}

func (x *SeasonFixtureRequest) Reset() {
	*x = SeasonFixtureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeasonFixtureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeasonFixtureRequest) ProtoMessage() {}

func (x *SeasonFixtureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeasonFixtureRequest.ProtoReflect.Descriptor instead.
func (*SeasonFixtureRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_requests_proto_rawDescGZIP(), []int{3}
}

func (x *SeasonFixtureRequest) GetSeasonId() uint64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *SeasonFixtureRequest) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *SeasonFixtureRequest) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

type SeasonCompetitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompetitionId uint64 `protobuf:"varint,1,opt,name=competition_id,json=competitionId,proto3" json:"competition_id,omitempty"`
	// Order the name column to return seasons in specific order
	Sort *wrappers.StringValue `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *SeasonCompetitionRequest) Reset() {
	*x = SeasonCompetitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeasonCompetitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeasonCompetitionRequest) ProtoMessage() {}

func (x *SeasonCompetitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeasonCompetitionRequest.ProtoReflect.Descriptor instead.
func (*SeasonCompetitionRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_requests_proto_rawDescGZIP(), []int{4}
}

func (x *SeasonCompetitionRequest) GetCompetitionId() uint64 {
	if x != nil {
		return x.CompetitionId
	}
	return 0
}

func (x *SeasonCompetitionRequest) GetSort() *wrappers.StringValue {
	if x != nil {
		return x.Sort
	}
	return nil
}

type SeasonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Season ID that the Result set relates to
	SeasonId int64 `protobuf:"varint,1,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
	// A filter to return Results before a specific date
	// RFC3339 formatted string i.e "2006-01-02T15:04:05Z07:00"
	DateBefore string `protobuf:"bytes,2,opt,name=date_before,json=dateBefore,proto3" json:"date_before,omitempty"`
}

func (x *SeasonRequest) Reset() {
	*x = SeasonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeasonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeasonRequest) ProtoMessage() {}

func (x *SeasonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeasonRequest.ProtoReflect.Descriptor instead.
func (*SeasonRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_requests_proto_rawDescGZIP(), []int{5}
}

func (x *SeasonRequest) GetSeasonId() int64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *SeasonRequest) GetDateBefore() string {
	if x != nil {
		return x.DateBefore
	}
	return ""
}

type SeasonTeamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeasonId uint64 `protobuf:"varint,1,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
}

func (x *SeasonTeamsRequest) Reset() {
	*x = SeasonTeamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeasonTeamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeasonTeamsRequest) ProtoMessage() {}

func (x *SeasonTeamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeasonTeamsRequest.ProtoReflect.Descriptor instead.
func (*SeasonTeamsRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_requests_proto_rawDescGZIP(), []int{6}
}

func (x *SeasonTeamsRequest) GetSeasonId() uint64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

type TeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId uint64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
}

func (x *TeamRequest) Reset() {
	*x = TeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamRequest) ProtoMessage() {}

func (x *TeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamRequest.ProtoReflect.Descriptor instead.
func (*TeamRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_requests_proto_rawDescGZIP(), []int{7}
}

func (x *TeamRequest) GetTeamId() uint64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

type TeamResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Team ID that the Result set relates to
	TeamId uint64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	// The number of results to return. If limit is not set the whole Result set for the Team
	// will be returned
	Limit *wrappers.UInt64Value `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// A filter to return Results before a specific date
	// RFC3339 formatted string i.e "2006-01-02T15:04:05Z07:00"
	DateBefore *wrappers.StringValue `protobuf:"bytes,3,opt,name=date_before,json=dateBefore,proto3" json:"date_before,omitempty"`
	// A filter to return Results after a specific date
	// RFC3339 formatted string i.e "2006-01-02T15:04:05Z07:00"
	DateAfter *wrappers.StringValue `protobuf:"bytes,4,opt,name=date_after,json=dateAfter,proto3" json:"date_after,omitempty"`
	// A filter to return based limited to either home or away results
	Venue *wrappers.StringValue `protobuf:"bytes,5,opt,name=venue,proto3" json:"venue,omitempty"`
	// A filter to limit the results returned associated to a specific season
	SeasonIds []uint64 `protobuf:"varint,6,rep,packed,name=season_ids,json=seasonIds,proto3" json:"season_ids,omitempty"`
	// Order the date column to return results in specific order
	Sort *wrappers.StringValue `protobuf:"bytes,7,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *TeamResultRequest) Reset() {
	*x = TeamResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamResultRequest) ProtoMessage() {}

func (x *TeamResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamResultRequest.ProtoReflect.Descriptor instead.
func (*TeamResultRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_requests_proto_rawDescGZIP(), []int{8}
}

func (x *TeamResultRequest) GetTeamId() uint64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *TeamResultRequest) GetLimit() *wrappers.UInt64Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *TeamResultRequest) GetDateBefore() *wrappers.StringValue {
	if x != nil {
		return x.DateBefore
	}
	return nil
}

func (x *TeamResultRequest) GetDateAfter() *wrappers.StringValue {
	if x != nil {
		return x.DateAfter
	}
	return nil
}

func (x *TeamResultRequest) GetVenue() *wrappers.StringValue {
	if x != nil {
		return x.Venue
	}
	return nil
}

func (x *TeamResultRequest) GetSeasonIds() []uint64 {
	if x != nil {
		return x.SeasonIds
	}
	return nil
}

func (x *TeamResultRequest) GetSort() *wrappers.StringValue {
	if x != nil {
		return x.Sort
	}
	return nil
}

type TeamSeasonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId uint64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	// Order the name column to return seasons in specific order
	Sort *wrappers.StringValue `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *TeamSeasonsRequest) Reset() {
	*x = TeamSeasonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamSeasonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamSeasonsRequest) ProtoMessage() {}

func (x *TeamSeasonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_requests_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamSeasonsRequest.ProtoReflect.Descriptor instead.
func (*TeamSeasonsRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_requests_proto_rawDescGZIP(), []int{9}
}

func (x *TeamSeasonsRequest) GetTeamId() uint64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *TeamSeasonsRequest) GetSort() *wrappers.StringValue {
	if x != nil {
		return x.Sort
	}
	return nil
}

var File_internal_app_grpc_proto_requests_proto protoreflect.FileDescriptor

var file_internal_app_grpc_proto_requests_proto_rawDesc = []byte{
	0x0a, 0x26, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9a, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x69, 0x73, 0x5f,
	0x63, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x69, 0x73, 0x43, 0x75, 0x70, 0x22, 0x2f, 0x0a, 0x0e,
	0x46, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x22, 0x94, 0x01,
	0x0a, 0x17, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x68, 0x6f, 0x6d,
	0x65, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x61,
	0x77, 0x61, 0x79, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x61, 0x77, 0x61, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x42, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x22, 0x69, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x46, 0x69,
	0x78, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x22,
	0x73, 0x0a, 0x18, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x22, 0x4d, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x42, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x22, 0x31, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x65, 0x61,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0xe1,
	0x02, 0x0a, 0x11, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x32, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x3d, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x12, 0x3b, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a,
	0x05, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x65, 0x6e, 0x75,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x73,
	0x12, 0x30, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x22, 0x5f, 0x0a, 0x12, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x12, 0x30, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x42, 0x19, 0x5a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_app_grpc_proto_requests_proto_rawDescOnce sync.Once
	file_internal_app_grpc_proto_requests_proto_rawDescData = file_internal_app_grpc_proto_requests_proto_rawDesc
)

func file_internal_app_grpc_proto_requests_proto_rawDescGZIP() []byte {
	file_internal_app_grpc_proto_requests_proto_rawDescOnce.Do(func() {
		file_internal_app_grpc_proto_requests_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_grpc_proto_requests_proto_rawDescData)
	})
	return file_internal_app_grpc_proto_requests_proto_rawDescData
}

var file_internal_app_grpc_proto_requests_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_internal_app_grpc_proto_requests_proto_goTypes = []interface{}{
	(*CompetitionRequest)(nil),       // 0: proto.CompetitionRequest
	(*FixtureRequest)(nil),           // 1: proto.FixtureRequest
	(*HistoricalResultRequest)(nil),  // 2: proto.HistoricalResultRequest
	(*SeasonFixtureRequest)(nil),     // 3: proto.SeasonFixtureRequest
	(*SeasonCompetitionRequest)(nil), // 4: proto.SeasonCompetitionRequest
	(*SeasonRequest)(nil),            // 5: proto.SeasonRequest
	(*SeasonTeamsRequest)(nil),       // 6: proto.SeasonTeamsRequest
	(*TeamRequest)(nil),              // 7: proto.TeamRequest
	(*TeamResultRequest)(nil),        // 8: proto.TeamResultRequest
	(*TeamSeasonsRequest)(nil),       // 9: proto.TeamSeasonsRequest
	(*wrappers.StringValue)(nil),     // 10: google.protobuf.StringValue
	(*wrappers.BoolValue)(nil),       // 11: google.protobuf.BoolValue
	(*wrappers.UInt64Value)(nil),     // 12: google.protobuf.UInt64Value
}
var file_internal_app_grpc_proto_requests_proto_depIdxs = []int32{
	10, // 0: proto.CompetitionRequest.sort:type_name -> google.protobuf.StringValue
	11, // 1: proto.CompetitionRequest.is_cup:type_name -> google.protobuf.BoolValue
	10, // 2: proto.SeasonCompetitionRequest.sort:type_name -> google.protobuf.StringValue
	12, // 3: proto.TeamResultRequest.limit:type_name -> google.protobuf.UInt64Value
	10, // 4: proto.TeamResultRequest.date_before:type_name -> google.protobuf.StringValue
	10, // 5: proto.TeamResultRequest.date_after:type_name -> google.protobuf.StringValue
	10, // 6: proto.TeamResultRequest.venue:type_name -> google.protobuf.StringValue
	10, // 7: proto.TeamResultRequest.sort:type_name -> google.protobuf.StringValue
	10, // 8: proto.TeamSeasonsRequest.sort:type_name -> google.protobuf.StringValue
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_internal_app_grpc_proto_requests_proto_init() }
func file_internal_app_grpc_proto_requests_proto_init() {
	if File_internal_app_grpc_proto_requests_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_app_grpc_proto_requests_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompetitionRequest); i {
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
		file_internal_app_grpc_proto_requests_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FixtureRequest); i {
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
		file_internal_app_grpc_proto_requests_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoricalResultRequest); i {
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
		file_internal_app_grpc_proto_requests_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeasonFixtureRequest); i {
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
		file_internal_app_grpc_proto_requests_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeasonCompetitionRequest); i {
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
		file_internal_app_grpc_proto_requests_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeasonRequest); i {
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
		file_internal_app_grpc_proto_requests_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeasonTeamsRequest); i {
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
		file_internal_app_grpc_proto_requests_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamRequest); i {
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
		file_internal_app_grpc_proto_requests_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamResultRequest); i {
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
		file_internal_app_grpc_proto_requests_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamSeasonsRequest); i {
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
			RawDescriptor: file_internal_app_grpc_proto_requests_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_app_grpc_proto_requests_proto_goTypes,
		DependencyIndexes: file_internal_app_grpc_proto_requests_proto_depIdxs,
		MessageInfos:      file_internal_app_grpc_proto_requests_proto_msgTypes,
	}.Build()
	File_internal_app_grpc_proto_requests_proto = out.File
	file_internal_app_grpc_proto_requests_proto_rawDesc = nil
	file_internal_app_grpc_proto_requests_proto_goTypes = nil
	file_internal_app_grpc_proto_requests_proto_depIdxs = nil
}
