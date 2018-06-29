// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/healthcheck.proto

package healthcheck // import "github.com/runconduit/conduit/controller/gen/common/healthcheck"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CheckStatus int32

const (
	CheckStatus_OK    CheckStatus = 0
	CheckStatus_FAIL  CheckStatus = 1
	CheckStatus_ERROR CheckStatus = 2
)

var CheckStatus_name = map[int32]string{
	0: "OK",
	1: "FAIL",
	2: "ERROR",
}
var CheckStatus_value = map[string]int32{
	"OK":    0,
	"FAIL":  1,
	"ERROR": 2,
}

func (x CheckStatus) String() string {
	return proto.EnumName(CheckStatus_name, int32(x))
}
func (CheckStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_healthcheck_de51aec14389c05b, []int{0}
}

type CheckResult struct {
	SubsystemName         string      `protobuf:"bytes,1,opt,name=SubsystemName" json:"SubsystemName,omitempty"`
	CheckDescription      string      `protobuf:"bytes,2,opt,name=CheckDescription" json:"CheckDescription,omitempty"`
	Status                CheckStatus `protobuf:"varint,3,opt,name=Status,enum=conduit.common.healthcheck.CheckStatus" json:"Status,omitempty"`
	FriendlyMessageToUser string      `protobuf:"bytes,4,opt,name=FriendlyMessageToUser" json:"FriendlyMessageToUser,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}    `json:"-"`
	XXX_unrecognized      []byte      `json:"-"`
	XXX_sizecache         int32       `json:"-"`
}

func (m *CheckResult) Reset()         { *m = CheckResult{} }
func (m *CheckResult) String() string { return proto.CompactTextString(m) }
func (*CheckResult) ProtoMessage()    {}
func (*CheckResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_healthcheck_de51aec14389c05b, []int{0}
}
func (m *CheckResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResult.Unmarshal(m, b)
}
func (m *CheckResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResult.Marshal(b, m, deterministic)
}
func (dst *CheckResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResult.Merge(dst, src)
}
func (m *CheckResult) XXX_Size() int {
	return xxx_messageInfo_CheckResult.Size(m)
}
func (m *CheckResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResult.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResult proto.InternalMessageInfo

func (m *CheckResult) GetSubsystemName() string {
	if m != nil {
		return m.SubsystemName
	}
	return ""
}

func (m *CheckResult) GetCheckDescription() string {
	if m != nil {
		return m.CheckDescription
	}
	return ""
}

func (m *CheckResult) GetStatus() CheckStatus {
	if m != nil {
		return m.Status
	}
	return CheckStatus_OK
}

func (m *CheckResult) GetFriendlyMessageToUser() string {
	if m != nil {
		return m.FriendlyMessageToUser
	}
	return ""
}

type SelfCheckRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelfCheckRequest) Reset()         { *m = SelfCheckRequest{} }
func (m *SelfCheckRequest) String() string { return proto.CompactTextString(m) }
func (*SelfCheckRequest) ProtoMessage()    {}
func (*SelfCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_healthcheck_de51aec14389c05b, []int{1}
}
func (m *SelfCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelfCheckRequest.Unmarshal(m, b)
}
func (m *SelfCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelfCheckRequest.Marshal(b, m, deterministic)
}
func (dst *SelfCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelfCheckRequest.Merge(dst, src)
}
func (m *SelfCheckRequest) XXX_Size() int {
	return xxx_messageInfo_SelfCheckRequest.Size(m)
}
func (m *SelfCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SelfCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SelfCheckRequest proto.InternalMessageInfo

type SelfCheckResponse struct {
	Results              []*CheckResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SelfCheckResponse) Reset()         { *m = SelfCheckResponse{} }
func (m *SelfCheckResponse) String() string { return proto.CompactTextString(m) }
func (*SelfCheckResponse) ProtoMessage()    {}
func (*SelfCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_healthcheck_de51aec14389c05b, []int{2}
}
func (m *SelfCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelfCheckResponse.Unmarshal(m, b)
}
func (m *SelfCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelfCheckResponse.Marshal(b, m, deterministic)
}
func (dst *SelfCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelfCheckResponse.Merge(dst, src)
}
func (m *SelfCheckResponse) XXX_Size() int {
	return xxx_messageInfo_SelfCheckResponse.Size(m)
}
func (m *SelfCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SelfCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SelfCheckResponse proto.InternalMessageInfo

func (m *SelfCheckResponse) GetResults() []*CheckResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckResult)(nil), "conduit.common.healthcheck.CheckResult")
	proto.RegisterType((*SelfCheckRequest)(nil), "conduit.common.healthcheck.SelfCheckRequest")
	proto.RegisterType((*SelfCheckResponse)(nil), "conduit.common.healthcheck.SelfCheckResponse")
	proto.RegisterEnum("conduit.common.healthcheck.CheckStatus", CheckStatus_name, CheckStatus_value)
}

func init() {
	proto.RegisterFile("common/healthcheck.proto", fileDescriptor_healthcheck_de51aec14389c05b)
}

var fileDescriptor_healthcheck_de51aec14389c05b = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4d, 0x4b, 0x3b, 0x31,
	0x10, 0xc6, 0xff, 0x69, 0xfb, 0xaf, 0x76, 0x8a, 0xb2, 0x06, 0x84, 0xe0, 0xa9, 0x14, 0xc1, 0xd2,
	0xc3, 0x2e, 0xa8, 0xf7, 0x52, 0x5f, 0x0a, 0xe2, 0x4b, 0x21, 0x55, 0x0f, 0xde, 0xb6, 0xdb, 0xb1,
	0xbb, 0x98, 0x4d, 0x6a, 0x26, 0x39, 0xf4, 0x8b, 0xfa, 0x79, 0xa4, 0xe9, 0x0a, 0x95, 0x2a, 0x78,
	0x4a, 0x78, 0xf2, 0xfc, 0x32, 0xcf, 0xcc, 0x80, 0xc8, 0x4c, 0x59, 0x1a, 0x9d, 0xe4, 0x98, 0x2a,
	0x97, 0x67, 0x39, 0x66, 0x6f, 0xf1, 0xc2, 0x1a, 0x67, 0xf8, 0x51, 0x66, 0xf4, 0xcc, 0x17, 0x2e,
	0x5e, 0x3b, 0xe2, 0x0d, 0x47, 0xf7, 0x83, 0x41, 0xfb, 0x72, 0x75, 0x93, 0x48, 0x5e, 0x39, 0x7e,
	0x0c, 0x7b, 0x13, 0x3f, 0xa5, 0x25, 0x39, 0x2c, 0x1f, 0xd2, 0x12, 0x05, 0xeb, 0xb0, 0x5e, 0x4b,
	0x7e, 0x17, 0x79, 0x1f, 0xa2, 0x00, 0x5d, 0x21, 0x65, 0xb6, 0x58, 0xb8, 0xc2, 0x68, 0x51, 0x0b,
	0xc6, 0x2d, 0x9d, 0x0f, 0xa0, 0x39, 0x71, 0xa9, 0xf3, 0x24, 0xea, 0x1d, 0xd6, 0xdb, 0x3f, 0x3d,
	0x89, 0x7f, 0x8f, 0x13, 0x07, 0x7a, 0x6d, 0x97, 0x15, 0xc6, 0xcf, 0xe1, 0x70, 0x64, 0x0b, 0xd4,
	0x33, 0xb5, 0xbc, 0x47, 0xa2, 0x74, 0x8e, 0x8f, 0xe6, 0x89, 0xd0, 0x8a, 0x46, 0xa8, 0xf8, 0xf3,
	0x63, 0x97, 0x43, 0x34, 0x41, 0xf5, 0x5a, 0xf5, 0xf6, 0xee, 0x91, 0x5c, 0xf7, 0x19, 0x0e, 0x36,
	0x34, 0x5a, 0x18, 0x4d, 0xc8, 0x87, 0xb0, 0x63, 0x43, 0xef, 0x24, 0x58, 0xa7, 0xde, 0x6b, 0xff,
	0x21, 0xe0, 0x7a, 0x56, 0xf2, 0x8b, 0xeb, 0xf7, 0xab, 0x19, 0x56, 0x81, 0x9b, 0x50, 0x1b, 0xdf,
	0x46, 0xff, 0xf8, 0x2e, 0x34, 0x46, 0xc3, 0x9b, 0xbb, 0x88, 0xf1, 0x16, 0xfc, 0xbf, 0x96, 0x72,
	0x2c, 0xa3, 0xda, 0xc5, 0xf0, 0x65, 0x30, 0x2f, 0x5c, 0xee, 0xa7, 0xab, 0xdf, 0x13, 0xeb, 0x75,
	0x55, 0x2c, 0xd9, 0x38, 0x9d, 0x35, 0x4a, 0xa1, 0x4d, 0xe6, 0xa8, 0x93, 0xed, 0xad, 0x4e, 0x9b,
	0x61, 0xad, 0x67, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x93, 0x2a, 0xb7, 0x43, 0xf2, 0x01, 0x00,
	0x00,
}
