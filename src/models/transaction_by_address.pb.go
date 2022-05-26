// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: transaction_by_address.proto

package models

import (
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
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

type TransactionByAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionHash string `protobuf:"bytes,1,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash"`
	Address         string `protobuf:"bytes,2,opt,name=address,proto3" json:"address"`
	BlockNumber     int64  `protobuf:"varint,3,opt,name=block_number,json=blockNumber,proto3" json:"block_number"`
}

func (x *TransactionByAddress) Reset() {
	*x = TransactionByAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_by_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionByAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionByAddress) ProtoMessage() {}

func (x *TransactionByAddress) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_by_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionByAddress.ProtoReflect.Descriptor instead.
func (*TransactionByAddress) Descriptor() ([]byte, []int) {
	return file_transaction_by_address_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionByAddress) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *TransactionByAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TransactionByAddress) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

var File_transaction_by_address_proto protoreflect.FileDescriptor

var file_transaction_by_address_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x79,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x0a, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x10, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x28, 0x01, 0x52,
	0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x22, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x28, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x52, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x2f, 0xba, 0xb9, 0x19, 0x2b,
	0x0a, 0x29, 0x52, 0x27, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x62, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_by_address_proto_rawDescOnce sync.Once
	file_transaction_by_address_proto_rawDescData = file_transaction_by_address_proto_rawDesc
)

func file_transaction_by_address_proto_rawDescGZIP() []byte {
	file_transaction_by_address_proto_rawDescOnce.Do(func() {
		file_transaction_by_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_by_address_proto_rawDescData)
	})
	return file_transaction_by_address_proto_rawDescData
}

var file_transaction_by_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transaction_by_address_proto_goTypes = []interface{}{
	(*TransactionByAddress)(nil), // 0: models.TransactionByAddress
}
var file_transaction_by_address_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transaction_by_address_proto_init() }
func file_transaction_by_address_proto_init() {
	if File_transaction_by_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transaction_by_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionByAddress); i {
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
			RawDescriptor: file_transaction_by_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_by_address_proto_goTypes,
		DependencyIndexes: file_transaction_by_address_proto_depIdxs,
		MessageInfos:      file_transaction_by_address_proto_msgTypes,
	}.Build()
	File_transaction_by_address_proto = out.File
	file_transaction_by_address_proto_rawDesc = nil
	file_transaction_by_address_proto_goTypes = nil
	file_transaction_by_address_proto_depIdxs = nil
}
