// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: platform.proto

package platform

import (
	providers "go.mondoo.com/cnquery/motor/providers"
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

type Platform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// FIXME: remove in v8.0 vv
	//
	// Deprecated: Do not use.
	Release string            `protobuf:"bytes,2,opt,name=release,proto3" json:"release,omitempty"`
	Arch    string            `protobuf:"bytes,3,opt,name=arch,proto3" json:"arch,omitempty"`
	Title   string            `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Family  []string          `protobuf:"bytes,5,rep,name=family,proto3" json:"family,omitempty"`
	Build   string            `protobuf:"bytes,6,opt,name=build,proto3" json:"build,omitempty"`
	Version string            `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	Kind    providers.Kind    `protobuf:"varint,20,opt,name=kind,proto3,enum=cnquery.motor.providers.v1.Kind" json:"kind,omitempty"`
	Runtime string            `protobuf:"bytes,21,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Labels  map[string]string `protobuf:"bytes,22,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Platform) Reset() {
	*x = Platform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform) ProtoMessage() {}

func (x *Platform) ProtoReflect() protoreflect.Message {
	mi := &file_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform.ProtoReflect.Descriptor instead.
func (*Platform) Descriptor() ([]byte, []int) {
	return file_platform_proto_rawDescGZIP(), []int{0}
}

func (x *Platform) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Deprecated: Do not use.
func (x *Platform) GetRelease() string {
	if x != nil {
		return x.Release
	}
	return ""
}

func (x *Platform) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *Platform) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Platform) GetFamily() []string {
	if x != nil {
		return x.Family
	}
	return nil
}

func (x *Platform) GetBuild() string {
	if x != nil {
		return x.Build
	}
	return ""
}

func (x *Platform) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Platform) GetKind() providers.Kind {
	if x != nil {
		return x.Kind
	}
	return providers.Kind(0)
}

func (x *Platform) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *Platform) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_platform_proto protoreflect.FileDescriptor

var file_platform_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x63, 0x6e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x6d, 0x6f, 0x74,
	0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03, 0x0a, 0x08,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x07,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72,
	0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63,
	0x6e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x6f, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_platform_proto_rawDescOnce sync.Once
	file_platform_proto_rawDescData = file_platform_proto_rawDesc
)

func file_platform_proto_rawDescGZIP() []byte {
	file_platform_proto_rawDescOnce.Do(func() {
		file_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_proto_rawDescData)
	})
	return file_platform_proto_rawDescData
}

var file_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_platform_proto_goTypes = []interface{}{
	(*Platform)(nil),    // 0: cnquery.motor.platform.v1.Platform
	nil,                 // 1: cnquery.motor.platform.v1.Platform.LabelsEntry
	(providers.Kind)(0), // 2: cnquery.motor.providers.v1.Kind
}
var file_platform_proto_depIdxs = []int32{
	2, // 0: cnquery.motor.platform.v1.Platform.kind:type_name -> cnquery.motor.providers.v1.Kind
	1, // 1: cnquery.motor.platform.v1.Platform.labels:type_name -> cnquery.motor.platform.v1.Platform.LabelsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_platform_proto_init() }
func file_platform_proto_init() {
	if File_platform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_platform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platform); i {
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
			RawDescriptor: file_platform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_platform_proto_goTypes,
		DependencyIndexes: file_platform_proto_depIdxs,
		MessageInfos:      file_platform_proto_msgTypes,
	}.Build()
	File_platform_proto = out.File
	file_platform_proto_rawDesc = nil
	file_platform_proto_goTypes = nil
	file_platform_proto_depIdxs = nil
}
