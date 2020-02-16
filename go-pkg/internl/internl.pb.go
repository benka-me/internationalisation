// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internl.proto

package internl

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
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

type Languages int32

const (
	EN Languages = 0
	FR Languages = 1
)

var Languages_name = map[int32]string{
	0: "EN",
	1: "FR",
}

var Languages_value = map[string]int32{
	"EN": 0,
	"FR": 1,
}

func (Languages) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bfe439df8e494cc1, []int{0}
}

type MessageRequest struct {
	ServiceRequester string    `protobuf:"bytes,1,opt,name=ServiceRequester,proto3" json:"ServiceRequester,omitempty"`
	Lang             Languages `protobuf:"varint,2,opt,name=Lang,proto3,enum=internl.Languages" json:"Lang,omitempty"`
	Code             uint32    `protobuf:"varint,3,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (m *MessageRequest) Reset()      { *m = MessageRequest{} }
func (*MessageRequest) ProtoMessage() {}
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe439df8e494cc1, []int{0}
}
func (m *MessageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MessageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRequest.Merge(m, src)
}
func (m *MessageRequest) XXX_Size() int {
	return m.Size()
}
func (m *MessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRequest proto.InternalMessageInfo

func (m *MessageRequest) GetServiceRequester() string {
	if m != nil {
		return m.ServiceRequester
	}
	return ""
}

func (m *MessageRequest) GetLang() Languages {
	if m != nil {
		return m.Lang
	}
	return EN
}

func (m *MessageRequest) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type MessageResponse struct {
	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *MessageResponse) Reset()      { *m = MessageResponse{} }
func (*MessageResponse) ProtoMessage() {}
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe439df8e494cc1, []int{1}
}
func (m *MessageResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MessageResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageResponse.Merge(m, src)
}
func (m *MessageResponse) XXX_Size() int {
	return m.Size()
}
func (m *MessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageResponse proto.InternalMessageInfo

func (m *MessageResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type PushMessage struct {
	Value string    `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	Lang  Languages `protobuf:"varint,2,opt,name=Lang,proto3,enum=internl.Languages" json:"Lang,omitempty"`
	Code  uint32    `protobuf:"varint,3,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (m *PushMessage) Reset()      { *m = PushMessage{} }
func (*PushMessage) ProtoMessage() {}
func (*PushMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe439df8e494cc1, []int{2}
}
func (m *PushMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMessage.Merge(m, src)
}
func (m *PushMessage) XXX_Size() int {
	return m.Size()
}
func (m *PushMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PushMessage proto.InternalMessageInfo

func (m *PushMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PushMessage) GetLang() Languages {
	if m != nil {
		return m.Lang
	}
	return EN
}

func (m *PushMessage) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type PushMessageReq struct {
	PushMessage *PushMessage `protobuf:"bytes,1,opt,name=PushMessage,proto3" json:"PushMessage,omitempty"`
}

func (m *PushMessageReq) Reset()      { *m = PushMessageReq{} }
func (*PushMessageReq) ProtoMessage() {}
func (*PushMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe439df8e494cc1, []int{3}
}
func (m *PushMessageReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushMessageReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMessageReq.Merge(m, src)
}
func (m *PushMessageReq) XXX_Size() int {
	return m.Size()
}
func (m *PushMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_PushMessageReq proto.InternalMessageInfo

func (m *PushMessageReq) GetPushMessage() *PushMessage {
	if m != nil {
		return m.PushMessage
	}
	return nil
}

type PushMessageRes struct {
}

func (m *PushMessageRes) Reset()      { *m = PushMessageRes{} }
func (*PushMessageRes) ProtoMessage() {}
func (*PushMessageRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe439df8e494cc1, []int{4}
}
func (m *PushMessageRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushMessageRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushMessageRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushMessageRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMessageRes.Merge(m, src)
}
func (m *PushMessageRes) XXX_Size() int {
	return m.Size()
}
func (m *PushMessageRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMessageRes.DiscardUnknown(m)
}

var xxx_messageInfo_PushMessageRes proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("internl.Languages", Languages_name, Languages_value)
	proto.RegisterType((*MessageRequest)(nil), "internl.MessageRequest")
	proto.RegisterType((*MessageResponse)(nil), "internl.MessageResponse")
	proto.RegisterType((*PushMessage)(nil), "internl.PushMessage")
	proto.RegisterType((*PushMessageReq)(nil), "internl.PushMessageReq")
	proto.RegisterType((*PushMessageRes)(nil), "internl.PushMessageRes")
}

func init() { proto.RegisterFile("internl.proto", fileDescriptor_bfe439df8e494cc1) }

var fileDescriptor_bfe439df8e494cc1 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0xcb, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x2a, 0xb8,
	0xf8, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0xb4, 0xb8, 0x04, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x61, 0x22, 0xa9, 0x45, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x18, 0xe2, 0x42, 0x6a, 0x5c, 0x2c, 0x3e, 0x89, 0x79, 0xe9, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0x42, 0x7a, 0x30, 0x4b, 0x40, 0x82, 0xa5, 0x89, 0xe9, 0xa9, 0xc5,
	0x41, 0x60, 0x79, 0x21, 0x21, 0x2e, 0x16, 0xe7, 0xfc, 0x94, 0x54, 0x09, 0x66, 0x05, 0x46, 0x0d,
	0xde, 0x20, 0x30, 0x5b, 0x49, 0x9d, 0x8b, 0x1f, 0x6e, 0x73, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa,
	0x90, 0x08, 0x17, 0x6b, 0x58, 0x62, 0x4e, 0x69, 0x2a, 0xd4, 0x3e, 0x08, 0x47, 0x29, 0x9e, 0x8b,
	0x3b, 0xa0, 0xb4, 0x38, 0x03, 0xaa, 0x18, 0xbb, 0x22, 0x8a, 0x5c, 0xe2, 0xc1, 0xc5, 0x87, 0x64,
	0x41, 0x50, 0x6a, 0xa1, 0x90, 0x19, 0x8a, 0x95, 0x60, 0x9b, 0xb8, 0x8d, 0x44, 0xe0, 0x86, 0x22,
	0xab, 0x46, 0x56, 0xa8, 0x24, 0x80, 0x66, 0x52, 0xb1, 0x96, 0x34, 0x17, 0x27, 0xdc, 0x09, 0x42,
	0x6c, 0x5c, 0x4c, 0xae, 0x7e, 0x02, 0x0c, 0x20, 0xda, 0x2d, 0x48, 0x80, 0xd1, 0x29, 0xf7, 0xc2,
	0x43, 0x39, 0x86, 0x1b, 0x0f, 0xe5, 0x18, 0x3e, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8,
	0xe2, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7,
	0xf8, 0xe2, 0x91, 0x1c, 0xc3, 0x87, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c,
	0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x79, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0x7e, 0x52, 0x6a, 0x5e, 0x76, 0xa2, 0x6e, 0x6e, 0xaa, 0x3e, 0xc4, 0x69, 0x89, 0x25,
	0x99, 0xf9, 0x79, 0x89, 0x39, 0x99, 0xc5, 0x60, 0x86, 0x7e, 0x7a, 0xbe, 0x6e, 0x41, 0x76, 0x3a,
	0x54, 0x2e, 0x27, 0x89, 0x0d, 0x1c, 0xf7, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x45,
	0x8c, 0xc3, 0x0c, 0x02, 0x00, 0x00,
}

func (x Languages) String() string {
	s, ok := Languages_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *MessageRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MessageRequest)
	if !ok {
		that2, ok := that.(MessageRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ServiceRequester != that1.ServiceRequester {
		return false
	}
	if this.Lang != that1.Lang {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	return true
}
func (this *MessageResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MessageResponse)
	if !ok {
		that2, ok := that.(MessageResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *PushMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PushMessage)
	if !ok {
		that2, ok := that.(PushMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.Lang != that1.Lang {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	return true
}
func (this *PushMessageReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PushMessageReq)
	if !ok {
		that2, ok := that.(PushMessageReq)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.PushMessage.Equal(that1.PushMessage) {
		return false
	}
	return true
}
func (this *PushMessageRes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PushMessageRes)
	if !ok {
		that2, ok := that.(PushMessageRes)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *MessageRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&internl.MessageRequest{")
	s = append(s, "ServiceRequester: "+fmt.Sprintf("%#v", this.ServiceRequester)+",\n")
	s = append(s, "Lang: "+fmt.Sprintf("%#v", this.Lang)+",\n")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *MessageResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&internl.MessageResponse{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PushMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&internl.PushMessage{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "Lang: "+fmt.Sprintf("%#v", this.Lang)+",\n")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PushMessageReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&internl.PushMessageReq{")
	if this.PushMessage != nil {
		s = append(s, "PushMessage: "+fmt.Sprintf("%#v", this.PushMessage)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PushMessageRes) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&internl.PushMessageRes{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringInternl(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *MessageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MessageRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		i = encodeVarintInternl(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x18
	}
	if m.Lang != 0 {
		i = encodeVarintInternl(dAtA, i, uint64(m.Lang))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ServiceRequester) > 0 {
		i -= len(m.ServiceRequester)
		copy(dAtA[i:], m.ServiceRequester)
		i = encodeVarintInternl(dAtA, i, uint64(len(m.ServiceRequester)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MessageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MessageResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintInternl(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PushMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		i = encodeVarintInternl(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x18
	}
	if m.Lang != 0 {
		i = encodeVarintInternl(dAtA, i, uint64(m.Lang))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintInternl(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PushMessageReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushMessageReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushMessageReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PushMessage != nil {
		{
			size, err := m.PushMessage.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInternl(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PushMessageRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushMessageRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushMessageRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintInternl(dAtA []byte, offset int, v uint64) int {
	offset -= sovInternl(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MessageRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceRequester)
	if l > 0 {
		n += 1 + l + sovInternl(uint64(l))
	}
	if m.Lang != 0 {
		n += 1 + sovInternl(uint64(m.Lang))
	}
	if m.Code != 0 {
		n += 1 + sovInternl(uint64(m.Code))
	}
	return n
}

func (m *MessageResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovInternl(uint64(l))
	}
	return n
}

func (m *PushMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovInternl(uint64(l))
	}
	if m.Lang != 0 {
		n += 1 + sovInternl(uint64(m.Lang))
	}
	if m.Code != 0 {
		n += 1 + sovInternl(uint64(m.Code))
	}
	return n
}

func (m *PushMessageReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PushMessage != nil {
		l = m.PushMessage.Size()
		n += 1 + l + sovInternl(uint64(l))
	}
	return n
}

func (m *PushMessageRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovInternl(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInternl(x uint64) (n int) {
	return sovInternl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *MessageRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MessageRequest{`,
		`ServiceRequester:` + fmt.Sprintf("%v", this.ServiceRequester) + `,`,
		`Lang:` + fmt.Sprintf("%v", this.Lang) + `,`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MessageResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MessageResponse{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PushMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PushMessage{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`Lang:` + fmt.Sprintf("%v", this.Lang) + `,`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PushMessageReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PushMessageReq{`,
		`PushMessage:` + strings.Replace(this.PushMessage.String(), "PushMessage", "PushMessage", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PushMessageRes) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PushMessageRes{`,
		`}`,
	}, "")
	return s
}
func valueToStringInternl(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *MessageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceRequester", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInternl
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInternl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceRequester = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lang", wireType)
			}
			m.Lang = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lang |= Languages(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInternl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternl
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInternl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MessageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MessageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInternl
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInternl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternl
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInternl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PushMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PushMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInternl
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInternl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lang", wireType)
			}
			m.Lang = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lang |= Languages(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInternl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternl
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInternl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PushMessageReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PushMessageReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushMessageReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PushMessage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInternl
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInternl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PushMessage == nil {
				m.PushMessage = &PushMessage{}
			}
			if err := m.PushMessage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternl
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInternl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PushMessageRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PushMessageRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushMessageRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInternl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternl
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInternl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipInternl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInternl
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInternl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInternl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthInternl
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInternl
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInternl
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInternl        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInternl          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInternl = fmt.Errorf("proto: unexpected end of group")
)
