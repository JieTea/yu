// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: txpool.proto

package goproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_txpool_proto protoreflect.FileDescriptor

var file_txpool_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x74,
	0x78, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc0, 0x02, 0x0a, 0x06, 0x54, 0x78, 0x70,
	0x6f, 0x6f, 0x6c, 0x12, 0x28, 0x0a, 0x08, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x04, 0x2e, 0x55, 0x36, 0x34, 0x12, 0x1d, 0x0a,
	0x09, 0x42, 0x61, 0x73, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0a, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x1a, 0x04, 0x2e, 0x45, 0x72, 0x72, 0x12, 0x20, 0x0a, 0x0c,
	0x54, 0x72, 0x69, 0x70, 0x6f, 0x64, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0a, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x1a, 0x04, 0x2e, 0x45, 0x72, 0x72, 0x12, 0x22,
	0x0a, 0x0e, 0x4e, 0x65, 0x63, 0x65, 0x73, 0x73, 0x61, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x12, 0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x1a, 0x04, 0x2e, 0x45,
	0x72, 0x72, 0x12, 0x1a, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x0a, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x1a, 0x04, 0x2e, 0x45, 0x72, 0x72, 0x12, 0x25,
	0x0a, 0x0b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x10, 0x2e,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x73, 0x1a,
	0x04, 0x2e, 0x45, 0x72, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x73,
	0x54, 0x78, 0x6e, 0x73, 0x12, 0x0b, 0x2e, 0x54, 0x78, 0x6e, 0x73, 0x48, 0x61, 0x73, 0x68, 0x65,
	0x73, 0x1a, 0x04, 0x2e, 0x45, 0x72, 0x72, 0x12, 0x1b, 0x0a, 0x04, 0x50, 0x61, 0x63, 0x6b, 0x12,
	0x04, 0x2e, 0x55, 0x36, 0x34, 0x1a, 0x0d, 0x2e, 0x54, 0x78, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x04, 0x2e, 0x45, 0x72, 0x72, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_txpool_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil),   // 0: google.protobuf.Empty
	(*SignedTxn)(nil),       // 1: SignedTxn
	(*BatchSignedTxns)(nil), // 2: BatchSignedTxns
	(*TxnsHashes)(nil),      // 3: TxnsHashes
	(*U64)(nil),             // 4: U64
	(*Err)(nil),             // 5: Err
	(*TxnsResponse)(nil),    // 6: TxnsResponse
}
var file_txpool_proto_depIdxs = []int32{
	0, // 0: Txpool.PoolSize:input_type -> google.protobuf.Empty
	1, // 1: Txpool.BaseCheck:input_type -> SignedTxn
	1, // 2: Txpool.TripodsCheck:input_type -> SignedTxn
	1, // 3: Txpool.NecessaryCheck:input_type -> SignedTxn
	1, // 4: Txpool.Insert:input_type -> SignedTxn
	2, // 5: Txpool.BatchInsert:input_type -> BatchSignedTxns
	3, // 6: Txpool.RemovesTxns:input_type -> TxnsHashes
	4, // 7: Txpool.Pack:input_type -> U64
	0, // 8: Txpool.Reset:input_type -> google.protobuf.Empty
	4, // 9: Txpool.PoolSize:output_type -> U64
	5, // 10: Txpool.BaseCheck:output_type -> Err
	5, // 11: Txpool.TripodsCheck:output_type -> Err
	5, // 12: Txpool.NecessaryCheck:output_type -> Err
	5, // 13: Txpool.Insert:output_type -> Err
	5, // 14: Txpool.BatchInsert:output_type -> Err
	5, // 15: Txpool.RemovesTxns:output_type -> Err
	6, // 16: Txpool.Pack:output_type -> TxnsResponse
	5, // 17: Txpool.Reset:output_type -> Err
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_txpool_proto_init() }
func file_txpool_proto_init() {
	if File_txpool_proto != nil {
		return
	}
	file_base_types_proto_init()
	file_txn_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_txpool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_txpool_proto_goTypes,
		DependencyIndexes: file_txpool_proto_depIdxs,
	}.Build()
	File_txpool_proto = out.File
	file_txpool_proto_rawDesc = nil
	file_txpool_proto_goTypes = nil
	file_txpool_proto_depIdxs = nil
}
