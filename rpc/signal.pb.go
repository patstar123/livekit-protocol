// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: rpc/signal.proto

package rpc

import (
	livekit "github.com/livekit/protocol/livekit"
	_ "github.com/livekit/psrpc/protoc-gen-psrpc/options"
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

type RelaySignalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartSession *livekit.StartSession    `protobuf:"bytes,1,opt,name=start_session,json=startSession,proto3" json:"start_session,omitempty"`
	Requests     []*livekit.SignalRequest `protobuf:"bytes,3,rep,name=requests,proto3" json:"requests,omitempty"`
	Seq          uint64                   `protobuf:"varint,4,opt,name=seq,proto3" json:"seq,omitempty"`
	Close        bool                     `protobuf:"varint,5,opt,name=close,proto3" json:"close,omitempty"`
}

func (x *RelaySignalRequest) Reset() {
	*x = RelaySignalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_signal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelaySignalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelaySignalRequest) ProtoMessage() {}

func (x *RelaySignalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_signal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelaySignalRequest.ProtoReflect.Descriptor instead.
func (*RelaySignalRequest) Descriptor() ([]byte, []int) {
	return file_rpc_signal_proto_rawDescGZIP(), []int{0}
}

func (x *RelaySignalRequest) GetStartSession() *livekit.StartSession {
	if x != nil {
		return x.StartSession
	}
	return nil
}

func (x *RelaySignalRequest) GetRequests() []*livekit.SignalRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

func (x *RelaySignalRequest) GetSeq() uint64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *RelaySignalRequest) GetClose() bool {
	if x != nil {
		return x.Close
	}
	return false
}

type RelaySignalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responses []*livekit.SignalResponse `protobuf:"bytes,2,rep,name=responses,proto3" json:"responses,omitempty"`
	Seq       uint64                    `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
	Close     bool                      `protobuf:"varint,4,opt,name=close,proto3" json:"close,omitempty"`
}

func (x *RelaySignalResponse) Reset() {
	*x = RelaySignalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_signal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelaySignalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelaySignalResponse) ProtoMessage() {}

func (x *RelaySignalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_signal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelaySignalResponse.ProtoReflect.Descriptor instead.
func (*RelaySignalResponse) Descriptor() ([]byte, []int) {
	return file_rpc_signal_proto_rawDescGZIP(), []int{1}
}

func (x *RelaySignalResponse) GetResponses() []*livekit.SignalResponse {
	if x != nil {
		return x.Responses
	}
	return nil
}

func (x *RelaySignalResponse) GetSeq() uint64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *RelaySignalResponse) GetClose() bool {
	if x != nil {
		return x.Close
	}
	return false
}

var File_rpc_signal_proto protoreflect.FileDescriptor

var file_rpc_signal_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x72, 0x74, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xac, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x22, 0x74, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x32, 0x63, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x12, 0x59, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12,
	0x17, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x6c, 0x61, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x17, 0xb2, 0x89, 0x01, 0x13, 0x10, 0x01, 0x1a, 0x0d, 0x12, 0x07, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x20, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rpc_signal_proto_rawDescOnce sync.Once
	file_rpc_signal_proto_rawDescData = file_rpc_signal_proto_rawDesc
)

func file_rpc_signal_proto_rawDescGZIP() []byte {
	file_rpc_signal_proto_rawDescOnce.Do(func() {
		file_rpc_signal_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_signal_proto_rawDescData)
	})
	return file_rpc_signal_proto_rawDescData
}

var file_rpc_signal_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_signal_proto_goTypes = []interface{}{
	(*RelaySignalRequest)(nil),     // 0: rpc.RelaySignalRequest
	(*RelaySignalResponse)(nil),    // 1: rpc.RelaySignalResponse
	(*livekit.StartSession)(nil),   // 2: livekit.StartSession
	(*livekit.SignalRequest)(nil),  // 3: livekit.SignalRequest
	(*livekit.SignalResponse)(nil), // 4: livekit.SignalResponse
}
var file_rpc_signal_proto_depIdxs = []int32{
	2, // 0: rpc.RelaySignalRequest.start_session:type_name -> livekit.StartSession
	3, // 1: rpc.RelaySignalRequest.requests:type_name -> livekit.SignalRequest
	4, // 2: rpc.RelaySignalResponse.responses:type_name -> livekit.SignalResponse
	0, // 3: rpc.Signal.RelaySignal:input_type -> rpc.RelaySignalRequest
	1, // 4: rpc.Signal.RelaySignal:output_type -> rpc.RelaySignalResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_signal_proto_init() }
func file_rpc_signal_proto_init() {
	if File_rpc_signal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_signal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelaySignalRequest); i {
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
		file_rpc_signal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelaySignalResponse); i {
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
			RawDescriptor: file_rpc_signal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_signal_proto_goTypes,
		DependencyIndexes: file_rpc_signal_proto_depIdxs,
		MessageInfos:      file_rpc_signal_proto_msgTypes,
	}.Build()
	File_rpc_signal_proto = out.File
	file_rpc_signal_proto_rawDesc = nil
	file_rpc_signal_proto_goTypes = nil
	file_rpc_signal_proto_depIdxs = nil
}
