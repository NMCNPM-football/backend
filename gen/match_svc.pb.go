// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: match_svc.proto

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

var File_match_svc_proto protoreflect.FileDescriptor

var file_match_svc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xfe, 0x07, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x23, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x7b, 0x73, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x7d, 0x12, 0x6c, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x1a,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x6c,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x1a, 0x16, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x12, 0x66, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x2e,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01,
	0x2a, 0x22, 0x15, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x63, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x61, 0x72, 0x64, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43,
	0x61, 0x72, 0x64, 0x1a, 0x17, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x12, 0x6b, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x7b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x65, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x0d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x12, 0x5d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x6f, 0x61, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x1a, 0x17, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x5d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x1a, 0x17, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x58, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x6f, 0x61, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x12, 0x16, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f,
	0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x58, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x76, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x42, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x76, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x4d,
	0x43, 0x4e, 0x50, 0x4d, 0x2d, 0x66, 0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x2f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa,
	0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xe2,
	0x02, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_match_svc_proto_goTypes = []interface{}{
	(*MatchCalendar)(nil),          // 0: proto.MatchCalendar
	(*ProgressScore)(nil),          // 1: proto.ProgressScore
	(*ProgressCard)(nil),           // 2: proto.ProgressCard
	(*ResultScore)(nil),            // 3: proto.ResultScore
	(*EmptyRequest)(nil),           // 4: EmptyRequest
	(*GoalType)(nil),               // 5: proto.GoalType
	(*CardType)(nil),               // 6: proto.CardType
	(*MatchCalendarRequest)(nil),   // 7: proto.MatchCalendarRequest
	(*MatchCalendarResponse)(nil),  // 8: proto.MatchCalendarResponse
	(*SuccessMessageResponse)(nil), // 9: SuccessMessageResponse
	(*GoalTypeResponse)(nil),       // 10: proto.GoalTypeResponse
	(*CardTypeResponse)(nil),       // 11: proto.CardTypeResponse
}
var file_match_svc_proto_depIdxs = []int32{
	0,  // 0: proto.MatchService.CreateMatchCalendar:input_type -> proto.MatchCalendar
	0,  // 1: proto.MatchService.UpdateMatchCalendar:input_type -> proto.MatchCalendar
	1,  // 2: proto.MatchService.CreateProgressScore:input_type -> proto.ProgressScore
	2,  // 3: proto.MatchService.CreateProgressCard:input_type -> proto.ProgressCard
	3,  // 4: proto.MatchService.CreateMatchResult:input_type -> proto.ResultScore
	4,  // 5: proto.MatchService.CreateAllMatchResults:input_type -> EmptyRequest
	5,  // 6: proto.MatchService.CreateGoalType:input_type -> proto.GoalType
	6,  // 7: proto.MatchService.CreateCardType:input_type -> proto.CardType
	4,  // 8: proto.MatchService.GetAllGoalType:input_type -> EmptyRequest
	4,  // 9: proto.MatchService.GetAllCardType:input_type -> EmptyRequest
	7,  // 10: proto.MatchService.CreateMatchCalendar:output_type -> proto.MatchCalendarRequest
	8,  // 11: proto.MatchService.UpdateMatchCalendar:output_type -> proto.MatchCalendarResponse
	9,  // 12: proto.MatchService.CreateProgressScore:output_type -> SuccessMessageResponse
	9,  // 13: proto.MatchService.CreateProgressCard:output_type -> SuccessMessageResponse
	9,  // 14: proto.MatchService.CreateMatchResult:output_type -> SuccessMessageResponse
	9,  // 15: proto.MatchService.CreateAllMatchResults:output_type -> SuccessMessageResponse
	9,  // 16: proto.MatchService.CreateGoalType:output_type -> SuccessMessageResponse
	9,  // 17: proto.MatchService.CreateCardType:output_type -> SuccessMessageResponse
	10, // 18: proto.MatchService.GetAllGoalType:output_type -> proto.GoalTypeResponse
	11, // 19: proto.MatchService.GetAllCardType:output_type -> proto.CardTypeResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_match_svc_proto_init() }
func file_match_svc_proto_init() {
	if File_match_svc_proto != nil {
		return
	}
	file_message_proto_init()
	file_match_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_match_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_match_svc_proto_goTypes,
		DependencyIndexes: file_match_svc_proto_depIdxs,
	}.Build()
	File_match_svc_proto = out.File
	file_match_svc_proto_rawDesc = nil
	file_match_svc_proto_goTypes = nil
	file_match_svc_proto_depIdxs = nil
}
