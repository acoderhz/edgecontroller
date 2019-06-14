// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// Identity defines a request to obtain authentication credentials. These
// credentials would be used to further communicate with endpoint(s) that are
// protected by a form of authentication.
//
// Any strings that are annotated as "PEM-encoded" implies that encoding format
// is used, with any newlines indicated with `\n` characters. Most languages
// provide encoders that correctly marshal this out. For more information,
// see the RFC here: https://tools.ietf.org/html/rfc7468
type Identity struct {
	// A PEM-encoded certificate signing request (CSR)
	Csr                  string   `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identity) Reset()         { *m = Identity{} }
func (m *Identity) String() string { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()    {}
func (*Identity) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *Identity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity.Unmarshal(m, b)
}
func (m *Identity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity.Marshal(b, m, deterministic)
}
func (m *Identity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity.Merge(m, src)
}
func (m *Identity) XXX_Size() int {
	return xxx_messageInfo_Identity.Size(m)
}
func (m *Identity) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity.DiscardUnknown(m)
}

var xxx_messageInfo_Identity proto.InternalMessageInfo

func (m *Identity) GetCsr() string {
	if m != nil {
		return m.Csr
	}
	return ""
}

// Credentials defines a response for a request to obtain authentication
// credentials. These credentials may be used to further communicate with
// endpoint(s) that are protected by a form of authentication.
//
// Any strings that are annotated as "PEM-encoded" implies that encoding format
// is used, with any newlines indicated with `\n` characters. Most languages
// provide encoders that correctly marshal this out. For more information,
// see the RFC here: https://tools.ietf.org/html/rfc7468
type Credentials struct {
	// An identifier for the set of credentials
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A PEM-encoded signed certificate for the CSR
	Certificate string `protobuf:"bytes,2,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// A PEM-encoded certificate chain, starting with the issuing CA and
	// ending with the root CA (inclusive)
	CaChain []string `protobuf:"bytes,3,rep,name=ca_chain,json=caChain,proto3" json:"ca_chain,omitempty"`
	// A PEM-encoded CAs to be added to the client's CA pool
	CaPool               []string `protobuf:"bytes,4,rep,name=ca_pool,json=caPool,proto3" json:"ca_pool,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credentials) Reset()         { *m = Credentials{} }
func (m *Credentials) String() string { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()    {}
func (*Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *Credentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Credentials.Unmarshal(m, b)
}
func (m *Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Credentials.Marshal(b, m, deterministic)
}
func (m *Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credentials.Merge(m, src)
}
func (m *Credentials) XXX_Size() int {
	return xxx_messageInfo_Credentials.Size(m)
}
func (m *Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_Credentials proto.InternalMessageInfo

func (m *Credentials) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Credentials) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func (m *Credentials) GetCaChain() []string {
	if m != nil {
		return m.CaChain
	}
	return nil
}

func (m *Credentials) GetCaPool() []string {
	if m != nil {
		return m.CaPool
	}
	return nil
}

func init() {
	proto.RegisterType((*Identity)(nil), "openness.auth.Identity")
	proto.RegisterType((*Credentials)(nil), "openness.auth.Credentials")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0x94, 0x6d, 0x9d, 0x0b, 0xd3, 0x64, 0x09, 0x56, 0xa2, 0x1e, 0xac, 0xc0, 0x61,
	0x2a, 0x34, 0x6e, 0xcb, 0x4e, 0xe5, 0x42, 0x56, 0xf5, 0x50, 0x34, 0xa1, 0xaa, 0x15, 0x17, 0x2e,
	0x93, 0xeb, 0xbc, 0x25, 0x86, 0xd4, 0x36, 0xb6, 0xc3, 0x04, 0x07, 0x0e, 0x7c, 0x83, 0x8d, 0x0f,
	0xc1, 0x07, 0xe2, 0xcc, 0x8d, 0x03, 0x17, 0xbe, 0x03, 0x72, 0x28, 0xa2, 0x30, 0x6d, 0xa7, 0x38,
	0xef, 0xf7, 0xcb, 0x7b, 0x2f, 0x7f, 0x19, 0x21, 0x56, 0xb9, 0x22, 0xd1, 0x46, 0x39, 0x85, 0xef,
	0x28, 0x0d, 0x52, 0x82, 0xb5, 0x89, 0x2f, 0x46, 0x9d, 0x5c, 0xa9, 0xbc, 0x04, 0xca, 0xb4, 0xa0,
	0x4c, 0x4a, 0xe5, 0x98, 0x13, 0x4a, 0xda, 0xdf, 0x72, 0xf4, 0xb8, 0x7e, 0xf0, 0x5e, 0x0e, 0xb2,
	0x67, 0xcf, 0x59, 0x9e, 0x83, 0xa1, 0x4a, 0xd7, 0xc6, 0x55, 0x3b, 0xee, 0xa0, 0xe6, 0x34, 0x03,
	0xe9, 0x84, 0x7b, 0x8f, 0xf7, 0x51, 0x83, 0x5b, 0xd3, 0x0e, 0x48, 0x70, 0xb8, 0x3b, 0xf7, 0xc7,
	0xd8, 0xa2, 0xd6, 0xd8, 0x40, 0xcd, 0x59, 0x69, 0xf1, 0x1e, 0x0a, 0x45, 0xb6, 0xe6, 0xa1, 0xc8,
	0x30, 0x41, 0x2d, 0x0e, 0xc6, 0x89, 0x33, 0xc1, 0x99, 0x83, 0x76, 0x58, 0x83, 0xcd, 0x12, 0xbe,
	0x8f, 0x9a, 0x9c, 0x9d, 0xf2, 0x82, 0x09, 0xd9, 0x6e, 0x90, 0xc6, 0xe1, 0xee, 0x7c, 0x87, 0xb3,
	0xb1, 0x7f, 0xc5, 0x07, 0x68, 0x87, 0xb3, 0x53, 0xad, 0x54, 0xd9, 0xbe, 0x55, 0x93, 0x6d, 0xce,
	0x66, 0x4a, 0x95, 0xc3, 0x9f, 0x01, 0x6a, 0xa5, 0x95, 0x2b, 0x16, 0x60, 0xde, 0x09, 0x0e, 0xf8,
	0x5b, 0x80, 0xf0, 0x1c, 0xde, 0x56, 0x60, 0xdd, 0xe6, 0x32, 0x07, 0xc9, 0x3f, 0xa9, 0x24, 0x7f,
	0x7e, 0x23, 0x8a, 0xfe, 0x03, 0x1b, 0x1f, 0xc5, 0x17, 0xc1, 0x65, 0xfa, 0x31, 0x8a, 0xd7, 0xed,
	0x88, 0xe7, 0x1e, 0xf1, 0x3a, 0x13, 0xc2, 0xff, 0x9a, 0xcf, 0x1f, 0xa1, 0xc6, 0xb0, 0x3f, 0xc0,
	0x0f, 0x51, 0x9c, 0x5e, 0x2b, 0xf9, 0x33, 0x73, 0x90, 0x79, 0xf9, 0xa8, 0x7f, 0xe4, 0xe5, 0x75,
	0x67, 0xc8, 0x88, 0x58, 0xef, 0x43, 0xa4, 0x72, 0xe4, 0x8d, 0x54, 0xe7, 0x92, 0x9e, 0xa9, 0x4a,
	0x66, 0x9f, 0xbe, 0x7e, 0xff, 0x1c, 0xa2, 0x78, 0x8b, 0xfa, 0xe1, 0xa3, 0xa0, 0x7b, 0x7c, 0x11,
	0x5e, 0xa6, 0x3f, 0x02, 0xfc, 0x25, 0x40, 0x4d, 0x3f, 0x8a, 0xa4, 0xb3, 0x69, 0x7c, 0x8c, 0xd0,
	0x62, 0xc5, 0x8c, 0x23, 0x93, 0x2c, 0x07, 0xdc, 0xc9, 0x85, 0x2b, 0xaa, 0x65, 0xc2, 0xd5, 0x8a,
	0x5a, 0x5f, 0x86, 0x2c, 0x87, 0x15, 0xf0, 0xba, 0x45, 0x74, 0xcf, 0x56, 0x5a, 0x2b, 0xe3, 0x9e,
	0xd5, 0xa8, 0xe7, 0x99, 0x37, 0xbb, 0x33, 0x84, 0x53, 0xcd, 0x78, 0x01, 0x64, 0x98, 0xf4, 0xc9,
	0x89, 0xe0, 0x20, 0x2d, 0xe0, 0x51, 0xe1, 0x9c, 0xb6, 0x23, 0x4a, 0xaf, 0xeb, 0x69, 0x79, 0x01,
	0x2b, 0x46, 0x97, 0xa5, 0x5a, 0xd2, 0x15, 0xb3, 0x0e, 0x0c, 0x3d, 0x99, 0x8e, 0x27, 0x2f, 0x16,
	0x93, 0xe1, 0xd6, 0x20, 0xe9, 0x27, 0xfd, 0x6e, 0x10, 0x0e, 0xf7, 0x99, 0xd6, 0xe5, 0x3a, 0x11,
	0xfa, 0xda, 0x2a, 0x39, 0xba, 0x52, 0x99, 0xdf, 0xf5, 0xa1, 0x0c, 0xf0, 0x1e, 0xba, 0xfd, 0x52,
	0xfa, 0x45, 0x95, 0x11, 0x1f, 0x20, 0x7b, 0xf5, 0xe0, 0xe6, 0xc1, 0x4f, 0xbd, 0xba, 0xdc, 0xae,
	0x6f, 0xe7, 0x93, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x6a, 0x38, 0x7f, 0x06, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	RequestCredentials(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Credentials, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) RequestCredentials(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Credentials, error) {
	out := new(Credentials)
	err := c.cc.Invoke(ctx, "/openness.auth.AuthService/RequestCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	RequestCredentials(context.Context, *Identity) (*Credentials, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_RequestCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RequestCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.auth.AuthService/RequestCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RequestCredentials(ctx, req.(*Identity))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openness.auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestCredentials",
			Handler:    _AuthService_RequestCredentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}