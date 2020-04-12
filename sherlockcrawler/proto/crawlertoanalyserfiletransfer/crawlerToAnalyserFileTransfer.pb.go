// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crawlerToAnalyserFileTransfer.proto

package crawlerproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//
//UploadStatusCode gives information. the receiver sends a ok if he received all chunks.
type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}

var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}

func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6af6f45688fe9b43, []int{0}
}

//
//Chunk represents a chunk of a http response.
type Chunk struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	TaskId               uint64   `protobuf:"varint,2,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_6af6f45688fe9b43, []int{0}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Chunk) GetTaskId() uint64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

//
//Infos represent all information in relation to the http response. it sends all information of a CrawlerTaskRequest that need to be sent once.
type Infos struct {
	TaskId               uint64         `protobuf:"varint,1,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
	Address              string         `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	Header               []*HeaderArray `protobuf:"bytes,3,rep,name=Header,proto3" json:"Header,omitempty"`
	StatusCode           int32          `protobuf:"varint,4,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	ResponseTime         int64          `protobuf:"varint,5,opt,name=ResponseTime,proto3" json:"ResponseTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Infos) Reset()         { *m = Infos{} }
func (m *Infos) String() string { return proto.CompactTextString(m) }
func (*Infos) ProtoMessage()    {}
func (*Infos) Descriptor() ([]byte, []int) {
	return fileDescriptor_6af6f45688fe9b43, []int{1}
}

func (m *Infos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Infos.Unmarshal(m, b)
}
func (m *Infos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Infos.Marshal(b, m, deterministic)
}
func (m *Infos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Infos.Merge(m, src)
}
func (m *Infos) XXX_Size() int {
	return xxx_messageInfo_Infos.Size(m)
}
func (m *Infos) XXX_DiscardUnknown() {
	xxx_messageInfo_Infos.DiscardUnknown(m)
}

var xxx_messageInfo_Infos proto.InternalMessageInfo

func (m *Infos) GetTaskId() uint64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *Infos) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Infos) GetHeader() []*HeaderArray {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Infos) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Infos) GetResponseTime() int64 {
	if m != nil {
		return m.ResponseTime
	}
	return 0
}

//
//HeaderArray consists of keys and values
type HeaderArray struct {
	Key                  string              `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	ValueArr             []*HeaderArrayValue `protobuf:"bytes,2,rep,name=ValueArr,proto3" json:"ValueArr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *HeaderArray) Reset()         { *m = HeaderArray{} }
func (m *HeaderArray) String() string { return proto.CompactTextString(m) }
func (*HeaderArray) ProtoMessage()    {}
func (*HeaderArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_6af6f45688fe9b43, []int{2}
}

func (m *HeaderArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderArray.Unmarshal(m, b)
}
func (m *HeaderArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderArray.Marshal(b, m, deterministic)
}
func (m *HeaderArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderArray.Merge(m, src)
}
func (m *HeaderArray) XXX_Size() int {
	return xxx_messageInfo_HeaderArray.Size(m)
}
func (m *HeaderArray) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderArray.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderArray proto.InternalMessageInfo

func (m *HeaderArray) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HeaderArray) GetValueArr() []*HeaderArrayValue {
	if m != nil {
		return m.ValueArr
	}
	return nil
}

//
//HeaderArray consists of keys and values
type HeaderArrayValue struct {
	Value                string   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderArrayValue) Reset()         { *m = HeaderArrayValue{} }
func (m *HeaderArrayValue) String() string { return proto.CompactTextString(m) }
func (*HeaderArrayValue) ProtoMessage()    {}
func (*HeaderArrayValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_6af6f45688fe9b43, []int{3}
}

func (m *HeaderArrayValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderArrayValue.Unmarshal(m, b)
}
func (m *HeaderArrayValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderArrayValue.Marshal(b, m, deterministic)
}
func (m *HeaderArrayValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderArrayValue.Merge(m, src)
}
func (m *HeaderArrayValue) XXX_Size() int {
	return xxx_messageInfo_HeaderArrayValue.Size(m)
}
func (m *HeaderArrayValue) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderArrayValue.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderArrayValue proto.InternalMessageInfo

func (m *HeaderArrayValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

//
//ErrorCase will be sent if there is a Taskerror in a CrawlerTaskRequest. no body will be sent in this case.
type ErrorCase struct {
	TaskId               uint64   `protobuf:"varint,1,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	TaskError            string   `protobuf:"bytes,3,opt,name=TaskError,proto3" json:"TaskError,omitempty"`
	ResponseTime         int64    `protobuf:"varint,4,opt,name=ResponseTime,proto3" json:"ResponseTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorCase) Reset()         { *m = ErrorCase{} }
func (m *ErrorCase) String() string { return proto.CompactTextString(m) }
func (*ErrorCase) ProtoMessage()    {}
func (*ErrorCase) Descriptor() ([]byte, []int) {
	return fileDescriptor_6af6f45688fe9b43, []int{4}
}

func (m *ErrorCase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorCase.Unmarshal(m, b)
}
func (m *ErrorCase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorCase.Marshal(b, m, deterministic)
}
func (m *ErrorCase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorCase.Merge(m, src)
}
func (m *ErrorCase) XXX_Size() int {
	return xxx_messageInfo_ErrorCase.Size(m)
}
func (m *ErrorCase) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorCase.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorCase proto.InternalMessageInfo

func (m *ErrorCase) GetTaskId() uint64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *ErrorCase) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ErrorCase) GetTaskError() string {
	if m != nil {
		return m.TaskError
	}
	return ""
}

func (m *ErrorCase) GetResponseTime() int64 {
	if m != nil {
		return m.ResponseTime
	}
	return 0
}

//
//UploadStatus is the response from the receiver.
type UploadStatus struct {
	Message              string           `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code                 UploadStatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=crawlerproto.UploadStatusCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadStatus) Reset()         { *m = UploadStatus{} }
func (m *UploadStatus) String() string { return proto.CompactTextString(m) }
func (*UploadStatus) ProtoMessage()    {}
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_6af6f45688fe9b43, []int{5}
}

func (m *UploadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadStatus.Unmarshal(m, b)
}
func (m *UploadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadStatus.Marshal(b, m, deterministic)
}
func (m *UploadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadStatus.Merge(m, src)
}
func (m *UploadStatus) XXX_Size() int {
	return xxx_messageInfo_UploadStatus.Size(m)
}
func (m *UploadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UploadStatus proto.InternalMessageInfo

func (m *UploadStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadStatus) GetCode() UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadStatusCode_Unknown
}

func init() {
	proto.RegisterEnum("crawlerproto.UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
	proto.RegisterType((*Chunk)(nil), "crawlerproto.Chunk")
	proto.RegisterType((*Infos)(nil), "crawlerproto.Infos")
	proto.RegisterType((*HeaderArray)(nil), "crawlerproto.HeaderArray")
	proto.RegisterType((*HeaderArrayValue)(nil), "crawlerproto.HeaderArrayValue")
	proto.RegisterType((*ErrorCase)(nil), "crawlerproto.ErrorCase")
	proto.RegisterType((*UploadStatus)(nil), "crawlerproto.UploadStatus")
}

func init() {
	proto.RegisterFile("crawlerToAnalyserFileTransfer.proto", fileDescriptor_6af6f45688fe9b43)
}

var fileDescriptor_6af6f45688fe9b43 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x52, 0x41, 0x6b, 0xd4, 0x40,
	0x18, 0x6d, 0x76, 0x37, 0xa9, 0xfb, 0x6d, 0xd0, 0xf0, 0x29, 0x1a, 0x8b, 0x94, 0x12, 0x2f, 0xc1,
	0xc3, 0x82, 0xdb, 0x93, 0x82, 0xe0, 0x12, 0x2c, 0x2d, 0x22, 0xc2, 0x34, 0xf5, 0xa2, 0x97, 0xd1,
	0x7c, 0xd5, 0x92, 0x38, 0xb3, 0xcc, 0x64, 0x29, 0x7b, 0xf2, 0x2f, 0xf9, 0x5f, 0xfc, 0x43, 0xce,
	0x4c, 0x92, 0x36, 0x1b, 0x6d, 0x0f, 0xbd, 0xcd, 0xf7, 0xde, 0x9b, 0x37, 0x8f, 0x37, 0x1f, 0x3c,
	0xff, 0xa6, 0xf8, 0x65, 0x45, 0x2a, 0x97, 0x4b, 0xc1, 0xab, 0x8d, 0x26, 0x75, 0x74, 0x51, 0x51,
	0xae, 0xb8, 0xd0, 0xe7, 0xa4, 0xe6, 0x2b, 0x25, 0x6b, 0x89, 0x61, 0x2b, 0x72, 0x53, 0xf2, 0x0a,
	0xfc, 0xec, 0xc7, 0x5a, 0x94, 0x18, 0xc3, 0x6e, 0x26, 0x45, 0x4d, 0xa2, 0x8e, 0xbd, 0x03, 0x2f,
	0x0d, 0x59, 0x37, 0xe2, 0x63, 0x08, 0x72, 0xae, 0xcb, 0x93, 0x22, 0x1e, 0x19, 0x62, 0xc2, 0xda,
	0x29, 0xf9, 0xed, 0x81, 0x7f, 0x22, 0xce, 0xa5, 0xee, 0x29, 0xbc, 0xbe, 0xc2, 0x7a, 0x2e, 0x8b,
	0x42, 0x91, 0xd6, 0xee, 0xea, 0x94, 0x75, 0x23, 0xbe, 0x84, 0xe0, 0x98, 0x78, 0x41, 0x2a, 0x1e,
	0x1f, 0x8c, 0xd3, 0xd9, 0xe2, 0xe9, 0xbc, 0x9f, 0x6a, 0xde, 0x70, 0x4b, 0xa5, 0xf8, 0x86, 0xb5,
	0x42, 0xdc, 0x07, 0x38, 0xad, 0x79, 0xbd, 0xd6, 0x99, 0x2c, 0x28, 0x9e, 0x18, 0x3f, 0x9f, 0xf5,
	0x10, 0x4c, 0x20, 0x64, 0xa4, 0x57, 0x52, 0x68, 0xca, 0x2f, 0x7e, 0x52, 0xec, 0x1b, 0xc5, 0x98,
	0x6d, 0x61, 0xc9, 0x67, 0x98, 0xf5, 0xac, 0x31, 0x82, 0xf1, 0x7b, 0xda, 0xb8, 0xd0, 0x53, 0x66,
	0x8f, 0xf8, 0x1a, 0xee, 0x7d, 0xe2, 0xd5, 0x9a, 0x0c, 0x6f, 0x22, 0xdb, 0x64, 0xfb, 0x37, 0x26,
	0x73, 0x42, 0x76, 0xa5, 0x4f, 0x52, 0x88, 0x86, 0x2c, 0x3e, 0x02, 0xdf, 0x1d, 0xda, 0x37, 0x9a,
	0x21, 0xf9, 0x05, 0xd3, 0x77, 0x4a, 0x49, 0x95, 0x71, 0x4d, 0x77, 0x28, 0xef, 0x19, 0x4c, 0xad,
	0xc6, 0x59, 0x98, 0xfe, 0x2c, 0x77, 0x0d, 0xfc, 0xd3, 0xc3, 0xe4, 0x3f, 0x3d, 0x7c, 0x81, 0xf0,
	0x6c, 0x55, 0x49, 0x5e, 0x34, 0xfd, 0xd9, 0xb7, 0x3e, 0x18, 0x67, 0xfe, 0xbd, 0x0b, 0xda, 0x8d,
	0xb8, 0x80, 0x89, 0xeb, 0xdb, 0x46, 0xb8, 0x3f, 0x2c, 0xa3, 0xef, 0x61, 0x55, 0xcc, 0x69, 0x5f,
	0x1c, 0x42, 0x34, 0x64, 0x70, 0x06, 0xbb, 0x67, 0xa2, 0x14, 0xf2, 0x52, 0x44, 0x3b, 0x18, 0xc0,
	0xe8, 0x63, 0x19, 0x79, 0x08, 0x10, 0x1c, 0x71, 0xb3, 0xa1, 0x45, 0x34, 0x5a, 0xfc, 0xf1, 0x20,
	0x38, 0x25, 0x61, 0x7f, 0xfa, 0x0d, 0x04, 0xcd, 0x7d, 0x7c, 0xb8, 0xfd, 0x9e, 0xdb, 0xd4, 0xbd,
	0xbd, 0x9b, 0x43, 0x24, 0x3b, 0xa9, 0x87, 0x6f, 0x61, 0xd6, 0x60, 0xcd, 0x72, 0x0e, 0x3c, 0x1c,
	0x78, 0xbb, 0x07, 0x1e, 0xc3, 0x83, 0x06, 0xb9, 0xfe, 0xa5, 0x27, 0xdb, 0x17, 0xae, 0x88, 0xdb,
	0x9d, 0xbe, 0x06, 0x0e, 0x3d, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x45, 0x30, 0xf8, 0x9a,
	0x03, 0x00, 0x00,
}
