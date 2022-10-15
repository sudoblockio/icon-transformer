// Code generated by protoc-gen-go. DO NOT EDIT.
// source: token_transfer.proto

package models

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
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

type TokenTransfer struct {
	TransactionHash      string   `protobuf:"bytes,1,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash"`
	LogIndex             int64    `protobuf:"varint,2,opt,name=log_index,json=logIndex,proto3" json:"log_index"`
	TokenContractAddress string   `protobuf:"bytes,3,opt,name=token_contract_address,json=tokenContractAddress,proto3" json:"token_contract_address"`
	FromAddress          string   `protobuf:"bytes,4,opt,name=from_address,json=fromAddress,proto3" json:"from_address"`
	ToAddress            string   `protobuf:"bytes,5,opt,name=to_address,json=toAddress,proto3" json:"to_address"`
	BlockNumber          int64    `protobuf:"varint,6,opt,name=block_number,json=blockNumber,proto3" json:"block_number"`
	Value                string   `protobuf:"bytes,7,opt,name=value,proto3" json:"value"`
	ValueDecimal         float64  `protobuf:"fixed64,8,opt,name=value_decimal,json=valueDecimal,proto3" json:"value_decimal"`
	BlockTimestamp       int64    `protobuf:"varint,9,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp"`
	TokenContractName    string   `protobuf:"bytes,10,opt,name=token_contract_name,json=tokenContractName,proto3" json:"token_contract_name"`
	TokenContractSymbol  string   `protobuf:"bytes,11,opt,name=token_contract_symbol,json=tokenContractSymbol,proto3" json:"token_contract_symbol"`
	TransactionFee       string   `protobuf:"bytes,12,opt,name=transaction_fee,json=transactionFee,proto3" json:"transaction_fee"`
	NftId                int64    `protobuf:"varint,13,opt,name=nft_id,json=nftId,proto3" json:"nft_id"`
}

func (m *TokenTransfer) Reset()         { *m = TokenTransfer{} }
func (m *TokenTransfer) String() string { return proto.CompactTextString(m) }
func (*TokenTransfer) ProtoMessage()    {}
func (*TokenTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dee5df8e2b2416f, []int{0}
}

func (m *TokenTransfer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenTransfer.Unmarshal(m, b)
}
func (m *TokenTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenTransfer.Marshal(b, m, deterministic)
}
func (m *TokenTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenTransfer.Merge(m, src)
}
func (m *TokenTransfer) XXX_Size() int {
	return xxx_messageInfo_TokenTransfer.Size(m)
}
func (m *TokenTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_TokenTransfer proto.InternalMessageInfo

func (m *TokenTransfer) GetTransactionHash() string {
	if m != nil {
		return m.TransactionHash
	}
	return ""
}

func (m *TokenTransfer) GetLogIndex() int64 {
	if m != nil {
		return m.LogIndex
	}
	return 0
}

func (m *TokenTransfer) GetTokenContractAddress() string {
	if m != nil {
		return m.TokenContractAddress
	}
	return ""
}

func (m *TokenTransfer) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *TokenTransfer) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *TokenTransfer) GetBlockNumber() int64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *TokenTransfer) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *TokenTransfer) GetValueDecimal() float64 {
	if m != nil {
		return m.ValueDecimal
	}
	return 0
}

func (m *TokenTransfer) GetBlockTimestamp() int64 {
	if m != nil {
		return m.BlockTimestamp
	}
	return 0
}

func (m *TokenTransfer) GetTokenContractName() string {
	if m != nil {
		return m.TokenContractName
	}
	return ""
}

func (m *TokenTransfer) GetTokenContractSymbol() string {
	if m != nil {
		return m.TokenContractSymbol
	}
	return ""
}

func (m *TokenTransfer) GetTransactionFee() string {
	if m != nil {
		return m.TransactionFee
	}
	return ""
}

func (m *TokenTransfer) GetNftId() int64 {
	if m != nil {
		return m.NftId
	}
	return 0
}

func init() {
	proto.RegisterType((*TokenTransfer)(nil), "models.TokenTransfer")
}

func init() {
	proto.RegisterFile("token_transfer.proto", fileDescriptor_4dee5df8e2b2416f)
}

var fileDescriptor_4dee5df8e2b2416f = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x55, 0xba, 0xb5, 0xa4, 0x5f, 0xdb, 0x0d, 0xbc, 0x0d, 0x39, 0x48, 0x68, 0x1d, 0xd3, 0xb4,
	0x22, 0xd4, 0x46, 0xb0, 0x1b, 0x27, 0x18, 0x13, 0x62, 0x1c, 0x76, 0x30, 0x3d, 0x71, 0x89, 0x9c,
	0xc4, 0x49, 0xa3, 0xc5, 0xfe, 0xaa, 0xc4, 0x45, 0xe5, 0xa7, 0x25, 0x47, 0x7e, 0x19, 0x8a, 0xbd,
	0x55, 0xe9, 0x54, 0xed, 0x96, 0xf7, 0xfc, 0xbe, 0xf7, 0x9e, 0x3f, 0x39, 0x70, 0xac, 0xf1, 0x5e,
	0xa8, 0x40, 0x17, 0x5c, 0x95, 0x89, 0x28, 0x66, 0xcb, 0x02, 0x35, 0x92, 0x9e, 0xc4, 0x58, 0xe4,
	0xe5, 0x9b, 0x2f, 0x69, 0xa6, 0x17, 0xab, 0x70, 0x16, 0xa1, 0xf4, 0x33, 0x95, 0x60, 0x98, 0xe3,
	0x1a, 0x97, 0x42, 0xf9, 0x46, 0x16, 0x4d, 0x53, 0xa1, 0xa6, 0x29, 0x16, 0xd2, 0x62, 0x1f, 0x97,
	0x3a, 0x43, 0x55, 0xfa, 0x0d, 0x65, 0x9d, 0xde, 0xfd, 0xeb, 0xc2, 0x68, 0xde, 0x44, 0xcc, 0x1f,
	0x12, 0xc8, 0x15, 0xbc, 0x34, 0x69, 0x3c, 0x6a, 0xc4, 0xc1, 0x82, 0x97, 0x0b, 0xea, 0x8c, 0x9d,
	0x49, 0xff, 0xda, 0xad, 0x2b, 0x6f, 0x1f, 0x3a, 0x13, 0x87, 0x1d, 0xb6, 0x14, 0x3f, 0x78, 0xb9,
	0x20, 0x17, 0xd0, 0xcf, 0x31, 0x0d, 0x32, 0x15, 0x8b, 0x35, 0xed, 0x8c, 0x9d, 0xc9, 0x5e, 0x4b,
	0xed, 0xe6, 0x98, 0xde, 0x36, 0x27, 0x24, 0x85, 0xd7, 0xf6, 0x3e, 0x11, 0x2a, 0x5d, 0xf0, 0x48,
	0x07, 0x3c, 0x8e, 0x0b, 0x51, 0x96, 0x74, 0xcf, 0x24, 0x7c, 0xac, 0x2b, 0x6f, 0x0a, 0x1f, 0xd8,
	0xfb, 0xed, 0x5b, 0x07, 0x59, 0xbc, 0x0e, 0x76, 0x0f, 0x32, 0xbb, 0xa0, 0x6f, 0x0f, 0xf4, 0x57,
	0xcb, 0x92, 0x9f, 0x30, 0x4c, 0x0a, 0x94, 0x1b, 0xfb, 0x7d, 0x63, 0x7f, 0x59, 0x57, 0xde, 0x39,
	0x9c, 0xb1, 0xd3, 0x1d, 0xf6, 0x6d, 0x39, 0x1b, 0x34, 0xe8, 0xd1, 0xeb, 0x06, 0x40, 0xe3, 0xc6,
	0xa9, 0x6b, 0x9c, 0x2e, 0xea, 0xca, 0x3b, 0x83, 0x53, 0xf6, 0x76, 0x67, 0xd1, 0x8d, 0x4f, 0x5f,
	0x63, 0xab, 0x51, 0x98, 0x63, 0x74, 0x1f, 0xa8, 0x95, 0x0c, 0x45, 0x41, 0x7b, 0x66, 0x49, 0xcf,
	0x35, 0x6a, 0xcb, 0xd9, 0xc0, 0xa0, 0x3b, 0x03, 0xc8, 0x31, 0x74, 0xff, 0xf0, 0x7c, 0x25, 0xe8,
	0x8b, 0xa6, 0x0c, 0xb3, 0x80, 0x9c, 0xc3, 0xc8, 0x7c, 0x04, 0xb1, 0x88, 0x32, 0xc9, 0x73, 0xea,
	0x8e, 0x9d, 0x89, 0xc3, 0x86, 0x86, 0xbc, 0xb1, 0x1c, 0xb9, 0x84, 0x43, 0xeb, 0xab, 0x33, 0x29,
	0x4a, 0xcd, 0xe5, 0x92, 0xf6, 0x9b, 0x26, 0xec, 0xc0, 0xd0, 0xf3, 0x47, 0x96, 0xcc, 0xe0, 0xe8,
	0xc9, 0xc6, 0x15, 0x97, 0x82, 0x82, 0x49, 0x7c, 0xb5, 0xb5, 0xf4, 0x3b, 0x2e, 0x05, 0xf9, 0x04,
	0x27, 0x4f, 0xf4, 0xe5, 0x5f, 0x19, 0x62, 0x4e, 0x07, 0x66, 0xe2, 0x68, 0x6b, 0xe2, 0x97, 0x39,
	0x6a, 0xca, 0xb4, 0x9f, 0x5a, 0x22, 0x04, 0x1d, 0x1a, 0xf5, 0x41, 0x8b, 0xfe, 0x2e, 0x04, 0x39,
	0x81, 0x9e, 0x4a, 0x74, 0x90, 0xc5, 0x74, 0x64, 0xca, 0x76, 0x55, 0xa2, 0x6f, 0xe3, 0xcf, 0xbd,
	0xba, 0xf2, 0x3a, 0xae, 0x73, 0x0d, 0xbf, 0xdd, 0x99, 0x6f, 0x7f, 0x89, 0xb0, 0x67, 0xde, 0xf5,
	0xd5, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0x66, 0x4b, 0x89, 0x39, 0x03, 0x00, 0x00,
}
