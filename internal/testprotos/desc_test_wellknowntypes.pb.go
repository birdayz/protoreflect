// Code generated by protoc-gen-go. DO NOT EDIT.
// source: desc_test_wellknowntypes.proto

package testprotos // import "github.com/jhump/protoreflect/internal/testprotos"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import duration "github.com/golang/protobuf/ptypes/duration"
import _struct "github.com/golang/protobuf/ptypes/struct"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TestWellKnownTypes struct {
	StartTime            *timestamp.Timestamp  `protobuf:"bytes,1,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	Elapsed              *duration.Duration    `protobuf:"bytes,2,opt,name=elapsed" json:"elapsed,omitempty"`
	Dbl                  *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=dbl" json:"dbl,omitempty"`
	Flt                  *wrappers.FloatValue  `protobuf:"bytes,4,opt,name=flt" json:"flt,omitempty"`
	Bl                   *wrappers.BoolValue   `protobuf:"bytes,5,opt,name=bl" json:"bl,omitempty"`
	I32                  *wrappers.Int32Value  `protobuf:"bytes,6,opt,name=i32" json:"i32,omitempty"`
	I64                  *wrappers.Int64Value  `protobuf:"bytes,7,opt,name=i64" json:"i64,omitempty"`
	U32                  *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=u32" json:"u32,omitempty"`
	U64                  *wrappers.UInt64Value `protobuf:"bytes,9,opt,name=u64" json:"u64,omitempty"`
	Str                  *wrappers.StringValue `protobuf:"bytes,10,opt,name=str" json:"str,omitempty"`
	Byt                  *wrappers.BytesValue  `protobuf:"bytes,11,opt,name=byt" json:"byt,omitempty"`
	Json                 []*_struct.Value      `protobuf:"bytes,12,rep,name=json" json:"json,omitempty"`
	Extras               []*any.Any            `protobuf:"bytes,13,rep,name=extras" json:"extras,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TestWellKnownTypes) Reset()         { *m = TestWellKnownTypes{} }
func (m *TestWellKnownTypes) String() string { return proto.CompactTextString(m) }
func (*TestWellKnownTypes) ProtoMessage()    {}
func (*TestWellKnownTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_desc_test_wellknowntypes_958e16afb8e92d7c, []int{0}
}
func (m *TestWellKnownTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestWellKnownTypes.Unmarshal(m, b)
}
func (m *TestWellKnownTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestWellKnownTypes.Marshal(b, m, deterministic)
}
func (dst *TestWellKnownTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestWellKnownTypes.Merge(dst, src)
}
func (m *TestWellKnownTypes) XXX_Size() int {
	return xxx_messageInfo_TestWellKnownTypes.Size(m)
}
func (m *TestWellKnownTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_TestWellKnownTypes.DiscardUnknown(m)
}

var xxx_messageInfo_TestWellKnownTypes proto.InternalMessageInfo

func (m *TestWellKnownTypes) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TestWellKnownTypes) GetElapsed() *duration.Duration {
	if m != nil {
		return m.Elapsed
	}
	return nil
}

func (m *TestWellKnownTypes) GetDbl() *wrappers.DoubleValue {
	if m != nil {
		return m.Dbl
	}
	return nil
}

func (m *TestWellKnownTypes) GetFlt() *wrappers.FloatValue {
	if m != nil {
		return m.Flt
	}
	return nil
}

func (m *TestWellKnownTypes) GetBl() *wrappers.BoolValue {
	if m != nil {
		return m.Bl
	}
	return nil
}

func (m *TestWellKnownTypes) GetI32() *wrappers.Int32Value {
	if m != nil {
		return m.I32
	}
	return nil
}

func (m *TestWellKnownTypes) GetI64() *wrappers.Int64Value {
	if m != nil {
		return m.I64
	}
	return nil
}

func (m *TestWellKnownTypes) GetU32() *wrappers.UInt32Value {
	if m != nil {
		return m.U32
	}
	return nil
}

func (m *TestWellKnownTypes) GetU64() *wrappers.UInt64Value {
	if m != nil {
		return m.U64
	}
	return nil
}

func (m *TestWellKnownTypes) GetStr() *wrappers.StringValue {
	if m != nil {
		return m.Str
	}
	return nil
}

func (m *TestWellKnownTypes) GetByt() *wrappers.BytesValue {
	if m != nil {
		return m.Byt
	}
	return nil
}

func (m *TestWellKnownTypes) GetJson() []*_struct.Value {
	if m != nil {
		return m.Json
	}
	return nil
}

func (m *TestWellKnownTypes) GetExtras() []*any.Any {
	if m != nil {
		return m.Extras
	}
	return nil
}

func init() {
	proto.RegisterType((*TestWellKnownTypes)(nil), "testprotos.TestWellKnownTypes")
}

func init() {
	proto.RegisterFile("desc_test_wellknowntypes.proto", fileDescriptor_desc_test_wellknowntypes_958e16afb8e92d7c)
}

var fileDescriptor_desc_test_wellknowntypes_958e16afb8e92d7c = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd3, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0x07, 0x70, 0xae, 0xbd, 0xeb, 0x79, 0x5b, 0x7d, 0x09, 0x22, 0x7b, 0xf5, 0x38, 0x0f, 0x9f,
	0x0e, 0xd1, 0x04, 0x93, 0x52, 0xf0, 0xd1, 0x22, 0x82, 0xf8, 0x56, 0xab, 0x82, 0x2f, 0x65, 0xb7,
	0x9d, 0xf6, 0x72, 0x4e, 0x77, 0xc3, 0xee, 0x2c, 0x35, 0x5f, 0xc0, 0xcf, 0x2d, 0xbb, 0xd9, 0xa8,
	0x34, 0xe4, 0x9e, 0x92, 0x65, 0x7e, 0xf3, 0xcf, 0x4c, 0x48, 0xd8, 0xf5, 0x06, 0xec, 0x7a, 0x45,
	0x60, 0x69, 0x75, 0x00, 0xc4, 0x9f, 0x4a, 0x1f, 0x14, 0xd5, 0x15, 0xd8, 0xb4, 0x32, 0x9a, 0x74,
	0xc2, 0x7c, 0x29, 0xdc, 0xda, 0xc9, 0xe5, 0x4e, 0xeb, 0x1d, 0x42, 0x16, 0x8e, 0xd2, 0x6d, 0x33,
	0xa1, 0xea, 0x86, 0x4d, 0xae, 0x8f, 0x4b, 0x1b, 0x67, 0x04, 0x95, 0x5a, 0xc5, 0xfa, 0x8b, 0xe3,
	0x3a, 0x95, 0x7b, 0xb0, 0x24, 0xf6, 0x55, 0x04, 0x57, 0xc7, 0xc0, 0x92, 0x71, 0x6b, 0xea, 0x8b,
	0x3f, 0x18, 0x51, 0x55, 0x60, 0xe2, 0x94, 0x2f, 0x7f, 0x9f, 0xb1, 0x64, 0x09, 0x96, 0xbe, 0x03,
	0xe2, 0x67, 0xbf, 0xc2, 0xd2, 0xaf, 0x90, 0xbc, 0x63, 0xcc, 0x92, 0x30, 0xb4, 0xf2, 0x4f, 0xe3,
	0x27, 0x37, 0x27, 0xb7, 0xe3, 0x7c, 0x92, 0x36, 0x59, 0x69, 0x9b, 0x95, 0x2e, 0xdb, 0x51, 0x16,
	0x17, 0x41, 0xfb, 0x73, 0x52, 0xb0, 0x73, 0x40, 0x51, 0x59, 0xd8, 0xf0, 0x41, 0xe8, 0xbb, 0xec,
	0xf4, 0x7d, 0x88, 0x2b, 0x2e, 0x5a, 0x99, 0xa4, 0x6c, 0xb8, 0x91, 0xc8, 0x87, 0xa1, 0xe1, 0xaa,
	0xdb, 0xa0, 0x9d, 0x44, 0xf8, 0x26, 0xd0, 0xc1, 0xc2, 0xc3, 0xe4, 0x0d, 0x1b, 0x6e, 0x91, 0xf8,
	0x69, 0xf0, 0xcf, 0x3b, 0xfe, 0x23, 0x6a, 0x41, 0x91, 0x6f, 0x91, 0x92, 0x57, 0x6c, 0x20, 0x91,
	0x9f, 0xf5, 0xac, 0x31, 0xd7, 0x1a, 0x1b, 0x3c, 0x68, 0xa2, 0xcb, 0x22, 0xe7, 0xa3, 0x9e, 0xe8,
	0x4f, 0x8a, 0x8a, 0x3c, 0x46, 0x97, 0x45, 0x1e, 0xf8, 0x6c, 0xca, 0xcf, 0xfb, 0xf9, 0x6c, 0xda,
	0xf2, 0xd9, 0xd4, 0x2f, 0xea, 0x8a, 0x9c, 0x3f, 0xea, 0x59, 0xf4, 0xeb, 0xff, 0xf1, 0xae, 0xc8,
	0x83, 0x9f, 0x4d, 0xf9, 0xc5, 0x03, 0xfe, 0x6f, 0xbe, 0x6b, 0xf2, 0x2d, 0x19, 0xce, 0x7a, 0xfc,
	0x17, 0x32, 0xa5, 0xda, 0x45, 0x6f, 0xc9, 0xf8, 0xf1, 0x65, 0x4d, 0x7c, 0xdc, 0x33, 0xfe, 0xbc,
	0x26, 0xb0, 0x91, 0xcb, 0xda, 0xbf, 0xc8, 0xd3, 0x7b, 0xab, 0x15, 0x7f, 0x7c, 0x33, 0xbc, 0x1d,
	0xe7, 0xcf, 0x3a, 0xbe, 0xa1, 0xc1, 0x24, 0xaf, 0xd9, 0x08, 0x7e, 0x91, 0x11, 0x96, 0x3f, 0x09,
	0xfa, 0x69, 0x47, 0xbf, 0x57, 0xf5, 0x22, 0x9a, 0x79, 0xf1, 0xe3, 0xed, 0xae, 0xa4, 0x3b, 0x27,
	0xd3, 0xb5, 0xde, 0x67, 0xf7, 0x77, 0x6e, 0x5f, 0x35, 0x1f, 0xad, 0x81, 0x2d, 0xc2, 0x9a, 0xb2,
	0x52, 0x11, 0x18, 0x25, 0x30, 0xfb, 0xf7, 0x5f, 0xc9, 0x51, 0xb8, 0x16, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xcb, 0xec, 0x5a, 0xd4, 0x8c, 0x03, 0x00, 0x00,
}
