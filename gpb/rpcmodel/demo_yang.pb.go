// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo_yang.proto

package rpcmodel

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

type User_LevelType int32

const (
	User_Default User_LevelType = 0
	User_Sun     User_LevelType = 1
	User_Earth   User_LevelType = 2
	User_Moon    User_LevelType = 3
	User_Star    User_LevelType = 4
)

var User_LevelType_name = map[int32]string{
	0: "Default",
	1: "Sun",
	2: "Earth",
	3: "Moon",
	4: "Star",
}
var User_LevelType_value = map[string]int32{
	"Default": 0,
	"Sun":     1,
	"Earth":   2,
	"Moon":    3,
	"Star":    4,
}

func (x User_LevelType) String() string {
	return proto.EnumName(User_LevelType_name, int32(x))
}
func (User_LevelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_demo_yang_3b7af0428ded2f0b, []int{0, 0}
}

type User_AddressStringType int32

const (
	User_qingdao  User_AddressStringType = 0
	User_beijing  User_AddressStringType = 1
	User_shanghai User_AddressStringType = 2
)

var User_AddressStringType_name = map[int32]string{
	0: "qingdao",
	1: "beijing",
	2: "shanghai",
}
var User_AddressStringType_value = map[string]int32{
	"qingdao":  0,
	"beijing":  1,
	"shanghai": 2,
}

func (x User_AddressStringType) String() string {
	return proto.EnumName(User_AddressStringType_name, int32(x))
}
func (User_AddressStringType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_demo_yang_3b7af0428ded2f0b, []int{0, 1}
}

type RpcRequest_OpType int32

const (
	RpcRequest_Edit RpcRequest_OpType = 0
	RpcRequest_Get  RpcRequest_OpType = 1
)

var RpcRequest_OpType_name = map[int32]string{
	0: "Edit",
	1: "Get",
}
var RpcRequest_OpType_value = map[string]int32{
	"Edit": 0,
	"Get":  1,
}

func (x RpcRequest_OpType) String() string {
	return proto.EnumName(RpcRequest_OpType_name, int32(x))
}
func (RpcRequest_OpType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_demo_yang_3b7af0428ded2f0b, []int{2, 0}
}

type RpcRequest_DBType int32

const (
	RpcRequest_running   RpcRequest_DBType = 0
	RpcRequest_candidate RpcRequest_DBType = 1
	RpcRequest_startup   RpcRequest_DBType = 2
)

var RpcRequest_DBType_name = map[int32]string{
	0: "running",
	1: "candidate",
	2: "startup",
}
var RpcRequest_DBType_value = map[string]int32{
	"running":   0,
	"candidate": 1,
	"startup":   2,
}

func (x RpcRequest_DBType) String() string {
	return proto.EnumName(RpcRequest_DBType_name, int32(x))
}
func (RpcRequest_DBType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_demo_yang_3b7af0428ded2f0b, []int{2, 1}
}

// [START messages]
type User struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             int32            `protobuf:"varint,2,opt,name=password,proto3" json:"password,omitempty"`
	Role                 int32            `protobuf:"varint,3,opt,name=role,proto3" json:"role,omitempty"`
	Level                User_LevelType   `protobuf:"varint,4,opt,name=level,proto3,enum=rpcmodel.User_LevelType" json:"level,omitempty"`
	IsDriver             bool             `protobuf:"varint,5,opt,name=is_driver,json=isDriver,proto3" json:"is_driver,omitempty"`
	Homeaddress          *UserAddressType `protobuf:"bytes,6,opt,name=homeaddress,proto3" json:"homeaddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_yang_3b7af0428ded2f0b, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() int32 {
	if m != nil {
		return m.Password
	}
	return 0
}

func (m *User) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *User) GetLevel() User_LevelType {
	if m != nil {
		return m.Level
	}
	return User_Default
}

func (m *User) GetIsDriver() bool {
	if m != nil {
		return m.IsDriver
	}
	return false
}

func (m *User) GetHomeaddress() *UserAddressType {
	if m != nil {
		return m.Homeaddress
	}
	return nil
}

type UserAddressType struct {
	// Types that are valid to be assigned to TestOneof:
	//	*UserAddressType_Address
	//	*UserAddressType_Name
	TestOneof            isUserAddressType_TestOneof `protobuf_oneof:"test_oneof"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *UserAddressType) Reset()         { *m = UserAddressType{} }
func (m *UserAddressType) String() string { return proto.CompactTextString(m) }
func (*UserAddressType) ProtoMessage()    {}
func (*UserAddressType) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_yang_3b7af0428ded2f0b, []int{0, 0}
}
func (m *UserAddressType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAddressType.Unmarshal(m, b)
}
func (m *UserAddressType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAddressType.Marshal(b, m, deterministic)
}
func (dst *UserAddressType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAddressType.Merge(dst, src)
}
func (m *UserAddressType) XXX_Size() int {
	return xxx_messageInfo_UserAddressType.Size(m)
}
func (m *UserAddressType) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAddressType.DiscardUnknown(m)
}

var xxx_messageInfo_UserAddressType proto.InternalMessageInfo

type isUserAddressType_TestOneof interface {
	isUserAddressType_TestOneof()
}

type UserAddressType_Address struct {
	Address User_AddressStringType `protobuf:"varint,4,opt,name=address,proto3,enum=rpcmodel.User_AddressStringType,oneof"`
}
type UserAddressType_Name struct {
	Name int32 `protobuf:"varint,9,opt,name=name,proto3,oneof"`
}

func (*UserAddressType_Address) isUserAddressType_TestOneof() {}
func (*UserAddressType_Name) isUserAddressType_TestOneof()    {}

func (m *UserAddressType) GetTestOneof() isUserAddressType_TestOneof {
	if m != nil {
		return m.TestOneof
	}
	return nil
}

func (m *UserAddressType) GetAddress() User_AddressStringType {
	if x, ok := m.GetTestOneof().(*UserAddressType_Address); ok {
		return x.Address
	}
	return User_qingdao
}

func (m *UserAddressType) GetName() int32 {
	if x, ok := m.GetTestOneof().(*UserAddressType_Name); ok {
		return x.Name
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UserAddressType) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UserAddressType_OneofMarshaler, _UserAddressType_OneofUnmarshaler, _UserAddressType_OneofSizer, []interface{}{
		(*UserAddressType_Address)(nil),
		(*UserAddressType_Name)(nil),
	}
}

func _UserAddressType_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UserAddressType)
	// test_oneof
	switch x := m.TestOneof.(type) {
	case *UserAddressType_Address:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Address))
	case *UserAddressType_Name:
		b.EncodeVarint(9<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Name))
	case nil:
	default:
		return fmt.Errorf("UserAddressType.TestOneof has unexpected type %T", x)
	}
	return nil
}

func _UserAddressType_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UserAddressType)
	switch tag {
	case 4: // test_oneof.address
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.TestOneof = &UserAddressType_Address{User_AddressStringType(x)}
		return true, err
	case 9: // test_oneof.name
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.TestOneof = &UserAddressType_Name{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _UserAddressType_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UserAddressType)
	// test_oneof
	switch x := m.TestOneof.(type) {
	case *UserAddressType_Address:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Address))
	case *UserAddressType_Name:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Name))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Validate struct {
	Xmlns                string   `protobuf:"bytes,1,opt,name=xmlns,proto3" json:"xmlns,omitempty"`
	User                 []*User  `protobuf:"bytes,2,rep,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Validate) Reset()         { *m = Validate{} }
func (m *Validate) String() string { return proto.CompactTextString(m) }
func (*Validate) ProtoMessage()    {}
func (*Validate) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_yang_3b7af0428ded2f0b, []int{1}
}
func (m *Validate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Validate.Unmarshal(m, b)
}
func (m *Validate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Validate.Marshal(b, m, deterministic)
}
func (dst *Validate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validate.Merge(dst, src)
}
func (m *Validate) XXX_Size() int {
	return xxx_messageInfo_Validate.Size(m)
}
func (m *Validate) XXX_DiscardUnknown() {
	xxx_messageInfo_Validate.DiscardUnknown(m)
}

var xxx_messageInfo_Validate proto.InternalMessageInfo

func (m *Validate) GetXmlns() string {
	if m != nil {
		return m.Xmlns
	}
	return ""
}

func (m *Validate) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

type RpcRequest struct {
	Operation            RpcRequest_OpType `protobuf:"varint,1,opt,name=operation,proto3,enum=rpcmodel.RpcRequest_OpType" json:"operation,omitempty"`
	Db                   RpcRequest_DBType `protobuf:"varint,2,opt,name=db,proto3,enum=rpcmodel.RpcRequest_DBType" json:"db,omitempty"`
	Validate             *Validate         `protobuf:"bytes,3,opt,name=validate,proto3" json:"validate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RpcRequest) Reset()         { *m = RpcRequest{} }
func (m *RpcRequest) String() string { return proto.CompactTextString(m) }
func (*RpcRequest) ProtoMessage()    {}
func (*RpcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_yang_3b7af0428ded2f0b, []int{2}
}
func (m *RpcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcRequest.Unmarshal(m, b)
}
func (m *RpcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcRequest.Marshal(b, m, deterministic)
}
func (dst *RpcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcRequest.Merge(dst, src)
}
func (m *RpcRequest) XXX_Size() int {
	return xxx_messageInfo_RpcRequest.Size(m)
}
func (m *RpcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RpcRequest proto.InternalMessageInfo

func (m *RpcRequest) GetOperation() RpcRequest_OpType {
	if m != nil {
		return m.Operation
	}
	return RpcRequest_Edit
}

func (m *RpcRequest) GetDb() RpcRequest_DBType {
	if m != nil {
		return m.Db
	}
	return RpcRequest_running
}

func (m *RpcRequest) GetValidate() *Validate {
	if m != nil {
		return m.Validate
	}
	return nil
}

type RpcReply struct {
	Result               bool      `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	FailReason           string    `protobuf:"bytes,2,opt,name=failReason,proto3" json:"failReason,omitempty"`
	Validate             *Validate `protobuf:"bytes,3,opt,name=validate,proto3" json:"validate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RpcReply) Reset()         { *m = RpcReply{} }
func (m *RpcReply) String() string { return proto.CompactTextString(m) }
func (*RpcReply) ProtoMessage()    {}
func (*RpcReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_yang_3b7af0428ded2f0b, []int{3}
}
func (m *RpcReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcReply.Unmarshal(m, b)
}
func (m *RpcReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcReply.Marshal(b, m, deterministic)
}
func (dst *RpcReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcReply.Merge(dst, src)
}
func (m *RpcReply) XXX_Size() int {
	return xxx_messageInfo_RpcReply.Size(m)
}
func (m *RpcReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcReply.DiscardUnknown(m)
}

var xxx_messageInfo_RpcReply proto.InternalMessageInfo

func (m *RpcReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *RpcReply) GetFailReason() string {
	if m != nil {
		return m.FailReason
	}
	return ""
}

func (m *RpcReply) GetValidate() *Validate {
	if m != nil {
		return m.Validate
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "rpcmodel.user")
	proto.RegisterType((*UserAddressType)(nil), "rpcmodel.user.addressType")
	proto.RegisterType((*Validate)(nil), "rpcmodel.validate")
	proto.RegisterType((*RpcRequest)(nil), "rpcmodel.rpcRequest")
	proto.RegisterType((*RpcReply)(nil), "rpcmodel.rpcReply")
	proto.RegisterEnum("rpcmodel.User_LevelType", User_LevelType_name, User_LevelType_value)
	proto.RegisterEnum("rpcmodel.User_AddressStringType", User_AddressStringType_name, User_AddressStringType_value)
	proto.RegisterEnum("rpcmodel.RpcRequest_OpType", RpcRequest_OpType_name, RpcRequest_OpType_value)
	proto.RegisterEnum("rpcmodel.RpcRequest_DBType", RpcRequest_DBType_name, RpcRequest_DBType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetingClient is the client API for Greeting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetingClient interface {
	// Sends a greeting
	ExecRPC(ctx context.Context, in *RpcRequest, opts ...grpc.CallOption) (*RpcReply, error)
}

type greetingClient struct {
	cc *grpc.ClientConn
}

func NewGreetingClient(cc *grpc.ClientConn) GreetingClient {
	return &greetingClient{cc}
}

func (c *greetingClient) ExecRPC(ctx context.Context, in *RpcRequest, opts ...grpc.CallOption) (*RpcReply, error) {
	out := new(RpcReply)
	err := c.cc.Invoke(ctx, "/rpcmodel.Greeting/ExecRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetingServer is the server API for Greeting service.
type GreetingServer interface {
	// Sends a greeting
	ExecRPC(context.Context, *RpcRequest) (*RpcReply, error)
}

func RegisterGreetingServer(s *grpc.Server, srv GreetingServer) {
	s.RegisterService(&_Greeting_serviceDesc, srv)
}

func _Greeting_ExecRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServer).ExecRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcmodel.Greeting/ExecRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServer).ExecRPC(ctx, req.(*RpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcmodel.Greeting",
	HandlerType: (*GreetingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecRPC",
			Handler:    _Greeting_ExecRPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo_yang.proto",
}

func init() { proto.RegisterFile("demo_yang.proto", fileDescriptor_demo_yang_3b7af0428ded2f0b) }

var fileDescriptor_demo_yang_3b7af0428ded2f0b = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x6e, 0xd2, 0xbf, 0xe4, 0x74, 0x8c, 0x60, 0x55, 0x28, 0xea, 0x24, 0x14, 0x45, 0x5c, 0x54,
	0x42, 0x8a, 0x44, 0x11, 0x17, 0x88, 0x5d, 0xb0, 0xd2, 0x6a, 0xbb, 0x00, 0x31, 0x65, 0xdc, 0x4f,
	0x6e, 0x73, 0x9a, 0x1a, 0x1c, 0x3b, 0xb3, 0x9d, 0xb1, 0xbe, 0x12, 0xef, 0xc1, 0x1b, 0xf1, 0x00,
	0x28, 0x4e, 0xbb, 0x96, 0x01, 0x17, 0xdc, 0xf9, 0xf8, 0xfb, 0xf1, 0x39, 0x9f, 0x9d, 0xc0, 0xe3,
	0x0c, 0x0b, 0x79, 0xbd, 0xa1, 0x22, 0x4f, 0x4a, 0x25, 0x8d, 0x24, 0x9e, 0x2a, 0x97, 0x85, 0xcc,
	0x90, 0xc7, 0x3f, 0xda, 0xd0, 0xa9, 0x34, 0x2a, 0x42, 0xa0, 0x23, 0x68, 0x81, 0xa1, 0x13, 0x39,
	0x63, 0x3f, 0xb5, 0x6b, 0x32, 0x02, 0xaf, 0xa4, 0x5a, 0x7f, 0x93, 0x2a, 0x0b, 0xdd, 0xc8, 0x19,
	0x77, 0xd3, 0xfb, 0xba, 0xe6, 0x2b, 0xc9, 0x31, 0x6c, 0xdb, 0x7d, 0xbb, 0x26, 0x09, 0x74, 0x39,
	0xde, 0x22, 0x0f, 0x3b, 0x91, 0x33, 0x3e, 0x9e, 0x84, 0xc9, 0xee, 0x98, 0xa4, 0x3e, 0x22, 0xf9,
	0x50, 0x63, 0x9f, 0x37, 0x25, 0xa6, 0x0d, 0x8d, 0x9c, 0x80, 0xcf, 0xf4, 0x75, 0xa6, 0xd8, 0x2d,
	0xaa, 0xb0, 0x1b, 0x39, 0x63, 0x2f, 0xf5, 0x98, 0x9e, 0xd9, 0x9a, 0x9c, 0xc2, 0x60, 0x2d, 0x0b,
	0xa4, 0x59, 0xa6, 0x50, 0xeb, 0xb0, 0x17, 0x39, 0xe3, 0xc1, 0x64, 0xf4, 0xc0, 0x72, 0x8b, 0x5a,
	0xd3, 0x43, 0xfa, 0x48, 0xc2, 0xe0, 0x00, 0x23, 0xa7, 0xd0, 0xdf, 0x19, 0x35, 0xbd, 0x45, 0x0f,
	0x8c, 0xce, 0x1a, 0xf4, 0xca, 0x28, 0x26, 0xf2, 0x5a, 0x72, 0xd1, 0x4a, 0x77, 0x12, 0x32, 0xdc,
	0x66, 0xe3, 0xd7, 0xb3, 0x5e, 0xb4, 0x9a, 0x74, 0xa6, 0x47, 0x00, 0x06, 0xb5, 0xb9, 0x96, 0x02,
	0xe5, 0x2a, 0x7e, 0x07, 0xfe, 0xfd, 0x7c, 0x64, 0x00, 0xfd, 0x19, 0xae, 0x68, 0xc5, 0x4d, 0xd0,
	0x22, 0x7d, 0x68, 0x5f, 0x55, 0x22, 0x70, 0x88, 0x0f, 0xdd, 0x39, 0x55, 0x66, 0x1d, 0xb8, 0xc4,
	0x83, 0xce, 0x47, 0x29, 0x45, 0xd0, 0xae, 0x57, 0x57, 0x86, 0xaa, 0xa0, 0x13, 0xbf, 0x85, 0x27,
	0x7f, 0x74, 0x51, 0x3b, 0xdd, 0x30, 0x91, 0x67, 0x54, 0x06, 0xad, 0xba, 0x58, 0x20, 0xfb, 0xc2,
	0x44, 0x1e, 0x38, 0xe4, 0x08, 0x3c, 0xbd, 0xa6, 0x22, 0x5f, 0x53, 0x16, 0xb8, 0xf1, 0x0c, 0xbc,
	0x5b, 0xca, 0x59, 0x46, 0x0d, 0x92, 0x21, 0x74, 0xef, 0x0a, 0x2e, 0xf4, 0xf6, 0x2e, 0x9b, 0x82,
	0xc4, 0xcd, 0x45, 0x87, 0x6e, 0xd4, 0x1e, 0x0f, 0x26, 0xc7, 0xbf, 0xcf, 0x9f, 0x5a, 0x2c, 0xfe,
	0xe9, 0x00, 0xa8, 0x72, 0x99, 0xe2, 0x4d, 0x85, 0xda, 0x90, 0x37, 0xe0, 0xcb, 0x12, 0x15, 0x35,
	0x4c, 0x0a, 0x6b, 0x76, 0x3c, 0x39, 0xd9, 0xeb, 0xf6, 0xc4, 0xe4, 0x53, 0x69, 0x6f, 0x60, 0xcf,
	0x26, 0x2f, 0xc0, 0xcd, 0x16, 0xf6, 0xd1, 0xfc, 0x4b, 0x33, 0x9b, 0x5a, 0x8d, 0x9b, 0x2d, 0x48,
	0xb2, 0x6f, 0xde, 0xbe, 0xa7, 0xc1, 0x84, 0xec, 0x25, 0x3b, 0x24, 0xbd, 0xe7, 0xc4, 0x27, 0xd0,
	0x6b, 0x4e, 0xac, 0xd3, 0x9b, 0x67, 0x6c, 0x9b, 0xf2, 0x39, 0x9a, 0xc0, 0x89, 0x5f, 0x42, 0xaf,
	0xb1, 0xae, 0xe3, 0x52, 0x95, 0x10, 0x75, 0x5c, 0x2d, 0xf2, 0x08, 0xfc, 0x25, 0x15, 0x99, 0x35,
	0x08, 0x9c, 0x1a, 0xd3, 0x86, 0x2a, 0x53, 0x95, 0x81, 0x1b, 0x2b, 0xf0, 0x6c, 0x63, 0x25, 0xdf,
	0x90, 0xa7, 0xd0, 0x53, 0xa8, 0x2b, 0x6e, 0xec, 0xc0, 0x5e, 0xba, 0xad, 0xc8, 0x33, 0x80, 0x15,
	0x65, 0x3c, 0x45, 0xaa, 0xa5, 0xb0, 0x83, 0xf9, 0xe9, 0xc1, 0xce, 0xff, 0xce, 0x30, 0x39, 0x03,
	0xef, 0x5c, 0x21, 0x1a, 0x26, 0x72, 0xf2, 0x1a, 0xfa, 0xf3, 0x3b, 0x5c, 0xa6, 0x97, 0xef, 0xc9,
	0xf0, 0x6f, 0x59, 0x8d, 0xc8, 0x83, 0xdd, 0x92, 0x6f, 0xe2, 0xd6, 0xf4, 0x12, 0x86, 0x4b, 0x59,
	0x24, 0x78, 0x47, 0x8b, 0x92, 0x63, 0x62, 0x2a, 0x23, 0x15, 0xa3, 0x7c, 0xba, 0x7b, 0x46, 0x53,
	0x29, 0xbf, 0x5e, 0xd6, 0xdf, 0xbb, 0xfe, 0xee, 0x3e, 0x3f, 0x97, 0x32, 0xe7, 0x98, 0xd8, 0x7a,
	0x51, 0xad, 0x92, 0x79, 0xa3, 0xd2, 0xc9, 0x01, 0x79, 0xd1, 0xb3, 0xbf, 0x87, 0x57, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x30, 0xb4, 0xc2, 0x87, 0x31, 0x04, 0x00, 0x00,
}