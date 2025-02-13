// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package xuperp2p

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

type XuperMessage_MessageType int32

const (
	XuperMessage_SENDBLOCK                XuperMessage_MessageType = 0
	XuperMessage_POSTTX                   XuperMessage_MessageType = 1
	XuperMessage_BATCHPOSTTX              XuperMessage_MessageType = 2
	XuperMessage_GET_BLOCK                XuperMessage_MessageType = 3
	XuperMessage_PING                     XuperMessage_MessageType = 4
	XuperMessage_GET_BLOCKCHAINSTATUS     XuperMessage_MessageType = 5
	XuperMessage_GET_BLOCK_RES            XuperMessage_MessageType = 6
	XuperMessage_GET_BLOCKCHAINSTATUS_RES XuperMessage_MessageType = 7
	// 向邻近确认区块是否为最新状态区块
	XuperMessage_CONFIRM_BLOCKCHAINSTATUS     XuperMessage_MessageType = 8
	XuperMessage_CONFIRM_BLOCKCHAINSTATUS_RES XuperMessage_MessageType = 9
	XuperMessage_MSG_TYPE_NONE                XuperMessage_MessageType = 10
	// 询问RPC端口信息
	XuperMessage_GET_RPC_PORT           XuperMessage_MessageType = 11
	XuperMessage_GET_RPC_PORT_RES       XuperMessage_MessageType = 12
	XuperMessage_GET_AUTHENTICATION     XuperMessage_MessageType = 13
	XuperMessage_GET_AUTHENTICATION_RES XuperMessage_MessageType = 14
)

var XuperMessage_MessageType_name = map[int32]string{
	0:  "SENDBLOCK",
	1:  "POSTTX",
	2:  "BATCHPOSTTX",
	3:  "GET_BLOCK",
	4:  "PING",
	5:  "GET_BLOCKCHAINSTATUS",
	6:  "GET_BLOCK_RES",
	7:  "GET_BLOCKCHAINSTATUS_RES",
	8:  "CONFIRM_BLOCKCHAINSTATUS",
	9:  "CONFIRM_BLOCKCHAINSTATUS_RES",
	10: "MSG_TYPE_NONE",
	11: "GET_RPC_PORT",
	12: "GET_RPC_PORT_RES",
	13: "GET_AUTHENTICATION",
	14: "GET_AUTHENTICATION_RES",
}

var XuperMessage_MessageType_value = map[string]int32{
	"SENDBLOCK":                    0,
	"POSTTX":                       1,
	"BATCHPOSTTX":                  2,
	"GET_BLOCK":                    3,
	"PING":                         4,
	"GET_BLOCKCHAINSTATUS":         5,
	"GET_BLOCK_RES":                6,
	"GET_BLOCKCHAINSTATUS_RES":     7,
	"CONFIRM_BLOCKCHAINSTATUS":     8,
	"CONFIRM_BLOCKCHAINSTATUS_RES": 9,
	"MSG_TYPE_NONE":                10,
	"GET_RPC_PORT":                 11,
	"GET_RPC_PORT_RES":             12,
	"GET_AUTHENTICATION":           13,
	"GET_AUTHENTICATION_RES":       14,
}

func (x XuperMessage_MessageType) String() string {
	return proto.EnumName(XuperMessage_MessageType_name, int32(x))
}

func (XuperMessage_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

type XuperMessage_ErrorType int32

const (
	// success
	XuperMessage_SUCCESS XuperMessage_ErrorType = 0
	XuperMessage_NONE    XuperMessage_ErrorType = 1
	// common error
	XuperMessage_UNKNOW_ERROR             XuperMessage_ErrorType = 2
	XuperMessage_CHECK_SUM_ERROR          XuperMessage_ErrorType = 3
	XuperMessage_UNMARSHAL_MSG_BODY_ERROR XuperMessage_ErrorType = 4
	XuperMessage_CONNECT_REFUSE           XuperMessage_ErrorType = 5
	// block error
	XuperMessage_GET_BLOCKCHAIN_ERROR           XuperMessage_ErrorType = 6
	XuperMessage_BLOCKCHAIN_NOTEXIST            XuperMessage_ErrorType = 7
	XuperMessage_GET_BLOCK_ERROR                XuperMessage_ErrorType = 8
	XuperMessage_CONFIRM_BLOCKCHAINSTATUS_ERROR XuperMessage_ErrorType = 9
	XuperMessage_GET_AUTHENTICATION_ERROR       XuperMessage_ErrorType = 10
)

var XuperMessage_ErrorType_name = map[int32]string{
	0:  "SUCCESS",
	1:  "NONE",
	2:  "UNKNOW_ERROR",
	3:  "CHECK_SUM_ERROR",
	4:  "UNMARSHAL_MSG_BODY_ERROR",
	5:  "CONNECT_REFUSE",
	6:  "GET_BLOCKCHAIN_ERROR",
	7:  "BLOCKCHAIN_NOTEXIST",
	8:  "GET_BLOCK_ERROR",
	9:  "CONFIRM_BLOCKCHAINSTATUS_ERROR",
	10: "GET_AUTHENTICATION_ERROR",
}

var XuperMessage_ErrorType_value = map[string]int32{
	"SUCCESS":                        0,
	"NONE":                           1,
	"UNKNOW_ERROR":                   2,
	"CHECK_SUM_ERROR":                3,
	"UNMARSHAL_MSG_BODY_ERROR":       4,
	"CONNECT_REFUSE":                 5,
	"GET_BLOCKCHAIN_ERROR":           6,
	"BLOCKCHAIN_NOTEXIST":            7,
	"GET_BLOCK_ERROR":                8,
	"CONFIRM_BLOCKCHAINSTATUS_ERROR": 9,
	"GET_AUTHENTICATION_ERROR":       10,
}

func (x XuperMessage_ErrorType) String() string {
	return proto.EnumName(XuperMessage_ErrorType_name, int32(x))
}

func (XuperMessage_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 1}
}

// XuperMessage is the message of Xuper p2p server
type XuperMessage struct {
	Header               *XuperMessage_MessageHeader `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Data                 *XuperMessage_MessageData   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *XuperMessage) Reset()         { *m = XuperMessage{} }
func (m *XuperMessage) String() string { return proto.CompactTextString(m) }
func (*XuperMessage) ProtoMessage()    {}
func (*XuperMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *XuperMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XuperMessage.Unmarshal(m, b)
}
func (m *XuperMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XuperMessage.Marshal(b, m, deterministic)
}
func (m *XuperMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XuperMessage.Merge(m, src)
}
func (m *XuperMessage) XXX_Size() int {
	return xxx_messageInfo_XuperMessage.Size(m)
}
func (m *XuperMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_XuperMessage.DiscardUnknown(m)
}

var xxx_messageInfo_XuperMessage proto.InternalMessageInfo

func (m *XuperMessage) GetHeader() *XuperMessage_MessageHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *XuperMessage) GetData() *XuperMessage_MessageData {
	if m != nil {
		return m.Data
	}
	return nil
}

// MessageHeader is the message header of Xuper p2p server
type XuperMessage_MessageHeader struct {
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// dataCheckSum is the message data checksum, it can be used check where the message have been received
	Logid                string                   `protobuf:"bytes,2,opt,name=logid,proto3" json:"logid,omitempty"`
	From                 string                   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	Bcname               string                   `protobuf:"bytes,4,opt,name=bcname,proto3" json:"bcname,omitempty"`
	Type                 XuperMessage_MessageType `protobuf:"varint,5,opt,name=type,proto3,enum=xuperp2p.XuperMessage_MessageType" json:"type,omitempty"`
	DataCheckSum         uint32                   `protobuf:"varint,6,opt,name=dataCheckSum,proto3" json:"dataCheckSum,omitempty"`
	ErrorType            XuperMessage_ErrorType   `protobuf:"varint,7,opt,name=errorType,proto3,enum=xuperp2p.XuperMessage_ErrorType" json:"errorType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *XuperMessage_MessageHeader) Reset()         { *m = XuperMessage_MessageHeader{} }
func (m *XuperMessage_MessageHeader) String() string { return proto.CompactTextString(m) }
func (*XuperMessage_MessageHeader) ProtoMessage()    {}
func (*XuperMessage_MessageHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

func (m *XuperMessage_MessageHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XuperMessage_MessageHeader.Unmarshal(m, b)
}
func (m *XuperMessage_MessageHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XuperMessage_MessageHeader.Marshal(b, m, deterministic)
}
func (m *XuperMessage_MessageHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XuperMessage_MessageHeader.Merge(m, src)
}
func (m *XuperMessage_MessageHeader) XXX_Size() int {
	return xxx_messageInfo_XuperMessage_MessageHeader.Size(m)
}
func (m *XuperMessage_MessageHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_XuperMessage_MessageHeader.DiscardUnknown(m)
}

var xxx_messageInfo_XuperMessage_MessageHeader proto.InternalMessageInfo

func (m *XuperMessage_MessageHeader) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *XuperMessage_MessageHeader) GetLogid() string {
	if m != nil {
		return m.Logid
	}
	return ""
}

func (m *XuperMessage_MessageHeader) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *XuperMessage_MessageHeader) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *XuperMessage_MessageHeader) GetType() XuperMessage_MessageType {
	if m != nil {
		return m.Type
	}
	return XuperMessage_SENDBLOCK
}

func (m *XuperMessage_MessageHeader) GetDataCheckSum() uint32 {
	if m != nil {
		return m.DataCheckSum
	}
	return 0
}

func (m *XuperMessage_MessageHeader) GetErrorType() XuperMessage_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return XuperMessage_SUCCESS
}

// MessageData is the message data of Xuper p2p server
type XuperMessage_MessageData struct {
	// msgInfo is the message infomation, use protobuf coding style
	MsgInfo              []byte   `protobuf:"bytes,3,opt,name=msgInfo,proto3" json:"msgInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XuperMessage_MessageData) Reset()         { *m = XuperMessage_MessageData{} }
func (m *XuperMessage_MessageData) String() string { return proto.CompactTextString(m) }
func (*XuperMessage_MessageData) ProtoMessage()    {}
func (*XuperMessage_MessageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 1}
}

func (m *XuperMessage_MessageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XuperMessage_MessageData.Unmarshal(m, b)
}
func (m *XuperMessage_MessageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XuperMessage_MessageData.Marshal(b, m, deterministic)
}
func (m *XuperMessage_MessageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XuperMessage_MessageData.Merge(m, src)
}
func (m *XuperMessage_MessageData) XXX_Size() int {
	return xxx_messageInfo_XuperMessage_MessageData.Size(m)
}
func (m *XuperMessage_MessageData) XXX_DiscardUnknown() {
	xxx_messageInfo_XuperMessage_MessageData.DiscardUnknown(m)
}

var xxx_messageInfo_XuperMessage_MessageData proto.InternalMessageInfo

func (m *XuperMessage_MessageData) GetMsgInfo() []byte {
	if m != nil {
		return m.MsgInfo
	}
	return nil
}

func init() {
	proto.RegisterEnum("xuperp2p.XuperMessage_MessageType", XuperMessage_MessageType_name, XuperMessage_MessageType_value)
	proto.RegisterEnum("xuperp2p.XuperMessage_ErrorType", XuperMessage_ErrorType_name, XuperMessage_ErrorType_value)
	proto.RegisterType((*XuperMessage)(nil), "xuperp2p.XuperMessage")
	proto.RegisterType((*XuperMessage_MessageHeader)(nil), "xuperp2p.XuperMessage.MessageHeader")
	proto.RegisterType((*XuperMessage_MessageData)(nil), "xuperp2p.XuperMessage.MessageData")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xd1, 0x4e, 0xdb, 0x30,
	0x14, 0x86, 0x69, 0x69, 0xd3, 0xe6, 0xb4, 0x01, 0xef, 0x80, 0x58, 0x84, 0xd0, 0x54, 0x45, 0x93,
	0xc6, 0x55, 0x2f, 0x98, 0xb4, 0xab, 0x69, 0x52, 0x30, 0x86, 0x44, 0x50, 0xa7, 0xb2, 0x1d, 0x0d,
	0xae, 0xa2, 0x00, 0x81, 0xa1, 0xad, 0x24, 0x4a, 0xcb, 0x34, 0x1e, 0x60, 0xda, 0xdd, 0x9e, 0x6c,
	0x0f, 0x35, 0xd9, 0x49, 0x4b, 0x19, 0x63, 0xbb, 0x6a, 0xcf, 0xff, 0x7f, 0xbf, 0xcf, 0xf1, 0xb1,
	0x14, 0x70, 0x26, 0xd9, 0x74, 0x9a, 0x5e, 0x67, 0xc3, 0xa2, 0xcc, 0x67, 0x39, 0x76, 0xbf, 0xdd,
	0x15, 0x59, 0x59, 0xec, 0x15, 0xde, 0xcf, 0x2e, 0xf4, 0x4f, 0x75, 0x31, 0xaa, 0x00, 0x7c, 0x0f,
	0x56, 0x90, 0xa5, 0x97, 0x59, 0xe9, 0x36, 0x06, 0x8d, 0xdd, 0xde, 0xde, 0xeb, 0xe1, 0x9c, 0x1d,
	0x2e, 0x73, 0xc3, 0xfa, 0xb7, 0x62, 0x45, 0x9d, 0xc1, 0x77, 0xd0, 0x3a, 0x48, 0x67, 0xa9, 0xdb,
	0x34, 0x59, 0xef, 0xdf, 0x59, 0x4d, 0x0a, 0xc3, 0x6f, 0xff, 0x68, 0x82, 0xf3, 0xe8, 0x44, 0x74,
	0xa1, 0xf3, 0x35, 0x2b, 0xa7, 0x37, 0xf9, 0xad, 0x19, 0xc4, 0x16, 0xf3, 0x12, 0x37, 0xa1, 0xfd,
	0x25, 0xbf, 0xbe, 0xb9, 0x34, 0x4d, 0x6c, 0x51, 0x15, 0x88, 0xd0, 0xba, 0x2a, 0xf3, 0x89, 0xbb,
	0x6a, 0x44, 0xf3, 0x1f, 0xb7, 0xc0, 0x3a, 0xbf, 0xb8, 0x4d, 0x27, 0x99, 0xdb, 0x32, 0x6a, 0x5d,
	0xe9, 0x29, 0x67, 0xf7, 0x45, 0xe6, 0xb6, 0x07, 0x8d, 0xdd, 0xb5, 0xff, 0x4d, 0xa9, 0xee, 0x8b,
	0x4c, 0x18, 0x1e, 0x3d, 0xe8, 0x5f, 0xa6, 0xb3, 0x94, 0x7e, 0xca, 0x2e, 0x3e, 0xcb, 0xbb, 0x89,
	0x6b, 0x0d, 0x1a, 0xbb, 0x8e, 0x78, 0xa4, 0xe1, 0x07, 0xb0, 0xb3, 0xb2, 0xcc, 0x4b, 0x1d, 0x73,
	0x3b, 0xa6, 0xc1, 0xe0, 0x99, 0x06, 0x6c, 0xce, 0x89, 0x87, 0xc8, 0xf6, 0x1b, 0xe8, 0x2d, 0xad,
	0x47, 0xaf, 0x61, 0x32, 0xbd, 0x0e, 0x6f, 0xaf, 0x72, 0x73, 0xb3, 0xbe, 0x98, 0x97, 0xde, 0xaf,
	0xe6, 0x82, 0xd4, 0x41, 0x74, 0xc0, 0x96, 0x8c, 0x1f, 0xec, 0x9f, 0x44, 0xf4, 0x98, 0xac, 0x20,
	0x80, 0x35, 0x8e, 0xa4, 0x52, 0xa7, 0xa4, 0x81, 0xeb, 0xd0, 0xdb, 0xf7, 0x15, 0x0d, 0x6a, 0xa1,
	0xa9, 0xd9, 0x23, 0xa6, 0x92, 0x8a, 0x5d, 0xc5, 0x2e, 0xb4, 0xc6, 0x21, 0x3f, 0x22, 0x2d, 0x74,
	0x61, 0x73, 0x61, 0xd0, 0xc0, 0x0f, 0xb9, 0x54, 0xbe, 0x8a, 0x25, 0x69, 0xe3, 0x0b, 0x70, 0x16,
	0x4e, 0x22, 0x98, 0x24, 0x16, 0xee, 0x80, 0xfb, 0x37, 0xd8, 0xb8, 0x1d, 0xed, 0xd2, 0x88, 0x1f,
	0x86, 0x62, 0xf4, 0xf4, 0xb8, 0x2e, 0x0e, 0x60, 0xe7, 0x39, 0xd7, 0xe4, 0x6d, 0xdd, 0x70, 0x24,
	0x8f, 0x12, 0x75, 0x36, 0x66, 0x09, 0x8f, 0x38, 0x23, 0x80, 0x04, 0xfa, 0xba, 0xa1, 0x18, 0xd3,
	0x64, 0x1c, 0x09, 0x45, 0x7a, 0xb8, 0x09, 0x64, 0x59, 0x31, 0xd1, 0x3e, 0x6e, 0x01, 0x6a, 0xd5,
	0x8f, 0x55, 0xc0, 0xb8, 0x0a, 0xa9, 0xaf, 0xc2, 0x88, 0x13, 0x07, 0xb7, 0x61, 0xeb, 0xa9, 0x6e,
	0x32, 0x6b, 0xde, 0xf7, 0x26, 0xd8, 0x8b, 0x07, 0xc1, 0x1e, 0x74, 0x64, 0x4c, 0x29, 0x93, 0x92,
	0xac, 0xe8, 0xf5, 0x98, 0x01, 0x1a, 0x7a, 0x80, 0x98, 0x1f, 0xf3, 0xe8, 0x63, 0xc2, 0x84, 0x88,
	0x04, 0x69, 0xe2, 0x06, 0xac, 0xd3, 0x80, 0xd1, 0xe3, 0x44, 0xc6, 0xa3, 0x5a, 0x5c, 0xd5, 0x57,
	0x8f, 0xf9, 0xc8, 0x17, 0x32, 0xf0, 0x4f, 0x12, 0x7d, 0x89, 0xfd, 0xe8, 0xe0, 0xac, 0x76, 0x5b,
	0x88, 0xb0, 0x46, 0x23, 0xce, 0x19, 0xd5, 0xe3, 0x1e, 0xc6, 0x92, 0x91, 0xf6, 0xd3, 0xbd, 0xd7,
	0xb4, 0x85, 0x2f, 0x61, 0x63, 0x49, 0xe5, 0x91, 0x62, 0xa7, 0xa1, 0x54, 0xa4, 0xa3, 0x3b, 0x3f,
	0x3c, 0x48, 0x45, 0x77, 0xd1, 0x83, 0x57, 0xcf, 0xae, 0xb5, 0x62, 0xec, 0xf9, 0xb3, 0xfd, 0xb1,
	0x85, 0xca, 0x85, 0x73, 0xcb, 0x7c, 0x21, 0xde, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x77,
	0xf5, 0x14, 0x32, 0x04, 0x00, 0x00,
}
