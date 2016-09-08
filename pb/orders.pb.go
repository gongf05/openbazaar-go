// Code generated by protoc-gen-go.
// source: orders.proto
// DO NOT EDIT!

/*
Package orders is a generated protocol buffer package.

It is generated from these files:
	orders.proto

It has these top-level messages:
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type OrderState int32

const (
	// New order, has not yet been accepted by the vendor
	OrderState_PENDING OrderState = 0
	// Vendor has accepted the order and is awaiting payment
	OrderState_CONFIRMED OrderState = 1
	// Buyer has funded the order
	OrderState_FUNDED OrderState = 2
	// Vendor has shipped the item/sent the digitial good
	OrderState_FULFILLED OrderState = 3
	// Buyer has completed the order, released the funds, and left a review
	OrderState_COMPLETE OrderState = 4
	// Contract is under active dispute
	OrderState_DISPUTED OrderState = 5
	// The moderator has resolved the dispute and we are waiting for the winning party to
	// accept the payout. After the payout is accepted the state will be set to COMPLETE
	OrderState_RESOLVED OrderState = 6
	// The vendor refunded the order
	OrderState_REFUNDED OrderState = 7
)

var OrderState_name = map[int32]string{
	0: "PENDING",
	1: "CONFIRMED",
	2: "FUNDED",
	3: "FULFILLED",
	4: "COMPLETE",
	5: "DISPUTED",
	6: "RESOLVED",
	7: "REFUNDED",
}
var OrderState_value = map[string]int32{
	"PENDING":   0,
	"CONFIRMED": 1,
	"FUNDED":    2,
	"FULFILLED": 3,
	"COMPLETE":  4,
	"DISPUTED":  5,
	"RESOLVED":  6,
	"REFUNDED":  7,
}

func (x OrderState) String() string {
	return proto.EnumName(OrderState_name, int32(x))
}
func (OrderState) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func init() {
	proto.RegisterEnum("OrderState", OrderState_name, OrderState_value)
}

var fileDescriptor6 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x2f, 0x4a, 0x49,
	0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xaa, 0xe6, 0xe2, 0xf2, 0x07, 0xf1, 0x83,
	0x4b, 0x12, 0x4b, 0x52, 0x85, 0xb8, 0xb9, 0xd8, 0x03, 0x5c, 0xfd, 0x5c, 0x3c, 0xfd, 0xdc, 0x05,
	0x18, 0x84, 0x78, 0xb9, 0x38, 0x9d, 0xfd, 0xfd, 0xdc, 0x3c, 0x83, 0x7c, 0x5d, 0x5d, 0x04, 0x18,
	0x85, 0xb8, 0xb8, 0xd8, 0xdc, 0x42, 0xfd, 0x5c, 0x80, 0x6c, 0x26, 0x90, 0x94, 0x5b, 0xa8, 0x8f,
	0x9b, 0xa7, 0x8f, 0x0f, 0x90, 0xcb, 0x2c, 0xc4, 0xc3, 0xc5, 0xe1, 0xec, 0xef, 0x1b, 0xe0, 0xe3,
	0x1a, 0xe2, 0x2a, 0xc0, 0x02, 0xe2, 0xb9, 0x78, 0x06, 0x07, 0x84, 0x86, 0x00, 0xe5, 0x58, 0x41,
	0xbc, 0x20, 0xd7, 0x60, 0x7f, 0x9f, 0x30, 0x20, 0x8f, 0x0d, 0xc2, 0x83, 0x1a, 0xc3, 0x9e, 0xc4,
	0x06, 0x76, 0x83, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xb6, 0x8c, 0x0d, 0x93, 0x00, 0x00,
	0x00,
}