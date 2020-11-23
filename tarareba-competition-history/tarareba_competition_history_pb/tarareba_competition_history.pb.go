// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: tarareba_competition_history.proto

package tarareba_competition_history_pb

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

type CompetitionHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. rated かどうか
	IsRated bool `protobuf:"varint,1,opt,name=is_rated,json=isRated,proto3" json:"is_rated,omitempty"`
	// Required. 順位
	Place uint32 `protobuf:"varint,2,opt,name=place,proto3" json:"place,omitempty"`
	// Required. 参加前のレート
	OldRating uint32 `protobuf:"varint,3,opt,name=old_rating,json=oldRating,proto3" json:"old_rating,omitempty"`
	// Required. 参加後のレート
	NewRating uint32 `protobuf:"varint,4,opt,name=new_rating,json=newRating,proto3" json:"new_rating,omitempty"`
	// Required. パフォーマンス
	Performance uint32 `protobuf:"varint,5,opt,name=performance,proto3" json:"performance,omitempty"`
	// Required. 内部パフォーマンス
	InnerPerformance uint32 `protobuf:"varint,6,opt,name=inner_performance,json=innerPerformance,proto3" json:"inner_performance,omitempty"`
	// Required. コンテストのスクリーンネーム（おそらく、フロント側でリンクをはるのに使う）
	ContestScreenName string `protobuf:"bytes,7,opt,name=contest_screen_name,json=contestScreenName,proto3" json:"contest_screen_name,omitempty"`
	// Required. コンテスト名
	ContestName string `protobuf:"bytes,8,opt,name=contest_name,json=contestName,proto3" json:"contest_name,omitempty"`
	// Required. コンテスト名（英語）
	ContestNameEn string `protobuf:"bytes,9,opt,name=contest_name_en,json=contestNameEn,proto3" json:"contest_name_en,omitempty"`
	// Required. コンテストの終了時刻
	EndTime string `protobuf:"bytes,10,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *CompetitionHistory) Reset() {
	*x = CompetitionHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarareba_competition_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompetitionHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompetitionHistory) ProtoMessage() {}

func (x *CompetitionHistory) ProtoReflect() protoreflect.Message {
	mi := &file_tarareba_competition_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompetitionHistory.ProtoReflect.Descriptor instead.
func (*CompetitionHistory) Descriptor() ([]byte, []int) {
	return file_tarareba_competition_history_proto_rawDescGZIP(), []int{0}
}

func (x *CompetitionHistory) GetIsRated() bool {
	if x != nil {
		return x.IsRated
	}
	return false
}

func (x *CompetitionHistory) GetPlace() uint32 {
	if x != nil {
		return x.Place
	}
	return 0
}

func (x *CompetitionHistory) GetOldRating() uint32 {
	if x != nil {
		return x.OldRating
	}
	return 0
}

func (x *CompetitionHistory) GetNewRating() uint32 {
	if x != nil {
		return x.NewRating
	}
	return 0
}

func (x *CompetitionHistory) GetPerformance() uint32 {
	if x != nil {
		return x.Performance
	}
	return 0
}

func (x *CompetitionHistory) GetInnerPerformance() uint32 {
	if x != nil {
		return x.InnerPerformance
	}
	return 0
}

func (x *CompetitionHistory) GetContestScreenName() string {
	if x != nil {
		return x.ContestScreenName
	}
	return ""
}

func (x *CompetitionHistory) GetContestName() string {
	if x != nil {
		return x.ContestName
	}
	return ""
}

func (x *CompetitionHistory) GetContestNameEn() string {
	if x != nil {
		return x.ContestNameEn
	}
	return ""
}

func (x *CompetitionHistory) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type GetCompetitionHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. AtCoder のユーザー ID。（例：monkukui）
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetCompetitionHistoryRequest) Reset() {
	*x = GetCompetitionHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarareba_competition_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompetitionHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompetitionHistoryRequest) ProtoMessage() {}

func (x *GetCompetitionHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tarareba_competition_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompetitionHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetCompetitionHistoryRequest) Descriptor() ([]byte, []int) {
	return file_tarareba_competition_history_proto_rawDescGZIP(), []int{1}
}

func (x *GetCompetitionHistoryRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetCompetitionHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. 出場したコンテストのリスト
	CompetitionHistory []*CompetitionHistory `protobuf:"bytes,1,rep,name=competition_history,json=competitionHistory,proto3" json:"competition_history,omitempty"`
}

func (x *GetCompetitionHistoryResponse) Reset() {
	*x = GetCompetitionHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarareba_competition_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompetitionHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompetitionHistoryResponse) ProtoMessage() {}

func (x *GetCompetitionHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tarareba_competition_history_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompetitionHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetCompetitionHistoryResponse) Descriptor() ([]byte, []int) {
	return file_tarareba_competition_history_proto_rawDescGZIP(), []int{2}
}

func (x *GetCompetitionHistoryResponse) GetCompetitionHistory() []*CompetitionHistory {
	if x != nil {
		return x.CompetitionHistory
	}
	return nil
}

var File_tarareba_competition_history_proto protoreflect.FileDescriptor

var file_tarareba_competition_history_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x61, 0x72, 0x61, 0x72, 0x65, 0x62, 0x61, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x65,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74, 0x61, 0x72, 0x61, 0x72, 0x65, 0x62, 0x61, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x22, 0xe8, 0x02, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x52,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x6c,
	0x64, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x6f, 0x6c, 0x64, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x77,
	0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6e,
	0x65, 0x77, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x65, 0x72, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x53, 0x63, 0x72,
	0x65, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x45, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x37, 0x0a,
	0x1c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x70,
	0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x61, 0x72, 0x61, 0x72, 0x65, 0x62, 0x61,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x32, 0xa6, 0x01, 0x0a, 0x0f,
	0x54, 0x61, 0x72, 0x61, 0x72, 0x65, 0x62, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x92, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3a, 0x2e, 0x74, 0x61, 0x72, 0x61,
	0x72, 0x65, 0x62, 0x61, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x74, 0x61, 0x72, 0x61, 0x72, 0x65, 0x62, 0x61,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x2e, 0x2e, 0x2f, 0x74, 0x61, 0x72, 0x61, 0x72,
	0x65, 0x62, 0x61, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tarareba_competition_history_proto_rawDescOnce sync.Once
	file_tarareba_competition_history_proto_rawDescData = file_tarareba_competition_history_proto_rawDesc
)

func file_tarareba_competition_history_proto_rawDescGZIP() []byte {
	file_tarareba_competition_history_proto_rawDescOnce.Do(func() {
		file_tarareba_competition_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_tarareba_competition_history_proto_rawDescData)
	})
	return file_tarareba_competition_history_proto_rawDescData
}

var file_tarareba_competition_history_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tarareba_competition_history_proto_goTypes = []interface{}{
	(*CompetitionHistory)(nil),            // 0: tarareba_competition_history.CompetitionHistory
	(*GetCompetitionHistoryRequest)(nil),  // 1: tarareba_competition_history.GetCompetitionHistoryRequest
	(*GetCompetitionHistoryResponse)(nil), // 2: tarareba_competition_history.GetCompetitionHistoryResponse
}
var file_tarareba_competition_history_proto_depIdxs = []int32{
	0, // 0: tarareba_competition_history.GetCompetitionHistoryResponse.competition_history:type_name -> tarareba_competition_history.CompetitionHistory
	1, // 1: tarareba_competition_history.TararebaService.GetCompetitionHistory:input_type -> tarareba_competition_history.GetCompetitionHistoryRequest
	2, // 2: tarareba_competition_history.TararebaService.GetCompetitionHistory:output_type -> tarareba_competition_history.GetCompetitionHistoryResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tarareba_competition_history_proto_init() }
func file_tarareba_competition_history_proto_init() {
	if File_tarareba_competition_history_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tarareba_competition_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompetitionHistory); i {
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
		file_tarareba_competition_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompetitionHistoryRequest); i {
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
		file_tarareba_competition_history_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompetitionHistoryResponse); i {
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
			RawDescriptor: file_tarareba_competition_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tarareba_competition_history_proto_goTypes,
		DependencyIndexes: file_tarareba_competition_history_proto_depIdxs,
		MessageInfos:      file_tarareba_competition_history_proto_msgTypes,
	}.Build()
	File_tarareba_competition_history_proto = out.File
	file_tarareba_competition_history_proto_rawDesc = nil
	file_tarareba_competition_history_proto_goTypes = nil
	file_tarareba_competition_history_proto_depIdxs = nil
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
	GetCompetitionHistory(ctx context.Context, in *GetCompetitionHistoryRequest, opts ...grpc.CallOption) (*GetCompetitionHistoryResponse, error)
}

type tararebaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTararebaServiceClient(cc grpc.ClientConnInterface) TararebaServiceClient {
	return &tararebaServiceClient{cc}
}

func (c *tararebaServiceClient) GetCompetitionHistory(ctx context.Context, in *GetCompetitionHistoryRequest, opts ...grpc.CallOption) (*GetCompetitionHistoryResponse, error) {
	out := new(GetCompetitionHistoryResponse)
	err := c.cc.Invoke(ctx, "/tarareba_competition_history.TararebaService/GetCompetitionHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TararebaServiceServer is the server API for TararebaService service.
type TararebaServiceServer interface {
	// コンテスト参加履歴を返す
	GetCompetitionHistory(context.Context, *GetCompetitionHistoryRequest) (*GetCompetitionHistoryResponse, error)
}

// UnimplementedTararebaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTararebaServiceServer struct {
}

func (*UnimplementedTararebaServiceServer) GetCompetitionHistory(context.Context, *GetCompetitionHistoryRequest) (*GetCompetitionHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompetitionHistory not implemented")
}

func RegisterTararebaServiceServer(s *grpc.Server, srv TararebaServiceServer) {
	s.RegisterService(&_TararebaService_serviceDesc, srv)
}

func _TararebaService_GetCompetitionHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompetitionHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TararebaServiceServer).GetCompetitionHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tarareba_competition_history.TararebaService/GetCompetitionHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TararebaServiceServer).GetCompetitionHistory(ctx, req.(*GetCompetitionHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TararebaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tarareba_competition_history.TararebaService",
	HandlerType: (*TararebaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompetitionHistory",
			Handler:    _TararebaService_GetCompetitionHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tarareba_competition_history.proto",
}
