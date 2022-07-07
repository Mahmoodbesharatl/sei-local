// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/asset_list.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/bank/types"
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

type AssetIBCInfo struct {
	SourceChannel string `protobuf:"bytes,1,opt,name=sourceChannel,proto3" json:"source_channel"`
	DstChannel    string `protobuf:"bytes,2,opt,name=dstChannel,proto3" json:"dst_channel"`
	SourceDenom   string `protobuf:"bytes,3,opt,name=sourceDenom,proto3" json:"source_denom"`
	SourceChainID string `protobuf:"bytes,4,opt,name=sourceChainID,proto3" json:"source_chain_id"`
}

func (m *AssetIBCInfo) Reset()         { *m = AssetIBCInfo{} }
func (m *AssetIBCInfo) String() string { return proto.CompactTextString(m) }
func (*AssetIBCInfo) ProtoMessage()    {}
func (*AssetIBCInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_65992190d1b6a2ed, []int{0}
}
func (m *AssetIBCInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetIBCInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetIBCInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetIBCInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetIBCInfo.Merge(m, src)
}
func (m *AssetIBCInfo) XXX_Size() int {
	return m.Size()
}
func (m *AssetIBCInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetIBCInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AssetIBCInfo proto.InternalMessageInfo

func (m *AssetIBCInfo) GetSourceChannel() string {
	if m != nil {
		return m.SourceChannel
	}
	return ""
}

func (m *AssetIBCInfo) GetDstChannel() string {
	if m != nil {
		return m.DstChannel
	}
	return ""
}

func (m *AssetIBCInfo) GetSourceDenom() string {
	if m != nil {
		return m.SourceDenom
	}
	return ""
}

func (m *AssetIBCInfo) GetSourceChainID() string {
	if m != nil {
		return m.SourceChainID
	}
	return ""
}

type AssetMetadata struct {
	IbcInfo   AssetIBCInfo   `protobuf:"bytes,1,opt,name=ibcInfo,proto3" json:"ibc_info"`
	TypeAsset string         `protobuf:"bytes,2,opt,name=type_asset,json=typeAsset,proto3" json:"type_asset"`
	Metadata  types.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata"`
}

func (m *AssetMetadata) Reset()         { *m = AssetMetadata{} }
func (m *AssetMetadata) String() string { return proto.CompactTextString(m) }
func (*AssetMetadata) ProtoMessage()    {}
func (*AssetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_65992190d1b6a2ed, []int{1}
}
func (m *AssetMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetMetadata.Merge(m, src)
}
func (m *AssetMetadata) XXX_Size() int {
	return m.Size()
}
func (m *AssetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_AssetMetadata proto.InternalMessageInfo

func (m *AssetMetadata) GetIbcInfo() AssetIBCInfo {
	if m != nil {
		return m.IbcInfo
	}
	return AssetIBCInfo{}
}

func (m *AssetMetadata) GetTypeAsset() string {
	if m != nil {
		return m.TypeAsset
	}
	return ""
}

func (m *AssetMetadata) GetMetadata() types.Metadata {
	if m != nil {
		return m.Metadata
	}
	return types.Metadata{}
}

func init() {
	proto.RegisterType((*AssetIBCInfo)(nil), "seiprotocol.seichain.dex.AssetIBCInfo")
	proto.RegisterType((*AssetMetadata)(nil), "seiprotocol.seichain.dex.AssetMetadata")
}

func init() { proto.RegisterFile("dex/asset_list.proto", fileDescriptor_65992190d1b6a2ed) }

var fileDescriptor_65992190d1b6a2ed = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x31, 0xaf, 0xd3, 0x30,
	0x10, 0x4e, 0x00, 0xc1, 0x7b, 0xce, 0x7b, 0x6d, 0x65, 0x3a, 0x44, 0x95, 0x48, 0x50, 0x07, 0xc4,
	0x52, 0x5b, 0x2d, 0x0b, 0x8c, 0xa4, 0x95, 0x50, 0x85, 0x58, 0x32, 0xb2, 0x44, 0x8e, 0xe3, 0xb6,
	0x16, 0x8d, 0x5d, 0xd5, 0x2e, 0x2a, 0xff, 0x82, 0x9f, 0xd5, 0xb1, 0x23, 0x53, 0x40, 0xed, 0x82,
	0xf2, 0x2b, 0x90, 0x9d, 0xa4, 0x4d, 0x07, 0xb6, 0xcf, 0x77, 0xf7, 0xdd, 0x7d, 0xdf, 0x9d, 0x41,
	0x3f, 0x63, 0x7b, 0x4c, 0x94, 0x62, 0x3a, 0x59, 0x73, 0xa5, 0xd1, 0x66, 0x2b, 0xb5, 0x84, 0xbe,
	0x62, 0xdc, 0x22, 0x2a, 0xd7, 0x48, 0x31, 0x4e, 0x57, 0x84, 0x0b, 0x94, 0xb1, 0xfd, 0x20, 0xa0,
	0x52, 0xe5, 0x52, 0xe1, 0x94, 0x88, 0x6f, 0xf8, 0xfb, 0x38, 0x65, 0x9a, 0x8c, 0xed, 0xa3, 0x62,
	0x0e, 0xfa, 0x4b, 0xb9, 0x94, 0x16, 0x62, 0x83, 0xea, 0x68, 0xd7, 0x4c, 0x61, 0x62, 0x97, 0xab,
	0x2a, 0x30, 0xfc, 0xeb, 0x82, 0x87, 0x8f, 0x66, 0xea, 0x3c, 0x9a, 0xce, 0xc5, 0x42, 0xc2, 0xf7,
	0xe0, 0x51, 0xc9, 0xdd, 0x96, 0xb2, 0xe9, 0x8a, 0x08, 0xc1, 0xd6, 0xbe, 0xfb, 0xda, 0x7d, 0x7b,
	0x1f, 0xc1, 0xb2, 0x08, 0x3b, 0x55, 0x22, 0xa1, 0x55, 0x26, 0xbe, 0x2d, 0x84, 0x18, 0x80, 0x4c,
	0xe9, 0x86, 0xf6, 0xc4, 0xd2, 0xba, 0x65, 0x11, 0x7a, 0x99, 0xd2, 0x17, 0x4e, 0xab, 0x04, 0x4e,
	0x80, 0x57, 0x75, 0x98, 0x31, 0x21, 0x73, 0xff, 0xa9, 0x65, 0xf4, 0xca, 0x22, 0x7c, 0xa8, 0x07,
	0x65, 0x26, 0x1e, 0xb7, 0x8b, 0xe0, 0x87, 0x96, 0x3c, 0x2e, 0xe6, 0x33, 0xff, 0x99, 0x65, 0xbd,
	0x2c, 0x8b, 0xb0, 0x7b, 0x95, 0xc7, 0x45, 0xc2, 0xb3, 0xf8, 0xb6, 0x72, 0xf8, 0xdb, 0x05, 0x8f,
	0xd6, 0xea, 0x17, 0xa6, 0x49, 0x46, 0x34, 0x81, 0x31, 0x78, 0xc1, 0x53, 0x6a, 0x6c, 0x5b, 0x97,
	0xde, 0xe4, 0x0d, 0xfa, 0xdf, 0xbe, 0x51, 0x7b, 0x49, 0x51, 0xef, 0x50, 0x84, 0x4e, 0x59, 0x84,
	0x77, 0x3c, 0xa5, 0x09, 0x17, 0x0b, 0x19, 0x37, 0x8d, 0xe0, 0x08, 0x00, 0xfd, 0x63, 0xc3, 0x12,
	0x7b, 0xca, 0x7a, 0x0b, 0x9d, 0xb2, 0x08, 0x5b, 0xd1, 0xf8, 0xde, 0x60, 0xdb, 0x10, 0x7e, 0x06,
	0x77, 0x79, 0x2d, 0xc7, 0x2e, 0xc0, 0x9b, 0xbc, 0x42, 0xd5, 0x65, 0x91, 0x3d, 0x66, 0x7d, 0x59,
	0xd4, 0x68, 0xbe, 0x8e, 0x6e, 0x68, 0xf1, 0x05, 0x45, 0x9f, 0x0e, 0xa7, 0xc0, 0x3d, 0x9e, 0x02,
	0xf7, 0xcf, 0x29, 0x70, 0x7f, 0x9e, 0x03, 0xe7, 0x78, 0x0e, 0x9c, 0x5f, 0xe7, 0xc0, 0xf9, 0x3a,
	0x5a, 0x72, 0xbd, 0xda, 0xa5, 0x88, 0xca, 0x1c, 0x2b, 0xc6, 0x47, 0x8d, 0x47, 0xfb, 0xb0, 0x26,
	0xf1, 0x1e, 0x9b, 0xbf, 0x61, 0x94, 0xa9, 0xf4, 0xb9, 0xcd, 0xbf, 0xfb, 0x17, 0x00, 0x00, 0xff,
	0xff, 0xc9, 0xc6, 0xd3, 0x25, 0x95, 0x02, 0x00, 0x00,
}

func (m *AssetIBCInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetIBCInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetIBCInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SourceChainID) > 0 {
		i -= len(m.SourceChainID)
		copy(dAtA[i:], m.SourceChainID)
		i = encodeVarintAssetList(dAtA, i, uint64(len(m.SourceChainID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SourceDenom) > 0 {
		i -= len(m.SourceDenom)
		copy(dAtA[i:], m.SourceDenom)
		i = encodeVarintAssetList(dAtA, i, uint64(len(m.SourceDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DstChannel) > 0 {
		i -= len(m.DstChannel)
		copy(dAtA[i:], m.DstChannel)
		i = encodeVarintAssetList(dAtA, i, uint64(len(m.DstChannel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintAssetList(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AssetMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAssetList(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.TypeAsset) > 0 {
		i -= len(m.TypeAsset)
		copy(dAtA[i:], m.TypeAsset)
		i = encodeVarintAssetList(dAtA, i, uint64(len(m.TypeAsset)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.IbcInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAssetList(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAssetList(dAtA []byte, offset int, v uint64) int {
	offset -= sovAssetList(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AssetIBCInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovAssetList(uint64(l))
	}
	l = len(m.DstChannel)
	if l > 0 {
		n += 1 + l + sovAssetList(uint64(l))
	}
	l = len(m.SourceDenom)
	if l > 0 {
		n += 1 + l + sovAssetList(uint64(l))
	}
	l = len(m.SourceChainID)
	if l > 0 {
		n += 1 + l + sovAssetList(uint64(l))
	}
	return n
}

func (m *AssetMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.IbcInfo.Size()
	n += 1 + l + sovAssetList(uint64(l))
	l = len(m.TypeAsset)
	if l > 0 {
		n += 1 + l + sovAssetList(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovAssetList(uint64(l))
	return n
}

func sovAssetList(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAssetList(x uint64) (n int) {
	return sovAssetList(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AssetIBCInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAssetList
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
			return fmt.Errorf("proto: AssetIBCInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetIBCInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetList
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
				return ErrInvalidLengthAssetList
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAssetList
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DstChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetList
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
				return ErrInvalidLengthAssetList
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAssetList
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DstChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetList
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
				return ErrInvalidLengthAssetList
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAssetList
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetList
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
				return ErrInvalidLengthAssetList
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAssetList
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAssetList(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAssetList
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
func (m *AssetMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAssetList
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
			return fmt.Errorf("proto: AssetMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetList
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
				return ErrInvalidLengthAssetList
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAssetList
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IbcInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetList
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
				return ErrInvalidLengthAssetList
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAssetList
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TypeAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetList
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
				return ErrInvalidLengthAssetList
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAssetList
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAssetList(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAssetList
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
func skipAssetList(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAssetList
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
					return 0, ErrIntOverflowAssetList
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
					return 0, ErrIntOverflowAssetList
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
				return 0, ErrInvalidLengthAssetList
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAssetList
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAssetList
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAssetList        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAssetList          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAssetList = fmt.Errorf("proto: unexpected end of group")
)
