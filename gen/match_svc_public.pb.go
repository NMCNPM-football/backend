// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: match_svc_public.proto

package gen

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_match_svc_public_proto protoreflect.FileDescriptor

var file_match_svc_public_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x76, 0x63, 0x5f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb1, 0x03, 0x0a, 0x12, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x12, 0x6e, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x73, 0x57, 0x69, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x7b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x7d,
	0x12, 0x5d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x0d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x61, 0x6c, 0x6c, 0x12,
	0x63, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14,
	0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x67, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x13, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x7b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x7c, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x13, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x76, 0x63, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x4d,
	0x43, 0x4e, 0x50, 0x4d, 0x2d, 0x66, 0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x2f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa,
	0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xe2,
	0x02, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_match_svc_public_proto_goTypes = []interface{}{
	(*StatusRequest)(nil),             // 0: StatusRequest
	(*EmptyRequest)(nil),              // 1: EmptyRequest
	(*MatchCalendarRequest)(nil),      // 2: MatchCalendarRequest
	(*ResultScoreRequest)(nil),        // 3: ResultScoreRequest
	(*MatchCalendarListResponse)(nil), // 4: MatchCalendarListResponse
	(*MatchCalendarResponse)(nil),     // 5: MatchCalendarResponse
	(*ResultScoreResponse)(nil),       // 6: ResultScoreResponse
}
var file_match_svc_public_proto_depIdxs = []int32{
	0, // 0: proto.MatchServicePublic.GetAllMatchCalendarsWithStatus:input_type -> StatusRequest
	1, // 1: proto.MatchServicePublic.GetAllMatchCalendar:input_type -> EmptyRequest
	2, // 2: proto.MatchServicePublic.GetMatchCalendarById:input_type -> MatchCalendarRequest
	3, // 3: proto.MatchServicePublic.GetMatchResultByID:input_type -> ResultScoreRequest
	4, // 4: proto.MatchServicePublic.GetAllMatchCalendarsWithStatus:output_type -> MatchCalendarListResponse
	4, // 5: proto.MatchServicePublic.GetAllMatchCalendar:output_type -> MatchCalendarListResponse
	5, // 6: proto.MatchServicePublic.GetMatchCalendarById:output_type -> MatchCalendarResponse
	6, // 7: proto.MatchServicePublic.GetMatchResultByID:output_type -> ResultScoreResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_match_svc_public_proto_init() }
func file_match_svc_public_proto_init() {
	if File_match_svc_public_proto != nil {
		return
	}
	file_message_proto_init()
	file_match_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_match_svc_public_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_match_svc_public_proto_goTypes,
		DependencyIndexes: file_match_svc_public_proto_depIdxs,
	}.Build()
	File_match_svc_public_proto = out.File
	file_match_svc_public_proto_rawDesc = nil
	file_match_svc_public_proto_goTypes = nil
	file_match_svc_public_proto_depIdxs = nil
}
