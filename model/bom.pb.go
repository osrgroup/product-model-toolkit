// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.0
// source: bom.proto

package model

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InputType int32

const (
	InputType_SPDX   InputType = 0 // create BoM as SPDX
	InputType_HUMAN  InputType = 1 // create BoM as human readable
	InputType_CUSTOM InputType = 2 // create BoM as custom report
)

// Enum value maps for InputType.
var (
	InputType_name = map[int32]string{
		0: "SPDX",
		1: "HUMAN",
		2: "CUSTOM",
	}
	InputType_value = map[string]int32{
		"SPDX":   0,
		"HUMAN":  1,
		"CUSTOM": 2,
	}
)

func (x InputType) Enum() *InputType {
	p := new(InputType)
	*p = x
	return p
}

func (x InputType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InputType) Descriptor() protoreflect.EnumDescriptor {
	return file_bom_proto_enumTypes[0].Descriptor()
}

func (InputType) Type() protoreflect.EnumType {
	return &file_bom_proto_enumTypes[0]
}

func (x InputType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InputType.Descriptor instead.
func (InputType) EnumDescriptor() ([]byte, []int) {
	return file_bom_proto_rawDescGZIP(), []int{0}
}

type InputValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string    `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Type     InputType `protobuf:"varint,2,opt,name=type,proto3,enum=product.InputType" json:"type,omitempty"`
}

func (x *InputValue) Reset() {
	*x = InputValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputValue) ProtoMessage() {}

func (x *InputValue) ProtoReflect() protoreflect.Message {
	mi := &file_bom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputValue.ProtoReflect.Descriptor instead.
func (*InputValue) Descriptor() ([]byte, []int) {
	return file_bom_proto_rawDescGZIP(), []int{0}
}

func (x *InputValue) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *InputValue) GetType() InputType {
	if x != nil {
		return x.Type
	}
	return InputType_SPDX
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created bool     `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Product *Product `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_bom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_bom_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

func (x *Response) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

var File_bom_proto protoreflect.FileDescriptor

var file_bom_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x1a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0a, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x50, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2a, 0x2c, 0x0a, 0x09, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x50, 0x44, 0x58, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x48, 0x55, 0x4d, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55,
	0x53, 0x54, 0x4f, 0x4d, 0x10, 0x02, 0x32, 0x43, 0x0a, 0x0a, 0x42, 0x6f, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6d, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bom_proto_rawDescOnce sync.Once
	file_bom_proto_rawDescData = file_bom_proto_rawDesc
)

func file_bom_proto_rawDescGZIP() []byte {
	file_bom_proto_rawDescOnce.Do(func() {
		file_bom_proto_rawDescData = protoimpl.X.CompressGZIP(file_bom_proto_rawDescData)
	})
	return file_bom_proto_rawDescData
}

var file_bom_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bom_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bom_proto_goTypes = []interface{}{
	(InputType)(0),     // 0: product.InputType
	(*InputValue)(nil), // 1: product.InputValue
	(*Response)(nil),   // 2: product.Response
	(*Product)(nil),    // 3: product.Product
}
var file_bom_proto_depIdxs = []int32{
	0, // 0: product.InputValue.type:type_name -> product.InputType
	3, // 1: product.Response.product:type_name -> product.Product
	1, // 2: product.BomService.CreateBom:input_type -> product.InputValue
	2, // 3: product.BomService.CreateBom:output_type -> product.Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bom_proto_init() }
func file_bom_proto_init() {
	if File_bom_proto != nil {
		return
	}
	file_product_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bom_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bom_proto_goTypes,
		DependencyIndexes: file_bom_proto_depIdxs,
		EnumInfos:         file_bom_proto_enumTypes,
		MessageInfos:      file_bom_proto_msgTypes,
	}.Build()
	File_bom_proto = out.File
	file_bom_proto_rawDesc = nil
	file_bom_proto_goTypes = nil
	file_bom_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BomServiceClient is the client API for BomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BomServiceClient interface {
	CreateBom(ctx context.Context, in *InputValue, opts ...grpc.CallOption) (*Response, error)
}

type bomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBomServiceClient(cc grpc.ClientConnInterface) BomServiceClient {
	return &bomServiceClient{cc}
}

func (c *bomServiceClient) CreateBom(ctx context.Context, in *InputValue, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/product.BomService/CreateBom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BomServiceServer is the server API for BomService service.
type BomServiceServer interface {
	CreateBom(context.Context, *InputValue) (*Response, error)
}

// UnimplementedBomServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBomServiceServer struct {
}

func (*UnimplementedBomServiceServer) CreateBom(context.Context, *InputValue) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBom not implemented")
}

func RegisterBomServiceServer(s *grpc.Server, srv BomServiceServer) {
	s.RegisterService(&_BomService_serviceDesc, srv)
}

func _BomService_CreateBom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BomServiceServer).CreateBom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.BomService/CreateBom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BomServiceServer).CreateBom(ctx, req.(*InputValue))
	}
	return interceptor(ctx, in, info, handler)
}

var _BomService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.BomService",
	HandlerType: (*BomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBom",
			Handler:    _BomService_CreateBom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bom.proto",
}
