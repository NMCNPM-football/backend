// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: match.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MatchCalendar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClubOneName string `protobuf:"bytes,2,opt,name=club_one_name,json=clubOneName,proto3" json:"club_one_name,omitempty"`
	ClubTwoName string `protobuf:"bytes,3,opt,name=club_two_name,json=clubTwoName,proto3" json:"club_two_name,omitempty"`
	MatchDate   string `protobuf:"bytes,4,opt,name=match_date,json=matchDate,proto3" json:"match_date,omitempty"`
	MatchRound  string `protobuf:"bytes,5,opt,name=match_round,json=matchRound,proto3" json:"match_round,omitempty"`
	MatchTurn   string `protobuf:"bytes,6,opt,name=match_turn,json=matchTurn,proto3" json:"match_turn,omitempty"`
	MatchStatus string `protobuf:"bytes,7,opt,name=match_status,json=matchStatus,proto3" json:"match_status,omitempty"`
}

func (x *MatchCalendar) Reset() {
	*x = MatchCalendar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchCalendar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchCalendar) ProtoMessage() {}

func (x *MatchCalendar) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchCalendar.ProtoReflect.Descriptor instead.
func (*MatchCalendar) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{0}
}

func (x *MatchCalendar) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MatchCalendar) GetClubOneName() string {
	if x != nil {
		return x.ClubOneName
	}
	return ""
}

func (x *MatchCalendar) GetClubTwoName() string {
	if x != nil {
		return x.ClubTwoName
	}
	return ""
}

func (x *MatchCalendar) GetMatchDate() string {
	if x != nil {
		return x.MatchDate
	}
	return ""
}

func (x *MatchCalendar) GetMatchRound() string {
	if x != nil {
		return x.MatchRound
	}
	return ""
}

func (x *MatchCalendar) GetMatchTurn() string {
	if x != nil {
		return x.MatchTurn
	}
	return ""
}

func (x *MatchCalendar) GetMatchStatus() string {
	if x != nil {
		return x.MatchStatus
	}
	return ""
}

type MatchCalendarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MatchCalendarRequest) Reset() {
	*x = MatchCalendarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchCalendarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchCalendarRequest) ProtoMessage() {}

func (x *MatchCalendarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchCalendarRequest.ProtoReflect.Descriptor instead.
func (*MatchCalendarRequest) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{1}
}

func (x *MatchCalendarRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{2}
}

func (x *StatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type MatchCalendarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    *MatchCalendarResponse_Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MatchCalendarResponse) Reset() {
	*x = MatchCalendarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchCalendarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchCalendarResponse) ProtoMessage() {}

func (x *MatchCalendarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchCalendarResponse.ProtoReflect.Descriptor instead.
func (*MatchCalendarResponse) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{3}
}

func (x *MatchCalendarResponse) GetData() *MatchCalendarResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MatchCalendarResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MatchCalendarListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []*MatchCalendarResponse_Data `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Message string                        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MatchCalendarListResponse) Reset() {
	*x = MatchCalendarListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchCalendarListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchCalendarListResponse) ProtoMessage() {}

func (x *MatchCalendarListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchCalendarListResponse.ProtoReflect.Descriptor instead.
func (*MatchCalendarListResponse) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{4}
}

func (x *MatchCalendarListResponse) GetData() []*MatchCalendarResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MatchCalendarListResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MatchCalendarResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClubOneName string `protobuf:"bytes,1,opt,name=club_one_name,json=clubOneName,proto3" json:"club_one_name,omitempty"`
	ClubTwoName string `protobuf:"bytes,2,opt,name=club_two_name,json=clubTwoName,proto3" json:"club_two_name,omitempty"`
	MatchDate   string `protobuf:"bytes,3,opt,name=match_date,json=matchDate,proto3" json:"match_date,omitempty"`
	MatchRound  string `protobuf:"bytes,4,opt,name=match_round,json=matchRound,proto3" json:"match_round,omitempty"`
	MatchTurn   string `protobuf:"bytes,5,opt,name=match_turn,json=matchTurn,proto3" json:"match_turn,omitempty"`
	MatchStatus string `protobuf:"bytes,6,opt,name=match_status,json=matchStatus,proto3" json:"match_status,omitempty"`
}

func (x *MatchCalendarResponse_Data) Reset() {
	*x = MatchCalendarResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchCalendarResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchCalendarResponse_Data) ProtoMessage() {}

func (x *MatchCalendarResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchCalendarResponse_Data.ProtoReflect.Descriptor instead.
func (*MatchCalendarResponse_Data) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{3, 0}
}

func (x *MatchCalendarResponse_Data) GetClubOneName() string {
	if x != nil {
		return x.ClubOneName
	}
	return ""
}

func (x *MatchCalendarResponse_Data) GetClubTwoName() string {
	if x != nil {
		return x.ClubTwoName
	}
	return ""
}

func (x *MatchCalendarResponse_Data) GetMatchDate() string {
	if x != nil {
		return x.MatchDate
	}
	return ""
}

func (x *MatchCalendarResponse_Data) GetMatchRound() string {
	if x != nil {
		return x.MatchRound
	}
	return ""
}

func (x *MatchCalendarResponse_Data) GetMatchTurn() string {
	if x != nil {
		return x.MatchTurn
	}
	return ""
}

func (x *MatchCalendarResponse_Data) GetMatchStatus() string {
	if x != nil {
		return x.MatchStatus
	}
	return ""
}

var File_match_proto protoreflect.FileDescriptor

var file_match_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x01, 0x0a, 0x0d, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0d,
	0x63, 0x6c, 0x75, 0x62, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x62, 0x4f, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6c, 0x75, 0x62, 0x5f, 0x74, 0x77, 0x6f, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x62, 0x54, 0x77, 0x6f,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x75,
	0x72, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54,
	0x75, 0x72, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x26, 0x0a, 0x14, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27,
	0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xb5, 0x02, 0x0a, 0x15, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0xd0, 0x01, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6c, 0x75, 0x62, 0x5f, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c,
	0x75, 0x62, 0x4f, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6c, 0x75,
	0x62, 0x5f, 0x74, 0x77, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6c, 0x75, 0x62, 0x54, 0x77, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x75, 0x72, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x66, 0x0a, 0x19, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x36, 0x42, 0x0a, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x4d, 0x43, 0x4e, 0x50, 0x4d, 0x2d, 0x66, 0x6f, 0x6f, 0x74, 0x62,
	0x61, 0x6c, 0x6c, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_match_proto_rawDescOnce sync.Once
	file_match_proto_rawDescData = file_match_proto_rawDesc
)

func file_match_proto_rawDescGZIP() []byte {
	file_match_proto_rawDescOnce.Do(func() {
		file_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_match_proto_rawDescData)
	})
	return file_match_proto_rawDescData
}

var file_match_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_match_proto_goTypes = []interface{}{
	(*MatchCalendar)(nil),              // 0: MatchCalendar
	(*MatchCalendarRequest)(nil),       // 1: MatchCalendarRequest
	(*StatusRequest)(nil),              // 2: StatusRequest
	(*MatchCalendarResponse)(nil),      // 3: MatchCalendarResponse
	(*MatchCalendarListResponse)(nil),  // 4: MatchCalendarListResponse
	(*MatchCalendarResponse_Data)(nil), // 5: MatchCalendarResponse.Data
}
var file_match_proto_depIdxs = []int32{
	5, // 0: MatchCalendarResponse.data:type_name -> MatchCalendarResponse.Data
	5, // 1: MatchCalendarListResponse.data:type_name -> MatchCalendarResponse.Data
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_match_proto_init() }
func file_match_proto_init() {
	if File_match_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchCalendar); i {
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
		file_match_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchCalendarRequest); i {
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
		file_match_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_match_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchCalendarResponse); i {
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
		file_match_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchCalendarListResponse); i {
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
		file_match_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchCalendarResponse_Data); i {
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
			RawDescriptor: file_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_match_proto_goTypes,
		DependencyIndexes: file_match_proto_depIdxs,
		MessageInfos:      file_match_proto_msgTypes,
	}.Build()
	File_match_proto = out.File
	file_match_proto_rawDesc = nil
	file_match_proto_goTypes = nil
	file_match_proto_depIdxs = nil
}
