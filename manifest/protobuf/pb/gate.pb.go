// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gate.proto

package pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// PacketHead PB的packet头部
type PacketHead struct {
	Cmd            int32             `protobuf:"varint,1,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
	Seq            int32             `protobuf:"varint,2,opt,name=Seq,proto3" json:"Seq,omitempty"`
	Time           int64             `protobuf:"varint,3,opt,name=Time,proto3" json:"Time,omitempty"`
	Opts           int64             `protobuf:"varint,4,opt,name=Opts,proto3" json:"Opts,omitempty"`
	RKey           string            `protobuf:"bytes,10,opt,name=RKey,proto3" json:"RKey,omitempty"`
	ClientVer      string            `protobuf:"bytes,11,opt,name=ClientVer,proto3" json:"ClientVer,omitempty"`
	ClientIP       string            `protobuf:"bytes,12,opt,name=ClientIP,proto3" json:"ClientIP,omitempty"`
	Mod            int32             `protobuf:"varint,13,opt,name=Mod,proto3" json:"Mod,omitempty"`
	ReqId          string            `protobuf:"bytes,14,opt,name=ReqId,proto3" json:"ReqId,omitempty"`
	Event          bool              `protobuf:"varint,15,opt,name=Event,proto3" json:"Event,omitempty"`
	ClientRevision string            `protobuf:"bytes,16,opt,name=ClientRevision,proto3" json:"ClientRevision,omitempty"`
	FlowID         uint64            `protobuf:"varint,17,opt,name=FlowID,proto3" json:"FlowID,omitempty"`
	Reserved       []byte            `protobuf:"bytes,18,opt,name=Reserved,proto3" json:"Reserved,omitempty"`
	UID            string            `protobuf:"bytes,20,opt,name=UID,proto3" json:"UID,omitempty"`
	MetaData       map[string]string `protobuf:"bytes,22,rep,name=MetaData,proto3" json:"MetaData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UseTrace       bool              `protobuf:"varint,23,opt,name=UseTrace,proto3" json:"UseTrace,omitempty"`
	// 安全网络
	SendID               int64    `protobuf:"varint,25,opt,name=SendID,proto3" json:"SendID,omitempty"`
	AckPeerID            int64    `protobuf:"varint,26,opt,name=AckPeerID,proto3" json:"AckPeerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PacketHead) Reset()         { *m = PacketHead{} }
func (m *PacketHead) String() string { return proto.CompactTextString(m) }
func (*PacketHead) ProtoMessage()    {}
func (*PacketHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{0}
}
func (m *PacketHead) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketHead.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketHead.Merge(m, src)
}
func (m *PacketHead) XXX_Size() int {
	return m.Size()
}
func (m *PacketHead) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketHead.DiscardUnknown(m)
}

var xxx_messageInfo_PacketHead proto.InternalMessageInfo

func (m *PacketHead) GetCmd() int32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *PacketHead) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PacketHead) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *PacketHead) GetOpts() int64 {
	if m != nil {
		return m.Opts
	}
	return 0
}

func (m *PacketHead) GetRKey() string {
	if m != nil {
		return m.RKey
	}
	return ""
}

func (m *PacketHead) GetClientVer() string {
	if m != nil {
		return m.ClientVer
	}
	return ""
}

func (m *PacketHead) GetClientIP() string {
	if m != nil {
		return m.ClientIP
	}
	return ""
}

func (m *PacketHead) GetMod() int32 {
	if m != nil {
		return m.Mod
	}
	return 0
}

func (m *PacketHead) GetReqId() string {
	if m != nil {
		return m.ReqId
	}
	return ""
}

func (m *PacketHead) GetEvent() bool {
	if m != nil {
		return m.Event
	}
	return false
}

func (m *PacketHead) GetClientRevision() string {
	if m != nil {
		return m.ClientRevision
	}
	return ""
}

func (m *PacketHead) GetFlowID() uint64 {
	if m != nil {
		return m.FlowID
	}
	return 0
}

func (m *PacketHead) GetReserved() []byte {
	if m != nil {
		return m.Reserved
	}
	return nil
}

func (m *PacketHead) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *PacketHead) GetMetaData() map[string]string {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func (m *PacketHead) GetUseTrace() bool {
	if m != nil {
		return m.UseTrace
	}
	return false
}

func (m *PacketHead) GetSendID() int64 {
	if m != nil {
		return m.SendID
	}
	return 0
}

func (m *PacketHead) GetAckPeerID() int64 {
	if m != nil {
		return m.AckPeerID
	}
	return 0
}

// Packet 网络层收到的一个帧格式为Packet结构
type Packet struct {
	Head                 *PacketHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	Body                 []byte      `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{1}
}
func (m *Packet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return m.Size()
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetHead() *PacketHead {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *Packet) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// RspHead 作为Response的一个通用字段
type RspHead struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	SubCode              int32    `protobuf:"varint,3,opt,name=SubCode,proto3" json:"SubCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RspHead) Reset()         { *m = RspHead{} }
func (m *RspHead) String() string { return proto.CompactTextString(m) }
func (*RspHead) ProtoMessage()    {}
func (*RspHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{2}
}
func (m *RspHead) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RspHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RspHead.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RspHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspHead.Merge(m, src)
}
func (m *RspHead) XXX_Size() int {
	return m.Size()
}
func (m *RspHead) XXX_DiscardUnknown() {
	xxx_messageInfo_RspHead.DiscardUnknown(m)
}

var xxx_messageInfo_RspHead proto.InternalMessageInfo

func (m *RspHead) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RspHead) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RspHead) GetSubCode() int32 {
	if m != nil {
		return m.SubCode
	}
	return 0
}

func init() {
	proto.RegisterType((*PacketHead)(nil), "zhanmengao.aihelp.proto.PacketHead")
	proto.RegisterMapType((map[string]string)(nil), "zhanmengao.aihelp.proto.PacketHead.MetaDataEntry")
	proto.RegisterType((*Packet)(nil), "zhanmengao.aihelp.proto.Packet")
	proto.RegisterType((*RspHead)(nil), "zhanmengao.aihelp.proto.RspHead")
}

func init() { proto.RegisterFile("gate.proto", fileDescriptor_743bb58a714d8b7d) }

var fileDescriptor_743bb58a714d8b7d = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xee, 0x36, 0x69, 0x9b, 0x6c, 0xd2, 0x52, 0x56, 0x55, 0xbb, 0x44, 0xc8, 0xb2, 0x82, 0x84,
	0x7c, 0x72, 0xd4, 0x72, 0x00, 0xc1, 0x89, 0x26, 0x45, 0x58, 0x55, 0x44, 0xb4, 0x69, 0x38, 0x70,
	0xdb, 0xc4, 0xd3, 0xc4, 0x4a, 0xec, 0x75, 0xec, 0x4d, 0x50, 0x78, 0x2a, 0x1e, 0x83, 0x63, 0x1f,
	0x01, 0x7c, 0xe2, 0xc8, 0x23, 0xa0, 0x9d, 0xcd, 0x8f, 0x40, 0x42, 0xe2, 0xf6, 0x7d, 0xdf, 0xfc,
	0xec, 0x37, 0x33, 0x4b, 0xe9, 0x58, 0x6a, 0xf0, 0xd3, 0x4c, 0x69, 0xc5, 0x2e, 0xbe, 0x4c, 0x64,
	0x12, 0x43, 0x32, 0x96, 0xca, 0x97, 0xd1, 0x04, 0x66, 0xa9, 0x0d, 0x34, 0xbf, 0x96, 0x29, 0xed,
	0xc9, 0xd1, 0x14, 0xf4, 0x7b, 0x90, 0x21, 0x3b, 0xa5, 0xa5, 0x76, 0x1c, 0x72, 0xe2, 0x12, 0xef,
	0x40, 0x18, 0x68, 0x94, 0x3e, 0xcc, 0xf9, 0xbe, 0x55, 0xfa, 0x30, 0x67, 0x8c, 0x96, 0xef, 0xa2,
	0x18, 0x78, 0xc9, 0x25, 0x5e, 0x49, 0x20, 0x36, 0xda, 0x87, 0x54, 0xe7, 0xbc, 0x6c, 0x35, 0x83,
	0x8d, 0x26, 0x6e, 0x61, 0xc5, 0xa9, 0x4b, 0xbc, 0xaa, 0x40, 0xcc, 0x9e, 0xd2, 0x6a, 0x7b, 0x16,
	0x41, 0xa2, 0x3f, 0x42, 0xc6, 0x6b, 0x18, 0xd8, 0x09, 0xac, 0x41, 0x2b, 0x96, 0x04, 0x3d, 0x5e,
	0xc7, 0xe0, 0x96, 0x1b, 0x1f, 0x5d, 0x15, 0xf2, 0x63, 0xeb, 0xa3, 0xab, 0x42, 0x76, 0x46, 0x0f,
	0x04, 0xcc, 0x83, 0x90, 0x9f, 0x60, 0xaa, 0x25, 0x46, 0xbd, 0x59, 0x42, 0xa2, 0xf9, 0x23, 0x97,
	0x78, 0x15, 0x61, 0x09, 0x7b, 0x4e, 0x4f, 0x6c, 0x27, 0x01, 0xcb, 0x28, 0x8f, 0x54, 0xc2, 0x4f,
	0xb1, 0xe8, 0x2f, 0x95, 0x9d, 0xd3, 0xc3, 0x77, 0x33, 0xf5, 0x39, 0xe8, 0xf0, 0xc7, 0x2e, 0xf1,
	0xca, 0x62, 0xcd, 0x8c, 0x33, 0x01, 0x39, 0x64, 0x4b, 0x08, 0x39, 0x73, 0x89, 0x57, 0x17, 0x5b,
	0x6e, 0x9c, 0x0d, 0x82, 0x0e, 0x3f, 0xc3, 0x86, 0x06, 0xb2, 0x2e, 0xad, 0x74, 0x41, 0xcb, 0x8e,
	0xd4, 0x92, 0x9f, 0xbb, 0x25, 0xaf, 0x76, 0x75, 0xe9, 0xff, 0xe3, 0x00, 0xfe, 0x6e, 0xf9, 0xfe,
	0xa6, 0xe6, 0x26, 0xd1, 0xd9, 0x4a, 0x6c, 0x5b, 0x98, 0xc7, 0x07, 0x39, 0xdc, 0x65, 0x72, 0x04,
	0xfc, 0x02, 0xa7, 0xda, 0x72, 0x63, 0xb8, 0x0f, 0x49, 0x18, 0x74, 0xf8, 0x13, 0x5c, 0xfd, 0x9a,
	0x99, 0x45, 0xbf, 0x1d, 0x4d, 0x7b, 0x00, 0x59, 0xd0, 0xe1, 0x0d, 0x0c, 0xed, 0x84, 0xc6, 0x1b,
	0x7a, 0xfc, 0xc7, 0x63, 0x66, 0x86, 0x29, 0xac, 0xf0, 0xee, 0x55, 0x61, 0xa0, 0xd9, 0xe3, 0x52,
	0xce, 0x16, 0x80, 0x97, 0xaf, 0x0a, 0x4b, 0x5e, 0xef, 0xbf, 0x22, 0xcd, 0x01, 0x3d, 0xb4, 0xa6,
	0xd9, 0x4b, 0x5a, 0x36, 0xc6, 0xb1, 0xac, 0x76, 0xf5, 0xec, 0x3f, 0x66, 0x14, 0x58, 0x60, 0xbe,
	0xc6, 0xb5, 0x0a, 0x57, 0xd8, 0xbb, 0x2e, 0x10, 0x37, 0x03, 0x7a, 0x24, 0xf2, 0x74, 0x13, 0x6e,
	0xab, 0x10, 0xd6, 0xdf, 0x10, 0x31, 0xde, 0x3f, 0x1f, 0xaf, 0xdd, 0x18, 0xc8, 0x38, 0x3d, 0xea,
	0x2f, 0x86, 0x98, 0x58, 0xc2, 0xc4, 0x0d, 0xbd, 0xbe, 0x7d, 0xf8, 0xe1, 0xec, 0x7d, 0x2b, 0x1c,
	0xf2, 0x50, 0x38, 0xe4, 0x7b, 0xe1, 0x90, 0x9f, 0x85, 0xb3, 0xf7, 0xab, 0x70, 0xc8, 0xa7, 0xcb,
	0x71, 0xa4, 0x27, 0x8b, 0xa1, 0x3f, 0x52, 0x71, 0x6b, 0xe7, 0xb6, 0x65, 0xdd, 0xb6, 0x62, 0x99,
	0x44, 0xf7, 0x90, 0xeb, 0x16, 0xda, 0x1e, 0x2e, 0xee, 0x5b, 0xe9, 0x70, 0x78, 0x88, 0xe4, 0xc5,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xfa, 0x52, 0xaa, 0x4f, 0x03, 0x00, 0x00,
}

func (m *PacketHead) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketHead) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketHead) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.AckPeerID != 0 {
		i = encodeVarintGate(dAtA, i, uint64(m.AckPeerID))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xd0
	}
	if m.SendID != 0 {
		i = encodeVarintGate(dAtA, i, uint64(m.SendID))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xc8
	}
	if m.UseTrace {
		i--
		if m.UseTrace {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb8
	}
	if len(m.MetaData) > 0 {
		for k := range m.MetaData {
			v := m.MetaData[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintGate(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGate(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGate(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0xb2
		}
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintGate(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	if len(m.Reserved) > 0 {
		i -= len(m.Reserved)
		copy(dAtA[i:], m.Reserved)
		i = encodeVarintGate(dAtA, i, uint64(len(m.Reserved)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if m.FlowID != 0 {
		i = encodeVarintGate(dAtA, i, uint64(m.FlowID))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if len(m.ClientRevision) > 0 {
		i -= len(m.ClientRevision)
		copy(dAtA[i:], m.ClientRevision)
		i = encodeVarintGate(dAtA, i, uint64(len(m.ClientRevision)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if m.Event {
		i--
		if m.Event {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x78
	}
	if len(m.ReqId) > 0 {
		i -= len(m.ReqId)
		copy(dAtA[i:], m.ReqId)
		i = encodeVarintGate(dAtA, i, uint64(len(m.ReqId)))
		i--
		dAtA[i] = 0x72
	}
	if m.Mod != 0 {
		i = encodeVarintGate(dAtA, i, uint64(m.Mod))
		i--
		dAtA[i] = 0x68
	}
	if len(m.ClientIP) > 0 {
		i -= len(m.ClientIP)
		copy(dAtA[i:], m.ClientIP)
		i = encodeVarintGate(dAtA, i, uint64(len(m.ClientIP)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.ClientVer) > 0 {
		i -= len(m.ClientVer)
		copy(dAtA[i:], m.ClientVer)
		i = encodeVarintGate(dAtA, i, uint64(len(m.ClientVer)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.RKey) > 0 {
		i -= len(m.RKey)
		copy(dAtA[i:], m.RKey)
		i = encodeVarintGate(dAtA, i, uint64(len(m.RKey)))
		i--
		dAtA[i] = 0x52
	}
	if m.Opts != 0 {
		i = encodeVarintGate(dAtA, i, uint64(m.Opts))
		i--
		dAtA[i] = 0x20
	}
	if m.Time != 0 {
		i = encodeVarintGate(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x18
	}
	if m.Seq != 0 {
		i = encodeVarintGate(dAtA, i, uint64(m.Seq))
		i--
		dAtA[i] = 0x10
	}
	if m.Cmd != 0 {
		i = encodeVarintGate(dAtA, i, uint64(m.Cmd))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Packet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Packet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Packet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintGate(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x12
	}
	if m.Head != nil {
		{
			size, err := m.Head.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGate(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RspHead) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RspHead) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RspHead) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SubCode != 0 {
		i = encodeVarintGate(dAtA, i, uint64(m.SubCode))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintGate(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintGate(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGate(dAtA []byte, offset int, v uint64) int {
	offset -= sovGate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PacketHead) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cmd != 0 {
		n += 1 + sovGate(uint64(m.Cmd))
	}
	if m.Seq != 0 {
		n += 1 + sovGate(uint64(m.Seq))
	}
	if m.Time != 0 {
		n += 1 + sovGate(uint64(m.Time))
	}
	if m.Opts != 0 {
		n += 1 + sovGate(uint64(m.Opts))
	}
	l = len(m.RKey)
	if l > 0 {
		n += 1 + l + sovGate(uint64(l))
	}
	l = len(m.ClientVer)
	if l > 0 {
		n += 1 + l + sovGate(uint64(l))
	}
	l = len(m.ClientIP)
	if l > 0 {
		n += 1 + l + sovGate(uint64(l))
	}
	if m.Mod != 0 {
		n += 1 + sovGate(uint64(m.Mod))
	}
	l = len(m.ReqId)
	if l > 0 {
		n += 1 + l + sovGate(uint64(l))
	}
	if m.Event {
		n += 2
	}
	l = len(m.ClientRevision)
	if l > 0 {
		n += 2 + l + sovGate(uint64(l))
	}
	if m.FlowID != 0 {
		n += 2 + sovGate(uint64(m.FlowID))
	}
	l = len(m.Reserved)
	if l > 0 {
		n += 2 + l + sovGate(uint64(l))
	}
	l = len(m.UID)
	if l > 0 {
		n += 2 + l + sovGate(uint64(l))
	}
	if len(m.MetaData) > 0 {
		for k, v := range m.MetaData {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGate(uint64(len(k))) + 1 + len(v) + sovGate(uint64(len(v)))
			n += mapEntrySize + 2 + sovGate(uint64(mapEntrySize))
		}
	}
	if m.UseTrace {
		n += 3
	}
	if m.SendID != 0 {
		n += 2 + sovGate(uint64(m.SendID))
	}
	if m.AckPeerID != 0 {
		n += 2 + sovGate(uint64(m.AckPeerID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Packet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Head != nil {
		l = m.Head.Size()
		n += 1 + l + sovGate(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovGate(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RspHead) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovGate(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovGate(uint64(l))
	}
	if m.SubCode != 0 {
		n += 1 + sovGate(uint64(m.SubCode))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGate(x uint64) (n int) {
	return sovGate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PacketHead) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGate
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
			return fmt.Errorf("proto: PacketHead: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketHead: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cmd", wireType)
			}
			m.Cmd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cmd |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			m.Seq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seq |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Opts", wireType)
			}
			m.Opts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Opts |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
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
				return ErrInvalidLengthGate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientVer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
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
				return ErrInvalidLengthGate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientVer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientIP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
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
				return ErrInvalidLengthGate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientIP = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mod", wireType)
			}
			m.Mod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mod |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
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
				return ErrInvalidLengthGate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReqId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Event = bool(v != 0)
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientRevision", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
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
				return ErrInvalidLengthGate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientRevision = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FlowID", wireType)
			}
			m.FlowID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FlowID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserved", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGate
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reserved = append(m.Reserved[:0], dAtA[iNdEx:postIndex]...)
			if m.Reserved == nil {
				m.Reserved = []byte{}
			}
			iNdEx = postIndex
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
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
				return ErrInvalidLengthGate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
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
				return ErrInvalidLengthGate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetaData == nil {
				m.MetaData = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGate
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGate
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGate
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGate
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGate
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthGate
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthGate
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGate(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGate
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MetaData[mapkey] = mapvalue
			iNdEx = postIndex
		case 23:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UseTrace", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UseTrace = bool(v != 0)
		case 25:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendID", wireType)
			}
			m.SendID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SendID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 26:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckPeerID", wireType)
			}
			m.AckPeerID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AckPeerID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Packet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGate
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
			return fmt.Errorf("proto: Packet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Packet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Head", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
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
				return ErrInvalidLengthGate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Head == nil {
				m.Head = &PacketHead{}
			}
			if err := m.Head.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGate
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RspHead) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGate
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
			return fmt.Errorf("proto: RspHead: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RspHead: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
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
				return ErrInvalidLengthGate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubCode", wireType)
			}
			m.SubCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGate
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
					return 0, ErrIntOverflowGate
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
					return 0, ErrIntOverflowGate
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
				return 0, ErrInvalidLengthGate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGate = fmt.Errorf("proto: unexpected end of group")
)
