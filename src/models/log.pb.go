// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log.proto

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

type Log struct {
	TransactionHash      string   `protobuf:"bytes,1,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash"`
	LogIndex             int64    `protobuf:"varint,2,opt,name=log_index,json=logIndex,proto3" json:"log_index"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address"`
	BlockNumber          int64    `protobuf:"varint,4,opt,name=block_number,json=blockNumber,proto3" json:"block_number"`
	Method               string   `protobuf:"bytes,5,opt,name=method,proto3" json:"method"`
	Data                 string   `protobuf:"bytes,9,opt,name=data,proto3" json:"data"`
	Indexed              string   `protobuf:"bytes,10,opt,name=indexed,proto3" json:"indexed"`
	BlockTimestamp       int64    `protobuf:"varint,11,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{0}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetTransactionHash() string {
	if m != nil {
		return m.TransactionHash
	}
	return ""
}

func (m *Log) GetLogIndex() int64 {
	if m != nil {
		return m.LogIndex
	}
	return 0
}

func (m *Log) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Log) GetBlockNumber() int64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *Log) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Log) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Log) GetIndexed() string {
	if m != nil {
		return m.Indexed
	}
	return ""
}

func (m *Log) GetBlockTimestamp() int64 {
	if m != nil {
		return m.BlockTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Log)(nil), "models.Log")
}

func init() {
	proto.RegisterFile("log.proto", fileDescriptor_a153da538f858886)
}

var fileDescriptor_a153da538f858886 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x6b, 0xdb, 0x40,
	0x10, 0x85, 0x91, 0xec, 0xca, 0xf2, 0xb8, 0xd8, 0xee, 0xb6, 0xb8, 0xab, 0x62, 0x8a, 0x31, 0x94,
	0xea, 0x62, 0x89, 0xd2, 0x5b, 0x2f, 0x0d, 0x3e, 0x25, 0x90, 0xe4, 0xb0, 0xe4, 0x94, 0x8b, 0x58,
	0x69, 0x37, 0x92, 0x88, 0xa4, 0x11, 0xda, 0x35, 0xf8, 0xaf, 0x49, 0x3f, 0x2e, 0x04, 0xad, 0x2c,
	0x30, 0xb9, 0xed, 0x7b, 0xf3, 0xcd, 0xbc, 0x59, 0x06, 0xe6, 0x05, 0xa6, 0x41, 0xdd, 0xa0, 0x46,
	0xe2, 0x94, 0x28, 0x64, 0xa1, 0x7e, 0xdc, 0xa4, 0xb9, 0xce, 0x4e, 0x71, 0x90, 0x60, 0x19, 0xe6,
	0xd5, 0x0b, 0xc6, 0x05, 0x9e, 0xb1, 0x96, 0x55, 0x68, 0xb0, 0xe4, 0x90, 0xca, 0xea, 0x90, 0x62,
	0x53, 0x0e, 0x3a, 0xc4, 0x5a, 0xe7, 0x58, 0xa9, 0xb0, 0xb7, 0x86, 0x49, 0xfb, 0x37, 0x1b, 0x26,
	0xf7, 0x98, 0x92, 0x07, 0x58, 0xeb, 0x86, 0x57, 0x8a, 0x27, 0x3d, 0x12, 0x65, 0x5c, 0x65, 0xd4,
	0xda, 0x59, 0xfe, 0xfc, 0xb8, 0xef, 0x5a, 0xef, 0x27, 0x6c, 0x7d, 0x8b, 0xd1, 0x02, 0xd3, 0x28,
	0x17, 0xe7, 0xe8, 0x23, 0xc9, 0x56, 0x57, 0xce, 0x2d, 0x57, 0x19, 0xf9, 0x65, 0xb6, 0x8d, 0xf2,
	0x4a, 0xc8, 0x33, 0xb5, 0x77, 0x96, 0x3f, 0x39, 0xba, 0x5d, 0xeb, 0x4d, 0xc1, 0xf6, 0x2d, 0xe6,
	0x16, 0x98, 0xde, 0xf5, 0x15, 0xf2, 0x07, 0x66, 0x5c, 0x88, 0x46, 0x2a, 0x45, 0x27, 0x26, 0xec,
	0x7b, 0xd7, 0x7a, 0x5f, 0xe1, 0x0b, 0x5b, 0x8d, 0x51, 0x97, 0x32, 0x1b, 0x39, 0xf2, 0x1f, 0x3e,
	0xc7, 0x05, 0x26, 0xaf, 0x51, 0x75, 0x2a, 0x63, 0xd9, 0xd0, 0xa9, 0x19, 0xbe, 0xed, 0x5a, 0x8f,
	0xc2, 0x86, 0x7d, 0x1b, 0xfb, 0xae, 0x19, 0xb6, 0x30, 0xea, 0xd1, 0x08, 0x12, 0x80, 0x53, 0x4a,
	0x9d, 0xa1, 0xa0, 0x9f, 0x4c, 0xe4, 0xa6, 0x6b, 0x3d, 0x02, 0x6b, 0xb6, 0x1c, 0x5b, 0x87, 0x2a,
	0xbb, 0x50, 0x84, 0xc0, 0x54, 0x70, 0xcd, 0xe9, 0xbc, 0xa7, 0x99, 0x79, 0x13, 0x0a, 0x33, 0xf3,
	0x35, 0x29, 0x28, 0x18, 0x7b, 0x94, 0xe4, 0x37, 0xac, 0x86, 0x68, 0x9d, 0x97, 0x52, 0x69, 0x5e,
	0xd6, 0x74, 0xd1, 0x6f, 0xc8, 0x96, 0xc6, 0x7e, 0x1a, 0xdd, 0x7f, 0x4e, 0xd7, 0x7a, 0xb6, 0x6b,
	0x1d, 0xe1, 0xd9, 0x0d, 0xc2, 0xe1, 0x9c, 0xb1, 0x63, 0x6e, 0xf2, 0xf7, 0x3d, 0x00, 0x00, 0xff,
	0xff, 0xeb, 0x44, 0x95, 0x1f, 0xea, 0x01, 0x00, 0x00,
}
