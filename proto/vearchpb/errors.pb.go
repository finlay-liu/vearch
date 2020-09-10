// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errors.proto

package vearchpb

import (
	bytes "bytes"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

type ErrorEnum int32

const (
	ErrorEnum_SUCCESS                              ErrorEnum = 0
	ErrorEnum_INTERNAL_ERROR                       ErrorEnum = 1
	ErrorEnum_NAME_OR_PASSWORD                     ErrorEnum = 2
	ErrorEnum_SYSBUSY                              ErrorEnum = 3
	ErrorEnum_PARAM_ERROR                          ErrorEnum = 4
	ErrorEnum_INVALID_CFG                          ErrorEnum = 5
	ErrorEnum_TIMEOUT                              ErrorEnum = 6
	ErrorEnum_SERVICE_UNAVAILABLE                  ErrorEnum = 7
	ErrorEnum_ZONE_NOT_EXISTS                      ErrorEnum = 8
	ErrorEnum_LOCAL_ZONE_OPS_FAILED                ErrorEnum = 9
	ErrorEnum_DUP_ZONE                             ErrorEnum = 10
	ErrorEnum_DUP_DB                               ErrorEnum = 11
	ErrorEnum_INVALID_ENGINE                       ErrorEnum = 12
	ErrorEnum_DB_NOTEXISTS                         ErrorEnum = 13
	ErrorEnum_DB_Not_Empty                         ErrorEnum = 14
	ErrorEnum_DUP_SPACE                            ErrorEnum = 15
	ErrorEnum_SPACE_NOTEXISTS                      ErrorEnum = 16
	ErrorEnum_PARTITION_HAS_TASK_NOW               ErrorEnum = 17
	ErrorEnum_REPLICA_NOT_EXISTS                   ErrorEnum = 18
	ErrorEnum_DUP_REPLICA                          ErrorEnum = 19
	ErrorEnum_PARTITION_REPLICA_LEADER_NOT_DELETE  ErrorEnum = 20
	ErrorEnum_PS_NOTEXISTS                         ErrorEnum = 21
	ErrorEnum_PS_Already_Exists                    ErrorEnum = 22
	ErrorEnum_LOCAL_SPACE_OPS_FAILED               ErrorEnum = 23
	ErrorEnum_Local_PS_Ops_Failed                  ErrorEnum = 24
	ErrorEnum_GENID_FAILED                         ErrorEnum = 25
	ErrorEnum_LOCALDB_OPTFAILED                    ErrorEnum = 26
	ErrorEnum_SPACE_SCHEMA_INVALID                 ErrorEnum = 27
	ErrorEnum_RPC_GET_CLIENT_FAILED                ErrorEnum = 28
	ErrorEnum_RPC_INVALID_RESP                     ErrorEnum = 29
	ErrorEnum_RPC_INVOKE_FAILED                    ErrorEnum = 30
	ErrorEnum_RPC_PARAM_ERROR                      ErrorEnum = 31
	ErrorEnum_METHOD_NOT_IMPLEMENT                 ErrorEnum = 32
	ErrorEnum_USER_NOT_EXISTS                      ErrorEnum = 33
	ErrorEnum_DUP_USER                             ErrorEnum = 34
	ErrorEnum_USER_OPS_FAILED                      ErrorEnum = 35
	ErrorEnum_AUTHENTICATION_FAILED                ErrorEnum = 36
	ErrorEnum_REGION_NOT_EXISTS                    ErrorEnum = 37
	ErrorEnum_MASTER_PS_CAN_NOT_SELECT             ErrorEnum = 38
	ErrorEnum_MASTER_PS_NOT_ENOUGH_SELECT          ErrorEnum = 39
	ErrorEnum_PARTITION_DUPLICATE                  ErrorEnum = 40
	ErrorEnum_PARTITION_NOT_EXIST                  ErrorEnum = 41
	ErrorEnum_PARTITION_NOT_LEADER                 ErrorEnum = 42
	ErrorEnum_PARTITION_NO_LEADER                  ErrorEnum = 43
	ErrorEnum_PARTITION_REQ_PARAM                  ErrorEnum = 44
	ErrorEnum_PARTITON_ENGINENAME_INVALID          ErrorEnum = 45
	ErrorEnum_UNKNOWN_PARTITION_RAFT_CMD_TYPE      ErrorEnum = 46
	ErrorEnum_MASTER_SERVER_IS_NOT_RUNNING         ErrorEnum = 47
	ErrorEnum_PARTITION_IS_INVALID                 ErrorEnum = 48
	ErrorEnum_PARTITION_IS_CLOSED                  ErrorEnum = 49
	ErrorEnum_DOCUMENT_NOT_EXIST                   ErrorEnum = 50
	ErrorEnum_DOCUMENT_EXIST                       ErrorEnum = 51
	ErrorEnum_DOCUMENT_MUST_HAS_SOURCE             ErrorEnum = 52
	ErrorEnum_PULL_OUT_VERSION_NOT_MATCH           ErrorEnum = 53
	ErrorEnum_FUNC_CAN_NOT_INVOKE_IN_FROZEN_ENGINE ErrorEnum = 54
	ErrorEnum_ROUTER_NO_PS_CLIENT                  ErrorEnum = 55
	ErrorEnum_ROUTER_CALL_PS_RPC_ERR               ErrorEnum = 56
	ErrorEnum_GAMMA_SEARCH_QUERY_NUM_LESS_0        ErrorEnum = 57
	ErrorEnum_GAMMA_SEARCH_NO_CREATE_INDEX         ErrorEnum = 58
	ErrorEnum_GAMMA_SEARCH_INDEX_QUERY_ERR         ErrorEnum = 59
	ErrorEnum_GAMMA_SEARCH_OTHER_ERR               ErrorEnum = 60
	ErrorEnum_Primary_IS_INVALID                   ErrorEnum = 61
	ErrorEnum_PARSING_RESULT_ERROR                 ErrorEnum = 62
	ErrorEnum_FORCE_MERGE_BUILD_INDEX_ERR          ErrorEnum = 63
	ErrorEnum_DELETE_BY_QUERY_SERACH_ERR           ErrorEnum = 64
	ErrorEnum_DELETE_BY_QUERY_SEARCH_ID_IS_0       ErrorEnum = 65
	ErrorEnum_FLUSH_ERR                            ErrorEnum = 66
	ErrorEnum_Create_RpcClient_Failed              ErrorEnum = 70
	ErrorEnum_Call_RpcClient_Failed                ErrorEnum = 71
	ErrorEnum_RECOVER                              ErrorEnum = 100
)

var ErrorEnum_name = map[int32]string{
	0:   "SUCCESS",
	1:   "INTERNAL_ERROR",
	2:   "NAME_OR_PASSWORD",
	3:   "SYSBUSY",
	4:   "PARAM_ERROR",
	5:   "INVALID_CFG",
	6:   "TIMEOUT",
	7:   "SERVICE_UNAVAILABLE",
	8:   "ZONE_NOT_EXISTS",
	9:   "LOCAL_ZONE_OPS_FAILED",
	10:  "DUP_ZONE",
	11:  "DUP_DB",
	12:  "INVALID_ENGINE",
	13:  "DB_NOTEXISTS",
	14:  "DB_Not_Empty",
	15:  "DUP_SPACE",
	16:  "SPACE_NOTEXISTS",
	17:  "PARTITION_HAS_TASK_NOW",
	18:  "REPLICA_NOT_EXISTS",
	19:  "DUP_REPLICA",
	20:  "PARTITION_REPLICA_LEADER_NOT_DELETE",
	21:  "PS_NOTEXISTS",
	22:  "PS_Already_Exists",
	23:  "LOCAL_SPACE_OPS_FAILED",
	24:  "Local_PS_Ops_Failed",
	25:  "GENID_FAILED",
	26:  "LOCALDB_OPTFAILED",
	27:  "SPACE_SCHEMA_INVALID",
	28:  "RPC_GET_CLIENT_FAILED",
	29:  "RPC_INVALID_RESP",
	30:  "RPC_INVOKE_FAILED",
	31:  "RPC_PARAM_ERROR",
	32:  "METHOD_NOT_IMPLEMENT",
	33:  "USER_NOT_EXISTS",
	34:  "DUP_USER",
	35:  "USER_OPS_FAILED",
	36:  "AUTHENTICATION_FAILED",
	37:  "REGION_NOT_EXISTS",
	38:  "MASTER_PS_CAN_NOT_SELECT",
	39:  "MASTER_PS_NOT_ENOUGH_SELECT",
	40:  "PARTITION_DUPLICATE",
	41:  "PARTITION_NOT_EXIST",
	42:  "PARTITION_NOT_LEADER",
	43:  "PARTITION_NO_LEADER",
	44:  "PARTITION_REQ_PARAM",
	45:  "PARTITON_ENGINENAME_INVALID",
	46:  "UNKNOWN_PARTITION_RAFT_CMD_TYPE",
	47:  "MASTER_SERVER_IS_NOT_RUNNING",
	48:  "PARTITION_IS_INVALID",
	49:  "PARTITION_IS_CLOSED",
	50:  "DOCUMENT_NOT_EXIST",
	51:  "DOCUMENT_EXIST",
	52:  "DOCUMENT_MUST_HAS_SOURCE",
	53:  "PULL_OUT_VERSION_NOT_MATCH",
	54:  "FUNC_CAN_NOT_INVOKE_IN_FROZEN_ENGINE",
	55:  "ROUTER_NO_PS_CLIENT",
	56:  "ROUTER_CALL_PS_RPC_ERR",
	57:  "GAMMA_SEARCH_QUERY_NUM_LESS_0",
	58:  "GAMMA_SEARCH_NO_CREATE_INDEX",
	59:  "GAMMA_SEARCH_INDEX_QUERY_ERR",
	60:  "GAMMA_SEARCH_OTHER_ERR",
	61:  "Primary_IS_INVALID",
	62:  "PARSING_RESULT_ERROR",
	63:  "FORCE_MERGE_BUILD_INDEX_ERR",
	64:  "DELETE_BY_QUERY_SERACH_ERR",
	65:  "DELETE_BY_QUERY_SEARCH_ID_IS_0",
	66:  "FLUSH_ERR",
	70:  "Create_RpcClient_Failed",
	71:  "Call_RpcClient_Failed",
	100: "RECOVER",
}

var ErrorEnum_value = map[string]int32{
	"SUCCESS":                              0,
	"INTERNAL_ERROR":                       1,
	"NAME_OR_PASSWORD":                     2,
	"SYSBUSY":                              3,
	"PARAM_ERROR":                          4,
	"INVALID_CFG":                          5,
	"TIMEOUT":                              6,
	"SERVICE_UNAVAILABLE":                  7,
	"ZONE_NOT_EXISTS":                      8,
	"LOCAL_ZONE_OPS_FAILED":                9,
	"DUP_ZONE":                             10,
	"DUP_DB":                               11,
	"INVALID_ENGINE":                       12,
	"DB_NOTEXISTS":                         13,
	"DB_Not_Empty":                         14,
	"DUP_SPACE":                            15,
	"SPACE_NOTEXISTS":                      16,
	"PARTITION_HAS_TASK_NOW":               17,
	"REPLICA_NOT_EXISTS":                   18,
	"DUP_REPLICA":                          19,
	"PARTITION_REPLICA_LEADER_NOT_DELETE":  20,
	"PS_NOTEXISTS":                         21,
	"PS_Already_Exists":                    22,
	"LOCAL_SPACE_OPS_FAILED":               23,
	"Local_PS_Ops_Failed":                  24,
	"GENID_FAILED":                         25,
	"LOCALDB_OPTFAILED":                    26,
	"SPACE_SCHEMA_INVALID":                 27,
	"RPC_GET_CLIENT_FAILED":                28,
	"RPC_INVALID_RESP":                     29,
	"RPC_INVOKE_FAILED":                    30,
	"RPC_PARAM_ERROR":                      31,
	"METHOD_NOT_IMPLEMENT":                 32,
	"USER_NOT_EXISTS":                      33,
	"DUP_USER":                             34,
	"USER_OPS_FAILED":                      35,
	"AUTHENTICATION_FAILED":                36,
	"REGION_NOT_EXISTS":                    37,
	"MASTER_PS_CAN_NOT_SELECT":             38,
	"MASTER_PS_NOT_ENOUGH_SELECT":          39,
	"PARTITION_DUPLICATE":                  40,
	"PARTITION_NOT_EXIST":                  41,
	"PARTITION_NOT_LEADER":                 42,
	"PARTITION_NO_LEADER":                  43,
	"PARTITION_REQ_PARAM":                  44,
	"PARTITON_ENGINENAME_INVALID":          45,
	"UNKNOWN_PARTITION_RAFT_CMD_TYPE":      46,
	"MASTER_SERVER_IS_NOT_RUNNING":         47,
	"PARTITION_IS_INVALID":                 48,
	"PARTITION_IS_CLOSED":                  49,
	"DOCUMENT_NOT_EXIST":                   50,
	"DOCUMENT_EXIST":                       51,
	"DOCUMENT_MUST_HAS_SOURCE":             52,
	"PULL_OUT_VERSION_NOT_MATCH":           53,
	"FUNC_CAN_NOT_INVOKE_IN_FROZEN_ENGINE": 54,
	"ROUTER_NO_PS_CLIENT":                  55,
	"ROUTER_CALL_PS_RPC_ERR":               56,
	"GAMMA_SEARCH_QUERY_NUM_LESS_0":        57,
	"GAMMA_SEARCH_NO_CREATE_INDEX":         58,
	"GAMMA_SEARCH_INDEX_QUERY_ERR":         59,
	"GAMMA_SEARCH_OTHER_ERR":               60,
	"Primary_IS_INVALID":                   61,
	"PARSING_RESULT_ERROR":                 62,
	"FORCE_MERGE_BUILD_INDEX_ERR":          63,
	"DELETE_BY_QUERY_SERACH_ERR":           64,
	"DELETE_BY_QUERY_SEARCH_ID_IS_0":       65,
	"FLUSH_ERR":                            66,
	"Create_RpcClient_Failed":              70,
	"Call_RpcClient_Failed":                71,
	"RECOVER":                              100,
}

func (x ErrorEnum) String() string {
	return proto.EnumName(ErrorEnum_name, int32(x))
}

func (ErrorEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{0}
}

type SearchResultCode int32

const (
	SearchResultCode_SEARCH_SUCCESS    SearchResultCode = 0
	SearchResultCode_INDEX_NOT_BUILDED SearchResultCode = 1
	SearchResultCode_SEARCH_ERROR      SearchResultCode = 2
)

var SearchResultCode_name = map[int32]string{
	0: "SEARCH_SUCCESS",
	1: "INDEX_NOT_BUILDED",
	2: "SEARCH_ERROR",
}

var SearchResultCode_value = map[string]int32{
	"SEARCH_SUCCESS":    0,
	"INDEX_NOT_BUILDED": 1,
	"SEARCH_ERROR":      2,
}

func (x SearchResultCode) String() string {
	return proto.EnumName(SearchResultCode_name, int32(x))
}

func (SearchResultCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{1}
}

type Error struct {
	Code                 ErrorEnum `protobuf:"varint,1,opt,name=code,proto3,enum=ErrorEnum" json:"code,omitempty"`
	Msg                  string    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Error) Reset()      { *m = Error{} }
func (*Error) ProtoMessage() {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{0}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Error.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return m.Size()
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("ErrorEnum", ErrorEnum_name, ErrorEnum_value)
	proto.RegisterEnum("SearchResultCode", SearchResultCode_name, SearchResultCode_value)
	proto.RegisterType((*Error)(nil), "Error")
}

func init() { proto.RegisterFile("errors.proto", fileDescriptor_24fe73c7f0ddb19c) }

var fileDescriptor_24fe73c7f0ddb19c = []byte{
	// 1131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x55, 0xcb, 0x76, 0x13, 0x47,
	0x13, 0xd6, 0x70, 0x77, 0x63, 0x4c, 0xd3, 0x60, 0x30, 0x06, 0x06, 0x73, 0xf9, 0x7f, 0x1c, 0x12,
	0x0c, 0x81, 0xdc, 0xc8, 0x95, 0x56, 0x4f, 0x49, 0x9a, 0xc3, 0x4c, 0xf7, 0xd0, 0xdd, 0x63, 0x30,
	0x9b, 0x3e, 0xb2, 0xad, 0x18, 0x9f, 0x23, 0x23, 0x1f, 0x49, 0xce, 0x89, 0x77, 0x79, 0x8c, 0x3c,
	0x42, 0x1e, 0x21, 0x8b, 0x2c, 0xb2, 0x64, 0x99, 0x65, 0x96, 0x58, 0x79, 0x81, 0x2c, 0xb3, 0xcc,
	0xa9, 0x9e, 0x19, 0x59, 0x8a, 0x77, 0x33, 0xf5, 0xd5, 0xe5, 0xab, 0xaf, 0xaa, 0x66, 0xc8, 0x6c,
	0xa7, 0xdf, 0xef, 0xf5, 0x07, 0x2b, 0xbb, 0xfd, 0xde, 0xb0, 0xb7, 0xf8, 0x60, 0x6b, 0x7b, 0xf8,
	0x66, 0x6f, 0x7d, 0x65, 0xa3, 0xb7, 0xf3, 0x70, 0xab, 0xb7, 0xd5, 0x7b, 0xe8, 0xcd, 0xeb, 0x7b,
	0xdf, 0xfb, 0x37, 0xff, 0xe2, 0x9f, 0x0a, 0xf7, 0xdb, 0x4f, 0xc9, 0x49, 0xc0, 0x70, 0x16, 0x92,
	0x13, 0x1b, 0xbd, 0xcd, 0xce, 0x42, 0xb0, 0x14, 0x2c, 0xcf, 0x3d, 0x26, 0x2b, 0xde, 0x0a, 0x6f,
	0xf7, 0x76, 0xb4, 0xb7, 0x33, 0x4a, 0x8e, 0xef, 0x0c, 0xb6, 0x16, 0x8e, 0x2d, 0x05, 0xcb, 0x33,
	0x1a, 0x1f, 0xef, 0xff, 0x76, 0x8e, 0xcc, 0x8c, 0xbd, 0xd8, 0x59, 0x72, 0xda, 0xe4, 0x42, 0x80,
	0x31, 0xb4, 0xc6, 0x18, 0x99, 0x8b, 0xa5, 0x05, 0x2d, 0x79, 0xe2, 0x40, 0x6b, 0xa5, 0x69, 0xc0,
	0x2e, 0x11, 0x2a, 0x79, 0x0a, 0x4e, 0x69, 0x97, 0x71, 0x63, 0x5e, 0x2a, 0x1d, 0xd1, 0x63, 0x3e,
	0x6c, 0xcd, 0xd4, 0x73, 0xb3, 0x46, 0x8f, 0xb3, 0xf3, 0xe4, 0x6c, 0xc6, 0x35, 0x4f, 0xcb, 0x98,
	0x13, 0x68, 0x88, 0xe5, 0x2a, 0x4f, 0xe2, 0xc8, 0x89, 0x46, 0x93, 0x9e, 0x44, 0x77, 0x1b, 0xa7,
	0xa0, 0x72, 0x4b, 0x4f, 0xb1, 0x2b, 0xe4, 0xa2, 0x01, 0xbd, 0x1a, 0x0b, 0x70, 0xb9, 0xe4, 0xab,
	0x3c, 0x4e, 0x78, 0x3d, 0x01, 0x7a, 0x9a, 0x5d, 0x24, 0xe7, 0x5f, 0x2b, 0x09, 0x4e, 0x2a, 0xeb,
	0xe0, 0x55, 0x6c, 0xac, 0xa1, 0x67, 0xd8, 0x55, 0x32, 0x9f, 0x28, 0xc1, 0x13, 0xe7, 0x21, 0x95,
	0x19, 0xd7, 0xe0, 0x71, 0x02, 0x11, 0x9d, 0x61, 0xb3, 0xe4, 0x4c, 0x94, 0x67, 0x1e, 0xa0, 0x84,
	0x11, 0x72, 0x0a, 0xdf, 0xa2, 0x3a, 0x3d, 0x5b, 0x34, 0x52, 0x10, 0x00, 0xd9, 0x8c, 0x25, 0xd0,
	0x59, 0x46, 0xc9, 0x6c, 0x54, 0xc7, 0xdc, 0x65, 0xea, 0x73, 0x95, 0xa5, 0x37, 0x74, 0xb0, 0xb3,
	0x3b, 0xdc, 0xa7, 0x73, 0xec, 0x1c, 0x99, 0xc1, 0x1c, 0x26, 0xe3, 0x02, 0xe8, 0x79, 0x24, 0xe4,
	0x1f, 0x27, 0xa2, 0x28, 0x5b, 0x24, 0x97, 0x33, 0xae, 0x6d, 0x6c, 0x63, 0x25, 0x5d, 0x8b, 0x1b,
	0x67, 0xb9, 0x79, 0xee, 0xa4, 0x7a, 0x49, 0x2f, 0xb0, 0xcb, 0x84, 0x69, 0xc8, 0x92, 0x58, 0xf0,
	0xc9, 0x26, 0x18, 0x0a, 0x82, 0x79, 0x4b, 0x8c, 0x5e, 0x64, 0xf7, 0xc8, 0x9d, 0xc3, 0x24, 0x55,
	0x48, 0x02, 0x3c, 0x02, 0xed, 0x23, 0x23, 0x48, 0xc0, 0x02, 0xbd, 0x84, 0x1c, 0x33, 0x33, 0x51,
	0x7f, 0x9e, 0xcd, 0x93, 0x0b, 0x99, 0x71, 0xbc, 0xdb, 0xef, 0xb4, 0x37, 0xf7, 0x1d, 0xfc, 0xb8,
	0x3d, 0x18, 0x0e, 0xe8, 0x65, 0xa4, 0x55, 0xe8, 0x54, 0x30, 0x9e, 0x10, 0xea, 0x0a, 0x2a, 0x9e,
	0xf4, 0x36, 0xda, 0x5d, 0x97, 0x19, 0xa7, 0x76, 0x07, 0xae, 0xd1, 0xde, 0xee, 0x76, 0x36, 0xe9,
	0x02, 0x66, 0x6f, 0x82, 0x8c, 0xa3, 0xca, 0xf5, 0x2a, 0x66, 0xf7, 0x69, 0xa2, 0xba, 0x53, 0x99,
	0x2d, 0xcd, 0x8b, 0x6c, 0x81, 0x5c, 0x2a, 0xf2, 0x1a, 0xd1, 0x82, 0x94, 0xbb, 0x52, 0x5d, 0x7a,
	0x0d, 0xe7, 0xa3, 0x33, 0xe1, 0x9a, 0x60, 0x9d, 0x48, 0x62, 0x90, 0xb6, 0xca, 0x75, 0x1d, 0x57,
	0x07, 0xa1, 0x6a, 0x12, 0x1a, 0x4c, 0x46, 0x6f, 0x60, 0x85, 0xd2, 0xaa, 0x9e, 0x43, 0xe5, 0x1c,
	0xa2, 0xd6, 0x68, 0x9e, 0x5c, 0xa4, 0x9b, 0x58, 0x36, 0x05, 0xdb, 0x52, 0x91, 0x17, 0x25, 0x4e,
	0xb3, 0x04, 0x52, 0x90, 0x96, 0x2e, 0xa1, 0x7b, 0x6e, 0x4a, 0xb1, 0x4a, 0x69, 0x6e, 0x55, 0x0b,
	0x81, 0x00, 0xbd, 0x3d, 0x76, 0x99, 0x90, 0xe2, 0x0e, 0xd2, 0xe5, 0xb9, 0x6d, 0x81, 0xb4, 0xb1,
	0xe0, 0x5e, 0xfd, 0x12, 0xba, 0xeb, 0x89, 0x41, 0x13, 0x4d, 0x13, 0x49, 0xff, 0xc7, 0xae, 0x93,
	0x85, 0x94, 0x1b, 0x0b, 0x1a, 0xd5, 0x13, 0xbc, 0x40, 0x0d, 0x24, 0x20, 0x2c, 0xfd, 0x3f, 0xbb,
	0x49, 0xae, 0x1d, 0xa2, 0x3e, 0x4e, 0xaa, 0xbc, 0xd9, 0xaa, 0x1c, 0xee, 0xa1, 0xf6, 0x87, 0x93,
	0x8e, 0x72, 0x3f, 0x69, 0x0b, 0x74, 0x79, 0x1a, 0x18, 0x57, 0xa4, 0x1f, 0x60, 0xd3, 0xd3, 0x40,
	0xb1, 0x17, 0xf4, 0xfe, 0x7f, 0x43, 0x2a, 0xe0, 0xc3, 0x69, 0x40, 0xc3, 0x8b, 0x42, 0x46, 0xfa,
	0x11, 0xd2, 0x2b, 0x00, 0x25, 0xcb, 0x4b, 0xf0, 0xc7, 0x5c, 0x8d, 0xef, 0x01, 0xbb, 0x43, 0x6e,
	0xe6, 0xf2, 0xb9, 0x54, 0x2f, 0xa5, 0x9b, 0xc8, 0xc0, 0x1b, 0xd6, 0x89, 0x34, 0x72, 0x76, 0x2d,
	0x03, 0xba, 0xc2, 0x96, 0xc8, 0xf5, 0xb2, 0x49, 0x3c, 0x5c, 0xd0, 0x2e, 0x2e, 0x7a, 0xd5, 0xb9,
	0x94, 0xb1, 0x6c, 0xd2, 0x87, 0xd3, 0x9c, 0x63, 0x33, 0x2e, 0xf0, 0x68, 0x9a, 0x5a, 0x6c, 0x9c,
	0x48, 0x94, 0x81, 0x88, 0x7e, 0x8c, 0xb7, 0x12, 0x29, 0x91, 0xe3, 0x3c, 0x27, 0xda, 0x7f, 0x8c,
	0xb7, 0x3b, 0xb6, 0x17, 0xb6, 0x27, 0x38, 0x83, 0xb1, 0x2d, 0xcd, 0x8d, 0xf5, 0x77, 0x67, 0x54,
	0xae, 0x05, 0xd0, 0x4f, 0x58, 0x48, 0x16, 0xb3, 0x3c, 0x49, 0x9c, 0xca, 0xad, 0x5b, 0x05, 0x6d,
	0x2a, 0xdd, 0x52, 0x6e, 0x45, 0x8b, 0x7e, 0xca, 0x96, 0xc9, 0xdd, 0x46, 0x2e, 0xc5, 0x78, 0x78,
	0xe5, 0xea, 0xc5, 0xd2, 0x35, 0xb4, 0x7a, 0x0d, 0x95, 0x32, 0xf4, 0x33, 0x24, 0xab, 0x55, 0x6e,
	0xfd, 0x5e, 0xf9, 0x71, 0xfb, 0x8d, 0xa6, 0x9f, 0xe3, 0x75, 0x95, 0x80, 0xe0, 0x49, 0x82, 0x10,
	0x2e, 0x2b, 0x68, 0x4d, 0xbf, 0x60, 0xb7, 0xc8, 0x8d, 0x26, 0x4f, 0x53, 0xee, 0x0c, 0x70, 0x2d,
	0x5a, 0xee, 0x45, 0x0e, 0x7a, 0xcd, 0xc9, 0x3c, 0x75, 0x09, 0x18, 0xe3, 0x1e, 0xd1, 0xa7, 0x28,
	0xe0, 0x94, 0x8b, 0x54, 0x4e, 0x68, 0xe0, 0x16, 0x49, 0x44, 0xf0, 0x8a, 0x7e, 0x79, 0xc4, 0xc3,
	0xdb, 0xcb, 0x54, 0x58, 0xe6, 0x2b, 0xa4, 0x30, 0xe5, 0xa1, 0x6c, 0x0b, 0xb4, 0xc7, 0xbe, 0x46,
	0x2d, 0xb3, 0xfe, 0xf6, 0x4e, 0xbb, 0xbf, 0x3f, 0x29, 0xfe, 0x37, 0xe5, 0x58, 0x4c, 0x2c, 0x9b,
	0x78, 0x7d, 0x79, 0x62, 0xcb, 0xcb, 0xfa, 0x16, 0x17, 0xa3, 0xa1, 0xb4, 0x00, 0x97, 0x82, 0x6e,
	0x82, 0xab, 0xe7, 0x71, 0x12, 0x95, 0x45, 0x31, 0xe5, 0x77, 0x28, 0x6a, 0xf1, 0x11, 0x72, 0xf5,
	0xb5, 0x92, 0x87, 0x01, 0xcd, 0x45, 0xcb, 0xe3, 0xcf, 0xd8, 0x6d, 0x12, 0x1e, 0xc5, 0x0b, 0xea,
	0x11, 0x92, 0x78, 0x44, 0x39, 0x7e, 0x4e, 0x1b, 0x49, 0x6e, 0x8a, 0x90, 0x3a, 0xbb, 0x46, 0xae,
	0x88, 0x7e, 0xa7, 0x3d, 0xec, 0x38, 0xbd, 0xbb, 0x21, 0xba, 0xdb, 0x9d, 0xb7, 0xc3, 0xea, 0x53,
	0xd4, 0xc0, 0xc3, 0x14, 0xed, 0x6e, 0xf7, 0x28, 0xd4, 0xc4, 0xbf, 0x87, 0x06, 0xa1, 0x56, 0x41,
	0xd3, 0xcd, 0xfb, 0x8a, 0x50, 0xd3, 0x69, 0xf7, 0x37, 0xde, 0xe8, 0xce, 0x60, 0xaf, 0x3b, 0x14,
	0xf8, 0x93, 0x63, 0x64, 0xae, 0xac, 0x7d, 0xf8, 0x2f, 0x9b, 0x27, 0x17, 0x8a, 0x76, 0x70, 0xe2,
	0xbe, 0x3d, 0x88, 0x68, 0x80, 0x5f, 0xbc, 0xd2, 0xb5, 0x50, 0xe2, 0x58, 0xfd, 0xd9, 0xbb, 0x83,
	0xb0, 0xf6, 0xe7, 0x41, 0x58, 0x7b, 0x7f, 0x10, 0xd6, 0xfe, 0x3e, 0x08, 0x6b, 0xff, 0x1c, 0x84,
	0xc1, 0x4f, 0xa3, 0x30, 0xf8, 0x65, 0x14, 0x06, 0xbf, 0x8e, 0xc2, 0xda, 0xef, 0xa3, 0xb0, 0xf6,
	0x6e, 0x14, 0x06, 0x7f, 0x8c, 0xc2, 0xe0, 0xfd, 0x28, 0x0c, 0x7e, 0xfe, 0x2b, 0xac, 0xb5, 0x82,
	0xd7, 0x67, 0x7e, 0xf0, 0x34, 0x76, 0xd7, 0xd7, 0x4f, 0xf9, 0x7f, 0xf2, 0x93, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xe3, 0x3e, 0x92, 0xce, 0xd2, 0x07, 0x00, 0x00,
}

func (this *Error) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Error)
	if !ok {
		that2, ok := that.(Error)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	if this.Msg != that1.Msg {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Error) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintErrors(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintErrors(dAtA []byte, offset int, v uint64) int {
	offset -= sovErrors(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedError(r randyErrors, easy bool) *Error {
	this := &Error{}
	this.Code = ErrorEnum([]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 70, 71, 100}[r.Intn(70)])
	this.Msg = string(randStringErrors(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedErrors(r, 3)
	}
	return this
}

type randyErrors interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneErrors(r randyErrors) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringErrors(r randyErrors) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneErrors(r)
	}
	return string(tmps)
}
func randUnrecognizedErrors(r randyErrors, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldErrors(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldErrors(dAtA []byte, r randyErrors, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateErrors(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateErrors(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateErrors(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateErrors(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateErrors(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateErrors(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateErrors(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovErrors(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovErrors(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErrors(x uint64) (n int) {
	return sovErrors(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Error) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Error{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Msg:` + fmt.Sprintf("%v", this.Msg) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringErrors(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= ErrorEnum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipErrors(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrors
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthErrors
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErrors
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErrors
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErrors        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrors          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErrors = fmt.Errorf("proto: unexpected end of group")
)
