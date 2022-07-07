// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/gov.proto

package types

import (
	fmt "fmt"
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

// RegisterPairsProposal is a gov Content type for adding a new whitelisted token
// pair to the dex module. It must specify a list of contract addresses and their respective
// token pairs to be registered.
type RegisterPairsProposal struct {
	Title             string              `protobuf:"bytes,1,opt,name=title,proto3" json:"title" yaml:"title"`
	Description       string              `protobuf:"bytes,2,opt,name=description,proto3" json:"description" yaml:"description"`
	Batchcontractpair []BatchContractPair `protobuf:"bytes,3,rep,name=batchcontractpair,proto3" json:"batch_contract_pair" yaml:"batch_contract_pair"`
}

func (m *RegisterPairsProposal) Reset()      { *m = RegisterPairsProposal{} }
func (*RegisterPairsProposal) ProtoMessage() {}
func (*RegisterPairsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_dab07ca1a96062d0, []int{0}
}
func (m *RegisterPairsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterPairsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterPairsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterPairsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterPairsProposal.Merge(m, src)
}
func (m *RegisterPairsProposal) XXX_Size() int {
	return m.Size()
}
func (m *RegisterPairsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterPairsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterPairsProposal proto.InternalMessageInfo

type UpdateTickSizeProposal struct {
	Title        string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title" yaml:"title"`
	Description  string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description" yaml:"description"`
	TickSizeList []TickSize `protobuf:"bytes,3,rep,name=tickSizeList,proto3" json:"tick_size_list" yaml:"tick_size_list"`
}

func (m *UpdateTickSizeProposal) Reset()      { *m = UpdateTickSizeProposal{} }
func (*UpdateTickSizeProposal) ProtoMessage() {}
func (*UpdateTickSizeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_dab07ca1a96062d0, []int{1}
}
func (m *UpdateTickSizeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateTickSizeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateTickSizeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateTickSizeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTickSizeProposal.Merge(m, src)
}
func (m *UpdateTickSizeProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateTickSizeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTickSizeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTickSizeProposal proto.InternalMessageInfo

// AddAssetMetadataProposal is a gov Content type for adding a new asset
// to the dex module's asset list.
type AddAssetMetadataProposal struct {
	Title       string          `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description string          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	AssetList   []AssetMetadata `protobuf:"bytes,3,rep,name=assetList,proto3" json:"assetList" yaml:"asset_list"`
}

func (m *AddAssetMetadataProposal) Reset()      { *m = AddAssetMetadataProposal{} }
func (*AddAssetMetadataProposal) ProtoMessage() {}
func (*AddAssetMetadataProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_dab07ca1a96062d0, []int{2}
}
func (m *AddAssetMetadataProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddAssetMetadataProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddAssetMetadataProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddAssetMetadataProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAssetMetadataProposal.Merge(m, src)
}
func (m *AddAssetMetadataProposal) XXX_Size() int {
	return m.Size()
}
func (m *AddAssetMetadataProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAssetMetadataProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AddAssetMetadataProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterPairsProposal)(nil), "seiprotocol.seichain.dex.RegisterPairsProposal")
	proto.RegisterType((*UpdateTickSizeProposal)(nil), "seiprotocol.seichain.dex.UpdateTickSizeProposal")
	proto.RegisterType((*AddAssetMetadataProposal)(nil), "seiprotocol.seichain.dex.AddAssetMetadataProposal")
}

func init() { proto.RegisterFile("dex/gov.proto", fileDescriptor_dab07ca1a96062d0) }

var fileDescriptor_dab07ca1a96062d0 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0xbf, 0x6e, 0xd4, 0x40,
	0x10, 0xc6, 0xed, 0x8b, 0x40, 0x8a, 0x73, 0x44, 0xc4, 0xf9, 0x23, 0xe7, 0x0a, 0x6f, 0xb4, 0x12,
	0x10, 0x09, 0xc5, 0x96, 0xa0, 0x81, 0x74, 0x31, 0x45, 0x1a, 0x90, 0x22, 0x03, 0x0d, 0xcd, 0xb1,
	0xb7, 0x5e, 0xf9, 0x56, 0x71, 0x6e, 0x2d, 0xef, 0x00, 0x97, 0x3c, 0x01, 0x25, 0x05, 0x05, 0x74,
	0xf7, 0x14, 0x3c, 0x43, 0xca, 0x94, 0x54, 0x2b, 0x74, 0xd7, 0xa0, 0x13, 0x95, 0x9f, 0x00, 0xed,
	0xda, 0x97, 0xf8, 0xe0, 0x4e, 0xa2, 0xa4, 0xdb, 0xf9, 0x66, 0x67, 0xf6, 0x9b, 0x9f, 0x66, 0x9d,
	0x3b, 0x09, 0x1b, 0x86, 0xa9, 0x78, 0x1f, 0xe4, 0x85, 0x00, 0xe1, 0x7a, 0x92, 0x71, 0x73, 0xa2,
	0x22, 0x0b, 0x24, 0xe3, 0xb4, 0x4f, 0xf8, 0x20, 0x48, 0xd8, 0xb0, 0xb3, 0x95, 0x8a, 0x54, 0x98,
	0x54, 0xa8, 0x4f, 0xd5, 0xfd, 0xce, 0xba, 0x2e, 0xcf, 0x09, 0x2f, 0xea, 0x78, 0x4b, 0xc7, 0x44,
	0x4a, 0x06, 0xdd, 0x8c, 0x4b, 0xa8, 0xd5, 0x4d, 0xad, 0x02, 0xa7, 0xa7, 0x5d, 0xc9, 0x2f, 0x58,
	0x25, 0xe2, 0x6f, 0x2d, 0x67, 0x3b, 0x66, 0x29, 0x97, 0xc0, 0x8a, 0x13, 0xc2, 0x0b, 0x79, 0x52,
	0x88, 0x5c, 0x48, 0x92, 0xb9, 0xa1, 0x73, 0x0b, 0x38, 0x64, 0xcc, 0xb3, 0xf7, 0xec, 0xfd, 0xd5,
	0x68, 0x77, 0xaa, 0x50, 0x25, 0x94, 0x0a, 0xb5, 0xcf, 0xc9, 0x59, 0x76, 0x88, 0x4d, 0x88, 0xe3,
	0x4a, 0x76, 0x8f, 0x9d, 0xb5, 0x84, 0x49, 0x5a, 0xf0, 0x1c, 0xb8, 0x18, 0x78, 0x2d, 0x53, 0x76,
	0x6f, 0xaa, 0x50, 0x53, 0x2e, 0x15, 0x72, 0xab, 0xe2, 0x86, 0x88, 0xe3, 0xe6, 0x15, 0xf7, 0xb3,
	0xed, 0x6c, 0xf4, 0x08, 0xd0, 0x3e, 0x15, 0x03, 0x28, 0x08, 0x05, 0x3d, 0x9a, 0xb7, 0xb2, 0xb7,
	0xb2, 0xbf, 0xf6, 0xe8, 0x61, 0xb0, 0x8c, 0x4d, 0x10, 0xe9, 0x92, 0x67, 0x75, 0x89, 0x9e, 0x25,
	0x7a, 0x7a, 0xa9, 0x90, 0x35, 0x55, 0x68, 0xd3, 0x74, 0xeb, 0xce, 0xda, 0x75, 0x75, 0xbf, 0x52,
	0xa1, 0x4e, 0x65, 0x64, 0x41, 0x12, 0xc7, 0x7f, 0x1b, 0x38, 0x6c, 0x7f, 0x1c, 0x21, 0xeb, 0xcb,
	0x08, 0x59, 0x3f, 0x47, 0xc8, 0xc2, 0x5f, 0x5b, 0xce, 0xce, 0xeb, 0x3c, 0x21, 0xc0, 0x5e, 0x71,
	0x7a, 0xfa, 0x92, 0x5f, 0xb0, 0xff, 0x80, 0xdc, 0x07, 0xa7, 0x0d, 0xb5, 0x9b, 0xe7, 0x5c, 0x42,
	0xcd, 0x0c, 0x2f, 0x67, 0x36, 0xf3, 0x1e, 0x85, 0x35, 0xaa, 0xf5, 0xeb, 0x05, 0x31, 0xab, 0x53,
	0x2a, 0xb4, 0x3d, 0x73, 0xdc, 0xd4, 0x71, 0x3c, 0xf7, 0xd0, 0x1f, 0x6c, 0x7e, 0xd9, 0x8e, 0x77,
	0x94, 0x24, 0x47, 0x7a, 0x03, 0x5f, 0x30, 0x20, 0x09, 0x01, 0x72, 0x4d, 0xe7, 0xfe, 0x3c, 0x9d,
	0xbb, 0xcb, 0xa0, 0x3c, 0x59, 0x04, 0x65, 0xe7, 0x5f, 0x28, 0xbc, 0x75, 0x56, 0xcd, 0xf2, 0x37,
	0x10, 0x3c, 0x58, 0x8e, 0x60, 0xce, 0x65, 0xb4, 0xab, 0x39, 0x94, 0x0a, 0x6d, 0x54, 0x8f, 0xdc,
	0x7c, 0x22, 0x1c, 0xdf, 0x34, 0x9d, 0x1f, 0x37, 0x3a, 0xbe, 0x1c, 0xfb, 0xf6, 0xd5, 0xd8, 0xb7,
	0x7f, 0x8c, 0x7d, 0xfb, 0xd3, 0xc4, 0xb7, 0xae, 0x26, 0xbe, 0xf5, 0x7d, 0xe2, 0x5b, 0x6f, 0x0e,
	0x52, 0x0e, 0xfd, 0x77, 0xbd, 0x80, 0x8a, 0xb3, 0x50, 0x32, 0x7e, 0x30, 0x73, 0x60, 0x02, 0x63,
	0x21, 0x1c, 0x86, 0xe6, 0x5b, 0x9e, 0xe7, 0x4c, 0xf6, 0x6e, 0x9b, 0xfc, 0xe3, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xd3, 0x4d, 0x06, 0xff, 0x0f, 0x04, 0x00, 0x00,
}

func (m *RegisterPairsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterPairsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterPairsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Batchcontractpair) > 0 {
		for iNdEx := len(m.Batchcontractpair) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Batchcontractpair[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateTickSizeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateTickSizeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateTickSizeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TickSizeList) > 0 {
		for iNdEx := len(m.TickSizeList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TickSizeList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AddAssetMetadataProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddAssetMetadataProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddAssetMetadataProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AssetList) > 0 {
		for iNdEx := len(m.AssetList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AssetList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegisterPairsProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.Batchcontractpair) > 0 {
		for _, e := range m.Batchcontractpair {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *UpdateTickSizeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.TickSizeList) > 0 {
		for _, e := range m.TickSizeList {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *AddAssetMetadataProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.AssetList) > 0 {
		for _, e := range m.AssetList {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterPairsProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: RegisterPairsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterPairsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Batchcontractpair", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Batchcontractpair = append(m.Batchcontractpair, BatchContractPair{})
			if err := m.Batchcontractpair[len(m.Batchcontractpair)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *UpdateTickSizeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: UpdateTickSizeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateTickSizeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickSizeList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TickSizeList = append(m.TickSizeList, TickSize{})
			if err := m.TickSizeList[len(m.TickSizeList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *AddAssetMetadataProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: AddAssetMetadataProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddAssetMetadataProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetList = append(m.AssetList, AssetMetadata{})
			if err := m.AssetList[len(m.AssetList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
