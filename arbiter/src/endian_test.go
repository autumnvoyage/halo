package main

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
	bS16 := encodeNetInt16(testS16)
	bU16 := encodeNetUint16(testU16)
	bS32 := encodeNetInt32(testS32)
	bU32 := encodeNetUint32(testU32)
	bS64 := encodeNetInt64(testS64)
	bU64 := encodeNetUint64(testU64)
	oS16 := decodeNetInt16(bS16)
	oU16 := decodeNetUint16(bU16)
	oS32 := decodeNetInt32(bS32)
	oU32 := decodeNetUint32(bU32)
	oS64 := decodeNetInt64(bS64)
	oU64 := decodeNetUint64(bU64)
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
