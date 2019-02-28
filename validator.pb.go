// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validator.proto

package validator

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FieldValidator struct {
	// Uses a Golang RE2-syntax regex to match the field contents.
	Regex *string `protobuf:"bytes,1,opt,name=regex" json:"regex,omitempty"`
	// Field value should not match regex. Uses a Golang RE2-syntax regex.
	NotRegex *string `protobuf:"bytes,2,opt,name=not_regex,json=notRegex" json:"not_regex,omitempty"`
	// Field value of integer strictly greater than this value.
	IntGt *int64 `protobuf:"varint,3,opt,name=int_gt,json=intGt" json:"int_gt,omitempty"`
	// Field value of integer strictly smaller than this value.
	IntLt *int64 `protobuf:"varint,4,opt,name=int_lt,json=intLt" json:"int_lt,omitempty"`
	// Used for nested message types, requires that the message type exists.
	MsgExists *bool `protobuf:"varint,5,opt,name=msg_exists,json=msgExists" json:"msg_exists,omitempty"`
	// Human error specifies a user-customizable error that is visible to the user.
	HumanError *string `protobuf:"bytes,6,opt,name=human_error,json=humanError" json:"human_error,omitempty"`
	// Field value of double strictly greater than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatGt *float64 `protobuf:"fixed64,7,opt,name=float_gt,json=floatGt" json:"float_gt,omitempty"`
	// Field value of double strictly smaller than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatLt *float64 `protobuf:"fixed64,8,opt,name=float_lt,json=floatLt" json:"float_lt,omitempty"`
	// Field value of double describing the epsilon within which
	// any comparison should be considered to be true. For example,
	// when using float_gt = 0.35, using a float_epsilon of 0.05
	// would mean that any value above 0.30 is acceptable. It can be
	// thought of as a {float_value_condition} +- {float_epsilon}.
	// If unset, no correction for floating point inaccuracies in
	// comparisons will be attempted.
	FloatEpsilon *float64 `protobuf:"fixed64,9,opt,name=float_epsilon,json=floatEpsilon" json:"float_epsilon,omitempty"`
	// Floating-point value compared to which the field content should be greater or equal.
	FloatGte *float64 `protobuf:"fixed64,10,opt,name=float_gte,json=floatGte" json:"float_gte,omitempty"`
	// Floating-point value compared to which the field content should be smaller or equal.
	FloatLte *float64 `protobuf:"fixed64,11,opt,name=float_lte,json=floatLte" json:"float_lte,omitempty"`
	// Used for string fields, requires the string to be not empty (i.e different from "").
	StringNotEmpty *bool `protobuf:"varint,12,opt,name=string_not_empty,json=stringNotEmpty" json:"string_not_empty,omitempty"`
	// Repeated field with at least this number of elements.
	RepeatedCountMin *int64 `protobuf:"varint,13,opt,name=repeated_count_min,json=repeatedCountMin" json:"repeated_count_min,omitempty"`
	// Repeated field with at most this number of elements.
	RepeatedCountMax *int64 `protobuf:"varint,14,opt,name=repeated_count_max,json=repeatedCountMax" json:"repeated_count_max,omitempty"`
	// Field value of length greater than this value.
	LengthGt *int64 `protobuf:"varint,15,opt,name=length_gt,json=lengthGt" json:"length_gt,omitempty"`
	// Field value of length smaller than this value.
	LengthLt *int64 `protobuf:"varint,16,opt,name=length_lt,json=lengthLt" json:"length_lt,omitempty"`
	// Field value of integer strictly equal this value.
	LengthEq *int64 `protobuf:"varint,17,opt,name=length_eq,json=lengthEq" json:"length_eq,omitempty"`
	// Requires that the value is in the enum.
	IsInEnum             *bool    `protobuf:"varint,18,opt,name=is_in_enum,json=isInEnum" json:"is_in_enum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldValidator) Reset()         { *m = FieldValidator{} }
func (m *FieldValidator) String() string { return proto.CompactTextString(m) }
func (*FieldValidator) ProtoMessage()    {}
func (*FieldValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_validator_09a3c791d54e1d35, []int{0}
}
func (m *FieldValidator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldValidator.Unmarshal(m, b)
}
func (m *FieldValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldValidator.Marshal(b, m, deterministic)
}
func (dst *FieldValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldValidator.Merge(dst, src)
}
func (m *FieldValidator) XXX_Size() int {
	return xxx_messageInfo_FieldValidator.Size(m)
}
func (m *FieldValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldValidator.DiscardUnknown(m)
}

var xxx_messageInfo_FieldValidator proto.InternalMessageInfo

func (m *FieldValidator) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *FieldValidator) GetNotRegex() string {
	if m != nil && m.NotRegex != nil {
		return *m.NotRegex
	}
	return ""
}

func (m *FieldValidator) GetIntGt() int64 {
	if m != nil && m.IntGt != nil {
		return *m.IntGt
	}
	return 0
}

func (m *FieldValidator) GetIntLt() int64 {
	if m != nil && m.IntLt != nil {
		return *m.IntLt
	}
	return 0
}

func (m *FieldValidator) GetMsgExists() bool {
	if m != nil && m.MsgExists != nil {
		return *m.MsgExists
	}
	return false
}

func (m *FieldValidator) GetHumanError() string {
	if m != nil && m.HumanError != nil {
		return *m.HumanError
	}
	return ""
}

func (m *FieldValidator) GetFloatGt() float64 {
	if m != nil && m.FloatGt != nil {
		return *m.FloatGt
	}
	return 0
}

func (m *FieldValidator) GetFloatLt() float64 {
	if m != nil && m.FloatLt != nil {
		return *m.FloatLt
	}
	return 0
}

func (m *FieldValidator) GetFloatEpsilon() float64 {
	if m != nil && m.FloatEpsilon != nil {
		return *m.FloatEpsilon
	}
	return 0
}

func (m *FieldValidator) GetFloatGte() float64 {
	if m != nil && m.FloatGte != nil {
		return *m.FloatGte
	}
	return 0
}

func (m *FieldValidator) GetFloatLte() float64 {
	if m != nil && m.FloatLte != nil {
		return *m.FloatLte
	}
	return 0
}

func (m *FieldValidator) GetStringNotEmpty() bool {
	if m != nil && m.StringNotEmpty != nil {
		return *m.StringNotEmpty
	}
	return false
}

func (m *FieldValidator) GetRepeatedCountMin() int64 {
	if m != nil && m.RepeatedCountMin != nil {
		return *m.RepeatedCountMin
	}
	return 0
}

func (m *FieldValidator) GetRepeatedCountMax() int64 {
	if m != nil && m.RepeatedCountMax != nil {
		return *m.RepeatedCountMax
	}
	return 0
}

func (m *FieldValidator) GetLengthGt() int64 {
	if m != nil && m.LengthGt != nil {
		return *m.LengthGt
	}
	return 0
}

func (m *FieldValidator) GetLengthLt() int64 {
	if m != nil && m.LengthLt != nil {
		return *m.LengthLt
	}
	return 0
}

func (m *FieldValidator) GetLengthEq() int64 {
	if m != nil && m.LengthEq != nil {
		return *m.LengthEq
	}
	return 0
}

func (m *FieldValidator) GetIsInEnum() bool {
	if m != nil && m.IsInEnum != nil {
		return *m.IsInEnum
	}
	return false
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*FieldValidator)(nil),
	Field:         65020,
	Name:          "validator.field",
	Tag:           "bytes,65020,opt,name=field",
	Filename:      "validator.proto",
}

func init() {
	proto.RegisterType((*FieldValidator)(nil), "validator.FieldValidator")
	proto.RegisterExtension(E_Field)
}

func init() { proto.RegisterFile("validator.proto", fileDescriptor_validator_09a3c791d54e1d35) }

var fileDescriptor_validator_09a3c791d54e1d35 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x86, 0x15, 0xda, 0x6d, 0x93, 0xd9, 0x76, 0xbb, 0x58, 0x20, 0xb9, 0x40, 0x45, 0x04, 0x97,
	0x1c, 0x50, 0x2a, 0x71, 0xe4, 0x08, 0x0a, 0x2b, 0xa4, 0xe5, 0x43, 0x39, 0x70, 0xe0, 0x62, 0x85,
	0xee, 0x6c, 0x6a, 0xc9, 0xb1, 0x53, 0x7b, 0x82, 0x96, 0x9f, 0xc7, 0xef, 0x82, 0x03, 0xca, 0x84,
	0xec, 0x06, 0xa9, 0x47, 0x3f, 0xcf, 0x1b, 0x5b, 0x33, 0x79, 0xe1, 0xe2, 0x47, 0x65, 0xf4, 0xa6,
	0x22, 0xe7, 0xf3, 0xd6, 0x3b, 0x72, 0x22, 0xd9, 0x83, 0x27, 0x69, 0xed, 0x5c, 0x6d, 0xf0, 0x9a,
	0xc5, 0xf7, 0x6e, 0x7b, 0xbd, 0xc1, 0x70, 0xe3, 0x75, 0xbb, 0x0f, 0xbf, 0xf8, 0x75, 0x0c, 0x8b,
	0xf7, 0x1a, 0xcd, 0xe6, 0xeb, 0xf8, 0x91, 0x78, 0x04, 0x33, 0x8f, 0x35, 0xee, 0x64, 0x94, 0x46,
	0x59, 0x52, 0x0e, 0x07, 0xf1, 0x14, 0x12, 0xeb, 0x48, 0x0d, 0xe6, 0x01, 0x9b, 0xd8, 0x3a, 0x2a,
	0x59, 0x3e, 0x86, 0x13, 0x6d, 0x49, 0xd5, 0x24, 0x8f, 0xd2, 0x28, 0x3b, 0x2a, 0x67, 0xda, 0xd2,
	0x8a, 0x46, 0x6c, 0x48, 0x1e, 0xef, 0xf1, 0x9a, 0xc4, 0x15, 0x40, 0x13, 0x6a, 0x85, 0x3b, 0x1d,
	0x28, 0xc8, 0x59, 0x1a, 0x65, 0x71, 0x99, 0x34, 0xa1, 0x2e, 0x18, 0x88, 0xe7, 0x30, 0xbf, 0xed,
	0x9a, 0xca, 0x2a, 0xf4, 0xde, 0x79, 0x79, 0xc2, 0x6f, 0x01, 0xa3, 0xa2, 0x27, 0xe2, 0x12, 0xe2,
	0xad, 0x71, 0x15, 0xbf, 0x77, 0x9a, 0x46, 0x59, 0x54, 0x9e, 0xf2, 0x79, 0x45, 0x07, 0x65, 0x48,
	0xc6, 0x13, 0xb5, 0x26, 0xf1, 0x12, 0xce, 0x07, 0x85, 0x6d, 0xd0, 0xc6, 0x59, 0x99, 0xb0, 0x3f,
	0x63, 0x58, 0x0c, 0xac, 0x9f, 0x72, 0xbc, 0x1a, 0x25, 0x70, 0x20, 0xfe, 0x77, 0x37, 0x1e, 0xa4,
	0x21, 0x94, 0xf3, 0x89, 0x5c, 0x13, 0x8a, 0x0c, 0x96, 0x81, 0xbc, 0xb6, 0xb5, 0xea, 0xd7, 0x84,
	0x4d, 0x4b, 0x3f, 0xe5, 0x19, 0x8f, 0xb6, 0x18, 0xf8, 0x27, 0x47, 0x45, 0x4f, 0xc5, 0x2b, 0x10,
	0x1e, 0x5b, 0xac, 0x08, 0x37, 0xea, 0xc6, 0x75, 0x96, 0x54, 0xa3, 0xad, 0x3c, 0xe7, 0x0d, 0x2d,
	0x47, 0xf3, 0xae, 0x17, 0x1f, 0xb5, 0xbd, 0x2f, 0x5d, 0xed, 0xe4, 0xe2, 0xbe, 0x74, 0xc5, 0x7f,
	0xc9, 0xa0, 0xad, 0xe9, 0xb6, 0xdf, 0xcd, 0x05, 0x87, 0xe2, 0x01, 0xac, 0x68, 0x22, 0x0d, 0xc9,
	0xe5, 0x54, 0xae, 0xa7, 0x12, 0xef, 0xe4, 0xc3, 0xa9, 0x2c, 0xee, 0xc4, 0x33, 0x00, 0x1d, 0x94,
	0xb6, 0x0a, 0x6d, 0xd7, 0x48, 0xc1, 0x63, 0xc5, 0x3a, 0x7c, 0xb0, 0x85, 0xed, 0x9a, 0x37, 0x5f,
	0x60, 0xb6, 0xed, 0x2b, 0x24, 0xae, 0xf2, 0xa1, 0x6f, 0xf9, 0xd8, 0xb7, 0x9c, 0xab, 0xf5, 0xb9,
	0x25, 0xed, 0x6c, 0x90, 0x7f, 0x7e, 0xf7, 0xed, 0x98, 0xbf, 0xbe, 0xcc, 0x0f, 0x95, 0xfd, 0xbf,
	0x7b, 0xe5, 0x70, 0xd1, 0xdb, 0xf9, 0xb7, 0x43, 0x89, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x4c,
	0xc1, 0xa1, 0xc7, 0xe1, 0x02, 0x00, 0x00,
}
