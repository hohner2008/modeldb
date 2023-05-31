// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.11.2
// source: uac/ResourceV2.proto

package uac

import (
	context "context"
	_ "github.com/VertaAI/modeldb/protos/gen/go/protos/public/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ResourceV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrganizationId string         `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	ResourceType   ResourceTypeV2 `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=ai.verta.uac.ResourceTypeV2" json:"resource_type,omitempty"`
}

func (x *ResourceV2) Reset() {
	*x = ResourceV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uac_ResourceV2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceV2) ProtoMessage() {}

func (x *ResourceV2) ProtoReflect() protoreflect.Message {
	mi := &file_uac_ResourceV2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceV2.ProtoReflect.Descriptor instead.
func (*ResourceV2) Descriptor() ([]byte, []int) {
	return file_uac_ResourceV2_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceV2) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResourceV2) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *ResourceV2) GetResourceType() ResourceTypeV2 {
	if x != nil {
		return x.ResourceType
	}
	return ResourceTypeV2_RESOURCE_TYPE_UNKNOWN
}

type GetResourcesV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           []string       `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	ResourceType ResourceTypeV2 `protobuf:"varint,2,opt,name=resource_type,json=resourceType,proto3,enum=ai.verta.uac.ResourceTypeV2" json:"resource_type,omitempty"`
}

func (x *GetResourcesV2) Reset() {
	*x = GetResourcesV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uac_ResourceV2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcesV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcesV2) ProtoMessage() {}

func (x *GetResourcesV2) ProtoReflect() protoreflect.Message {
	mi := &file_uac_ResourceV2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcesV2.ProtoReflect.Descriptor instead.
func (*GetResourcesV2) Descriptor() ([]byte, []int) {
	return file_uac_ResourceV2_proto_rawDescGZIP(), []int{1}
}

func (x *GetResourcesV2) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *GetResourcesV2) GetResourceType() ResourceTypeV2 {
	if x != nil {
		return x.ResourceType
	}
	return ResourceTypeV2_RESOURCE_TYPE_UNKNOWN
}

type GetResourcesV2_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*ResourceV2 `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *GetResourcesV2_Response) Reset() {
	*x = GetResourcesV2_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uac_ResourceV2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcesV2_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcesV2_Response) ProtoMessage() {}

func (x *GetResourcesV2_Response) ProtoReflect() protoreflect.Message {
	mi := &file_uac_ResourceV2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcesV2_Response.ProtoReflect.Descriptor instead.
func (*GetResourcesV2_Response) Descriptor() ([]byte, []int) {
	return file_uac_ResourceV2_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetResourcesV2_Response) GetResources() []*ResourceV2 {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_uac_ResourceV2_proto protoreflect.FileDescriptor

var file_uac_ResourceV2_proto_rawDesc = []byte{
	0x0a, 0x14, 0x75, 0x61, 0x63, 0x2f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x32,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61,
	0x2e, 0x75, 0x61, 0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x75, 0x61, 0x63, 0x2f, 0x55, 0x41, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x75, 0x61, 0x63, 0x2f, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x75, 0x61, 0x63, 0x2f, 0x54, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x75, 0x61, 0x63, 0x2f, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x75, 0x61, 0x63, 0x2f, 0x52, 0x6f, 0x6c,
	0x65, 0x56, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x56, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x41, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e,
	0x75, 0x61, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x56, 0x32, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x22, 0xa7, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x56, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x41, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x69, 0x2e,
	0x76, 0x65, 0x72, 0x74, 0x61, 0x2e, 0x75, 0x61, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x56, 0x32, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x42, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61,
	0x2e, 0x75, 0x61, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x32, 0x52,
	0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x32, 0x82, 0x01, 0x0a, 0x11, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x32,
	0x12, 0x6d, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x1c, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e, 0x75, 0x61, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x56, 0x32, 0x1a, 0x25,
	0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e, 0x75, 0x61, 0x63, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f,
	0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x42,
	0x3e, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x56, 0x65, 0x72, 0x74, 0x61, 0x41, 0x49, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x62, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x75, 0x61, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uac_ResourceV2_proto_rawDescOnce sync.Once
	file_uac_ResourceV2_proto_rawDescData = file_uac_ResourceV2_proto_rawDesc
)

func file_uac_ResourceV2_proto_rawDescGZIP() []byte {
	file_uac_ResourceV2_proto_rawDescOnce.Do(func() {
		file_uac_ResourceV2_proto_rawDescData = protoimpl.X.CompressGZIP(file_uac_ResourceV2_proto_rawDescData)
	})
	return file_uac_ResourceV2_proto_rawDescData
}

var file_uac_ResourceV2_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_uac_ResourceV2_proto_goTypes = []interface{}{
	(*ResourceV2)(nil),              // 0: ai.verta.uac.ResourceV2
	(*GetResourcesV2)(nil),          // 1: ai.verta.uac.GetResourcesV2
	(*GetResourcesV2_Response)(nil), // 2: ai.verta.uac.GetResourcesV2.Response
	(ResourceTypeV2)(0),             // 3: ai.verta.uac.ResourceTypeV2
}
var file_uac_ResourceV2_proto_depIdxs = []int32{
	3, // 0: ai.verta.uac.ResourceV2.resource_type:type_name -> ai.verta.uac.ResourceTypeV2
	3, // 1: ai.verta.uac.GetResourcesV2.resource_type:type_name -> ai.verta.uac.ResourceTypeV2
	0, // 2: ai.verta.uac.GetResourcesV2.Response.resources:type_name -> ai.verta.uac.ResourceV2
	1, // 3: ai.verta.uac.ResourceServiceV2.getResources:input_type -> ai.verta.uac.GetResourcesV2
	2, // 4: ai.verta.uac.ResourceServiceV2.getResources:output_type -> ai.verta.uac.GetResourcesV2.Response
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_uac_ResourceV2_proto_init() }
func file_uac_ResourceV2_proto_init() {
	if File_uac_ResourceV2_proto != nil {
		return
	}
	file_uac_UACService_proto_init()
	file_uac_Organization_proto_init()
	file_uac_Team_proto_init()
	file_uac_RoleService_proto_init()
	file_uac_RoleV2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_uac_ResourceV2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceV2); i {
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
		file_uac_ResourceV2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourcesV2); i {
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
		file_uac_ResourceV2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourcesV2_Response); i {
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
			RawDescriptor: file_uac_ResourceV2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_uac_ResourceV2_proto_goTypes,
		DependencyIndexes: file_uac_ResourceV2_proto_depIdxs,
		MessageInfos:      file_uac_ResourceV2_proto_msgTypes,
	}.Build()
	File_uac_ResourceV2_proto = out.File
	file_uac_ResourceV2_proto_rawDesc = nil
	file_uac_ResourceV2_proto_goTypes = nil
	file_uac_ResourceV2_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ResourceServiceV2Client is the client API for ResourceServiceV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceServiceV2Client interface {
	GetResources(ctx context.Context, in *GetResourcesV2, opts ...grpc.CallOption) (*GetResourcesV2_Response, error)
}

type resourceServiceV2Client struct {
	cc grpc.ClientConnInterface
}

func NewResourceServiceV2Client(cc grpc.ClientConnInterface) ResourceServiceV2Client {
	return &resourceServiceV2Client{cc}
}

func (c *resourceServiceV2Client) GetResources(ctx context.Context, in *GetResourcesV2, opts ...grpc.CallOption) (*GetResourcesV2_Response, error) {
	out := new(GetResourcesV2_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.uac.ResourceServiceV2/getResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceServiceV2Server is the server API for ResourceServiceV2 service.
type ResourceServiceV2Server interface {
	GetResources(context.Context, *GetResourcesV2) (*GetResourcesV2_Response, error)
}

// UnimplementedResourceServiceV2Server can be embedded to have forward compatible implementations.
type UnimplementedResourceServiceV2Server struct {
}

func (*UnimplementedResourceServiceV2Server) GetResources(context.Context, *GetResourcesV2) (*GetResourcesV2_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResources not implemented")
}

func RegisterResourceServiceV2Server(s *grpc.Server, srv ResourceServiceV2Server) {
	s.RegisterService(&_ResourceServiceV2_serviceDesc, srv)
}

func _ResourceServiceV2_GetResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourcesV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceV2Server).GetResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.uac.ResourceServiceV2/GetResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceV2Server).GetResources(ctx, req.(*GetResourcesV2))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceServiceV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.verta.uac.ResourceServiceV2",
	HandlerType: (*ResourceServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getResources",
			Handler:    _ResourceServiceV2_GetResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uac/ResourceV2.proto",
}
