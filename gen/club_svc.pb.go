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
	0x1a, 0x0a, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xab, 0x01, 0x0a,
	0x0b, 0x43, 0x6c, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0d,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x43, 0x6c, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x63, 0x6c,
	0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x13, 0x2e, 0x43, 0x6c, 0x75, 0x62, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x43, 0x6c, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0d, 0x2f, 0x63, 0x6c,
	0x75, 0x62, 0x2f, 0x63, 0x6c, 0x75, 0x62, 0x5f, 0x69, 0x64, 0x42, 0x75, 0x0a, 0x09, 0x63, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0c, 0x43, 0x6c, 0x75, 0x62, 0x53, 0x76, 0x63,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x4d, 0x43, 0x4e, 0x50, 0x4d, 0x2d, 0x66, 0x6f, 0x6f, 0x74, 0x62,
	0x61, 0x6c, 0x6c, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0xa2,
	0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x05,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0xe2, 0x02, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_club_svc_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil),        // 0: EmptyRequest
	(*ClubProfileRequest)(nil),  // 1: ClubProfileRequest
	(*ClubProfileResponse)(nil), // 2: ClubProfileResponse
}
var file_club_svc_proto_depIdxs = []int32{
	0, // 0: proto.ClubService.GetClubProfile:input_type -> EmptyRequest
	1, // 1: proto.ClubService.UpdateClub:input_type -> ClubProfileRequest
	2, // 2: proto.ClubService.GetClubProfile:output_type -> ClubProfileResponse
	2, // 3: proto.ClubService.UpdateClub:output_type -> ClubProfileResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
