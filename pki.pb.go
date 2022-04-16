// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pki.proto

package portal_apis

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RevokeSmimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Identifier:
	//	*RevokeSmimeRequest_OrderNumber
	//	*RevokeSmimeRequest_Serial
	//	*RevokeSmimeRequest_Email
	Identifier isRevokeSmimeRequest_Identifier `protobuf_oneof:"identifier"`
}

func (x *RevokeSmimeRequest) Reset() {
	*x = RevokeSmimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pki_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeSmimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeSmimeRequest) ProtoMessage() {}

func (x *RevokeSmimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pki_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeSmimeRequest.ProtoReflect.Descriptor instead.
func (*RevokeSmimeRequest) Descriptor() ([]byte, []int) {
	return file_pki_proto_rawDescGZIP(), []int{0}
}

func (m *RevokeSmimeRequest) GetIdentifier() isRevokeSmimeRequest_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (x *RevokeSmimeRequest) GetOrderNumber() int32 {
	if x, ok := x.GetIdentifier().(*RevokeSmimeRequest_OrderNumber); ok {
		return x.OrderNumber
	}
	return 0
}

func (x *RevokeSmimeRequest) GetSerial() string {
	if x, ok := x.GetIdentifier().(*RevokeSmimeRequest_Serial); ok {
		return x.Serial
	}
	return ""
}

func (x *RevokeSmimeRequest) GetEmail() string {
	if x, ok := x.GetIdentifier().(*RevokeSmimeRequest_Email); ok {
		return x.Email
	}
	return ""
}

type isRevokeSmimeRequest_Identifier interface {
	isRevokeSmimeRequest_Identifier()
}

type RevokeSmimeRequest_OrderNumber struct {
	OrderNumber int32 `protobuf:"varint,1,opt,name=orderNumber,proto3,oneof"`
}

type RevokeSmimeRequest_Serial struct {
	Serial string `protobuf:"bytes,2,opt,name=serial,proto3,oneof"`
}

type RevokeSmimeRequest_Email struct {
	Email string `protobuf:"bytes,3,opt,name=email,proto3,oneof"`
}

func (*RevokeSmimeRequest_OrderNumber) isRevokeSmimeRequest_Identifier() {}

func (*RevokeSmimeRequest_Serial) isRevokeSmimeRequest_Identifier() {}

func (*RevokeSmimeRequest_Email) isRevokeSmimeRequest_Identifier() {}

type ListSmimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ListSmimeRequest) Reset() {
	*x = ListSmimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pki_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSmimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSmimeRequest) ProtoMessage() {}

func (x *ListSmimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pki_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSmimeRequest.ProtoReflect.Descriptor instead.
func (*ListSmimeRequest) Descriptor() ([]byte, []int) {
	return file_pki_proto_rawDescGZIP(), []int{1}
}

func (x *ListSmimeRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type IssueSmimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Csr        string `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	Email      string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FirstName  string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName   string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	MiddleName string `protobuf:"bytes,5,opt,name=middleName,proto3" json:"middleName,omitempty"`
}

func (x *IssueSmimeRequest) Reset() {
	*x = IssueSmimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pki_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueSmimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueSmimeRequest) ProtoMessage() {}

func (x *IssueSmimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pki_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueSmimeRequest.ProtoReflect.Descriptor instead.
func (*IssueSmimeRequest) Descriptor() ([]byte, []int) {
	return file_pki_proto_rawDescGZIP(), []int{2}
}

func (x *IssueSmimeRequest) GetCsr() string {
	if x != nil {
		return x.Csr
	}
	return ""
}

func (x *IssueSmimeRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *IssueSmimeRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *IssueSmimeRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *IssueSmimeRequest) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

type IssueSmimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificate string `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
}

func (x *IssueSmimeResponse) Reset() {
	*x = IssueSmimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pki_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueSmimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueSmimeResponse) ProtoMessage() {}

func (x *IssueSmimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pki_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueSmimeResponse.ProtoReflect.Descriptor instead.
func (*IssueSmimeResponse) Descriptor() ([]byte, []int) {
	return file_pki_proto_rawDescGZIP(), []int{3}
}

func (x *IssueSmimeResponse) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

type ListSmimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status  string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Serial  string                 `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`
	Expires *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expires,proto3" json:"expires,omitempty"`
}

func (x *ListSmimeResponse) Reset() {
	*x = ListSmimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pki_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSmimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSmimeResponse) ProtoMessage() {}

func (x *ListSmimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pki_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSmimeResponse.ProtoReflect.Descriptor instead.
func (*ListSmimeResponse) Descriptor() ([]byte, []int) {
	return file_pki_proto_rawDescGZIP(), []int{4}
}

func (x *ListSmimeResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListSmimeResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ListSmimeResponse) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *ListSmimeResponse) GetExpires() *timestamppb.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

type ListSslRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domains []string `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty"`
	Status  string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ListSslRequest) Reset() {
	*x = ListSslRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pki_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSslRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSslRequest) ProtoMessage() {}

func (x *ListSslRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pki_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSslRequest.ProtoReflect.Descriptor instead.
func (*ListSslRequest) Descriptor() ([]byte, []int) {
	return file_pki_proto_rawDescGZIP(), []int{5}
}

func (x *ListSslRequest) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *ListSslRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CertificateDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CommonName              string                 `protobuf:"bytes,2,opt,name=common_name,json=commonName,proto3" json:"common_name,omitempty"`
	Status                  string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Serial                  string                 `protobuf:"bytes,4,opt,name=serial,proto3" json:"serial,omitempty"`
	SubjectAlternativeNames []string               `protobuf:"bytes,5,rep,name=subject_alternative_names,json=subjectAlternativeNames,proto3" json:"subject_alternative_names,omitempty"`
	Expires                 *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=expires,proto3" json:"expires,omitempty"`
	NotBefore               *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=notBefore,proto3" json:"notBefore,omitempty"`
}

func (x *CertificateDetails) Reset() {
	*x = CertificateDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pki_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateDetails) ProtoMessage() {}

func (x *CertificateDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pki_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateDetails.ProtoReflect.Descriptor instead.
func (*CertificateDetails) Descriptor() ([]byte, []int) {
	return file_pki_proto_rawDescGZIP(), []int{6}
}

func (x *CertificateDetails) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CertificateDetails) GetCommonName() string {
	if x != nil {
		return x.CommonName
	}
	return ""
}

func (x *CertificateDetails) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CertificateDetails) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *CertificateDetails) GetSubjectAlternativeNames() []string {
	if x != nil {
		return x.SubjectAlternativeNames
	}
	return nil
}

func (x *CertificateDetails) GetExpires() *timestamppb.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

func (x *CertificateDetails) GetNotBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.NotBefore
	}
	return nil
}

type ListSslResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CertificateDetails `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListSslResponse) Reset() {
	*x = ListSslResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pki_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSslResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSslResponse) ProtoMessage() {}

func (x *ListSslResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pki_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSslResponse.ProtoReflect.Descriptor instead.
func (*ListSslResponse) Descriptor() ([]byte, []int) {
	return file_pki_proto_rawDescGZIP(), []int{7}
}

func (x *ListSslResponse) GetItems() []*CertificateDetails {
	if x != nil {
		return x.Items
	}
	return nil
}

type IssueSslRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Csr                     string   `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	SubjectAlternativeNames []string `protobuf:"bytes,2,rep,name=subject_alternative_names,json=subjectAlternativeNames,proto3" json:"subject_alternative_names,omitempty"`
}

func (x *IssueSslRequest) Reset() {
	*x = IssueSslRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pki_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueSslRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueSslRequest) ProtoMessage() {}

func (x *IssueSslRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pki_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueSslRequest.ProtoReflect.Descriptor instead.
func (*IssueSslRequest) Descriptor() ([]byte, []int) {
	return file_pki_proto_rawDescGZIP(), []int{8}
}

func (x *IssueSslRequest) GetCsr() string {
	if x != nil {
		return x.Csr
	}
	return ""
}

func (x *IssueSslRequest) GetSubjectAlternativeNames() []string {
	if x != nil {
		return x.SubjectAlternativeNames
	}
	return nil
}

type IssueSslResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificate string `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
}

func (x *IssueSslResponse) Reset() {
	*x = IssueSslResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pki_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueSslResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueSslResponse) ProtoMessage() {}

func (x *IssueSslResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pki_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueSslResponse.ProtoReflect.Descriptor instead.
func (*IssueSslResponse) Descriptor() ([]byte, []int) {
	return file_pki_proto_rawDescGZIP(), []int{9}
}

func (x *IssueSslResponse) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

type RevokeSslRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Identifier:
	//	*RevokeSslRequest_Id
	//	*RevokeSslRequest_Serial
	//	*RevokeSslRequest_CommonName
	Identifier isRevokeSslRequest_Identifier `protobuf_oneof:"identifier"`
}

func (x *RevokeSslRequest) Reset() {
	*x = RevokeSslRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pki_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeSslRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeSslRequest) ProtoMessage() {}

func (x *RevokeSslRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pki_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeSslRequest.ProtoReflect.Descriptor instead.
func (*RevokeSslRequest) Descriptor() ([]byte, []int) {
	return file_pki_proto_rawDescGZIP(), []int{10}
}

func (m *RevokeSslRequest) GetIdentifier() isRevokeSslRequest_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (x *RevokeSslRequest) GetId() int32 {
	if x, ok := x.GetIdentifier().(*RevokeSslRequest_Id); ok {
		return x.Id
	}
	return 0
}

func (x *RevokeSslRequest) GetSerial() string {
	if x, ok := x.GetIdentifier().(*RevokeSslRequest_Serial); ok {
		return x.Serial
	}
	return ""
}

func (x *RevokeSslRequest) GetCommonName() string {
	if x, ok := x.GetIdentifier().(*RevokeSslRequest_CommonName); ok {
		return x.CommonName
	}
	return ""
}

type isRevokeSslRequest_Identifier interface {
	isRevokeSslRequest_Identifier()
}

type RevokeSslRequest_Id struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3,oneof"`
}

type RevokeSslRequest_Serial struct {
	Serial string `protobuf:"bytes,2,opt,name=serial,proto3,oneof"`
}

type RevokeSslRequest_CommonName struct {
	CommonName string `protobuf:"bytes,3,opt,name=common_name,json=commonName,proto3,oneof"`
}

func (*RevokeSslRequest_Id) isRevokeSslRequest_Identifier() {}

func (*RevokeSslRequest_Serial) isRevokeSslRequest_Identifier() {}

func (*RevokeSslRequest_CommonName) isRevokeSslRequest_Identifier() {}

var File_pki_proto protoreflect.FileDescriptor

var file_pki_proto_rawDesc = []byte{
	0x0a, 0x09, 0x70, 0x6b, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x6b, 0x69,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x12, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x53,
	0x6d, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x42, 0x0c, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22,
	0x28, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6d, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x95, 0x01, 0x0a, 0x11, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x53, 0x6d, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x73, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x73,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x36, 0x0a, 0x12, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x6d, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x6d, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12,
	0x34, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x73, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa1, 0x02, 0x0a, 0x12, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x12, 0x3a, 0x0a, 0x19, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6c, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x17, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6c, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x34, 0x0a,
	0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x22, 0x47, 0x0a,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x73, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x70, 0x6b, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x5f, 0x0a, 0x0f, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53,
	0x73, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x73, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x73, 0x72, 0x12, 0x3a, 0x0a, 0x19, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x17,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x10, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x53, 0x73, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0x6f, 0x0a,
	0x10, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x53, 0x73, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x21, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x42, 0x0c, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x32, 0xf3,
	0x01, 0x0a, 0x0a, 0x53, 0x53, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x1a, 0x2e, 0x70, 0x6b, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x73, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x6b, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x73, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1b,
	0x2e, 0x70, 0x6b, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x53, 0x73, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6b,
	0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x73,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x11, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x2e, 0x70, 0x6b, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x53, 0x73, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x32, 0xff, 0x01, 0x0a, 0x0c, 0x53, 0x6d, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x6b, 0x69, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6d, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x6b, 0x69, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6d, 0x69, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x6b, 0x69,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x6d, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x6b, 0x69, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x6d, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x11, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x2e, 0x70, 0x6b, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x53, 0x6d, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6d, 0x2d, 0x65, 0x64, 0x75, 0x2f, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pki_proto_rawDescOnce sync.Once
	file_pki_proto_rawDescData = file_pki_proto_rawDesc
)

func file_pki_proto_rawDescGZIP() []byte {
	file_pki_proto_rawDescOnce.Do(func() {
		file_pki_proto_rawDescData = protoimpl.X.CompressGZIP(file_pki_proto_rawDescData)
	})
	return file_pki_proto_rawDescData
}

var file_pki_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pki_proto_goTypes = []interface{}{
	(*RevokeSmimeRequest)(nil),    // 0: pkiService.RevokeSmimeRequest
	(*ListSmimeRequest)(nil),      // 1: pkiService.ListSmimeRequest
	(*IssueSmimeRequest)(nil),     // 2: pkiService.IssueSmimeRequest
	(*IssueSmimeResponse)(nil),    // 3: pkiService.IssueSmimeResponse
	(*ListSmimeResponse)(nil),     // 4: pkiService.ListSmimeResponse
	(*ListSslRequest)(nil),        // 5: pkiService.ListSslRequest
	(*CertificateDetails)(nil),    // 6: pkiService.CertificateDetails
	(*ListSslResponse)(nil),       // 7: pkiService.ListSslResponse
	(*IssueSslRequest)(nil),       // 8: pkiService.IssueSslRequest
	(*IssueSslResponse)(nil),      // 9: pkiService.IssueSslResponse
	(*RevokeSslRequest)(nil),      // 10: pkiService.RevokeSslRequest
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_pki_proto_depIdxs = []int32{
	11, // 0: pkiService.ListSmimeResponse.expires:type_name -> google.protobuf.Timestamp
	11, // 1: pkiService.CertificateDetails.expires:type_name -> google.protobuf.Timestamp
	11, // 2: pkiService.CertificateDetails.notBefore:type_name -> google.protobuf.Timestamp
	6,  // 3: pkiService.ListSslResponse.items:type_name -> pkiService.CertificateDetails
	5,  // 4: pkiService.SSLService.ListCertificates:input_type -> pkiService.ListSslRequest
	8,  // 5: pkiService.SSLService.IssueCertificate:input_type -> pkiService.IssueSslRequest
	10, // 6: pkiService.SSLService.RevokeCertificate:input_type -> pkiService.RevokeSslRequest
	1,  // 7: pkiService.SmimeService.ListCertificates:input_type -> pkiService.ListSmimeRequest
	2,  // 8: pkiService.SmimeService.IssueCertificate:input_type -> pkiService.IssueSmimeRequest
	0,  // 9: pkiService.SmimeService.RevokeCertificate:input_type -> pkiService.RevokeSmimeRequest
	7,  // 10: pkiService.SSLService.ListCertificates:output_type -> pkiService.ListSslResponse
	9,  // 11: pkiService.SSLService.IssueCertificate:output_type -> pkiService.IssueSslResponse
	12, // 12: pkiService.SSLService.RevokeCertificate:output_type -> google.protobuf.Empty
	4,  // 13: pkiService.SmimeService.ListCertificates:output_type -> pkiService.ListSmimeResponse
	3,  // 14: pkiService.SmimeService.IssueCertificate:output_type -> pkiService.IssueSmimeResponse
	12, // 15: pkiService.SmimeService.RevokeCertificate:output_type -> google.protobuf.Empty
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pki_proto_init() }
func file_pki_proto_init() {
	if File_pki_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pki_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeSmimeRequest); i {
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
		file_pki_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSmimeRequest); i {
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
		file_pki_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueSmimeRequest); i {
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
		file_pki_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueSmimeResponse); i {
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
		file_pki_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSmimeResponse); i {
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
		file_pki_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSslRequest); i {
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
		file_pki_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateDetails); i {
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
		file_pki_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSslResponse); i {
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
		file_pki_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueSslRequest); i {
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
		file_pki_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueSslResponse); i {
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
		file_pki_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeSslRequest); i {
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
	file_pki_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RevokeSmimeRequest_OrderNumber)(nil),
		(*RevokeSmimeRequest_Serial)(nil),
		(*RevokeSmimeRequest_Email)(nil),
	}
	file_pki_proto_msgTypes[10].OneofWrappers = []interface{}{
		(*RevokeSslRequest_Id)(nil),
		(*RevokeSslRequest_Serial)(nil),
		(*RevokeSslRequest_CommonName)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pki_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pki_proto_goTypes,
		DependencyIndexes: file_pki_proto_depIdxs,
		MessageInfos:      file_pki_proto_msgTypes,
	}.Build()
	File_pki_proto = out.File
	file_pki_proto_rawDesc = nil
	file_pki_proto_goTypes = nil
	file_pki_proto_depIdxs = nil
}
