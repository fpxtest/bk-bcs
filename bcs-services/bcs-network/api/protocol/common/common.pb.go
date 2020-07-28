// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// ErrCode error code
type ErrCode int32

const (
	// 0 Sunccess
	ErrCode_ERROR_OK ErrCode = 0
	// 81000 cloud netservice, invalid parameter
	ErrCode_ERROR_CLOUD_NETSERVICE_INVALID_PARAMS ErrCode = 81000
	// 81001 cloud netservice, storage operation failed
	ErrCode_ERROR_CLOUD_NETSERVICE_STOREOPS_FAILED ErrCode = 81001
	// 81002 cloud netservice, call cloud api failed
	ErrCode_ERROR_CLOUD_NETSERVICE_CLOUDAPI_FAILED ErrCode = 81002
	// 81003 cloud netservice, assign ip failed
	ErrCode_ERROR_CLOUD_NETSERVICE_CLOUDAPI_ASSIGNIP_FAILED ErrCode = 81003
	// 81004 cloud netserivce, unassign ip failed
	ErrCode_ERROR_CLOUD_NETSERVICE_CLOUDAPI_UNASSIGNIP_FAILED ErrCode = 81004
	// 81005 cloud netservice, query eni failed
	ErrCode_ERROR_CLOUD_NETSERVICE_CLOUDAPI_QUERY_ENI_FAILED ErrCode = 81005
	// 81006 cloud netservice, eni info not match
	ErrCode_ERROR_CLOUD_NETSERVICE_CLOUDAPI_ENI_INFO_NOTMATCH ErrCode = 81006
	// 81007 cloud netservice, query subnet from store failed
	ErrCode_ERROR_CLOUD_NETSERVICE_CLOUDAPI_QUERY_SUBNET_FROM_STORE_FAILED ErrCode = 81007
	// 81008 cloud netservice, try to delete subent with active ip
	ErrCode_ERROR_CLOUD_NETSERVICE_TRY_TO_DELETE_ACTIVE_SUBNET ErrCode = 81008
	// 81009 cloud netservice, try to delete enabled subnet
	ErrCode_ERROR_CLOUD_NETSERVICE_TRY_TO_DELETE_ENABLED_SUBNET ErrCode = 81009
	// 81010 cloud netservice, subnet is disabled
	ErrCode_ERROR_CLOUD_NETSERVICE_SUBNET_IS_DISABLED ErrCode = 81010
	// 81011 cloud netservice, try to allocate active ip
	ErrCode_ERROR_CLOUD_NETSERVICE_TRY_TO_ALLOCATE_ACTIVE_IP ErrCode = 81011
	// 81012 cloud netservice, migrate ip failed
	ErrCode_ERROR_CLOUD_NETSERVICE_MIGRATE_IP_FAILED ErrCode = 81012
	// 81013 cloud netservice, allocated ip info not match request
	ErrCode_ERROR_CLOUD_NETSERVICE_ALLOCATE_IP_NOT_MATCH ErrCode = 81013
	// 81013 cloud netservice, try to clean active ip
	ErrCode_ERROR_CLOUD_NETSERVICE_TRY_TO_CLEAN_ACTIVE_IP ErrCode = 81014
	// 82000 cloud netagent, find no pod
	ErrCode_ERROR_CLOUD_NETAGENT_POD_NOT_FOUND ErrCode = 82000
	// 82001 cloud netagent, get ip failed
	ErrCode_ERROR_CLOUD_NETAGENT_ALLOCATE_IP_FAILED ErrCode = 82001
	// 82002 cloud netagent, release ip failed
	ErrCode_ERROR_CLOUD_NETAGENT_RELEASE_IP_FAILED ErrCode = 82002
	// 82003 cloud netagent, k8s api server ops failed
	ErrCode_ERROR_CLOUD_NETAGENT_K8S_API_SERVER_OPS_FAILED ErrCode = 82003
	// 82004 cloud netagent, nodenetwork not available
	ErrCode_ERROR_CLOUD_NETAGENT_NODENETWORK_NOT_AVAILABLE ErrCode = 82004
	// 82005 cloud netagent, pod workload not found
	ErrCode_ERROR_CLOUD_NETAGENT_POD_WORKLOAD_NOT_FOUND ErrCode = 82005
	// 82006 cloud netagent, ip info invalid
	ErrCode_ERROR_CLOUD_NETAGENT_INVALID_IP_INFO ErrCode = 82006
	// 82007 cloud netagent, invalid parameter
	ErrCode_ERROR_CLOUD_NETAGENT_INVALID_PARAMS ErrCode = 82007
	// 82008 cloud netagent, workload not support fixed ip feature
	ErrCode_ERROR_CLOUD_NETAGENT_WORKLOAD_NOT_SUPPORT_FIXED_IP_FEATURE ErrCode = 82008
)

var ErrCode_name = map[int32]string{
	0:     "ERROR_OK",
	81000: "ERROR_CLOUD_NETSERVICE_INVALID_PARAMS",
	81001: "ERROR_CLOUD_NETSERVICE_STOREOPS_FAILED",
	81002: "ERROR_CLOUD_NETSERVICE_CLOUDAPI_FAILED",
	81003: "ERROR_CLOUD_NETSERVICE_CLOUDAPI_ASSIGNIP_FAILED",
	81004: "ERROR_CLOUD_NETSERVICE_CLOUDAPI_UNASSIGNIP_FAILED",
	81005: "ERROR_CLOUD_NETSERVICE_CLOUDAPI_QUERY_ENI_FAILED",
	81006: "ERROR_CLOUD_NETSERVICE_CLOUDAPI_ENI_INFO_NOTMATCH",
	81007: "ERROR_CLOUD_NETSERVICE_CLOUDAPI_QUERY_SUBNET_FROM_STORE_FAILED",
	81008: "ERROR_CLOUD_NETSERVICE_TRY_TO_DELETE_ACTIVE_SUBNET",
	81009: "ERROR_CLOUD_NETSERVICE_TRY_TO_DELETE_ENABLED_SUBNET",
	81010: "ERROR_CLOUD_NETSERVICE_SUBNET_IS_DISABLED",
	81011: "ERROR_CLOUD_NETSERVICE_TRY_TO_ALLOCATE_ACTIVE_IP",
	81012: "ERROR_CLOUD_NETSERVICE_MIGRATE_IP_FAILED",
	81013: "ERROR_CLOUD_NETSERVICE_ALLOCATE_IP_NOT_MATCH",
	81014: "ERROR_CLOUD_NETSERVICE_TRY_TO_CLEAN_ACTIVE_IP",
	82000: "ERROR_CLOUD_NETAGENT_POD_NOT_FOUND",
	82001: "ERROR_CLOUD_NETAGENT_ALLOCATE_IP_FAILED",
	82002: "ERROR_CLOUD_NETAGENT_RELEASE_IP_FAILED",
	82003: "ERROR_CLOUD_NETAGENT_K8S_API_SERVER_OPS_FAILED",
	82004: "ERROR_CLOUD_NETAGENT_NODENETWORK_NOT_AVAILABLE",
	82005: "ERROR_CLOUD_NETAGENT_POD_WORKLOAD_NOT_FOUND",
	82006: "ERROR_CLOUD_NETAGENT_INVALID_IP_INFO",
	82007: "ERROR_CLOUD_NETAGENT_INVALID_PARAMS",
	82008: "ERROR_CLOUD_NETAGENT_WORKLOAD_NOT_SUPPORT_FIXED_IP_FEATURE",
}

var ErrCode_value = map[string]int32{
	"ERROR_OK":                                                       0,
	"ERROR_CLOUD_NETSERVICE_INVALID_PARAMS":                          81000,
	"ERROR_CLOUD_NETSERVICE_STOREOPS_FAILED":                         81001,
	"ERROR_CLOUD_NETSERVICE_CLOUDAPI_FAILED":                         81002,
	"ERROR_CLOUD_NETSERVICE_CLOUDAPI_ASSIGNIP_FAILED":                81003,
	"ERROR_CLOUD_NETSERVICE_CLOUDAPI_UNASSIGNIP_FAILED":              81004,
	"ERROR_CLOUD_NETSERVICE_CLOUDAPI_QUERY_ENI_FAILED":               81005,
	"ERROR_CLOUD_NETSERVICE_CLOUDAPI_ENI_INFO_NOTMATCH":              81006,
	"ERROR_CLOUD_NETSERVICE_CLOUDAPI_QUERY_SUBNET_FROM_STORE_FAILED": 81007,
	"ERROR_CLOUD_NETSERVICE_TRY_TO_DELETE_ACTIVE_SUBNET":             81008,
	"ERROR_CLOUD_NETSERVICE_TRY_TO_DELETE_ENABLED_SUBNET":            81009,
	"ERROR_CLOUD_NETSERVICE_SUBNET_IS_DISABLED":                      81010,
	"ERROR_CLOUD_NETSERVICE_TRY_TO_ALLOCATE_ACTIVE_IP":               81011,
	"ERROR_CLOUD_NETSERVICE_MIGRATE_IP_FAILED":                       81012,
	"ERROR_CLOUD_NETSERVICE_ALLOCATE_IP_NOT_MATCH":                   81013,
	"ERROR_CLOUD_NETSERVICE_TRY_TO_CLEAN_ACTIVE_IP":                  81014,
	"ERROR_CLOUD_NETAGENT_POD_NOT_FOUND":                             82000,
	"ERROR_CLOUD_NETAGENT_ALLOCATE_IP_FAILED":                        82001,
	"ERROR_CLOUD_NETAGENT_RELEASE_IP_FAILED":                         82002,
	"ERROR_CLOUD_NETAGENT_K8S_API_SERVER_OPS_FAILED":                 82003,
	"ERROR_CLOUD_NETAGENT_NODENETWORK_NOT_AVAILABLE":                 82004,
	"ERROR_CLOUD_NETAGENT_POD_WORKLOAD_NOT_FOUND":                    82005,
	"ERROR_CLOUD_NETAGENT_INVALID_IP_INFO":                           82006,
	"ERROR_CLOUD_NETAGENT_INVALID_PARAMS":                            82007,
	"ERROR_CLOUD_NETAGENT_WORKLOAD_NOT_SUPPORT_FIXED_IP_FEATURE":     82008,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type CloudSubnet struct {
	VpcID                string   `protobuf:"bytes,1,opt,name=vpcID,proto3" json:"vpcID,omitempty"`
	Region               string   `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	Zone                 string   `protobuf:"bytes,3,opt,name=zone,proto3" json:"zone,omitempty"`
	SubnetID             string   `protobuf:"bytes,4,opt,name=subnetID,proto3" json:"subnetID,omitempty"`
	SubnetCidr           string   `protobuf:"bytes,5,opt,name=subnetCidr,proto3" json:"subnetCidr,omitempty"`
	AvailableIPNum       uint64   `protobuf:"varint,6,opt,name=availableIPNum,proto3" json:"availableIPNum,omitempty"`
	State                int32    `protobuf:"varint,7,opt,name=state,proto3" json:"state,omitempty"`
	CreateTime           string   `protobuf:"bytes,8,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime           string   `protobuf:"bytes,9,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloudSubnet) Reset()         { *m = CloudSubnet{} }
func (m *CloudSubnet) String() string { return proto.CompactTextString(m) }
func (*CloudSubnet) ProtoMessage()    {}
func (*CloudSubnet) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *CloudSubnet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudSubnet.Unmarshal(m, b)
}
func (m *CloudSubnet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudSubnet.Marshal(b, m, deterministic)
}
func (m *CloudSubnet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudSubnet.Merge(m, src)
}
func (m *CloudSubnet) XXX_Size() int {
	return xxx_messageInfo_CloudSubnet.Size(m)
}
func (m *CloudSubnet) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudSubnet.DiscardUnknown(m)
}

var xxx_messageInfo_CloudSubnet proto.InternalMessageInfo

func (m *CloudSubnet) GetVpcID() string {
	if m != nil {
		return m.VpcID
	}
	return ""
}

func (m *CloudSubnet) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *CloudSubnet) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *CloudSubnet) GetSubnetID() string {
	if m != nil {
		return m.SubnetID
	}
	return ""
}

func (m *CloudSubnet) GetSubnetCidr() string {
	if m != nil {
		return m.SubnetCidr
	}
	return ""
}

func (m *CloudSubnet) GetAvailableIPNum() uint64 {
	if m != nil {
		return m.AvailableIPNum
	}
	return 0
}

func (m *CloudSubnet) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *CloudSubnet) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *CloudSubnet) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

type IPObject struct {
	VpcID                string   `protobuf:"bytes,1,opt,name=vpcID,proto3" json:"vpcID,omitempty"`
	Region               string   `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	SubnetID             string   `protobuf:"bytes,3,opt,name=subnetID,proto3" json:"subnetID,omitempty"`
	SubnetCidr           string   `protobuf:"bytes,4,opt,name=subnetCidr,proto3" json:"subnetCidr,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Cluster              string   `protobuf:"bytes,6,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace            string   `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	PodName              string   `protobuf:"bytes,8,opt,name=podName,proto3" json:"podName,omitempty"`
	WorkloadName         string   `protobuf:"bytes,9,opt,name=workloadName,proto3" json:"workloadName,omitempty"`
	WorkloadKind         string   `protobuf:"bytes,10,opt,name=workloadKind,proto3" json:"workloadKind,omitempty"`
	Host                 string   `protobuf:"bytes,11,opt,name=host,proto3" json:"host,omitempty"`
	ContainerID          string   `protobuf:"bytes,12,opt,name=containerID,proto3" json:"containerID,omitempty"`
	EniID                string   `protobuf:"bytes,13,opt,name=eniID,proto3" json:"eniID,omitempty"`
	CreateTime           string   `protobuf:"bytes,14,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime           string   `protobuf:"bytes,15,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	IsFixed              bool     `protobuf:"varint,16,opt,name=isFixed,proto3" json:"isFixed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPObject) Reset()         { *m = IPObject{} }
func (m *IPObject) String() string { return proto.CompactTextString(m) }
func (*IPObject) ProtoMessage()    {}
func (*IPObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *IPObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPObject.Unmarshal(m, b)
}
func (m *IPObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPObject.Marshal(b, m, deterministic)
}
func (m *IPObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPObject.Merge(m, src)
}
func (m *IPObject) XXX_Size() int {
	return xxx_messageInfo_IPObject.Size(m)
}
func (m *IPObject) XXX_DiscardUnknown() {
	xxx_messageInfo_IPObject.DiscardUnknown(m)
}

var xxx_messageInfo_IPObject proto.InternalMessageInfo

func (m *IPObject) GetVpcID() string {
	if m != nil {
		return m.VpcID
	}
	return ""
}

func (m *IPObject) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *IPObject) GetSubnetID() string {
	if m != nil {
		return m.SubnetID
	}
	return ""
}

func (m *IPObject) GetSubnetCidr() string {
	if m != nil {
		return m.SubnetCidr
	}
	return ""
}

func (m *IPObject) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IPObject) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *IPObject) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *IPObject) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *IPObject) GetWorkloadName() string {
	if m != nil {
		return m.WorkloadName
	}
	return ""
}

func (m *IPObject) GetWorkloadKind() string {
	if m != nil {
		return m.WorkloadKind
	}
	return ""
}

func (m *IPObject) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *IPObject) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *IPObject) GetEniID() string {
	if m != nil {
		return m.EniID
	}
	return ""
}

func (m *IPObject) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *IPObject) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *IPObject) GetIsFixed() bool {
	if m != nil {
		return m.IsFixed
	}
	return false
}

type IPClaim struct {
	VpcID                string   `protobuf:"bytes,1,opt,name=vpcID,proto3" json:"vpcID,omitempty"`
	Region               string   `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	SubnetID             string   `protobuf:"bytes,3,opt,name=subnetID,proto3" json:"subnetID,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Cluster              string   `protobuf:"bytes,5,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace            string   `protobuf:"bytes,6,opt,name=namespace,proto3" json:"namespace,omitempty"`
	PodName              string   `protobuf:"bytes,7,opt,name=podName,proto3" json:"podName,omitempty"`
	WorkloadName         string   `protobuf:"bytes,8,opt,name=workloadName,proto3" json:"workloadName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPClaim) Reset()         { *m = IPClaim{} }
func (m *IPClaim) String() string { return proto.CompactTextString(m) }
func (*IPClaim) ProtoMessage()    {}
func (*IPClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *IPClaim) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPClaim.Unmarshal(m, b)
}
func (m *IPClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPClaim.Marshal(b, m, deterministic)
}
func (m *IPClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPClaim.Merge(m, src)
}
func (m *IPClaim) XXX_Size() int {
	return xxx_messageInfo_IPClaim.Size(m)
}
func (m *IPClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_IPClaim.DiscardUnknown(m)
}

var xxx_messageInfo_IPClaim proto.InternalMessageInfo

func (m *IPClaim) GetVpcID() string {
	if m != nil {
		return m.VpcID
	}
	return ""
}

func (m *IPClaim) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *IPClaim) GetSubnetID() string {
	if m != nil {
		return m.SubnetID
	}
	return ""
}

func (m *IPClaim) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IPClaim) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *IPClaim) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *IPClaim) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *IPClaim) GetWorkloadName() string {
	if m != nil {
		return m.WorkloadName
	}
	return ""
}

func init() {
	proto.RegisterEnum("common.ErrCode", ErrCode_name, ErrCode_value)
	proto.RegisterType((*CloudSubnet)(nil), "common.CloudSubnet")
	proto.RegisterType((*IPObject)(nil), "common.IPObject")
	proto.RegisterType((*IPClaim)(nil), "common.IPClaim")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 885 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4b, 0x6f, 0xe3, 0x54,
	0x14, 0x26, 0x53, 0x3b, 0x8f, 0xd3, 0x32, 0x58, 0x57, 0x08, 0x59, 0x08, 0xa1, 0x28, 0xc0, 0x90,
	0x79, 0x75, 0x98, 0x29, 0x8f, 0x81, 0x05, 0xe2, 0x8e, 0x7d, 0x53, 0xae, 0xea, 0xda, 0xe6, 0xda,
	0x29, 0xcc, 0xea, 0xca, 0x8d, 0x2d, 0x30, 0x24, 0x76, 0xe4, 0x24, 0x03, 0x62, 0x95, 0x2d, 0x7f,
	0x0e, 0xf1, 0x86, 0x3d, 0x1b, 0xde, 0x6f, 0x90, 0x37, 0x48, 0xe8, 0xda, 0x71, 0xe3, 0x66, 0x12,
	0xb7, 0x48, 0xb3, 0xf3, 0xf9, 0xce, 0xf7, 0x9d, 0xd7, 0x17, 0x5b, 0x81, 0x9d, 0x41, 0x3c, 0x1a,
	0xc5, 0xd1, 0xee, 0x38, 0x89, 0xa7, 0x31, 0xaa, 0xe7, 0x51, 0xe7, 0xa3, 0x0b, 0xb0, 0xad, 0x0d,
	0xe3, 0x99, 0xef, 0xcc, 0x8e, 0xa3, 0x60, 0x8a, 0x1e, 0x05, 0xf9, 0xde, 0x78, 0x40, 0x75, 0xb5,
	0xd6, 0xae, 0x75, 0x5b, 0x2c, 0x0f, 0xd0, 0x63, 0x50, 0x4f, 0x82, 0xb7, 0xc3, 0x38, 0x52, 0x2f,
	0x64, 0xf0, 0x22, 0x42, 0x08, 0xa4, 0x0f, 0xe3, 0x28, 0x50, 0xb7, 0x32, 0x34, 0x7b, 0x46, 0x8f,
	0x43, 0x73, 0x92, 0xd5, 0xa2, 0xba, 0x2a, 0x65, 0xf8, 0x49, 0x8c, 0x9e, 0x04, 0xc8, 0x9f, 0xb5,
	0xd0, 0x4f, 0x54, 0x39, 0xcb, 0x96, 0x10, 0x74, 0x09, 0x2e, 0x7a, 0xf7, 0xbc, 0x70, 0xe8, 0x1d,
	0x0f, 0x03, 0x6a, 0x9b, 0xb3, 0x91, 0x5a, 0x6f, 0xd7, 0xba, 0x12, 0x5b, 0x41, 0xc5, 0x94, 0x93,
	0xa9, 0x37, 0x0d, 0xd4, 0x46, 0xbb, 0xd6, 0x95, 0x59, 0x1e, 0x88, 0xea, 0x83, 0x24, 0xf0, 0xa6,
	0x81, 0x1b, 0x8e, 0x02, 0xb5, 0x99, 0x57, 0x5f, 0x22, 0x22, 0x3f, 0x1b, 0xfb, 0x45, 0xbe, 0x95,
	0xe7, 0x97, 0x48, 0xe7, 0xe3, 0x2d, 0x68, 0x52, 0xdb, 0x3a, 0x7e, 0x37, 0x18, 0xfc, 0xdf, 0x43,
	0x94, 0x97, 0xde, 0xaa, 0x5c, 0x5a, 0xba, 0x6f, 0x69, 0x15, 0x1a, 0x9e, 0xef, 0x27, 0xc1, 0x64,
	0xb2, 0xb8, 0x48, 0x11, 0x8a, 0xcc, 0x60, 0x38, 0x9b, 0x4c, 0x83, 0x24, 0xbb, 0x43, 0x8b, 0x15,
	0x21, 0x7a, 0x02, 0x5a, 0x91, 0x37, 0x0a, 0x26, 0x63, 0x6f, 0x90, 0x1f, 0xa1, 0xc5, 0x96, 0x80,
	0xd0, 0x8d, 0x63, 0xdf, 0xf4, 0x4e, 0xae, 0x50, 0x84, 0xa8, 0x03, 0x3b, 0xef, 0xc7, 0xc9, 0x7b,
	0xc3, 0xd8, 0xcb, 0xd3, 0xf9, 0x11, 0x4e, 0x61, 0x65, 0xce, 0x41, 0x18, 0xf9, 0x2a, 0x9c, 0xe6,
	0x08, 0x4c, 0x18, 0xff, 0x4e, 0x3c, 0x99, 0xaa, 0xdb, 0xb9, 0xf1, 0xe2, 0x19, 0xb5, 0x61, 0x7b,
	0x10, 0x47, 0x53, 0x2f, 0x8c, 0x82, 0x84, 0xea, 0xea, 0x4e, 0x96, 0x2a, 0x43, 0xe2, 0xa6, 0x41,
	0x14, 0x52, 0x5d, 0x7d, 0x38, 0xbf, 0x69, 0x16, 0xac, 0xd8, 0x76, 0xf1, 0x0c, 0xdb, 0x1e, 0x59,
	0xb5, 0x4d, 0x6c, 0x1b, 0x4e, 0x7a, 0xe1, 0x07, 0x81, 0xaf, 0x2a, 0xed, 0x5a, 0xb7, 0xc9, 0x8a,
	0xb0, 0xf3, 0x6d, 0x0d, 0x1a, 0xd4, 0xd6, 0x86, 0x5e, 0x38, 0x7a, 0x80, 0x7e, 0x96, 0xfc, 0x92,
	0x36, 0xfa, 0x25, 0x57, 0xf8, 0x55, 0xaf, 0xf0, 0xab, 0x51, 0xed, 0x57, 0xf3, 0x7e, 0xbf, 0xae,
	0xfc, 0xdb, 0x82, 0x06, 0x49, 0x12, 0x2d, 0xf6, 0x03, 0xb4, 0x03, 0x4d, 0xc2, 0x98, 0xc5, 0xb8,
	0x75, 0xa0, 0x3c, 0x84, 0xae, 0xc2, 0x33, 0x79, 0xa4, 0x19, 0x56, 0x5f, 0xe7, 0x26, 0x71, 0x1d,
	0xc2, 0x8e, 0xa8, 0x46, 0x38, 0x35, 0x8f, 0xb0, 0x41, 0x75, 0x6e, 0x63, 0x86, 0x0f, 0x1d, 0xe5,
	0xbb, 0x54, 0x42, 0xd7, 0xe0, 0xd2, 0x06, 0xb2, 0xe3, 0x5a, 0x8c, 0x58, 0xb6, 0xc3, 0x7b, 0x98,
	0x1a, 0x44, 0x57, 0xbe, 0xaf, 0x64, 0x67, 0x00, 0xb6, 0x69, 0xc1, 0xfe, 0x21, 0x95, 0xd0, 0x0b,
	0x70, 0xe3, 0x2c, 0x36, 0x76, 0x1c, 0xba, 0x6f, 0x52, 0xbb, 0x90, 0xfd, 0x98, 0x4a, 0xe8, 0x25,
	0xb8, 0x79, 0x96, 0xac, 0x6f, 0xae, 0x0a, 0x7f, 0x4a, 0x25, 0xf4, 0x22, 0x3c, 0x77, 0x96, 0xf0,
	0x8d, 0x3e, 0x61, 0x77, 0x39, 0x31, 0x4f, 0xe6, 0xfc, 0xf9, 0x7c, 0x0d, 0x85, 0x82, 0x9a, 0x3d,
	0x8b, 0x9b, 0x96, 0x7b, 0x88, 0x5d, 0xed, 0x75, 0xe5, 0x97, 0x54, 0x42, 0x3a, 0xbc, 0x7a, 0xbe,
	0x86, 0x4e, 0xff, 0x8e, 0x49, 0x5c, 0xde, 0x63, 0xd6, 0x61, 0x7e, 0xd7, 0xa2, 0xfd, 0xaf, 0xa9,
	0x84, 0x6e, 0xc3, 0xad, 0x0d, 0x55, 0x5c, 0x76, 0x97, 0xbb, 0x16, 0xd7, 0x89, 0x41, 0x5c, 0xc2,
	0xb1, 0xe6, 0xd2, 0x23, 0xb2, 0xa8, 0xa5, 0xfc, 0x96, 0x4a, 0xe8, 0x65, 0xd8, 0x3b, 0x97, 0x92,
	0x98, 0xf8, 0x8e, 0x41, 0xf4, 0x42, 0xfa, 0x7b, 0x2a, 0xa1, 0x1b, 0x70, 0x79, 0x93, 0xef, 0xf9,
	0xac, 0xd4, 0xe1, 0x3a, 0x75, 0x32, 0x9d, 0xf2, 0x47, 0xe5, 0x71, 0x17, 0xbd, 0xb0, 0x61, 0x58,
	0x1a, 0x5e, 0xce, 0x49, 0x6d, 0xe5, 0xcf, 0x54, 0x42, 0xbb, 0xd0, 0xdd, 0xa0, 0x3b, 0xa4, 0xfb,
	0x4c, 0xf0, 0x97, 0x26, 0xfe, 0x95, 0x4a, 0xe8, 0x16, 0x5c, 0xdb, 0xc0, 0x3f, 0x69, 0x40, 0x6d,
	0x61, 0x03, 0xcf, 0x7d, 0xf8, 0x3b, 0x95, 0xd0, 0x1e, 0x5c, 0xaf, 0x9e, 0x4d, 0x33, 0x08, 0x36,
	0x4b, 0x83, 0xfd, 0x93, 0x4a, 0xa8, 0x0b, 0x9d, 0x15, 0x11, 0xde, 0x27, 0xa6, 0xcb, 0x6d, 0x4b,
	0xcf, 0xca, 0xf7, 0xac, 0xbe, 0xa9, 0x2b, 0x9f, 0xcc, 0x65, 0x74, 0x1d, 0x9e, 0x5d, 0xcb, 0x2c,
	0x0f, 0xb4, 0xd8, 0xe0, 0xd3, 0xb9, 0xbc, 0xe6, 0x25, 0xc9, 0xe9, 0x8c, 0x18, 0x04, 0x3b, 0x65,
	0xf6, 0x67, 0x73, 0x19, 0x3d, 0x0f, 0xbb, 0x6b, 0xd9, 0x07, 0xb7, 0x1d, 0x2e, 0x7e, 0x40, 0x62,
	0x13, 0xc2, 0x78, 0xe9, 0x45, 0xfc, 0xbc, 0x42, 0x65, 0x5a, 0x3a, 0x31, 0x89, 0xfb, 0xa6, 0xc5,
	0x0e, 0xb2, 0x25, 0xf0, 0x11, 0xa6, 0x86, 0x30, 0x51, 0xf9, 0x62, 0x2e, 0xa3, 0x9b, 0x70, 0x75,
	0xe3, 0xca, 0x42, 0x62, 0x58, 0xb8, 0xbc, 0xfb, 0x97, 0x73, 0x19, 0x5d, 0x81, 0xa7, 0xd7, 0x4a,
	0x8a, 0x4f, 0x09, 0xb5, 0xb3, 0xf7, 0x42, 0xf9, 0x6a, 0x2e, 0xa3, 0xcb, 0xf0, 0x54, 0x25, 0x77,
	0xf1, 0xd9, 0xf9, 0x7a, 0x2e, 0xa3, 0xd7, 0xe0, 0x95, 0xb5, 0xd4, 0x53, 0x53, 0x38, 0x7d, 0xdb,
	0xb6, 0x98, 0xcb, 0x7b, 0xf4, 0x2d, 0x92, 0x75, 0xea, 0x11, 0xec, 0xf6, 0x19, 0x51, 0xbe, 0x99,
	0xcb, 0xc7, 0xf5, 0xec, 0x1f, 0xcd, 0xde, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xa0, 0xa7,
	0x6e, 0xe1, 0x08, 0x00, 0x00,
}
