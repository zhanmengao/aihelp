// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enum.proto

package pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
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

type ChatType int32

const (
	ChatTypeInit   ChatType = 0
	ChatTypeBook   ChatType = 1
	ChatTypeCos    ChatType = 2
	ChatTypeWrite  ChatType = 3
	ChatTypeCoding ChatType = 4
)

var ChatType_name = map[int32]string{
	0: "ChatTypeInit",
	1: "ChatTypeBook",
	2: "ChatTypeCos",
	3: "ChatTypeWrite",
	4: "ChatTypeCoding",
}

var ChatType_value = map[string]int32{
	"ChatTypeInit":   0,
	"ChatTypeBook":   1,
	"ChatTypeCos":    2,
	"ChatTypeWrite":  3,
	"ChatTypeCoding": 4,
}

func (ChatType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_13a9f1b5947140c8, []int{0}
}

type CosPlayRole int32

const (
	CosPlayRoleInit   CosPlayRole = 0
	CosPlayBoyFriend  CosPlayRole = 1
	CosPlayGirlFriend CosPlayRole = 2
	CosPlayCat        CosPlayRole = 3
)

var CosPlayRole_name = map[int32]string{
	0: "CosPlayRoleInit",
	1: "CosPlayBoyFriend",
	2: "CosPlayGirlFriend",
	3: "CosPlayCat",
}

var CosPlayRole_value = map[string]int32{
	"CosPlayRoleInit":   0,
	"CosPlayBoyFriend":  1,
	"CosPlayGirlFriend": 2,
	"CosPlayCat":        3,
}

func (CosPlayRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_13a9f1b5947140c8, []int{1}
}

// 写作类型
type WriteType int32

const (
	WriteInit        WriteType = 0
	WriteWorkSummary WriteType = 1
	WriteTopic       WriteType = 2
)

var WriteType_name = map[int32]string{
	0: "WriteInit",
	1: "WriteWorkSummary",
	2: "WriteTopic",
}

var WriteType_value = map[string]int32{
	"WriteInit":        0,
	"WriteWorkSummary": 1,
	"WriteTopic":       2,
}

func (WriteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_13a9f1b5947140c8, []int{2}
}

// 写作风格
type WriteStyle int32

const (
	WriteStyleInit     WriteStyle = 0
	WriteStyleSolemnly WriteStyle = 1
	WriteStyleHumorous WriteStyle = 2
)

var WriteStyle_name = map[int32]string{
	0: "WriteStyleInit",
	1: "WriteStyleSolemnly",
	2: "WriteStyleHumorous",
}

var WriteStyle_value = map[string]int32{
	"WriteStyleInit":     0,
	"WriteStyleSolemnly": 1,
	"WriteStyleHumorous": 2,
}

func (WriteStyle) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_13a9f1b5947140c8, []int{3}
}

func init() {
	proto.RegisterEnum("zhanmengao.aihelp.proto.ChatType", ChatType_name, ChatType_value)
	proto.RegisterEnum("zhanmengao.aihelp.proto.CosPlayRole", CosPlayRole_name, CosPlayRole_value)
	proto.RegisterEnum("zhanmengao.aihelp.proto.WriteType", WriteType_name, WriteType_value)
	proto.RegisterEnum("zhanmengao.aihelp.proto.WriteStyle", WriteStyle_name, WriteStyle_value)
}

func init() { proto.RegisterFile("enum.proto", fileDescriptor_13a9f1b5947140c8) }

var fileDescriptor_13a9f1b5947140c8 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4b, 0x6e, 0xea, 0x30,
	0x14, 0x86, 0x71, 0xb8, 0xba, 0xba, 0xf7, 0x50, 0xc0, 0xb8, 0xaf, 0x59, 0x16, 0x90, 0x01, 0x51,
	0xd5, 0x0d, 0x54, 0x44, 0xea, 0x43, 0x9d, 0xa0, 0x52, 0x09, 0xa9, 0x93, 0xca, 0x01, 0x43, 0x2c,
	0x6c, 0x9f, 0xc8, 0x71, 0x06, 0xe9, 0x0a, 0x3b, 0x64, 0x09, 0x6d, 0x46, 0x1d, 0x76, 0x09, 0x15,
	0xc1, 0x28, 0x74, 0xe6, 0xef, 0xb3, 0xfd, 0x1f, 0x1d, 0xfd, 0x00, 0xc2, 0x94, 0x7a, 0x9c, 0x5b,
	0x74, 0xc8, 0x2e, 0xdf, 0x32, 0x6e, 0xb4, 0x30, 0x6b, 0x8e, 0x63, 0x2e, 0x33, 0xa1, 0xf2, 0xfd,
	0x45, 0xb4, 0x82, 0x7f, 0x49, 0xc6, 0xdd, 0x73, 0x95, 0x0b, 0x46, 0xe1, 0xe4, 0x70, 0x7e, 0x30,
	0xd2, 0xd1, 0xce, 0xb1, 0x99, 0x20, 0x6e, 0x28, 0x61, 0x43, 0xe8, 0x1d, 0x4c, 0x82, 0x05, 0x0d,
	0xd8, 0x08, 0xfa, 0x07, 0x31, 0xb7, 0xd2, 0x09, 0xda, 0x65, 0x0c, 0x06, 0xed, 0x9b, 0xa5, 0x34,
	0x6b, 0xfa, 0x27, 0x7a, 0x85, 0x5e, 0x82, 0xc5, 0x54, 0xf1, 0xea, 0x09, 0x95, 0x60, 0xa7, 0x30,
	0x3c, 0x42, 0x3f, 0xed, 0x0c, 0xa8, 0x97, 0x13, 0xac, 0x6e, 0xad, 0x14, 0x66, 0x49, 0x09, 0x3b,
	0x87, 0x91, 0xb7, 0x77, 0xd2, 0x2a, 0xaf, 0x03, 0x36, 0x00, 0xf0, 0x3a, 0xe1, 0x8e, 0x76, 0xa3,
	0x1b, 0xf8, 0xdf, 0xcc, 0x6f, 0x36, 0xe9, 0x7b, 0x68, 0x83, 0x1b, 0x9c, 0xa3, 0xdd, 0xcc, 0x4a,
	0xad, 0xb9, 0xad, 0x28, 0xd9, 0x25, 0xec, 0x7f, 0x60, 0x2e, 0x17, 0x34, 0x88, 0xa6, 0x9e, 0x67,
	0xae, 0x52, 0x62, 0xb7, 0x44, 0x4b, 0x3e, 0xe7, 0x02, 0x58, 0xeb, 0x66, 0xa8, 0x84, 0x36, 0x6a,
	0x97, 0xf4, 0xcb, 0xdf, 0x97, 0x1a, 0x2d, 0x96, 0x05, 0x0d, 0x26, 0x8f, 0xdb, 0xcf, 0xb0, 0xf3,
	0x5e, 0x87, 0x64, 0x5b, 0x87, 0xe4, 0xa3, 0x0e, 0xc9, 0x57, 0x1d, 0x76, 0xbe, 0xeb, 0x90, 0xbc,
	0x5c, 0xad, 0xa5, 0xcb, 0xca, 0x74, 0xbc, 0x40, 0x1d, 0xb7, 0xd5, 0xc4, 0xfb, 0x6a, 0x62, 0xcd,
	0x8d, 0x5c, 0x89, 0xc2, 0xc5, 0x4d, 0x47, 0x69, 0xb9, 0x8a, 0xf3, 0x34, 0xfd, 0xdb, 0xc0, 0xf5,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0x16, 0xba, 0x47, 0xd7, 0x01, 0x00, 0x00,
}

func (x ChatType) String() string {
	s, ok := ChatType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x CosPlayRole) String() string {
	s, ok := CosPlayRole_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x WriteType) String() string {
	s, ok := WriteType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x WriteStyle) String() string {
	s, ok := WriteStyle_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
