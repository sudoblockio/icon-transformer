// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction_internal_by_address.proto

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

type TransactionInternalByAddress struct {
	TransactionHash      string   `protobuf:"bytes,1,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	LogIndex             int64    `protobuf:"varint,2,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	BlockNumber          int64    `protobuf:"varint,4,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
}

func (m *TransactionInternalByAddress) Reset()         { *m = TransactionInternalByAddress{} }
func (m *TransactionInternalByAddress) String() string { return proto.CompactTextString(m) }
func (*TransactionInternalByAddress) ProtoMessage()    {}
func (*TransactionInternalByAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_45c032238ff9de70, []int{0}
}

func (m *TransactionInternalByAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionInternalByAddress.Unmarshal(m, b)
}
func (m *TransactionInternalByAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionInternalByAddress.Marshal(b, m, deterministic)
}
func (m *TransactionInternalByAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionInternalByAddress.Merge(m, src)
}
func (m *TransactionInternalByAddress) XXX_Size() int {
	return xxx_messageInfo_TransactionInternalByAddress.Size(m)
}
func (m *TransactionInternalByAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionInternalByAddress.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionInternalByAddress proto.InternalMessageInfo

func (m *TransactionInternalByAddress) GetTransactionHash() string {
	if m != nil {
		return m.TransactionHash
	}
	return ""
}

func (m *TransactionInternalByAddress) GetLogIndex() int64 {
	if m != nil {
		return m.LogIndex
	}
	return 0
}

func (m *TransactionInternalByAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *TransactionInternalByAddress) GetBlockNumber() int64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*TransactionInternalByAddress)(nil), "models.TransactionInternalByAddress")
}

func init() {
	proto.RegisterFile("transaction_internal_by_address.proto", fileDescriptor_45c032238ff9de70)
}

var fileDescriptor_45c032238ff9de70 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4f, 0x84, 0x30,
	0x1c, 0xc5, 0x03, 0x77, 0x41, 0xae, 0x9a, 0x68, 0x98, 0xc0, 0x38, 0x5c, 0x2e, 0xb9, 0x84, 0xe5,
	0xc0, 0x78, 0x0e, 0xc6, 0x49, 0x99, 0xbc, 0xc5, 0x81, 0x38, 0xe9, 0xd0, 0xb4, 0x50, 0xa1, 0xb1,
	0xf4, 0x4f, 0x5a, 0x2e, 0xe1, 0x66, 0x3f, 0x15, 0x7c, 0x3a, 0x43, 0x39, 0x23, 0x71, 0xb9, 0xf1,
	0xf5, 0xbd, 0xfe, 0x5e, 0xfb, 0xd0, 0xba, 0x51, 0x44, 0x6a, 0x92, 0x35, 0x1c, 0x24, 0xe6, 0xb2,
	0x61, 0x4a, 0x12, 0x81, 0xe9, 0x01, 0x93, 0x3c, 0x57, 0x4c, 0xeb, 0xa8, 0x56, 0xd0, 0x80, 0xe7,
	0x54, 0x90, 0x33, 0xa1, 0xaf, 0x9f, 0x0a, 0xde, 0x94, 0x7b, 0x1a, 0x65, 0x50, 0xc5, 0x5c, 0x7e,
	0x02, 0x15, 0xd0, 0x42, 0xcd, 0x64, 0x6c, 0x62, 0xd9, 0xa6, 0x60, 0x72, 0x53, 0x80, 0xaa, 0x46,
	0x1d, 0x43, 0x3d, 0x70, 0x75, 0x3c, 0x1c, 0x8d, 0xa4, 0xd5, 0xb7, 0x8d, 0x6e, 0xde, 0xfe, 0x3a,
	0x77, 0xc7, 0xca, 0xe4, 0xf0, 0x3c, 0x16, 0x7a, 0x5b, 0x74, 0x35, 0x7d, 0x53, 0x49, 0x74, 0xe9,
	0x5b, 0x4b, 0x2b, 0x5c, 0x24, 0x6e, 0xdf, 0x05, 0x73, 0x64, 0x87, 0x56, 0x7a, 0x39, 0x49, 0xbc,
	0x10, 0x5d, 0x7a, 0x6b, 0xb4, 0x10, 0x50, 0x60, 0x2e, 0x73, 0xd6, 0xfa, 0xf6, 0xd2, 0x0a, 0x67,
	0x93, 0xb4, 0x2b, 0xa0, 0xd8, 0x0d, 0x8e, 0xb7, 0x42, 0x67, 0xc7, 0x7f, 0xf9, 0xb3, 0x7f, 0xc8,
	0x5f, 0xc3, 0xfb, 0x40, 0x17, 0x54, 0x40, 0xf6, 0x85, 0xe5, 0xbe, 0xa2, 0x4c, 0xf9, 0x73, 0x43,
	0x7b, 0xe8, 0xbb, 0xe0, 0x1e, 0xdd, 0xa5, 0xb7, 0x27, 0xf6, 0xc2, 0x3c, 0x6f, 0xf1, 0xf4, 0x7e,
	0x7a, 0x6e, 0xd4, 0xab, 0x11, 0x8f, 0x4e, 0xdf, 0x05, 0xb6, 0x6b, 0x25, 0xe8, 0xdd, 0x8d, 0xe2,
	0x71, 0x53, 0xea, 0x98, 0x61, 0xb6, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xf7, 0xa1, 0x8e,
	0x8b, 0x01, 0x00, 0x00,
}
