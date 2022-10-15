// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contract_processed.proto

package models

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Used for proto and gorm
type ContractProcessed struct {
	// Address related items
	Address              string `protobuf:"bytes,1,opt,name=address,proto3" json:"address"`
	Name                 string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	CreatedTimestamp     int64  `protobuf:"varint,3,opt,name=created_timestamp,json=createdTimestamp,proto3" json:"created_timestamp"`
	Status               string `protobuf:"bytes,4,opt,name=status,proto3" json:"status"`
	IsToken              bool   `protobuf:"varint,5,opt,name=is_token,json=isToken,proto3" json:"is_token"`
	ContractUpdatedBlock int64  `protobuf:"varint,6,opt,name=contract_updated_block,json=contractUpdatedBlock,proto3" json:"contract_updated_block"`
	ContractType         string `protobuf:"bytes,7,opt,name=contract_type,json=contractType,proto3" json:"contract_type"`
	TokenStandard        string `protobuf:"bytes,8,opt,name=token_standard,json=tokenStandard,proto3" json:"token_standard"`
	Symbol               string `protobuf:"bytes,9,opt,name=symbol,proto3" json:"symbol"`
	// Transaction related items
	TransactionHash      string   `protobuf:"bytes,15,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash"`
	IsCreation           bool     `protobuf:"varint,16,opt,name=is_creation,json=isCreation,proto3" json:"is_creation"`
}

func (m *ContractProcessed) Reset()         { *m = ContractProcessed{} }
func (m *ContractProcessed) String() string { return proto.CompactTextString(m) }
func (*ContractProcessed) ProtoMessage()    {}
func (*ContractProcessed) Descriptor() ([]byte, []int) {
	return fileDescriptor_5993f9666cfc1736, []int{0}
}

func (m *ContractProcessed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractProcessed.Unmarshal(m, b)
}
func (m *ContractProcessed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractProcessed.Marshal(b, m, deterministic)
}
func (m *ContractProcessed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractProcessed.Merge(m, src)
}
func (m *ContractProcessed) XXX_Size() int {
	return xxx_messageInfo_ContractProcessed.Size(m)
}
func (m *ContractProcessed) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractProcessed.DiscardUnknown(m)
}

var xxx_messageInfo_ContractProcessed proto.InternalMessageInfo

func (m *ContractProcessed) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ContractProcessed) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContractProcessed) GetCreatedTimestamp() int64 {
	if m != nil {
		return m.CreatedTimestamp
	}
	return 0
}

func (m *ContractProcessed) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ContractProcessed) GetIsToken() bool {
	if m != nil {
		return m.IsToken
	}
	return false
}

func (m *ContractProcessed) GetContractUpdatedBlock() int64 {
	if m != nil {
		return m.ContractUpdatedBlock
	}
	return 0
}

func (m *ContractProcessed) GetContractType() string {
	if m != nil {
		return m.ContractType
	}
	return ""
}

func (m *ContractProcessed) GetTokenStandard() string {
	if m != nil {
		return m.TokenStandard
	}
	return ""
}

func (m *ContractProcessed) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ContractProcessed) GetTransactionHash() string {
	if m != nil {
		return m.TransactionHash
	}
	return ""
}

func (m *ContractProcessed) GetIsCreation() bool {
	if m != nil {
		return m.IsCreation
	}
	return false
}

func init() {
	proto.RegisterType((*ContractProcessed)(nil), "models.ContractProcessed")
}

func init() {
	proto.RegisterFile("contract_processed.proto", fileDescriptor_5993f9666cfc1736)
}

var fileDescriptor_5993f9666cfc1736 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0x53, 0xc1, 0x52, 0x46, 0x11, 0xd8, 0x18, 0x32, 0x9e, 0x24, 0x1a, 0x13, 0x8c, 0x09,
	0x1e, 0xf4, 0x17, 0xc0, 0xc5, 0xa3, 0x41, 0xbc, 0x78, 0x69, 0x96, 0xee, 0x26, 0x6c, 0xa0, 0xbb,
	0xcd, 0xce, 0x70, 0xe8, 0x2f, 0xf0, 0x6f, 0x1b, 0x86, 0x96, 0x78, 0xeb, 0xfb, 0xde, 0x6b, 0xdf,
	0x4c, 0x07, 0xb0, 0x08, 0x9e, 0xa3, 0x2e, 0x38, 0xaf, 0x62, 0x28, 0x2c, 0x91, 0x35, 0xf3, 0x2a,
	0x06, 0x0e, 0x2a, 0x2d, 0x83, 0xb1, 0x7b, 0x7a, 0xf8, 0xed, 0xc0, 0x78, 0xd9, 0x84, 0x3e, 0xdb,
	0x8c, 0x42, 0xe8, 0x69, 0x63, 0xa2, 0x25, 0xc2, 0x64, 0x9a, 0xcc, 0xfa, 0xab, 0x56, 0x2a, 0x05,
	0x5d, 0xaf, 0x4b, 0x8b, 0x17, 0x82, 0xe5, 0x59, 0xbd, 0xc0, 0xb8, 0x88, 0x56, 0xb3, 0x35, 0x39,
	0xbb, 0xd2, 0x12, 0xeb, 0xb2, 0xc2, 0xce, 0x34, 0x99, 0x75, 0x56, 0xa3, 0xc6, 0x58, 0xb7, 0x5c,
	0x4d, 0x20, 0x25, 0xd6, 0x7c, 0x20, 0xec, 0xca, 0x27, 0x1a, 0xa5, 0xee, 0x20, 0x73, 0x94, 0x73,
	0xd8, 0x59, 0x8f, 0x97, 0xd3, 0x64, 0x96, 0xad, 0x7a, 0x8e, 0xd6, 0x47, 0xa9, 0xde, 0x61, 0x72,
	0xde, 0xe3, 0x50, 0x19, 0x29, 0xda, 0xec, 0x43, 0xb1, 0xc3, 0x54, 0x4a, 0x6e, 0x5b, 0xf7, 0xfb,
	0x64, 0x2e, 0x8e, 0x9e, 0x7a, 0x84, 0xc1, 0xf9, 0x2d, 0xae, 0x2b, 0x8b, 0x3d, 0xe9, 0xbb, 0x6e,
	0xe1, 0xba, 0xae, 0xac, 0x7a, 0x82, 0x1b, 0xa9, 0xcc, 0x89, 0xb5, 0x37, 0x3a, 0x1a, 0xcc, 0x24,
	0x35, 0x10, 0xfa, 0xd5, 0x40, 0x19, 0xba, 0x2e, 0x37, 0x61, 0x8f, 0xfd, 0x66, 0x68, 0x51, 0xea,
	0x19, 0x46, 0x1c, 0xb5, 0x27, 0x5d, 0xb0, 0x0b, 0x3e, 0xdf, 0x6a, 0xda, 0xe2, 0x50, 0x12, 0xc3,
	0x7f, 0xfc, 0x43, 0xd3, 0x56, 0xdd, 0xc3, 0x95, 0xa3, 0x5c, 0x7e, 0x87, 0x0b, 0x1e, 0x47, 0xb2,
	0x22, 0x38, 0x5a, 0x36, 0x64, 0x01, 0x3f, 0xd9, 0xfc, 0xf5, 0x74, 0x95, 0x4d, 0x2a, 0x47, 0x7a,
	0xfb, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x86, 0xaa, 0xe5, 0xc0, 0x01, 0x00, 0x00,
}
