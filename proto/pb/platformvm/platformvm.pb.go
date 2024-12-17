// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: platformvm/platformvm.proto

package platformvm

import (
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

type L1ValidatorRegistrationJustification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Preimage:
	//
	//	*L1ValidatorRegistrationJustification_ConvertSubnetToL1TxData
	//	*L1ValidatorRegistrationJustification_RegisterL1ValidatorMessage
	Preimage isL1ValidatorRegistrationJustification_Preimage `protobuf_oneof:"preimage"`
}

func (x *L1ValidatorRegistrationJustification) Reset() {
	*x = L1ValidatorRegistrationJustification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platformvm_platformvm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *L1ValidatorRegistrationJustification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*L1ValidatorRegistrationJustification) ProtoMessage() {}

func (x *L1ValidatorRegistrationJustification) ProtoReflect() protoreflect.Message {
	mi := &file_platformvm_platformvm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use L1ValidatorRegistrationJustification.ProtoReflect.Descriptor instead.
func (*L1ValidatorRegistrationJustification) Descriptor() ([]byte, []int) {
	return file_platformvm_platformvm_proto_rawDescGZIP(), []int{0}
}

func (m *L1ValidatorRegistrationJustification) GetPreimage() isL1ValidatorRegistrationJustification_Preimage {
	if m != nil {
		return m.Preimage
	}
	return nil
}

func (x *L1ValidatorRegistrationJustification) GetConvertSubnetToL1TxData() *SubnetIDIndex {
	if x, ok := x.GetPreimage().(*L1ValidatorRegistrationJustification_ConvertSubnetToL1TxData); ok {
		return x.ConvertSubnetToL1TxData
	}
	return nil
}

func (x *L1ValidatorRegistrationJustification) GetRegisterL1ValidatorMessage() []byte {
	if x, ok := x.GetPreimage().(*L1ValidatorRegistrationJustification_RegisterL1ValidatorMessage); ok {
		return x.RegisterL1ValidatorMessage
	}
	return nil
}

type isL1ValidatorRegistrationJustification_Preimage interface {
	isL1ValidatorRegistrationJustification_Preimage()
}

type L1ValidatorRegistrationJustification_ConvertSubnetToL1TxData struct {
	// This should be set to obtain an attestation that a validator specified in
	// a ConvertSubnetToL1Tx has been removed from the validator set.
	ConvertSubnetToL1TxData *SubnetIDIndex `protobuf:"bytes,1,opt,name=convert_subnet_to_l1_tx_data,json=convertSubnetToL1TxData,proto3,oneof"`
}

type L1ValidatorRegistrationJustification_RegisterL1ValidatorMessage struct {
	// This should be set to a RegisterL1ValidatorMessage to obtain an
	// attestation that a validator is not currently registered and can never be
	// registered. This can be because the validator was successfully added and
	// then later removed, or because the validator was never added and the
	// registration expired.
	RegisterL1ValidatorMessage []byte `protobuf:"bytes,2,opt,name=register_l1_validator_message,json=registerL1ValidatorMessage,proto3,oneof"`
}

func (*L1ValidatorRegistrationJustification_ConvertSubnetToL1TxData) isL1ValidatorRegistrationJustification_Preimage() {
}

func (*L1ValidatorRegistrationJustification_RegisterL1ValidatorMessage) isL1ValidatorRegistrationJustification_Preimage() {
}

type SubnetIDIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubnetId []byte `protobuf:"bytes,1,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	Index    uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *SubnetIDIndex) Reset() {
	*x = SubnetIDIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platformvm_platformvm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubnetIDIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubnetIDIndex) ProtoMessage() {}

func (x *SubnetIDIndex) ProtoReflect() protoreflect.Message {
	mi := &file_platformvm_platformvm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubnetIDIndex.ProtoReflect.Descriptor instead.
func (*SubnetIDIndex) Descriptor() ([]byte, []int) {
	return file_platformvm_platformvm_proto_rawDescGZIP(), []int{1}
}

func (x *SubnetIDIndex) GetSubnetId() []byte {
	if x != nil {
		return x.SubnetId
	}
	return nil
}

func (x *SubnetIDIndex) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

var File_platformvm_platformvm_proto protoreflect.FileDescriptor

var file_platformvm_platformvm_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x76, 0x6d, 0x2f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x76, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x76, 0x6d, 0x22, 0xd3, 0x01, 0x0a, 0x24, 0x4c, 0x31,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x5a, 0x0a, 0x1c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x31, 0x5f, 0x74, 0x78, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x76, 0x6d, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x48, 0x00, 0x52, 0x17, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x53, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x54, 0x6f, 0x4c, 0x31, 0x54, 0x78, 0x44, 0x61, 0x74, 0x61, 0x12, 0x43,
	0x0a, 0x1d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x31, 0x5f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x1a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x4c, 0x31, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22,
	0x42, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x76, 0x61, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x61, 0x76, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x68, 0x65, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x76, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_platformvm_platformvm_proto_rawDescOnce sync.Once
	file_platformvm_platformvm_proto_rawDescData = file_platformvm_platformvm_proto_rawDesc
)

func file_platformvm_platformvm_proto_rawDescGZIP() []byte {
	file_platformvm_platformvm_proto_rawDescOnce.Do(func() {
		file_platformvm_platformvm_proto_rawDescData = protoimpl.X.CompressGZIP(file_platformvm_platformvm_proto_rawDescData)
	})
	return file_platformvm_platformvm_proto_rawDescData
}

var file_platformvm_platformvm_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_platformvm_platformvm_proto_goTypes = []interface{}{
	(*L1ValidatorRegistrationJustification)(nil), // 0: platformvm.L1ValidatorRegistrationJustification
	(*SubnetIDIndex)(nil),                        // 1: platformvm.SubnetIDIndex
}
var file_platformvm_platformvm_proto_depIdxs = []int32{
	1, // 0: platformvm.L1ValidatorRegistrationJustification.convert_subnet_to_l1_tx_data:type_name -> platformvm.SubnetIDIndex
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_platformvm_platformvm_proto_init() }
func file_platformvm_platformvm_proto_init() {
	if File_platformvm_platformvm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_platformvm_platformvm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*L1ValidatorRegistrationJustification); i {
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
		file_platformvm_platformvm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubnetIDIndex); i {
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
	file_platformvm_platformvm_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*L1ValidatorRegistrationJustification_ConvertSubnetToL1TxData)(nil),
		(*L1ValidatorRegistrationJustification_RegisterL1ValidatorMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_platformvm_platformvm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_platformvm_platformvm_proto_goTypes,
		DependencyIndexes: file_platformvm_platformvm_proto_depIdxs,
		MessageInfos:      file_platformvm_platformvm_proto_msgTypes,
	}.Build()
	File_platformvm_platformvm_proto = out.File
	file_platformvm_platformvm_proto_rawDesc = nil
	file_platformvm_platformvm_proto_goTypes = nil
	file_platformvm_platformvm_proto_depIdxs = nil
}