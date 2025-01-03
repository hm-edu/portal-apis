// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: eab-mgmt.proto

package portal_apis

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

type CheckEABPermissionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Domains       []string               `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty"`
	EabKey        string                 `protobuf:"bytes,2,opt,name=eab_key,json=eabKey,proto3" json:"eab_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckEABPermissionRequest) Reset() {
	*x = CheckEABPermissionRequest{}
	mi := &file_eab_mgmt_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckEABPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEABPermissionRequest) ProtoMessage() {}

func (x *CheckEABPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eab_mgmt_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEABPermissionRequest.ProtoReflect.Descriptor instead.
func (*CheckEABPermissionRequest) Descriptor() ([]byte, []int) {
	return file_eab_mgmt_proto_rawDescGZIP(), []int{0}
}

func (x *CheckEABPermissionRequest) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *CheckEABPermissionRequest) GetEabKey() string {
	if x != nil {
		return x.EabKey
	}
	return ""
}

type CheckEABPermissionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Missing       []string               `protobuf:"bytes,1,rep,name=missing,proto3" json:"missing,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckEABPermissionResponse) Reset() {
	*x = CheckEABPermissionResponse{}
	mi := &file_eab_mgmt_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckEABPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEABPermissionResponse) ProtoMessage() {}

func (x *CheckEABPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eab_mgmt_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEABPermissionResponse.ProtoReflect.Descriptor instead.
func (*CheckEABPermissionResponse) Descriptor() ([]byte, []int) {
	return file_eab_mgmt_proto_rawDescGZIP(), []int{1}
}

func (x *CheckEABPermissionResponse) GetMissing() []string {
	if x != nil {
		return x.Missing
	}
	return nil
}

type ResolveAccountIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountId     string                 `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResolveAccountIdRequest) Reset() {
	*x = ResolveAccountIdRequest{}
	mi := &file_eab_mgmt_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResolveAccountIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveAccountIdRequest) ProtoMessage() {}

func (x *ResolveAccountIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eab_mgmt_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveAccountIdRequest.ProtoReflect.Descriptor instead.
func (*ResolveAccountIdRequest) Descriptor() ([]byte, []int) {
	return file_eab_mgmt_proto_rawDescGZIP(), []int{2}
}

func (x *ResolveAccountIdRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type ResolveAccountIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EabKey        string                 `protobuf:"bytes,1,opt,name=eab_key,json=eabKey,proto3" json:"eab_key,omitempty"`
	User          string                 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResolveAccountIdResponse) Reset() {
	*x = ResolveAccountIdResponse{}
	mi := &file_eab_mgmt_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResolveAccountIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveAccountIdResponse) ProtoMessage() {}

func (x *ResolveAccountIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eab_mgmt_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveAccountIdResponse.ProtoReflect.Descriptor instead.
func (*ResolveAccountIdResponse) Descriptor() ([]byte, []int) {
	return file_eab_mgmt_proto_rawDescGZIP(), []int{3}
}

func (x *ResolveAccountIdResponse) GetEabKey() string {
	if x != nil {
		return x.EabKey
	}
	return ""
}

func (x *ResolveAccountIdResponse) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

var File_eab_mgmt_proto protoreflect.FileDescriptor

var file_eab_mgmt_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x61, 0x62, 0x2d, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x65, 0x61, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x4e, 0x0a, 0x19,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x41, 0x42, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x61, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x61, 0x62, 0x4b, 0x65, 0x79, 0x22, 0x36, 0x0a, 0x1a,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x41, 0x42, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x22, 0x38, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x47,
	0x0a, 0x18, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x61,
	0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x61, 0x62,
	0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xd1, 0x01, 0x0a, 0x0a, 0x45, 0x41, 0x42, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45,
	0x41, 0x42, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x2e,
	0x65, 0x61, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x45, 0x41, 0x42, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x61, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x41, 0x42, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x10,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x23, 0x2e, 0x65, 0x61, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x61, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6d, 0x2d, 0x65, 0x64, 0x75,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eab_mgmt_proto_rawDescOnce sync.Once
	file_eab_mgmt_proto_rawDescData = file_eab_mgmt_proto_rawDesc
)

func file_eab_mgmt_proto_rawDescGZIP() []byte {
	file_eab_mgmt_proto_rawDescOnce.Do(func() {
		file_eab_mgmt_proto_rawDescData = protoimpl.X.CompressGZIP(file_eab_mgmt_proto_rawDescData)
	})
	return file_eab_mgmt_proto_rawDescData
}

var file_eab_mgmt_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eab_mgmt_proto_goTypes = []any{
	(*CheckEABPermissionRequest)(nil),  // 0: eabService.CheckEABPermissionRequest
	(*CheckEABPermissionResponse)(nil), // 1: eabService.CheckEABPermissionResponse
	(*ResolveAccountIdRequest)(nil),    // 2: eabService.ResolveAccountIdRequest
	(*ResolveAccountIdResponse)(nil),   // 3: eabService.ResolveAccountIdResponse
}
var file_eab_mgmt_proto_depIdxs = []int32{
	0, // 0: eabService.EABService.CheckEABPermissions:input_type -> eabService.CheckEABPermissionRequest
	2, // 1: eabService.EABService.ResolveAccountId:input_type -> eabService.ResolveAccountIdRequest
	1, // 2: eabService.EABService.CheckEABPermissions:output_type -> eabService.CheckEABPermissionResponse
	3, // 3: eabService.EABService.ResolveAccountId:output_type -> eabService.ResolveAccountIdResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eab_mgmt_proto_init() }
func file_eab_mgmt_proto_init() {
	if File_eab_mgmt_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eab_mgmt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eab_mgmt_proto_goTypes,
		DependencyIndexes: file_eab_mgmt_proto_depIdxs,
		MessageInfos:      file_eab_mgmt_proto_msgTypes,
	}.Build()
	File_eab_mgmt_proto = out.File
	file_eab_mgmt_proto_rawDesc = nil
	file_eab_mgmt_proto_goTypes = nil
	file_eab_mgmt_proto_depIdxs = nil
}
