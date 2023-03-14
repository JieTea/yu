// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: blockchain.proto

package goproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block *CompactBlock `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Error string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *BlockResponse) Reset() {
	*x = BlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResponse) ProtoMessage() {}

func (x *BlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResponse.ProtoReflect.Descriptor instead.
func (*BlockResponse) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{0}
}

func (x *BlockResponse) GetBlock() *CompactBlock {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *BlockResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type BlocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*CompactBlock `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Error  string          `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *BlocksResponse) Reset() {
	*x = BlocksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlocksResponse) ProtoMessage() {}

func (x *BlocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlocksResponse.ProtoReflect.Descriptor instead.
func (*BlocksResponse) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{1}
}

func (x *BlocksResponse) GetBlocks() []*CompactBlock {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *BlocksResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type RangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartHeight uint64 `protobuf:"varint,1,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	EndHeight   uint64 `protobuf:"varint,2,opt,name=end_height,json=endHeight,proto3" json:"end_height,omitempty"`
}

func (x *RangeRequest) Reset() {
	*x = RangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RangeRequest) ProtoMessage() {}

func (x *RangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RangeRequest.ProtoReflect.Descriptor instead.
func (*RangeRequest) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{2}
}

func (x *RangeRequest) GetStartHeight() uint64 {
	if x != nil {
		return x.StartHeight
	}
	return 0
}

func (x *RangeRequest) GetEndHeight() uint64 {
	if x != nil {
		return x.EndHeight
	}
	return 0
}

var File_blockchain_proto protoreflect.FileDescriptor

var file_blockchain_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a,
	0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4d, 0x0a, 0x0e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x50, 0x0a, 0x0c, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x6e, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0xe4, 0x03, 0x0a, 0x0a,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x34, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x0d,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x04, 0x2e,
	0x45, 0x72, 0x72, 0x12, 0x22, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x0d, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x1a, 0x04, 0x2e, 0x45, 0x72, 0x72, 0x12, 0x26, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x0a, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x1a,
	0x0e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0a,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x05, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x12, 0x22, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x0d, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x1a,
	0x04, 0x2e, 0x45, 0x72, 0x72, 0x12, 0x27, 0x0a, 0x08, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x12, 0x0a, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x0f, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x08, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x12, 0x0a, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x04, 0x2e, 0x45, 0x72, 0x72, 0x12, 0x3b, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x45, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x12, 0x0d, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_proto_rawDescOnce sync.Once
	file_blockchain_proto_rawDescData = file_blockchain_proto_rawDesc
)

func file_blockchain_proto_rawDescGZIP() []byte {
	file_blockchain_proto_rawDescOnce.Do(func() {
		file_blockchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_proto_rawDescData)
	})
	return file_blockchain_proto_rawDescData
}

var file_blockchain_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_blockchain_proto_goTypes = []interface{}{
	(*BlockResponse)(nil),  // 0: BlockResponse
	(*BlocksResponse)(nil), // 1: BlocksResponse
	(*RangeRequest)(nil),   // 2: RangeRequest
	(*CompactBlock)(nil),   // 3: CompactBlock
	(*emptypb.Empty)(nil),  // 4: google.protobuf.Empty
	(*BlockHash)(nil),      // 5: BlockHash
	(*Err)(nil),            // 6: Err
	(*Bool)(nil),           // 7: Bool
}
var file_blockchain_proto_depIdxs = []int32{
	3,  // 0: BlockResponse.block:type_name -> CompactBlock
	3,  // 1: BlocksResponse.blocks:type_name -> CompactBlock
	4,  // 2: BlockChain.GetGenesis:input_type -> google.protobuf.Empty
	3,  // 3: BlockChain.SetGenesis:input_type -> CompactBlock
	3,  // 4: BlockChain.AppendBlock:input_type -> CompactBlock
	5,  // 5: BlockChain.GetBlock:input_type -> BlockHash
	5,  // 6: BlockChain.ExistsBlock:input_type -> BlockHash
	3,  // 7: BlockChain.UpdateBlock:input_type -> CompactBlock
	5,  // 8: BlockChain.Children:input_type -> BlockHash
	5,  // 9: BlockChain.Finalize:input_type -> BlockHash
	4,  // 10: BlockChain.GetFinalizedBlock:input_type -> google.protobuf.Empty
	4,  // 11: BlockChain.GetEndBlock:input_type -> google.protobuf.Empty
	2,  // 12: BlockChain.GetRangeBlocks:input_type -> RangeRequest
	0,  // 13: BlockChain.GetGenesis:output_type -> BlockResponse
	6,  // 14: BlockChain.SetGenesis:output_type -> Err
	6,  // 15: BlockChain.AppendBlock:output_type -> Err
	0,  // 16: BlockChain.GetBlock:output_type -> BlockResponse
	7,  // 17: BlockChain.ExistsBlock:output_type -> Bool
	6,  // 18: BlockChain.UpdateBlock:output_type -> Err
	1,  // 19: BlockChain.Children:output_type -> BlocksResponse
	6,  // 20: BlockChain.Finalize:output_type -> Err
	0,  // 21: BlockChain.GetFinalizedBlock:output_type -> BlockResponse
	0,  // 22: BlockChain.GetEndBlock:output_type -> BlockResponse
	1,  // 23: BlockChain.GetRangeBlocks:output_type -> BlocksResponse
	13, // [13:24] is the sub-list for method output_type
	2,  // [2:13] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_blockchain_proto_init() }
func file_blockchain_proto_init() {
	if File_blockchain_proto != nil {
		return
	}
	file_block_proto_init()
	file_base_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_blockchain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResponse); i {
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
		file_blockchain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlocksResponse); i {
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
		file_blockchain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RangeRequest); i {
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
			RawDescriptor: file_blockchain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blockchain_proto_goTypes,
		DependencyIndexes: file_blockchain_proto_depIdxs,
		MessageInfos:      file_blockchain_proto_msgTypes,
	}.Build()
	File_blockchain_proto = out.File
	file_blockchain_proto_rawDesc = nil
	file_blockchain_proto_goTypes = nil
	file_blockchain_proto_depIdxs = nil
}
