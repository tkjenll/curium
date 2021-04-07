// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: synchronizer/synchronizer-value.proto

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

type SynchronizerValue struct {
	Value    string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Bookmark string `protobuf:"bytes,2,opt,name=bookmark,proto3" json:"bookmark,omitempty"`
}

func (m *SynchronizerValue) Reset()         { *m = SynchronizerValue{} }
func (m *SynchronizerValue) String() string { return proto.CompactTextString(m) }
func (*SynchronizerValue) ProtoMessage()    {}
func (*SynchronizerValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc9f8b307c14bdf7, []int{0}
}
func (m *SynchronizerValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SynchronizerValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SynchronizerValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SynchronizerValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynchronizerValue.Merge(m, src)
}
func (m *SynchronizerValue) XXX_Size() int {
	return m.Size()
}
func (m *SynchronizerValue) XXX_DiscardUnknown() {
	xxx_messageInfo_SynchronizerValue.DiscardUnknown(m)
}

var xxx_messageInfo_SynchronizerValue proto.InternalMessageInfo

func (m *SynchronizerValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *SynchronizerValue) GetBookmark() string {
	if m != nil {
		return m.Bookmark
	}
	return ""
}

func init() {
	proto.RegisterType((*SynchronizerValue)(nil), "bluzelle.curium.synchronizer.SynchronizerValue")
}

func init() {
	proto.RegisterFile("synchronizer/synchronizer-value.proto", fileDescriptor_dc9f8b307c14bdf7)
}

var fileDescriptor_dc9f8b307c14bdf7 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0xae, 0xcc, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcb, 0xac, 0x4a, 0x2d, 0xd2, 0x47, 0xe6, 0xe8, 0x96, 0x25, 0xe6, 0x94,
	0xa6, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xc9, 0x24, 0xe5, 0x94, 0x56, 0xa5, 0xe6, 0xe4,
	0xa4, 0xea, 0x25, 0x97, 0x16, 0x65, 0x96, 0xe6, 0xea, 0x21, 0xab, 0x94, 0x12, 0x49, 0xcf, 0x4f,
	0xcf, 0x07, 0x2b, 0xd4, 0x07, 0xb1, 0x20, 0x7a, 0x94, 0x5c, 0xb9, 0x04, 0x83, 0x91, 0x54, 0x85,
	0x81, 0x8c, 0x13, 0x12, 0xe1, 0x62, 0x05, 0x9b, 0x2b, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0xe1, 0x08, 0x49, 0x71, 0x71, 0x24, 0xe5, 0xe7, 0x67, 0xe7, 0x26, 0x16, 0x65, 0x4b, 0x30, 0x81,
	0x25, 0xe0, 0x7c, 0x27, 0xcf, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2,
	0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0xb9, 0x4f, 0x1f, 0xe2,
	0x3e, 0xfd, 0x0a, 0x14, 0xbf, 0xe8, 0x97, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x1d, 0x66,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x0d, 0x96, 0xe1, 0xf5, 0x00, 0x00, 0x00,
}

func (m *SynchronizerValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SynchronizerValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SynchronizerValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bookmark) > 0 {
		i -= len(m.Bookmark)
		copy(dAtA[i:], m.Bookmark)
		i = encodeVarintSynchronizerValue(dAtA, i, uint64(len(m.Bookmark)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintSynchronizerValue(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSynchronizerValue(dAtA []byte, offset int, v uint64) int {
	offset -= sovSynchronizerValue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SynchronizerValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovSynchronizerValue(uint64(l))
	}
	l = len(m.Bookmark)
	if l > 0 {
		n += 1 + l + sovSynchronizerValue(uint64(l))
	}
	return n
}

func sovSynchronizerValue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSynchronizerValue(x uint64) (n int) {
	return sovSynchronizerValue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SynchronizerValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSynchronizerValue
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
			return fmt.Errorf("proto: SynchronizerValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SynchronizerValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSynchronizerValue
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
				return ErrInvalidLengthSynchronizerValue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSynchronizerValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bookmark", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSynchronizerValue
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
				return ErrInvalidLengthSynchronizerValue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSynchronizerValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bookmark = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSynchronizerValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSynchronizerValue
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
func skipSynchronizerValue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSynchronizerValue
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
					return 0, ErrIntOverflowSynchronizerValue
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
					return 0, ErrIntOverflowSynchronizerValue
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
				return 0, ErrInvalidLengthSynchronizerValue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSynchronizerValue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSynchronizerValue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSynchronizerValue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSynchronizerValue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSynchronizerValue = fmt.Errorf("proto: unexpected end of group")
)