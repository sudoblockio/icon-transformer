// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.21.1
// source: token_transfer.proto

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

type TokenTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionHash      string  `protobuf:"bytes,1,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash"`
	LogIndex             int64   `protobuf:"varint,2,opt,name=log_index,json=logIndex,proto3" json:"log_index"`
	TokenContractAddress string  `protobuf:"bytes,3,opt,name=token_contract_address,json=tokenContractAddress,proto3" json:"token_contract_address"`
	FromAddress          string  `protobuf:"bytes,4,opt,name=from_address,json=fromAddress,proto3" json:"from_address"`
	ToAddress            string  `protobuf:"bytes,5,opt,name=to_address,json=toAddress,proto3" json:"to_address"`
	BlockNumber          int64   `protobuf:"varint,6,opt,name=block_number,json=blockNumber,proto3" json:"block_number"`
	Value                string  `protobuf:"bytes,7,opt,name=value,proto3" json:"value"`
	ValueDecimal         float64 `protobuf:"fixed64,8,opt,name=value_decimal,json=valueDecimal,proto3" json:"value_decimal"`
	BlockTimestamp       int64   `protobuf:"varint,9,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp"`
	TokenContractName    string  `protobuf:"bytes,10,opt,name=token_contract_name,json=tokenContractName,proto3" json:"token_contract_name"`
	TokenContractSymbol  string  `protobuf:"bytes,11,opt,name=token_contract_symbol,json=tokenContractSymbol,proto3" json:"token_contract_symbol"`
	TransactionFee       string  `protobuf:"bytes,12,opt,name=transaction_fee,json=transactionFee,proto3" json:"transaction_fee"`
}

func (x *TokenTransfer) Reset() {
	*x = TokenTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenTransfer) ProtoMessage() {}

func (x *TokenTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_token_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenTransfer.ProtoReflect.Descriptor instead.
func (*TokenTransfer) Descriptor() ([]byte, []int) {
	return file_token_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *TokenTransfer) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *TokenTransfer) GetLogIndex() int64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *TokenTransfer) GetTokenContractAddress() string {
	if x != nil {
		return x.TokenContractAddress
	}
	return ""
}

func (x *TokenTransfer) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *TokenTransfer) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *TokenTransfer) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *TokenTransfer) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TokenTransfer) GetValueDecimal() float64 {
	if x != nil {
		return x.ValueDecimal
	}
	return 0
}

func (x *TokenTransfer) GetBlockTimestamp() int64 {
	if x != nil {
		return x.BlockTimestamp
	}
	return 0
}

func (x *TokenTransfer) GetTokenContractName() string {
	if x != nil {
		return x.TokenContractName
	}
	return ""
}

func (x *TokenTransfer) GetTokenContractSymbol() string {
	if x != nil {
		return x.TokenContractSymbol
	}
	return ""
}

func (x *TokenTransfer) GetTransactionFee() string {
	if x != nil {
		return x.TransactionFee
	}
	return ""
}

var File_token_transfer_proto protoreflect.FileDescriptor

var file_token_transfer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62,
	0x6c, 0x6f, 0x78, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x05, 0x0a, 0x0d, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x10,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x28, 0x01,
	0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x25, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x28, 0x01, 0x52, 0x08,
	0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x67, 0x0a, 0x16, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xba, 0xb9, 0x19, 0x2d, 0x0a, 0x2b,
	0x52, 0x29, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x14, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x4a, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xba, 0xb9, 0x19, 0x23, 0x0a, 0x21, 0x52,
	0x1f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x78, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x44, 0x0a,
	0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x25, 0xba, 0xb9, 0x19, 0x21, 0x0a, 0x1f, 0x52, 0x1d, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x74, 0x6f,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x4a, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x27, 0xba, 0xb9, 0x19, 0x23, 0x0a,
	0x21, 0x52, 0x1f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x64,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65,
	0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_transfer_proto_rawDescOnce sync.Once
	file_token_transfer_proto_rawDescData = file_token_transfer_proto_rawDesc
)

func file_token_transfer_proto_rawDescGZIP() []byte {
	file_token_transfer_proto_rawDescOnce.Do(func() {
		file_token_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_transfer_proto_rawDescData)
	})
	return file_token_transfer_proto_rawDescData
}

var file_token_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_token_transfer_proto_goTypes = []interface{}{
	(*TokenTransfer)(nil), // 0: models.TokenTransfer
}
var file_token_transfer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_token_transfer_proto_init() }
func file_token_transfer_proto_init() {
	if File_token_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_token_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenTransfer); i {
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
			RawDescriptor: file_token_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_token_transfer_proto_goTypes,
		DependencyIndexes: file_token_transfer_proto_depIdxs,
		MessageInfos:      file_token_transfer_proto_msgTypes,
	}.Build()
	File_token_transfer_proto = out.File
	file_token_transfer_proto_rawDesc = nil
	file_token_transfer_proto_goTypes = nil
	file_token_transfer_proto_depIdxs = nil
}
