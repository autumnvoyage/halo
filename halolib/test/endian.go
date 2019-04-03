package halo

import (
	"testing"
)

const (
	testS16 = -4555
	testU16 = 40170
	testS32 = -100500
	testU32 = 640576
	testS64 = -4000000001
	testU64 = 5987654321
)

func TestEndian(t *testing.T) {
	bS16 := EncodeNetInt16(testS16)
	bU16 := EncodeNetUint16(testU16)
	bS32 := EncodeNetInt32(testS32)
	bU32 := EncodeNetUint32(testU32)
	bS64 := EncodeNetInt64(testS64)
	bU64 := EncodeNetUint64(testU64)
	oS16 := DecodeNetInt16(bS16)
	oU16 := DecodeNetUint16(bU16)
	oS32 := DecodeNetInt32(bS32)
	oU32 := DecodeNetUint32(bU32)
	oS64 := DecodeNetInt64(bS64)
	oU64 := DecodeNetUint64(bU64)
	if oS16 != testS16 {
		t.Errorf("Mismatched s16\nExpected: %v\nGot: %v\nBytes: %x", testS16,
			oS16, bS16)
	}
	if oU16 != testU16 {
		t.Errorf("Mismatched u16\nExpected: %v\nGot: %v\nBytes: %x", testU16,
			oU16, bU16)
	}
	if oS32 != testS32 {
		t.Errorf("Mismatched s32\nExpected: %v\nGot: %v\nBytes: %x", testS32,
			oS32, bS32)
	}
	if oU32 != testU32 {
		t.Errorf("Mismatched u32\nExpected: %v\nGot: %v\nBytes: %x", testU32,
			oU32, bU32)
	}
	if oS64 != testS64 {
		t.Errorf("Mismatched s64\nExpected: %v\nGot: %v\nBytes: %x", testS64,
			oS64, bS64)
	}
	if oU64 != testU64 {
		t.Errorf("Mismatched u64\nExpected: %v\nGot: %v\nBytes: %x", testU64,
			oU64, bU64)
	}
}
