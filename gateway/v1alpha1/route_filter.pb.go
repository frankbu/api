// Copyright Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: gateway/v1alpha1/route_filter.proto

// $schema: istio.gateway.v1alpha1.RouteFilter
// $title: RouteFilter
// $description: Provides Istio route filter extensions for Gateway API.
// $location: https://istio.io/docs/reference/config/gateway/route-filter.html
// $aliases: [/docs/reference/config/gateway/v1alpha1/route-filter]
// $mode: none

// `RouteFilter` is an experimental API used to configure Istio extensions to the
// Kubernetes [Gateway API](https://gateway-api.sigs.k8s.io/).
// It can be referenced using the `extensionRef` field of an
// [HTTPRouteFilter](https://gateway-api.sigs.k8s.io/references/spec/#gateway.networking.k8s.io%2fv1beta1.HTTPRouteFilter)
// to configure Istio features that are not currently supported in the Gateway API
// itself.
//
// ```yaml
// apiVersion: gateway.networking.k8s.io/v1beta1
// kind: HTTPRoute
// metadata:
//   name: reviews
// spec:
//   parentRefs:
//   - kind: Service
//     name: reviews
//     port: 9080
//   rules:
//   - filters:
//     - type: ExtensionRef
//       extensionRef:
//         group: gateway.istio.io/v1alpha1
//         kind: RouteFilter
//         name: reviews
//     backendRefs:
//     - name: reviews-v1
//       port: 9080
// ---
// apiVersion: gateway.istio.io/v1alpha1
// kind: RouteFilter
// metadata:
//   name: reviews
// spec:
//   timeout: 10s
//   retries:
//     attempts: 3
//     perTryTimeout: 2s
//     retryOn: connect-failure,refused-stream,503
// ```
//
// Fields in the `RouteFilter` type should be considered both experimental and
// potentially temporary. New fields will be added as needed and others will likely
// be removed in the future if Gateway API adds support for their corresponding
// features.
//

package v1alpha1

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1beta1 "istio.io/api/networking/v1beta1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// `RouteFilter` is an experimental API used to configure Istio extensions to the
// Kubernetes [Gateway API](https://gateway-api.sigs.k8s.io/).
//
// <!-- crd generation tags
// +cue-gen:RouteFilter:groupName:gateway.istio.io
// +cue-gen:RouteFilter:version:v1alpha1
// +cue-gen:RouteFilter:storageVersion
// +cue-gen:RouteFilter:annotations:helm.sh/resource-policy=keep
// +cue-gen:RouteFilter:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:RouteFilter:subresource:status
// +cue-gen:RouteFilter:scope:Namespaced
// +cue-gen:RouteFilter:resource:categories=istio-io,gateway-istio-io,plural=RouteFilters
// +cue-gen:RouteFilter:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=gateway.istio.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type RouteFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timeout for HTTP requests, default is disabled.
	Timeout *duration.Duration `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Retry policy for HTTP requests.
	Retries *v1beta1.HTTPRetry `protobuf:"bytes,2,opt,name=retries,proto3" json:"retries,omitempty"`
	// Fault injection policy to apply on HTTP traffic at the client side.
	// Note that timeouts or retries will not be enabled when faults are
	// enabled on the client side.
	Fault *v1beta1.HTTPFaultInjection `protobuf:"bytes,3,opt,name=fault,proto3" json:"fault,omitempty"`
}

func (x *RouteFilter) Reset() {
	*x = RouteFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_v1alpha1_route_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteFilter) ProtoMessage() {}

func (x *RouteFilter) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1alpha1_route_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteFilter.ProtoReflect.Descriptor instead.
func (*RouteFilter) Descriptor() ([]byte, []int) {
	return file_gateway_v1alpha1_route_filter_proto_rawDescGZIP(), []int{0}
}

func (x *RouteFilter) GetTimeout() *duration.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *RouteFilter) GetRetries() *v1beta1.HTTPRetry {
	if x != nil {
		return x.Retries
	}
	return nil
}

func (x *RouteFilter) GetFault() *v1beta1.HTTPFaultInjection {
	if x != nil {
		return x.Fault
	}
	return nil
}

var File_gateway_v1alpha1_route_filter_proto protoreflect.FileDescriptor

var file_gateway_v1alpha1_route_filter_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3d, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x05, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x73, 0x74,
	0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x49,
	0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42,
	0x1f, 0x5a, 0x1d, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_v1alpha1_route_filter_proto_rawDescOnce sync.Once
	file_gateway_v1alpha1_route_filter_proto_rawDescData = file_gateway_v1alpha1_route_filter_proto_rawDesc
)

func file_gateway_v1alpha1_route_filter_proto_rawDescGZIP() []byte {
	file_gateway_v1alpha1_route_filter_proto_rawDescOnce.Do(func() {
		file_gateway_v1alpha1_route_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_v1alpha1_route_filter_proto_rawDescData)
	})
	return file_gateway_v1alpha1_route_filter_proto_rawDescData
}

var file_gateway_v1alpha1_route_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gateway_v1alpha1_route_filter_proto_goTypes = []interface{}{
	(*RouteFilter)(nil),                // 0: istio.gateway.v1alpha1.RouteFilter
	(*duration.Duration)(nil),          // 1: google.protobuf.Duration
	(*v1beta1.HTTPRetry)(nil),          // 2: istio.networking.v1beta1.HTTPRetry
	(*v1beta1.HTTPFaultInjection)(nil), // 3: istio.networking.v1beta1.HTTPFaultInjection
}
var file_gateway_v1alpha1_route_filter_proto_depIdxs = []int32{
	1, // 0: istio.gateway.v1alpha1.RouteFilter.timeout:type_name -> google.protobuf.Duration
	2, // 1: istio.gateway.v1alpha1.RouteFilter.retries:type_name -> istio.networking.v1beta1.HTTPRetry
	3, // 2: istio.gateway.v1alpha1.RouteFilter.fault:type_name -> istio.networking.v1beta1.HTTPFaultInjection
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gateway_v1alpha1_route_filter_proto_init() }
func file_gateway_v1alpha1_route_filter_proto_init() {
	if File_gateway_v1alpha1_route_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_v1alpha1_route_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteFilter); i {
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
			RawDescriptor: file_gateway_v1alpha1_route_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_v1alpha1_route_filter_proto_goTypes,
		DependencyIndexes: file_gateway_v1alpha1_route_filter_proto_depIdxs,
		MessageInfos:      file_gateway_v1alpha1_route_filter_proto_msgTypes,
	}.Build()
	File_gateway_v1alpha1_route_filter_proto = out.File
	file_gateway_v1alpha1_route_filter_proto_rawDesc = nil
	file_gateway_v1alpha1_route_filter_proto_goTypes = nil
	file_gateway_v1alpha1_route_filter_proto_depIdxs = nil
}
