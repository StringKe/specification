// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.0
// source: circuitbreaker.proto

package faulttolerance

import (
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

var File_circuitbreaker_proto protoreflect.FileDescriptor

var file_circuitbreaker_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x74, 0x6f, 0x6c,
	0x65, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x74, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_circuitbreaker_proto_goTypes = []interface{}{}
var file_circuitbreaker_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_circuitbreaker_proto_init() }
func file_circuitbreaker_proto_init() {
	if File_circuitbreaker_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_circuitbreaker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_circuitbreaker_proto_goTypes,
		DependencyIndexes: file_circuitbreaker_proto_depIdxs,
	}.Build()
	File_circuitbreaker_proto = out.File
	file_circuitbreaker_proto_rawDesc = nil
	file_circuitbreaker_proto_goTypes = nil
	file_circuitbreaker_proto_depIdxs = nil
}