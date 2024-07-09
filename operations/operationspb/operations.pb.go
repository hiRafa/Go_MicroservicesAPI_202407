// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: operations.proto

package operationspb

import (
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

type OperationsByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OperationsByIdRequest) Reset() {
	*x = OperationsByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationsByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationsByIdRequest) ProtoMessage() {}

func (x *OperationsByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationsByIdRequest.ProtoReflect.Descriptor instead.
func (*OperationsByIdRequest) Descriptor() ([]byte, []int) {
	return file_operations_proto_rawDescGZIP(), []int{0}
}

func (x *OperationsByIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateOperationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ShortName string `protobuf:"bytes,2,opt,name=shortName,proto3" json:"shortName,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateOperationsRequest) Reset() {
	*x = CreateOperationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOperationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOperationsRequest) ProtoMessage() {}

func (x *CreateOperationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOperationsRequest.ProtoReflect.Descriptor instead.
func (*CreateOperationsRequest) Descriptor() ([]byte, []int) {
	return file_operations_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOperationsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateOperationsRequest) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *CreateOperationsRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type OperationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ShortName string `protobuf:"bytes,3,opt,name=shortName,proto3" json:"shortName,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *OperationsResponse) Reset() {
	*x = OperationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationsResponse) ProtoMessage() {}

func (x *OperationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationsResponse.ProtoReflect.Descriptor instead.
func (*OperationsResponse) Descriptor() ([]byte, []int) {
	return file_operations_proto_rawDescGZIP(), []int{2}
}

func (x *OperationsResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OperationsResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OperationsResponse) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *OperationsResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_operations_proto protoreflect.FileDescriptor

var file_operations_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x27,
	0x0a, 0x15, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x6c, 0x0a, 0x12, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0xc4, 0x01, 0x0a, 0x11, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x21, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0f, 0x5a, 0x0d, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_operations_proto_rawDescOnce sync.Once
	file_operations_proto_rawDescData = file_operations_proto_rawDesc
)

func file_operations_proto_rawDescGZIP() []byte {
	file_operations_proto_rawDescOnce.Do(func() {
		file_operations_proto_rawDescData = protoimpl.X.CompressGZIP(file_operations_proto_rawDescData)
	})
	return file_operations_proto_rawDescData
}

var file_operations_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_operations_proto_goTypes = []any{
	(*OperationsByIdRequest)(nil),   // 0: operations.OperationsByIdRequest
	(*CreateOperationsRequest)(nil), // 1: operations.CreateOperationsRequest
	(*OperationsResponse)(nil),      // 2: operations.OperationsResponse
}
var file_operations_proto_depIdxs = []int32{
	0, // 0: operations.OperationsService.GetOperationsById:input_type -> operations.OperationsByIdRequest
	1, // 1: operations.OperationsService.CreateOperations:input_type -> operations.CreateOperationsRequest
	2, // 2: operations.OperationsService.GetOperationsById:output_type -> operations.OperationsResponse
	2, // 3: operations.OperationsService.CreateOperations:output_type -> operations.OperationsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_operations_proto_init() }
func file_operations_proto_init() {
	if File_operations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_operations_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*OperationsByIdRequest); i {
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
		file_operations_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOperationsRequest); i {
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
		file_operations_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*OperationsResponse); i {
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
			RawDescriptor: file_operations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_operations_proto_goTypes,
		DependencyIndexes: file_operations_proto_depIdxs,
		MessageInfos:      file_operations_proto_msgTypes,
	}.Build()
	File_operations_proto = out.File
	file_operations_proto_rawDesc = nil
	file_operations_proto_goTypes = nil
	file_operations_proto_depIdxs = nil
}