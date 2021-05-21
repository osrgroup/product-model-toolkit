// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package product

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Filename struct {
	FileName             string   `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filename) Reset()         { *m = Filename{} }
func (m *Filename) String() string { return proto.CompactTextString(m) }
func (*Filename) ProtoMessage()    {}
func (*Filename) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *Filename) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filename.Unmarshal(m, b)
}
func (m *Filename) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filename.Marshal(b, m, deterministic)
}
func (m *Filename) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filename.Merge(m, src)
}
func (m *Filename) XXX_Size() int {
	return xxx_messageInfo_Filename.Size(m)
}
func (m *Filename) XXX_DiscardUnknown() {
	xxx_messageInfo_Filename.DiscardUnknown(m)
}

var xxx_messageInfo_Filename proto.InternalMessageInfo

func (m *Filename) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type Product struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version              string       `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Vcs                  string       `protobuf:"bytes,3,opt,name=vcs,proto3" json:"vcs,omitempty"`
	Description          string       `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Comment              string       `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	HomePageUrl          string       `protobuf:"bytes,6,opt,name=home_page_url,json=homePageUrl,proto3" json:"home_page_url,omitempty"`
	ExternalRef          string       `protobuf:"bytes,7,opt,name=external_ref,json=externalRef,proto3" json:"external_ref,omitempty"`
	Components           []*Component `protobuf:"bytes,8,rep,name=components,proto3" json:"components,omitempty"`
	UsageTypes           string       `protobuf:"bytes,9,opt,name=usage_types,json=usageTypes,proto3" json:"usage_types,omitempty"`
	ClearingState        string       `protobuf:"bytes,10,opt,name=clearing_state,json=clearingState,proto3" json:"clearing_state,omitempty"`
	DepGraph             string       `protobuf:"bytes,11,opt,name=dep_graph,json=depGraph,proto3" json:"dep_graph,omitempty"`
	Infrastructure       string       `protobuf:"bytes,12,opt,name=infrastructure,proto3" json:"infrastructure,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Product) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Product) GetVcs() string {
	if m != nil {
		return m.Vcs
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Product) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Product) GetHomePageUrl() string {
	if m != nil {
		return m.HomePageUrl
	}
	return ""
}

func (m *Product) GetExternalRef() string {
	if m != nil {
		return m.ExternalRef
	}
	return ""
}

func (m *Product) GetComponents() []*Component {
	if m != nil {
		return m.Components
	}
	return nil
}

func (m *Product) GetUsageTypes() string {
	if m != nil {
		return m.UsageTypes
	}
	return ""
}

func (m *Product) GetClearingState() string {
	if m != nil {
		return m.ClearingState
	}
	return ""
}

func (m *Product) GetDepGraph() string {
	if m != nil {
		return m.DepGraph
	}
	return ""
}

func (m *Product) GetInfrastructure() string {
	if m != nil {
		return m.Infrastructure
	}
	return ""
}

type Component struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Package              string   `protobuf:"bytes,3,opt,name=package,proto3" json:"package,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	License              *License `protobuf:"bytes,5,opt,name=license,proto3" json:"license,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Component) Reset()         { *m = Component{} }
func (m *Component) String() string { return proto.CompactTextString(m) }
func (*Component) ProtoMessage()    {}
func (*Component) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{2}
}

func (m *Component) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Component.Unmarshal(m, b)
}
func (m *Component) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Component.Marshal(b, m, deterministic)
}
func (m *Component) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Component.Merge(m, src)
}
func (m *Component) XXX_Size() int {
	return xxx_messageInfo_Component.Size(m)
}
func (m *Component) XXX_DiscardUnknown() {
	xxx_messageInfo_Component.DiscardUnknown(m)
}

var xxx_messageInfo_Component proto.InternalMessageInfo

func (m *Component) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Component) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Component) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *Component) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Component) GetLicense() *License {
	if m != nil {
		return m.License
	}
	return nil
}

type License struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeclaredLicense      string   `protobuf:"bytes,2,opt,name=declared_license,json=declaredLicense,proto3" json:"declared_license,omitempty"`
	ConcludedLicense     string   `protobuf:"bytes,3,opt,name=concluded_license,json=concludedLicense,proto3" json:"concluded_license,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *License) Reset()         { *m = License{} }
func (m *License) String() string { return proto.CompactTextString(m) }
func (*License) ProtoMessage()    {}
func (*License) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{3}
}

func (m *License) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_License.Unmarshal(m, b)
}
func (m *License) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_License.Marshal(b, m, deterministic)
}
func (m *License) XXX_Merge(src proto.Message) {
	xxx_messageInfo_License.Merge(m, src)
}
func (m *License) XXX_Size() int {
	return xxx_messageInfo_License.Size(m)
}
func (m *License) XXX_DiscardUnknown() {
	xxx_messageInfo_License.DiscardUnknown(m)
}

var xxx_messageInfo_License proto.InternalMessageInfo

func (m *License) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *License) GetDeclaredLicense() string {
	if m != nil {
		return m.DeclaredLicense
	}
	return ""
}

func (m *License) GetConcludedLicense() string {
	if m != nil {
		return m.ConcludedLicense
	}
	return ""
}

func init() {
	proto.RegisterType((*Filename)(nil), "product.Filename")
	proto.RegisterType((*Product)(nil), "product.Product")
	proto.RegisterType((*Component)(nil), "product.Component")
	proto.RegisterType((*License)(nil), "product.License")
}

func init() {
	proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5)
}

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0xb5, 0x6d, 0xd9, 0xb4, 0x93, 0x6d, 0x09, 0x3e, 0x59, 0xe2, 0x40, 0x89, 0x04, 0x2c,
	0x20, 0xed, 0xa1, 0x3c, 0x02, 0x12, 0x5c, 0x10, 0x5a, 0x75, 0xe1, 0x1c, 0x19, 0x7b, 0x9a, 0xb5,
	0x70, 0x6c, 0x63, 0x3b, 0x15, 0x3c, 0x02, 0x27, 0x5e, 0x19, 0xd9, 0x8e, 0xab, 0xc2, 0xde, 0xf2,
	0x7f, 0xf3, 0xcf, 0x64, 0x34, 0xbf, 0x61, 0x6d, 0x9d, 0x11, 0x23, 0x0f, 0x37, 0xd6, 0x99, 0x60,
	0x48, 0x35, 0xc9, 0xf6, 0x15, 0x2c, 0x3f, 0x48, 0x85, 0x9a, 0x0d, 0x48, 0x9e, 0xc2, 0xea, 0x20,
	0x15, 0x76, 0x51, 0xd0, 0x8b, 0xed, 0xc5, 0xf5, 0x6a, 0xbf, 0x8c, 0xe0, 0x33, 0x1b, 0xb0, 0xfd,
	0x3d, 0x87, 0xea, 0x36, 0x37, 0x91, 0x0d, 0xcc, 0xa4, 0x98, 0x1c, 0x33, 0x29, 0x08, 0x85, 0xea,
	0x88, 0xce, 0x4b, 0xa3, 0xe9, 0x2c, 0xc1, 0x22, 0x49, 0x03, 0xf3, 0x23, 0xf7, 0x74, 0x9e, 0x68,
	0xfc, 0x24, 0x5b, 0xa8, 0x05, 0x7a, 0xee, 0xa4, 0x0d, 0xd1, 0xbf, 0x48, 0x95, 0x73, 0x14, 0xa7,
	0x71, 0x33, 0x0c, 0xa8, 0x03, 0x7d, 0x94, 0xa7, 0x4d, 0x92, 0xb4, 0xb0, 0xbe, 0x37, 0x03, 0x76,
	0x96, 0xf5, 0xd8, 0x8d, 0x4e, 0xd1, 0xcb, 0xdc, 0x1d, 0xe1, 0x2d, 0xeb, 0xf1, 0xab, 0x53, 0xe4,
	0x39, 0x5c, 0xe1, 0xcf, 0x80, 0x4e, 0x33, 0xd5, 0x39, 0x3c, 0xd0, 0x2a, 0x5b, 0x0a, 0xdb, 0xe3,
	0x81, 0xec, 0x00, 0xb8, 0x19, 0xac, 0xd1, 0xa8, 0x83, 0xa7, 0xcb, 0xed, 0xfc, 0xba, 0xde, 0x91,
	0x9b, 0x72, 0xa0, 0xf7, 0xa5, 0xb4, 0x3f, 0x73, 0x91, 0x67, 0x50, 0x8f, 0x3e, 0xfe, 0x36, 0xfc,
	0xb2, 0xe8, 0xe9, 0x2a, 0x4d, 0x85, 0x84, 0xbe, 0x44, 0x42, 0x5e, 0xc0, 0x86, 0x2b, 0x64, 0x4e,
	0xea, 0xbe, 0xf3, 0x81, 0x05, 0xa4, 0x90, 0x3c, 0xeb, 0x42, 0xef, 0x22, 0x8c, 0x37, 0x16, 0x68,
	0xbb, 0xde, 0x31, 0x7b, 0x4f, 0xeb, 0x7c, 0x63, 0x81, 0xf6, 0x63, 0xd4, 0xe4, 0x25, 0x6c, 0xa4,
	0x3e, 0x38, 0xe6, 0x83, 0x1b, 0x79, 0x18, 0x1d, 0xd2, 0xab, 0xe4, 0xf8, 0x8f, 0xb6, 0x7f, 0x2e,
	0x60, 0x75, 0x5a, 0xf3, 0x41, 0x1a, 0x04, 0x16, 0x29, 0xc1, 0x1c, 0x45, 0xfa, 0x8e, 0x37, 0xb5,
	0x8c, 0x7f, 0x67, 0x3d, 0x4e, 0x59, 0x14, 0x79, 0x9e, 0xdd, 0xe2, 0xdf, 0xec, 0xde, 0x40, 0xa5,
	0x24, 0x47, 0xed, 0x31, 0xe5, 0x50, 0xef, 0x9a, 0xd3, 0x8d, 0x3e, 0x65, 0xbe, 0x2f, 0x86, 0xf6,
	0x07, 0x54, 0x13, 0x7b, 0xb0, 0xce, 0x6b, 0x68, 0x04, 0x72, 0xc5, 0x1c, 0x8a, 0xae, 0xcc, 0xcb,
	0xab, 0x3d, 0x2e, 0xbc, 0xb4, 0xbe, 0x85, 0x27, 0xdc, 0x68, 0xae, 0x46, 0x71, 0xe6, 0xcd, 0xfb,
	0x36, 0xa7, 0xc2, 0x64, 0xde, 0x35, 0xb0, 0x99, 0xde, 0xe3, 0x1d, 0xba, 0xa3, 0xe4, 0xf8, 0xed,
	0x32, 0xbd, 0xed, 0x77, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xbc, 0x3d, 0x6e, 0xec, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "product.proto",
}
