// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: Product.proto

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

type CreatContainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alias   string `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
	Images  string `protobuf:"bytes,2,opt,name=images,proto3" json:"images,omitempty"`
	Port1   string `protobuf:"bytes,3,opt,name=port1,proto3" json:"port1,omitempty"`
	Port2   string `protobuf:"bytes,4,opt,name=port2,proto3" json:"port2,omitempty"`
	SqlPort string `protobuf:"bytes,5,opt,name=SqlPort,proto3" json:"SqlPort,omitempty"`
}

func (x *CreatContainRequest) Reset() {
	*x = CreatContainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatContainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatContainRequest) ProtoMessage() {}

func (x *CreatContainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatContainRequest.ProtoReflect.Descriptor instead.
func (*CreatContainRequest) Descriptor() ([]byte, []int) {
	return file_Product_proto_rawDescGZIP(), []int{0}
}

func (x *CreatContainRequest) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *CreatContainRequest) GetImages() string {
	if x != nil {
		return x.Images
	}
	return ""
}

func (x *CreatContainRequest) GetPort1() string {
	if x != nil {
		return x.Port1
	}
	return ""
}

func (x *CreatContainRequest) GetPort2() string {
	if x != nil {
		return x.Port2
	}
	return ""
}

func (x *CreatContainRequest) GetSqlPort() string {
	if x != nil {
		return x.SqlPort
	}
	return ""
}

type CreatContainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID  string `protobuf:"bytes,1,opt,name=containerID,proto3" json:"containerID,omitempty"`
	ContainerUrl string `protobuf:"bytes,2,opt,name=containerUrl,proto3" json:"containerUrl,omitempty"`
}

func (x *CreatContainResponse) Reset() {
	*x = CreatContainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatContainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatContainResponse) ProtoMessage() {}

func (x *CreatContainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatContainResponse.ProtoReflect.Descriptor instead.
func (*CreatContainResponse) Descriptor() ([]byte, []int) {
	return file_Product_proto_rawDescGZIP(), []int{1}
}

func (x *CreatContainResponse) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *CreatContainResponse) GetContainerUrl() string {
	if x != nil {
		return x.ContainerUrl
	}
	return ""
}

type StartContainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string `protobuf:"bytes,1,opt,name=containerID,proto3" json:"containerID,omitempty"`
}

func (x *StartContainRequest) Reset() {
	*x = StartContainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartContainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartContainRequest) ProtoMessage() {}

func (x *StartContainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartContainRequest.ProtoReflect.Descriptor instead.
func (*StartContainRequest) Descriptor() ([]byte, []int) {
	return file_Product_proto_rawDescGZIP(), []int{2}
}

func (x *StartContainRequest) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

type StartContainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainUrl string `protobuf:"bytes,1,opt,name=ContainUrl,proto3" json:"ContainUrl,omitempty"`
}

func (x *StartContainResponse) Reset() {
	*x = StartContainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartContainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartContainResponse) ProtoMessage() {}

func (x *StartContainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartContainResponse.ProtoReflect.Descriptor instead.
func (*StartContainResponse) Descriptor() ([]byte, []int) {
	return file_Product_proto_rawDescGZIP(), []int{3}
}

func (x *StartContainResponse) GetContainUrl() string {
	if x != nil {
		return x.ContainUrl
	}
	return ""
}

var File_Product_proto protoreflect.FileDescriptor

var file_Product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x67, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x89, 0x01, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x32,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x71, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x53, 0x71, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x5c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x37, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x22, 0x36,
	0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x32, 0xc1, 0x01, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x57, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x2e,
	0x67, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x67, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x58, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Product_proto_rawDescOnce sync.Once
	file_Product_proto_rawDescData = file_Product_proto_rawDesc
)

func file_Product_proto_rawDescGZIP() []byte {
	file_Product_proto_rawDescOnce.Do(func() {
		file_Product_proto_rawDescData = protoimpl.X.CompressGZIP(file_Product_proto_rawDescData)
	})
	return file_Product_proto_rawDescData
}

var file_Product_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Product_proto_goTypes = []interface{}{
	(*CreatContainRequest)(nil),  // 0: gRPCServer.CreatContainRequest
	(*CreatContainResponse)(nil), // 1: gRPCServer.CreatContainResponse
	(*StartContainRequest)(nil),  // 2: gRPCServer.StartContainRequest
	(*StartContainResponse)(nil), // 3: gRPCServer.StartContainResponse
}
var file_Product_proto_depIdxs = []int32{
	2, // 0: gRPCServer.DockerServer.GetStartContainUrl:input_type -> gRPCServer.StartContainRequest
	0, // 1: gRPCServer.DockerServer.CreatContainMessage:input_type -> gRPCServer.CreatContainRequest
	3, // 2: gRPCServer.DockerServer.GetStartContainUrl:output_type -> gRPCServer.StartContainResponse
	1, // 3: gRPCServer.DockerServer.CreatContainMessage:output_type -> gRPCServer.CreatContainResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Product_proto_init() }
func file_Product_proto_init() {
	if File_Product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatContainRequest); i {
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
		file_Product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatContainResponse); i {
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
		file_Product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartContainRequest); i {
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
		file_Product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartContainResponse); i {
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
			RawDescriptor: file_Product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Product_proto_goTypes,
		DependencyIndexes: file_Product_proto_depIdxs,
		MessageInfos:      file_Product_proto_msgTypes,
	}.Build()
	File_Product_proto = out.File
	file_Product_proto_rawDesc = nil
	file_Product_proto_goTypes = nil
	file_Product_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DockerServerClient is the client API for DockerServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DockerServerClient interface {
	//启动容器获取容器URL请求
	GetStartContainUrl(ctx context.Context, in *StartContainRequest, opts ...grpc.CallOption) (*StartContainResponse, error)
	//创建容器并获得容器信息以及URL
	CreatContainMessage(ctx context.Context, in *CreatContainRequest, opts ...grpc.CallOption) (*CreatContainResponse, error)
}

type dockerServerClient struct {
	cc grpc.ClientConnInterface
}

func NewDockerServerClient(cc grpc.ClientConnInterface) DockerServerClient {
	return &dockerServerClient{cc}
}

func (c *dockerServerClient) GetStartContainUrl(ctx context.Context, in *StartContainRequest, opts ...grpc.CallOption) (*StartContainResponse, error) {
	out := new(StartContainResponse)
	err := c.cc.Invoke(ctx, "/gRPCServer.DockerServer/GetStartContainUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServerClient) CreatContainMessage(ctx context.Context, in *CreatContainRequest, opts ...grpc.CallOption) (*CreatContainResponse, error) {
	out := new(CreatContainResponse)
	err := c.cc.Invoke(ctx, "/gRPCServer.DockerServer/CreatContainMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DockerServerServer is the server API for DockerServer service.
type DockerServerServer interface {
	//启动容器获取容器URL请求
	GetStartContainUrl(context.Context, *StartContainRequest) (*StartContainResponse, error)
	//创建容器并获得容器信息以及URL
	CreatContainMessage(context.Context, *CreatContainRequest) (*CreatContainResponse, error)
}

// UnimplementedDockerServerServer can be embedded to have forward compatible implementations.
type UnimplementedDockerServerServer struct {
}

func (*UnimplementedDockerServerServer) GetStartContainUrl(context.Context, *StartContainRequest) (*StartContainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStartContainUrl not implemented")
}
func (*UnimplementedDockerServerServer) CreatContainMessage(context.Context, *CreatContainRequest) (*CreatContainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatContainMessage not implemented")
}

func RegisterDockerServerServer(s *grpc.Server, srv DockerServerServer) {
	s.RegisterService(&_DockerServer_serviceDesc, srv)
}

func _DockerServer_GetStartContainUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartContainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServerServer).GetStartContainUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gRPCServer.DockerServer/GetStartContainUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServerServer).GetStartContainUrl(ctx, req.(*StartContainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerServer_CreatContainMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatContainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServerServer).CreatContainMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gRPCServer.DockerServer/CreatContainMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServerServer).CreatContainMessage(ctx, req.(*CreatContainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DockerServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gRPCServer.DockerServer",
	HandlerType: (*DockerServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStartContainUrl",
			Handler:    _DockerServer_GetStartContainUrl_Handler,
		},
		{
			MethodName: "CreatContainMessage",
			Handler:    _DockerServer_CreatContainMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Product.proto",
}
