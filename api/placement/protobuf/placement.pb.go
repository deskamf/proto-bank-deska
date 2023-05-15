// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/placement.proto

package __

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

type SignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReffNumber        string               `protobuf:"bytes,1,opt,name=reff_number,json=reffNumber,proto3" json:"reff_number,omitempty"`
	DocFlow           string               `protobuf:"bytes,2,opt,name=doc_flow,json=docFlow,proto3" json:"doc_flow,omitempty"`
	SignPositions     []*SignPositions     `protobuf:"bytes,3,rep,name=sign_positions,json=signPositions,proto3" json:"sign_positions,omitempty"`
	EmeteraiPositions []*EmeteraiPositions `protobuf:"bytes,4,rep,name=emeterai_positions,json=emeteraiPositions,proto3" json:"emeterai_positions,omitempty"`
	SignType          *string              `protobuf:"bytes,5,opt,name=sign_type,json=signType,proto3,oneof" json:"sign_type,omitempty"`
}

func (x *SignRequest) Reset() {
	*x = SignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRequest) ProtoMessage() {}

func (x *SignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRequest.ProtoReflect.Descriptor instead.
func (*SignRequest) Descriptor() ([]byte, []int) {
	return file_proto_placement_proto_rawDescGZIP(), []int{0}
}

func (x *SignRequest) GetReffNumber() string {
	if x != nil {
		return x.ReffNumber
	}
	return ""
}

func (x *SignRequest) GetDocFlow() string {
	if x != nil {
		return x.DocFlow
	}
	return ""
}

func (x *SignRequest) GetSignPositions() []*SignPositions {
	if x != nil {
		return x.SignPositions
	}
	return nil
}

func (x *SignRequest) GetEmeteraiPositions() []*EmeteraiPositions {
	if x != nil {
		return x.EmeteraiPositions
	}
	return nil
}

func (x *SignRequest) GetSignType() string {
	if x != nil && x.SignType != nil {
		return *x.SignType
	}
	return ""
}

type SignPositions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nik         *string `protobuf:"bytes,1,opt,name=nik,proto3,oneof" json:"nik,omitempty"`
	Name        *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Dob         *string `protobuf:"bytes,3,opt,name=dob,proto3,oneof" json:"dob,omitempty"`
	PhoneNumber *string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3,oneof" json:"phone_number,omitempty"`
	Email       *string `protobuf:"bytes,5,opt,name=email,proto3,oneof" json:"email,omitempty"`
	PosX        float32 `protobuf:"fixed32,6,opt,name=pos_x,json=posX,proto3" json:"pos_x,omitempty"`
	PosY        float32 `protobuf:"fixed32,7,opt,name=pos_y,json=posY,proto3" json:"pos_y,omitempty"`
	SignPage    string  `protobuf:"bytes,8,opt,name=sign_page,json=signPage,proto3" json:"sign_page,omitempty"`
}

func (x *SignPositions) Reset() {
	*x = SignPositions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignPositions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignPositions) ProtoMessage() {}

func (x *SignPositions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignPositions.ProtoReflect.Descriptor instead.
func (*SignPositions) Descriptor() ([]byte, []int) {
	return file_proto_placement_proto_rawDescGZIP(), []int{1}
}

func (x *SignPositions) GetNik() string {
	if x != nil && x.Nik != nil {
		return *x.Nik
	}
	return ""
}

func (x *SignPositions) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *SignPositions) GetDob() string {
	if x != nil && x.Dob != nil {
		return *x.Dob
	}
	return ""
}

func (x *SignPositions) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *SignPositions) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *SignPositions) GetPosX() float32 {
	if x != nil {
		return x.PosX
	}
	return 0
}

func (x *SignPositions) GetPosY() float32 {
	if x != nil {
		return x.PosY
	}
	return 0
}

func (x *SignPositions) GetSignPage() string {
	if x != nil {
		return x.SignPage
	}
	return ""
}

type EmeteraiPositions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosX float32 `protobuf:"fixed32,1,opt,name=pos_x,json=posX,proto3" json:"pos_x,omitempty"`
	PosY float32 `protobuf:"fixed32,2,opt,name=pos_y,json=posY,proto3" json:"pos_y,omitempty"`
	Page string  `protobuf:"bytes,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *EmeteraiPositions) Reset() {
	*x = EmeteraiPositions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmeteraiPositions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmeteraiPositions) ProtoMessage() {}

func (x *EmeteraiPositions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmeteraiPositions.ProtoReflect.Descriptor instead.
func (*EmeteraiPositions) Descriptor() ([]byte, []int) {
	return file_proto_placement_proto_rawDescGZIP(), []int{2}
}

func (x *EmeteraiPositions) GetPosX() float32 {
	if x != nil {
		return x.PosX
	}
	return 0
}

func (x *EmeteraiPositions) GetPosY() float32 {
	if x != nil {
		return x.PosY
	}
	return 0
}

func (x *EmeteraiPositions) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

type SignResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReffNumber string `protobuf:"bytes,1,opt,name=reff_number,json=reffNumber,proto3" json:"reff_number,omitempty"`
	DocToken   string `protobuf:"bytes,2,opt,name=doc_token,json=docToken,proto3" json:"doc_token,omitempty"`
}

func (x *SignResponse) Reset() {
	*x = SignResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignResponse) ProtoMessage() {}

func (x *SignResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignResponse.ProtoReflect.Descriptor instead.
func (*SignResponse) Descriptor() ([]byte, []int) {
	return file_proto_placement_proto_rawDescGZIP(), []int{3}
}

func (x *SignResponse) GetReffNumber() string {
	if x != nil {
		return x.ReffNumber
	}
	return ""
}

func (x *SignResponse) GetDocToken() string {
	if x != nil {
		return x.DocToken
	}
	return ""
}

type Otp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReffNumber string `protobuf:"bytes,1,opt,name=reff_number,json=reffNumber,proto3" json:"reff_number,omitempty"`
	DocToken   string `protobuf:"bytes,2,opt,name=doc_token,json=docToken,proto3" json:"doc_token,omitempty"`
	OtpChannel int32  `protobuf:"varint,3,opt,name=otp_channel,json=otpChannel,proto3" json:"otp_channel,omitempty"`
}

func (x *Otp) Reset() {
	*x = Otp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Otp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Otp) ProtoMessage() {}

func (x *Otp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Otp.ProtoReflect.Descriptor instead.
func (*Otp) Descriptor() ([]byte, []int) {
	return file_proto_placement_proto_rawDescGZIP(), []int{4}
}

func (x *Otp) GetReffNumber() string {
	if x != nil {
		return x.ReffNumber
	}
	return ""
}

func (x *Otp) GetDocToken() string {
	if x != nil {
		return x.DocToken
	}
	return ""
}

func (x *Otp) GetOtpChannel() int32 {
	if x != nil {
		return x.OtpChannel
	}
	return 0
}

type OtpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId  string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	ReffNumber     string `protobuf:"bytes,2,opt,name=reff_number,json=reffNumber,proto3" json:"reff_number,omitempty"`
	RequestAttempt string `protobuf:"bytes,3,opt,name=request_attempt,json=requestAttempt,proto3" json:"request_attempt,omitempty"`
}

func (x *OtpResponse) Reset() {
	*x = OtpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtpResponse) ProtoMessage() {}

func (x *OtpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtpResponse.ProtoReflect.Descriptor instead.
func (*OtpResponse) Descriptor() ([]byte, []int) {
	return file_proto_placement_proto_rawDescGZIP(), []int{5}
}

func (x *OtpResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *OtpResponse) GetReffNumber() string {
	if x != nil {
		return x.ReffNumber
	}
	return ""
}

func (x *OtpResponse) GetRequestAttempt() string {
	if x != nil {
		return x.RequestAttempt
	}
	return ""
}

type OtpValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReffNumber    string `protobuf:"bytes,1,opt,name=reff_number,json=reffNumber,proto3" json:"reff_number,omitempty"`
	OtpCode       string `protobuf:"bytes,2,opt,name=otp_code,json=otpCode,proto3" json:"otp_code,omitempty"`
	TransactionId string `protobuf:"bytes,3,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *OtpValidateRequest) Reset() {
	*x = OtpValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtpValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtpValidateRequest) ProtoMessage() {}

func (x *OtpValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtpValidateRequest.ProtoReflect.Descriptor instead.
func (*OtpValidateRequest) Descriptor() ([]byte, []int) {
	return file_proto_placement_proto_rawDescGZIP(), []int{6}
}

func (x *OtpValidateRequest) GetReffNumber() string {
	if x != nil {
		return x.ReffNumber
	}
	return ""
}

func (x *OtpValidateRequest) GetOtpCode() string {
	if x != nil {
		return x.OtpCode
	}
	return ""
}

func (x *OtpValidateRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type OtpValidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId  string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	ReffNumber     string `protobuf:"bytes,2,opt,name=reff_number,json=reffNumber,proto3" json:"reff_number,omitempty"`
	RequestAttempt string `protobuf:"bytes,3,opt,name=request_attempt,json=requestAttempt,proto3" json:"request_attempt,omitempty"`
}

func (x *OtpValidateResponse) Reset() {
	*x = OtpValidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtpValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtpValidateResponse) ProtoMessage() {}

func (x *OtpValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtpValidateResponse.ProtoReflect.Descriptor instead.
func (*OtpValidateResponse) Descriptor() ([]byte, []int) {
	return file_proto_placement_proto_rawDescGZIP(), []int{7}
}

func (x *OtpValidateResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *OtpValidateResponse) GetReffNumber() string {
	if x != nil {
		return x.ReffNumber
	}
	return ""
}

func (x *OtpValidateResponse) GetRequestAttempt() string {
	if x != nil {
		return x.RequestAttempt
	}
	return ""
}

type OtpResendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *OtpResendRequest) Reset() {
	*x = OtpResendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtpResendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtpResendRequest) ProtoMessage() {}

func (x *OtpResendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtpResendRequest.ProtoReflect.Descriptor instead.
func (*OtpResendRequest) Descriptor() ([]byte, []int) {
	return file_proto_placement_proto_rawDescGZIP(), []int{8}
}

func (x *OtpResendRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type OtpResendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId  string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	ReffNumber     string `protobuf:"bytes,2,opt,name=reff_number,json=reffNumber,proto3" json:"reff_number,omitempty"`
	RequestAttempt string `protobuf:"bytes,3,opt,name=request_attempt,json=requestAttempt,proto3" json:"request_attempt,omitempty"`
}

func (x *OtpResendResponse) Reset() {
	*x = OtpResendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtpResendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtpResendResponse) ProtoMessage() {}

func (x *OtpResendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtpResendResponse.ProtoReflect.Descriptor instead.
func (*OtpResendResponse) Descriptor() ([]byte, []int) {
	return file_proto_placement_proto_rawDescGZIP(), []int{9}
}

func (x *OtpResendResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *OtpResendResponse) GetReffNumber() string {
	if x != nil {
		return x.ReffNumber
	}
	return ""
}

func (x *OtpResendResponse) GetRequestAttempt() string {
	if x != nil {
		return x.RequestAttempt
	}
	return ""
}

var File_proto_placement_proto protoreflect.FileDescriptor

var file_proto_placement_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x66, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x66, 0x66, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x5f,
	0x66, 0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x63, 0x46,
	0x6c, 0x6f, 0x77, 0x12, 0x35, 0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x73, 0x69, 0x67,
	0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x41, 0x0a, 0x12, 0x65, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x61, 0x69, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x45, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x61,
	0x69, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x11, 0x65, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x61, 0x69, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x94, 0x02,
	0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x15, 0x0a, 0x03, 0x6e, 0x69, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03,
	0x6e, 0x69, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x15, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x03,
	0x64, 0x6f, 0x62, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x6f, 0x73,
	0x5f, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x58, 0x12, 0x13,
	0x0a, 0x05, 0x70, 0x6f, 0x73, 0x5f, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x70,
	0x6f, 0x73, 0x59, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x50, 0x61, 0x67, 0x65,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6e, 0x69, 0x6b, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x64, 0x6f, 0x62, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x51, 0x0a, 0x11, 0x45, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x61, 0x69,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x6f, 0x73,
	0x5f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x58, 0x12, 0x13,
	0x0a, 0x05, 0x70, 0x6f, 0x73, 0x5f, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x70,
	0x6f, 0x73, 0x59, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x4c, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x66, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x66, 0x66, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x63,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x64, 0x0a, 0x03, 0x4f, 0x74, 0x70, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x66, 0x66, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x66, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x6f, 0x63, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x6f, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x74,
	0x70, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x6f, 0x74, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x7e, 0x0a, 0x0b, 0x4f,
	0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x66, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x66, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x74,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x22, 0x77, 0x0a, 0x12, 0x4f,
	0x74, 0x70, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x66, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x66, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x13, 0x4f, 0x74, 0x70, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x66, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x66, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x22, 0x39, 0x0a,
	0x10, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x11, 0x4f, 0x74, 0x70,
	0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x66, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x66,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x32,
	0xc0, 0x01, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a,
	0x04, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x0c, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x0a, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x04, 0x2e, 0x4f, 0x74, 0x70, 0x1a, 0x0c, 0x2e, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x4f, 0x74, 0x70, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x4f, 0x74, 0x70, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x4f, 0x74, 0x70, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x09, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x11, 0x2e, 0x4f, 0x74,
	0x70, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_placement_proto_rawDescOnce sync.Once
	file_proto_placement_proto_rawDescData = file_proto_placement_proto_rawDesc
)

func file_proto_placement_proto_rawDescGZIP() []byte {
	file_proto_placement_proto_rawDescOnce.Do(func() {
		file_proto_placement_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_placement_proto_rawDescData)
	})
	return file_proto_placement_proto_rawDescData
}

var file_proto_placement_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_placement_proto_goTypes = []interface{}{
	(*SignRequest)(nil),         // 0: SignRequest
	(*SignPositions)(nil),       // 1: SignPositions
	(*EmeteraiPositions)(nil),   // 2: EmeteraiPositions
	(*SignResponse)(nil),        // 3: SignResponse
	(*Otp)(nil),                 // 4: Otp
	(*OtpResponse)(nil),         // 5: OtpResponse
	(*OtpValidateRequest)(nil),  // 6: OtpValidateRequest
	(*OtpValidateResponse)(nil), // 7: OtpValidateResponse
	(*OtpResendRequest)(nil),    // 8: OtpResendRequest
	(*OtpResendResponse)(nil),   // 9: OtpResendResponse
}
var file_proto_placement_proto_depIdxs = []int32{
	1, // 0: SignRequest.sign_positions:type_name -> SignPositions
	2, // 1: SignRequest.emeterai_positions:type_name -> EmeteraiPositions
	0, // 2: Placement.Sign:input_type -> SignRequest
	4, // 3: Placement.OtpRequest:input_type -> Otp
	6, // 4: Placement.OtpValidate:input_type -> OtpValidateRequest
	8, // 5: Placement.OtpResend:input_type -> OtpResendRequest
	3, // 6: Placement.Sign:output_type -> SignResponse
	5, // 7: Placement.OtpRequest:output_type -> OtpResponse
	7, // 8: Placement.OtpValidate:output_type -> OtpValidateResponse
	9, // 9: Placement.OtpResend:output_type -> OtpResendResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_placement_proto_init() }
func file_proto_placement_proto_init() {
	if File_proto_placement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_placement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRequest); i {
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
		file_proto_placement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignPositions); i {
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
		file_proto_placement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmeteraiPositions); i {
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
		file_proto_placement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignResponse); i {
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
		file_proto_placement_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Otp); i {
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
		file_proto_placement_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtpResponse); i {
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
		file_proto_placement_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtpValidateRequest); i {
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
		file_proto_placement_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtpValidateResponse); i {
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
		file_proto_placement_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtpResendRequest); i {
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
		file_proto_placement_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtpResendResponse); i {
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
	file_proto_placement_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_placement_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_placement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_placement_proto_goTypes,
		DependencyIndexes: file_proto_placement_proto_depIdxs,
		MessageInfos:      file_proto_placement_proto_msgTypes,
	}.Build()
	File_proto_placement_proto = out.File
	file_proto_placement_proto_rawDesc = nil
	file_proto_placement_proto_goTypes = nil
	file_proto_placement_proto_depIdxs = nil
}
