// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto_files/user/user.proto

package auth_srv_user

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

type User struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Birthday             int64    `protobuf:"varint,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_a69ae77dbe5ca84e, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetBirthday() int64 {
	if m != nil {
		return m.Birthday
	}
	return 0
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a69ae77dbe5ca84e, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a69ae77dbe5ca84e, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Birthday             int64    `protobuf:"varint,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Nickname             string   `protobuf:"bytes,6,opt,name=nickname,proto3" json:"nickname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a69ae77dbe5ca84e, []int{3}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *LoginResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginResponse) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *LoginResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginResponse) GetBirthday() int64 {
	if m != nil {
		return m.Birthday
	}
	return 0
}

func (m *LoginResponse) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "auth.srv.user.User")
	proto.RegisterType((*Response)(nil), "auth.srv.user.Response")
	proto.RegisterType((*LoginRequest)(nil), "auth.srv.user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "auth.srv.user.LoginResponse")
}

func init() { proto.RegisterFile("proto_files/user/user.proto", fileDescriptor_a69ae77dbe5ca84e) }

var fileDescriptor_a69ae77dbe5ca84e = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x75, 0x9b, 0x36, 0xb6, 0xa3, 0x05, 0x59, 0x45, 0x97, 0xd6, 0x43, 0xc9, 0xa9, 0xa7, 0x08,
	0x0a, 0x9e, 0xbc, 0x09, 0xa2, 0xe0, 0x29, 0xe2, 0x59, 0xb6, 0xe9, 0xd8, 0x2e, 0xb6, 0xd9, 0xb8,
	0x93, 0x54, 0xfc, 0x10, 0x7f, 0xc0, 0x8b, 0xbf, 0x29, 0xbb, 0x49, 0x8a, 0x5b, 0x6a, 0x2f, 0x21,
	0xef, 0xbd, 0xc9, 0x30, 0xef, 0xbd, 0xc0, 0x30, 0x37, 0xba, 0xd0, 0x2f, 0xaf, 0x6a, 0x81, 0x74,
	0x51, 0x12, 0x1a, 0xf7, 0x88, 0x1d, 0xcb, 0xfb, 0xb2, 0x2c, 0xe6, 0x31, 0x99, 0x55, 0x6c, 0xc9,
	0xe8, 0x9b, 0x41, 0xfb, 0x99, 0xd0, 0xf0, 0x01, 0x74, 0x33, 0x95, 0xbe, 0x65, 0x72, 0x89, 0x82,
	0x8d, 0xd8, 0xb8, 0x97, 0xac, 0xb1, 0xd5, 0xec, 0xb0, 0xd3, 0x5a, 0x95, 0xd6, 0x60, 0xab, 0xe5,
	0x92, 0xe8, 0x43, 0x9b, 0xa9, 0x08, 0x2a, 0xad, 0xc1, 0xfc, 0x04, 0x3a, 0xf9, 0x5c, 0x67, 0x28,
	0xda, 0x4e, 0xa8, 0x80, 0x65, 0x71, 0x29, 0xd5, 0x42, 0x74, 0x2a, 0xd6, 0x01, 0xbb, 0x67, 0xa2,
	0x4c, 0x31, 0x9f, 0xca, 0x4f, 0x11, 0x8e, 0xd8, 0x38, 0x48, 0xd6, 0x38, 0xba, 0x86, 0x6e, 0x82,
	0x94, 0xeb, 0x8c, 0x90, 0x0b, 0xd8, 0xa7, 0x32, 0x4d, 0x91, 0xc8, 0x9d, 0xd9, 0x4d, 0x1a, 0xc8,
	0x8f, 0x20, 0x58, 0xd2, 0xac, 0x3e, 0xd0, 0xbe, 0x46, 0x77, 0x70, 0xf8, 0xa8, 0x67, 0x2a, 0x4b,
	0xf0, 0xbd, 0x44, 0x2a, 0x3c, 0x1f, 0x6c, 0x87, 0x8f, 0x96, 0xef, 0x23, 0xfa, 0x61, 0xd0, 0xaf,
	0x17, 0xd5, 0x57, 0x9c, 0x42, 0x68, 0xbf, 0x7c, 0x98, 0xd6, 0x7b, 0x6a, 0xb4, 0x33, 0xa9, 0x75,
	0x1a, 0xc1, 0xd6, 0x34, 0xda, 0xff, 0xa5, 0xd1, 0xf1, 0xd3, 0xf0, 0x9a, 0x0a, 0xfd, 0xa6, 0x2e,
	0xbf, 0x18, 0x1c, 0xd8, 0x3a, 0x9f, 0xd0, 0xac, 0x54, 0x8a, 0xfc, 0x06, 0xe0, 0xd6, 0xa0, 0x2c,
	0xd0, 0x75, 0x7c, 0x1c, 0x7b, 0xe5, 0xc7, 0x96, 0x1c, 0x9c, 0x6d, 0x90, 0x8d, 0xc7, 0x68, 0x8f,
	0xdf, 0x43, 0xcf, 0x8e, 0x38, 0xeb, 0x7c, 0xb8, 0x31, 0xf7, 0x37, 0xd9, 0xc1, 0xf9, 0x76, 0xb1,
	0xd9, 0x34, 0x09, 0xdd, 0xcf, 0x77, 0xf5, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x01, 0x70, 0x07, 0xd9,
	0x9b, 0x02, 0x00, 0x00,
}
