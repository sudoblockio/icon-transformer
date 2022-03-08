// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: block_etl.proto

package models

import (
	_ "github.com/mwitkow/go-proto-validators"
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

type BlockETL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number         int64             `protobuf:"varint,1,opt,name=number,proto3" json:"number"`
	Hash           string            `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash"`
	ParentHash     string            `protobuf:"bytes,3,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash"`
	MerkleRootHash string            `protobuf:"bytes,4,opt,name=merkle_root_hash,json=merkleRootHash,proto3" json:"merkle_root_hash"`
	PeerId         string            `protobuf:"bytes,5,opt,name=peer_id,json=peerId,proto3" json:"peer_id"`
	Signature      string            `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature"`
	Timestamp      int64             `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp"`
	Version        string            `protobuf:"bytes,8,opt,name=version,proto3" json:"version"`
	Transactions   []*TransactionETL `protobuf:"bytes,9,rep,name=transactions,proto3" json:"transactions"`
}

func (x *BlockETL) Reset() {
	*x = BlockETL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_etl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockETL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockETL) ProtoMessage() {}

func (x *BlockETL) ProtoReflect() protoreflect.Message {
	mi := &file_block_etl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockETL.ProtoReflect.Descriptor instead.
func (*BlockETL) Descriptor() ([]byte, []int) {
	return file_block_etl_proto_rawDescGZIP(), []int{0}
}

func (x *BlockETL) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *BlockETL) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *BlockETL) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *BlockETL) GetMerkleRootHash() string {
	if x != nil {
		return x.MerkleRootHash
	}
	return ""
}

func (x *BlockETL) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *BlockETL) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *BlockETL) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BlockETL) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *BlockETL) GetTransactions() []*TransactionETL {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type TransactionETL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash               string    `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash"`
	Timestamp          int64     `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp"`
	TransactionIndex   int64     `protobuf:"varint,3,opt,name=transaction_index,json=transactionIndex,proto3" json:"transaction_index"`
	Nonce              string    `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce"`
	Nid                string    `protobuf:"bytes,5,opt,name=nid,proto3" json:"nid"`
	FromAddress        string    `protobuf:"bytes,6,opt,name=from_address,json=fromAddress,proto3" json:"from_address"`
	ToAddress          string    `protobuf:"bytes,7,opt,name=to_address,json=toAddress,proto3" json:"to_address"`
	Value              string    `protobuf:"bytes,8,opt,name=value,proto3" json:"value"`
	Status             string    `protobuf:"bytes,9,opt,name=status,proto3" json:"status"`
	StepPrice          string    `protobuf:"bytes,10,opt,name=step_price,json=stepPrice,proto3" json:"step_price"`
	StepUsed           string    `protobuf:"bytes,11,opt,name=step_used,json=stepUsed,proto3" json:"step_used"`
	StepLimit          string    `protobuf:"bytes,12,opt,name=step_limit,json=stepLimit,proto3" json:"step_limit"`
	CumulativeStepUsed string    `protobuf:"bytes,13,opt,name=cumulative_step_used,json=cumulativeStepUsed,proto3" json:"cumulative_step_used"`
	LogsBloom          string    `protobuf:"bytes,14,opt,name=logs_bloom,json=logsBloom,proto3" json:"logs_bloom"`
	Data               string    `protobuf:"bytes,15,opt,name=data,proto3" json:"data"`
	DataType           string    `protobuf:"bytes,16,opt,name=data_type,json=dataType,proto3" json:"data_type"`
	ScoreAddress       string    `protobuf:"bytes,17,opt,name=score_address,json=scoreAddress,proto3" json:"score_address"`
	Signature          string    `protobuf:"bytes,18,opt,name=signature,proto3" json:"signature"`
	Version            string    `protobuf:"bytes,19,opt,name=version,proto3" json:"version"`
	Logs               []*LogETL `protobuf:"bytes,20,rep,name=logs,proto3" json:"logs"`
}

func (x *TransactionETL) Reset() {
	*x = TransactionETL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_etl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionETL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionETL) ProtoMessage() {}

func (x *TransactionETL) ProtoReflect() protoreflect.Message {
	mi := &file_block_etl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionETL.ProtoReflect.Descriptor instead.
func (*TransactionETL) Descriptor() ([]byte, []int) {
	return file_block_etl_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionETL) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *TransactionETL) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *TransactionETL) GetTransactionIndex() int64 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

func (x *TransactionETL) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *TransactionETL) GetNid() string {
	if x != nil {
		return x.Nid
	}
	return ""
}

func (x *TransactionETL) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *TransactionETL) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *TransactionETL) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TransactionETL) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TransactionETL) GetStepPrice() string {
	if x != nil {
		return x.StepPrice
	}
	return ""
}

func (x *TransactionETL) GetStepUsed() string {
	if x != nil {
		return x.StepUsed
	}
	return ""
}

func (x *TransactionETL) GetStepLimit() string {
	if x != nil {
		return x.StepLimit
	}
	return ""
}

func (x *TransactionETL) GetCumulativeStepUsed() string {
	if x != nil {
		return x.CumulativeStepUsed
	}
	return ""
}

func (x *TransactionETL) GetLogsBloom() string {
	if x != nil {
		return x.LogsBloom
	}
	return ""
}

func (x *TransactionETL) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *TransactionETL) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *TransactionETL) GetScoreAddress() string {
	if x != nil {
		return x.ScoreAddress
	}
	return ""
}

func (x *TransactionETL) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *TransactionETL) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *TransactionETL) GetLogs() []*LogETL {
	if x != nil {
		return x.Logs
	}
	return nil
}

type LogETL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address"`
	Indexed []string `protobuf:"bytes,2,rep,name=indexed,proto3" json:"indexed"`
	Data    []string `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *LogETL) Reset() {
	*x = LogETL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_etl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogETL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogETL) ProtoMessage() {}

func (x *LogETL) ProtoReflect() protoreflect.Message {
	mi := &file_block_etl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogETL.ProtoReflect.Descriptor instead.
func (*LogETL) Descriptor() ([]byte, []int) {
	return file_block_etl_proto_rawDescGZIP(), []int{2}
}

func (x *LogETL) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *LogETL) GetIndexed() []string {
	if x != nil {
		return x.Indexed
	}
	return nil
}

func (x *LogETL) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_block_etl_proto protoreflect.FileDescriptor

var file_block_etl_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x65, 0x74, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x96, 0x03, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x54, 0x4c, 0x12, 0x1e,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xe2, 0xdf,
	0x1f, 0x10, 0x0a, 0x0e, 0x5e, 0x5b, 0x61, 0x2d, 0x66, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x36, 0x34,
	0x7d, 0x24, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x35, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xe2,
	0xdf, 0x1f, 0x10, 0x0a, 0x0e, 0x5e, 0x5b, 0x61, 0x2d, 0x66, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x36,
	0x34, 0x7d, 0x24, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x3e, 0x0a, 0x10, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xe2, 0xdf, 0x1f, 0x10, 0x0a,
	0x0e, 0x5e, 0x5b, 0x61, 0x2d, 0x66, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x36, 0x34, 0x7d, 0x24, 0x52,
	0x0e, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x2f, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x16, 0xe2, 0xdf, 0x1f, 0x12, 0x0a, 0x10, 0x5e, 0x68, 0x78, 0x5b, 0x61, 0x2d, 0x66, 0x30,
	0x2d, 0x39, 0x5d, 0x7b, 0x34, 0x30, 0x7d, 0x24, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x24,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3a,
	0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x54, 0x4c, 0x52, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x95, 0x05, 0x0a, 0x0e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x54, 0x4c, 0x12, 0x2a, 0x0a,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xe2, 0xdf, 0x1f,
	0x12, 0x0a, 0x10, 0x5e, 0x30, 0x78, 0x5b, 0x61, 0x2d, 0x66, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x36,
	0x34, 0x7d, 0x24, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x10, 0x00, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6e, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x26, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xe2, 0xdf,
	0x1f, 0x0a, 0x0a, 0x08, 0x5e, 0x30, 0x78, 0x5b, 0x30, 0x31, 0x5d, 0x24, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x65, 0x70, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x65, 0x70, 0x55, 0x73, 0x65, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x65, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x30, 0x0a, 0x14, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x74,
	0x65, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63,
	0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x65, 0x70, 0x55, 0x73, 0x65,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x6c, 0x6f, 0x6f, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22,
	0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x54, 0x4c, 0x52, 0x04, 0x6c, 0x6f,
	0x67, 0x73, 0x22, 0x68, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x45, 0x54, 0x4c, 0x12, 0x30, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xe2,
	0xdf, 0x1f, 0x12, 0x0a, 0x10, 0x5e, 0x63, 0x78, 0x5b, 0x61, 0x2d, 0x66, 0x30, 0x2d, 0x39, 0x5d,
	0x7b, 0x34, 0x30, 0x7d, 0x24, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_block_etl_proto_rawDescOnce sync.Once
	file_block_etl_proto_rawDescData = file_block_etl_proto_rawDesc
)

func file_block_etl_proto_rawDescGZIP() []byte {
	file_block_etl_proto_rawDescOnce.Do(func() {
		file_block_etl_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_etl_proto_rawDescData)
	})
	return file_block_etl_proto_rawDescData
}

var file_block_etl_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_block_etl_proto_goTypes = []interface{}{
	(*BlockETL)(nil),       // 0: models.BlockETL
	(*TransactionETL)(nil), // 1: models.TransactionETL
	(*LogETL)(nil),         // 2: models.LogETL
}
var file_block_etl_proto_depIdxs = []int32{
	1, // 0: models.BlockETL.transactions:type_name -> models.TransactionETL
	2, // 1: models.TransactionETL.logs:type_name -> models.LogETL
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_block_etl_proto_init() }
func file_block_etl_proto_init() {
	if File_block_etl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_block_etl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockETL); i {
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
		file_block_etl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionETL); i {
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
		file_block_etl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogETL); i {
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
			RawDescriptor: file_block_etl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_block_etl_proto_goTypes,
		DependencyIndexes: file_block_etl_proto_depIdxs,
		MessageInfos:      file_block_etl_proto_msgTypes,
	}.Build()
	File_block_etl_proto = out.File
	file_block_etl_proto_rawDesc = nil
	file_block_etl_proto_goTypes = nil
	file_block_etl_proto_depIdxs = nil
}
