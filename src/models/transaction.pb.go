// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: transaction.proto

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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash               string  `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash"`
	LogIndex           int64   `protobuf:"varint,2,opt,name=log_index,json=logIndex,proto3" json:"log_index"`
	Type               string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`
	Method             string  `protobuf:"bytes,4,opt,name=method,proto3" json:"method"`
	FromAddress        string  `protobuf:"bytes,5,opt,name=from_address,json=fromAddress,proto3" json:"from_address"`
	ToAddress          string  `protobuf:"bytes,6,opt,name=to_address,json=toAddress,proto3" json:"to_address"`
	BlockNumber        int64   `protobuf:"varint,7,opt,name=block_number,json=blockNumber,proto3" json:"block_number"`
	Version            string  `protobuf:"bytes,8,opt,name=version,proto3" json:"version"`
	Value              string  `protobuf:"bytes,9,opt,name=value,proto3" json:"value"`
	ValueDecimal       float64 `protobuf:"fixed64,10,opt,name=value_decimal,json=valueDecimal,proto3" json:"value_decimal"`
	StepLimit          string  `protobuf:"bytes,11,opt,name=step_limit,json=stepLimit,proto3" json:"step_limit"`
	Timestamp          int64   `protobuf:"varint,12,opt,name=timestamp,proto3" json:"timestamp"`
	BlockTimestamp     int64   `protobuf:"varint,13,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp"`
	Nid                string  `protobuf:"bytes,14,opt,name=nid,proto3" json:"nid"`
	Nonce              string  `protobuf:"bytes,15,opt,name=nonce,proto3" json:"nonce"`
	TransactionIndex   int64   `protobuf:"varint,16,opt,name=transaction_index,json=transactionIndex,proto3" json:"transaction_index"`
	BlockHash          string  `protobuf:"bytes,17,opt,name=block_hash,json=blockHash,proto3" json:"block_hash"`
	TransactionFee     string  `protobuf:"bytes,18,opt,name=transaction_fee,json=transactionFee,proto3" json:"transaction_fee"`
	Signature          string  `protobuf:"bytes,19,opt,name=signature,proto3" json:"signature"`
	DataType           string  `protobuf:"bytes,20,opt,name=data_type,json=dataType,proto3" json:"data_type"`
	Data               string  `protobuf:"bytes,21,opt,name=data,proto3" json:"data"`
	CumulativeStepUsed string  `protobuf:"bytes,22,opt,name=cumulative_step_used,json=cumulativeStepUsed,proto3" json:"cumulative_step_used"`
	StepUsed           string  `protobuf:"bytes,23,opt,name=step_used,json=stepUsed,proto3" json:"step_used"`
	StepPrice          string  `protobuf:"bytes,24,opt,name=step_price,json=stepPrice,proto3" json:"step_price"`
	ScoreAddress       string  `protobuf:"bytes,25,opt,name=score_address,json=scoreAddress,proto3" json:"score_address"`
	LogsBloom          string  `protobuf:"bytes,26,opt,name=logs_bloom,json=logsBloom,proto3" json:"logs_bloom"`
	Status             string  `protobuf:"bytes,27,opt,name=status,proto3" json:"status"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Transaction) GetLogIndex() int64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *Transaction) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Transaction) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Transaction) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *Transaction) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *Transaction) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *Transaction) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Transaction) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Transaction) GetValueDecimal() float64 {
	if x != nil {
		return x.ValueDecimal
	}
	return 0
}

func (x *Transaction) GetStepLimit() string {
	if x != nil {
		return x.StepLimit
	}
	return ""
}

func (x *Transaction) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Transaction) GetBlockTimestamp() int64 {
	if x != nil {
		return x.BlockTimestamp
	}
	return 0
}

func (x *Transaction) GetNid() string {
	if x != nil {
		return x.Nid
	}
	return ""
}

func (x *Transaction) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *Transaction) GetTransactionIndex() int64 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

func (x *Transaction) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *Transaction) GetTransactionFee() string {
	if x != nil {
		return x.TransactionFee
	}
	return ""
}

func (x *Transaction) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Transaction) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *Transaction) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Transaction) GetCumulativeStepUsed() string {
	if x != nil {
		return x.CumulativeStepUsed
	}
	return ""
}

func (x *Transaction) GetStepUsed() string {
	if x != nil {
		return x.StepUsed
	}
	return ""
}

func (x *Transaction) GetStepPrice() string {
	if x != nil {
		return x.StepPrice
	}
	return ""
}

func (x *Transaction) GetScoreAddress() string {
	if x != nil {
		return x.ScoreAddress
	}
	return ""
}

func (x *Transaction) GetLogsBloom() string {
	if x != nil {
		return x.LogsBloom
	}
	return ""
}

func (x *Transaction) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_transaction_proto protoreflect.FileDescriptor

var file_transaction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78,
	0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x08, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x28, 0x01, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x25, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02,
	0x28, 0x01, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x30, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xba, 0xb9, 0x19, 0x18,
	0x0a, 0x16, 0x52, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x36,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e,
	0xba, 0xb9, 0x19, 0x1a, 0x0a, 0x18, 0x52, 0x16, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x47, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xba, 0xb9,
	0x19, 0x20, 0x0a, 0x1e, 0x52, 0x1c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x41, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x22, 0xba, 0xb9, 0x19, 0x1e, 0x0a, 0x1c, 0x52, 0x1a, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x74, 0x6f, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x47, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x24, 0xba, 0xb9, 0x19, 0x20, 0x0a, 0x1e,
	0x52, 0x1c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x78, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x65, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x27, 0x0a,
	0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x64, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x2b,
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x65, 0x70,
	0x55, 0x73, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x75, 0x73, 0x65,
	0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x65, 0x70, 0x55, 0x73, 0x65,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x65, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x62, 0x6c,
	0x6f, 0x6f, 0x6d, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x73, 0x42,
	0x6c, 0x6f, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x1b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x06, 0xba, 0xb9,
	0x19, 0x02, 0x08, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_proto_rawDescOnce sync.Once
	file_transaction_proto_rawDescData = file_transaction_proto_rawDesc
)

func file_transaction_proto_rawDescGZIP() []byte {
	file_transaction_proto_rawDescOnce.Do(func() {
		file_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_proto_rawDescData)
	})
	return file_transaction_proto_rawDescData
}

var file_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transaction_proto_goTypes = []interface{}{
	(*Transaction)(nil), // 0: models.Transaction
}
var file_transaction_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transaction_proto_init() }
func file_transaction_proto_init() {
	if File_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
			RawDescriptor: file_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_proto_goTypes,
		DependencyIndexes: file_transaction_proto_depIdxs,
		MessageInfos:      file_transaction_proto_msgTypes,
	}.Build()
	File_transaction_proto = out.File
	file_transaction_proto_rawDesc = nil
	file_transaction_proto_goTypes = nil
	file_transaction_proto_depIdxs = nil
}