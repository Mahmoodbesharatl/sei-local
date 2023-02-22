// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mint/v1beta1/mint.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Minter represents the most recent
type Minter struct {
	LastMintAmount github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=last_mint_amount,json=lastMintAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"last_mint_amount" yaml:"last_mint_amount"`
	LastMintDate   string                                 `protobuf:"bytes,2,opt,name=last_mint_date,json=lastMintDate,proto3" json:"last_mint_date,omitempty" yaml:"last_mint_date"`
	LastMintHeight int64                                  `protobuf:"varint,3,opt,name=last_mint_height,json=lastMintHeight,proto3" json:"last_mint_height,omitempty" yaml:"last_mint_height"`
	Denom          string                                 `protobuf:"bytes,4,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_06339c129491fd39, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

func (m *Minter) GetLastMintDate() string {
	if m != nil {
		return m.LastMintDate
	}
	return ""
}

func (m *Minter) GetLastMintHeight() int64 {
	if m != nil {
		return m.LastMintHeight
	}
	return 0
}

func (m *Minter) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type ScheduledTokenRelease struct {
	Date               string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	TokenReleaseAmount int64  `protobuf:"varint,2,opt,name=token_release_amount,json=tokenReleaseAmount,proto3" json:"token_release_amount,omitempty"`
}

func (m *ScheduledTokenRelease) Reset()         { *m = ScheduledTokenRelease{} }
func (m *ScheduledTokenRelease) String() string { return proto.CompactTextString(m) }
func (*ScheduledTokenRelease) ProtoMessage()    {}
func (*ScheduledTokenRelease) Descriptor() ([]byte, []int) {
	return fileDescriptor_06339c129491fd39, []int{1}
}
func (m *ScheduledTokenRelease) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduledTokenRelease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScheduledTokenRelease.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScheduledTokenRelease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduledTokenRelease.Merge(m, src)
}
func (m *ScheduledTokenRelease) XXX_Size() int {
	return m.Size()
}
func (m *ScheduledTokenRelease) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduledTokenRelease.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduledTokenRelease proto.InternalMessageInfo

func (m *ScheduledTokenRelease) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *ScheduledTokenRelease) GetTokenReleaseAmount() int64 {
	if m != nil {
		return m.TokenReleaseAmount
	}
	return 0
}

// Params holds parameters for the mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// List of token release schedules
	TokenReleaseSchedule []ScheduledTokenRelease `protobuf:"bytes,2,rep,name=token_release_schedule,json=tokenReleaseSchedule,proto3" json:"token_release_schedule" yaml:"token_release_schedule"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_06339c129491fd39, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetTokenReleaseSchedule() []ScheduledTokenRelease {
	if m != nil {
		return m.TokenReleaseSchedule
	}
	return nil
}

func init() {
	proto.RegisterType((*Minter)(nil), "seiprotocol.seichain.mint.Minter")
	proto.RegisterType((*ScheduledTokenRelease)(nil), "seiprotocol.seichain.mint.ScheduledTokenRelease")
	proto.RegisterType((*Params)(nil), "seiprotocol.seichain.mint.Params")
}

func init() { proto.RegisterFile("mint/v1beta1/mint.proto", fileDescriptor_06339c129491fd39) }

var fileDescriptor_06339c129491fd39 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xd5, 0xda, 0xae, 0x21, 0xdb, 0x10, 0xc2, 0xe2, 0x34, 0x4a, 0x4b, 0x24, 0xb3, 0xd0, 0xe0,
	0x4b, 0x56, 0x49, 0x7b, 0xcb, 0xa5, 0xd4, 0xb8, 0x90, 0x1e, 0x0a, 0x45, 0xed, 0xa9, 0x50, 0xcc,
	0x5a, 0x1a, 0xac, 0x25, 0x92, 0x36, 0x78, 0xd7, 0xa5, 0xf9, 0x88, 0x42, 0x8f, 0x3d, 0xf6, 0x07,
	0xfa, 0x09, 0xbd, 0xe7, 0x98, 0x63, 0xe9, 0x41, 0x14, 0xfb, 0x0f, 0xf4, 0x05, 0x65, 0x77, 0x23,
	0xaa, 0x06, 0xe7, 0xa4, 0xd9, 0x37, 0x6f, 0xde, 0xbc, 0x19, 0x0d, 0xde, 0x2f, 0x44, 0xa9, 0xa3,
	0x4f, 0xa7, 0x33, 0xd0, 0xfc, 0x34, 0x32, 0x0f, 0x76, 0xb9, 0x90, 0x5a, 0x92, 0x03, 0x05, 0xc2,
	0x46, 0x89, 0xcc, 0x99, 0x02, 0x91, 0x64, 0x5c, 0x94, 0xcc, 0x10, 0x1e, 0x0f, 0xe6, 0x72, 0x2e,
	0x6d, 0x2e, 0x32, 0x91, 0x2b, 0xa0, 0x3f, 0x3a, 0xb8, 0xff, 0x46, 0x94, 0x1a, 0x16, 0x44, 0xe1,
	0xdd, 0x9c, 0x2b, 0x3d, 0x35, 0xec, 0x29, 0x2f, 0xe4, 0xb2, 0xd4, 0x3e, 0x1a, 0xa2, 0xd1, 0xd6,
	0xf8, 0xf5, 0x75, 0x15, 0x7a, 0xbf, 0xab, 0xf0, 0x68, 0x2e, 0x74, 0xb6, 0x9c, 0xb1, 0x44, 0x16,
	0x51, 0x22, 0x55, 0x21, 0xd5, 0xed, 0xe7, 0x58, 0xa5, 0x17, 0x91, 0xbe, 0xba, 0x04, 0xc5, 0x26,
	0x90, 0xd4, 0x55, 0xb8, 0x7f, 0xc5, 0x8b, 0xfc, 0x8c, 0xde, 0xd5, 0xa3, 0xf1, 0x8e, 0x81, 0x4c,
	0xc3, 0x97, 0x16, 0x20, 0x2f, 0xf0, 0xce, 0x3f, 0x52, 0xca, 0x35, 0xf8, 0x1d, 0xdb, 0xf2, 0xa0,
	0xae, 0xc2, 0xbd, 0xbb, 0x22, 0x26, 0x4f, 0xe3, 0xed, 0x46, 0x62, 0xc2, 0x35, 0x90, 0x57, 0x6d,
	0xd7, 0x19, 0x88, 0x79, 0xa6, 0xfd, 0xee, 0x10, 0x8d, 0xba, 0xe3, 0x27, 0x9b, 0x7c, 0x38, 0x46,
	0xcb, 0xc7, 0xb9, 0x05, 0xc8, 0x11, 0x7e, 0x90, 0x42, 0x29, 0x0b, 0xbf, 0x67, 0xdb, 0xef, 0xd6,
	0x55, 0xb8, 0xed, 0x6a, 0x2d, 0x4c, 0x63, 0x97, 0xa6, 0x1f, 0xf1, 0xde, 0xbb, 0x24, 0x83, 0x74,
	0x99, 0x43, 0xfa, 0x5e, 0x5e, 0x40, 0x19, 0x43, 0x0e, 0x5c, 0x01, 0x21, 0xb8, 0x67, 0xed, 0xdb,
	0x8d, 0xc5, 0x36, 0x26, 0x27, 0x78, 0xa0, 0x0d, 0x67, 0xba, 0x70, 0xa4, 0x66, 0xab, 0x66, 0xc4,
	0x6e, 0x4c, 0x74, 0xab, 0xde, 0xad, 0x83, 0xfe, 0x44, 0xb8, 0xff, 0x96, 0x2f, 0x78, 0xa1, 0xc8,
	0x21, 0xc6, 0x6e, 0x68, 0x6b, 0xcb, 0xc9, 0x6e, 0x19, 0x64, 0x62, 0x00, 0xf2, 0x05, 0xe1, 0x47,
	0xff, 0x8b, 0xab, 0x5b, 0x5f, 0x7e, 0x67, 0xd8, 0x1d, 0x3d, 0x7c, 0x76, 0xc2, 0xee, 0xbd, 0x05,
	0xb6, 0x71, 0x84, 0xf1, 0x53, 0xf3, 0x9b, 0xeb, 0x2a, 0x3c, 0x74, 0x83, 0x6f, 0x56, 0xa7, 0xf1,
	0xa0, 0xed, 0xbb, 0x51, 0x3a, 0xeb, 0x7d, 0xfb, 0x1e, 0x7a, 0xe3, 0xf3, 0xeb, 0x55, 0x80, 0x6e,
	0x56, 0x01, 0xfa, 0xb3, 0x0a, 0xd0, 0xd7, 0x75, 0xe0, 0xdd, 0xac, 0x03, 0xef, 0xd7, 0x3a, 0xf0,
	0x3e, 0xb0, 0xd6, 0xed, 0x28, 0x10, 0xc7, 0x8d, 0x33, 0xfb, 0xb0, 0xd6, 0xa2, 0xcf, 0xf6, 0x92,
	0xdd, 0x1d, 0xcd, 0xfa, 0x96, 0xf0, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x7e, 0xc4,
	0xcc, 0xeb, 0x02, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x22
	}
	if m.LastMintHeight != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.LastMintHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.LastMintDate) > 0 {
		i -= len(m.LastMintDate)
		copy(dAtA[i:], m.LastMintDate)
		i = encodeVarintMint(dAtA, i, uint64(len(m.LastMintDate)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.LastMintAmount.Size()
		i -= size
		if _, err := m.LastMintAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ScheduledTokenRelease) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduledTokenRelease) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduledTokenRelease) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TokenReleaseAmount != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.TokenReleaseAmount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Date) > 0 {
		i -= len(m.Date)
		copy(dAtA[i:], m.Date)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Date)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenReleaseSchedule) > 0 {
		for iNdEx := len(m.TokenReleaseSchedule) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenReleaseSchedule[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMint(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.LastMintAmount.Size()
	n += 1 + l + sovMint(uint64(l))
	l = len(m.LastMintDate)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if m.LastMintHeight != 0 {
		n += 1 + sovMint(uint64(m.LastMintHeight))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	return n
}

func (m *ScheduledTokenRelease) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Date)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if m.TokenReleaseAmount != 0 {
		n += 1 + sovMint(uint64(m.TokenReleaseAmount))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if len(m.TokenReleaseSchedule) > 0 {
		for _, e := range m.TokenReleaseSchedule {
			l = e.Size()
			n += 1 + l + sovMint(uint64(l))
		}
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastMintAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastMintAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastMintDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastMintDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastMintHeight", wireType)
			}
			m.LastMintHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastMintHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *ScheduledTokenRelease) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: ScheduledTokenRelease: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduledTokenRelease: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Date = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenReleaseAmount", wireType)
			}
			m.TokenReleaseAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenReleaseAmount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenReleaseSchedule", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenReleaseSchedule = append(m.TokenReleaseSchedule, ScheduledTokenRelease{})
			if err := m.TokenReleaseSchedule[len(m.TokenReleaseSchedule)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
