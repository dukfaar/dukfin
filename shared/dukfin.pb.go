// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: dukfin.proto

package shared

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type TransactionType int32

const (
	TransactionType_TRANSFER_IN  TransactionType = 0
	TransactionType_TRANSFER_OUT TransactionType = 1
	TransactionType_DEPOSIT      TransactionType = 2
	TransactionType_WITHDRAW     TransactionType = 3
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "TRANSFER_IN",
		1: "TRANSFER_OUT",
		2: "DEPOSIT",
		3: "WITHDRAW",
	}
	TransactionType_value = map[string]int32{
		"TRANSFER_IN":  0,
		"TRANSFER_OUT": 1,
		"DEPOSIT":      2,
		"WITHDRAW":     3,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_dukfin_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_dukfin_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_dukfin_proto_rawDescGZIP(), []int{0}
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dukfin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dukfin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_dukfin_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateAccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateAccountReply) Reset() {
	*x = CreateAccountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dukfin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountReply) ProtoMessage() {}

func (x *CreateAccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_dukfin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountReply.ProtoReflect.Descriptor instead.
func (*CreateAccountReply) Descriptor() ([]byte, []int) {
	return file_dukfin_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateAccountReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAccountsRequest) Reset() {
	*x = GetAccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dukfin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountsRequest) ProtoMessage() {}

func (x *GetAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dukfin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountsRequest.ProtoReflect.Descriptor instead.
func (*GetAccountsRequest) Descriptor() ([]byte, []int) {
	return file_dukfin_proto_rawDescGZIP(), []int{2}
}

type GetAccountsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *GetAccountsReply) Reset() {
	*x = GetAccountsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dukfin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountsReply) ProtoMessage() {}

func (x *GetAccountsReply) ProtoReflect() protoreflect.Message {
	mi := &file_dukfin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountsReply.ProtoReflect.Descriptor instead.
func (*GetAccountsReply) Descriptor() ([]byte, []int) {
	return file_dukfin_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountsReply) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dukfin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_dukfin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_dukfin_proto_rawDescGZIP(), []int{4}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DepositRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value        string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	CurrencyCode string `protobuf:"bytes,2,opt,name=currencyCode,proto3" json:"currencyCode,omitempty"`
	AccountId    string `protobuf:"bytes,3,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *DepositRequest) Reset() {
	*x = DepositRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dukfin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRequest) ProtoMessage() {}

func (x *DepositRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dukfin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRequest.ProtoReflect.Descriptor instead.
func (*DepositRequest) Descriptor() ([]byte, []int) {
	return file_dukfin_proto_rawDescGZIP(), []int{5}
}

func (x *DepositRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *DepositRequest) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *DepositRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type WithdrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value        string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	CurrencyCode string `protobuf:"bytes,2,opt,name=currencyCode,proto3" json:"currencyCode,omitempty"`
	AccountId    string `protobuf:"bytes,3,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *WithdrawRequest) Reset() {
	*x = WithdrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dukfin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawRequest) ProtoMessage() {}

func (x *WithdrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dukfin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawRequest.ProtoReflect.Descriptor instead.
func (*WithdrawRequest) Descriptor() ([]byte, []int) {
	return file_dukfin_proto_rawDescGZIP(), []int{6}
}

func (x *WithdrawRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *WithdrawRequest) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *WithdrawRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type TransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value         string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	CurrencyCode  string `protobuf:"bytes,2,opt,name=currencyCode,proto3" json:"currencyCode,omitempty"`
	FromAccountId string `protobuf:"bytes,3,opt,name=fromAccountId,proto3" json:"fromAccountId,omitempty"`
	ToAccountId   string `protobuf:"bytes,4,opt,name=toAccountId,proto3" json:"toAccountId,omitempty"`
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dukfin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dukfin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_dukfin_proto_rawDescGZIP(), []int{7}
}

func (x *TransferRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TransferRequest) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *TransferRequest) GetFromAccountId() string {
	if x != nil {
		return x.FromAccountId
	}
	return ""
}

func (x *TransferRequest) GetToAccountId() string {
	if x != nil {
		return x.ToAccountId
	}
	return ""
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Time *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Type TransactionType        `protobuf:"varint,3,opt,name=type,proto3,enum=TransactionType" json:"type,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dukfin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_dukfin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_dukfin_proto_rawDescGZIP(), []int{8}
}

func (x *Transaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transaction) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Transaction) GetType() TransactionType {
	if x != nil {
		return x.Type
	}
	return TransactionType_TRANSFER_IN
}

type TransactionsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *TransactionsReply) Reset() {
	*x = TransactionsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dukfin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsReply) ProtoMessage() {}

func (x *TransactionsReply) ProtoReflect() protoreflect.Message {
	mi := &file_dukfin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsReply.ProtoReflect.Descriptor instead.
func (*TransactionsReply) Descriptor() ([]byte, []int) {
	return file_dukfin_proto_rawDescGZIP(), []int{9}
}

func (x *TransactionsReply) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

var File_dukfin_proto protoreflect.FileDescriptor

var file_dukfin_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x75, 0x6b, 0x66, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x24, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x2d, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x68, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x69,
	0x0a, 0x0f, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x0f, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x66, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x74, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x73, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x45, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x4f, 0x0a, 0x0f, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f,
	0x0a, 0x0b, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x10, 0x00, 0x12,
	0x10, 0x0a, 0x0c, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x4f, 0x55, 0x54, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x10, 0x03, 0x32, 0x9f, 0x02, 0x0a,
	0x06, 0x44, 0x75, 0x6b, 0x46, 0x69, 0x6e, 0x12, 0x32, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x08, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x08, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x10, 0x44,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x0f, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x13, 0x57, 0x69, 0x74, 0x68, 0x44, 0x72,
	0x61, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x2e,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x10, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x22,
	0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x6b,
	0x66, 0x61, 0x61, 0x72, 0x2f, 0x64, 0x75, 0x6b, 0x66, 0x69, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dukfin_proto_rawDescOnce sync.Once
	file_dukfin_proto_rawDescData = file_dukfin_proto_rawDesc
)

func file_dukfin_proto_rawDescGZIP() []byte {
	file_dukfin_proto_rawDescOnce.Do(func() {
		file_dukfin_proto_rawDescData = protoimpl.X.CompressGZIP(file_dukfin_proto_rawDescData)
	})
	return file_dukfin_proto_rawDescData
}

var file_dukfin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dukfin_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_dukfin_proto_goTypes = []interface{}{
	(TransactionType)(0),          // 0: TransactionType
	(*CreateAccountRequest)(nil),  // 1: CreateAccountRequest
	(*CreateAccountReply)(nil),    // 2: CreateAccountReply
	(*GetAccountsRequest)(nil),    // 3: GetAccountsRequest
	(*GetAccountsReply)(nil),      // 4: GetAccountsReply
	(*Account)(nil),               // 5: Account
	(*DepositRequest)(nil),        // 6: DepositRequest
	(*WithdrawRequest)(nil),       // 7: WithdrawRequest
	(*TransferRequest)(nil),       // 8: TransferRequest
	(*Transaction)(nil),           // 9: Transaction
	(*TransactionsReply)(nil),     // 10: TransactionsReply
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_dukfin_proto_depIdxs = []int32{
	5,  // 0: GetAccountsReply.accounts:type_name -> Account
	11, // 1: Transaction.time:type_name -> google.protobuf.Timestamp
	0,  // 2: Transaction.type:type_name -> TransactionType
	9,  // 3: TransactionsReply.transactions:type_name -> Transaction
	1,  // 4: DukFin.CreateAccount:input_type -> CreateAccountRequest
	3,  // 5: DukFin.GetAccounts:input_type -> GetAccountsRequest
	6,  // 6: DukFin.DepositToAccount:input_type -> DepositRequest
	7,  // 7: DukFin.WithDrawFromAccount:input_type -> WithdrawRequest
	8,  // 8: DukFin.TransferMoney:input_type -> TransferRequest
	5,  // 9: DukFin.CreateAccount:output_type -> Account
	5,  // 10: DukFin.GetAccounts:output_type -> Account
	10, // 11: DukFin.DepositToAccount:output_type -> TransactionsReply
	10, // 12: DukFin.WithDrawFromAccount:output_type -> TransactionsReply
	10, // 13: DukFin.TransferMoney:output_type -> TransactionsReply
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_dukfin_proto_init() }
func file_dukfin_proto_init() {
	if File_dukfin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dukfin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_dukfin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountReply); i {
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
		file_dukfin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountsRequest); i {
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
		file_dukfin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountsReply); i {
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
		file_dukfin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_dukfin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositRequest); i {
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
		file_dukfin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawRequest); i {
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
		file_dukfin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferRequest); i {
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
		file_dukfin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_dukfin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionsReply); i {
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
			RawDescriptor: file_dukfin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dukfin_proto_goTypes,
		DependencyIndexes: file_dukfin_proto_depIdxs,
		EnumInfos:         file_dukfin_proto_enumTypes,
		MessageInfos:      file_dukfin_proto_msgTypes,
	}.Build()
	File_dukfin_proto = out.File
	file_dukfin_proto_rawDesc = nil
	file_dukfin_proto_goTypes = nil
	file_dukfin_proto_depIdxs = nil
}