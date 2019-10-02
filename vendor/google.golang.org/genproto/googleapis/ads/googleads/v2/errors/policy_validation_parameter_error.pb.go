// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/policy_validation_parameter_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible policy validation parameter errors.
type PolicyValidationParameterErrorEnum_PolicyValidationParameterError int32

const (
	// Enum unspecified.
	PolicyValidationParameterErrorEnum_UNSPECIFIED PolicyValidationParameterErrorEnum_PolicyValidationParameterError = 0
	// The received error code is not known in this version.
	PolicyValidationParameterErrorEnum_UNKNOWN PolicyValidationParameterErrorEnum_PolicyValidationParameterError = 1
	// Ignorable policy topics are not supported for the ad type.
	PolicyValidationParameterErrorEnum_UNSUPPORTED_AD_TYPE_FOR_IGNORABLE_POLICY_TOPICS PolicyValidationParameterErrorEnum_PolicyValidationParameterError = 2
	// Exempt policy violation keys are not supported for the ad type.
	PolicyValidationParameterErrorEnum_UNSUPPORTED_AD_TYPE_FOR_EXEMPT_POLICY_VIOLATION_KEYS PolicyValidationParameterErrorEnum_PolicyValidationParameterError = 3
	// Cannot set ignorable policy topics and exempt policy violation keys in
	// the same policy violation parameter.
	PolicyValidationParameterErrorEnum_CANNOT_SET_BOTH_IGNORABLE_POLICY_TOPICS_AND_EXEMPT_POLICY_VIOLATION_KEYS PolicyValidationParameterErrorEnum_PolicyValidationParameterError = 4
)

var PolicyValidationParameterErrorEnum_PolicyValidationParameterError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "UNSUPPORTED_AD_TYPE_FOR_IGNORABLE_POLICY_TOPICS",
	3: "UNSUPPORTED_AD_TYPE_FOR_EXEMPT_POLICY_VIOLATION_KEYS",
	4: "CANNOT_SET_BOTH_IGNORABLE_POLICY_TOPICS_AND_EXEMPT_POLICY_VIOLATION_KEYS",
}

var PolicyValidationParameterErrorEnum_PolicyValidationParameterError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"UNSUPPORTED_AD_TYPE_FOR_IGNORABLE_POLICY_TOPICS":                          2,
	"UNSUPPORTED_AD_TYPE_FOR_EXEMPT_POLICY_VIOLATION_KEYS":                     3,
	"CANNOT_SET_BOTH_IGNORABLE_POLICY_TOPICS_AND_EXEMPT_POLICY_VIOLATION_KEYS": 4,
}

func (x PolicyValidationParameterErrorEnum_PolicyValidationParameterError) String() string {
	return proto.EnumName(PolicyValidationParameterErrorEnum_PolicyValidationParameterError_name, int32(x))
}

func (PolicyValidationParameterErrorEnum_PolicyValidationParameterError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_672b335e6c2f0519, []int{0, 0}
}

// Container for enum describing possible policy validation parameter errors.
type PolicyValidationParameterErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyValidationParameterErrorEnum) Reset()         { *m = PolicyValidationParameterErrorEnum{} }
func (m *PolicyValidationParameterErrorEnum) String() string { return proto.CompactTextString(m) }
func (*PolicyValidationParameterErrorEnum) ProtoMessage()    {}
func (*PolicyValidationParameterErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_672b335e6c2f0519, []int{0}
}

func (m *PolicyValidationParameterErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyValidationParameterErrorEnum.Unmarshal(m, b)
}
func (m *PolicyValidationParameterErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyValidationParameterErrorEnum.Marshal(b, m, deterministic)
}
func (m *PolicyValidationParameterErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyValidationParameterErrorEnum.Merge(m, src)
}
func (m *PolicyValidationParameterErrorEnum) XXX_Size() int {
	return xxx_messageInfo_PolicyValidationParameterErrorEnum.Size(m)
}
func (m *PolicyValidationParameterErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyValidationParameterErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyValidationParameterErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.PolicyValidationParameterErrorEnum_PolicyValidationParameterError", PolicyValidationParameterErrorEnum_PolicyValidationParameterError_name, PolicyValidationParameterErrorEnum_PolicyValidationParameterError_value)
	proto.RegisterType((*PolicyValidationParameterErrorEnum)(nil), "google.ads.googleads.v2.errors.PolicyValidationParameterErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/policy_validation_parameter_error.proto", fileDescriptor_672b335e6c2f0519)
}

var fileDescriptor_672b335e6c2f0519 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x8a, 0xdb, 0x30,
	0x18, 0x85, 0x1b, 0x4f, 0x69, 0x41, 0xb3, 0x68, 0xf0, 0xb2, 0x94, 0x2c, 0xdc, 0xbd, 0x0c, 0x9e,
	0x2e, 0x8a, 0xba, 0x52, 0x62, 0x25, 0x63, 0x26, 0x95, 0x44, 0xac, 0xb8, 0x4d, 0x31, 0x08, 0x75,
	0x6c, 0x8c, 0xc1, 0xb1, 0x8c, 0xe5, 0x06, 0x7a, 0x9d, 0x59, 0xf6, 0x28, 0x3d, 0x4a, 0xaf, 0x50,
	0x28, 0xc5, 0xd6, 0xd8, 0xbb, 0x78, 0xa5, 0x87, 0x78, 0xef, 0x7b, 0xe2, 0xd7, 0x0f, 0xb6, 0x85,
	0xd6, 0x45, 0x95, 0xfb, 0x2a, 0x33, 0xbe, 0x95, 0xbd, 0xba, 0x04, 0x7e, 0xde, 0xb6, 0xba, 0x35,
	0x7e, 0xa3, 0xab, 0xf2, 0xf1, 0xa7, 0xbc, 0xa8, 0xaa, 0xcc, 0x54, 0x57, 0xea, 0x5a, 0x36, 0xaa,
	0x55, 0xe7, 0xbc, 0xcb, 0x5b, 0x39, 0x58, 0x60, 0xd3, 0xea, 0x4e, 0xbb, 0x2b, 0x1b, 0x86, 0x2a,
	0x33, 0x70, 0xe2, 0xc0, 0x4b, 0x00, 0x2d, 0xe7, 0xed, 0xbb, 0xb1, 0xa7, 0x29, 0x7d, 0x55, 0xd7,
	0xba, 0x1b, 0x60, 0xc6, 0xa6, 0xbd, 0x27, 0x07, 0x78, 0x7c, 0x68, 0x4a, 0xa6, 0x22, 0x3e, 0xf6,
	0x90, 0x9e, 0x40, 0xea, 0x1f, 0x67, 0xef, 0xef, 0x02, 0xac, 0xe6, 0x6d, 0xee, 0x1b, 0x70, 0x7b,
	0xa4, 0x31, 0x27, 0x9b, 0x68, 0x1b, 0x91, 0x70, 0xf9, 0xc2, 0xbd, 0x05, 0xaf, 0x8f, 0xf4, 0x81,
	0xb2, 0x2f, 0x74, 0xb9, 0x70, 0xef, 0x80, 0x7f, 0xa4, 0xf1, 0x91, 0x73, 0x76, 0x10, 0x24, 0x94,
	0x38, 0x94, 0xe2, 0xc4, 0x89, 0xdc, 0xb2, 0x83, 0x8c, 0x76, 0x94, 0x1d, 0xf0, 0x7a, 0x4f, 0x24,
	0x67, 0xfb, 0x68, 0x73, 0x92, 0x82, 0xf1, 0x68, 0x13, 0x2f, 0x1d, 0xf7, 0x23, 0xf8, 0x70, 0x2d,
	0x44, 0xbe, 0x92, 0xcf, 0x5c, 0x8c, 0x89, 0x24, 0x62, 0x7b, 0x2c, 0x22, 0x46, 0xe5, 0x03, 0x39,
	0xc5, 0xcb, 0x1b, 0x77, 0x0f, 0xee, 0x37, 0x98, 0x52, 0x26, 0x64, 0x4c, 0x84, 0x5c, 0x33, 0x71,
	0x7f, 0xad, 0x46, 0x62, 0x1a, 0xce, 0xd3, 0x5e, 0xae, 0xff, 0x2d, 0x80, 0xf7, 0xa8, 0xcf, 0x70,
	0x7e, 0xd2, 0xeb, 0xf7, 0xf3, 0x13, 0xe2, 0xfd, 0xc0, 0xf9, 0xe2, 0x5b, 0xf8, 0x8c, 0x29, 0x74,
	0xa5, 0xea, 0x02, 0xea, 0xb6, 0xf0, 0x8b, 0xbc, 0x1e, 0xbe, 0x63, 0x5c, 0x84, 0xa6, 0x34, 0xd7,
	0xf6, 0xe2, 0x93, 0x3d, 0x9e, 0x9c, 0x9b, 0x1d, 0xc6, 0xbf, 0x9c, 0xd5, 0xce, 0xc2, 0x70, 0x66,
	0xa0, 0x95, 0xbd, 0x4a, 0x02, 0x38, 0x54, 0x9a, 0xdf, 0xa3, 0x21, 0xc5, 0x99, 0x49, 0x27, 0x43,
	0x9a, 0x04, 0xa9, 0x35, 0xfc, 0x71, 0x3c, 0x7b, 0x8b, 0x10, 0xce, 0x0c, 0x42, 0x93, 0x05, 0xa1,
	0x24, 0x40, 0xc8, 0x9a, 0xbe, 0xbf, 0x1a, 0x5e, 0x77, 0xf7, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xd9,
	0x88, 0x75, 0x46, 0xb4, 0x02, 0x00, 0x00,
}
