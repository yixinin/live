// Code generated by protoc-gen-go. DO NOT EDIT.
// source: room.proto

package protocol

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RoomUser struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	VideoPush            bool     `protobuf:"varint,5,opt,name=VideoPush,proto3" json:"VideoPush,omitempty"`
	AudioPush            bool     `protobuf:"varint,6,opt,name=AudioPush,proto3" json:"AudioPush,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomUser) Reset()         { *m = RoomUser{} }
func (m *RoomUser) String() string { return proto.CompactTextString(m) }
func (*RoomUser) ProtoMessage()    {}
func (*RoomUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{0}
}

func (m *RoomUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomUser.Unmarshal(m, b)
}
func (m *RoomUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomUser.Marshal(b, m, deterministic)
}
func (m *RoomUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomUser.Merge(m, src)
}
func (m *RoomUser) XXX_Size() int {
	return xxx_messageInfo_RoomUser.Size(m)
}
func (m *RoomUser) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomUser.DiscardUnknown(m)
}

var xxx_messageInfo_RoomUser proto.InternalMessageInfo

func (m *RoomUser) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RoomUser) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RoomUser) GetVideoPush() bool {
	if m != nil {
		return m.VideoPush
	}
	return false
}

func (m *RoomUser) GetAudioPush() bool {
	if m != nil {
		return m.AudioPush
	}
	return false
}

type CreateRoomReq struct {
	Users                []*RoomUser `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Protocol             string      `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateRoomReq) Reset()         { *m = CreateRoomReq{} }
func (m *CreateRoomReq) String() string { return proto.CompactTextString(m) }
func (*CreateRoomReq) ProtoMessage()    {}
func (*CreateRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{1}
}

func (m *CreateRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomReq.Unmarshal(m, b)
}
func (m *CreateRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomReq.Marshal(b, m, deterministic)
}
func (m *CreateRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomReq.Merge(m, src)
}
func (m *CreateRoomReq) XXX_Size() int {
	return xxx_messageInfo_CreateRoomReq.Size(m)
}
func (m *CreateRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomReq proto.InternalMessageInfo

func (m *CreateRoomReq) GetUsers() []*RoomUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *CreateRoomReq) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type CreateRoomAck struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomAck) Reset()         { *m = CreateRoomAck{} }
func (m *CreateRoomAck) String() string { return proto.CompactTextString(m) }
func (*CreateRoomAck) ProtoMessage()    {}
func (*CreateRoomAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{2}
}

func (m *CreateRoomAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomAck.Unmarshal(m, b)
}
func (m *CreateRoomAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomAck.Marshal(b, m, deterministic)
}
func (m *CreateRoomAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomAck.Merge(m, src)
}
func (m *CreateRoomAck) XXX_Size() int {
	return xxx_messageInfo_CreateRoomAck.Size(m)
}
func (m *CreateRoomAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomAck.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomAck proto.InternalMessageInfo

func (m *CreateRoomAck) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *CreateRoomAck) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type JoinRoomReq struct {
	Header               *CallAckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	User                 *RoomUser      `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	RoomId               int32          `protobuf:"varint,3,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Protocol             string         `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *JoinRoomReq) Reset()         { *m = JoinRoomReq{} }
func (m *JoinRoomReq) String() string { return proto.CompactTextString(m) }
func (*JoinRoomReq) ProtoMessage()    {}
func (*JoinRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{3}
}

func (m *JoinRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoomReq.Unmarshal(m, b)
}
func (m *JoinRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoomReq.Marshal(b, m, deterministic)
}
func (m *JoinRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomReq.Merge(m, src)
}
func (m *JoinRoomReq) XXX_Size() int {
	return xxx_messageInfo_JoinRoomReq.Size(m)
}
func (m *JoinRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomReq proto.InternalMessageInfo

func (m *JoinRoomReq) GetHeader() *CallAckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *JoinRoomReq) GetUser() *RoomUser {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *JoinRoomReq) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *JoinRoomReq) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type JoinRoomAck struct {
	Header               *CallAckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Addr                 string         `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *JoinRoomAck) Reset()         { *m = JoinRoomAck{} }
func (m *JoinRoomAck) String() string { return proto.CompactTextString(m) }
func (*JoinRoomAck) ProtoMessage()    {}
func (*JoinRoomAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{4}
}

func (m *JoinRoomAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoomAck.Unmarshal(m, b)
}
func (m *JoinRoomAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoomAck.Marshal(b, m, deterministic)
}
func (m *JoinRoomAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomAck.Merge(m, src)
}
func (m *JoinRoomAck) XXX_Size() int {
	return xxx_messageInfo_JoinRoomAck.Size(m)
}
func (m *JoinRoomAck) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomAck.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomAck proto.InternalMessageInfo

func (m *JoinRoomAck) GetHeader() *CallAckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *JoinRoomAck) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type LeaveRoomReq struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	RoomId               int32    `protobuf:"varint,2,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRoomReq) Reset()         { *m = LeaveRoomReq{} }
func (m *LeaveRoomReq) String() string { return proto.CompactTextString(m) }
func (*LeaveRoomReq) ProtoMessage()    {}
func (*LeaveRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{5}
}

func (m *LeaveRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRoomReq.Unmarshal(m, b)
}
func (m *LeaveRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRoomReq.Marshal(b, m, deterministic)
}
func (m *LeaveRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoomReq.Merge(m, src)
}
func (m *LeaveRoomReq) XXX_Size() int {
	return xxx_messageInfo_LeaveRoomReq.Size(m)
}
func (m *LeaveRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoomReq proto.InternalMessageInfo

func (m *LeaveRoomReq) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *LeaveRoomReq) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

//最后倒数第二个人离开房间时解散
type LeaveRoomAck struct {
	Header               *CallAckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *LeaveRoomAck) Reset()         { *m = LeaveRoomAck{} }
func (m *LeaveRoomAck) String() string { return proto.CompactTextString(m) }
func (*LeaveRoomAck) ProtoMessage()    {}
func (*LeaveRoomAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{6}
}

func (m *LeaveRoomAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRoomAck.Unmarshal(m, b)
}
func (m *LeaveRoomAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRoomAck.Marshal(b, m, deterministic)
}
func (m *LeaveRoomAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoomAck.Merge(m, src)
}
func (m *LeaveRoomAck) XXX_Size() int {
	return xxx_messageInfo_LeaveRoomAck.Size(m)
}
func (m *LeaveRoomAck) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoomAck.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoomAck proto.InternalMessageInfo

func (m *LeaveRoomAck) GetHeader() *CallAckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

//互相请求 中断通话/标记房间解散
type DiscardRoomReq struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscardRoomReq) Reset()         { *m = DiscardRoomReq{} }
func (m *DiscardRoomReq) String() string { return proto.CompactTextString(m) }
func (*DiscardRoomReq) ProtoMessage()    {}
func (*DiscardRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{7}
}

func (m *DiscardRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscardRoomReq.Unmarshal(m, b)
}
func (m *DiscardRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscardRoomReq.Marshal(b, m, deterministic)
}
func (m *DiscardRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscardRoomReq.Merge(m, src)
}
func (m *DiscardRoomReq) XXX_Size() int {
	return xxx_messageInfo_DiscardRoomReq.Size(m)
}
func (m *DiscardRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscardRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_DiscardRoomReq proto.InternalMessageInfo

func (m *DiscardRoomReq) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

type DiscardRoomAck struct {
	Header               *CallAckHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DiscardRoomAck) Reset()         { *m = DiscardRoomAck{} }
func (m *DiscardRoomAck) String() string { return proto.CompactTextString(m) }
func (*DiscardRoomAck) ProtoMessage()    {}
func (*DiscardRoomAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{8}
}

func (m *DiscardRoomAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscardRoomAck.Unmarshal(m, b)
}
func (m *DiscardRoomAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscardRoomAck.Marshal(b, m, deterministic)
}
func (m *DiscardRoomAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscardRoomAck.Merge(m, src)
}
func (m *DiscardRoomAck) XXX_Size() int {
	return xxx_messageInfo_DiscardRoomAck.Size(m)
}
func (m *DiscardRoomAck) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscardRoomAck.DiscardUnknown(m)
}

var xxx_messageInfo_DiscardRoomAck proto.InternalMessageInfo

func (m *DiscardRoomAck) GetHeader() *CallAckHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*RoomUser)(nil), "protocol.RoomUser")
	proto.RegisterType((*CreateRoomReq)(nil), "protocol.CreateRoomReq")
	proto.RegisterType((*CreateRoomAck)(nil), "protocol.CreateRoomAck")
	proto.RegisterType((*JoinRoomReq)(nil), "protocol.JoinRoomReq")
	proto.RegisterType((*JoinRoomAck)(nil), "protocol.JoinRoomAck")
	proto.RegisterType((*LeaveRoomReq)(nil), "protocol.LeaveRoomReq")
	proto.RegisterType((*LeaveRoomAck)(nil), "protocol.LeaveRoomAck")
	proto.RegisterType((*DiscardRoomReq)(nil), "protocol.DiscardRoomReq")
	proto.RegisterType((*DiscardRoomAck)(nil), "protocol.DiscardRoomAck")
}

func init() { proto.RegisterFile("room.proto", fileDescriptor_c5fd27dd97284ef4) }

var fileDescriptor_c5fd27dd97284ef4 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0x41, 0x8f, 0x93, 0x40,
	0x14, 0xc7, 0xa5, 0x2d, 0x84, 0x3e, 0xaa, 0x31, 0x2f, 0x5a, 0x09, 0xf1, 0xd0, 0xcc, 0xc1, 0x70,
	0xaa, 0x09, 0x5e, 0x8c, 0xc6, 0x28, 0xc1, 0x83, 0x1a, 0x0f, 0x66, 0x4c, 0xbd, 0x23, 0x4c, 0x52,
	0x42, 0xdb, 0xd1, 0x01, 0xfa, 0x35, 0xf6, 0xbe, 0xdf, 0x69, 0xbf, 0xd3, 0x86, 0x19, 0x60, 0xa6,
	0xbb, 0xec, 0x61, 0x9b, 0x4d, 0xf6, 0x04, 0xf3, 0xfe, 0xef, 0xbd, 0xf9, 0xfd, 0xdf, 0x1b, 0x00,
	0xc1, 0xf9, 0x7e, 0xfd, 0x4f, 0xf0, 0x9a, 0xa3, 0x2b, 0x3f, 0x19, 0xdf, 0x05, 0x8b, 0x2d, 0x4b,
	0x73, 0x26, 0x54, 0x9c, 0x1c, 0xc0, 0xa5, 0x9c, 0xef, 0x37, 0x15, 0x13, 0xf8, 0x1c, 0xa6, 0x4d,
	0x91, 0xfb, 0xd6, 0xca, 0x0a, 0xe7, 0xb4, 0xfd, 0xc5, 0x17, 0x60, 0xd7, 0xbc, 0x64, 0x07, 0x7f,
	0x22, 0x63, 0xea, 0x80, 0xaf, 0x61, 0xfe, 0xa7, 0xc8, 0x19, 0xff, 0xd5, 0x54, 0x5b, 0xdf, 0x5e,
	0x59, 0xa1, 0x4b, 0x75, 0xa0, 0x55, 0xe3, 0x26, 0x2f, 0x94, 0xea, 0x28, 0x75, 0x08, 0x90, 0x0d,
	0x3c, 0x4d, 0x04, 0x4b, 0x6b, 0xd6, 0xde, 0x4a, 0xd9, 0x7f, 0x0c, 0xc1, 0x6e, 0x2a, 0x26, 0x2a,
	0xdf, 0x5a, 0x4d, 0x43, 0x2f, 0xc2, 0x75, 0x0f, 0xba, 0xee, 0xb9, 0xa8, 0x4a, 0xc0, 0x00, 0x06,
	0x13, 0x1d, 0xcf, 0x70, 0x26, 0x1f, 0xcd, 0xb6, 0x71, 0x56, 0xe2, 0x12, 0x9c, 0xd6, 0xfd, 0x77,
	0x65, 0xc7, 0xa6, 0xdd, 0x09, 0x11, 0x66, 0x69, 0x9e, 0x8b, 0xae, 0x81, 0xfc, 0x27, 0x97, 0x16,
	0x78, 0x3f, 0x78, 0x71, 0xe8, 0x91, 0xde, 0x82, 0xa3, 0x66, 0x24, 0x6b, 0xbd, 0xe8, 0x95, 0x66,
	0x4a, 0xd2, 0xdd, 0x2e, 0xce, 0xca, 0x6f, 0x52, 0xa6, 0x5d, 0x1a, 0xbe, 0x81, 0x59, 0x8b, 0x28,
	0x9b, 0x8e, 0x5b, 0x90, 0xba, 0x01, 0x35, 0x3d, 0x81, 0x32, 0x9d, 0xcd, 0x6e, 0x38, 0xa3, 0x9a,
	0xad, 0xf5, 0x75, 0x6f, 0xb6, 0x31, 0xc3, 0xef, 0x61, 0xf1, 0x93, 0xa5, 0xc7, 0x61, 0x07, 0xb7,
	0x17, 0xaf, 0x49, 0x27, 0x26, 0x29, 0xf9, 0x6c, 0x54, 0x9e, 0x83, 0x43, 0x42, 0x78, 0xf6, 0xb5,
	0xa8, 0xb2, 0x54, 0xe4, 0xfd, 0xe5, 0x77, 0x6c, 0x8a, 0xc4, 0x27, 0x99, 0xe7, 0x5c, 0x16, 0x5d,
	0x4c, 0xc0, 0x6b, 0x8b, 0x7f, 0x33, 0x71, 0x2c, 0x32, 0x86, 0x5f, 0x00, 0xf4, 0x2b, 0x41, 0xb3,
	0xdc, 0x7c, 0x92, 0xc1, 0xa8, 0x10, 0x67, 0x25, 0x79, 0x82, 0x1f, 0xc0, 0xed, 0xb7, 0x81, 0x2f,
	0x75, 0x9a, 0xf1, 0x7a, 0x82, 0x91, 0xb0, 0xaa, 0xfd, 0x04, 0xf3, 0x61, 0x76, 0xb8, 0xd4, 0x59,
	0xe6, 0x2a, 0x82, 0xb1, 0xb8, 0x2a, 0x4f, 0xc0, 0x33, 0xe6, 0x81, 0xbe, 0x4e, 0x3c, 0x1d, 0x68,
	0x30, 0xae, 0xc8, 0x26, 0xd1, 0x95, 0x05, 0x5e, 0xb2, 0x4d, 0xeb, 0x7e, 0x22, 0x0f, 0xd1, 0xf4,
	0x11, 0x87, 0xf2, 0xd7, 0x91, 0xc2, 0xbb, 0xeb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x94, 0x7c,
	0xb5, 0xeb, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoomServiceClient is the client API for RoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoomServiceClient interface {
	//创建房间
	CreateRoom(ctx context.Context, in *CreateRoomReq, opts ...grpc.CallOption) (*CreateRoomAck, error)
	//加入通话
	JoinRoom(ctx context.Context, in *JoinRoomReq, opts ...grpc.CallOption) (*JoinRoomAck, error)
	//挂断 当挂断至房间只有一个人时 解散房间
	LeaveRoom(ctx context.Context, in *LeaveRoomReq, opts ...grpc.CallOption) (*LeaveRoomAck, error)
	//强制中断通话
	DiscardRoom(ctx context.Context, in *DiscardRoomReq, opts ...grpc.CallOption) (*DiscardRoomAck, error)
}

type roomServiceClient struct {
	cc *grpc.ClientConn
}

func NewRoomServiceClient(cc *grpc.ClientConn) RoomServiceClient {
	return &roomServiceClient{cc}
}

func (c *roomServiceClient) CreateRoom(ctx context.Context, in *CreateRoomReq, opts ...grpc.CallOption) (*CreateRoomAck, error) {
	out := new(CreateRoomAck)
	err := c.cc.Invoke(ctx, "/protocol.RoomService/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) JoinRoom(ctx context.Context, in *JoinRoomReq, opts ...grpc.CallOption) (*JoinRoomAck, error) {
	out := new(JoinRoomAck)
	err := c.cc.Invoke(ctx, "/protocol.RoomService/JoinRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) LeaveRoom(ctx context.Context, in *LeaveRoomReq, opts ...grpc.CallOption) (*LeaveRoomAck, error) {
	out := new(LeaveRoomAck)
	err := c.cc.Invoke(ctx, "/protocol.RoomService/LeaveRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) DiscardRoom(ctx context.Context, in *DiscardRoomReq, opts ...grpc.CallOption) (*DiscardRoomAck, error) {
	out := new(DiscardRoomAck)
	err := c.cc.Invoke(ctx, "/protocol.RoomService/DiscardRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServiceServer is the server API for RoomService service.
type RoomServiceServer interface {
	//创建房间
	CreateRoom(context.Context, *CreateRoomReq) (*CreateRoomAck, error)
	//加入通话
	JoinRoom(context.Context, *JoinRoomReq) (*JoinRoomAck, error)
	//挂断 当挂断至房间只有一个人时 解散房间
	LeaveRoom(context.Context, *LeaveRoomReq) (*LeaveRoomAck, error)
	//强制中断通话
	DiscardRoom(context.Context, *DiscardRoomReq) (*DiscardRoomAck, error)
}

// UnimplementedRoomServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRoomServiceServer struct {
}

func (*UnimplementedRoomServiceServer) CreateRoom(ctx context.Context, req *CreateRoomReq) (*CreateRoomAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (*UnimplementedRoomServiceServer) JoinRoom(ctx context.Context, req *JoinRoomReq) (*JoinRoomAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (*UnimplementedRoomServiceServer) LeaveRoom(ctx context.Context, req *LeaveRoomReq) (*LeaveRoomAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveRoom not implemented")
}
func (*UnimplementedRoomServiceServer) DiscardRoom(ctx context.Context, req *DiscardRoomReq) (*DiscardRoomAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscardRoom not implemented")
}

func RegisterRoomServiceServer(s *grpc.Server, srv RoomServiceServer) {
	s.RegisterService(&_RoomService_serviceDesc, srv)
}

func _RoomService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.RoomService/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).CreateRoom(ctx, req.(*CreateRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_JoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).JoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.RoomService/JoinRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).JoinRoom(ctx, req.(*JoinRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_LeaveRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).LeaveRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.RoomService/LeaveRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).LeaveRoom(ctx, req.(*LeaveRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_DiscardRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscardRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).DiscardRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.RoomService/DiscardRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).DiscardRoom(ctx, req.(*DiscardRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoomService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.RoomService",
	HandlerType: (*RoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _RoomService_CreateRoom_Handler,
		},
		{
			MethodName: "JoinRoom",
			Handler:    _RoomService_JoinRoom_Handler,
		},
		{
			MethodName: "LeaveRoom",
			Handler:    _RoomService_LeaveRoom_Handler,
		},
		{
			MethodName: "DiscardRoom",
			Handler:    _RoomService_DiscardRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "room.proto",
}

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	//解散房间
	DiscardRoom(ctx context.Context, in *DiscardRoomReq, opts ...grpc.CallOption) (*DiscardRoomAck, error)
	//加入通话
	JoinRoom(ctx context.Context, in *JoinRoomReq, opts ...grpc.CallOption) (*JoinRoomAck, error)
	//挂断 当挂断至房间只有一个人时 解散房间
	LeaveRoom(ctx context.Context, in *LeaveRoomReq, opts ...grpc.CallOption) (*LeaveRoomAck, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) DiscardRoom(ctx context.Context, in *DiscardRoomReq, opts ...grpc.CallOption) (*DiscardRoomAck, error) {
	out := new(DiscardRoomAck)
	err := c.cc.Invoke(ctx, "/protocol.ChatService/DiscardRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) JoinRoom(ctx context.Context, in *JoinRoomReq, opts ...grpc.CallOption) (*JoinRoomAck, error) {
	out := new(JoinRoomAck)
	err := c.cc.Invoke(ctx, "/protocol.ChatService/JoinRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) LeaveRoom(ctx context.Context, in *LeaveRoomReq, opts ...grpc.CallOption) (*LeaveRoomAck, error) {
	out := new(LeaveRoomAck)
	err := c.cc.Invoke(ctx, "/protocol.ChatService/LeaveRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	//解散房间
	DiscardRoom(context.Context, *DiscardRoomReq) (*DiscardRoomAck, error)
	//加入通话
	JoinRoom(context.Context, *JoinRoomReq) (*JoinRoomAck, error)
	//挂断 当挂断至房间只有一个人时 解散房间
	LeaveRoom(context.Context, *LeaveRoomReq) (*LeaveRoomAck, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) DiscardRoom(ctx context.Context, req *DiscardRoomReq) (*DiscardRoomAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscardRoom not implemented")
}
func (*UnimplementedChatServiceServer) JoinRoom(ctx context.Context, req *JoinRoomReq) (*JoinRoomAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (*UnimplementedChatServiceServer) LeaveRoom(ctx context.Context, req *LeaveRoomReq) (*LeaveRoomAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveRoom not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_DiscardRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscardRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DiscardRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.ChatService/DiscardRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DiscardRoom(ctx, req.(*DiscardRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_JoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).JoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.ChatService/JoinRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).JoinRoom(ctx, req.(*JoinRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_LeaveRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).LeaveRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.ChatService/LeaveRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).LeaveRoom(ctx, req.(*LeaveRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DiscardRoom",
			Handler:    _ChatService_DiscardRoom_Handler,
		},
		{
			MethodName: "JoinRoom",
			Handler:    _ChatService_JoinRoom_Handler,
		},
		{
			MethodName: "LeaveRoom",
			Handler:    _ChatService_LeaveRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "room.proto",
}
