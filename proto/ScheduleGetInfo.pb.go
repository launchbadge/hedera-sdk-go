// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: proto/ScheduleGetInfo.proto

package proto

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

//
//Gets information about Scheduled Transaction instance
//
//Answered with status INVALID_SCHEDULE_ID if the Scheduled Transaction is deleted, expires or is executed.
type ScheduleGetInfoQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header     *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`         // standard info sent from client to node including the signed payment, and what kind of response is requested (cost, state proof, both, or neither).
	ScheduleID *ScheduleID  `protobuf:"bytes,2,opt,name=scheduleID,proto3" json:"scheduleID,omitempty"` // The ID of the Scheduled Entity
}

func (x *ScheduleGetInfoQuery) Reset() {
	*x = ScheduleGetInfoQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ScheduleGetInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleGetInfoQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleGetInfoQuery) ProtoMessage() {}

func (x *ScheduleGetInfoQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ScheduleGetInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleGetInfoQuery.ProtoReflect.Descriptor instead.
func (*ScheduleGetInfoQuery) Descriptor() ([]byte, []int) {
	return file_proto_ScheduleGetInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleGetInfoQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ScheduleGetInfoQuery) GetScheduleID() *ScheduleID {
	if x != nil {
		return x.ScheduleID
	}
	return nil
}

//
//Metadata about a Scheduled Transaction instance
type ScheduleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleID       *ScheduleID `protobuf:"bytes,1,opt,name=scheduleID,proto3" json:"scheduleID,omitempty"`             // The ID of the Scheduled Entity
	CreatorAccountID *AccountID  `protobuf:"bytes,2,opt,name=creatorAccountID,proto3" json:"creatorAccountID,omitempty"` // The Account ID which created the Scheduled TX
	PayerAccountID   *AccountID  `protobuf:"bytes,3,opt,name=payerAccountID,proto3" json:"payerAccountID,omitempty"`     // The account which is going to pay for the execution of the Scheduled TX
	TransactionBody  []byte      `protobuf:"bytes,4,opt,name=transactionBody,proto3" json:"transactionBody,omitempty"`   // The transaction serialized into bytes that must be signed
	Signatories      *KeyList    `protobuf:"bytes,5,opt,name=signatories,proto3" json:"signatories,omitempty"`           // The signatories that have provided signatures so far for the Scheduled TX
	AdminKey         *Key        `protobuf:"bytes,6,opt,name=adminKey,proto3" json:"adminKey,omitempty"`                 // The Key which is able to delete the Scheduled Transaction if set
	Memo             string      `protobuf:"bytes,7,opt,name=memo,proto3" json:"memo,omitempty"`                         // Publicly visible information about the Scheduled entity, up to 100 bytes. No guarantee of uniqueness.
	ExpirationTime   *Timestamp  `protobuf:"bytes,8,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`     // The epoch second at which the schedule will expire
}

func (x *ScheduleInfo) Reset() {
	*x = ScheduleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ScheduleGetInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleInfo) ProtoMessage() {}

func (x *ScheduleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ScheduleGetInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleInfo.ProtoReflect.Descriptor instead.
func (*ScheduleInfo) Descriptor() ([]byte, []int) {
	return file_proto_ScheduleGetInfo_proto_rawDescGZIP(), []int{1}
}

func (x *ScheduleInfo) GetScheduleID() *ScheduleID {
	if x != nil {
		return x.ScheduleID
	}
	return nil
}

func (x *ScheduleInfo) GetCreatorAccountID() *AccountID {
	if x != nil {
		return x.CreatorAccountID
	}
	return nil
}

func (x *ScheduleInfo) GetPayerAccountID() *AccountID {
	if x != nil {
		return x.PayerAccountID
	}
	return nil
}

func (x *ScheduleInfo) GetTransactionBody() []byte {
	if x != nil {
		return x.TransactionBody
	}
	return nil
}

func (x *ScheduleInfo) GetSignatories() *KeyList {
	if x != nil {
		return x.Signatories
	}
	return nil
}

func (x *ScheduleInfo) GetAdminKey() *Key {
	if x != nil {
		return x.AdminKey
	}
	return nil
}

func (x *ScheduleInfo) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *ScheduleInfo) GetExpirationTime() *Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

//
//Response that gets returned, when the client sends the node a ScheduleGetInfoQuery
type ScheduleGetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`             // Standard response from node to client, including the requested fields: cost, or state proof, or both, or neither
	ScheduleInfo *ScheduleInfo   `protobuf:"bytes,2,opt,name=scheduleInfo,proto3" json:"scheduleInfo,omitempty"` // The information requested about this schedule instance
}

func (x *ScheduleGetInfoResponse) Reset() {
	*x = ScheduleGetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ScheduleGetInfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleGetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleGetInfoResponse) ProtoMessage() {}

func (x *ScheduleGetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ScheduleGetInfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleGetInfoResponse.ProtoReflect.Descriptor instead.
func (*ScheduleGetInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_ScheduleGetInfo_proto_rawDescGZIP(), []int{2}
}

func (x *ScheduleGetInfoResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ScheduleGetInfoResponse) GetScheduleInfo() *ScheduleInfo {
	if x != nil {
		return x.ScheduleInfo
	}
	return nil
}

var File_proto_ScheduleGetInfo_proto protoreflect.FileDescriptor

var file_proto_ScheduleGetInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x14, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x52,
	0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x22, 0x8b, 0x03, 0x0a, 0x0c,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x0a,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x49, 0x44, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x12,
	0x3c, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x10, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x38, 0x0a,
	0x0e, 0x70, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x30, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b,
	0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65,
	0x79, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12,
	0x38, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x17, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x50, 0x0a,
	0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a,
	0x61, 0x76, 0x61, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ScheduleGetInfo_proto_rawDescOnce sync.Once
	file_proto_ScheduleGetInfo_proto_rawDescData = file_proto_ScheduleGetInfo_proto_rawDesc
)

func file_proto_ScheduleGetInfo_proto_rawDescGZIP() []byte {
	file_proto_ScheduleGetInfo_proto_rawDescOnce.Do(func() {
		file_proto_ScheduleGetInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ScheduleGetInfo_proto_rawDescData)
	})
	return file_proto_ScheduleGetInfo_proto_rawDescData
}

var file_proto_ScheduleGetInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_ScheduleGetInfo_proto_goTypes = []interface{}{
	(*ScheduleGetInfoQuery)(nil),    // 0: proto.ScheduleGetInfoQuery
	(*ScheduleInfo)(nil),            // 1: proto.ScheduleInfo
	(*ScheduleGetInfoResponse)(nil), // 2: proto.ScheduleGetInfoResponse
	(*QueryHeader)(nil),             // 3: proto.QueryHeader
	(*ScheduleID)(nil),              // 4: proto.ScheduleID
	(*AccountID)(nil),               // 5: proto.AccountID
	(*KeyList)(nil),                 // 6: proto.KeyList
	(*Key)(nil),                     // 7: proto.Key
	(*Timestamp)(nil),               // 8: proto.Timestamp
	(*ResponseHeader)(nil),          // 9: proto.ResponseHeader
}
var file_proto_ScheduleGetInfo_proto_depIdxs = []int32{
	3,  // 0: proto.ScheduleGetInfoQuery.header:type_name -> proto.QueryHeader
	4,  // 1: proto.ScheduleGetInfoQuery.scheduleID:type_name -> proto.ScheduleID
	4,  // 2: proto.ScheduleInfo.scheduleID:type_name -> proto.ScheduleID
	5,  // 3: proto.ScheduleInfo.creatorAccountID:type_name -> proto.AccountID
	5,  // 4: proto.ScheduleInfo.payerAccountID:type_name -> proto.AccountID
	6,  // 5: proto.ScheduleInfo.signatories:type_name -> proto.KeyList
	7,  // 6: proto.ScheduleInfo.adminKey:type_name -> proto.Key
	8,  // 7: proto.ScheduleInfo.expirationTime:type_name -> proto.Timestamp
	9,  // 8: proto.ScheduleGetInfoResponse.header:type_name -> proto.ResponseHeader
	1,  // 9: proto.ScheduleGetInfoResponse.scheduleInfo:type_name -> proto.ScheduleInfo
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_proto_ScheduleGetInfo_proto_init() }
func file_proto_ScheduleGetInfo_proto_init() {
	if File_proto_ScheduleGetInfo_proto != nil {
		return
	}
	file_proto_BasicTypes_proto_init()
	file_proto_Timestamp_proto_init()
	file_proto_QueryHeader_proto_init()
	file_proto_ResponseHeader_proto_init()
	file_proto_ScheduleCreate_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_ScheduleGetInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleGetInfoQuery); i {
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
		file_proto_ScheduleGetInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleInfo); i {
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
		file_proto_ScheduleGetInfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleGetInfoResponse); i {
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
			RawDescriptor: file_proto_ScheduleGetInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ScheduleGetInfo_proto_goTypes,
		DependencyIndexes: file_proto_ScheduleGetInfo_proto_depIdxs,
		MessageInfos:      file_proto_ScheduleGetInfo_proto_msgTypes,
	}.Build()
	File_proto_ScheduleGetInfo_proto = out.File
	file_proto_ScheduleGetInfo_proto_rawDesc = nil
	file_proto_ScheduleGetInfo_proto_goTypes = nil
	file_proto_ScheduleGetInfo_proto_depIdxs = nil
}
