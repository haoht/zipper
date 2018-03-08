// Code generated by protoc-gen-go. DO NOT EDIT.
// source: block.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	block.proto
	message.proto
	transaction.proto

It has these top-level messages:
	BlockHeader
	Block
	Header
	Message
	GetBlocksMsg
	GetInvMsg
	OnBlockMsg
	OnTransactionMsg
	GetDataMsg
	ContractSpec
	TxHeader
	Transaction
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type BlockHeader struct {
	PreviousHash  string `protobuf:"bytes,1,opt,name=previousHash" json:"previousHash,omitempty"`
	TimeStamp     uint32 `protobuf:"varint,2,opt,name=timeStamp" json:"timeStamp,omitempty"`
	Nonce         uint32 `protobuf:"varint,3,opt,name=nonce" json:"nonce,omitempty"`
	TxsMerkleHash string `protobuf:"bytes,4,opt,name=txsMerkleHash" json:"txsMerkleHash,omitempty"`
	Height        uint32 `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto1.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BlockHeader) GetPreviousHash() string {
	if m != nil {
		return m.PreviousHash
	}
	return ""
}

func (m *BlockHeader) GetTimeStamp() uint32 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *BlockHeader) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *BlockHeader) GetTxsMerkleHash() string {
	if m != nil {
		return m.TxsMerkleHash
	}
	return ""
}

func (m *BlockHeader) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type Block struct {
	Header       *BlockHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto1.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto1.RegisterType((*BlockHeader)(nil), "proto.BlockHeader")
	proto1.RegisterType((*Block)(nil), "proto.Block")
}

func init() { proto1.RegisterFile("block.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xca, 0xc9, 0x4f,
	0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x82, 0x25, 0x45, 0x89,
	0x79, 0xc5, 0x89, 0xc9, 0x25, 0x99, 0xf9, 0x79, 0x10, 0x19, 0xa5, 0xc5, 0x8c, 0x5c, 0xdc, 0x4e,
	0x20, 0x95, 0x1e, 0xa9, 0x89, 0x29, 0xa9, 0x45, 0x42, 0x4a, 0x5c, 0x3c, 0x05, 0x45, 0xa9, 0x65,
	0x99, 0xf9, 0xa5, 0xc5, 0x1e, 0x89, 0xc5, 0x19, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x28,
	0x62, 0x42, 0x32, 0x5c, 0x9c, 0x25, 0x99, 0xb9, 0xa9, 0xc1, 0x25, 0x89, 0xb9, 0x05, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x08, 0x01, 0x21, 0x11, 0x2e, 0xd6, 0xbc, 0xfc, 0xbc, 0xe4, 0x54,
	0x09, 0x66, 0xb0, 0x0c, 0x84, 0x23, 0xa4, 0xc2, 0xc5, 0x5b, 0x52, 0x51, 0xec, 0x9b, 0x5a, 0x94,
	0x9d, 0x93, 0x0a, 0x36, 0x98, 0x05, 0x6c, 0x30, 0xaa, 0xa0, 0x90, 0x18, 0x17, 0x5b, 0x46, 0x6a,
	0x66, 0x7a, 0x46, 0x89, 0x04, 0x2b, 0x58, 0x33, 0x94, 0xa7, 0x94, 0xcd, 0xc5, 0x0a, 0x76, 0xa4,
	0x90, 0x16, 0x48, 0x01, 0xc8, 0xa1, 0x60, 0x87, 0x71, 0x1b, 0x09, 0x41, 0xbc, 0xa1, 0x87, 0xe4,
	0x85, 0x20, 0xa8, 0x0a, 0x21, 0x33, 0x2e, 0x1e, 0x24, 0xff, 0x16, 0x4b, 0x30, 0x29, 0x30, 0x23,
	0xe9, 0x08, 0x41, 0x48, 0x05, 0xa1, 0xa8, 0x4b, 0x62, 0x03, 0x2b, 0x30, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x0e, 0x73, 0xf6, 0x2d, 0x42, 0x01, 0x00, 0x00,
}
