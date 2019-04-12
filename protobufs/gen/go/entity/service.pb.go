// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity/service.proto

package entitypb

import (
	context "context"
	fmt "fmt"
	entity "github.com/centrifuge/centrifuge-protobufs/gen/go/entity"
	document "github.com/centrifuge/go-centrifuge/protobufs/gen/go/document"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
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

type GetRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type GetRequestRelationship struct {
	RelationshipIdentifier string   `protobuf:"bytes,1,opt,name=relationship_identifier,json=relationshipIdentifier,proto3" json:"relationship_identifier,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *GetRequestRelationship) Reset()         { *m = GetRequestRelationship{} }
func (m *GetRequestRelationship) String() string { return proto.CompactTextString(m) }
func (*GetRequestRelationship) ProtoMessage()    {}
func (*GetRequestRelationship) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{1}
}

func (m *GetRequestRelationship) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequestRelationship.Unmarshal(m, b)
}
func (m *GetRequestRelationship) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequestRelationship.Marshal(b, m, deterministic)
}
func (m *GetRequestRelationship) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequestRelationship.Merge(m, src)
}
func (m *GetRequestRelationship) XXX_Size() int {
	return xxx_messageInfo_GetRequestRelationship.Size(m)
}
func (m *GetRequestRelationship) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequestRelationship.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequestRelationship proto.InternalMessageInfo

func (m *GetRequestRelationship) GetRelationshipIdentifier() string {
	if m != nil {
		return m.RelationshipIdentifier
	}
	return ""
}

type GetVersionRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionRequest) Reset()         { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()    {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{2}
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

type EntityCreatePayload struct {
	ReadAccess           *document.ReadAccess  `protobuf:"bytes,1,opt,name=read_access,json=readAccess,proto3" json:"read_access,omitempty"`
	WriteAccess          *document.WriteAccess `protobuf:"bytes,2,opt,name=write_access,json=writeAccess,proto3" json:"write_access,omitempty"`
	Data                 *EntityData           `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *EntityCreatePayload) Reset()         { *m = EntityCreatePayload{} }
func (m *EntityCreatePayload) String() string { return proto.CompactTextString(m) }
func (*EntityCreatePayload) ProtoMessage()    {}
func (*EntityCreatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{3}
}

func (m *EntityCreatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityCreatePayload.Unmarshal(m, b)
}
func (m *EntityCreatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityCreatePayload.Marshal(b, m, deterministic)
}
func (m *EntityCreatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityCreatePayload.Merge(m, src)
}
func (m *EntityCreatePayload) XXX_Size() int {
	return xxx_messageInfo_EntityCreatePayload.Size(m)
}
func (m *EntityCreatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityCreatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_EntityCreatePayload proto.InternalMessageInfo

func (m *EntityCreatePayload) GetReadAccess() *document.ReadAccess {
	if m != nil {
		return m.ReadAccess
	}
	return nil
}

func (m *EntityCreatePayload) GetWriteAccess() *document.WriteAccess {
	if m != nil {
		return m.WriteAccess
	}
	return nil
}

func (m *EntityCreatePayload) GetData() *EntityData {
	if m != nil {
		return m.Data
	}
	return nil
}

type EntityUpdatePayload struct {
	Identifier           string                `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	ReadAccess           *document.ReadAccess  `protobuf:"bytes,2,opt,name=read_access,json=readAccess,proto3" json:"read_access,omitempty"`
	WriteAccess          *document.WriteAccess `protobuf:"bytes,3,opt,name=write_access,json=writeAccess,proto3" json:"write_access,omitempty"`
	Data                 *EntityData           `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *EntityUpdatePayload) Reset()         { *m = EntityUpdatePayload{} }
func (m *EntityUpdatePayload) String() string { return proto.CompactTextString(m) }
func (*EntityUpdatePayload) ProtoMessage()    {}
func (*EntityUpdatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{4}
}

func (m *EntityUpdatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityUpdatePayload.Unmarshal(m, b)
}
func (m *EntityUpdatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityUpdatePayload.Marshal(b, m, deterministic)
}
func (m *EntityUpdatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityUpdatePayload.Merge(m, src)
}
func (m *EntityUpdatePayload) XXX_Size() int {
	return xxx_messageInfo_EntityUpdatePayload.Size(m)
}
func (m *EntityUpdatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityUpdatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_EntityUpdatePayload proto.InternalMessageInfo

func (m *EntityUpdatePayload) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *EntityUpdatePayload) GetReadAccess() *document.ReadAccess {
	if m != nil {
		return m.ReadAccess
	}
	return nil
}

func (m *EntityUpdatePayload) GetWriteAccess() *document.WriteAccess {
	if m != nil {
		return m.WriteAccess
	}
	return nil
}

func (m *EntityUpdatePayload) GetData() *EntityData {
	if m != nil {
		return m.Data
	}
	return nil
}

type EntityResponse struct {
	Header               *document.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 *EntityDataResponse      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *EntityResponse) Reset()         { *m = EntityResponse{} }
func (m *EntityResponse) String() string { return proto.CompactTextString(m) }
func (*EntityResponse) ProtoMessage()    {}
func (*EntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{5}
}

func (m *EntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityResponse.Unmarshal(m, b)
}
func (m *EntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityResponse.Marshal(b, m, deterministic)
}
func (m *EntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityResponse.Merge(m, src)
}
func (m *EntityResponse) XXX_Size() int {
	return xxx_messageInfo_EntityResponse.Size(m)
}
func (m *EntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EntityResponse proto.InternalMessageInfo

func (m *EntityResponse) GetHeader() *document.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *EntityResponse) GetData() *EntityDataResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

type Relationship struct {
	Identity             string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Active               bool     `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Relationship) Reset()         { *m = Relationship{} }
func (m *Relationship) String() string { return proto.CompactTextString(m) }
func (*Relationship) ProtoMessage()    {}
func (*Relationship) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{6}
}

func (m *Relationship) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Relationship.Unmarshal(m, b)
}
func (m *Relationship) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Relationship.Marshal(b, m, deterministic)
}
func (m *Relationship) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Relationship.Merge(m, src)
}
func (m *Relationship) XXX_Size() int {
	return xxx_messageInfo_Relationship.Size(m)
}
func (m *Relationship) XXX_DiscardUnknown() {
	xxx_messageInfo_Relationship.DiscardUnknown(m)
}

var xxx_messageInfo_Relationship proto.InternalMessageInfo

func (m *Relationship) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *Relationship) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

// EntityData is the default entity schema
type EntityData struct {
	Identity  string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	LegalName string `protobuf:"bytes,2,opt,name=legal_name,json=legalName,proto3" json:"legal_name,omitempty"`
	// address
	Addresses []*entity.Address `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// tax information
	PaymentDetails []*entity.PaymentDetail `protobuf:"bytes,4,rep,name=payment_details,json=paymentDetails,proto3" json:"payment_details,omitempty"`
	// Entity contact list
	Contacts             []*entity.Contact `protobuf:"bytes,5,rep,name=contacts,proto3" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EntityData) Reset()         { *m = EntityData{} }
func (m *EntityData) String() string { return proto.CompactTextString(m) }
func (*EntityData) ProtoMessage()    {}
func (*EntityData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{7}
}

func (m *EntityData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityData.Unmarshal(m, b)
}
func (m *EntityData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityData.Marshal(b, m, deterministic)
}
func (m *EntityData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityData.Merge(m, src)
}
func (m *EntityData) XXX_Size() int {
	return xxx_messageInfo_EntityData.Size(m)
}
func (m *EntityData) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityData.DiscardUnknown(m)
}

var xxx_messageInfo_EntityData proto.InternalMessageInfo

func (m *EntityData) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *EntityData) GetLegalName() string {
	if m != nil {
		return m.LegalName
	}
	return ""
}

func (m *EntityData) GetAddresses() []*entity.Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *EntityData) GetPaymentDetails() []*entity.PaymentDetail {
	if m != nil {
		return m.PaymentDetails
	}
	return nil
}

func (m *EntityData) GetContacts() []*entity.Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

// Entity Relationships
type EntityDataResponse struct {
	Entity               *EntityData     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Relationships        []*Relationship `protobuf:"bytes,2,rep,name=relationships,proto3" json:"relationships,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EntityDataResponse) Reset()         { *m = EntityDataResponse{} }
func (m *EntityDataResponse) String() string { return proto.CompactTextString(m) }
func (*EntityDataResponse) ProtoMessage()    {}
func (*EntityDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{8}
}

func (m *EntityDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityDataResponse.Unmarshal(m, b)
}
func (m *EntityDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityDataResponse.Marshal(b, m, deterministic)
}
func (m *EntityDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityDataResponse.Merge(m, src)
}
func (m *EntityDataResponse) XXX_Size() int {
	return xxx_messageInfo_EntityDataResponse.Size(m)
}
func (m *EntityDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EntityDataResponse proto.InternalMessageInfo

func (m *EntityDataResponse) GetEntity() *EntityData {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *EntityDataResponse) GetRelationships() []*Relationship {
	if m != nil {
		return m.Relationships
	}
	return nil
}

type RelationshipPayload struct {
	// entity identifier
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	TargetIdentity       string   `protobuf:"bytes,2,opt,name=target_identity,json=targetIdentity,proto3" json:"target_identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelationshipPayload) Reset()         { *m = RelationshipPayload{} }
func (m *RelationshipPayload) String() string { return proto.CompactTextString(m) }
func (*RelationshipPayload) ProtoMessage()    {}
func (*RelationshipPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{9}
}

func (m *RelationshipPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationshipPayload.Unmarshal(m, b)
}
func (m *RelationshipPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationshipPayload.Marshal(b, m, deterministic)
}
func (m *RelationshipPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationshipPayload.Merge(m, src)
}
func (m *RelationshipPayload) XXX_Size() int {
	return xxx_messageInfo_RelationshipPayload.Size(m)
}
func (m *RelationshipPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationshipPayload.DiscardUnknown(m)
}

var xxx_messageInfo_RelationshipPayload proto.InternalMessageInfo

func (m *RelationshipPayload) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *RelationshipPayload) GetTargetIdentity() string {
	if m != nil {
		return m.TargetIdentity
	}
	return ""
}

type RelationshipData struct {
	// DID of relationship owner
	OwnerIdentity string `protobuf:"bytes,1,opt,name=owner_identity,json=ownerIdentity,proto3" json:"owner_identity,omitempty"`
	// DID of target identity
	TargetIdentity string `protobuf:"bytes,2,opt,name=target_identity,json=targetIdentity,proto3" json:"target_identity,omitempty"`
	// identifier of Entity whose data can be accessed via this relationship
	EntityIdentifier     string   `protobuf:"bytes,3,opt,name=entity_identifier,json=entityIdentifier,proto3" json:"entity_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelationshipData) Reset()         { *m = RelationshipData{} }
func (m *RelationshipData) String() string { return proto.CompactTextString(m) }
func (*RelationshipData) ProtoMessage()    {}
func (*RelationshipData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{10}
}

func (m *RelationshipData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationshipData.Unmarshal(m, b)
}
func (m *RelationshipData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationshipData.Marshal(b, m, deterministic)
}
func (m *RelationshipData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationshipData.Merge(m, src)
}
func (m *RelationshipData) XXX_Size() int {
	return xxx_messageInfo_RelationshipData.Size(m)
}
func (m *RelationshipData) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationshipData.DiscardUnknown(m)
}

var xxx_messageInfo_RelationshipData proto.InternalMessageInfo

func (m *RelationshipData) GetOwnerIdentity() string {
	if m != nil {
		return m.OwnerIdentity
	}
	return ""
}

func (m *RelationshipData) GetTargetIdentity() string {
	if m != nil {
		return m.TargetIdentity
	}
	return ""
}

func (m *RelationshipData) GetEntityIdentifier() string {
	if m != nil {
		return m.EntityIdentifier
	}
	return ""
}

type RelationshipResponse struct {
	Header               *document.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Relationship         []*RelationshipData      `protobuf:"bytes,2,rep,name=relationship,proto3" json:"relationship,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RelationshipResponse) Reset()         { *m = RelationshipResponse{} }
func (m *RelationshipResponse) String() string { return proto.CompactTextString(m) }
func (*RelationshipResponse) ProtoMessage()    {}
func (*RelationshipResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{11}
}

func (m *RelationshipResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationshipResponse.Unmarshal(m, b)
}
func (m *RelationshipResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationshipResponse.Marshal(b, m, deterministic)
}
func (m *RelationshipResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationshipResponse.Merge(m, src)
}
func (m *RelationshipResponse) XXX_Size() int {
	return xxx_messageInfo_RelationshipResponse.Size(m)
}
func (m *RelationshipResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationshipResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RelationshipResponse proto.InternalMessageInfo

func (m *RelationshipResponse) GetHeader() *document.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *RelationshipResponse) GetRelationship() []*RelationshipData {
	if m != nil {
		return m.Relationship
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "entity.GetRequest")
	proto.RegisterType((*GetRequestRelationship)(nil), "entity.GetRequestRelationship")
	proto.RegisterType((*GetVersionRequest)(nil), "entity.GetVersionRequest")
	proto.RegisterType((*EntityCreatePayload)(nil), "entity.EntityCreatePayload")
	proto.RegisterType((*EntityUpdatePayload)(nil), "entity.EntityUpdatePayload")
	proto.RegisterType((*EntityResponse)(nil), "entity.EntityResponse")
	proto.RegisterType((*Relationship)(nil), "entity.Relationship")
	proto.RegisterType((*EntityData)(nil), "entity.EntityData")
	proto.RegisterType((*EntityDataResponse)(nil), "entity.EntityDataResponse")
	proto.RegisterType((*RelationshipPayload)(nil), "entity.RelationshipPayload")
	proto.RegisterType((*RelationshipData)(nil), "entity.RelationshipData")
	proto.RegisterType((*RelationshipResponse)(nil), "entity.RelationshipResponse")
}

func init() { proto.RegisterFile("entity/service.proto", fileDescriptor_c1b437217b9e14a2) }

var fileDescriptor_c1b437217b9e14a2 = []byte{
	// 989 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x17, 0x06, 0x2d, 0x47, 0xb6, 0x8f, 0x64, 0x3b, 0x1e, 0x3b, 0xb2, 0x7e, 0x26, 0x7f, 0xc2, 0x12,
	0x4d, 0xe2, 0xc6, 0xb1, 0x14, 0xb8, 0x08, 0x52, 0x04, 0x45, 0x01, 0x39, 0x29, 0xdc, 0x14, 0xbd,
	0xb8, 0x0c, 0x7a, 0x41, 0x0b, 0x54, 0x18, 0x93, 0x27, 0x14, 0x5b, 0x89, 0x64, 0x67, 0xc6, 0x16,
	0x84, 0x34, 0x9b, 0x2e, 0xba, 0x0d, 0xa0, 0x3c, 0x43, 0xb7, 0x79, 0x94, 0x6e, 0xba, 0xec, 0xb2,
	0x7d, 0x90, 0x82, 0x33, 0x43, 0x71, 0x28, 0x4b, 0x89, 0x9b, 0xae, 0xa4, 0x39, 0xb7, 0xef, 0x3b,
	0x97, 0x39, 0x43, 0xd8, 0xc2, 0x58, 0x44, 0x62, 0xd4, 0xe6, 0xc8, 0x4e, 0x23, 0x1f, 0x5b, 0x29,
	0x4b, 0x44, 0x42, 0xaa, 0x4a, 0x6a, 0x37, 0x82, 0xc4, 0x3f, 0x19, 0x60, 0x2c, 0xca, 0x7a, 0x7b,
	0x53, 0x7b, 0xa9, 0x1f, 0x2d, 0xbc, 0x12, 0x26, 0x49, 0xd8, 0xc7, 0x36, 0x4d, 0xa3, 0x36, 0x8d,
	0xe3, 0x44, 0x50, 0x11, 0x25, 0x31, 0xd7, 0xda, 0xdb, 0xf2, 0xc7, 0xdf, 0x0b, 0x31, 0xde, 0xe3,
	0x43, 0x1a, 0x86, 0xc8, 0xda, 0x49, 0x2a, 0x2d, 0xce, 0x5a, 0xbb, 0xb7, 0x01, 0x0e, 0x51, 0x78,
	0xf8, 0xd3, 0x09, 0x72, 0x41, 0xae, 0x02, 0x44, 0x41, 0x86, 0xf5, 0x24, 0x42, 0xd6, 0xb4, 0x1c,
	0x6b, 0x67, 0xc5, 0x33, 0x24, 0xee, 0x17, 0xd0, 0x28, 0xac, 0x3d, 0xec, 0xab, 0x50, 0xbd, 0x28,
	0x25, 0xf7, 0x60, 0x9b, 0x19, 0xe7, 0xee, 0x99, 0x30, 0x0d, 0x53, 0xfd, 0xa8, 0x08, 0xf9, 0x29,
	0x6c, 0x1c, 0xa2, 0xf8, 0x0a, 0x19, 0x8f, 0x92, 0xf8, 0x9c, 0x3c, 0x48, 0x13, 0x96, 0x4e, 0x95,
	0x47, 0x73, 0x41, 0x2a, 0xf3, 0xa3, 0xfb, 0xd2, 0x82, 0xcd, 0x0f, 0x65, 0xb1, 0x1e, 0x30, 0xa4,
	0x02, 0x8f, 0xe8, 0xa8, 0x9f, 0xd0, 0x80, 0xdc, 0x85, 0x1a, 0x43, 0x1a, 0x74, 0xa9, 0xef, 0x23,
	0xe7, 0x32, 0x64, 0x6d, 0x7f, 0xab, 0x95, 0x97, 0xbd, 0xe5, 0x21, 0x0d, 0x3a, 0x52, 0xe7, 0x01,
	0x9b, 0xfc, 0x27, 0xef, 0x41, 0x7d, 0xc8, 0x22, 0x81, 0xb9, 0xdf, 0x82, 0xf4, 0xbb, 0x54, 0xf8,
	0x7d, 0x9d, 0x69, 0xb5, 0x63, 0x6d, 0x58, 0x1c, 0xc8, 0x0d, 0x58, 0x0c, 0xa8, 0xa0, 0xcd, 0x8a,
	0xf4, 0x20, 0x2d, 0xdd, 0x41, 0xc5, 0xed, 0x21, 0x15, 0xd4, 0x93, 0x7a, 0xf7, 0xf7, 0x09, 0xe1,
	0x2f, 0xd3, 0xc0, 0x20, 0xfc, 0xba, 0x12, 0x4c, 0x25, 0xb4, 0xf0, 0x86, 0x09, 0x55, 0xfe, 0x75,
	0x42, 0x8b, 0xaf, 0x49, 0x88, 0xc1, 0x9a, 0x92, 0x79, 0xc8, 0xd3, 0x24, 0xe6, 0x48, 0xee, 0x40,
	0xb5, 0x87, 0x34, 0xd0, 0x69, 0xd4, 0xf6, 0x9b, 0x26, 0x4b, 0x65, 0xf3, 0x91, 0xd4, 0x7b, 0xda,
	0x8e, 0xb4, 0x34, 0x96, 0xca, 0xca, 0x9e, 0x81, 0xa5, 0xfd, 0x34, 0xe6, 0x01, 0xd4, 0x4b, 0xd3,
	0x68, 0xc3, 0xb2, 0x2a, 0x95, 0x18, 0xe9, 0xd2, 0x4d, 0xce, 0xa4, 0x01, 0x55, 0xea, 0x8b, 0xe8,
	0x14, 0x65, 0xf4, 0x65, 0x4f, 0x9f, 0xdc, 0xbf, 0x2c, 0x80, 0x02, 0xe0, 0x95, 0x21, 0xfe, 0x0f,
	0xd0, 0xc7, 0x90, 0xf6, 0xbb, 0x31, 0x1d, 0xa0, 0x9e, 0xc0, 0x15, 0x29, 0xf9, 0x8c, 0x0e, 0x90,
	0xec, 0xc1, 0x0a, 0x0d, 0x02, 0x86, 0x9c, 0x63, 0x56, 0xe0, 0xca, 0x4e, 0x6d, 0x7f, 0x3d, 0x4f,
	0xa1, 0xa3, 0x14, 0x5e, 0x61, 0x41, 0x3e, 0x80, 0xf5, 0x94, 0x8e, 0xb2, 0x72, 0x74, 0x03, 0x14,
	0x34, 0xea, 0xf3, 0xe6, 0xa2, 0x74, 0xba, 0x94, 0x3b, 0x1d, 0x29, 0xf5, 0x43, 0xa9, 0xf5, 0xd6,
	0x52, 0xf3, 0xc8, 0xc9, 0x2e, 0x2c, 0xfb, 0x49, 0x2c, 0xa8, 0x2f, 0x78, 0xf3, 0x42, 0x19, 0xed,
	0x81, 0x92, 0x7b, 0x13, 0x03, 0xf7, 0x67, 0x20, 0x67, 0xab, 0x48, 0x6e, 0x41, 0xd5, 0x48, 0x75,
	0x76, 0x77, 0xb5, 0x05, 0xb9, 0x0f, 0xab, 0xe6, 0x55, 0xce, 0x46, 0xaf, 0x22, 0x47, 0x4f, 0xbb,
	0x98, 0x8d, 0xf0, 0xca, 0xa6, 0xee, 0xf7, 0xb0, 0x69, 0xaa, 0xcf, 0x3b, 0xeb, 0x37, 0x61, 0x5d,
	0x50, 0x16, 0xa2, 0xe8, 0x4e, 0x5a, 0xa2, 0x8a, 0xbe, 0xa6, 0xc4, 0x8f, 0xb4, 0xd4, 0x7d, 0x6e,
	0xc1, 0x45, 0x13, 0x40, 0x76, 0xf2, 0x3a, 0xac, 0x25, 0xc3, 0x18, 0x59, 0x77, 0xaa, 0x9f, 0xab,
	0x52, 0x9a, 0xfb, 0x9e, 0x1b, 0x84, 0xec, 0xc2, 0x86, 0xfa, 0x67, 0x2e, 0xb9, 0x8a, 0x34, 0xbd,
	0xa8, 0x14, 0xc6, 0x7a, 0xfb, 0xd5, 0x82, 0xad, 0x52, 0x45, 0xde, 0xfc, 0x52, 0xbc, 0x0f, 0x75,
	0xb3, 0x9a, 0xba, 0xee, 0xcd, 0x59, 0x75, 0x97, 0x0d, 0x2b, 0x59, 0xef, 0xff, 0xb9, 0x04, 0xab,
	0xaa, 0x9b, 0x8f, 0xd5, 0x0b, 0x43, 0x42, 0xa8, 0xaa, 0x1d, 0x49, 0x2e, 0x97, 0xdb, 0x5d, 0xda,
	0x9c, 0x76, 0xa3, 0xac, 0xcc, 0xc9, 0xb9, 0x3b, 0xe3, 0xce, 0xa6, 0xbd, 0xa1, 0x6c, 0xb9, 0x43,
	0x63, 0x47, 0x99, 0xfd, 0xf2, 0xc7, 0xdf, 0x2f, 0x16, 0xea, 0xee, 0x92, 0x7e, 0xb2, 0xee, 0x5b,
	0xb7, 0x88, 0x80, 0xaa, 0xda, 0x6d, 0xd3, 0x40, 0xa5, 0x8d, 0x37, 0x17, 0xe8, 0xae, 0x04, 0x52,
	0xb6, 0xd3, 0x40, 0xff, 0xb3, 0xb7, 0x34, 0x50, 0xfb, 0x69, 0xd1, 0x8d, 0x67, 0x19, 0xea, 0x0f,
	0x50, 0x39, 0x44, 0x41, 0x26, 0xa3, 0x5c, 0x3c, 0x5c, 0x73, 0x91, 0xee, 0x8d, 0x3b, 0x4d, 0x3b,
	0x7b, 0xe1, 0x1c, 0xd1, 0x43, 0xc7, 0x3f, 0x61, 0x0c, 0x63, 0x61, 0xc2, 0x35, 0xc8, 0x4c, 0x38,
	0xf2, 0xd2, 0x82, 0xed, 0x43, 0x14, 0x2a, 0xdc, 0xc1, 0xa8, 0xb4, 0x8b, 0xae, 0x9e, 0x25, 0x60,
	0xea, 0xe7, 0x92, 0xf9, 0x6e, 0xdc, 0x71, 0x6d, 0x27, 0x23, 0xa3, 0xf4, 0xce, 0x13, 0x96, 0x0c,
	0x9c, 0xe3, 0x13, 0x1e, 0xc5, 0xc8, 0xb9, 0x93, 0x52, 0x26, 0x62, 0x64, 0x92, 0xd6, 0x1d, 0xd2,
	0x6a, 0x9b, 0xcd, 0x6e, 0x3f, 0x9d, 0xf3, 0x16, 0x3f, 0xd3, 0xec, 0xc9, 0x0b, 0x0b, 0x2e, 0x3c,
	0xee, 0x51, 0x66, 0xb4, 0x64, 0xc6, 0xc5, 0xb4, 0xaf, 0xcc, 0xbc, 0xd4, 0x39, 0xc3, 0xcf, 0xc7,
	0x9d, 0x9b, 0xf6, 0x75, 0x19, 0x46, 0x16, 0x4c, 0xf3, 0xcc, 0xe7, 0xd8, 0x19, 0x46, 0xa2, 0xe7,
	0x24, 0xa2, 0x87, 0x8c, 0x4b, 0x9a, 0xd7, 0x5c, 0x7b, 0x56, 0xf5, 0xda, 0x3c, 0x8b, 0x90, 0xb5,
	0xec, 0xb9, 0x05, 0x55, 0x0f, 0x4f, 0x93, 0x1f, 0xff, 0x13, 0xad, 0x8f, 0xc7, 0x9d, 0xb7, 0xec,
	0x6b, 0x4c, 0xc6, 0x29, 0xc6, 0xa5, 0xa0, 0x25, 0xe1, 0x24, 0x21, 0xc7, 0xbd, 0x3c, 0x93, 0x90,
	0xf2, 0xcd, 0x18, 0xfd, 0x66, 0xc1, 0xf6, 0x27, 0x11, 0x17, 0x79, 0x6f, 0x8c, 0x65, 0x36, 0x73,
	0xb2, 0x5e, 0xcd, 0xec, 0x9b, 0x71, 0x67, 0xd7, 0x7e, 0xa7, 0x1f, 0x71, 0xe1, 0x94, 0xf6, 0xa2,
	0x73, 0x8c, 0xfd, 0x24, 0x0e, 0xa3, 0x38, 0x74, 0x44, 0x32, 0x35, 0xe1, 0x6f, 0x13, 0x77, 0x0e,
	0x47, 0x23, 0xc0, 0xc1, 0x0d, 0x00, 0x3f, 0x19, 0x68, 0xf0, 0x83, 0xba, 0xbe, 0xe2, 0x47, 0xd9,
	0x27, 0xde, 0x91, 0xf5, 0xed, 0xb2, 0x92, 0xa7, 0xc7, 0xc7, 0x55, 0xf9, 0xd5, 0xf7, 0xee, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xad, 0x69, 0xff, 0x6e, 0x8e, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EntityServiceClient is the client API for EntityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EntityServiceClient interface {
	Create(ctx context.Context, in *EntityCreatePayload, opts ...grpc.CallOption) (*EntityResponse, error)
	Update(ctx context.Context, in *EntityUpdatePayload, opts ...grpc.CallOption) (*EntityResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*EntityResponse, error)
	// Entity Relation Get
	GetEntityByRelationship(ctx context.Context, in *GetRequestRelationship, opts ...grpc.CallOption) (*EntityResponse, error)
	// Entity Relation Share
	Share(ctx context.Context, in *RelationshipPayload, opts ...grpc.CallOption) (*RelationshipResponse, error)
	// Entity Relation Revoke
	Revoke(ctx context.Context, in *RelationshipPayload, opts ...grpc.CallOption) (*RelationshipResponse, error)
	// ListEntityRelationships
	ListEntityRelationships(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*RelationshipResponse, error)
}

type entityServiceClient struct {
	cc *grpc.ClientConn
}

func NewEntityServiceClient(cc *grpc.ClientConn) EntityServiceClient {
	return &entityServiceClient{cc}
}

func (c *entityServiceClient) Create(ctx context.Context, in *EntityCreatePayload, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) Update(ctx context.Context, in *EntityUpdatePayload, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) GetEntityByRelationship(ctx context.Context, in *GetRequestRelationship, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityService/GetEntityByRelationship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) Share(ctx context.Context, in *RelationshipPayload, opts ...grpc.CallOption) (*RelationshipResponse, error) {
	out := new(RelationshipResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityService/Share", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) Revoke(ctx context.Context, in *RelationshipPayload, opts ...grpc.CallOption) (*RelationshipResponse, error) {
	out := new(RelationshipResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityService/Revoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) ListEntityRelationships(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*RelationshipResponse, error) {
	out := new(RelationshipResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityService/ListEntityRelationships", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityServiceServer is the server API for EntityService service.
type EntityServiceServer interface {
	Create(context.Context, *EntityCreatePayload) (*EntityResponse, error)
	Update(context.Context, *EntityUpdatePayload) (*EntityResponse, error)
	Get(context.Context, *GetRequest) (*EntityResponse, error)
	// Entity Relation Get
	GetEntityByRelationship(context.Context, *GetRequestRelationship) (*EntityResponse, error)
	// Entity Relation Share
	Share(context.Context, *RelationshipPayload) (*RelationshipResponse, error)
	// Entity Relation Revoke
	Revoke(context.Context, *RelationshipPayload) (*RelationshipResponse, error)
	// ListEntityRelationships
	ListEntityRelationships(context.Context, *GetRequest) (*RelationshipResponse, error)
}

func RegisterEntityServiceServer(s *grpc.Server, srv EntityServiceServer) {
	s.RegisterService(&_EntityService_serviceDesc, srv)
}

func _EntityService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntityCreatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).Create(ctx, req.(*EntityCreatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntityUpdatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).Update(ctx, req.(*EntityUpdatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_GetEntityByRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestRelationship)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).GetEntityByRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityService/GetEntityByRelationship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).GetEntityByRelationship(ctx, req.(*GetRequestRelationship))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_Share_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationshipPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).Share(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityService/Share",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).Share(ctx, req.(*RelationshipPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationshipPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityService/Revoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).Revoke(ctx, req.(*RelationshipPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityService_ListEntityRelationships_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).ListEntityRelationships(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityService/ListEntityRelationships",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).ListEntityRelationships(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EntityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "entity.EntityService",
	HandlerType: (*EntityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EntityService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EntityService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _EntityService_Get_Handler,
		},
		{
			MethodName: "GetEntityByRelationship",
			Handler:    _EntityService_GetEntityByRelationship_Handler,
		},
		{
			MethodName: "Share",
			Handler:    _EntityService_Share_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _EntityService_Revoke_Handler,
		},
		{
			MethodName: "ListEntityRelationships",
			Handler:    _EntityService_ListEntityRelationships_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entity/service.proto",
}
