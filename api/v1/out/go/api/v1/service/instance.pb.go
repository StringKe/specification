// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.0
// source: instance.proto

package service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Geolocation information for nearby routing
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Zone   *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	Campus *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=campus,proto3" json:"campus,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_instance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_instance_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetRegion() *wrapperspb.StringValue {
	if x != nil {
		return x.Region
	}
	return nil
}

func (x *Location) GetZone() *wrapperspb.StringValue {
	if x != nil {
		return x.Zone
	}
	return nil
}

func (x *Location) GetCampus() *wrapperspb.StringValue {
	if x != nil {
		return x.Campus
	}
	return nil
}

// Instance, a node that can provide service interface network calls, uniquely
// identified by IP:PORT
type Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// namespace name
	Namespace *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// service name
	Service *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// IPv4 or IPv6 address on which the instance listens
	Host *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	// instance corresponds to the listening port of the calling protocol
	Port *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	// calling protocol exposed by the instance
	Protocol *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Instance version, which can be used in routing rules
	Version *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	// Instance weight information for load balancing
	Weight *wrapperspb.UInt32Value `protobuf:"bytes,7,opt,name=weight,proto3" json:"weight,omitempty"`
	// The health status of the instance, through different health check methods
	// to determine whether the instance can provide external services
	Healthy *wrapperspb.BoolValue `protobuf:"bytes,8,opt,name=healthy,proto3" json:"healthy,omitempty"`
	// When this value is true, the instance is in an isolated state and should
	// not be returned to the caller
	Isolate *wrapperspb.BoolValue `protobuf:"bytes,9,opt,name=isolate,proto3" json:"isolate,omitempty"`
	// The location information of the instance, which is used to process the
	// nearest route
	Location *Location `protobuf:"bytes,10,opt,name=location,proto3" json:"location,omitempty"`
	// Instance metadata, which is used to store additional attribute data of the
	// instance, such as instance grouping information, instance startup time
	Metadata map[string]string `protobuf:"bytes,11,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Instance) Reset() {
	*x = Instance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instance) ProtoMessage() {}

func (x *Instance) ProtoReflect() protoreflect.Message {
	mi := &file_instance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instance.ProtoReflect.Descriptor instead.
func (*Instance) Descriptor() ([]byte, []int) {
	return file_instance_proto_rawDescGZIP(), []int{1}
}

func (x *Instance) GetNamespace() *wrapperspb.StringValue {
	if x != nil {
		return x.Namespace
	}
	return nil
}

func (x *Instance) GetService() *wrapperspb.StringValue {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *Instance) GetHost() *wrapperspb.StringValue {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *Instance) GetPort() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *Instance) GetProtocol() *wrapperspb.StringValue {
	if x != nil {
		return x.Protocol
	}
	return nil
}

func (x *Instance) GetVersion() *wrapperspb.StringValue {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *Instance) GetWeight() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Weight
	}
	return nil
}

func (x *Instance) GetHealthy() *wrapperspb.BoolValue {
	if x != nil {
		return x.Healthy
	}
	return nil
}

func (x *Instance) GetIsolate() *wrapperspb.BoolValue {
	if x != nil {
		return x.Isolate
	}
	return nil
}

func (x *Instance) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Instance) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_instance_proto protoreflect.FileDescriptor

var file_instance_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x08, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x04,
	0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x34,
	0x0a, 0x06, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x63, 0x61,
	0x6d, 0x70, 0x75, 0x73, 0x22, 0x9f, 0x05, 0x0a, 0x08, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x36, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x12, 0x36, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x34, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x12, 0x34, 0x0a, 0x07, 0x69, 0x73, 0x6f, 0x6c, 0x61,
	0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_instance_proto_rawDescOnce sync.Once
	file_instance_proto_rawDescData = file_instance_proto_rawDesc
)

func file_instance_proto_rawDescGZIP() []byte {
	file_instance_proto_rawDescOnce.Do(func() {
		file_instance_proto_rawDescData = protoimpl.X.CompressGZIP(file_instance_proto_rawDescData)
	})
	return file_instance_proto_rawDescData
}

var file_instance_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_instance_proto_goTypes = []interface{}{
	(*Location)(nil),               // 0: service.Location
	(*Instance)(nil),               // 1: service.Instance
	nil,                            // 2: service.Instance.MetadataEntry
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
	(*wrapperspb.UInt32Value)(nil), // 4: google.protobuf.UInt32Value
	(*wrapperspb.BoolValue)(nil),   // 5: google.protobuf.BoolValue
}
var file_instance_proto_depIdxs = []int32{
	3,  // 0: service.Location.region:type_name -> google.protobuf.StringValue
	3,  // 1: service.Location.zone:type_name -> google.protobuf.StringValue
	3,  // 2: service.Location.campus:type_name -> google.protobuf.StringValue
	3,  // 3: service.Instance.namespace:type_name -> google.protobuf.StringValue
	3,  // 4: service.Instance.service:type_name -> google.protobuf.StringValue
	3,  // 5: service.Instance.host:type_name -> google.protobuf.StringValue
	4,  // 6: service.Instance.port:type_name -> google.protobuf.UInt32Value
	3,  // 7: service.Instance.protocol:type_name -> google.protobuf.StringValue
	3,  // 8: service.Instance.version:type_name -> google.protobuf.StringValue
	4,  // 9: service.Instance.weight:type_name -> google.protobuf.UInt32Value
	5,  // 10: service.Instance.healthy:type_name -> google.protobuf.BoolValue
	5,  // 11: service.Instance.isolate:type_name -> google.protobuf.BoolValue
	0,  // 12: service.Instance.location:type_name -> service.Location
	2,  // 13: service.Instance.metadata:type_name -> service.Instance.MetadataEntry
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_instance_proto_init() }
func file_instance_proto_init() {
	if File_instance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_instance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_instance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instance); i {
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
			RawDescriptor: file_instance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_instance_proto_goTypes,
		DependencyIndexes: file_instance_proto_depIdxs,
		MessageInfos:      file_instance_proto_msgTypes,
	}.Build()
	File_instance_proto = out.File
	file_instance_proto_rawDesc = nil
	file_instance_proto_goTypes = nil
	file_instance_proto_depIdxs = nil
}