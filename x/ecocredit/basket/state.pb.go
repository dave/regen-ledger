// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/basket/v1/state.proto

package basket

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Basket represents a basket in state.
type Basket struct {
	// id is the uint64 ID of the basket. It is used internally for reducing
	// storage space.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// basket_denom is the basket bank denom.
	BasketDenom string `protobuf:"bytes,2,opt,name=basket_denom,json=basketDenom,proto3" json:"basket_denom,omitempty"`
	// disable_auto_retire indicates whether or not the credits will be retired upon withdraw from the basket.
	DisableAutoRetire bool `protobuf:"varint,3,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// credit_type_name is the name of the credit type this basket is able to hold.
	CreditTypeName string `protobuf:"bytes,4,opt,name=credit_type_name,json=creditTypeName,proto3" json:"credit_type_name,omitempty"`
	// date_criteria is the date criteria for batches admitted to the basket.
	DateCriteria *DateCriteria `protobuf:"bytes,5,opt,name=date_criteria,json=dateCriteria,proto3" json:"date_criteria,omitempty"`
	// exponent is the exponent for converting credits to/from basket tokens.
	Exponent uint32 `protobuf:"varint,6,opt,name=exponent,proto3" json:"exponent,omitempty"`
}

func (m *Basket) Reset()         { *m = Basket{} }
func (m *Basket) String() string { return proto.CompactTextString(m) }
func (*Basket) ProtoMessage()    {}
func (*Basket) Descriptor() ([]byte, []int) {
	return fileDescriptor_c416a19075224f85, []int{0}
}
func (m *Basket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Basket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Basket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Basket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Basket.Merge(m, src)
}
func (m *Basket) XXX_Size() int {
	return m.Size()
}
func (m *Basket) XXX_DiscardUnknown() {
	xxx_messageInfo_Basket.DiscardUnknown(m)
}

var xxx_messageInfo_Basket proto.InternalMessageInfo

func (m *Basket) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Basket) GetBasketDenom() string {
	if m != nil {
		return m.BasketDenom
	}
	return ""
}

func (m *Basket) GetDisableAutoRetire() bool {
	if m != nil {
		return m.DisableAutoRetire
	}
	return false
}

func (m *Basket) GetCreditTypeName() string {
	if m != nil {
		return m.CreditTypeName
	}
	return ""
}

func (m *Basket) GetDateCriteria() *DateCriteria {
	if m != nil {
		return m.DateCriteria
	}
	return nil
}

func (m *Basket) GetExponent() uint32 {
	if m != nil {
		return m.Exponent
	}
	return 0
}

// BasketClass describes a credit class that can be deposited in a basket.
type BasketClass struct {
	// basket_id is the ID of the basket
	BasketId uint64 `protobuf:"varint,1,opt,name=basket_id,json=basketId,proto3" json:"basket_id,omitempty"`
	// class_id is the id of the credit class that is allowed to be deposited in the basket
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (m *BasketClass) Reset()         { *m = BasketClass{} }
func (m *BasketClass) String() string { return proto.CompactTextString(m) }
func (*BasketClass) ProtoMessage()    {}
func (*BasketClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_c416a19075224f85, []int{1}
}
func (m *BasketClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasketClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasketClass.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasketClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasketClass.Merge(m, src)
}
func (m *BasketClass) XXX_Size() int {
	return m.Size()
}
func (m *BasketClass) XXX_DiscardUnknown() {
	xxx_messageInfo_BasketClass.DiscardUnknown(m)
}

var xxx_messageInfo_BasketClass proto.InternalMessageInfo

func (m *BasketClass) GetBasketId() uint64 {
	if m != nil {
		return m.BasketId
	}
	return 0
}

func (m *BasketClass) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

// BasketBalance stores the amount of credits from a batch in a basket
type BasketBalance struct {
	// basket_id is the ID of the basket
	BasketId uint64 `protobuf:"varint,1,opt,name=basket_id,json=basketId,proto3" json:"basket_id,omitempty"`
	// batch_denom is the denom of the credit batch
	BatchDenom string `protobuf:"bytes,2,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// balance is the amount of ecocredits held in the basket
	Balance string `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
	// batch_start_date is the start date of the batch. This field is used
	// to create an index which is used to remove the oldest credits first.
	BatchStartDate *types.Timestamp `protobuf:"bytes,4,opt,name=batch_start_date,json=batchStartDate,proto3" json:"batch_start_date,omitempty"`
}

func (m *BasketBalance) Reset()         { *m = BasketBalance{} }
func (m *BasketBalance) String() string { return proto.CompactTextString(m) }
func (*BasketBalance) ProtoMessage()    {}
func (*BasketBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_c416a19075224f85, []int{2}
}
func (m *BasketBalance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasketBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasketBalance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasketBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasketBalance.Merge(m, src)
}
func (m *BasketBalance) XXX_Size() int {
	return m.Size()
}
func (m *BasketBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_BasketBalance.DiscardUnknown(m)
}

var xxx_messageInfo_BasketBalance proto.InternalMessageInfo

func (m *BasketBalance) GetBasketId() uint64 {
	if m != nil {
		return m.BasketId
	}
	return 0
}

func (m *BasketBalance) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

func (m *BasketBalance) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func (m *BasketBalance) GetBatchStartDate() *types.Timestamp {
	if m != nil {
		return m.BatchStartDate
	}
	return nil
}

func init() {
	proto.RegisterType((*Basket)(nil), "regen.ecocredit.basket.v1.Basket")
	proto.RegisterType((*BasketClass)(nil), "regen.ecocredit.basket.v1.BasketClass")
	proto.RegisterType((*BasketBalance)(nil), "regen.ecocredit.basket.v1.BasketBalance")
}

func init() {
	proto.RegisterFile("regen/ecocredit/basket/v1/state.proto", fileDescriptor_c416a19075224f85)
}

var fileDescriptor_c416a19075224f85 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xce, 0x26, 0xfd, 0xa5, 0xc9, 0xa4, 0x89, 0xf2, 0x5b, 0x8a, 0x70, 0x83, 0x70, 0x43, 0x04,
	0xc2, 0x07, 0x58, 0x2b, 0xe5, 0x00, 0xca, 0x8d, 0x34, 0x97, 0x4a, 0x88, 0x83, 0xe9, 0x89, 0x8b,
	0xb5, 0xb6, 0x87, 0xc4, 0xaa, 0xed, 0xb5, 0xd6, 0x9b, 0xd0, 0xbe, 0x04, 0xe2, 0x09, 0xb8, 0xf1,
	0x2e, 0x1c, 0x2b, 0x71, 0xe1, 0x88, 0x12, 0x89, 0x07, 0xe0, 0x09, 0x90, 0x77, 0x93, 0xb4, 0x55,
	0x45, 0x8f, 0xdf, 0xcc, 0xf7, 0xcd, 0x9f, 0x6f, 0x06, 0x9e, 0x4a, 0x9c, 0x62, 0xe6, 0x62, 0x28,
	0x42, 0x89, 0x51, 0xac, 0xdc, 0x80, 0x17, 0x67, 0xa8, 0xdc, 0xc5, 0xd0, 0x2d, 0x14, 0x57, 0xc8,
	0x72, 0x29, 0x94, 0xa0, 0x07, 0x9a, 0xc6, 0xb6, 0x34, 0x66, 0x68, 0x6c, 0x31, 0xec, 0x3d, 0x0a,
	0x45, 0x91, 0x8a, 0xc2, 0x15, 0x32, 0x75, 0x17, 0x43, 0x9e, 0xe4, 0x33, 0x3e, 0x2c, 0x81, 0x51,
	0xf6, 0x0e, 0xa7, 0x42, 0x4c, 0x13, 0x74, 0x35, 0x0a, 0xe6, 0x1f, 0x5d, 0x15, 0xa7, 0x58, 0x28,
	0x9e, 0xe6, 0x6b, 0xc2, 0x1d, 0x13, 0xa8, 0x8b, 0x1c, 0x0b, 0x43, 0x1b, 0x7c, 0xab, 0x42, 0x7d,
	0xac, 0x33, 0xb4, 0x03, 0xd5, 0x38, 0xb2, 0x48, 0x9f, 0x38, 0x3b, 0x5e, 0x35, 0x8e, 0xe8, 0x63,
	0xd8, 0x33, 0x1a, 0x3f, 0xc2, 0x4c, 0xa4, 0x56, 0xb5, 0x4f, 0x9c, 0xa6, 0xd7, 0x32, 0xb1, 0x49,
	0x19, 0xa2, 0x0c, 0xee, 0x45, 0x71, 0xc1, 0x83, 0x04, 0x7d, 0x3e, 0x57, 0xc2, 0x97, 0xa8, 0x62,
	0x89, 0x56, 0xad, 0x4f, 0x9c, 0x86, 0xf7, 0xff, 0x3a, 0xf5, 0x66, 0xae, 0x84, 0xa7, 0x13, 0xd4,
	0x81, 0xae, 0x99, 0xc6, 0x2f, 0x67, 0xf0, 0x33, 0x9e, 0xa2, 0xb5, 0xa3, 0xcb, 0x76, 0x4c, 0xfc,
	0xf4, 0x22, 0xc7, 0x77, 0x3c, 0x45, 0xfa, 0x16, 0xda, 0x11, 0x57, 0xe8, 0x87, 0x32, 0x56, 0x28,
	0x63, 0x6e, 0xfd, 0xd7, 0x27, 0x4e, 0xeb, 0xe8, 0x19, 0xfb, 0xa7, 0x63, 0x6c, 0xc2, 0x15, 0x1e,
	0xaf, 0xe9, 0xde, 0x5e, 0x74, 0x0d, 0xd1, 0x1e, 0x34, 0xf0, 0x3c, 0x17, 0x19, 0x66, 0xca, 0xaa,
	0xf7, 0x89, 0xd3, 0xf6, 0xb6, 0x78, 0xf4, 0xe4, 0xcf, 0xd7, 0x1f, 0x9f, 0x6b, 0x36, 0xd4, 0xcb,
	0xf5, 0xbb, 0x84, 0xd2, 0x9b, 0x6b, 0x77, 0x89, 0x45, 0x2c, 0x32, 0x40, 0x68, 0x19, 0x9b, 0x8e,
	0x13, 0x5e, 0x14, 0xf4, 0x21, 0x34, 0xd7, 0xa4, 0xad, 0x65, 0x0d, 0x13, 0x38, 0x89, 0xe8, 0x01,
	0x34, 0xc2, 0x92, 0x55, 0xe6, 0x8c, 0x69, 0xbb, 0x1a, 0x9f, 0x44, 0x23, 0x5b, 0x37, 0xb3, 0x60,
	0x1f, 0xe8, 0x56, 0xff, 0xfc, 0x8a, 0x3c, 0xf8, 0x4d, 0xa0, 0x6d, 0xfa, 0x8c, 0x79, 0xc2, 0xb3,
	0x10, 0xef, 0xee, 0x74, 0x08, 0xad, 0x80, 0xab, 0x70, 0x76, 0xe3, 0x42, 0xa0, 0x43, 0xe6, 0x40,
	0x16, 0xec, 0x06, 0xa6, 0x90, 0x3e, 0x4a, 0xd3, 0xdb, 0x40, 0x3a, 0x81, 0xae, 0x91, 0x16, 0x8a,
	0x4b, 0xe5, 0x97, 0x76, 0xe9, 0x53, 0xb4, 0x8e, 0x7a, 0xcc, 0xfc, 0x16, 0xdb, 0xfc, 0x16, 0x3b,
	0xdd, 0xfc, 0x96, 0xd7, 0xd1, 0x9a, 0xf7, 0xa5, 0xa4, 0xb4, 0x7b, 0xf4, 0x4a, 0xef, 0x33, 0x84,
	0x07, 0x70, 0xff, 0x6a, 0x9f, 0x6b, 0x23, 0xd1, 0xfd, 0xdb, 0x6d, 0xba, 0xc4, 0xaa, 0x8d, 0xbd,
	0xef, 0x4b, 0x9b, 0x5c, 0x2e, 0x6d, 0xf2, 0x6b, 0x69, 0x93, 0x2f, 0x2b, 0xbb, 0x72, 0xb9, 0xb2,
	0x2b, 0x3f, 0x57, 0x76, 0xe5, 0xc3, 0xeb, 0x69, 0xac, 0x66, 0xf3, 0x80, 0x85, 0x22, 0x75, 0xf5,
	0xb1, 0x5f, 0x64, 0xa8, 0x3e, 0x09, 0x79, 0xb6, 0x46, 0x09, 0x46, 0x53, 0x94, 0xee, 0xf9, 0xad,
	0xd7, 0x0e, 0xea, 0x7a, 0xe0, 0x97, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x02, 0xd5, 0xcc, 0x76,
	0x7d, 0x03, 0x00, 0x00,
}

func (m *Basket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Basket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Basket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Exponent != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Exponent))
		i--
		dAtA[i] = 0x30
	}
	if m.DateCriteria != nil {
		{
			size, err := m.DateCriteria.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CreditTypeName) > 0 {
		i -= len(m.CreditTypeName)
		copy(dAtA[i:], m.CreditTypeName)
		i = encodeVarintState(dAtA, i, uint64(len(m.CreditTypeName)))
		i--
		dAtA[i] = 0x22
	}
	if m.DisableAutoRetire {
		i--
		if m.DisableAutoRetire {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.BasketDenom) > 0 {
		i -= len(m.BasketDenom)
		copy(dAtA[i:], m.BasketDenom)
		i = encodeVarintState(dAtA, i, uint64(len(m.BasketDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BasketClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasketClass) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasketClass) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintState(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0x12
	}
	if m.BasketId != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.BasketId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BasketBalance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasketBalance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasketBalance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BatchStartDate != nil {
		{
			size, err := m.BatchStartDate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Balance) > 0 {
		i -= len(m.Balance)
		copy(dAtA[i:], m.Balance)
		i = encodeVarintState(dAtA, i, uint64(len(m.Balance)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintState(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.BasketId != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.BasketId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Basket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovState(uint64(m.Id))
	}
	l = len(m.BasketDenom)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.DisableAutoRetire {
		n += 2
	}
	l = len(m.CreditTypeName)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.DateCriteria != nil {
		l = m.DateCriteria.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.Exponent != 0 {
		n += 1 + sovState(uint64(m.Exponent))
	}
	return n
}

func (m *BasketClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BasketId != 0 {
		n += 1 + sovState(uint64(m.BasketId))
	}
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func (m *BasketBalance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BasketId != 0 {
		n += 1 + sovState(uint64(m.BasketId))
	}
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Balance)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.BatchStartDate != nil {
		l = m.BatchStartDate.Size()
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Basket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: Basket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Basket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasketDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BasketDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableAutoRetire", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableAutoRetire = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditTypeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreditTypeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DateCriteria", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DateCriteria == nil {
				m.DateCriteria = &DateCriteria{}
			}
			if err := m.DateCriteria.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exponent", wireType)
			}
			m.Exponent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Exponent |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BasketClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: BasketClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasketClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasketId", wireType)
			}
			m.BasketId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BasketId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BasketBalance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: BasketBalance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasketBalance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasketId", wireType)
			}
			m.BasketId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BasketId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchStartDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BatchStartDate == nil {
				m.BatchStartDate = &types.Timestamp{}
			}
			if err := m.BatchStartDate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)