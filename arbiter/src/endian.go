package main

func decodeNetInt16(in [2]byte) int16 {
	var out int16 = int16(in[0] << 8 | (in[1] & 0x7F))
	if in[0] >> 7 == 1 {
		return out * -1
	} else {
		return out
	}
}

func encodeNetInt16(in int16) [2]byte {
	return [2]byte{
		byte((in & (0x7F << 8)) >> 8),
		byte(in & 0xFF),
	}
}

func decodeNetUint16(in [2]byte) uint16 {
	return uint16(in[0] | (in[1] << 8))
}

func encodeNetUint16(in uint16) [2]byte {
	return [2]byte{
		byte((in & (0xFF << 8)) >> 8),
		byte(in & 0xFF),
	}
}

func decodeNetInt32(in [4]byte) int32 {
	var out = int32(in[0] & 0x7F | (in[1] << 8) | (in[2] << 16) |
		(in[3] << 24))
	if in[0] >> 7 == 1 {
		return out * -1
	} else {
		return out
	}
}

func encodeNetInt32(in int32) [4]byte {
	return [4]byte{
		byte((in & (0x7F << 24)) >> 24),
		byte((in & (0xFF << 16)) >> 16),
		byte((in & (0xFF << 8)) >> 8),
		byte(in & 0xFF),
	}
}

func decodeNetUint32(in [4]byte) uint32 {
	return uint32(in[0] | (in[1] << 8) | (in[2] << 16) | (in[3] << 24))
}

func encodeNetUint32(in uint32) [4]byte {
	return [4]byte{
		byte((in & (0xFF << 24)) >> 24),
		byte((in & (0xFF << 16)) >> 16),
		byte((in & (0xFF << 8)) >> 8),
		byte(in & 0xFF),
	}
}

func decodeNetInt64(in [8]byte) int64 {
	var out = int64(in[0] & 0x7F | (in[1] << 8) | (in[2] << 16) |
		(in[3] << 24) | (in[4] << 32) | (in[5] << 40) | (in[6] << 48) |
		(in[7] << 56))
	if in[0] >> 7 == 1 {
		return out * -1
	} else {
		return out
	}
}

func encodeNetInt64(in int64) [8]byte {
	return [8]byte{
		byte((in & (0x7F << 56)) >> 56),
		byte((in & (0xFF << 48)) >> 48),
		byte((in & (0xFF << 40)) >> 40),
		byte((in & (0xFF << 32)) >> 32),
		byte((in & (0xFF << 24)) >> 24),
		byte((in & (0xFF << 16)) >> 16),
		byte((in & (0xFF << 8)) >> 8),
		byte(in & 0xFF),
	}
}

func decodeNetUint64(in [8]byte) uint64 {
	return uint64(in[0] | (in[1] << 8) | (in[2] << 16) | (in[3] << 24) |
		(in[4] << 32) | (in[5] << 40) | (in[6] << 48) | (in[7] << 56))
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
