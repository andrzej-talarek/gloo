// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/retries/retries.proto

package retries

import (
	reflect "reflect"
	sync "sync"

	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RetryPolicyInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the max interval for a retry
	MaxInterval *duration.Duration `protobuf:"bytes,1,opt,name=max_interval,json=maxInterval,proto3" json:"max_interval,omitempty"`
	// Specifies the base interval for a retry
	BaseInterval *duration.Duration `protobuf:"bytes,2,opt,name=base_interval,json=baseInterval,proto3" json:"base_interval,omitempty"`
}

func (x *RetryPolicyInterval) Reset() {
	*x = RetryPolicyInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryPolicyInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryPolicyInterval) ProtoMessage() {}

func (x *RetryPolicyInterval) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryPolicyInterval.ProtoReflect.Descriptor instead.
func (*RetryPolicyInterval) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescGZIP(), []int{0}
}

func (x *RetryPolicyInterval) GetMaxInterval() *duration.Duration {
	if x != nil {
		return x.MaxInterval
	}
	return nil
}

func (x *RetryPolicyInterval) GetBaseInterval() *duration.Duration {
	if x != nil {
		return x.BaseInterval
	}
	return nil
}

// Retry Policy applied at the Route and/or Virtual Hosts levels.
type RetryPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the conditions under which retry takes place. These are the same
	// conditions [documented for Envoy](https://www.envoyproxy.io/docs/envoy/v1.14.1/configuration/http/http_filters/router_filter#config-http-filters-router-x-envoy-retry-on)
	RetryOn string `protobuf:"bytes,1,opt,name=retry_on,json=retryOn,proto3" json:"retry_on,omitempty"`
	// Specifies the allowed number of retries. This parameter is optional and
	// defaults to 1. These are the same conditions [documented for Envoy](https://www.envoyproxy.io/docs/envoy/v1.14.1/configuration/http/http_filters/router_filter#config-http-filters-router-x-envoy-retry-on)
	NumRetries uint32 `protobuf:"varint,2,opt,name=num_retries,json=numRetries,proto3" json:"num_retries,omitempty"`
	// Specifies a non-zero upstream timeout per retry attempt. This parameter is optional.
	PerTryTimeout       *duration.Duration   `protobuf:"bytes,3,opt,name=per_try_timeout,json=perTryTimeout,proto3" json:"per_try_timeout,omitempty"`
	RetryPolicyInterval *RetryPolicyInterval `protobuf:"bytes,4,opt,name=retry_policy_interval,json=retryPolicyInterval,proto3" json:"retry_policy_interval,omitempty"`
}

func (x *RetryPolicy) Reset() {
	*x = RetryPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryPolicy) ProtoMessage() {}

func (x *RetryPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryPolicy.ProtoReflect.Descriptor instead.
func (*RetryPolicy) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescGZIP(), []int{1}
}

func (x *RetryPolicy) GetRetryOn() string {
	if x != nil {
		return x.RetryOn
	}
	return ""
}

func (x *RetryPolicy) GetNumRetries() uint32 {
	if x != nil {
		return x.NumRetries
	}
	return 0
}

func (x *RetryPolicy) GetPerTryTimeout() *duration.Duration {
	if x != nil {
		return x.PerTryTimeout
	}
	return nil
}

func (x *RetryPolicy) GetRetryPolicyInterval() *RetryPolicyInterval {
	if x != nil {
		return x.RetryPolicyInterval
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x72,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x72, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93,
	0x01, 0x0a, 0x13, 0x52, 0x65, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x3c, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x12, 0x3e, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x22, 0xf3, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x74, 0x72, 0x79, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x79, 0x4f, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x41, 0x0a, 0x0f, 0x70, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x65, 0x72, 0x54, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x65, 0x0a, 0x15, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x13, 0x72, 0x65, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x4e, 0x5a, 0x40, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0xc0, 0xf5,
	0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_goTypes = []interface{}{
	(*RetryPolicyInterval)(nil), // 0: retries.options.gloo.solo.io.RetryPolicyInterval
	(*RetryPolicy)(nil),         // 1: retries.options.gloo.solo.io.RetryPolicy
	(*duration.Duration)(nil),   // 2: google.protobuf.Duration
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_depIdxs = []int32{
	2, // 0: retries.options.gloo.solo.io.RetryPolicyInterval.max_interval:type_name -> google.protobuf.Duration
	2, // 1: retries.options.gloo.solo.io.RetryPolicyInterval.base_interval:type_name -> google.protobuf.Duration
	2, // 2: retries.options.gloo.solo.io.RetryPolicy.per_try_timeout:type_name -> google.protobuf.Duration
	0, // 3: retries.options.gloo.solo.io.RetryPolicy.retry_policy_interval:type_name -> retries.options.gloo.solo.io.RetryPolicyInterval
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryPolicyInterval); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryPolicy); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_depIdxs = nil
}
