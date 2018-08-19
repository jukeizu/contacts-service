// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contacts.proto

package contacts

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Contact struct {
	ServerId             string   `protobuf:"bytes,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_contacts_1beb50f9b8842eda, []int{0}
}
func (m *Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contact.Unmarshal(m, b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
}
func (dst *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(dst, src)
}
func (m *Contact) XXX_Size() int {
	return xxx_messageInfo_Contact.Size(m)
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func (m *Contact) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *Contact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contact) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Contact) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type SetAddressRequest struct {
	ServerId             string   `protobuf:"bytes,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetAddressRequest) Reset()         { *m = SetAddressRequest{} }
func (m *SetAddressRequest) String() string { return proto.CompactTextString(m) }
func (*SetAddressRequest) ProtoMessage()    {}
func (*SetAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_contacts_1beb50f9b8842eda, []int{1}
}
func (m *SetAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetAddressRequest.Unmarshal(m, b)
}
func (m *SetAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetAddressRequest.Marshal(b, m, deterministic)
}
func (dst *SetAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetAddressRequest.Merge(dst, src)
}
func (m *SetAddressRequest) XXX_Size() int {
	return xxx_messageInfo_SetAddressRequest.Size(m)
}
func (m *SetAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetAddressRequest proto.InternalMessageInfo

func (m *SetAddressRequest) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *SetAddressRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetAddressRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type SetAddressReply struct {
	Contact              *Contact `protobuf:"bytes,1,opt,name=contact,proto3" json:"contact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetAddressReply) Reset()         { *m = SetAddressReply{} }
func (m *SetAddressReply) String() string { return proto.CompactTextString(m) }
func (*SetAddressReply) ProtoMessage()    {}
func (*SetAddressReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_contacts_1beb50f9b8842eda, []int{2}
}
func (m *SetAddressReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetAddressReply.Unmarshal(m, b)
}
func (m *SetAddressReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetAddressReply.Marshal(b, m, deterministic)
}
func (dst *SetAddressReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetAddressReply.Merge(dst, src)
}
func (m *SetAddressReply) XXX_Size() int {
	return xxx_messageInfo_SetAddressReply.Size(m)
}
func (m *SetAddressReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SetAddressReply.DiscardUnknown(m)
}

var xxx_messageInfo_SetAddressReply proto.InternalMessageInfo

func (m *SetAddressReply) GetContact() *Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

type SetPhoneRequest struct {
	ServerId             string   `protobuf:"bytes,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetPhoneRequest) Reset()         { *m = SetPhoneRequest{} }
func (m *SetPhoneRequest) String() string { return proto.CompactTextString(m) }
func (*SetPhoneRequest) ProtoMessage()    {}
func (*SetPhoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_contacts_1beb50f9b8842eda, []int{3}
}
func (m *SetPhoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetPhoneRequest.Unmarshal(m, b)
}
func (m *SetPhoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetPhoneRequest.Marshal(b, m, deterministic)
}
func (dst *SetPhoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetPhoneRequest.Merge(dst, src)
}
func (m *SetPhoneRequest) XXX_Size() int {
	return xxx_messageInfo_SetPhoneRequest.Size(m)
}
func (m *SetPhoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetPhoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetPhoneRequest proto.InternalMessageInfo

func (m *SetPhoneRequest) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *SetPhoneRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetPhoneRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type SetPhoneReply struct {
	Contact              *Contact `protobuf:"bytes,1,opt,name=contact,proto3" json:"contact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetPhoneReply) Reset()         { *m = SetPhoneReply{} }
func (m *SetPhoneReply) String() string { return proto.CompactTextString(m) }
func (*SetPhoneReply) ProtoMessage()    {}
func (*SetPhoneReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_contacts_1beb50f9b8842eda, []int{4}
}
func (m *SetPhoneReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetPhoneReply.Unmarshal(m, b)
}
func (m *SetPhoneReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetPhoneReply.Marshal(b, m, deterministic)
}
func (dst *SetPhoneReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetPhoneReply.Merge(dst, src)
}
func (m *SetPhoneReply) XXX_Size() int {
	return xxx_messageInfo_SetPhoneReply.Size(m)
}
func (m *SetPhoneReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SetPhoneReply.DiscardUnknown(m)
}

var xxx_messageInfo_SetPhoneReply proto.InternalMessageInfo

func (m *SetPhoneReply) GetContact() *Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

type QueryRequest struct {
	ServerId             string   `protobuf:"bytes,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_contacts_1beb50f9b8842eda, []int{5}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (dst *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(dst, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

type QueryReply struct {
	Contacts             []*Contact `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *QueryReply) Reset()         { *m = QueryReply{} }
func (m *QueryReply) String() string { return proto.CompactTextString(m) }
func (*QueryReply) ProtoMessage()    {}
func (*QueryReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_contacts_1beb50f9b8842eda, []int{6}
}
func (m *QueryReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryReply.Unmarshal(m, b)
}
func (m *QueryReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryReply.Marshal(b, m, deterministic)
}
func (dst *QueryReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReply.Merge(dst, src)
}
func (m *QueryReply) XXX_Size() int {
	return xxx_messageInfo_QueryReply.Size(m)
}
func (m *QueryReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReply proto.InternalMessageInfo

func (m *QueryReply) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type RemoveContactRequest struct {
	ServerId             string   `protobuf:"bytes,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveContactRequest) Reset()         { *m = RemoveContactRequest{} }
func (m *RemoveContactRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveContactRequest) ProtoMessage()    {}
func (*RemoveContactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_contacts_1beb50f9b8842eda, []int{7}
}
func (m *RemoveContactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveContactRequest.Unmarshal(m, b)
}
func (m *RemoveContactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveContactRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveContactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveContactRequest.Merge(dst, src)
}
func (m *RemoveContactRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveContactRequest.Size(m)
}
func (m *RemoveContactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveContactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveContactRequest proto.InternalMessageInfo

func (m *RemoveContactRequest) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *RemoveContactRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RemoveContactReply struct {
	Removed              bool     `protobuf:"varint,1,opt,name=removed,proto3" json:"removed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveContactReply) Reset()         { *m = RemoveContactReply{} }
func (m *RemoveContactReply) String() string { return proto.CompactTextString(m) }
func (*RemoveContactReply) ProtoMessage()    {}
func (*RemoveContactReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_contacts_1beb50f9b8842eda, []int{8}
}
func (m *RemoveContactReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveContactReply.Unmarshal(m, b)
}
func (m *RemoveContactReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveContactReply.Marshal(b, m, deterministic)
}
func (dst *RemoveContactReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveContactReply.Merge(dst, src)
}
func (m *RemoveContactReply) XXX_Size() int {
	return xxx_messageInfo_RemoveContactReply.Size(m)
}
func (m *RemoveContactReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveContactReply.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveContactReply proto.InternalMessageInfo

func (m *RemoveContactReply) GetRemoved() bool {
	if m != nil {
		return m.Removed
	}
	return false
}

func init() {
	proto.RegisterType((*Contact)(nil), "contacts.Contact")
	proto.RegisterType((*SetAddressRequest)(nil), "contacts.SetAddressRequest")
	proto.RegisterType((*SetAddressReply)(nil), "contacts.SetAddressReply")
	proto.RegisterType((*SetPhoneRequest)(nil), "contacts.SetPhoneRequest")
	proto.RegisterType((*SetPhoneReply)(nil), "contacts.SetPhoneReply")
	proto.RegisterType((*QueryRequest)(nil), "contacts.QueryRequest")
	proto.RegisterType((*QueryReply)(nil), "contacts.QueryReply")
	proto.RegisterType((*RemoveContactRequest)(nil), "contacts.RemoveContactRequest")
	proto.RegisterType((*RemoveContactReply)(nil), "contacts.RemoveContactReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContactsClient is the client API for Contacts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContactsClient interface {
	SetAddress(ctx context.Context, in *SetAddressRequest, opts ...grpc.CallOption) (*SetAddressReply, error)
	SetPhone(ctx context.Context, in *SetPhoneRequest, opts ...grpc.CallOption) (*SetPhoneReply, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error)
	RemoveContact(ctx context.Context, in *RemoveContactRequest, opts ...grpc.CallOption) (*RemoveContactReply, error)
}

type contactsClient struct {
	cc *grpc.ClientConn
}

func NewContactsClient(cc *grpc.ClientConn) ContactsClient {
	return &contactsClient{cc}
}

func (c *contactsClient) SetAddress(ctx context.Context, in *SetAddressRequest, opts ...grpc.CallOption) (*SetAddressReply, error) {
	out := new(SetAddressReply)
	err := c.cc.Invoke(ctx, "/contacts.Contacts/SetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) SetPhone(ctx context.Context, in *SetPhoneRequest, opts ...grpc.CallOption) (*SetPhoneReply, error) {
	out := new(SetPhoneReply)
	err := c.cc.Invoke(ctx, "/contacts.Contacts/SetPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error) {
	out := new(QueryReply)
	err := c.cc.Invoke(ctx, "/contacts.Contacts/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) RemoveContact(ctx context.Context, in *RemoveContactRequest, opts ...grpc.CallOption) (*RemoveContactReply, error) {
	out := new(RemoveContactReply)
	err := c.cc.Invoke(ctx, "/contacts.Contacts/RemoveContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactsServer is the server API for Contacts service.
type ContactsServer interface {
	SetAddress(context.Context, *SetAddressRequest) (*SetAddressReply, error)
	SetPhone(context.Context, *SetPhoneRequest) (*SetPhoneReply, error)
	Query(context.Context, *QueryRequest) (*QueryReply, error)
	RemoveContact(context.Context, *RemoveContactRequest) (*RemoveContactReply, error)
}

func RegisterContactsServer(s *grpc.Server, srv ContactsServer) {
	s.RegisterService(&_Contacts_serviceDesc, srv)
}

func _Contacts_SetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).SetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contacts.Contacts/SetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).SetAddress(ctx, req.(*SetAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_SetPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).SetPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contacts.Contacts/SetPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).SetPhone(ctx, req.(*SetPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contacts.Contacts/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_RemoveContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).RemoveContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contacts.Contacts/RemoveContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).RemoveContact(ctx, req.(*RemoveContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contacts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contacts.Contacts",
	HandlerType: (*ContactsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetAddress",
			Handler:    _Contacts_SetAddress_Handler,
		},
		{
			MethodName: "SetPhone",
			Handler:    _Contacts_SetPhone_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Contacts_Query_Handler,
		},
		{
			MethodName: "RemoveContact",
			Handler:    _Contacts_RemoveContact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contacts.proto",
}

func init() { proto.RegisterFile("contacts.proto", fileDescriptor_contacts_1beb50f9b8842eda) }

var fileDescriptor_contacts_1beb50f9b8842eda = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xfd, 0xda, 0xb4, 0x5f, 0xe2, 0x68, 0x95, 0x0e, 0x41, 0xd7, 0x28, 0x52, 0xf6, 0x24, 0x8a,
	0x39, 0xd4, 0x83, 0x07, 0x45, 0x14, 0xa1, 0xe0, 0x49, 0x8d, 0x07, 0x4f, 0x1e, 0x62, 0xb3, 0xa0,
	0xd0, 0x26, 0x71, 0x37, 0x2d, 0xe4, 0xcf, 0xf8, 0x5b, 0x25, 0xbb, 0xd9, 0x6c, 0x6a, 0x23, 0x68,
	0xf1, 0xb6, 0x33, 0xf3, 0xf6, 0xcd, 0x9b, 0x37, 0xbb, 0xb0, 0x39, 0x4e, 0xe2, 0x2c, 0x1c, 0x67,
	0xc2, 0x4f, 0x79, 0x92, 0x25, 0xe8, 0xe8, 0x98, 0xbe, 0x81, 0x7d, 0xa3, 0xce, 0xe8, 0x81, 0x23,
	0x18, 0x9f, 0x33, 0x7e, 0x1b, 0x91, 0xd6, 0xa0, 0x75, 0xb8, 0x16, 0x54, 0x31, 0x22, 0x74, 0xe2,
	0x70, 0xca, 0x48, 0x5b, 0xe6, 0xe5, 0x19, 0x09, 0xd8, 0x61, 0x14, 0x71, 0x26, 0x04, 0xb1, 0x64,
	0x5a, 0x87, 0xe8, 0x42, 0x37, 0x7d, 0x4d, 0x62, 0x46, 0x3a, 0x32, 0xaf, 0x02, 0xfa, 0x0c, 0xfd,
	0x47, 0x96, 0x5d, 0x2b, 0x4c, 0xc0, 0xde, 0x67, 0x4c, 0xfc, 0x61, 0x53, 0x7a, 0x09, 0x5b, 0x75,
	0xfa, 0x74, 0x92, 0xe3, 0x31, 0xd8, 0xe5, 0xa0, 0x92, 0x7b, 0x7d, 0xd8, 0xf7, 0x2b, 0x23, 0xca,
	0xa9, 0x03, 0x8d, 0xa0, 0x4f, 0xf2, 0xfe, 0x7d, 0x21, 0x75, 0x55, 0x71, 0xd5, 0xdc, 0x56, 0x7d,
	0xee, 0x0b, 0xe8, 0x19, 0xe2, 0x5f, 0xcb, 0x3a, 0x82, 0x8d, 0x87, 0x19, 0xe3, 0xf9, 0x0f, 0x34,
	0xd1, 0x73, 0x80, 0x12, 0x5b, 0xb4, 0x39, 0x81, 0x6a, 0xcd, 0xa4, 0x35, 0xb0, 0x9a, 0xfb, 0x98,
	0x97, 0x30, 0x02, 0x37, 0x60, 0xd3, 0x64, 0xce, 0x74, 0x69, 0x35, 0x13, 0xa8, 0x0f, 0xf8, 0x85,
	0xa7, 0x10, 0x43, 0xc0, 0xe6, 0x32, 0xab, 0x48, 0x9c, 0x40, 0x87, 0xc3, 0x8f, 0x36, 0x38, 0x25,
	0x54, 0xe0, 0x08, 0xc0, 0x2c, 0x11, 0xf7, 0x8c, 0xde, 0xa5, 0x97, 0xe3, 0xed, 0x36, 0x17, 0xd3,
	0x49, 0x4e, 0xff, 0xe1, 0x15, 0x38, 0xda, 0x73, 0x5c, 0x04, 0xd6, 0x17, 0xec, 0xed, 0x34, 0x95,
	0x14, 0xc3, 0x19, 0x74, 0xa5, 0x97, 0xb8, 0x6d, 0x30, 0xf5, 0x45, 0x78, 0xee, 0x52, 0x5e, 0x5d,
	0xbc, 0x83, 0xde, 0xc2, 0xfc, 0x78, 0x60, 0x80, 0x4d, 0x06, 0x7b, 0xfb, 0xdf, 0xd6, 0x25, 0xe1,
	0xcb, 0x7f, 0xf9, 0x67, 0x4f, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xca, 0x98, 0x37, 0x08, 0xc5,
	0x03, 0x00, 0x00,
}