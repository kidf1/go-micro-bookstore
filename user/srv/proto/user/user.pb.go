// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_bookstore_srv_user

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Profile struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Fullname             string   `protobuf:"bytes,3,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Image                string   `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	CreatedAt            int64    `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Profile) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Profile) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *Profile) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Profile) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Profile) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Profile) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type CreateProfileRequest struct {
	Profile              *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProfileRequest) Reset()         { *m = CreateProfileRequest{} }
func (m *CreateProfileRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProfileRequest) ProtoMessage()    {}
func (*CreateProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *CreateProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProfileRequest.Unmarshal(m, b)
}
func (m *CreateProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProfileRequest.Marshal(b, m, deterministic)
}
func (m *CreateProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProfileRequest.Merge(m, src)
}
func (m *CreateProfileRequest) XXX_Size() int {
	return xxx_messageInfo_CreateProfileRequest.Size(m)
}
func (m *CreateProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProfileRequest proto.InternalMessageInfo

func (m *CreateProfileRequest) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type CreateProfileResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	Profile              *Profile `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProfileResponse) Reset()         { *m = CreateProfileResponse{} }
func (m *CreateProfileResponse) String() string { return proto.CompactTextString(m) }
func (*CreateProfileResponse) ProtoMessage()    {}
func (*CreateProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *CreateProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProfileResponse.Unmarshal(m, b)
}
func (m *CreateProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProfileResponse.Marshal(b, m, deterministic)
}
func (m *CreateProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProfileResponse.Merge(m, src)
}
func (m *CreateProfileResponse) XXX_Size() int {
	return xxx_messageInfo_CreateProfileResponse.Size(m)
}
func (m *CreateProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProfileResponse proto.InternalMessageInfo

func (m *CreateProfileResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateProfileResponse) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *CreateProfileResponse) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type GetProfileRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProfileRequest) Reset()         { *m = GetProfileRequest{} }
func (m *GetProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetProfileRequest) ProtoMessage()    {}
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *GetProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileRequest.Unmarshal(m, b)
}
func (m *GetProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileRequest.Merge(m, src)
}
func (m *GetProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetProfileRequest.Size(m)
}
func (m *GetProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileRequest proto.InternalMessageInfo

func (m *GetProfileRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetProfileResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	Profile              *Profile `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProfileResponse) Reset()         { *m = GetProfileResponse{} }
func (m *GetProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetProfileResponse) ProtoMessage()    {}
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *GetProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileResponse.Unmarshal(m, b)
}
func (m *GetProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileResponse.Merge(m, src)
}
func (m *GetProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetProfileResponse.Size(m)
}
func (m *GetProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileResponse proto.InternalMessageInfo

func (m *GetProfileResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetProfileResponse) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *GetProfileResponse) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type UpdateProfileRequest struct {
	Profile              *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateProfileRequest) Reset()         { *m = UpdateProfileRequest{} }
func (m *UpdateProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateProfileRequest) ProtoMessage()    {}
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *UpdateProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProfileRequest.Unmarshal(m, b)
}
func (m *UpdateProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProfileRequest.Merge(m, src)
}
func (m *UpdateProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateProfileRequest.Size(m)
}
func (m *UpdateProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProfileRequest proto.InternalMessageInfo

func (m *UpdateProfileRequest) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type UpdateProfileResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	Profile              *Profile `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateProfileResponse) Reset()         { *m = UpdateProfileResponse{} }
func (m *UpdateProfileResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateProfileResponse) ProtoMessage()    {}
func (*UpdateProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{6}
}

func (m *UpdateProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProfileResponse.Unmarshal(m, b)
}
func (m *UpdateProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProfileResponse.Marshal(b, m, deterministic)
}
func (m *UpdateProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProfileResponse.Merge(m, src)
}
func (m *UpdateProfileResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateProfileResponse.Size(m)
}
func (m *UpdateProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProfileResponse proto.InternalMessageInfo

func (m *UpdateProfileResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateProfileResponse) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *UpdateProfileResponse) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type DeleteProfileRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProfileRequest) Reset()         { *m = DeleteProfileRequest{} }
func (m *DeleteProfileRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProfileRequest) ProtoMessage()    {}
func (*DeleteProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{7}
}

func (m *DeleteProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProfileRequest.Unmarshal(m, b)
}
func (m *DeleteProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProfileRequest.Marshal(b, m, deterministic)
}
func (m *DeleteProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProfileRequest.Merge(m, src)
}
func (m *DeleteProfileRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProfileRequest.Size(m)
}
func (m *DeleteProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProfileRequest proto.InternalMessageInfo

func (m *DeleteProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteProfileResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProfileResponse) Reset()         { *m = DeleteProfileResponse{} }
func (m *DeleteProfileResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteProfileResponse) ProtoMessage()    {}
func (*DeleteProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{8}
}

func (m *DeleteProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProfileResponse.Unmarshal(m, b)
}
func (m *DeleteProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProfileResponse.Marshal(b, m, deterministic)
}
func (m *DeleteProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProfileResponse.Merge(m, src)
}
func (m *DeleteProfileResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteProfileResponse.Size(m)
}
func (m *DeleteProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProfileResponse proto.InternalMessageInfo

func (m *DeleteProfileResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteProfileResponse) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type SearchProfileRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchProfileRequest) Reset()         { *m = SearchProfileRequest{} }
func (m *SearchProfileRequest) String() string { return proto.CompactTextString(m) }
func (*SearchProfileRequest) ProtoMessage()    {}
func (*SearchProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{9}
}

func (m *SearchProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchProfileRequest.Unmarshal(m, b)
}
func (m *SearchProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchProfileRequest.Marshal(b, m, deterministic)
}
func (m *SearchProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchProfileRequest.Merge(m, src)
}
func (m *SearchProfileRequest) XXX_Size() int {
	return xxx_messageInfo_SearchProfileRequest.Size(m)
}
func (m *SearchProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchProfileRequest proto.InternalMessageInfo

func (m *SearchProfileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SearchProfileResponse struct {
	Code                 int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string     `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	Profile              []*Profile `protobuf:"bytes,3,rep,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SearchProfileResponse) Reset()         { *m = SearchProfileResponse{} }
func (m *SearchProfileResponse) String() string { return proto.CompactTextString(m) }
func (*SearchProfileResponse) ProtoMessage()    {}
func (*SearchProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{10}
}

func (m *SearchProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchProfileResponse.Unmarshal(m, b)
}
func (m *SearchProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchProfileResponse.Marshal(b, m, deterministic)
}
func (m *SearchProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchProfileResponse.Merge(m, src)
}
func (m *SearchProfileResponse) XXX_Size() int {
	return xxx_messageInfo_SearchProfileResponse.Size(m)
}
func (m *SearchProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchProfileResponse proto.InternalMessageInfo

func (m *SearchProfileResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SearchProfileResponse) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *SearchProfileResponse) GetProfile() []*Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func init() {
	proto.RegisterType((*Profile)(nil), "go.micro.bookstore.srv.user.Profile")
	proto.RegisterType((*CreateProfileRequest)(nil), "go.micro.bookstore.srv.user.CreateProfileRequest")
	proto.RegisterType((*CreateProfileResponse)(nil), "go.micro.bookstore.srv.user.CreateProfileResponse")
	proto.RegisterType((*GetProfileRequest)(nil), "go.micro.bookstore.srv.user.GetProfileRequest")
	proto.RegisterType((*GetProfileResponse)(nil), "go.micro.bookstore.srv.user.GetProfileResponse")
	proto.RegisterType((*UpdateProfileRequest)(nil), "go.micro.bookstore.srv.user.UpdateProfileRequest")
	proto.RegisterType((*UpdateProfileResponse)(nil), "go.micro.bookstore.srv.user.UpdateProfileResponse")
	proto.RegisterType((*DeleteProfileRequest)(nil), "go.micro.bookstore.srv.user.DeleteProfileRequest")
	proto.RegisterType((*DeleteProfileResponse)(nil), "go.micro.bookstore.srv.user.DeleteProfileResponse")
	proto.RegisterType((*SearchProfileRequest)(nil), "go.micro.bookstore.srv.user.SearchProfileRequest")
	proto.RegisterType((*SearchProfileResponse)(nil), "go.micro.bookstore.srv.user.SearchProfileResponse")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xcf, 0xea, 0xd3, 0x40,
	0x10, 0xfe, 0xa5, 0x49, 0x5b, 0x7f, 0x23, 0x15, 0x5c, 0x52, 0x09, 0x15, 0xa1, 0x2c, 0x22, 0xc5,
	0xc3, 0x16, 0xeb, 0x5d, 0x28, 0x15, 0xbc, 0x4a, 0xa4, 0xbd, 0x4a, 0x9a, 0x4c, 0x6b, 0x30, 0xe9,
	0xc6, 0xcd, 0x46, 0x7a, 0x14, 0x7c, 0x2c, 0x1f, 0xc6, 0x57, 0x91, 0xec, 0x26, 0x35, 0x49, 0x43,
	0x63, 0x15, 0xe9, 0xa5, 0x64, 0xfe, 0xec, 0x7c, 0xf3, 0xcd, 0x7c, 0x43, 0x61, 0x9c, 0x08, 0x2e,
	0xf9, 0x3c, 0x4b, 0x51, 0xa8, 0x1f, 0xa6, 0x6c, 0xf2, 0x74, 0xcf, 0x59, 0x1c, 0xfa, 0x82, 0xb3,
	0x2d, 0xe7, 0x9f, 0x53, 0xc9, 0x05, 0xb2, 0x54, 0x7c, 0x65, 0x79, 0x0a, 0xfd, 0x61, 0xc0, 0xf0,
	0xbd, 0xe0, 0xbb, 0x30, 0x42, 0xf2, 0x08, 0x7a, 0x61, 0xe0, 0x18, 0x53, 0x63, 0x76, 0xef, 0xf6,
	0xc2, 0x80, 0x4c, 0xe0, 0x41, 0x9e, 0x73, 0xf0, 0x62, 0x74, 0x7a, 0xca, 0x7b, 0xb2, 0xf3, 0xd8,
	0x2e, 0x8b, 0x22, 0x15, 0x33, 0x75, 0xac, 0xb4, 0x89, 0x0d, 0x7d, 0x8c, 0xbd, 0x30, 0x72, 0x2c,
	0x15, 0xd0, 0x46, 0xee, 0x0d, 0x63, 0x6f, 0x8f, 0x4e, 0x5f, 0x7b, 0x95, 0x41, 0x9e, 0x01, 0xf8,
	0x02, 0x3d, 0x89, 0xc1, 0x47, 0x4f, 0x3a, 0x83, 0xa9, 0x31, 0x33, 0xdd, 0xfb, 0xc2, 0xb3, 0x94,
	0x79, 0x38, 0x4b, 0x82, 0x32, 0x3c, 0xd4, 0xe1, 0xc2, 0xb3, 0x94, 0x74, 0x03, 0xf6, 0x4a, 0xe5,
	0x16, 0x14, 0x5c, 0xfc, 0x92, 0x61, 0x2a, 0xc9, 0x1b, 0x18, 0x26, 0xda, 0xa3, 0xe8, 0x3c, 0x5c,
	0x3c, 0x67, 0x17, 0x86, 0xc0, 0xca, 0xd7, 0xe5, 0x23, 0xfa, 0xdd, 0x80, 0x71, 0xa3, 0x70, 0x9a,
	0xf0, 0x43, 0x8a, 0x84, 0x80, 0xe5, 0xf3, 0x40, 0x97, 0xed, 0xbb, 0xea, 0x9b, 0x3c, 0x81, 0x41,
	0x80, 0x32, 0x27, 0xac, 0xa7, 0x54, 0x58, 0xd5, 0x2e, 0xcc, 0xbf, 0xe9, 0x62, 0x0e, 0x8f, 0xdf,
	0xa1, 0x6c, 0x50, 0xab, 0x2e, 0xc5, 0xa8, 0x2f, 0x85, 0x7e, 0x33, 0x80, 0x54, 0x5f, 0xdc, 0xa0,
	0xe7, 0x0d, 0xd8, 0x6b, 0xb5, 0x9e, 0xff, 0xb0, 0x91, 0x46, 0xe1, 0x1b, 0xb0, 0x7b, 0x01, 0xf6,
	0x5b, 0x8c, 0xf0, 0x8c, 0x5d, 0xe3, 0x72, 0xe8, 0x0a, 0xc6, 0x8d, 0xbc, 0xeb, 0x9b, 0xa5, 0x2f,
	0xc1, 0xfe, 0x80, 0x9e, 0xf0, 0x3f, 0x35, 0xc0, 0x08, 0x58, 0x95, 0xed, 0xab, 0x6f, 0x35, 0x9e,
	0x46, 0xf2, 0xbf, 0x8e, 0xc7, 0xbc, 0x7a, 0x3c, 0x8b, 0x9f, 0x16, 0x58, 0xeb, 0x14, 0x05, 0x39,
	0xc2, 0xa8, 0x76, 0x3e, 0xe4, 0xd5, 0xc5, 0x42, 0x6d, 0x37, 0x3c, 0x59, 0x5c, 0xf3, 0x44, 0x93,
	0xa5, 0x77, 0x84, 0x03, 0xfc, 0xbe, 0x00, 0xc2, 0x2e, 0xd6, 0x38, 0x3b, 0xae, 0xc9, 0xfc, 0x8f,
	0xf3, 0x4f, 0x80, 0x47, 0x18, 0xd5, 0x74, 0xd9, 0x41, 0xb5, 0xed, 0x38, 0x3a, 0xa8, 0xb6, 0xca,
	0x5e, 0x23, 0xd7, 0x44, 0xd6, 0x81, 0xdc, 0x26, 0xdc, 0x0e, 0xe4, 0x56, 0x0d, 0x6b, 0xe4, 0x9a,
	0xd8, 0x3a, 0x90, 0xdb, 0x54, 0xdc, 0x81, 0xdc, 0xaa, 0x65, 0x7a, 0xb7, 0x1d, 0xa8, 0xbf, 0xb4,
	0xd7, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x23, 0x1e, 0x2a, 0xeb, 0x06, 0x00, 0x00,
}