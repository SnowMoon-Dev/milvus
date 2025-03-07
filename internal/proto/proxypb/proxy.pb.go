// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proxy.proto

package proxypb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/milvus-io/milvus/internal/proto/commonpb"
	internalpb "github.com/milvus-io/milvus/internal/proto/internalpb"
	milvuspb "github.com/milvus-io/milvus/internal/proto/milvuspb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type InvalidateCollMetaCacheRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	DbName               string            `protobuf:"bytes,2,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	CollectionName       string            `protobuf:"bytes,3,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	CollectionID         int64             `protobuf:"varint,4,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *InvalidateCollMetaCacheRequest) Reset()         { *m = InvalidateCollMetaCacheRequest{} }
func (m *InvalidateCollMetaCacheRequest) String() string { return proto.CompactTextString(m) }
func (*InvalidateCollMetaCacheRequest) ProtoMessage()    {}
func (*InvalidateCollMetaCacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{0}
}

func (m *InvalidateCollMetaCacheRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvalidateCollMetaCacheRequest.Unmarshal(m, b)
}
func (m *InvalidateCollMetaCacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvalidateCollMetaCacheRequest.Marshal(b, m, deterministic)
}
func (m *InvalidateCollMetaCacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvalidateCollMetaCacheRequest.Merge(m, src)
}
func (m *InvalidateCollMetaCacheRequest) XXX_Size() int {
	return xxx_messageInfo_InvalidateCollMetaCacheRequest.Size(m)
}
func (m *InvalidateCollMetaCacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvalidateCollMetaCacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvalidateCollMetaCacheRequest proto.InternalMessageInfo

func (m *InvalidateCollMetaCacheRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *InvalidateCollMetaCacheRequest) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *InvalidateCollMetaCacheRequest) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *InvalidateCollMetaCacheRequest) GetCollectionID() int64 {
	if m != nil {
		return m.CollectionID
	}
	return 0
}

type ReleaseDQLMessageStreamRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	DbID                 int64             `protobuf:"varint,2,opt,name=dbID,proto3" json:"dbID,omitempty"`
	CollectionID         int64             `protobuf:"varint,3,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReleaseDQLMessageStreamRequest) Reset()         { *m = ReleaseDQLMessageStreamRequest{} }
func (m *ReleaseDQLMessageStreamRequest) String() string { return proto.CompactTextString(m) }
func (*ReleaseDQLMessageStreamRequest) ProtoMessage()    {}
func (*ReleaseDQLMessageStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{1}
}

func (m *ReleaseDQLMessageStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseDQLMessageStreamRequest.Unmarshal(m, b)
}
func (m *ReleaseDQLMessageStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseDQLMessageStreamRequest.Marshal(b, m, deterministic)
}
func (m *ReleaseDQLMessageStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseDQLMessageStreamRequest.Merge(m, src)
}
func (m *ReleaseDQLMessageStreamRequest) XXX_Size() int {
	return xxx_messageInfo_ReleaseDQLMessageStreamRequest.Size(m)
}
func (m *ReleaseDQLMessageStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseDQLMessageStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseDQLMessageStreamRequest proto.InternalMessageInfo

func (m *ReleaseDQLMessageStreamRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *ReleaseDQLMessageStreamRequest) GetDbID() int64 {
	if m != nil {
		return m.DbID
	}
	return 0
}

func (m *ReleaseDQLMessageStreamRequest) GetCollectionID() int64 {
	if m != nil {
		return m.CollectionID
	}
	return 0
}

type InvalidateCredCacheRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Username             string            `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *InvalidateCredCacheRequest) Reset()         { *m = InvalidateCredCacheRequest{} }
func (m *InvalidateCredCacheRequest) String() string { return proto.CompactTextString(m) }
func (*InvalidateCredCacheRequest) ProtoMessage()    {}
func (*InvalidateCredCacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{2}
}

func (m *InvalidateCredCacheRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvalidateCredCacheRequest.Unmarshal(m, b)
}
func (m *InvalidateCredCacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvalidateCredCacheRequest.Marshal(b, m, deterministic)
}
func (m *InvalidateCredCacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvalidateCredCacheRequest.Merge(m, src)
}
func (m *InvalidateCredCacheRequest) XXX_Size() int {
	return xxx_messageInfo_InvalidateCredCacheRequest.Size(m)
}
func (m *InvalidateCredCacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvalidateCredCacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvalidateCredCacheRequest proto.InternalMessageInfo

func (m *InvalidateCredCacheRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *InvalidateCredCacheRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UpdateCredCacheRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Username             string            `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string            `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateCredCacheRequest) Reset()         { *m = UpdateCredCacheRequest{} }
func (m *UpdateCredCacheRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCredCacheRequest) ProtoMessage()    {}
func (*UpdateCredCacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{3}
}

func (m *UpdateCredCacheRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCredCacheRequest.Unmarshal(m, b)
}
func (m *UpdateCredCacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCredCacheRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCredCacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCredCacheRequest.Merge(m, src)
}
func (m *UpdateCredCacheRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCredCacheRequest.Size(m)
}
func (m *UpdateCredCacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCredCacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCredCacheRequest proto.InternalMessageInfo

func (m *UpdateCredCacheRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *UpdateCredCacheRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UpdateCredCacheRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RefreshPolicyInfoCacheRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	OpType               int32             `protobuf:"varint,2,opt,name=opType,proto3" json:"opType,omitempty"`
	OpKey                string            `protobuf:"bytes,3,opt,name=opKey,proto3" json:"opKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RefreshPolicyInfoCacheRequest) Reset()         { *m = RefreshPolicyInfoCacheRequest{} }
func (m *RefreshPolicyInfoCacheRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshPolicyInfoCacheRequest) ProtoMessage()    {}
func (*RefreshPolicyInfoCacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{4}
}

func (m *RefreshPolicyInfoCacheRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshPolicyInfoCacheRequest.Unmarshal(m, b)
}
func (m *RefreshPolicyInfoCacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshPolicyInfoCacheRequest.Marshal(b, m, deterministic)
}
func (m *RefreshPolicyInfoCacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshPolicyInfoCacheRequest.Merge(m, src)
}
func (m *RefreshPolicyInfoCacheRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshPolicyInfoCacheRequest.Size(m)
}
func (m *RefreshPolicyInfoCacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshPolicyInfoCacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshPolicyInfoCacheRequest proto.InternalMessageInfo

func (m *RefreshPolicyInfoCacheRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *RefreshPolicyInfoCacheRequest) GetOpType() int32 {
	if m != nil {
		return m.OpType
	}
	return 0
}

func (m *RefreshPolicyInfoCacheRequest) GetOpKey() string {
	if m != nil {
		return m.OpKey
	}
	return ""
}

func init() {
	proto.RegisterType((*InvalidateCollMetaCacheRequest)(nil), "milvus.proto.proxy.InvalidateCollMetaCacheRequest")
	proto.RegisterType((*ReleaseDQLMessageStreamRequest)(nil), "milvus.proto.proxy.ReleaseDQLMessageStreamRequest")
	proto.RegisterType((*InvalidateCredCacheRequest)(nil), "milvus.proto.proxy.InvalidateCredCacheRequest")
	proto.RegisterType((*UpdateCredCacheRequest)(nil), "milvus.proto.proxy.UpdateCredCacheRequest")
	proto.RegisterType((*RefreshPolicyInfoCacheRequest)(nil), "milvus.proto.proxy.RefreshPolicyInfoCacheRequest")
}

func init() { proto.RegisterFile("proxy.proto", fileDescriptor_700b50b08ed8dbaf) }

var fileDescriptor_700b50b08ed8dbaf = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x5f, 0x4f, 0x13, 0x4f,
	0x14, 0x65, 0x7f, 0x05, 0x7e, 0x72, 0x69, 0xd0, 0x8c, 0x08, 0x58, 0x85, 0x90, 0xd5, 0x28, 0x21,
	0xb1, 0xc5, 0xea, 0x27, 0xa0, 0x4d, 0x48, 0xa3, 0x18, 0xdc, 0x4a, 0x4c, 0xf4, 0xc1, 0xcc, 0xee,
	0x5e, 0xda, 0x21, 0xb3, 0x33, 0xcb, 0xcc, 0x14, 0xed, 0x93, 0x89, 0x8f, 0xc6, 0x0f, 0xe4, 0xc7,
	0x33, 0xfb, 0xa7, 0x4b, 0xb7, 0xdd, 0x65, 0x55, 0xe2, 0xdb, 0x9c, 0x99, 0x73, 0xe7, 0x9c, 0x3b,
	0x77, 0xee, 0x85, 0xd5, 0x50, 0xc9, 0x2f, 0xe3, 0x66, 0xa8, 0xa4, 0x91, 0x84, 0x04, 0x8c, 0x5f,
	0x8e, 0x74, 0x82, 0x9a, 0xf1, 0x49, 0xa3, 0xee, 0xc9, 0x20, 0x90, 0x22, 0xd9, 0x6b, 0xac, 0x31,
	0x61, 0x50, 0x09, 0xca, 0x53, 0x5c, 0x9f, 0x8e, 0xb0, 0x7f, 0x5a, 0xb0, 0xd3, 0x13, 0x97, 0x94,
	0x33, 0x9f, 0x1a, 0xec, 0x48, 0xce, 0x8f, 0xd1, 0xd0, 0x0e, 0xf5, 0x86, 0xe8, 0xe0, 0xc5, 0x08,
	0xb5, 0x21, 0x07, 0xb0, 0xe8, 0x52, 0x8d, 0x5b, 0xd6, 0xae, 0xb5, 0xb7, 0xda, 0x7e, 0xd8, 0xcc,
	0x29, 0xa6, 0x52, 0xc7, 0x7a, 0x70, 0x48, 0x35, 0x3a, 0x31, 0x93, 0x6c, 0xc2, 0xff, 0xbe, 0xfb,
	0x49, 0xd0, 0x00, 0xb7, 0xfe, 0xdb, 0xb5, 0xf6, 0x56, 0x9c, 0x65, 0xdf, 0x7d, 0x43, 0x03, 0x24,
	0x4f, 0xe1, 0xb6, 0x27, 0x39, 0x47, 0xcf, 0x30, 0x29, 0x12, 0x42, 0x2d, 0x26, 0xac, 0x5d, 0x6d,
	0xc7, 0x44, 0x1b, 0xea, 0x57, 0x3b, 0xbd, 0xee, 0xd6, 0xe2, 0xae, 0xb5, 0x57, 0x73, 0x72, 0x7b,
	0xf6, 0x77, 0x0b, 0x76, 0x1c, 0xe4, 0x48, 0x35, 0x76, 0xdf, 0xbe, 0x3e, 0x46, 0xad, 0xe9, 0x00,
	0xfb, 0x46, 0x21, 0x0d, 0xfe, 0xde, 0x3a, 0x81, 0x45, 0xdf, 0xed, 0x75, 0x63, 0xdf, 0x35, 0x27,
	0x5e, 0xcf, 0x99, 0xa9, 0x15, 0x98, 0x39, 0x87, 0xc6, 0xd4, 0x33, 0x2a, 0xf4, 0x6f, 0xf8, 0x84,
	0x0d, 0xb8, 0x35, 0xd2, 0x51, 0xd9, 0xb2, 0x37, 0xcc, 0xb0, 0xfd, 0xcd, 0x82, 0x8d, 0xd3, 0xf0,
	0xdf, 0x0b, 0x45, 0x67, 0x21, 0xd5, 0xfa, 0xb3, 0x54, 0x7e, 0x5a, 0xa7, 0x0c, 0xdb, 0x5f, 0x61,
	0xdb, 0xc1, 0x33, 0x85, 0x7a, 0x78, 0x22, 0x39, 0xf3, 0xc6, 0x3d, 0x71, 0x26, 0x6f, 0x68, 0x65,
	0x03, 0x96, 0x65, 0xf8, 0x6e, 0x1c, 0x26, 0x46, 0x96, 0x9c, 0x14, 0x91, 0x75, 0x58, 0x92, 0xe1,
	0x2b, 0x1c, 0xa7, 0x1e, 0x12, 0xd0, 0xfe, 0xb1, 0x02, 0x4b, 0x27, 0xd1, 0x7f, 0x27, 0x21, 0x90,
	0x23, 0x34, 0x1d, 0x19, 0x84, 0x52, 0xa0, 0x30, 0x7d, 0x43, 0x0d, 0x6a, 0x72, 0x90, 0x57, 0xcc,
	0xba, 0x60, 0x9e, 0x9a, 0x3a, 0x6e, 0x3c, 0x29, 0x89, 0x98, 0xa1, 0xdb, 0x0b, 0xe4, 0x02, 0xd6,
	0x8f, 0x30, 0x86, 0x4c, 0x1b, 0xe6, 0xe9, 0xce, 0x90, 0x0a, 0x81, 0x9c, 0xb4, 0xcb, 0x35, 0xe7,
	0xc8, 0x13, 0xd5, 0x47, 0xf9, 0x98, 0x14, 0xf4, 0x8d, 0x62, 0x62, 0xe0, 0xa0, 0x0e, 0xa5, 0xd0,
	0x68, 0x2f, 0x10, 0x05, 0xdb, 0xf9, 0x3e, 0x4d, 0xbe, 0x5e, 0xd6, 0xad, 0xb3, 0xda, 0xc9, 0x90,
	0xb8, 0xbe, 0xb5, 0x1b, 0x0f, 0x0a, 0xab, 0x12, 0x59, 0x1d, 0x45, 0x69, 0x52, 0xa8, 0x1f, 0xa1,
	0xe9, 0xfa, 0x93, 0xf4, 0xf6, 0xcb, 0xd3, 0xcb, 0x48, 0x7f, 0x98, 0x16, 0x87, 0xcd, 0x92, 0x1e,
	0x2e, 0x4e, 0xe8, 0xfa, 0x86, 0xaf, 0x4a, 0xe8, 0x3d, 0xdc, 0xe9, 0xa3, 0xf0, 0xfb, 0x48, 0x95,
	0x37, 0x74, 0x50, 0x8f, 0xb8, 0x21, 0x8f, 0x4b, 0x92, 0x9a, 0x26, 0xe9, 0xaa, 0x8b, 0x3f, 0x02,
	0x89, 0x2e, 0x76, 0xd0, 0x28, 0x86, 0x97, 0x98, 0x5e, 0x5d, 0xf6, 0xa1, 0xf2, 0xb4, 0xca, 0xcb,
	0xcf, 0xe1, 0x7e, 0x7e, 0xb6, 0xa0, 0x30, 0x8c, 0xf2, 0xa4, 0xec, 0xcd, 0x8a, 0xb2, 0xcf, 0x4c,
	0x88, 0x2a, 0x2d, 0x17, 0xee, 0x5d, 0x8d, 0x96, 0x69, 0x9d, 0xfd, 0x22, 0x9d, 0xe2, 0x29, 0x54,
	0xa5, 0x31, 0x80, 0xbb, 0x1d, 0x8e, 0x54, 0x45, 0x71, 0xa7, 0x1a, 0x95, 0x4e, 0x14, 0x9e, 0x97,
	0xb5, 0xdf, 0x3c, 0xf7, 0x37, 0x85, 0xce, 0x61, 0xa3, 0x78, 0x46, 0xcd, 0x6a, 0x4d, 0xfe, 0xd6,
	0x35, 0xf3, 0xac, 0x42, 0xeb, 0xf0, 0xe5, 0x87, 0xf6, 0x80, 0x99, 0xe1, 0xc8, 0x8d, 0x4e, 0x5a,
	0x09, 0xf5, 0x19, 0x93, 0xe9, 0xaa, 0x35, 0xc9, 0xa6, 0x15, 0x47, 0xb7, 0x62, 0xc1, 0xd0, 0x75,
	0x97, 0x63, 0xf8, 0xe2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0xd6, 0xdb, 0x10, 0xd4, 0x07,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProxyClient is the client API for Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxyClient interface {
	GetComponentStates(ctx context.Context, in *internalpb.GetComponentStatesRequest, opts ...grpc.CallOption) (*internalpb.ComponentStates, error)
	GetStatisticsChannel(ctx context.Context, in *internalpb.GetStatisticsChannelRequest, opts ...grpc.CallOption) (*milvuspb.StringResponse, error)
	InvalidateCollectionMetaCache(ctx context.Context, in *InvalidateCollMetaCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	GetDdChannel(ctx context.Context, in *internalpb.GetDdChannelRequest, opts ...grpc.CallOption) (*milvuspb.StringResponse, error)
	ReleaseDQLMessageStream(ctx context.Context, in *ReleaseDQLMessageStreamRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	SendSearchResult(ctx context.Context, in *internalpb.SearchResults, opts ...grpc.CallOption) (*commonpb.Status, error)
	SendRetrieveResult(ctx context.Context, in *internalpb.RetrieveResults, opts ...grpc.CallOption) (*commonpb.Status, error)
	InvalidateCredentialCache(ctx context.Context, in *InvalidateCredCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	UpdateCredentialCache(ctx context.Context, in *UpdateCredCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	ClearCredUsersCache(ctx context.Context, in *internalpb.ClearCredUsersCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
	RefreshPolicyInfoCache(ctx context.Context, in *RefreshPolicyInfoCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error)
}

type proxyClient struct {
	cc *grpc.ClientConn
}

func NewProxyClient(cc *grpc.ClientConn) ProxyClient {
	return &proxyClient{cc}
}

func (c *proxyClient) GetComponentStates(ctx context.Context, in *internalpb.GetComponentStatesRequest, opts ...grpc.CallOption) (*internalpb.ComponentStates, error) {
	out := new(internalpb.ComponentStates)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/GetComponentStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) GetStatisticsChannel(ctx context.Context, in *internalpb.GetStatisticsChannelRequest, opts ...grpc.CallOption) (*milvuspb.StringResponse, error) {
	out := new(milvuspb.StringResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/GetStatisticsChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) InvalidateCollectionMetaCache(ctx context.Context, in *InvalidateCollMetaCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/InvalidateCollectionMetaCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) GetDdChannel(ctx context.Context, in *internalpb.GetDdChannelRequest, opts ...grpc.CallOption) (*milvuspb.StringResponse, error) {
	out := new(milvuspb.StringResponse)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/GetDdChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) ReleaseDQLMessageStream(ctx context.Context, in *ReleaseDQLMessageStreamRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/ReleaseDQLMessageStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) SendSearchResult(ctx context.Context, in *internalpb.SearchResults, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/SendSearchResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) SendRetrieveResult(ctx context.Context, in *internalpb.RetrieveResults, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/SendRetrieveResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) InvalidateCredentialCache(ctx context.Context, in *InvalidateCredCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/InvalidateCredentialCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) UpdateCredentialCache(ctx context.Context, in *UpdateCredCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/UpdateCredentialCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) ClearCredUsersCache(ctx context.Context, in *internalpb.ClearCredUsersCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/ClearCredUsersCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) RefreshPolicyInfoCache(ctx context.Context, in *RefreshPolicyInfoCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	out := new(commonpb.Status)
	err := c.cc.Invoke(ctx, "/milvus.proto.proxy.Proxy/RefreshPolicyInfoCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyServer is the server API for Proxy service.
type ProxyServer interface {
	GetComponentStates(context.Context, *internalpb.GetComponentStatesRequest) (*internalpb.ComponentStates, error)
	GetStatisticsChannel(context.Context, *internalpb.GetStatisticsChannelRequest) (*milvuspb.StringResponse, error)
	InvalidateCollectionMetaCache(context.Context, *InvalidateCollMetaCacheRequest) (*commonpb.Status, error)
	GetDdChannel(context.Context, *internalpb.GetDdChannelRequest) (*milvuspb.StringResponse, error)
	ReleaseDQLMessageStream(context.Context, *ReleaseDQLMessageStreamRequest) (*commonpb.Status, error)
	SendSearchResult(context.Context, *internalpb.SearchResults) (*commonpb.Status, error)
	SendRetrieveResult(context.Context, *internalpb.RetrieveResults) (*commonpb.Status, error)
	InvalidateCredentialCache(context.Context, *InvalidateCredCacheRequest) (*commonpb.Status, error)
	UpdateCredentialCache(context.Context, *UpdateCredCacheRequest) (*commonpb.Status, error)
	ClearCredUsersCache(context.Context, *internalpb.ClearCredUsersCacheRequest) (*commonpb.Status, error)
	RefreshPolicyInfoCache(context.Context, *RefreshPolicyInfoCacheRequest) (*commonpb.Status, error)
}

// UnimplementedProxyServer can be embedded to have forward compatible implementations.
type UnimplementedProxyServer struct {
}

func (*UnimplementedProxyServer) GetComponentStates(ctx context.Context, req *internalpb.GetComponentStatesRequest) (*internalpb.ComponentStates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComponentStates not implemented")
}
func (*UnimplementedProxyServer) GetStatisticsChannel(ctx context.Context, req *internalpb.GetStatisticsChannelRequest) (*milvuspb.StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatisticsChannel not implemented")
}
func (*UnimplementedProxyServer) InvalidateCollectionMetaCache(ctx context.Context, req *InvalidateCollMetaCacheRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateCollectionMetaCache not implemented")
}
func (*UnimplementedProxyServer) GetDdChannel(ctx context.Context, req *internalpb.GetDdChannelRequest) (*milvuspb.StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDdChannel not implemented")
}
func (*UnimplementedProxyServer) ReleaseDQLMessageStream(ctx context.Context, req *ReleaseDQLMessageStreamRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseDQLMessageStream not implemented")
}
func (*UnimplementedProxyServer) SendSearchResult(ctx context.Context, req *internalpb.SearchResults) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSearchResult not implemented")
}
func (*UnimplementedProxyServer) SendRetrieveResult(ctx context.Context, req *internalpb.RetrieveResults) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRetrieveResult not implemented")
}
func (*UnimplementedProxyServer) InvalidateCredentialCache(ctx context.Context, req *InvalidateCredCacheRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateCredentialCache not implemented")
}
func (*UnimplementedProxyServer) UpdateCredentialCache(ctx context.Context, req *UpdateCredCacheRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCredentialCache not implemented")
}
func (*UnimplementedProxyServer) ClearCredUsersCache(ctx context.Context, req *internalpb.ClearCredUsersCacheRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearCredUsersCache not implemented")
}
func (*UnimplementedProxyServer) RefreshPolicyInfoCache(ctx context.Context, req *RefreshPolicyInfoCacheRequest) (*commonpb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshPolicyInfoCache not implemented")
}

func RegisterProxyServer(s *grpc.Server, srv ProxyServer) {
	s.RegisterService(&_Proxy_serviceDesc, srv)
}

func _Proxy_GetComponentStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.GetComponentStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetComponentStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/GetComponentStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetComponentStates(ctx, req.(*internalpb.GetComponentStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_GetStatisticsChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.GetStatisticsChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetStatisticsChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/GetStatisticsChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetStatisticsChannel(ctx, req.(*internalpb.GetStatisticsChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_InvalidateCollectionMetaCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateCollMetaCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).InvalidateCollectionMetaCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/InvalidateCollectionMetaCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).InvalidateCollectionMetaCache(ctx, req.(*InvalidateCollMetaCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_GetDdChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.GetDdChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetDdChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/GetDdChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetDdChannel(ctx, req.(*internalpb.GetDdChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_ReleaseDQLMessageStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseDQLMessageStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).ReleaseDQLMessageStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/ReleaseDQLMessageStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).ReleaseDQLMessageStream(ctx, req.(*ReleaseDQLMessageStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_SendSearchResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.SearchResults)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).SendSearchResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/SendSearchResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).SendSearchResult(ctx, req.(*internalpb.SearchResults))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_SendRetrieveResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.RetrieveResults)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).SendRetrieveResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/SendRetrieveResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).SendRetrieveResult(ctx, req.(*internalpb.RetrieveResults))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_InvalidateCredentialCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateCredCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).InvalidateCredentialCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/InvalidateCredentialCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).InvalidateCredentialCache(ctx, req.(*InvalidateCredCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_UpdateCredentialCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCredCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).UpdateCredentialCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/UpdateCredentialCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).UpdateCredentialCache(ctx, req.(*UpdateCredCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_ClearCredUsersCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internalpb.ClearCredUsersCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).ClearCredUsersCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/ClearCredUsersCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).ClearCredUsersCache(ctx, req.(*internalpb.ClearCredUsersCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_RefreshPolicyInfoCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshPolicyInfoCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).RefreshPolicyInfoCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milvus.proto.proxy.Proxy/RefreshPolicyInfoCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).RefreshPolicyInfoCache(ctx, req.(*RefreshPolicyInfoCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Proxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.proxy.Proxy",
	HandlerType: (*ProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetComponentStates",
			Handler:    _Proxy_GetComponentStates_Handler,
		},
		{
			MethodName: "GetStatisticsChannel",
			Handler:    _Proxy_GetStatisticsChannel_Handler,
		},
		{
			MethodName: "InvalidateCollectionMetaCache",
			Handler:    _Proxy_InvalidateCollectionMetaCache_Handler,
		},
		{
			MethodName: "GetDdChannel",
			Handler:    _Proxy_GetDdChannel_Handler,
		},
		{
			MethodName: "ReleaseDQLMessageStream",
			Handler:    _Proxy_ReleaseDQLMessageStream_Handler,
		},
		{
			MethodName: "SendSearchResult",
			Handler:    _Proxy_SendSearchResult_Handler,
		},
		{
			MethodName: "SendRetrieveResult",
			Handler:    _Proxy_SendRetrieveResult_Handler,
		},
		{
			MethodName: "InvalidateCredentialCache",
			Handler:    _Proxy_InvalidateCredentialCache_Handler,
		},
		{
			MethodName: "UpdateCredentialCache",
			Handler:    _Proxy_UpdateCredentialCache_Handler,
		},
		{
			MethodName: "ClearCredUsersCache",
			Handler:    _Proxy_ClearCredUsersCache_Handler,
		},
		{
			MethodName: "RefreshPolicyInfoCache",
			Handler:    _Proxy_RefreshPolicyInfoCache_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proxy.proto",
}
