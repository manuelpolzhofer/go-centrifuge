// Code generated by protoc-gen-go. DO NOT EDIT.
// source: funding/funding.proto

package fundingpb

import (
	context "context"
	fmt "fmt"
	math "math"

	document "github.com/centrifuge/go-centrifuge/protobufs/gen/go/document"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type FundingCreatePayload struct {
	Identifier           string                `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	ReadAccess           *document.ReadAccess  `protobuf:"bytes,2,opt,name=read_access,json=readAccess,proto3" json:"read_access,omitempty"`
	WriteAccess          *document.WriteAccess `protobuf:"bytes,3,opt,name=write_access,json=writeAccess,proto3" json:"write_access,omitempty"`
	Data                 *FundingData          `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FundingCreatePayload) Reset()         { *m = FundingCreatePayload{} }
func (m *FundingCreatePayload) String() string { return proto.CompactTextString(m) }
func (*FundingCreatePayload) ProtoMessage()    {}
func (*FundingCreatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{0}
}

func (m *FundingCreatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingCreatePayload.Unmarshal(m, b)
}
func (m *FundingCreatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingCreatePayload.Marshal(b, m, deterministic)
}
func (m *FundingCreatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingCreatePayload.Merge(m, src)
}
func (m *FundingCreatePayload) XXX_Size() int {
	return xxx_messageInfo_FundingCreatePayload.Size(m)
}
func (m *FundingCreatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingCreatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_FundingCreatePayload proto.InternalMessageInfo

func (m *FundingCreatePayload) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *FundingCreatePayload) GetReadAccess() *document.ReadAccess {
	if m != nil {
		return m.ReadAccess
	}
	return nil
}

func (m *FundingCreatePayload) GetWriteAccess() *document.WriteAccess {
	if m != nil {
		return m.WriteAccess
	}
	return nil
}

func (m *FundingCreatePayload) GetData() *FundingData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FundingUpdatePayload struct {
	Identifier           string                `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	FundingId            string                `protobuf:"bytes,2,opt,name=funding_id,json=fundingId,proto3" json:"funding_id,omitempty"`
	ReadAccess           *document.ReadAccess  `protobuf:"bytes,3,opt,name=read_access,json=readAccess,proto3" json:"read_access,omitempty"`
	WriteAccess          *document.WriteAccess `protobuf:"bytes,4,opt,name=write_access,json=writeAccess,proto3" json:"write_access,omitempty"`
	Data                 *FundingData          `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FundingUpdatePayload) Reset()         { *m = FundingUpdatePayload{} }
func (m *FundingUpdatePayload) String() string { return proto.CompactTextString(m) }
func (*FundingUpdatePayload) ProtoMessage()    {}
func (*FundingUpdatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{1}
}

func (m *FundingUpdatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingUpdatePayload.Unmarshal(m, b)
}
func (m *FundingUpdatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingUpdatePayload.Marshal(b, m, deterministic)
}
func (m *FundingUpdatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingUpdatePayload.Merge(m, src)
}
func (m *FundingUpdatePayload) XXX_Size() int {
	return xxx_messageInfo_FundingUpdatePayload.Size(m)
}
func (m *FundingUpdatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingUpdatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_FundingUpdatePayload proto.InternalMessageInfo

func (m *FundingUpdatePayload) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *FundingUpdatePayload) GetFundingId() string {
	if m != nil {
		return m.FundingId
	}
	return ""
}

func (m *FundingUpdatePayload) GetReadAccess() *document.ReadAccess {
	if m != nil {
		return m.ReadAccess
	}
	return nil
}

func (m *FundingUpdatePayload) GetWriteAccess() *document.WriteAccess {
	if m != nil {
		return m.WriteAccess
	}
	return nil
}

func (m *FundingUpdatePayload) GetData() *FundingData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FundingResponse struct {
	Header               *document.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 *FundingResponseData     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FundingResponse) Reset()         { *m = FundingResponse{} }
func (m *FundingResponse) String() string { return proto.CompactTextString(m) }
func (*FundingResponse) ProtoMessage()    {}
func (*FundingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{2}
}

func (m *FundingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingResponse.Unmarshal(m, b)
}
func (m *FundingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingResponse.Marshal(b, m, deterministic)
}
func (m *FundingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingResponse.Merge(m, src)
}
func (m *FundingResponse) XXX_Size() int {
	return xxx_messageInfo_FundingResponse.Size(m)
}
func (m *FundingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FundingResponse proto.InternalMessageInfo

func (m *FundingResponse) GetHeader() *document.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *FundingResponse) GetData() *FundingResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FundingListResponse struct {
	Header               *document.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 []*FundingResponseData   `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FundingListResponse) Reset()         { *m = FundingListResponse{} }
func (m *FundingListResponse) String() string { return proto.CompactTextString(m) }
func (*FundingListResponse) ProtoMessage()    {}
func (*FundingListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{3}
}

func (m *FundingListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingListResponse.Unmarshal(m, b)
}
func (m *FundingListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingListResponse.Marshal(b, m, deterministic)
}
func (m *FundingListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingListResponse.Merge(m, src)
}
func (m *FundingListResponse) XXX_Size() int {
	return xxx_messageInfo_FundingListResponse.Size(m)
}
func (m *FundingListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FundingListResponse proto.InternalMessageInfo

func (m *FundingListResponse) GetHeader() *document.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *FundingListResponse) GetData() []*FundingResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type Request struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	FundingId            string   `protobuf:"bytes,2,opt,name=funding_id,json=fundingId,proto3" json:"funding_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{4}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Request) GetFundingId() string {
	if m != nil {
		return m.FundingId
	}
	return ""
}

type GetVersionRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	FundingId            string   `protobuf:"bytes,3,opt,name=funding_id,json=fundingId,proto3" json:"funding_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionRequest) Reset()         { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()    {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{5}
}

func (m *GetVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionRequest.Unmarshal(m, b)
}
func (m *GetVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionRequest.Marshal(b, m, deterministic)
}
func (m *GetVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionRequest.Merge(m, src)
}
func (m *GetVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetVersionRequest.Size(m)
}
func (m *GetVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionRequest proto.InternalMessageInfo

func (m *GetVersionRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *GetVersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GetVersionRequest) GetFundingId() string {
	if m != nil {
		return m.FundingId
	}
	return ""
}

type GetListRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListRequest) Reset()         { *m = GetListRequest{} }
func (m *GetListRequest) String() string { return proto.CompactTextString(m) }
func (*GetListRequest) ProtoMessage()    {}
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{6}
}

func (m *GetListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListRequest.Unmarshal(m, b)
}
func (m *GetListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListRequest.Marshal(b, m, deterministic)
}
func (m *GetListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListRequest.Merge(m, src)
}
func (m *GetListRequest) XXX_Size() int {
	return xxx_messageInfo_GetListRequest.Size(m)
}
func (m *GetListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetListRequest proto.InternalMessageInfo

func (m *GetListRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type GetListVersionRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListVersionRequest) Reset()         { *m = GetListVersionRequest{} }
func (m *GetListVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetListVersionRequest) ProtoMessage()    {}
func (*GetListVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{7}
}

func (m *GetListVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListVersionRequest.Unmarshal(m, b)
}
func (m *GetListVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListVersionRequest.Marshal(b, m, deterministic)
}
func (m *GetListVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListVersionRequest.Merge(m, src)
}
func (m *GetListVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetListVersionRequest.Size(m)
}
func (m *GetListVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetListVersionRequest proto.InternalMessageInfo

func (m *GetListVersionRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *GetListVersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// FundingData is the default funding extension schema
type FundingData struct {
	FundingId             string   `protobuf:"bytes,1,opt,name=funding_id,json=fundingId,proto3" json:"funding_id,omitempty"`
	Amount                string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Apr                   string   `protobuf:"bytes,3,opt,name=apr,proto3" json:"apr,omitempty"`
	Days                  string   `protobuf:"bytes,4,opt,name=days,proto3" json:"days,omitempty"`
	Fee                   string   `protobuf:"bytes,5,opt,name=fee,proto3" json:"fee,omitempty"`
	RepaymentDueDate      string   `protobuf:"bytes,7,opt,name=repayment_due_date,json=repaymentDueDate,proto3" json:"repayment_due_date,omitempty"`
	RepaymentOccurredDate string   `protobuf:"bytes,8,opt,name=repayment_occurred_date,json=repaymentOccurredDate,proto3" json:"repayment_occurred_date,omitempty"`
	RepaymentAmount       string   `protobuf:"bytes,9,opt,name=repayment_amount,json=repaymentAmount,proto3" json:"repayment_amount,omitempty"`
	Currency              string   `protobuf:"bytes,10,opt,name=currency,proto3" json:"currency,omitempty"`
	NftAddress            string   `protobuf:"bytes,11,opt,name=nft_address,json=nftAddress,proto3" json:"nft_address,omitempty"`
	PaymentDetailsId      string   `protobuf:"bytes,12,opt,name=payment_details_id,json=paymentDetailsId,proto3" json:"payment_details_id,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *FundingData) Reset()         { *m = FundingData{} }
func (m *FundingData) String() string { return proto.CompactTextString(m) }
func (*FundingData) ProtoMessage()    {}
func (*FundingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{8}
}

func (m *FundingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingData.Unmarshal(m, b)
}
func (m *FundingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingData.Marshal(b, m, deterministic)
}
func (m *FundingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingData.Merge(m, src)
}
func (m *FundingData) XXX_Size() int {
	return xxx_messageInfo_FundingData.Size(m)
}
func (m *FundingData) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingData.DiscardUnknown(m)
}

var xxx_messageInfo_FundingData proto.InternalMessageInfo

func (m *FundingData) GetFundingId() string {
	if m != nil {
		return m.FundingId
	}
	return ""
}

func (m *FundingData) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *FundingData) GetApr() string {
	if m != nil {
		return m.Apr
	}
	return ""
}

func (m *FundingData) GetDays() string {
	if m != nil {
		return m.Days
	}
	return ""
}

func (m *FundingData) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *FundingData) GetRepaymentDueDate() string {
	if m != nil {
		return m.RepaymentDueDate
	}
	return ""
}

func (m *FundingData) GetRepaymentOccurredDate() string {
	if m != nil {
		return m.RepaymentOccurredDate
	}
	return ""
}

func (m *FundingData) GetRepaymentAmount() string {
	if m != nil {
		return m.RepaymentAmount
	}
	return ""
}

func (m *FundingData) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *FundingData) GetNftAddress() string {
	if m != nil {
		return m.NftAddress
	}
	return ""
}

func (m *FundingData) GetPaymentDetailsId() string {
	if m != nil {
		return m.PaymentDetailsId
	}
	return ""
}

type FundingResponseData struct {
	Funding              *FundingData        `protobuf:"bytes,1,opt,name=funding,proto3" json:"funding,omitempty"`
	Signatures           []*FundingSignature `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FundingResponseData) Reset()         { *m = FundingResponseData{} }
func (m *FundingResponseData) String() string { return proto.CompactTextString(m) }
func (*FundingResponseData) ProtoMessage()    {}
func (*FundingResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{9}
}

func (m *FundingResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingResponseData.Unmarshal(m, b)
}
func (m *FundingResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingResponseData.Marshal(b, m, deterministic)
}
func (m *FundingResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingResponseData.Merge(m, src)
}
func (m *FundingResponseData) XXX_Size() int {
	return xxx_messageInfo_FundingResponseData.Size(m)
}
func (m *FundingResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_FundingResponseData proto.InternalMessageInfo

func (m *FundingResponseData) GetFunding() *FundingData {
	if m != nil {
		return m.Funding
	}
	return nil
}

func (m *FundingResponseData) GetSignatures() []*FundingSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type FundingSignature struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	OutdatedSignature    bool     `protobuf:"varint,2,opt,name=outdated_signature,json=outdatedSignature,proto3" json:"outdated_signature,omitempty"`
	Identity             string   `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	SignedVersion        string   `protobuf:"bytes,4,opt,name=signed_version,json=signedVersion,proto3" json:"signed_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FundingSignature) Reset()         { *m = FundingSignature{} }
func (m *FundingSignature) String() string { return proto.CompactTextString(m) }
func (*FundingSignature) ProtoMessage()    {}
func (*FundingSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{10}
}

func (m *FundingSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingSignature.Unmarshal(m, b)
}
func (m *FundingSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingSignature.Marshal(b, m, deterministic)
}
func (m *FundingSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingSignature.Merge(m, src)
}
func (m *FundingSignature) XXX_Size() int {
	return xxx_messageInfo_FundingSignature.Size(m)
}
func (m *FundingSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingSignature.DiscardUnknown(m)
}

var xxx_messageInfo_FundingSignature proto.InternalMessageInfo

func (m *FundingSignature) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *FundingSignature) GetOutdatedSignature() bool {
	if m != nil {
		return m.OutdatedSignature
	}
	return false
}

func (m *FundingSignature) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *FundingSignature) GetSignedVersion() string {
	if m != nil {
		return m.SignedVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*FundingCreatePayload)(nil), "funding.FundingCreatePayload")
	proto.RegisterType((*FundingUpdatePayload)(nil), "funding.FundingUpdatePayload")
	proto.RegisterType((*FundingResponse)(nil), "funding.FundingResponse")
	proto.RegisterType((*FundingListResponse)(nil), "funding.FundingListResponse")
	proto.RegisterType((*Request)(nil), "funding.Request")
	proto.RegisterType((*GetVersionRequest)(nil), "funding.GetVersionRequest")
	proto.RegisterType((*GetListRequest)(nil), "funding.GetListRequest")
	proto.RegisterType((*GetListVersionRequest)(nil), "funding.GetListVersionRequest")
	proto.RegisterType((*FundingData)(nil), "funding.FundingData")
	proto.RegisterType((*FundingResponseData)(nil), "funding.FundingResponseData")
	proto.RegisterType((*FundingSignature)(nil), "funding.FundingSignature")
}

func init() { proto.RegisterFile("funding/funding.proto", fileDescriptor_e74e73a4c22d632c) }

var fileDescriptor_e74e73a4c22d632c = []byte{
	// 1004 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcd, 0x6e, 0x23, 0xc5,
	0x13, 0xc0, 0x35, 0x71, 0xd6, 0x1f, 0xe5, 0xfc, 0xb3, 0xd9, 0xfe, 0x27, 0xbb, 0x66, 0x94, 0xcd,
	0x0e, 0x83, 0x56, 0x24, 0x24, 0xf1, 0x98, 0xa0, 0x5d, 0x2d, 0xdc, 0x66, 0x59, 0x91, 0x5d, 0x84,
	0x44, 0x18, 0x04, 0x48, 0x70, 0xb0, 0x7a, 0xdd, 0x65, 0x33, 0x92, 0x33, 0x63, 0x66, 0x7a, 0x12,
	0x59, 0xd1, 0x4a, 0x88, 0x3b, 0x07, 0xcc, 0x05, 0x90, 0x38, 0xf3, 0x02, 0x88, 0xf7, 0x40, 0xbc,
	0x02, 0x0f, 0xc0, 0x91, 0x23, 0xea, 0xaf, 0xf1, 0x64, 0x9c, 0x38, 0x0e, 0xec, 0xc9, 0xd3, 0x5d,
	0x55, 0x5d, 0xbf, 0xaa, 0xae, 0xea, 0x32, 0x6c, 0xf4, 0xb3, 0x88, 0x85, 0xd1, 0xc0, 0xd3, 0xbf,
	0xed, 0x51, 0x12, 0xf3, 0x98, 0xd4, 0xf4, 0xd2, 0xbe, 0xcd, 0xe2, 0x5e, 0x76, 0x8c, 0x11, 0xf7,
	0x52, 0x4c, 0x4e, 0xc2, 0x1e, 0x2a, 0x05, 0x7b, 0x73, 0x10, 0xc7, 0x83, 0x21, 0x7a, 0x74, 0x14,
	0x7a, 0x34, 0x8a, 0x62, 0x4e, 0x79, 0x18, 0x47, 0xa9, 0x96, 0xee, 0xc9, 0x9f, 0xde, 0xfe, 0x00,
	0xa3, 0xfd, 0xf4, 0x94, 0x0e, 0x06, 0x98, 0x78, 0xf1, 0x48, 0x6a, 0xcc, 0x6a, 0xbb, 0xbf, 0x5b,
	0xb0, 0xfe, 0x9e, 0xf2, 0xf7, 0x6e, 0x82, 0x94, 0xe3, 0x11, 0x1d, 0x0f, 0x63, 0xca, 0xc8, 0x16,
	0x40, 0xc8, 0x30, 0xe2, 0x61, 0x3f, 0xc4, 0xa4, 0x65, 0x39, 0xd6, 0x76, 0x23, 0x28, 0xec, 0x90,
	0x07, 0xd0, 0x4c, 0x90, 0xb2, 0x2e, 0xed, 0xf5, 0x30, 0x4d, 0x5b, 0x4b, 0x8e, 0xb5, 0xdd, 0x3c,
	0x58, 0x6f, 0x1b, 0xe4, 0x76, 0x80, 0x94, 0xf9, 0x52, 0x16, 0x40, 0x92, 0x7f, 0x93, 0x47, 0xb0,
	0x72, 0x9a, 0x84, 0x1c, 0x8d, 0x5d, 0x45, 0xda, 0x6d, 0x4c, 0xed, 0x3e, 0x13, 0x52, 0x6d, 0xd8,
	0x3c, 0x9d, 0x2e, 0xc8, 0x36, 0x2c, 0x33, 0xca, 0x69, 0x6b, 0x59, 0x7b, 0x32, 0x49, 0xd3, 0xf4,
	0x4f, 0x28, 0xa7, 0x81, 0xd4, 0x70, 0xff, 0x9a, 0xc6, 0xf4, 0xc9, 0x88, 0x5d, 0x23, 0xa6, 0xbb,
	0x00, 0xfa, 0xd4, 0x6e, 0xc8, 0x64, 0x48, 0x8d, 0xa0, 0xa1, 0x77, 0x9e, 0xb1, 0x72, 0xc8, 0x95,
	0x7f, 0x19, 0xf2, 0xf2, 0xb5, 0x43, 0xbe, 0x71, 0x65, 0xc8, 0x19, 0xdc, 0xd4, 0x9b, 0x01, 0xa6,
	0xa3, 0x38, 0x4a, 0x91, 0x74, 0xa0, 0xfa, 0x25, 0x52, 0xa6, 0x03, 0x6d, 0x1e, 0xb4, 0x8a, 0xa0,
	0x4a, 0xe7, 0xa9, 0x94, 0x07, 0x5a, 0x8f, 0x74, 0xb4, 0x3b, 0x75, 0x97, 0x9b, 0x65, 0x77, 0xc6,
	0xaa, 0xe0, 0x76, 0x0c, 0xff, 0xd7, 0xc2, 0x0f, 0xc2, 0x94, 0xbf, 0x14, 0xd7, 0x95, 0x05, 0x5d,
	0x3f, 0x85, 0x5a, 0x80, 0x5f, 0x65, 0x98, 0xf2, 0xff, 0x78, 0xad, 0xee, 0x10, 0x6e, 0x1d, 0x22,
	0xff, 0x14, 0x93, 0x34, 0x8c, 0xa3, 0x45, 0xcf, 0x6c, 0x41, 0xed, 0x44, 0x59, 0xe8, 0x03, 0xcd,
	0xb2, 0xe4, 0xad, 0x52, 0xf6, 0xd6, 0x81, 0xd5, 0x43, 0xe4, 0x2a, 0x5d, 0x0b, 0xb9, 0x72, 0x3f,
	0x82, 0x0d, 0x6d, 0xf1, 0xb2, 0x18, 0xdd, 0xbf, 0x97, 0xa0, 0x59, 0x28, 0xa2, 0x12, 0xb3, 0x55,
	0x2e, 0xfc, 0xdb, 0x50, 0xa5, 0xc7, 0x71, 0x16, 0x71, 0x7d, 0x8e, 0x5e, 0x91, 0x35, 0xa8, 0xd0,
	0x51, 0xa2, 0x63, 0x14, 0x9f, 0x84, 0x88, 0x7b, 0x1c, 0xab, 0x1a, 0x6f, 0x04, 0xf2, 0x5b, 0x68,
	0xf5, 0x11, 0x65, 0x11, 0x37, 0x02, 0xf1, 0x49, 0xf6, 0x80, 0x24, 0x38, 0xa2, 0x63, 0x51, 0x11,
	0x5d, 0x96, 0x61, 0x57, 0x34, 0x69, 0xab, 0x26, 0x15, 0xd6, 0x72, 0xc9, 0x93, 0x4c, 0xdc, 0x37,
	0x92, 0x87, 0x70, 0x67, 0xaa, 0x1d, 0xf7, 0x7a, 0x59, 0x92, 0x20, 0x53, 0x26, 0x75, 0x69, 0xb2,
	0x91, 0x8b, 0x3f, 0xd4, 0x52, 0x69, 0xb7, 0x03, 0xd3, 0xb3, 0xba, 0x9a, 0xbf, 0x21, 0x0d, 0x6e,
	0xe6, 0xfb, 0xbe, 0x0a, 0xc4, 0x86, 0xba, 0x34, 0x8c, 0x7a, 0xe3, 0x16, 0x48, 0x95, 0x7c, 0x4d,
	0xee, 0x41, 0x33, 0xea, 0xf3, 0x2e, 0x65, 0x2c, 0x11, 0xdd, 0xdb, 0x54, 0x69, 0x8e, 0xfa, 0xdc,
	0x57, 0x3b, 0x22, 0x9a, 0x3c, 0x16, 0xe4, 0x34, 0x1c, 0xa6, 0x22, 0x89, 0x2b, 0x2a, 0x1a, 0x13,
	0x8b, 0x12, 0x3c, 0x63, 0xee, 0xd7, 0x56, 0xde, 0x33, 0xc5, 0xaa, 0x26, 0x6d, 0x30, 0xef, 0xbe,
	0x6e, 0x9a, 0x8b, 0xdb, 0xdd, 0x28, 0x91, 0xb7, 0x01, 0xd2, 0x70, 0x10, 0x51, 0x9e, 0x25, 0x98,
	0xea, 0xbe, 0x79, 0xa5, 0x6c, 0xf2, 0xb1, 0xd1, 0x08, 0x0a, 0xca, 0xee, 0x4f, 0x16, 0xac, 0x95,
	0x15, 0xc8, 0x3a, 0xdc, 0x38, 0xa1, 0x43, 0x7d, 0xfb, 0xf5, 0x40, 0x2d, 0xc8, 0x3e, 0x90, 0x38,
	0xe3, 0x22, 0xd7, 0xac, 0x9b, 0x9f, 0x20, 0xab, 0xa0, 0x1e, 0xdc, 0x32, 0x92, 0xe9, 0x21, 0x36,
	0xd4, 0x55, 0xfd, 0xf1, 0xb1, 0xae, 0x8a, 0x7c, 0x4d, 0xee, 0xc3, 0xaa, 0x38, 0x01, 0x59, 0xd7,
	0x14, 0xa5, 0x2a, 0x92, 0xff, 0xa9, 0x5d, 0x5d, 0xdb, 0x07, 0xdf, 0x35, 0x60, 0xd5, 0xc0, 0xa9,
	0xa9, 0x47, 0xbe, 0xb5, 0xa0, 0xaa, 0x86, 0x13, 0xb9, 0x5b, 0x8e, 0xf0, 0xdc, 0xd0, 0xb2, 0x5b,
	0x97, 0x3d, 0x1c, 0xee, 0xfb, 0x13, 0x7f, 0xcb, 0xde, 0xf4, 0x19, 0x4b, 0x1d, 0xea, 0x68, 0x25,
	0x87, 0xc7, 0x0e, 0x75, 0xcc, 0xd3, 0xf4, 0xcd, 0x1f, 0x7f, 0x7e, 0xbf, 0xf4, 0x9a, 0xbb, 0xe5,
	0xe5, 0x53, 0xf7, 0x6c, 0xda, 0x4b, 0x2f, 0xcc, 0x8c, 0x7e, 0xc7, 0x7a, 0x83, 0xfc, 0x62, 0x41,
	0x55, 0x0d, 0x96, 0x59, 0x9e, 0x73, 0x03, 0x67, 0x0e, 0x0f, 0x9b, 0xf8, 0xbb, 0xf6, 0x8e, 0xd2,
	0x2e, 0x22, 0xd1, 0x41, 0x82, 0x28, 0xbc, 0x3b, 0x61, 0x54, 0x86, 0xeb, 0xd8, 0xbb, 0xf3, 0xe1,
	0xbc, 0xb3, 0x69, 0x17, 0xbf, 0x10, 0xa4, 0x3f, 0x5a, 0xb0, 0x2c, 0x6e, 0x87, 0xac, 0xe5, 0x20,
	0xfa, 0xf1, 0x98, 0x83, 0x86, 0x13, 0xff, 0x75, 0xfb, 0xbe, 0x30, 0x4b, 0x17, 0xc2, 0x7a, 0xe0,
	0x76, 0xae, 0x81, 0xe5, 0x89, 0xeb, 0x16, 0x6c, 0x3f, 0x58, 0x50, 0x39, 0x44, 0x7e, 0x5d, 0xb4,
	0xb6, 0xbd, 0x77, 0x88, 0xfc, 0xc2, 0x8c, 0xc5, 0x7d, 0x87, 0x3a, 0x43, 0x91, 0x50, 0x7e, 0x9e,
	0xb0, 0x4d, 0xf6, 0x72, 0xc2, 0xf4, 0x6a, 0x44, 0xf2, 0xab, 0x05, 0x30, 0x1d, 0x09, 0xc4, 0xce,
	0x79, 0x66, 0xe6, 0xc4, 0x1c, 0xd6, 0xd1, 0xc4, 0xf7, 0xec, 0xfd, 0xb9, 0xac, 0x86, 0xc9, 0xd1,
	0x1d, 0x21, 0x61, 0x1f, 0x91, 0x87, 0x97, 0xc1, 0x9e, 0x69, 0xbd, 0x4b, 0xb0, 0x7f, 0xb6, 0xa0,
	0xa6, 0x27, 0x05, 0xb9, 0x53, 0x64, 0x2e, 0x4c, 0x1b, 0x7b, 0x66, 0xb6, 0x16, 0x27, 0xb7, 0xfb,
	0xc5, 0xc4, 0x7f, 0xd3, 0xf6, 0x24, 0xf4, 0x70, 0x38, 0x8b, 0x9d, 0x5e, 0x9e, 0xe3, 0x57, 0xc9,
	0xbd, 0x2b, 0x72, 0x4c, 0x7e, 0xb3, 0xf2, 0xd9, 0x67, 0x52, 0xbb, 0x55, 0xc6, 0x2c, 0xa5, 0x77,
	0x3e, 0x6d, 0x7f, 0xe2, 0x1f, 0xd8, 0x9d, 0xab, 0x68, 0x2f, 0xcc, 0xf2, 0x2e, 0xd9, 0x59, 0x38,
	0xcb, 0x8f, 0xb7, 0xa1, 0xd9, 0x8b, 0x8f, 0x0d, 0xca, 0xe3, 0x15, 0xcd, 0x72, 0x24, 0xfe, 0x41,
	0x1f, 0x59, 0x9f, 0x9b, 0x49, 0x39, 0x7a, 0xfe, 0xbc, 0x2a, 0xff, 0x55, 0xbf, 0xf5, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0xde, 0xe1, 0xec, 0xdb, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FundingServiceClient is the client API for FundingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FundingServiceClient interface {
	Create(ctx context.Context, in *FundingCreatePayload, opts ...grpc.CallOption) (*FundingResponse, error)
	Update(ctx context.Context, in *FundingUpdatePayload, opts ...grpc.CallOption) (*FundingResponse, error)
	Sign(ctx context.Context, in *Request, opts ...grpc.CallOption) (*FundingResponse, error)
	Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*FundingResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*FundingResponse, error)
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*FundingListResponse, error)
	GetListVersion(ctx context.Context, in *GetListVersionRequest, opts ...grpc.CallOption) (*FundingListResponse, error)
}

type fundingServiceClient struct {
	cc *grpc.ClientConn
}

func NewFundingServiceClient(cc *grpc.ClientConn) FundingServiceClient {
	return &fundingServiceClient{cc}
}

func (c *fundingServiceClient) Create(ctx context.Context, in *FundingCreatePayload, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/funding.FundingService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) Update(ctx context.Context, in *FundingUpdatePayload, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/funding.FundingService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) Sign(ctx context.Context, in *Request, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/funding.FundingService/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/funding.FundingService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/funding.FundingService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*FundingListResponse, error) {
	out := new(FundingListResponse)
	err := c.cc.Invoke(ctx, "/funding.FundingService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) GetListVersion(ctx context.Context, in *GetListVersionRequest, opts ...grpc.CallOption) (*FundingListResponse, error) {
	out := new(FundingListResponse)
	err := c.cc.Invoke(ctx, "/funding.FundingService/GetListVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FundingServiceServer is the server API for FundingService service.
type FundingServiceServer interface {
	Create(context.Context, *FundingCreatePayload) (*FundingResponse, error)
	Update(context.Context, *FundingUpdatePayload) (*FundingResponse, error)
	Sign(context.Context, *Request) (*FundingResponse, error)
	Get(context.Context, *Request) (*FundingResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*FundingResponse, error)
	GetList(context.Context, *GetListRequest) (*FundingListResponse, error)
	GetListVersion(context.Context, *GetListVersionRequest) (*FundingListResponse, error)
}

func RegisterFundingServiceServer(s *grpc.Server, srv FundingServiceServer) {
	s.RegisterService(&_FundingService_serviceDesc, srv)
}

func _FundingService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FundingCreatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funding.FundingService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).Create(ctx, req.(*FundingCreatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FundingUpdatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funding.FundingService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).Update(ctx, req.(*FundingUpdatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funding.FundingService/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).Sign(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funding.FundingService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).Get(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funding.FundingService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funding.FundingService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_GetListVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).GetListVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funding.FundingService/GetListVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).GetListVersion(ctx, req.(*GetListVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FundingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "funding.FundingService",
	HandlerType: (*FundingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FundingService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FundingService_Update_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _FundingService_Sign_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _FundingService_Get_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _FundingService_GetVersion_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _FundingService_GetList_Handler,
		},
		{
			MethodName: "GetListVersion",
			Handler:    _FundingService_GetListVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "funding/funding.proto",
}
