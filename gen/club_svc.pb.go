// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: club_svc.proto

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

var File_club_svc_proto protoreflect.FileDescriptor

var file_club_svc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x62, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd7, 0x07, 0x0a,
	0x0b, 0x43, 0x6c, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0d,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x43, 0x6c, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x63, 0x6c,
	0x75, 0x62, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x50, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x13, 0x2e, 0x43, 0x6c, 0x75, 0x62, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x43, 0x6c, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x1a, 0x0c,
	0x2f, 0x63, 0x6c, 0x75, 0x62, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x57, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6c, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x0d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x43, 0x6c, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x12, 0x11, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2f, 0x61, 0x6c, 0x6c, 0x12, 0x52, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x50, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a,
	0x22, 0x07, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x5d, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x50, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x50, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x3a, 0x01, 0x2a, 0x1a, 0x13, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x54, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x50, 0x4c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x55,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x50, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x50, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x12, 0x11, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12,
	0x10, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x61, 0x6c,
	0x6c, 0x12, 0x5b, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68,
	0x12, 0x14, 0x2e, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x1a, 0x12, 0x2f, 0x63, 0x6f, 0x61,
	0x63, 0x68, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x51,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x0d, 0x2e, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12,
	0x10, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x63, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x61, 0x63,
	0x68, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x43, 0x6f, 0x61, 0x63, 0x68,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x61,
	0x63, 0x68, 0x2f, 0x61, 0x6c, 0x6c, 0x42, 0x75, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x42, 0x0c, 0x43, 0x6c, 0x75, 0x62, 0x53, 0x76, 0x63, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x4d, 0x43, 0x4e, 0x50, 0x4d, 0x2d, 0x66, 0x6f, 0x6f, 0x74, 0x62, 0x61, 0x6c, 0x6c, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0xa2, 0x02, 0x03, 0x50, 0x58,
	0x58, 0xaa, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0xe2, 0x02, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_club_svc_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil),              // 0: EmptyRequest
	(*ClubProfileRequest)(nil),        // 1: ClubProfileRequest
	(*PLayerProfileRequest)(nil),      // 2: PLayerProfileRequest
	(*PLayerRequest)(nil),             // 3: PLayerRequest
	(*CoachProfileRequest)(nil),       // 4: CoachProfileRequest
	(*CoachRequest)(nil),              // 5: CoachRequest
	(*CoachProfileListRequest)(nil),   // 6: CoachProfileListRequest
	(*ClubProfileResponse)(nil),       // 7: ClubProfileResponse
	(*ClubProfileListResponse)(nil),   // 8: ClubProfileListResponse
	(*SuccessMessageResponse)(nil),    // 9: SuccessMessageResponse
	(*PLayerProfileResponse)(nil),     // 10: PLayerProfileResponse
	(*PlayerProfileListResponse)(nil), // 11: PlayerProfileListResponse
	(*CoachProfileResponse)(nil),      // 12: CoachProfileResponse
	(*CoachProfileListResponse)(nil),  // 13: CoachProfileListResponse
}
var file_club_svc_proto_depIdxs = []int32{
	0,  // 0: proto.ClubService.GetClubProfile:input_type -> EmptyRequest
	1,  // 1: proto.ClubService.UpdateClub:input_type -> ClubProfileRequest
	0,  // 2: proto.ClubService.GetAllClubProfile:input_type -> EmptyRequest
	2,  // 3: proto.ClubService.CreatePlayer:input_type -> PLayerProfileRequest
	2,  // 4: proto.ClubService.UpdatePlayer:input_type -> PLayerProfileRequest
	3,  // 5: proto.ClubService.DeletePlayer:input_type -> PLayerRequest
	3,  // 6: proto.ClubService.GetPlayerProfile:input_type -> PLayerRequest
	0,  // 7: proto.ClubService.GetAllPlayerProfile:input_type -> EmptyRequest
	4,  // 8: proto.ClubService.UpdateCoach:input_type -> CoachProfileRequest
	5,  // 9: proto.ClubService.GetCoachProfile:input_type -> CoachRequest
	6,  // 10: proto.ClubService.GetListCoachProfile:input_type -> CoachProfileListRequest
	7,  // 11: proto.ClubService.GetClubProfile:output_type -> ClubProfileResponse
	7,  // 12: proto.ClubService.UpdateClub:output_type -> ClubProfileResponse
	8,  // 13: proto.ClubService.GetAllClubProfile:output_type -> ClubProfileListResponse
	9,  // 14: proto.ClubService.CreatePlayer:output_type -> SuccessMessageResponse
	10, // 15: proto.ClubService.UpdatePlayer:output_type -> PLayerProfileResponse
	9,  // 16: proto.ClubService.DeletePlayer:output_type -> SuccessMessageResponse
	10, // 17: proto.ClubService.GetPlayerProfile:output_type -> PLayerProfileResponse
	11, // 18: proto.ClubService.GetAllPlayerProfile:output_type -> PlayerProfileListResponse
	9,  // 19: proto.ClubService.UpdateCoach:output_type -> SuccessMessageResponse
	12, // 20: proto.ClubService.GetCoachProfile:output_type -> CoachProfileResponse
	13, // 21: proto.ClubService.GetListCoachProfile:output_type -> CoachProfileListResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_club_svc_proto_init() }
func file_club_svc_proto_init() {
	if File_club_svc_proto != nil {
		return
	}
	file_message_proto_init()
	file_user_proto_init()
	file_club_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_club_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_club_svc_proto_goTypes,
		DependencyIndexes: file_club_svc_proto_depIdxs,
	}.Build()
	File_club_svc_proto = out.File
	file_club_svc_proto_rawDesc = nil
	file_club_svc_proto_goTypes = nil
	file_club_svc_proto_depIdxs = nil
}
