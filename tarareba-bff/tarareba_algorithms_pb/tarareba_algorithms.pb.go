// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: tarareba_algorithms.proto

package tarareba_algorithms_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ActualHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. rated かどうか
	IsRated bool `protobuf:"varint,1,opt,name=is_rated,json=isRated,proto3" json:"is_rated,omitempty"`
	// Required. パフォーマンス
	Performance int32 `protobuf:"varint,5,opt,name=performance,proto3" json:"performance,omitempty"`
	// Required. 内部パフォーマンス
	InnerPerformance int32 `protobuf:"varint,6,opt,name=inner_performance,json=innerPerformance,proto3" json:"inner_performance,omitempty"`
}

func (x *ActualHistory) Reset() {
	*x = ActualHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarareba_algorithms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActualHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActualHistory) ProtoMessage() {}

func (x *ActualHistory) ProtoReflect() protoreflect.Message {
	mi := &file_tarareba_algorithms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActualHistory.ProtoReflect.Descriptor instead.
func (*ActualHistory) Descriptor() ([]byte, []int) {
	return file_tarareba_algorithms_proto_rawDescGZIP(), []int{0}
}

func (x *ActualHistory) GetIsRated() bool {
	if x != nil {
		return x.IsRated
	}
	return false
}

func (x *ActualHistory) GetPerformance() int32 {
	if x != nil {
		return x.Performance
	}
	return 0
}

func (x *ActualHistory) GetInnerPerformance() int32 {
	if x != nil {
		return x.InnerPerformance
	}
	return 0
}

type OptimalHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. 参加前のレート
	OldRating int32 `protobuf:"varint,3,opt,name=old_rating,json=oldRating,proto3" json:"old_rating,omitempty"`
	// Required. 参加後のレート
	NewRating int32 `protobuf:"varint,4,opt,name=new_rating,json=newRating,proto3" json:"new_rating,omitempty"`
	// Required. 空想上で、参加するかどうか
	IsParticipated bool `protobuf:"varint,7,opt,name=isParticipated,proto3" json:"isParticipated,omitempty"`
}

func (x *OptimalHistory) Reset() {
	*x = OptimalHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarareba_algorithms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptimalHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptimalHistory) ProtoMessage() {}

func (x *OptimalHistory) ProtoReflect() protoreflect.Message {
	mi := &file_tarareba_algorithms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptimalHistory.ProtoReflect.Descriptor instead.
func (*OptimalHistory) Descriptor() ([]byte, []int) {
	return file_tarareba_algorithms_proto_rawDescGZIP(), []int{1}
}

func (x *OptimalHistory) GetOldRating() int32 {
	if x != nil {
		return x.OldRating
	}
	return 0
}

func (x *OptimalHistory) GetNewRating() int32 {
	if x != nil {
		return x.NewRating
	}
	return 0
}

func (x *OptimalHistory) GetIsParticipated() bool {
	if x != nil {
		return x.IsParticipated
	}
	return false
}

type GetOptimalHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. 実際のコンテスト履歴のリスト
	ActualHistory []*ActualHistory `protobuf:"bytes,1,rep,name=actual_history,json=actualHistory,proto3" json:"actual_history,omitempty"`
}

func (x *GetOptimalHistoryRequest) Reset() {
	*x = GetOptimalHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarareba_algorithms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOptimalHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOptimalHistoryRequest) ProtoMessage() {}

func (x *GetOptimalHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tarareba_algorithms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOptimalHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetOptimalHistoryRequest) Descriptor() ([]byte, []int) {
	return file_tarareba_algorithms_proto_rawDescGZIP(), []int{2}
}

func (x *GetOptimalHistoryRequest) GetActualHistory() []*ActualHistory {
	if x != nil {
		return x.ActualHistory
	}
	return nil
}

type GetOptimalHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. 最適化後の出場コンテストのリスト
	OptimalHistory []*OptimalHistory `protobuf:"bytes,1,rep,name=optimal_history,json=optimalHistory,proto3" json:"optimal_history,omitempty"`
}

func (x *GetOptimalHistoryResponse) Reset() {
	*x = GetOptimalHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarareba_algorithms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOptimalHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOptimalHistoryResponse) ProtoMessage() {}

func (x *GetOptimalHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tarareba_algorithms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOptimalHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetOptimalHistoryResponse) Descriptor() ([]byte, []int) {
	return file_tarareba_algorithms_proto_rawDescGZIP(), []int{3}
}

func (x *GetOptimalHistoryResponse) GetOptimalHistory() []*OptimalHistory {
	if x != nil {
		return x.OptimalHistory
	}
	return nil
}

var File_tarareba_algorithms_proto protoreflect.FileDescriptor

var file_tarareba_algorithms_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x61, 0x72, 0x61, 0x72, 0x65, 0x62, 0x61, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x61, 0x72,
	0x61, 0x72, 0x65, 0x62, 0x61, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73,
	0x22, 0x79, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x52, 0x61, 0x74, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2b,
	0x0a, 0x11, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x76, 0x0a, 0x0e, 0x4f,
	0x70, 0x74, 0x69, 0x6d, 0x61, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x6f, 0x6c, 0x64, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x6f, 0x6c, 0x64, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a,
	0x6e, 0x65, 0x77, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x6e, 0x65, 0x77, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x69,
	0x73, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x74, 0x65, 0x64, 0x22, 0x65, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x61,
	0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x49, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x61, 0x72, 0x61, 0x72, 0x65,
	0x62, 0x61, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x2e, 0x41, 0x63,
	0x74, 0x75, 0x61, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0d, 0x61, 0x63, 0x74,
	0x75, 0x61, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x69, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x61, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6d,
	0x61, 0x6c, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x74, 0x61, 0x72, 0x61, 0x72, 0x65, 0x62, 0x61, 0x5f, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x61, 0x6c, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x61, 0x6c, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x32, 0x87, 0x01, 0x0a, 0x0f, 0x54, 0x61, 0x72, 0x61, 0x72, 0x65,
	0x62, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6d, 0x61, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2d,
	0x2e, 0x74, 0x61, 0x72, 0x61, 0x72, 0x65, 0x62, 0x61, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x61, 0x6c, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x74, 0x61, 0x72, 0x61, 0x72, 0x65, 0x62, 0x61, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x61, 0x6c, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x1b, 0x5a, 0x19, 0x2e, 0x2e, 0x2f, 0x74, 0x61, 0x72, 0x61, 0x72, 0x65, 0x62, 0x61, 0x5f, 0x61,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tarareba_algorithms_proto_rawDescOnce sync.Once
	file_tarareba_algorithms_proto_rawDescData = file_tarareba_algorithms_proto_rawDesc
)

func file_tarareba_algorithms_proto_rawDescGZIP() []byte {
	file_tarareba_algorithms_proto_rawDescOnce.Do(func() {
		file_tarareba_algorithms_proto_rawDescData = protoimpl.X.CompressGZIP(file_tarareba_algorithms_proto_rawDescData)
	})
	return file_tarareba_algorithms_proto_rawDescData
}

var file_tarareba_algorithms_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tarareba_algorithms_proto_goTypes = []interface{}{
	(*ActualHistory)(nil),             // 0: tarareba_algorithms.ActualHistory
	(*OptimalHistory)(nil),            // 1: tarareba_algorithms.OptimalHistory
	(*GetOptimalHistoryRequest)(nil),  // 2: tarareba_algorithms.GetOptimalHistoryRequest
	(*GetOptimalHistoryResponse)(nil), // 3: tarareba_algorithms.GetOptimalHistoryResponse
}
var file_tarareba_algorithms_proto_depIdxs = []int32{
	0, // 0: tarareba_algorithms.GetOptimalHistoryRequest.actual_history:type_name -> tarareba_algorithms.ActualHistory
	1, // 1: tarareba_algorithms.GetOptimalHistoryResponse.optimal_history:type_name -> tarareba_algorithms.OptimalHistory
	2, // 2: tarareba_algorithms.TararebaService.GetOptimalHistory:input_type -> tarareba_algorithms.GetOptimalHistoryRequest
	3, // 3: tarareba_algorithms.TararebaService.GetOptimalHistory:output_type -> tarareba_algorithms.GetOptimalHistoryResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tarareba_algorithms_proto_init() }
func file_tarareba_algorithms_proto_init() {
	if File_tarareba_algorithms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tarareba_algorithms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActualHistory); i {
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
		file_tarareba_algorithms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptimalHistory); i {
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
		file_tarareba_algorithms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOptimalHistoryRequest); i {
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
		file_tarareba_algorithms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOptimalHistoryResponse); i {
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
			RawDescriptor: file_tarareba_algorithms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tarareba_algorithms_proto_goTypes,
		DependencyIndexes: file_tarareba_algorithms_proto_depIdxs,
		MessageInfos:      file_tarareba_algorithms_proto_msgTypes,
	}.Build()
	File_tarareba_algorithms_proto = out.File
	file_tarareba_algorithms_proto_rawDesc = nil
	file_tarareba_algorithms_proto_goTypes = nil
	file_tarareba_algorithms_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TararebaServiceClient is the client API for TararebaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TararebaServiceClient interface {
	// コンテスト参加履歴を返す
	GetOptimalHistory(ctx context.Context, in *GetOptimalHistoryRequest, opts ...grpc.CallOption) (*GetOptimalHistoryResponse, error)
}

type tararebaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTararebaServiceClient(cc grpc.ClientConnInterface) TararebaServiceClient {
	return &tararebaServiceClient{cc}
}

func (c *tararebaServiceClient) GetOptimalHistory(ctx context.Context, in *GetOptimalHistoryRequest, opts ...grpc.CallOption) (*GetOptimalHistoryResponse, error) {
	out := new(GetOptimalHistoryResponse)
	err := c.cc.Invoke(ctx, "/tarareba_algorithms.TararebaService/GetOptimalHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TararebaServiceServer is the server API for TararebaService service.
type TararebaServiceServer interface {
	// コンテスト参加履歴を返す
	GetOptimalHistory(context.Context, *GetOptimalHistoryRequest) (*GetOptimalHistoryResponse, error)
}

// UnimplementedTararebaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTararebaServiceServer struct {
}

func (*UnimplementedTararebaServiceServer) GetOptimalHistory(context.Context, *GetOptimalHistoryRequest) (*GetOptimalHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOptimalHistory not implemented")
}

func RegisterTararebaServiceServer(s *grpc.Server, srv TararebaServiceServer) {
	s.RegisterService(&_TararebaService_serviceDesc, srv)
}

func _TararebaService_GetOptimalHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOptimalHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TararebaServiceServer).GetOptimalHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarareba_algorithms.TararebaService/GetOptimalHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TararebaServiceServer).GetOptimalHistory(ctx, req.(*GetOptimalHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TararebaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tarareba_algorithms.TararebaService",
	HandlerType: (*TararebaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOptimalHistory",
			Handler:    _TararebaService_GetOptimalHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tarareba_algorithms.proto",
}
