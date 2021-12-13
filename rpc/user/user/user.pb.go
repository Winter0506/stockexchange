// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type PasswordCheckInfo struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	EncryptedPassword    string   `protobuf:"bytes,2,opt,name=encryptedPassword,proto3" json:"encryptedPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordCheckInfo) Reset()         { *m = PasswordCheckInfo{} }
func (m *PasswordCheckInfo) String() string { return proto.CompactTextString(m) }
func (*PasswordCheckInfo) ProtoMessage()    {}
func (*PasswordCheckInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *PasswordCheckInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordCheckInfo.Unmarshal(m, b)
}
func (m *PasswordCheckInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordCheckInfo.Marshal(b, m, deterministic)
}
func (m *PasswordCheckInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordCheckInfo.Merge(m, src)
}
func (m *PasswordCheckInfo) XXX_Size() int {
	return xxx_messageInfo_PasswordCheckInfo.Size(m)
}
func (m *PasswordCheckInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordCheckInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordCheckInfo proto.InternalMessageInfo

func (m *PasswordCheckInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *PasswordCheckInfo) GetEncryptedPassword() string {
	if m != nil {
		return m.EncryptedPassword
	}
	return ""
}

type CheckResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (m *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(m, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type PageInfo struct {
	Pn                   uint32   `protobuf:"varint,1,opt,name=pn,proto3" json:"pn,omitempty"`
	PSize                uint32   `protobuf:"varint,2,opt,name=pSize,proto3" json:"pSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageInfo) Reset()         { *m = PageInfo{} }
func (m *PageInfo) String() string { return proto.CompactTextString(m) }
func (*PageInfo) ProtoMessage()    {}
func (*PageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *PageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageInfo.Unmarshal(m, b)
}
func (m *PageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageInfo.Marshal(b, m, deterministic)
}
func (m *PageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageInfo.Merge(m, src)
}
func (m *PageInfo) XXX_Size() int {
	return xxx_messageInfo_PageInfo.Size(m)
}
func (m *PageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PageInfo proto.InternalMessageInfo

func (m *PageInfo) GetPn() uint32 {
	if m != nil {
		return m.Pn
	}
	return 0
}

func (m *PageInfo) GetPSize() uint32 {
	if m != nil {
		return m.PSize
	}
	return 0
}

type EmailRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailRequest) Reset()         { *m = EmailRequest{} }
func (m *EmailRequest) String() string { return proto.CompactTextString(m) }
func (*EmailRequest) ProtoMessage()    {}
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *EmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailRequest.Unmarshal(m, b)
}
func (m *EmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailRequest.Marshal(b, m, deterministic)
}
func (m *EmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailRequest.Merge(m, src)
}
func (m *EmailRequest) XXX_Size() int {
	return xxx_messageInfo_EmailRequest.Size(m)
}
func (m *EmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailRequest proto.InternalMessageInfo

func (m *EmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type IdRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdRequest) Reset()         { *m = IdRequest{} }
func (m *IdRequest) String() string { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()    {}
func (*IdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *IdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdRequest.Unmarshal(m, b)
}
func (m *IdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdRequest.Marshal(b, m, deterministic)
}
func (m *IdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdRequest.Merge(m, src)
}
func (m *IdRequest) XXX_Size() int {
	return xxx_messageInfo_IdRequest.Size(m)
}
func (m *IdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdRequest proto.InternalMessageInfo

func (m *IdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateUserInfo struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	PassWord             string   `protobuf:"bytes,2,opt,name=passWord,proto3" json:"passWord,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Gender               string   `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserInfo) Reset()         { *m = CreateUserInfo{} }
func (m *CreateUserInfo) String() string { return proto.CompactTextString(m) }
func (*CreateUserInfo) ProtoMessage()    {}
func (*CreateUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *CreateUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserInfo.Unmarshal(m, b)
}
func (m *CreateUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserInfo.Marshal(b, m, deterministic)
}
func (m *CreateUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserInfo.Merge(m, src)
}
func (m *CreateUserInfo) XXX_Size() int {
	return xxx_messageInfo_CreateUserInfo.Size(m)
}
func (m *CreateUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserInfo proto.InternalMessageInfo

func (m *CreateUserInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CreateUserInfo) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

func (m *CreateUserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserInfo) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

type UpdateUserInfo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	PassWord             string   `protobuf:"bytes,3,opt,name=passWord,proto3" json:"passWord,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Gender               string   `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserInfo) Reset()         { *m = UpdateUserInfo{} }
func (m *UpdateUserInfo) String() string { return proto.CompactTextString(m) }
func (*UpdateUserInfo) ProtoMessage()    {}
func (*UpdateUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UpdateUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserInfo.Unmarshal(m, b)
}
func (m *UpdateUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserInfo.Marshal(b, m, deterministic)
}
func (m *UpdateUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserInfo.Merge(m, src)
}
func (m *UpdateUserInfo) XXX_Size() int {
	return xxx_messageInfo_UpdateUserInfo.Size(m)
}
func (m *UpdateUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserInfo proto.InternalMessageInfo

func (m *UpdateUserInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateUserInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UpdateUserInfo) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

func (m *UpdateUserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpdateUserInfo) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

type UserInfoResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	PassWord             string   `protobuf:"bytes,3,opt,name=passWord,proto3" json:"passWord,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Gender               string   `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Role                 int32    `protobuf:"varint,6,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(m, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfoResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserInfoResponse) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

func (m *UserInfoResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfoResponse) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UserInfoResponse) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

type UserListResponse struct {
	Total                int32               `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data                 []*UserInfoResponse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (m *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(m, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *UserListResponse) GetData() []*UserInfoResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "user.Empty")
	proto.RegisterType((*PasswordCheckInfo)(nil), "user.PasswordCheckInfo")
	proto.RegisterType((*CheckResponse)(nil), "user.CheckResponse")
	proto.RegisterType((*PageInfo)(nil), "user.PageInfo")
	proto.RegisterType((*EmailRequest)(nil), "user.EmailRequest")
	proto.RegisterType((*IdRequest)(nil), "user.IdRequest")
	proto.RegisterType((*CreateUserInfo)(nil), "user.CreateUserInfo")
	proto.RegisterType((*UpdateUserInfo)(nil), "user.UpdateUserInfo")
	proto.RegisterType((*UserInfoResponse)(nil), "user.UserInfoResponse")
	proto.RegisterType((*UserListResponse)(nil), "user.UserListResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x15, 0xc7, 0x4e, 0xd3, 0x09, 0x71, 0xe9, 0x10, 0x15, 0xcb, 0x5c, 0xaa, 0x15, 0x87,
	0x82, 0x50, 0x41, 0x45, 0x70, 0x40, 0xea, 0xa5, 0x55, 0x85, 0x22, 0x01, 0x8a, 0x0c, 0x15, 0x27,
	0x0e, 0x6e, 0x3c, 0x04, 0x8b, 0xc4, 0xbb, 0xec, 0x6e, 0x40, 0xe6, 0x01, 0x78, 0x0a, 0x1e, 0x8e,
	0x47, 0x41, 0xde, 0xf5, 0x3a, 0x76, 0x20, 0x5c, 0xb9, 0xf9, 0xb7, 0xff, 0xd9, 0xf9, 0x66, 0xfd,
	0xef, 0x02, 0xac, 0x15, 0xc9, 0x53, 0x21, 0xb9, 0xe6, 0xe8, 0x57, 0xcf, 0x6c, 0x0f, 0x82, 0xab,
	0x95, 0xd0, 0x25, 0xfb, 0x00, 0x87, 0xb3, 0x54, 0xa9, 0x6f, 0x5c, 0x66, 0x97, 0x9f, 0x68, 0xfe,
	0x79, 0x5a, 0x7c, 0xe4, 0x18, 0xc3, 0x50, 0xd4, 0x2f, 0xa3, 0xde, 0x71, 0xef, 0x64, 0x3f, 0x69,
	0x34, 0x3e, 0x82, 0x43, 0x2a, 0xe6, 0xb2, 0x14, 0x9a, 0x32, 0x57, 0x19, 0x79, 0xc6, 0xf4, 0xe7,
	0x07, 0xf6, 0x00, 0xc6, 0x66, 0xd9, 0x84, 0x94, 0xe0, 0x85, 0x22, 0x8c, 0x60, 0x4f, 0xad, 0xe7,
	0x73, 0x52, 0xca, 0xac, 0x3c, 0x4c, 0x9c, 0x64, 0x4f, 0x60, 0x38, 0x4b, 0x17, 0x64, 0x00, 0x42,
	0xf0, 0x44, 0x61, 0x0c, 0xe3, 0xc4, 0x13, 0x05, 0x4e, 0x20, 0x10, 0x6f, 0xf3, 0xef, 0x64, 0x1a,
	0x8d, 0x13, 0x2b, 0xd8, 0x7d, 0xb8, 0x75, 0xb5, 0x4a, 0xf3, 0x65, 0x42, 0x5f, 0xd6, 0xa4, 0x74,
	0xe5, 0xa2, 0x4a, 0xd7, 0xcc, 0x56, 0xb0, 0x7b, 0xb0, 0x3f, 0xcd, 0x9c, 0x25, 0x04, 0x2f, 0xb7,
	0x33, 0x05, 0x89, 0x97, 0x67, 0xec, 0x2b, 0x84, 0x97, 0x92, 0x52, 0x4d, 0xd7, 0x8a, 0xa4, 0x9b,
	0xbd, 0xda, 0xa1, 0x37, 0xe9, 0x8a, 0xdc, 0xec, 0x4e, 0xbb, 0x7d, 0x79, 0xbf, 0x19, 0xb9, 0xd1,
	0x9b, 0xe6, 0xfd, 0x56, 0x73, 0x3c, 0x82, 0xc1, 0x82, 0x8a, 0x8c, 0x64, 0xe4, 0x9b, 0xd7, 0xb5,
	0x62, 0x3f, 0x7a, 0x10, 0x5e, 0x8b, 0xac, 0xdd, 0x78, 0x0b, 0xad, 0x03, 0xe2, 0xfd, 0x03, 0xa4,
	0xbf, 0x0b, 0xc4, 0xff, 0x3b, 0x48, 0xd0, 0x01, 0xf9, 0xd9, 0x83, 0xdb, 0x0e, 0xa1, 0xf9, 0x49,
	0xff, 0x05, 0x05, 0x11, 0x7c, 0xc9, 0x97, 0x14, 0x0d, 0x4c, 0x5f, 0xf3, 0xcc, 0xde, 0x59, 0xba,
	0x57, 0xb9, 0xd2, 0x0d, 0xdd, 0x04, 0x02, 0xcd, 0x75, 0xba, 0xac, 0x01, 0xad, 0xc0, 0x87, 0xe0,
	0x67, 0xa9, 0x4e, 0x23, 0xef, 0xb8, 0x7f, 0x32, 0x3a, 0x3b, 0x3a, 0x35, 0x91, 0xdf, 0x9e, 0x2c,
	0x31, 0x9e, 0xb3, 0x5f, 0x1e, 0xf8, 0xd5, 0x27, 0x7c, 0x06, 0xa3, 0x97, 0xa4, 0x5d, 0x07, 0x0c,
	0x6d, 0x95, 0x8b, 0x61, 0xdc, 0x5a, 0xa5, 0x43, 0x70, 0x0e, 0x07, 0x75, 0xd9, 0x45, 0xf9, 0x9a,
	0xdf, 0xe4, 0x4b, 0x42, 0xb4, 0xd6, 0x76, 0x1e, 0xe3, 0x1d, 0x10, 0xf8, 0xbc, 0xe9, 0x7a, 0x51,
	0x4e, 0x33, 0x3c, 0xb0, 0xb6, 0x26, 0xa4, 0x3b, 0xeb, 0x5e, 0x00, 0x6c, 0xc2, 0x8a, 0x13, 0xeb,
	0xea, 0xc6, 0x77, 0x67, 0xed, 0x63, 0x80, 0x4d, 0xde, 0x5c, 0x6d, 0x37, 0x81, 0xf1, 0xc8, 0xcd,
	0x20, 0x74, 0x89, 0xe7, 0xf5, 0xc9, 0x9d, 0xb9, 0x9f, 0x79, 0xd7, 0x6d, 0xce, 0xd6, 0x6d, 0x11,
	0xdf, 0xa9, 0x41, 0xda, 0xe7, 0xfc, 0x66, 0x60, 0x6e, 0x9b, 0xa7, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xea, 0x3a, 0x84, 0x84, 0x7b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResponse, error)
	GetUserByMobile(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	GetUserById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*Empty, error)
	CheckPassWord(ctx context.Context, in *PasswordCheckInfo, opts ...grpc.CallOption) (*CheckResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/user.User/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserByMobile(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/user.User/GetUserByMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/user.User/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/user.User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/user.User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckPassWord(ctx context.Context, in *PasswordCheckInfo, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/user.User/CheckPassWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetUserList(context.Context, *PageInfo) (*UserListResponse, error)
	GetUserByMobile(context.Context, *EmailRequest) (*UserInfoResponse, error)
	GetUserById(context.Context, *IdRequest) (*UserInfoResponse, error)
	CreateUser(context.Context, *CreateUserInfo) (*UserInfoResponse, error)
	UpdateUser(context.Context, *UpdateUserInfo) (*Empty, error)
	CheckPassWord(context.Context, *PasswordCheckInfo) (*CheckResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUserList(ctx context.Context, req *PageInfo) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (*UnimplementedUserServer) GetUserByMobile(ctx context.Context, req *EmailRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByMobile not implemented")
}
func (*UnimplementedUserServer) GetUserById(ctx context.Context, req *IdRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (*UnimplementedUserServer) CreateUser(ctx context.Context, req *CreateUserInfo) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServer) UpdateUser(ctx context.Context, req *UpdateUserInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUserServer) CheckPassWord(ctx context.Context, req *PasswordCheckInfo) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassWord not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserList(ctx, req.(*PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetUserByMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByMobile(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckPassWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordCheckInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckPassWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/CheckPassWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckPassWord(ctx, req.(*PasswordCheckInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _User_GetUserList_Handler,
		},
		{
			MethodName: "GetUserByMobile",
			Handler:    _User_GetUserByMobile_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _User_GetUserById_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "CheckPassWord",
			Handler:    _User_CheckPassWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
