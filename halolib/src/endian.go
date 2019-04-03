package halo

func DecodeNetInt16(in [2]byte) int16 {
	out := [2]int16{
		int16(in[0]),
		int16(in[1]),
	}
	return int16(out[1] | (out[0] << 8))
}

func EncodeNetInt16(in int16) [2]byte {
	return [2]byte{
		byte((uint16(in) & (0xFF << 8)) >> 8),
		byte(in & 0xFF),
	}
}

func DecodeNetUint16(in [2]byte) uint16 {
	out := [2]uint16{
		uint16(in[0]),
		uint16(in[1]),
	}
	return uint16(out[1] | (out[0] << 8))
}

func EncodeNetUint16(in uint16) [2]byte {
	return [2]byte{
		byte((in & (0xFF << 8)) >> 8),
		byte(in & 0xFF),
	}
}

func DecodeNetInt32(in [4]byte) int32 {
	out := [4]int32{
		int32(in[0]),
		int32(in[1]),
		int32(in[2]),
		int32(in[3]),
	}
	return out[3] | (out[2] << 8) | (out[1] << 16) | (out[0] << 24)
}

func EncodeNetInt32(in int32) [4]byte {
	return [4]byte{
		byte((uint32(in) & (0xFF << 24)) >> 24),
		byte((in & (0xFF << 16)) >> 16),
		byte((in & (0xFF << 8)) >> 8),
		byte(in & 0xFF),
	}
}

func DecodeNetUint32(in [4]byte) uint32 {
	out := [4]uint32{
		uint32(in[0]),
		uint32(in[1]),
		uint32(in[2]),
		uint32(in[3]),
	}
	return out[3] | (out[2] << 8) | (out[1] << 16) | (out[0] << 24)
}

func EncodeNetUint32(in uint32) [4]byte {
	return [4]byte{
		byte((in & (0xFF << 24)) >> 24),
		byte((in & (0xFF << 16)) >> 16),
		byte((in & (0xFF << 8)) >> 8),
		byte(in & 0xFF),
	}
}

func DecodeNetInt64(in [8]byte) int64 {
	out := [8]int64{
		int64(in[0]),
		int64(in[1]),
		int64(in[2]),
		int64(in[3]),
		int64(in[4]),
		int64(in[5]),
		int64(in[6]),
		int64(in[7]),
	}
	return out[7] | (out[6] << 8) | (out[5] << 16) | (out[4] << 24) |
		(out[3] << 32) | (out[2] << 40) | (out[1] << 48) | (out[0] << 56)
}

func EncodeNetInt64(in int64) [8]byte {
	return [8]byte{
		byte((uint64(in) & (0xFF << 56)) >> 56),
		byte((in & (0xFF << 48)) >> 48),
		byte((in & (0xFF << 40)) >> 40),
		byte((in & (0xFF << 32)) >> 32),
		byte((in & (0xFF << 24)) >> 24),
		byte((in & (0xFF << 16)) >> 16),
		byte((in & (0xFF << 8)) >> 8),
		byte(in & 0xFF),
	}
}

func DecodeNetUint64(in [8]byte) uint64 {
	out := [8]uint64{
		uint64(in[0]),
		uint64(in[1]),
		uint64(in[2]),
		uint64(in[3]),
		uint64(in[4]),
		uint64(in[5]),
		uint64(in[6]),
		uint64(in[7]),
	}
	return out[7] | (out[6] << 8) | (out[5] << 16) | (out[4] << 24) |
		(out[3] << 32) | (out[2] << 40) | (out[1] << 48) | (out[0] << 56)
}

func encodeNetUint64(in uint64) [8]byte {
	return [8]byte{
		byte((in & (0xFF << 56)) >> 56),
		byte((in & (0xFF << 48)) >> 48),
		byte((in & (0xFF << 40)) >> 40),
		byte((in & (0xFF << 32)) >> 32),
		byte((in & (0xFF << 24)) >> 24),
		byte((in & (0xFF << 16)) >> 16),
		byte((in & (0xFF << 8)) >> 8),
		byte(in & 0xFF),
	}
}
