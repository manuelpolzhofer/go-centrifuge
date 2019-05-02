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
	Identifier           string       `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Data                 *FundingData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
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

func (m *FundingCreatePayload) GetData() *FundingData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FundingResponse struct {
	Header               *document.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 *FundingData             `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FundingResponse) Reset()         { *m = FundingResponse{} }
func (m *FundingResponse) String() string { return proto.CompactTextString(m) }
func (*FundingResponse) ProtoMessage()    {}
func (*FundingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{1}
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

func (m *FundingResponse) GetData() *FundingData {
	if m != nil {
		return m.Data
	}
	return nil
}

// FundingData is the default funding extension schema
type FundingData struct {
	AgreementId           string   `protobuf:"bytes,1,opt,name=agreement_id,json=agreementId,proto3" json:"agreement_id,omitempty"`
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
	return fileDescriptor_e74e73a4c22d632c, []int{2}
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

func (m *FundingData) GetAgreementId() string {
	if m != nil {
		return m.AgreementId
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

func init() {
	proto.RegisterType((*FundingCreatePayload)(nil), "funding.FundingCreatePayload")
	proto.RegisterType((*FundingResponse)(nil), "funding.FundingResponse")
	proto.RegisterType((*FundingData)(nil), "funding.FundingData")
}

func init() { proto.RegisterFile("funding/funding.proto", fileDescriptor_e74e73a4c22d632c) }

var fileDescriptor_e74e73a4c22d632c = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0xba, 0xd1, 0xad, 0x37, 0x15, 0x9b, 0xac, 0x6d, 0x44, 0x15, 0x1f, 0xa3, 0x4f, 0x05,
	0x6d, 0x0d, 0x2a, 0x12, 0x48, 0x48, 0x3c, 0x6c, 0x54, 0x88, 0x3d, 0x51, 0x95, 0x37, 0x5e, 0xca,
	0x5d, 0x7c, 0x13, 0x2c, 0xb5, 0x76, 0x64, 0x3b, 0x40, 0x85, 0x90, 0x10, 0x7f, 0x61, 0x3f, 0x8d,
	0xbf, 0xc0, 0xcf, 0xe0, 0x01, 0xc5, 0x76, 0xd2, 0x31, 0x9e, 0x78, 0x4a, 0x7c, 0xce, 0xb9, 0xc7,
	0x3e, 0xf6, 0xbd, 0x70, 0x98, 0x57, 0x92, 0x0b, 0x59, 0xa4, 0xe1, 0x3b, 0x2e, 0xb5, 0xb2, 0x8a,
	0xed, 0x84, 0xe5, 0xe0, 0x88, 0xab, 0xac, 0x5a, 0x91, 0xb4, 0xa9, 0x21, 0xfd, 0x49, 0x64, 0xe4,
	0x05, 0x83, 0xbb, 0x85, 0x52, 0xc5, 0x92, 0x52, 0x2c, 0x45, 0x8a, 0x52, 0x2a, 0x8b, 0x56, 0x28,
	0x69, 0x02, 0x7b, 0xe2, 0x3e, 0xd9, 0x69, 0x41, 0xf2, 0xd4, 0x7c, 0xc6, 0xa2, 0x20, 0x9d, 0xaa,
	0xd2, 0x29, 0xfe, 0x55, 0x0f, 0x3f, 0xc0, 0xc1, 0x6b, 0xbf, 0xdd, 0x2b, 0x4d, 0x68, 0x69, 0x86,
	0xeb, 0xa5, 0x42, 0xce, 0xee, 0x03, 0x08, 0x4e, 0xd2, 0x8a, 0x5c, 0x90, 0x4e, 0xa2, 0xe3, 0x68,
	0xd4, 0x9b, 0x5f, 0x43, 0xd8, 0x08, 0xb6, 0x39, 0x5a, 0x4c, 0x3a, 0xc7, 0xd1, 0x28, 0x9e, 0x1c,
	0x8c, 0x9b, 0x08, 0xc1, 0x6c, 0x8a, 0x16, 0xe7, 0x4e, 0x31, 0x5c, 0xc1, 0x5e, 0x00, 0xe7, 0x64,
	0x4a, 0x25, 0x0d, 0xb1, 0x27, 0xd0, 0xfd, 0x48, 0xc8, 0x83, 0x71, 0x3c, 0x49, 0xc6, 0x4d, 0xd2,
	0x71, 0xa3, 0x79, 0xe3, 0xf8, 0x79, 0xd0, 0xfd, 0xc7, 0x76, 0xbf, 0x3b, 0x10, 0x5f, 0x43, 0xd9,
	0x43, 0xe8, 0x63, 0xa1, 0x89, 0x6a, 0xf7, 0x85, 0xe0, 0x21, 0x4a, 0xdc, 0x62, 0x17, 0x9c, 0x1d,
	0x41, 0x17, 0x57, 0xaa, 0x92, 0xd6, 0xd9, 0xf7, 0xe6, 0x61, 0xc5, 0xf6, 0x61, 0x0b, 0x4b, 0x9d,
	0x6c, 0x39, 0xb0, 0xfe, 0x65, 0xac, 0x3e, 0xc6, 0xda, 0x24, 0xdb, 0x0e, 0x72, 0xff, 0xb5, 0x2a,
	0x27, 0x4a, 0x6e, 0x79, 0x55, 0x4e, 0xc4, 0x4e, 0x80, 0x69, 0x2a, 0x71, 0xed, 0xb6, 0xe4, 0x15,
	0x2d, 0x38, 0x5a, 0x4a, 0x76, 0x9c, 0x60, 0xbf, 0x65, 0xa6, 0x15, 0x4d, 0xd1, 0x12, 0x7b, 0x06,
	0x77, 0x36, 0x6a, 0x95, 0x65, 0x95, 0xd6, 0xc4, 0x7d, 0xc9, 0xae, 0x2b, 0x39, 0x6c, 0xe9, 0xb7,
	0x81, 0x75, 0x75, 0x8f, 0x60, 0xe3, 0xb5, 0x08, 0xe7, 0xef, 0xb9, 0x82, 0xbd, 0x16, 0x3f, 0xf3,
	0x41, 0x06, 0xb0, 0xeb, 0x0a, 0x65, 0xb6, 0x4e, 0xc0, 0x49, 0xda, 0x35, 0x7b, 0x00, 0xb1, 0xcc,
	0xed, 0x02, 0x39, 0xd7, 0x64, 0x4c, 0x12, 0xfb, 0x97, 0x96, 0xb9, 0x3d, 0xf3, 0x48, 0x9d, 0xa6,
	0xcd, 0x42, 0x16, 0xc5, 0xd2, 0xd4, 0xd7, 0xd8, 0xf7, 0x69, 0x9a, 0x2c, 0x9e, 0xb8, 0xe0, 0x93,
	0xab, 0x08, 0x6e, 0x87, 0xeb, 0x7f, 0xe7, 0x9b, 0x96, 0x7d, 0x8f, 0xa0, 0xeb, 0x9b, 0x8b, 0xdd,
	0xbb, 0xf9, 0x70, 0x7f, 0x35, 0xdd, 0x20, 0xb9, 0x49, 0x37, 0xdd, 0x30, 0x7c, 0xf9, 0xe3, 0xe7,
	0xaf, 0xab, 0xce, 0xf3, 0xe1, 0x24, 0x6d, 0x67, 0xe2, 0xeb, 0xa6, 0x1b, 0xbf, 0xa5, 0xf4, 0xc5,
	0x92, 0x34, 0x42, 0xc9, 0x66, 0x96, 0xd2, 0xf6, 0x81, 0x5f, 0x44, 0x8f, 0xcf, 0x47, 0x10, 0x67,
	0x6a, 0xd5, 0xb8, 0x9f, 0xf7, 0x83, 0xfd, 0xac, 0x1e, 0x81, 0x59, 0xf4, 0xbe, 0x17, 0x88, 0xf2,
	0xf2, 0xb2, 0xeb, 0xc6, 0xe2, 0xe9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x73, 0x68, 0xab,
	0x9c, 0x03, 0x00, 0x00,
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

// FundingServiceServer is the server API for FundingService service.
type FundingServiceServer interface {
	Create(context.Context, *FundingCreatePayload) (*FundingResponse, error)
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

var _FundingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "funding.FundingService",
	HandlerType: (*FundingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FundingService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "funding/funding.proto",
}
