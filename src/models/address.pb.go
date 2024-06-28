// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: address.proto

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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// BlockETL
	Address                  string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address"`
	IsContract               bool    `protobuf:"varint,2,opt,name=is_contract,json=isContract,proto3" json:"is_contract"`
	TransactionCount         int64   `protobuf:"varint,3,opt,name=transaction_count,json=transactionCount,proto3" json:"transaction_count"`
	TransactionInternalCount int64   `protobuf:"varint,4,opt,name=transaction_internal_count,json=transactionInternalCount,proto3" json:"transaction_internal_count"`
	LogCount                 int64   `protobuf:"varint,5,opt,name=log_count,json=logCount,proto3" json:"log_count"`
	TokenTransferCount       int64   `protobuf:"varint,6,opt,name=token_transfer_count,json=tokenTransferCount,proto3" json:"token_transfer_count"`
	Balance                  float64 `protobuf:"fixed64,7,opt,name=balance,proto3" json:"balance"`
	Type                     string  `protobuf:"bytes,8,opt,name=type,proto3" json:"type"`
	// Contracts Processed
	Name                 string `protobuf:"bytes,9,opt,name=name,proto3" json:"name"`
	CreatedTimestamp     int64  `protobuf:"varint,11,opt,name=created_timestamp,json=createdTimestamp,proto3" json:"created_timestamp"`
	IsToken              bool   `protobuf:"varint,12,opt,name=is_token,json=isToken,proto3" json:"is_token"`
	IsNft                bool   `protobuf:"varint,19,opt,name=is_nft,json=isNft,proto3" json:"is_nft"`
	ContractUpdatedBlock int64  `protobuf:"varint,14,opt,name=contract_updated_block,json=contractUpdatedBlock,proto3" json:"contract_updated_block"`
	AuditTxHash          string `protobuf:"bytes,20,opt,name=audit_tx_hash,json=auditTxHash,proto3" json:"audit_tx_hash"`
	CodeHash             string `protobuf:"bytes,21,opt,name=code_hash,json=codeHash,proto3" json:"code_hash"`
	DeployTxHash         string `protobuf:"bytes,22,opt,name=deploy_tx_hash,json=deployTxHash,proto3" json:"deploy_tx_hash"`
	ContractType         string `protobuf:"bytes,15,opt,name=contract_type,json=contractType,proto3" json:"contract_type"`
	Status               string `protobuf:"bytes,10,opt,name=status,proto3" json:"status"`
	Owner                string `protobuf:"bytes,23,opt,name=owner,proto3" json:"owner"`
	// Governance Prep Processed
	IsPrep bool `protobuf:"varint,13,opt,name=is_prep,json=isPrep,proto3" json:"is_prep"`
	// Tokens
	TokenStandard string `protobuf:"bytes,16,opt,name=token_standard,json=tokenStandard,proto3" json:"token_standard"`
	Symbol        string `protobuf:"bytes,17,opt,name=symbol,proto3" json:"symbol"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Address) GetIsContract() bool {
	if x != nil {
		return x.IsContract
	}
	return false
}

func (x *Address) GetTransactionCount() int64 {
	if x != nil {
		return x.TransactionCount
	}
	return 0
}

func (x *Address) GetTransactionInternalCount() int64 {
	if x != nil {
		return x.TransactionInternalCount
	}
	return 0
}

func (x *Address) GetLogCount() int64 {
	if x != nil {
		return x.LogCount
	}
	return 0
}

func (x *Address) GetTokenTransferCount() int64 {
	if x != nil {
		return x.TokenTransferCount
	}
	return 0
}

func (x *Address) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Address) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Address) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Address) GetCreatedTimestamp() int64 {
	if x != nil {
		return x.CreatedTimestamp
	}
	return 0
}

func (x *Address) GetIsToken() bool {
	if x != nil {
		return x.IsToken
	}
	return false
}

func (x *Address) GetIsNft() bool {
	if x != nil {
		return x.IsNft
	}
	return false
}

func (x *Address) GetContractUpdatedBlock() int64 {
	if x != nil {
		return x.ContractUpdatedBlock
	}
	return 0
}

func (x *Address) GetAuditTxHash() string {
	if x != nil {
		return x.AuditTxHash
	}
	return ""
}

func (x *Address) GetCodeHash() string {
	if x != nil {
		return x.CodeHash
	}
	return ""
}

func (x *Address) GetDeployTxHash() string {
	if x != nil {
		return x.DeployTxHash
	}
	return ""
}

func (x *Address) GetContractType() string {
	if x != nil {
		return x.ContractType
	}
	return ""
}

func (x *Address) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Address) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Address) GetIsPrep() bool {
	if x != nil {
		return x.IsPrep
	}
	return false
}

func (x *Address) GetTokenStandard() string {
	if x != nil {
		return x.TokenStandard
	}
	return ""
}

func (x *Address) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

var File_address_proto protoreflect.FileDescriptor

var file_address_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x09, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x28, 0x01,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x69, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x1f,
	0xba, 0xb9, 0x19, 0x1b, 0x0a, 0x19, 0x52, 0x17, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f,
	0x69, 0x64, 0x78, 0x5f, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52,
	0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x52, 0x0a, 0x11, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x25, 0xba, 0xb9, 0x19, 0x21, 0x0a, 0x1f, 0x52, 0x1d,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x10, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x6c, 0x0a, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x2e, 0xba, 0xb9, 0x19, 0x2a, 0x0a, 0x28, 0x52, 0x26, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x18, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a,
	0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x1d, 0xba, 0xb9, 0x19, 0x19, 0x0a, 0x17, 0x52, 0x15, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x08, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x5a, 0x0a, 0x14, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x28, 0xba, 0xb9, 0x19, 0x24, 0x0a, 0x22, 0x52,
	0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x42, 0x1b, 0xba, 0xb9, 0x19, 0x17, 0x0a, 0x15, 0x52, 0x13,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x35, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21,
	0xba, 0xb9, 0x19, 0x1d, 0x0a, 0x1b, 0x52, 0x19, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f,
	0x69, 0x64, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x25, 0xba, 0xb9, 0x19, 0x21, 0x0a, 0x1f, 0x52, 0x1d, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x37, 0x0a, 0x08, 0x69,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x42, 0x1c, 0xba,
	0xb9, 0x19, 0x18, 0x0a, 0x16, 0x52, 0x14, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69,
	0x64, 0x78, 0x5f, 0x69, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x07, 0x69, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x6e, 0x66, 0x74, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x1a, 0xba, 0xb9, 0x19, 0x16, 0x0a, 0x14, 0x52, 0x12, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x69, 0x73, 0x5f, 0x6e, 0x66, 0x74,
	0x52, 0x05, 0x69, 0x73, 0x4e, 0x66, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x22, 0x0a,
	0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x54, 0x78, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x24,
	0x0a, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x54, 0x78,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x46, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xba, 0xb9, 0x19,
	0x1d, 0x0a, 0x1b, 0x52, 0x19, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x78,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xba, 0xb9,
	0x19, 0x16, 0x0a, 0x14, 0x52, 0x12, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64,
	0x78, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65,
	0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x42, 0x26, 0xba, 0xb9, 0x19, 0x22, 0x0a, 0x20, 0x52,
	0x1e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x69, 0x73, 0x5f,
	0x67, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x70, 0x52,
	0x06, 0x69, 0x73, 0x50, 0x72, 0x65, 0x70, 0x12, 0x49, 0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x22, 0xba, 0xb9, 0x19, 0x1e, 0x0a, 0x1c, 0x52, 0x1a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x69, 0x64, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02,
	0x08, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_address_proto_rawDescOnce sync.Once
	file_address_proto_rawDescData = file_address_proto_rawDesc
)

func file_address_proto_rawDescGZIP() []byte {
	file_address_proto_rawDescOnce.Do(func() {
		file_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_address_proto_rawDescData)
	})
	return file_address_proto_rawDescData
}

var file_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_address_proto_goTypes = []interface{}{
	(*Address)(nil), // 0: models.Address
}
var file_address_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_address_proto_init() }
func file_address_proto_init() {
	if File_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_address_proto_goTypes,
		DependencyIndexes: file_address_proto_depIdxs,
		MessageInfos:      file_address_proto_msgTypes,
	}.Build()
	File_address_proto = out.File
	file_address_proto_rawDesc = nil
	file_address_proto_goTypes = nil
	file_address_proto_depIdxs = nil
}
