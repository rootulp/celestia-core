// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/types/validator.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	keys "github.com/lazyledger/lazyledger-core/proto/crypto/keys"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ValidatorSet struct {
	Validators           []*Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
	Proposer             *Validator   `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	TotalVotingPower     int64        `protobuf:"varint,3,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ValidatorSet) Reset()         { *m = ValidatorSet{} }
func (m *ValidatorSet) String() string { return proto.CompactTextString(m) }
func (*ValidatorSet) ProtoMessage()    {}
func (*ValidatorSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e7c6b38c20e5406, []int{0}
}
func (m *ValidatorSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorSet.Unmarshal(m, b)
}
func (m *ValidatorSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorSet.Marshal(b, m, deterministic)
}
func (m *ValidatorSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSet.Merge(m, src)
}
func (m *ValidatorSet) XXX_Size() int {
	return xxx_messageInfo_ValidatorSet.Size(m)
}
func (m *ValidatorSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSet.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSet proto.InternalMessageInfo

func (m *ValidatorSet) GetValidators() []*Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *ValidatorSet) GetProposer() *Validator {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *ValidatorSet) GetTotalVotingPower() int64 {
	if m != nil {
		return m.TotalVotingPower
	}
	return 0
}

type Validator struct {
	Address              []byte         `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey               keys.PublicKey `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key"`
	VotingPower          int64          `protobuf:"varint,3,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
	ProposerPriority     int64          `protobuf:"varint,4,opt,name=proposer_priority,json=proposerPriority,proto3" json:"proposer_priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e7c6b38c20e5406, []int{1}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Validator.Unmarshal(m, b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return xxx_messageInfo_Validator.Size(m)
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Validator) GetPubKey() keys.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return keys.PublicKey{}
}

func (m *Validator) GetVotingPower() int64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func (m *Validator) GetProposerPriority() int64 {
	if m != nil {
		return m.ProposerPriority
	}
	return 0
}

func init() {
	proto.RegisterType((*ValidatorSet)(nil), "tendermint.proto.types.ValidatorSet")
	proto.RegisterType((*Validator)(nil), "tendermint.proto.types.Validator")
}

func init() { proto.RegisterFile("proto/types/validator.proto", fileDescriptor_2e7c6b38c20e5406) }

var fileDescriptor_2e7c6b38c20e5406 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbd, 0x4e, 0x02, 0x41,
	0x10, 0xc7, 0x5d, 0x21, 0xa0, 0x0b, 0x85, 0x6e, 0x61, 0x2e, 0x18, 0x23, 0x50, 0x28, 0x89, 0x64,
	0x2f, 0xd1, 0xda, 0x42, 0x0a, 0x1b, 0x1a, 0x72, 0x26, 0x14, 0x36, 0x97, 0x3b, 0x6e, 0x73, 0x6c,
	0xf8, 0x98, 0xcd, 0xdc, 0x1c, 0x66, 0x1f, 0x4e, 0x6b, 0x9f, 0xc2, 0x67, 0x31, 0xdc, 0x72, 0x07,
	0x89, 0x14, 0x76, 0x33, 0xff, 0xff, 0x7c, 0xfc, 0x76, 0xb2, 0xfc, 0xda, 0x20, 0x10, 0xf8, 0x64,
	0x8d, 0xca, 0xfc, 0x4d, 0xb4, 0xd4, 0x49, 0x44, 0x80, 0xb2, 0x50, 0xc5, 0x15, 0xa9, 0x75, 0xa2,
	0x70, 0xa5, 0xd7, 0xe4, 0x14, 0x59, 0xd4, 0x75, 0xee, 0x68, 0xae, 0x31, 0x09, 0x4d, 0x84, 0x64,
	0x7d, 0x37, 0x20, 0x85, 0x14, 0xf6, 0x91, 0xab, 0xee, 0xdc, 0x38, 0x65, 0x86, 0xd6, 0x10, 0xf8,
	0x0b, 0x65, 0x33, 0xb7, 0xc8, 0xd9, 0xfd, 0x2f, 0xc6, 0xdb, 0xd3, 0x72, 0xe5, 0x9b, 0x22, 0xf1,
	0xc2, 0x79, 0x85, 0x90, 0x79, 0xac, 0x5b, 0x1b, 0xb4, 0x1e, 0x7b, 0xf2, 0x38, 0x84, 0xac, 0x3a,
	0x83, 0x83, 0x26, 0xf1, 0xcc, 0xcf, 0x0c, 0x82, 0x81, 0x4c, 0xa1, 0x77, 0xda, 0x65, 0xff, 0x1b,
	0x50, 0xb5, 0x88, 0x21, 0x17, 0x04, 0x14, 0x2d, 0xc3, 0x0d, 0x90, 0x5e, 0xa7, 0xa1, 0x81, 0x0f,
	0x85, 0x5e, 0xad, 0xcb, 0x06, 0xb5, 0xe0, 0xa2, 0x70, 0xa6, 0x85, 0x31, 0xd9, 0xea, 0xfd, 0x4f,
	0xc6, 0xcf, 0xab, 0x29, 0xc2, 0xe3, 0xcd, 0x28, 0x49, 0x50, 0x65, 0x5b, 0x74, 0x36, 0x68, 0x07,
	0x65, 0x2a, 0x5e, 0x79, 0xd3, 0xe4, 0x71, 0xb8, 0x50, 0x76, 0xc7, 0x74, 0xff, 0x97, 0xc9, 0x1d,
	0x49, 0x6e, 0x8f, 0x24, 0x27, 0x79, 0xbc, 0xd4, 0xb3, 0xb1, 0xb2, 0xa3, 0xfa, 0xf7, 0xcf, 0xed,
	0x49, 0xd0, 0x30, 0x79, 0x3c, 0x56, 0x56, 0xf4, 0x78, 0xfb, 0x08, 0x57, 0x6b, 0xb3, 0x47, 0x12,
	0x0f, 0xfc, 0xb2, 0x7c, 0x4c, 0x68, 0x50, 0x03, 0x6a, 0xb2, 0x5e, 0xdd, 0xf1, 0x97, 0xc6, 0x64,
	0xa7, 0x8f, 0xe4, 0xfb, 0x30, 0xd5, 0x34, 0xcf, 0x63, 0x39, 0x83, 0x95, 0xbf, 0x47, 0x3a, 0x0c,
	0x0f, 0xfe, 0x47, 0xdc, 0x28, 0x92, 0xa7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x28, 0xeb,
	0x40, 0x35, 0x02, 0x00, 0x00,
}
